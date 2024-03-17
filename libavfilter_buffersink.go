// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
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

//#cgo pkg-config: libavfilter
//#include "libavfilter/avfilter.h"
//#include "libavfilter/buffersink.h"
import "C"
import (
    "unsafe"
)

const AV_BUFFERSINK_FLAG_PEEK = 1
const AV_BUFFERSINK_FLAG_NO_REQUEST = 2


                             
                             

/**
 * @file
 * @ingroup lavfi_buffersink
 * memory buffer sink API for audio and video
 */

                     

/**
 * @defgroup lavfi_buffersink Buffer sink API
 * @ingroup lavfi
 * @{
 *
 * The buffersink and abuffersink filters are there to connect filter graphs
 * to applications. They have a single input, connected to the graph, and no
 * output. Frames must be extracted using av_buffersink_get_frame() or
 * av_buffersink_get_samples().
 *
 * The format negotiated by the graph during configuration can be obtained
 * using the accessor functions:
 * - av_buffersink_get_time_base(),
 * - av_buffersink_get_format(),
 * - av_buffersink_get_frame_rate(),
 * - av_buffersink_get_w(),
 * - av_buffersink_get_h(),
 * - av_buffersink_get_sample_aspect_ratio(),
 * - av_buffersink_get_channels(),
 * - av_buffersink_get_ch_layout(),
 * - av_buffersink_get_sample_rate().
 *
 * The layout returned by av_buffersink_get_ch_layout() must de uninitialized
 * by the caller.
 *
 * The format can be constrained by setting options, using av_opt_set() and
 * related functions with the AV_OPT_SEARCH_CHILDREN flag.
 *  - pix_fmts (int list),
 *  - sample_fmts (int list),
 *  - sample_rates (int list),
 *  - ch_layouts (string),
 *  - channel_counts (int list),
 *  - all_channel_counts (bool).
 * Most of these options are of type binary, and should be set using
 * av_opt_set_int_list() or av_opt_set_bin(). If they are not set, all
 * corresponding formats are accepted.
 *
 * As a special case, if ch_layouts is not set, all valid channel layouts are
 * accepted except for UNSPEC layouts, unless all_channel_counts is set.
 */

/**
 * Get a frame with filtered data from sink and put it in frame.
 *
 * @param ctx    pointer to a buffersink or abuffersink filter context.
 * @param frame  pointer to an allocated frame that will be filled with data.
 *               The data must be freed using av_frame_unref() / av_frame_free()
 * @param flags  a combination of AV_BUFFERSINK_FLAG_* flags
 *
 * @return  >= 0 in for success, a negative AVERROR code for failure.
 */
func Av_buffersink_get_frame_flags(ctx *AVFilterContext, frame *AVFrame, flags int32) int32 {
    return int32(C.av_buffersink_get_frame_flags(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * Tell av_buffersink_get_buffer_ref() to read video/samples buffer
 * reference, but not remove it from the buffer. This is useful if you
 * need only to read a video/samples buffer, without to fetch it.
 */


/**
 * Tell av_buffersink_get_buffer_ref() not to request a frame from its input.
 * If a frame is already buffered, it is read (and removed from the buffer),
 * but if no frame is present, return AVERROR(EAGAIN).
 */


/**
 * Set the frame size for an audio buffer sink.
 *
 * All calls to av_buffersink_get_buffer_ref will return a buffer with
 * exactly the specified number of samples, or AVERROR(EAGAIN) if there is
 * not enough. The last buffer at EOF will be padded with 0.
 */
func Av_buffersink_set_frame_size(ctx *AVFilterContext, frame_size uint32)  {
    C.av_buffersink_set_frame_size((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        C.unsigned(frame_size))
}

/**
 * @defgroup lavfi_buffersink_accessors Buffer sink accessors
 * Get the properties of the stream
 * @{
 */

func Av_buffersink_get_type                (ctx *AVFilterContext) AVMediaType {
    return AVMediaType(C.av_buffersink_get_type(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}
func       Av_buffersink_get_time_base           (ctx *AVFilterContext) AVRational {
    _ret := C.av_buffersink_get_time_base(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}
func              Av_buffersink_get_format              (ctx *AVFilterContext) int32 {
    return int32(C.av_buffersink_get_format(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}

func       Av_buffersink_get_frame_rate          (ctx *AVFilterContext) AVRational {
    _ret := C.av_buffersink_get_frame_rate(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}
func              Av_buffersink_get_w                   (ctx *AVFilterContext) int32 {
    return int32(C.av_buffersink_get_w(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}
func              Av_buffersink_get_h                   (ctx *AVFilterContext) int32 {
    return int32(C.av_buffersink_get_h(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}
func       Av_buffersink_get_sample_aspect_ratio (ctx *AVFilterContext) AVRational {
    _ret := C.av_buffersink_get_sample_aspect_ratio(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}

func              Av_buffersink_get_channels            (ctx *AVFilterContext) int32 {
    return int32(C.av_buffersink_get_channels(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}
                             

func         Av_buffersink_get_channel_layout      (ctx *AVFilterContext) uint64 {
    return uint64(C.av_buffersink_get_channel_layout(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}
      
func              Av_buffersink_get_ch_layout           (ctx *AVFilterContext,
                                                        ch_layout *AVChannelLayout) int32 {
    return int32(C.av_buffersink_get_ch_layout(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVChannelLayout)(unsafe.Pointer(ch_layout))))
}
func              Av_buffersink_get_sample_rate         (ctx *AVFilterContext) int32 {
    return int32(C.av_buffersink_get_sample_rate(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx))))
}

func     Av_buffersink_get_hw_frames_ctx       (ctx *AVFilterContext) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffersink_get_hw_frames_ctx(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)))))
}

/** @} */

/**
 * Get a frame with filtered data from sink and put it in frame.
 *
 * @param ctx pointer to a context of a buffersink or abuffersink AVFilter.
 * @param frame pointer to an allocated frame that will be filled with data.
 *              The data must be freed using av_frame_unref() / av_frame_free()
 *
 * @return
 *         - >= 0 if a frame was successfully returned.
 *         - AVERROR(EAGAIN) if no frames are available at this point; more
 *           input frames must be added to the filtergraph to get more output.
 *         - AVERROR_EOF if there will be no more output frames on this sink.
 *         - A different negative AVERROR code in other failure cases.
 */
func Av_buffersink_get_frame(ctx *AVFilterContext, frame *AVFrame) int32 {
    return int32(C.av_buffersink_get_frame(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Same as av_buffersink_get_frame(), but with the ability to specify the number
 * of samples read. This function is less efficient than
 * av_buffersink_get_frame(), because it copies the data around.
 *
 * @param ctx pointer to a context of the abuffersink AVFilter.
 * @param frame pointer to an allocated frame that will be filled with data.
 *              The data must be freed using av_frame_unref() / av_frame_free()
 *              frame will contain exactly nb_samples audio samples, except at
 *              the end of stream, when it can contain less than nb_samples.
 *
 * @return The return codes have the same meaning as for
 *         av_buffersink_get_frame().
 *
 * @warning do not mix this function with av_buffersink_get_frame(). Use only one or
 * the other with a single sink, not both.
 */
func Av_buffersink_get_samples(ctx *AVFilterContext, frame *AVFrame, nb_samples int32) int32 {
    return int32(C.av_buffersink_get_samples(
        (*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(nb_samples)))
}

/**
 * @}
 */

                                  
