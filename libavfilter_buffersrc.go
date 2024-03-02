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
//#include "libavfilter/buffersrc.h"
import "C"
import (
    "unsafe"
)



                            
                            

/**
 * @file
 * @ingroup lavfi_buffersrc
 * Memory buffer source API.
 */

                     

/**
 * @defgroup lavfi_buffersrc Buffer source API
 * @ingroup lavfi
 * @{
 */



/**
 * Get the number of failed requests.
 *
 * A failed request is when the request_frame method is called while no
 * frame is present in the buffer.
 * The number is reset when a frame is added.
 */
func Av_buffersrc_get_nb_failed_requests(buffer_src *AVFilterContext) uint32 {
    return uint32(C.av_buffersrc_get_nb_failed_requests(
        (*C.AVFilterContext)(unsafe.Pointer(buffer_src))))
}

/**
 * This structure contains the parameters describing the frames that will be
 * passed to this filter.
 *
 * It should be allocated with av_buffersrc_parameters_alloc() and freed with
 * av_free(). All the allocated fields in it remain owned by the caller.
 */
type AVBufferSrcParameters C.struct_AVBufferSrcParameters

/**
 * Allocate a new AVBufferSrcParameters instance. It should be freed by the
 * caller with av_free().
 */
func Av_buffersrc_parameters_alloc() *AVBufferSrcParameters {
    return (*AVBufferSrcParameters)(unsafe.Pointer(C.av_buffersrc_parameters_alloc()))
}

/**
 * Initialize the buffersrc or abuffersrc filter with the provided parameters.
 * This function may be called multiple times, the later calls override the
 * previous ones. Some of the parameters may also be set through AVOptions, then
 * whatever method is used last takes precedence.
 *
 * @param ctx an instance of the buffersrc or abuffersrc filter
 * @param param the stream parameters. The frames later passed to this filter
 *              must conform to those parameters. All the allocated fields in
 *              param remain owned by the caller, libavfilter will make internal
 *              copies or references when necessary.
 * @return 0 on success, a negative AVERROR code on failure.
 */
func Av_buffersrc_parameters_set(ctx *AVFilterContext, param *AVBufferSrcParameters) int32 {
    return int32(C.av_buffersrc_parameters_set(
        (*C.AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.AVBufferSrcParameters)(unsafe.Pointer(param))))
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will make a new reference to it. Otherwise the frame data will be
 * copied.
 *
 * @return 0 on success, a negative AVERROR on error
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() with the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
func Av_buffersrc_write_frame(ctx *AVFilterContext, frame *AVFrame) int32 {
    return int32(C.av_buffersrc_write_frame(
        (*C.AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will take ownership of the reference(s) and reset the frame.
 * Otherwise the frame data will be copied. If this function returns an error,
 * the input frame is not touched.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @note the difference between this function and av_buffersrc_write_frame() is
 * that av_buffersrc_write_frame() creates a new reference to the input frame,
 * while this function takes ownership of the reference passed to it.
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() without the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
func Av_buffersrc_add_frame(ctx *AVFilterContext, frame *AVFrame) int32 {
    return int32(C.av_buffersrc_add_frame(
        (*C.AVFilterContext)(unsafe.Pointer(ctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Add a frame to the buffer source.
 *
 * By default, if the frame is reference-counted, this function will take
 * ownership of the reference(s) and reset the frame. This can be controlled
 * using the flags.
 *
 * If this function returns an error, the input frame is not touched.
 *
 * @param buffer_src  pointer to a buffer source context
 * @param frame       a frame, or NULL to mark EOF
 * @param flags       a combination of AV_BUFFERSRC_FLAG_*
 * @return            >= 0 in case of success, a negative AVERROR code
 *                    in case of failure
 */
func Av_buffersrc_add_frame_flags(buffer_src *AVFilterContext,
                                 frame *AVFrame, flags int32) int32 {
    return int32(C.av_buffersrc_add_frame_flags(
        (*C.AVFilterContext)(unsafe.Pointer(buffer_src)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * Close the buffer source after EOF.
 *
 * This is similar to passing NULL to av_buffersrc_add_frame_flags()
 * except it takes the timestamp of the EOF, i.e. the timestamp of the end
 * of the last frame.
 */
func Av_buffersrc_close(ctx *AVFilterContext, pts int64, flags uint32) int32 {
    return int32(C.av_buffersrc_close((*C.AVFilterContext)(unsafe.Pointer(ctx)), 
        C.longlong(pts), C.unsigned(flags)))
}

/**
 * @}
 */

                                 
