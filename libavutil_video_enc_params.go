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
//#include "libavutil/avassert.h"
//#include "libavutil/frame.h"
//#include "libavutil/video_enc_params.h"
import "C"
import (
    "unsafe"
)



                                 
                                 

                   
                   

                               
                            

type AVVideoEncParamsType C.enum_AVVideoEncParamsType

/**
 * Video encoding parameters for a given frame. This struct is allocated along
 * with an optional array of per-block AVVideoBlockParams descriptors.
 * Must be allocated with av_video_enc_params_alloc().
 */
type AVVideoEncParams C.struct_AVVideoEncParams

/**
 * Data structure for storing block-level encoding information.
 * It is allocated as a part of AVVideoEncParams and should be retrieved with
 * av_video_enc_params_block().
 *
 * sizeof(AVVideoBlockParams) is not a part of the ABI and new fields may be
 * added to it.
 */
type AVVideoBlockParams C.struct_AVVideoBlockParams

/**
 * Get the block at the specified {@code idx}. Must be between 0 and nb_blocks - 1.
 */
// *av_video_enc_params_block(AVVideoEncParams*par,unsignedintidx)

/**
 * Allocates memory for AVVideoEncParams of the given type, plus an array of
 * {@code nb_blocks} AVVideoBlockParams and initializes the variables. Can be
 * freed with a normal av_free() call.
 *
 * @param out_size if non-NULL, the size in bytes of the resulting data array is
 * written here.
 */
func Av_video_enc_params_alloc(typex AVVideoEncParamsType,
                                            nb_blocks uint32, out_size *uint64) *AVVideoEncParams {
    return (*AVVideoEncParams)(unsafe.Pointer(C.av_video_enc_params_alloc(
        C.enum_AVVideoEncParamsType(typex), C.uint(nb_blocks), 
        (*C.ulonglong)(unsafe.Pointer(out_size)))))
}

/**
 * Allocates memory for AVEncodeInfoFrame plus an array of
 * {@code nb_blocks} AVEncodeInfoBlock in the given AVFrame {@code frame}
 * as AVFrameSideData of type AV_FRAME_DATA_VIDEO_ENC_PARAMS
 * and initializes the variables.
 */
func
Av_video_enc_params_create_side_data(frame *AVFrame, typex AVVideoEncParamsType,
                                     nb_blocks uint32) *AVVideoEncParams {
    return (*AVVideoEncParams)(unsafe.Pointer(C.av_video_enc_params_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)), C.enum_AVVideoEncParamsType(typex), 
        C.uint(nb_blocks))))
}

                                      
