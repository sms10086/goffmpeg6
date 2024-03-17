// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * filter layer
 * Copyright (c) 2007 Bobby Bingham
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

package goffmpeg6

//#cgo pkg-config: libavutil libavfilter
//#include <stddef.h>
//#include "libavutil/attributes.h"
//#include "libavutil/avutil.h"
//#include "libavutil/buffer.h"
//#include "libavutil/dict.h"
//#include "libavutil/frame.h"
//#include "libavutil/log.h"
//#include "libavutil/samplefmt.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavfilter/version_major.h"
//#include "libavfilter/version.h"
//#include "libavfilter/avfilter.h"
import "C"
import (
    "unsafe"
)

const AVFILTER_FLAG_DYNAMIC_INPUTS =         (1 << 0) 
const AVFILTER_FLAG_DYNAMIC_OUTPUTS =        (1 << 1) 
const AVFILTER_FLAG_SLICE_THREADS =          (1 << 2) 
const AVFILTER_FLAG_METADATA_ONLY =          (1 << 3) 
const AVFILTER_FLAG_HWDEVICE =               (1 << 4) 
const AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC =   (1 << 16) 
const AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL =  (1 << 17) 
const AVFILTER_FLAG_SUPPORT_TIMELINE =  (AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC | AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL) 
const AVFILTER_THREAD_SLICE =  (1 << 0) 
const AVFILTER_CMD_FLAG_ONE = 1
const AVFILTER_CMD_FLAG_FAST = 2


                           
                           

/**
 * @file
 * @ingroup lavfi
 * Main libavfilter public API header
 */

/**
 * @defgroup lavfi libavfilter
 * Graph-based frame editing library.
 *
 * @{
 */

                   

                                 
                             
                             
                           
                            
                          
                                
                             
                               

                                      
                        
/* When included as part of the ffmpeg build, only include the major version
 * to avoid unnecessary rebuilds. When included externally, keep including
 * the full version information. */
                                
      

/**
 * Return the LIBAVFILTER_VERSION_INT constant.
 */
func Avfilter_version() uint32 {
    return uint32(C.avfilter_version())
}

/**
 * Return the libavfilter build-time configuration.
 */
func Avfilter_configuration() string {
    return C.GoString(C.avfilter_configuration())
}

/**
 * Return the libavfilter license.
 */
func Avfilter_license() string {
    return C.GoString(C.avfilter_license())
}

// type AVFilterContext
// type AVFilterLink
type AVFilterPad struct {
}

type AVFilterFormats struct {
}

type AVFilterChannelLayouts struct {
}


/**
 * Get the name of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return name of the pad_idx'th pad in pads
 */
func Avfilter_pad_get_name(pads *AVFilterPad, pad_idx int32) string {
    return C.GoString(C.avfilter_pad_get_name(
        (*C.struct_AVFilterPad)(unsafe.Pointer(pads)), C.int(pad_idx)))
}

/**
 * Get the type of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return type of the pad_idx'th pad in pads
 */
func Avfilter_pad_get_type(pads *AVFilterPad, pad_idx int32) AVMediaType {
    return AVMediaType(C.avfilter_pad_get_type(
        (*C.struct_AVFilterPad)(unsafe.Pointer(pads)), C.int(pad_idx)))
}

/**
 * The number of the filter inputs is not determined just by AVFilter.inputs.
 * The filter might add additional inputs during initialization depending on the
 * options supplied to it.
 */

/**
 * The number of the filter outputs is not determined just by AVFilter.outputs.
 * The filter might add additional outputs during initialization depending on
 * the options supplied to it.
 */

/**
 * The filter supports multithreading by splitting frames into multiple parts
 * and processing them concurrently.
 */

/**
 * The filter is a "metadata" filter - it does not modify the frame data in any
 * way. It may only affect the metadata (i.e. those fields copied by
 * av_frame_copy_props()).
 *
 * More precisely, this means:
 * - video: the data of any frame output by the filter must be exactly equal to
 *   some frame that is received on one of its inputs. Furthermore, all frames
 *   produced on a given output must correspond to frames received on the same
 *   input and their order must be unchanged. Note that the filter may still
 *   drop or duplicate the frames.
 * - audio: the data produced by the filter on any of its outputs (viewed e.g.
 *   as an array of interleaved samples) must be exactly equal to the data
 *   received by the filter on one of its inputs.
 */


/**
 * The filter can create hardware frames using AVFilterContext.hw_device_ctx.
 */

/**
 * Some filters support a generic "enable" expression option that can be used
 * to enable or disable a filter in the timeline. Filters supporting this
 * option have this flag set. When the enable expression is false, the default
 * no-op filter_frame() function is called in place of the filter_frame()
 * callback defined on each input pad, thus the frame is passed unchanged to
 * the next filters.
 */

/**
 * Same as AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC, except that the filter will
 * have its filter_frame() callback(s) called as usual even when the enable
 * expression is false. The filter will disable filtering within the
 * filter_frame() callback(s) itself, for example executing code depending on
 * the AVFilterContext->is_disabled value.
 */

/**
 * Handy mask to test whether the filter supports or no the timeline feature
 * (internally or generically).
 */


/**
 * Filter definition. This defines the pads a filter contains, and all the
 * callback functions used to interact with the filter.
 */
type AVFilter struct {
    Name *byte
    Description *byte
    Inputs *AVFilterPad
    Outputs *AVFilterPad
    Priv_class *AVClass
    Flags int32
    Nb_inputs uint8
    Nb_outputs uint8
    Formats_state uint8
    Preinit uintptr
    Init uintptr
    Uninit uintptr
    Formats uintptr
    Priv_size int32
    Flags_internal int32
    Process_command uintptr
    Activate uintptr
}


/**
 * Get the number of elements in an AVFilter's inputs or outputs array.
 */
func Avfilter_filter_pad_count(filter *AVFilter, is_output int32) uint32 {
    return uint32(C.avfilter_filter_pad_count(
        (*C.struct_AVFilter)(unsafe.Pointer(filter)), C.int(is_output)))
}

/**
 * Process multiple parts of the frame concurrently.
 */


type AVFilterInternal struct {
}


/** An instance of a filter */
type AVFilterContext struct {
    Av_class *AVClass
    Filter *AVFilter
    Name *byte
    Input_pads *AVFilterPad
    Inputs **AVFilterLink
    Nb_inputs uint32
    Output_pads *AVFilterPad
    Outputs **AVFilterLink
    Nb_outputs uint32
    Priv unsafe.Pointer
    Graph *AVFilterGraph
    Thread_type int32
    Internal *AVFilterInternal
    Command_queue *AVFilterCommand
    Enable_str *byte
    Enable unsafe.Pointer
    Var_values *float64
    Is_disabled int32
    Hw_device_ctx *AVBufferRef
    Nb_threads int32
    Ready uint32
    Extra_hw_frames int32
}


/**
 * Lists of formats / etc. supported by an end of a link.
 *
 * This structure is directly part of AVFilterLink, in two copies:
 * one for the source filter, one for the destination filter.

 * These lists are used for negotiating the format to actually be used,
 * which will be loaded into the format and channel_layout members of
 * AVFilterLink, when chosen.
 */
type AVFilterFormatsConfig struct {
    Formats *AVFilterFormats
    Samplerates *AVFilterFormats
    Channel_layouts *AVFilterChannelLayouts
}


/**
 * A link between two filters. This contains pointers to the source and
 * destination filters between which this link exists, and the indexes of
 * the pads involved. In addition, this link also contains the parameters
 * which have been negotiated and agreed upon between the filter, such as
 * image dimensions, format, etc.
 *
 * Applications must not normally access the link structure directly.
 * Use the buffersrc and buffersink API instead.
 * In the future, access to the header may be reserved for filters
 * implementation.
 */
type AVFilterLink struct {
    Src *AVFilterContext
    Srcpad *AVFilterPad
    Dst *AVFilterContext
    Dstpad *AVFilterPad
    Type AVMediaType
    W int32
    H int32
    Sample_aspect_ratio AVRational
    Channel_layout uint64
    Sample_rate int32
    Format int32
    Time_base AVRational
    Ch_layout AVChannelLayout
    Incfg AVFilterFormatsConfig
    Outcfg AVFilterFormatsConfig
    Init_state int32
    Graph *AVFilterGraph
    Current_pts int64
    Current_pts_us int64
    Age_index int32
    Frame_rate AVRational
    Min_samples int32
    Max_samples int32
    Frame_count_in int64
    Frame_count_out int64
    Sample_count_in int64
    Sample_count_out int64
    Frame_pool unsafe.Pointer
    Frame_wanted_out int32
    Hw_frames_ctx *AVBufferRef
    Reserved [0xF000]byte
}


/**
 * Link two filters together.
 *
 * @param src    the source filter
 * @param srcpad index of the output pad on the source filter
 * @param dst    the destination filter
 * @param dstpad index of the input pad on the destination filter
 * @return       zero on success
 */
func Avfilter_link(src *AVFilterContext, srcpad uint32,
                  dst *AVFilterContext, dstpad uint32) int32 {
    return int32(C.avfilter_link((*C.struct_AVFilterContext)(unsafe.Pointer(src)), 
        C.unsigned(srcpad), (*C.struct_AVFilterContext)(unsafe.Pointer(dst)), 
        C.unsigned(dstpad)))
}

/**
 * Free the link in *link, and set its pointer to NULL.
 */
func Avfilter_link_free(link **AVFilterLink)  {
    C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(link)))
}

/**
 * Negotiate the media format, dimensions, etc of all inputs to a filter.
 *
 * @param filter the filter to negotiate the properties for its inputs
 * @return       zero on successful negotiation
 */
func Avfilter_config_links(filter *AVFilterContext) int32 {
    return int32(C.avfilter_config_links(
        (*C.struct_AVFilterContext)(unsafe.Pointer(filter))))
}

///< Stop once a filter understood the command (for target=all for example), fast filters are favored automatically
///< Only execute command when its fast (like a video out that supports contrast adjustment in hw)

/**
 * Make the filter instance process a command.
 * It is recommended to use avfilter_graph_send_command().
 */
func Avfilter_process_command(filter *AVFilterContext, cmd *byte, arg *byte, res *byte, res_len int32, flags int32) int32 {
    return int32(C.avfilter_process_command(
        (*C.struct_AVFilterContext)(unsafe.Pointer(filter)), 
        (*C.char)(unsafe.Pointer(cmd)), (*C.char)(unsafe.Pointer(arg)), 
        (*C.char)(unsafe.Pointer(res)), C.int(res_len), C.int(flags)))
}

/**
 * Iterate over all registered filters.
 *
 * @param opaque a pointer where libavfilter will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered filter or NULL when the iteration is
 *         finished
 */
func Av_filter_iterate(opaque *unsafe.Pointer) *AVFilter {
    return (*AVFilter)(unsafe.Pointer(C.av_filter_iterate(opaque)))
}

/**
 * Get a filter definition matching the given name.
 *
 * @param name the filter name to find
 * @return     the filter definition, if any matching one is registered.
 *             NULL if none found.
 */
func Avfilter_get_by_name(name *byte) *AVFilter {
    return (*AVFilter)(unsafe.Pointer(C.avfilter_get_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}


/**
 * Initialize a filter with the supplied parameters.
 *
 * @param ctx  uninitialized filter context to initialize
 * @param args Options to initialize the filter with. This must be a
 *             ':'-separated list of options in the 'key=value' form.
 *             May be NULL if the options have been set directly using the
 *             AVOptions API or there are no options that need to be set.
 * @return 0 on success, a negative AVERROR on failure
 */
func Avfilter_init_str(ctx *AVFilterContext, args *byte) int32 {
    return int32(C.avfilter_init_str(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.char)(unsafe.Pointer(args))))
}

/**
 * Initialize a filter with the supplied dictionary of options.
 *
 * @param ctx     uninitialized filter context to initialize
 * @param options An AVDictionary filled with options for this filter. On
 *                return this parameter will be destroyed and replaced with
 *                a dict containing options that were not found. This dictionary
 *                must be freed by the caller.
 *                May be NULL, then this function is equivalent to
 *                avfilter_init_str() with the second parameter set to NULL.
 * @return 0 on success, a negative AVERROR on failure
 *
 * @note This function and avfilter_init_str() do essentially the same thing,
 * the difference is in manner in which the options are passed. It is up to the
 * calling code to choose whichever is more preferable. The two functions also
 * behave differently when some of the provided options are not declared as
 * supported by the filter. In such a case, avfilter_init_str() will fail, but
 * this function will leave those extra options in the options AVDictionary and
 * continue as usual.
 */
func Avfilter_init_dict(ctx *AVFilterContext, options **AVDictionary) int32 {
    return int32(C.avfilter_init_dict(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Free a filter context. This will also remove the filter from its
 * filtergraph's list of filters.
 *
 * @param filter the filter to free
 */
func Avfilter_free(filter *AVFilterContext)  {
    C.avfilter_free((*C.struct_AVFilterContext)(unsafe.Pointer(filter)))
}

/**
 * Insert a filter in the middle of an existing link.
 *
 * @param link the link into which the filter should be inserted
 * @param filt the filter to be inserted
 * @param filt_srcpad_idx the input pad on the filter to connect
 * @param filt_dstpad_idx the output pad on the filter to connect
 * @return     zero on success
 */
func Avfilter_insert_filter(link *AVFilterLink, filt *AVFilterContext,
                           filt_srcpad_idx uint32, filt_dstpad_idx uint32) int32 {
    return int32(C.avfilter_insert_filter(
        (*C.struct_AVFilterLink)(unsafe.Pointer(link)), 
        (*C.struct_AVFilterContext)(unsafe.Pointer(filt)), 
        C.unsigned(filt_srcpad_idx), C.unsigned(filt_dstpad_idx)))
}

/**
 * @return AVClass for AVFilterContext.
 *
 * @see av_opt_find().
 */
func Avfilter_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avfilter_get_class()))
}

type AVFilterGraphInternal struct {
}


/**
 * A function pointer passed to the @ref AVFilterGraph.execute callback to be
 * executed multiple times, possibly in parallel.
 *
 * @param ctx the filter context the job belongs to
 * @param arg an opaque parameter passed through from @ref
 *            AVFilterGraph.execute
 * @param jobnr the index of the job being executed
 * @param nb_jobs the total number of jobs
 *
 * @return 0 on success, a negative AVERROR on error
 */
type avfilter_action_func C.avfilter_action_func

/**
 * A function executing multiple jobs, possibly in parallel.
 *
 * @param ctx the filter context to which the jobs belong
 * @param func the function to be called multiple times
 * @param arg the argument to be passed to func
 * @param ret a nb_jobs-sized array to be filled with return values from each
 *            invocation of func
 * @param nb_jobs the number of jobs to execute
 *
 * @return 0 on success, a negative AVERROR on error
 */
type avfilter_execute_func C.avfilter_execute_func

type AVFilterGraph struct {
    Av_class *AVClass
    Filters **AVFilterContext
    Nb_filters uint32
    Scale_sws_opts *byte
    Thread_type int32
    Nb_threads int32
    Internal *AVFilterGraphInternal
    Opaque unsafe.Pointer
    Execute *avfilter_execute_func
    Aresample_swr_opts *byte
    Sink_links **AVFilterLink
    Sink_links_count int32
    Disable_auto_convert uint32
}


/**
 * Allocate a filter graph.
 *
 * @return the allocated filter graph on success or NULL.
 */
func Avfilter_graph_alloc() *AVFilterGraph {
    return (*AVFilterGraph)(unsafe.Pointer(C.avfilter_graph_alloc()))
}

/**
 * Create a new filter instance in a filter graph.
 *
 * @param graph graph in which the new filter will be used
 * @param filter the filter to create an instance of
 * @param name Name to give to the new instance (will be copied to
 *             AVFilterContext.name). This may be used by the caller to identify
 *             different filters, libavfilter itself assigns no semantics to
 *             this parameter. May be NULL.
 *
 * @return the context of the newly created filter instance (note that it is
 *         also retrievable directly through AVFilterGraph.filters or with
 *         avfilter_graph_get_filter()) on success or NULL on failure.
 */
func Avfilter_graph_alloc_filter(graph *AVFilterGraph,
                                             filter *AVFilter,
                                             name *byte) *AVFilterContext {
    return (*AVFilterContext)(unsafe.Pointer(C.avfilter_graph_alloc_filter(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.struct_AVFilter)(unsafe.Pointer(filter)), 
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Get a filter instance identified by instance name from graph.
 *
 * @param graph filter graph to search through.
 * @param name filter instance name (should be unique in the graph).
 * @return the pointer to the found filter instance or NULL if it
 * cannot be found.
 */
func Avfilter_graph_get_filter(graph *AVFilterGraph, name *byte) *AVFilterContext {
    return (*AVFilterContext)(unsafe.Pointer(C.avfilter_graph_get_filter(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Create and add a filter instance into an existing graph.
 * The filter instance is created from the filter filt and inited
 * with the parameter args. opaque is currently ignored.
 *
 * In case of success put in *filt_ctx the pointer to the created
 * filter instance, otherwise set *filt_ctx to NULL.
 *
 * @param name the instance name to give to the created filter instance
 * @param graph_ctx the filter graph
 * @return a negative AVERROR error code in case of failure, a non
 * negative value otherwise
 */
func Avfilter_graph_create_filter(filt_ctx **AVFilterContext, filt *AVFilter,
                                 name *byte, args *byte, opaque unsafe.Pointer,
                                 graph_ctx *AVFilterGraph) int32 {
    return int32(C.avfilter_graph_create_filter(
        (**C.struct_AVFilterContext)(unsafe.Pointer(filt_ctx)), 
        (*C.struct_AVFilter)(unsafe.Pointer(filt)), 
        (*C.char)(unsafe.Pointer(name)), (*C.char)(unsafe.Pointer(args)), opaque, 
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph_ctx))))
}

/**
 * Enable or disable automatic format conversion inside the graph.
 *
 * Note that format conversion can still happen inside explicitly inserted
 * scale and aresample filters.
 *
 * @param flags  any of the AVFILTER_AUTO_CONVERT_* constants
 */
func Avfilter_graph_set_auto_convert(graph *AVFilterGraph, flags uint32)  {
    C.avfilter_graph_set_auto_convert(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), C.unsigned(flags))
}

const (
    AVFILTER_AUTO_CONVERT_ALL  = 0 + iota
    AVFILTER_AUTO_CONVERT_NONE = -1
)


/**
 * Check validity and configure all the links and formats in the graph.
 *
 * @param graphctx the filter graph
 * @param log_ctx context used for logging
 * @return >= 0 in case of success, a negative AVERROR code otherwise
 */
func Avfilter_graph_config(graphctx *AVFilterGraph, log_ctx unsafe.Pointer) int32 {
    return int32(C.avfilter_graph_config(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graphctx)), log_ctx))
}

/**
 * Free a graph, destroy its links, and set *graph to NULL.
 * If *graph is NULL, do nothing.
 */
func Avfilter_graph_free(graph **AVFilterGraph)  {
    C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(graph)))
}

/**
 * A linked-list of the inputs/outputs of the filter chain.
 *
 * This is mainly useful for avfilter_graph_parse() / avfilter_graph_parse2(),
 * where it is used to communicate open (unlinked) inputs and outputs from and
 * to the caller.
 * This struct specifies, per each not connected pad contained in the graph, the
 * filter context and the pad index required for establishing a link.
 */
type AVFilterInOut struct {
    Name *byte
    Filter_ctx *AVFilterContext
    Pad_idx int32
    Next *AVFilterInOut
}


/**
 * Allocate a single AVFilterInOut entry.
 * Must be freed with avfilter_inout_free().
 * @return allocated AVFilterInOut on success, NULL on failure.
 */
func Avfilter_inout_alloc() *AVFilterInOut {
    return (*AVFilterInOut)(unsafe.Pointer(C.avfilter_inout_alloc()))
}

/**
 * Free the supplied list of AVFilterInOut and set *inout to NULL.
 * If *inout is NULL, do nothing.
 */
func Avfilter_inout_free(inout **AVFilterInOut)  {
    C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(inout)))
}

/**
 * Add a graph described by a string to a graph.
 *
 * @note The caller must provide the lists of inputs and outputs,
 * which therefore must be known before calling the function.
 *
 * @note The inputs parameter describes inputs of the already existing
 * part of the graph; i.e. from the point of view of the newly created
 * part, they are outputs. Similarly the outputs parameter describes
 * outputs of the already existing filters, which are provided as
 * inputs to the parsed filters.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  linked list to the inputs of the graph
 * @param outputs linked list to the outputs of the graph
 * @return zero on success, a negative AVERROR code on error
 */
func Avfilter_graph_parse(graph *AVFilterGraph, filters *byte,
                         inputs *AVFilterInOut, outputs *AVFilterInOut,
                         log_ctx unsafe.Pointer) int32 {
    return int32(C.avfilter_graph_parse(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(filters)), 
        (*C.struct_AVFilterInOut)(unsafe.Pointer(inputs)), 
        (*C.struct_AVFilterInOut)(unsafe.Pointer(outputs)), log_ctx))
}

/**
 * Add a graph described by a string to a graph.
 *
 * In the graph filters description, if the input label of the first
 * filter is not specified, "in" is assumed; if the output label of
 * the last filter is not specified, "out" is assumed.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  pointer to a linked list to the inputs of the graph, may be NULL.
 *                If non-NULL, *inputs is updated to contain the list of open inputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @param outputs pointer to a linked list to the outputs of the graph, may be NULL.
 *                If non-NULL, *outputs is updated to contain the list of open outputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @return non negative on success, a negative AVERROR code on error
 */
func Avfilter_graph_parse_ptr(graph *AVFilterGraph, filters *byte,
                             inputs **AVFilterInOut, outputs **AVFilterInOut,
                             log_ctx unsafe.Pointer) int32 {
    return int32(C.avfilter_graph_parse_ptr(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(filters)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(outputs)), log_ctx))
}

/**
 * Add a graph described by a string to a graph.
 *
 * @param[in]  graph   the filter graph where to link the parsed graph context
 * @param[in]  filters string to be parsed
 * @param[out] inputs  a linked list of all free (unlinked) inputs of the
 *                     parsed graph will be returned here. It is to be freed
 *                     by the caller using avfilter_inout_free().
 * @param[out] outputs a linked list of all free (unlinked) outputs of the
 *                     parsed graph will be returned here. It is to be freed by the
 *                     caller using avfilter_inout_free().
 * @return zero on success, a negative AVERROR code on error
 *
 * @note This function returns the inputs and outputs that are left
 * unlinked after parsing the graph and the caller then deals with
 * them.
 * @note This function makes no reference whatsoever to already
 * existing parts of the graph and the inputs parameter will on return
 * contain inputs of the newly parsed part of the graph.  Analogously
 * the outputs parameter will contain outputs of the newly created
 * filters.
 */
func Avfilter_graph_parse2(graph *AVFilterGraph, filters *byte,
                          inputs **AVFilterInOut,
                          outputs **AVFilterInOut) int32 {
    return int32(C.avfilter_graph_parse2(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(filters)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

/**
 * Parameters of a filter's input or output pad.
 *
 * Created as a child of AVFilterParams by avfilter_graph_segment_parse().
 * Freed in avfilter_graph_segment_free().
 */
type AVFilterPadParams struct {
    Label *byte
}


/**
 * Parameters describing a filter to be created in a filtergraph.
 *
 * Created as a child of AVFilterGraphSegment by avfilter_graph_segment_parse().
 * Freed in avfilter_graph_segment_free().
 */
type AVFilterParams struct {
    Filter *AVFilterContext
    Filter_name *byte
    Instance_name *byte
    Opts *AVDictionary
    Inputs **AVFilterPadParams
    Nb_inputs uint32
    Outputs **AVFilterPadParams
    Nb_outputs uint32
}


/**
 * A filterchain is a list of filter specifications.
 *
 * Created as a child of AVFilterGraphSegment by avfilter_graph_segment_parse().
 * Freed in avfilter_graph_segment_free().
 */
type AVFilterChain struct {
    Filters **AVFilterParams
    Nb_filters uint64
}


/**
 * A parsed representation of a filtergraph segment.
 *
 * A filtergraph segment is conceptually a list of filterchains, with some
 * supplementary information (e.g. format conversion flags).
 *
 * Created by avfilter_graph_segment_parse(). Must be freed with
 * avfilter_graph_segment_free().
 */
type AVFilterGraphSegment struct {
    Graph *AVFilterGraph
    Chains **AVFilterChain
    Nb_chains uint64
    Scale_sws_opts *byte
}


/**
 * Parse a textual filtergraph description into an intermediate form.
 *
 * This intermediate representation is intended to be modified by the caller as
 * described in the documentation of AVFilterGraphSegment and its children, and
 * then applied to the graph either manually or with other
 * avfilter_graph_segment_*() functions. See the documentation for
 * avfilter_graph_segment_apply() for the canonical way to apply
 * AVFilterGraphSegment.
 *
 * @param graph Filter graph the parsed segment is associated with. Will only be
 *              used for logging and similar auxiliary purposes. The graph will
 *              not be actually modified by this function - the parsing results
 *              are instead stored in seg for further processing.
 * @param graph_str a string describing the filtergraph segment
 * @param flags reserved for future use, caller must set to 0 for now
 * @param seg A pointer to the newly-created AVFilterGraphSegment is written
 *            here on success. The graph segment is owned by the caller and must
 *            be freed with avfilter_graph_segment_free() before graph itself is
 *            freed.
 *
 * @retval "non-negative number" success
 * @retval "negative error code" failure
 */
func Avfilter_graph_segment_parse(graph *AVFilterGraph, graph_str *byte,
                                 flags int32, seg **AVFilterGraphSegment) int32 {
    return int32(C.avfilter_graph_segment_parse(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(graph_str)), C.int(flags), 
        (**C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg))))
}

/**
 * Create filters specified in a graph segment.
 *
 * Walk through the creation-pending AVFilterParams in the segment and create
 * new filter instances for them.
 * Creation-pending params are those where AVFilterParams.filter_name is
 * non-NULL (and hence AVFilterParams.filter is NULL). All other AVFilterParams
 * instances are ignored.
 *
 * For any filter created by this function, the corresponding
 * AVFilterParams.filter is set to the newly-created filter context,
 * AVFilterParams.filter_name and AVFilterParams.instance_name are freed and set
 * to NULL.
 *
 * @param seg the filtergraph segment to process
 * @param flags reserved for future use, caller must set to 0 for now
 *
 * @retval "non-negative number" Success, all creation-pending filters were
 *                               successfully created
 * @retval AVERROR_FILTER_NOT_FOUND some filter's name did not correspond to a
 *                                  known filter
 * @retval "another negative error code" other failures
 *
 * @note Calling this function multiple times is safe, as it is idempotent.
 */
func Avfilter_graph_segment_create_filters(seg *AVFilterGraphSegment, flags int32) int32 {
    return int32(C.avfilter_graph_segment_create_filters(
        (*C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)), C.int(flags)))
}

/**
 * Apply parsed options to filter instances in a graph segment.
 *
 * Walk through all filter instances in the graph segment that have option
 * dictionaries associated with them and apply those options with
 * av_opt_set_dict2(..., AV_OPT_SEARCH_CHILDREN). AVFilterParams.opts is
 * replaced by the dictionary output by av_opt_set_dict2(), which should be
 * empty (NULL) if all options were successfully applied.
 *
 * If any options could not be found, this function will continue processing all
 * other filters and finally return AVERROR_OPTION_NOT_FOUND (unless another
 * error happens). The calling program may then deal with unapplied options as
 * it wishes.
 *
 * Any creation-pending filters (see avfilter_graph_segment_create_filters())
 * present in the segment will cause this function to fail. AVFilterParams with
 * no associated filter context are simply skipped.
 *
 * @param seg the filtergraph segment to process
 * @param flags reserved for future use, caller must set to 0 for now
 *
 * @retval "non-negative number" Success, all options were successfully applied.
 * @retval AVERROR_OPTION_NOT_FOUND some options were not found in a filter
 * @retval "another negative error code" other failures
 *
 * @note Calling this function multiple times is safe, as it is idempotent.
 */
func Avfilter_graph_segment_apply_opts(seg *AVFilterGraphSegment, flags int32) int32 {
    return int32(C.avfilter_graph_segment_apply_opts(
        (*C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)), C.int(flags)))
}

/**
 * Initialize all filter instances in a graph segment.
 *
 * Walk through all filter instances in the graph segment and call
 * avfilter_init_dict(..., NULL) on those that have not been initialized yet.
 *
 * Any creation-pending filters (see avfilter_graph_segment_create_filters())
 * present in the segment will cause this function to fail. AVFilterParams with
 * no associated filter context or whose filter context is already initialized,
 * are simply skipped.
 *
 * @param seg the filtergraph segment to process
 * @param flags reserved for future use, caller must set to 0 for now
 *
 * @retval "non-negative number" Success, all filter instances were successfully
 *                               initialized
 * @retval "negative error code" failure
 *
 * @note Calling this function multiple times is safe, as it is idempotent.
 */
func Avfilter_graph_segment_init(seg *AVFilterGraphSegment, flags int32) int32 {
    return int32(C.avfilter_graph_segment_init(
        (*C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)), C.int(flags)))
}

/**
 * Link filters in a graph segment.
 *
 * Walk through all filter instances in the graph segment and try to link all
 * unlinked input and output pads. Any creation-pending filters (see
 * avfilter_graph_segment_create_filters()) present in the segment will cause
 * this function to fail. Disabled filters and already linked pads are skipped.
 *
 * Every filter output pad that has a corresponding AVFilterPadParams with a
 * non-NULL label is
 * - linked to the input with the matching label, if one exists;
 * - exported in the outputs linked list otherwise, with the label preserved.
 * Unlabeled outputs are
 * - linked to the first unlinked unlabeled input in the next non-disabled
 *   filter in the chain, if one exists
 * - exported in the ouputs linked list otherwise, with NULL label
 *
 * Similarly, unlinked input pads are exported in the inputs linked list.
 *
 * @param seg the filtergraph segment to process
 * @param flags reserved for future use, caller must set to 0 for now
 * @param[out] inputs  a linked list of all free (unlinked) inputs of the
 *                     filters in this graph segment will be returned here. It
 *                     is to be freed by the caller using avfilter_inout_free().
 * @param[out] outputs a linked list of all free (unlinked) outputs of the
 *                     filters in this graph segment will be returned here. It
 *                     is to be freed by the caller using avfilter_inout_free().
 *
 * @retval "non-negative number" success
 * @retval "negative error code" failure
 *
 * @note Calling this function multiple times is safe, as it is idempotent.
 */
func Avfilter_graph_segment_link(seg *AVFilterGraphSegment, flags int32,
                                inputs **AVFilterInOut,
                                outputs **AVFilterInOut) int32 {
    return int32(C.avfilter_graph_segment_link(
        (*C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)), C.int(flags), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

/**
 * Apply all filter/link descriptions from a graph segment to the associated filtergraph.
 *
 * This functions is currently equivalent to calling the following in sequence:
 * - avfilter_graph_segment_create_filters();
 * - avfilter_graph_segment_apply_opts();
 * - avfilter_graph_segment_init();
 * - avfilter_graph_segment_link();
 * failing if any of them fails. This list may be extended in the future.
 *
 * Since the above functions are idempotent, the caller may call some of them
 * manually, then do some custom processing on the filtergraph, then call this
 * function to do the rest.
 *
 * @param seg the filtergraph segment to process
 * @param flags reserved for future use, caller must set to 0 for now
 * @param[out] inputs passed to avfilter_graph_segment_link()
 * @param[out] outputs passed to avfilter_graph_segment_link()
 *
 * @retval "non-negative number" success
 * @retval "negative error code" failure
 *
 * @note Calling this function multiple times is safe, as it is idempotent.
 */
func Avfilter_graph_segment_apply(seg *AVFilterGraphSegment, flags int32,
                                 inputs **AVFilterInOut,
                                 outputs **AVFilterInOut) int32 {
    return int32(C.avfilter_graph_segment_apply(
        (*C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)), C.int(flags), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)), 
        (**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

/**
 * Free the provided AVFilterGraphSegment and everything associated with it.
 *
 * @param seg double pointer to the AVFilterGraphSegment to be freed. NULL will
 * be written to this pointer on exit from this function.
 *
 * @note
 * The filter contexts (AVFilterParams.filter) are owned by AVFilterGraph rather
 * than AVFilterGraphSegment, so they are not freed.
 */
func Avfilter_graph_segment_free(seg **AVFilterGraphSegment)  {
    C.avfilter_graph_segment_free(
        (**C.struct_AVFilterGraphSegment)(unsafe.Pointer(seg)))
}

/**
 * Send a command to one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to send, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param res    a buffer with size res_size where the filter(s) can return a response.
 *
 * @returns >=0 on success otherwise an error code.
 *              AVERROR(ENOSYS) on unsupported commands
 */
func Avfilter_graph_send_command(graph *AVFilterGraph, target *byte, cmd *byte, arg *byte, res *byte, res_len int32, flags int32) int32 {
    return int32(C.avfilter_graph_send_command(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(target)), (*C.char)(unsafe.Pointer(cmd)), 
        (*C.char)(unsafe.Pointer(arg)), (*C.char)(unsafe.Pointer(res)), 
        C.int(res_len), C.int(flags)))
}

/**
 * Queue a command for one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to sent, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param ts     time at which the command should be sent to the filter
 *
 * @note As this executes commands after this function returns, no return code
 *       from the filter is provided, also AVFILTER_CMD_FLAG_ONE is not supported.
 */
func Avfilter_graph_queue_command(graph *AVFilterGraph, target *byte, cmd *byte, arg *byte, flags int32, ts float64) int32 {
    return int32(C.avfilter_graph_queue_command(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(target)), (*C.char)(unsafe.Pointer(cmd)), 
        (*C.char)(unsafe.Pointer(arg)), C.int(flags), C.double(ts)))
}


/**
 * Dump a graph into a human-readable string representation.
 *
 * @param graph    the graph to dump
 * @param options  formatting options; currently ignored
 * @return  a string, or NULL in case of memory allocation failure;
 *          the string must be freed using av_free
 */
func Avfilter_graph_dump(graph *AVFilterGraph, options *byte) string {
    return C.GoString(C.avfilter_graph_dump(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph)), 
        (*C.char)(unsafe.Pointer(options))))
}

/**
 * Request a frame on the oldest sink link.
 *
 * If the request returns AVERROR_EOF, try the next.
 *
 * Note that this function is not meant to be the sole scheduling mechanism
 * of a filtergraph, only a convenience function to help drain a filtergraph
 * in a balanced way under normal circumstances.
 *
 * Also note that AVERROR_EOF does not mean that frames did not arrive on
 * some of the sinks during the process.
 * When there are multiple sink links, in case the requested link
 * returns an EOF, this may cause a filter to flush pending frames
 * which are sent to another sink link, although unrequested.
 *
 * @return  the return value of ff_request_frame(),
 *          or AVERROR_EOF if all links returned AVERROR_EOF
 */
func Avfilter_graph_request_oldest(graph *AVFilterGraph) int32 {
    return int32(C.avfilter_graph_request_oldest(
        (*C.struct_AVFilterGraph)(unsafe.Pointer(graph))))
}

/**
 * @}
 */

                                
