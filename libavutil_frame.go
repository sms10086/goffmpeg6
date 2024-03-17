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

//#cgo pkg-config: libavutil
//#include <stddef.h>
//#include <stdint.h>
//#include "libavutil/avutil.h"
//#include "libavutil/buffer.h"
//#include "libavutil/channel_layout.h"
//#include "libavutil/dict.h"
//#include "libavutil/rational.h"
//#include "libavutil/samplefmt.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/version.h"
//#include "libavutil/frame.h"
import "C"
import (
    "unsafe"
)

const AV_NUM_DATA_POINTERS = 8
const AV_FRAME_FLAG_CORRUPT =        (1 << 0) 
const AV_FRAME_FLAG_KEY =  (1 << 1) 
const AV_FRAME_FLAG_DISCARD =    (1 << 2) 
const AV_FRAME_FLAG_INTERLACED =  (1 << 3) 
const AV_FRAME_FLAG_TOP_FIELD_FIRST =  (1 << 4) 
const FF_DECODE_ERROR_INVALID_BITSTREAM = 1
const FF_DECODE_ERROR_MISSING_REFERENCE = 2
const FF_DECODE_ERROR_CONCEALMENT_ACTIVE = 4
const FF_DECODE_ERROR_DECODE_SLICES = 8


/**
 * @file
 * @ingroup lavu_frame
 * reference-counted frame API
 */

                      
                      

                   
                   

                   
                   
                           
                 
                     
                      
                   
                    


/**
 * @defgroup lavu_frame AVFrame
 * @ingroup lavu_data
 *
 * @{
 * AVFrame is an abstraction for reference-counted raw multimedia data.
 */

type AVFrameSideDataType int32
const (
    AV_FRAME_DATA_PANSCAN AVFrameSideDataType = iota
    AV_FRAME_DATA_A53_CC
    AV_FRAME_DATA_STEREO3D
    AV_FRAME_DATA_MATRIXENCODING
    AV_FRAME_DATA_DOWNMIX_INFO
    AV_FRAME_DATA_REPLAYGAIN
    AV_FRAME_DATA_DISPLAYMATRIX
    AV_FRAME_DATA_AFD
    AV_FRAME_DATA_MOTION_VECTORS
    AV_FRAME_DATA_SKIP_SAMPLES
    AV_FRAME_DATA_AUDIO_SERVICE_TYPE
    AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
    AV_FRAME_DATA_GOP_TIMECODE
    AV_FRAME_DATA_SPHERICAL
    AV_FRAME_DATA_CONTENT_LIGHT_LEVEL
    AV_FRAME_DATA_ICC_PROFILE
    AV_FRAME_DATA_S12M_TIMECODE
    AV_FRAME_DATA_DYNAMIC_HDR_PLUS
    AV_FRAME_DATA_REGIONS_OF_INTEREST
    AV_FRAME_DATA_VIDEO_ENC_PARAMS
    AV_FRAME_DATA_SEI_UNREGISTERED
    AV_FRAME_DATA_FILM_GRAIN_PARAMS
    AV_FRAME_DATA_DETECTION_BBOXES
    AV_FRAME_DATA_DOVI_RPU_BUFFER
    AV_FRAME_DATA_DOVI_METADATA
    AV_FRAME_DATA_DYNAMIC_HDR_VIVID
    AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT
    AV_FRAME_DATA_VIDEO_HINT
)


type AVActiveFormatDescription int32
const (
    AV_AFD_SAME AVActiveFormatDescription = 8 + iota
    AV_AFD_4_3 = 9
    AV_AFD_16_9 = 10
    AV_AFD_14_9 = 11
    AV_AFD_4_3_SP_14_9 = 13
    AV_AFD_16_9_SP_14_9 = 14
    AV_AFD_SP_4_3 = 15
)



/**
 * Structure to hold side data for an AVFrame.
 *
 * sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 */
type AVFrameSideData struct {
    Type AVFrameSideDataType
    Data *uint8
    Size uint64
    Metadata *AVDictionary
    Buf *AVBufferRef
}


/**
 * Structure describing a single Region Of Interest.
 *
 * When multiple regions are defined in a single side-data block, they
 * should be ordered from most to least important - some encoders are only
 * capable of supporting a limited number of distinct regions, so will have
 * to truncate the list.
 *
 * When overlapping regions are defined, the first region containing a given
 * area of the frame applies.
 */
type AVRegionOfInterest struct {
    Self_size uint32
    Top int32
    Bottom int32
    Left int32
    Right int32
    Qoffset AVRational
}


/**
 * This structure describes decoded (raw) audio or video data.
 *
 * AVFrame must be allocated using av_frame_alloc(). Note that this only
 * allocates the AVFrame itself, the buffers for the data must be managed
 * through other means (see below).
 * AVFrame must be freed with av_frame_free().
 *
 * AVFrame is typically allocated once and then reused multiple times to hold
 * different data (e.g. a single AVFrame to hold frames received from a
 * decoder). In such a case, av_frame_unref() will free any references held by
 * the frame and reset it to its original clean state before it
 * is reused again.
 *
 * The data described by an AVFrame is usually reference counted through the
 * AVBuffer API. The underlying buffer references are stored in AVFrame.buf /
 * AVFrame.extended_buf. An AVFrame is considered to be reference counted if at
 * least one reference is set, i.e. if AVFrame.buf[0] != NULL. In such a case,
 * every single data plane must be contained in one of the buffers in
 * AVFrame.buf or AVFrame.extended_buf.
 * There may be a single buffer for all the data, or one separate buffer for
 * each plane, or anything in between.
 *
 * sizeof(AVFrame) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 *
 * Fields can be accessed through AVOptions, the name string used, matches the
 * C structure field name for fields accessible through AVOptions. The AVClass
 * for AVFrame can be obtained from avcodec_get_frame_class()
 */
type AVFrame struct {
    Data [AV_NUM_DATA_POINTERS]*uint8
    Linesize [AV_NUM_DATA_POINTERS]int32
    Extended_data **uint8
    Width int32
    Height int32
    Nb_samples int32
    Format int32
    Key_frame int32
    Pict_type AVPictureType
    Sample_aspect_ratio AVRational
    Pts int64
    Pkt_dts int64
    Time_base AVRational
    Coded_picture_number int32
    Display_picture_number int32
    Quality int32
    Opaque unsafe.Pointer
    Repeat_pict int32
    Interlaced_frame int32
    Top_field_first int32
    Palette_has_changed int32
    Reordered_opaque int64
    Sample_rate int32
    Channel_layout uint64
    Buf [AV_NUM_DATA_POINTERS]*AVBufferRef
    Extended_buf **AVBufferRef
    Nb_extended_buf int32
    Side_data **AVFrameSideData
    Nb_side_data int32
    Flags int32
    Color_range AVColorRange
    Color_primaries AVColorPrimaries
    Color_trc AVColorTransferCharacteristic
    Colorspace AVColorSpace
    Chroma_location AVChromaLocation
    Best_effort_timestamp int64
    Pkt_pos int64
    Pkt_duration int64
    Metadata *AVDictionary
    Decode_error_flags int32
    Channels int32
    Pkt_size int32
    Hw_frames_ctx *AVBufferRef
    Opaque_ref *AVBufferRef
    Crop_top uint64
    Crop_bottom uint64
    Crop_left uint64
    Crop_right uint64
    Private_ref *AVBufferRef
    Ch_layout AVChannelLayout
    Duration int64
}



/**
 * Allocate an AVFrame and set its fields to default values.  The resulting
 * struct must be freed using av_frame_free().
 *
 * @return An AVFrame filled with default values or NULL on failure.
 *
 * @note this only allocates the AVFrame itself, not the data buffers. Those
 * must be allocated through other means, e.g. with av_frame_get_buffer() or
 * manually.
 */
func Av_frame_alloc() *AVFrame {
    return (*AVFrame)(unsafe.Pointer(C.av_frame_alloc()))
}

/**
 * Free the frame and any dynamically allocated objects in it,
 * e.g. extended_data. If the frame is reference counted, it will be
 * unreferenced first.
 *
 * @param frame frame to be freed. The pointer will be set to NULL.
 */
func Av_frame_free(frame **AVFrame)  {
    C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(frame)))
}

/**
 * Set up a new reference to the data described by the source frame.
 *
 * Copy frame properties from src to dst and create a new reference for each
 * AVBufferRef from src.
 *
 * If src is not reference counted, new buffers are allocated and the data is
 * copied.
 *
 * @warning: dst MUST have been either unreferenced with av_frame_unref(dst),
 *           or newly allocated with av_frame_alloc() before calling this
 *           function, or undefined behavior will occur.
 *
 * @return 0 on success, a negative AVERROR on error
 */
func Av_frame_ref(dst *AVFrame, src *AVFrame) int32 {
    return int32(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(dst)), 
        (*C.struct_AVFrame)(unsafe.Pointer(src))))
}

/**
 * Ensure the destination frame refers to the same data described by the source
 * frame, either by creating a new reference for each AVBufferRef from src if
 * they differ from those in dst, by allocating new buffers and copying data if
 * src is not reference counted, or by unrefencing it if src is empty.
 *
 * Frame properties on dst will be replaced by those from src.
 *
 * @return 0 on success, a negative AVERROR on error. On error, dst is
 *         unreferenced.
 */
func Av_frame_replace(dst *AVFrame, src *AVFrame) int32 {
    return int32(C.av_frame_replace((*C.struct_AVFrame)(unsafe.Pointer(dst)), 
        (*C.struct_AVFrame)(unsafe.Pointer(src))))
}

/**
 * Create a new frame that references the same data as src.
 *
 * This is a shortcut for av_frame_alloc()+av_frame_ref().
 *
 * @return newly created AVFrame on success, NULL on error.
 */
func Av_frame_clone(src *AVFrame) *AVFrame {
    return (*AVFrame)(unsafe.Pointer(C.av_frame_clone(
        (*C.struct_AVFrame)(unsafe.Pointer(src)))))
}

/**
 * Unreference all the buffers referenced by frame and reset the frame fields.
 */
func Av_frame_unref(frame *AVFrame)  {
    C.av_frame_unref((*C.struct_AVFrame)(unsafe.Pointer(frame)))
}

/**
 * Move everything contained in src to dst and reset src.
 *
 * @warning: dst is not unreferenced, but directly overwritten without reading
 *           or deallocating its contents. Call av_frame_unref(dst) manually
 *           before calling this function to ensure that no memory is leaked.
 */
func Av_frame_move_ref(dst *AVFrame, src *AVFrame)  {
    C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(dst)), 
        (*C.struct_AVFrame)(unsafe.Pointer(src)))
}

/**
 * Allocate new buffer(s) for audio or video data.
 *
 * The following fields must be set on frame before calling this function:
 * - format (pixel format for video, sample format for audio)
 * - width and height for video
 * - nb_samples and ch_layout for audio
 *
 * This function will fill AVFrame.data and AVFrame.buf arrays and, if
 * necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
 * For planar formats, one buffer will be allocated for each plane.
 *
 * @warning: if frame already has been allocated, calling this function will
 *           leak memory. In addition, undefined behavior can occur in certain
 *           cases.
 *
 * @param frame frame in which to store the new buffers.
 * @param align Required buffer size alignment. If equal to 0, alignment will be
 *              chosen automatically for the current CPU. It is highly
 *              recommended to pass 0 here unless you know what you are doing.
 *
 * @return 0 on success, a negative AVERROR on error.
 */
func Av_frame_get_buffer(frame *AVFrame, align int32) int32 {
    return int32(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(frame)), 
        C.int(align)))
}

/**
 * Check if the frame data is writable.
 *
 * @return A positive value if the frame data is writable (which is true if and
 * only if each of the underlying buffers has only one reference, namely the one
 * stored in this frame). Return 0 otherwise.
 *
 * If 1 is returned the answer is valid until av_buffer_ref() is called on any
 * of the underlying AVBufferRefs (e.g. through av_frame_ref() or directly).
 *
 * @see av_frame_make_writable(), av_buffer_is_writable()
 */
func Av_frame_is_writable(frame *AVFrame) int32 {
    return int32(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Ensure that the frame data is writable, avoiding data copy if possible.
 *
 * Do nothing if the frame is writable, allocate new buffers and copy the data
 * if it is not. Non-refcounted frames behave as non-writable, i.e. a copy
 * is always made.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_frame_is_writable(), av_buffer_is_writable(),
 * av_buffer_make_writable()
 */
func Av_frame_make_writable(frame *AVFrame) int32 {
    return int32(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Copy the frame data from src to dst.
 *
 * This function does not allocate anything, dst must be already initialized and
 * allocated with the same parameters as src.
 *
 * This function only copies the frame data (i.e. the contents of the data /
 * extended data arrays), not any other properties.
 *
 * @return >= 0 on success, a negative AVERROR on error.
 */
func Av_frame_copy(dst *AVFrame, src *AVFrame) int32 {
    return int32(C.av_frame_copy((*C.struct_AVFrame)(unsafe.Pointer(dst)), 
        (*C.struct_AVFrame)(unsafe.Pointer(src))))
}

/**
 * Copy only "metadata" fields from src to dst.
 *
 * Metadata for the purpose of this function are those fields that do not affect
 * the data layout in the buffers.  E.g. pts, sample rate (for audio) or sample
 * aspect ratio (for video), but not width/height or channel layout.
 * Side data is also copied.
 */
func Av_frame_copy_props(dst *AVFrame, src *AVFrame) int32 {
    return int32(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(dst)), 
        (*C.struct_AVFrame)(unsafe.Pointer(src))))
}

/**
 * Get the buffer reference a given data plane is stored in.
 *
 * @param frame the frame to get the plane's buffer from
 * @param plane index of the data plane of interest in frame->extended_data.
 *
 * @return the buffer reference that contains the plane or NULL if the input
 * frame is not valid.
 */
func Av_frame_get_plane_buffer(frame *AVFrame, plane int32) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_frame_get_plane_buffer(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(plane))))
}

/**
 * Add a new side data to a frame.
 *
 * @param frame a frame to which the side data should be added
 * @param type type of the added side data
 * @param size size of the side data
 *
 * @return newly added side data on success, NULL on error
 */
func Av_frame_new_side_data(frame *AVFrame,
                                        typex AVFrameSideDataType,
                                        size uint64) *AVFrameSideData {
    return (*AVFrameSideData)(unsafe.Pointer(C.av_frame_new_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), 
        C.enum_AVFrameSideDataType(typex), C.ulonglong(size))))
}

/**
 * Add a new side data to a frame from an existing AVBufferRef
 *
 * @param frame a frame to which the side data should be added
 * @param type  the type of the added side data
 * @param buf   an AVBufferRef to add as side data. The ownership of
 *              the reference is transferred to the frame.
 *
 * @return newly added side data on success, NULL on error. On failure
 *         the frame is unchanged and the AVBufferRef remains owned by
 *         the caller.
 */
func Av_frame_new_side_data_from_buf(frame *AVFrame,
                                                 typex AVFrameSideDataType,
                                                 buf *AVBufferRef) *AVFrameSideData {
    return (*AVFrameSideData)(unsafe.Pointer(C.av_frame_new_side_data_from_buf(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), 
        C.enum_AVFrameSideDataType(typex), 
        (*C.struct_AVBufferRef)(unsafe.Pointer(buf)))))
}

/**
 * @return a pointer to the side data of a given type on success, NULL if there
 * is no side data with such type in this frame.
 */
func Av_frame_get_side_data(frame *AVFrame,
                                        typex AVFrameSideDataType) *AVFrameSideData {
    return (*AVFrameSideData)(unsafe.Pointer(C.av_frame_get_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), 
        C.enum_AVFrameSideDataType(typex))))
}

/**
 * Remove and free all side data instances of the given type.
 */
func Av_frame_remove_side_data(frame *AVFrame, typex AVFrameSideDataType)  {
    C.av_frame_remove_side_data((*C.struct_AVFrame)(unsafe.Pointer(frame)), 
        C.enum_AVFrameSideDataType(typex))
}


/**
 * Flags for frame cropping.
 */
const (
    AV_FRAME_CROP_UNALIGNED  = 1<<0 + iota
)


/**
 * Crop the given video AVFrame according to its crop_left/crop_top/crop_right/
 * crop_bottom fields. If cropping is successful, the function will adjust the
 * data pointers and the width/height fields, and set the crop fields to 0.
 *
 * In all cases, the cropping boundaries will be rounded to the inherent
 * alignment of the pixel format. In some cases, such as for opaque hwaccel
 * formats, the left/top cropping is ignored. The crop fields are set to 0 even
 * if the cropping was rounded or ignored.
 *
 * @param frame the frame which should be cropped
 * @param flags Some combination of AV_FRAME_CROP_* flags, or 0.
 *
 * @return >= 0 on success, a negative AVERROR on error. If the cropping fields
 * were invalid, AVERROR(ERANGE) is returned, and nothing is changed.
 */
func Av_frame_apply_cropping(frame *AVFrame, flags int32) int32 {
    return int32(C.av_frame_apply_cropping(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * @return a string identifying the side data type
 */
func Av_frame_side_data_name(typex AVFrameSideDataType) string {
    return C.GoString(C.av_frame_side_data_name(C.enum_AVFrameSideDataType(typex)))
}

/**
 * @}
 */

                           
