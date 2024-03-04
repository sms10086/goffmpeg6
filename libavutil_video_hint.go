// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/**
 * Copyright 2023 Elias Carotti <eliascrt at amazon dot it>
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

//#cgo pkg-config: libavutil
//#include <stddef.h>
//#include <stdint.h>
//#include "libavutil/avassert.h"
//#include "libavutil/frame.h"
//#include "libavutil/video_hint.h"
import "C"
import (
    "unsafe"
)



                           
                           

                   
                   
                               
                            

type AVVideoRect C.struct_AVVideoRect

type AVVideoHintType C.enum_AVVideoHintType

type AVVideoHint C.struct_AVVideoHint

// *av_video_hint_rects(constAVVideoHint*hints)

// *av_video_hint_get_rect(constAVVideoHint*hints,size_tidx)

/**
 * Allocate memory for the AVVideoHint struct along with an nb_rects-sized
 * arrays of AVVideoRect.
 *
 * The side data contains a list of rectangles for the portions of the frame
 * which changed from the last encoded one (and the remainder are assumed to be
 * changed), or, alternately (depending on the type parameter) the unchanged
 * ones (and the remanining ones are those which changed).
 * Macroblocks will thus be hinted either to be P_SKIP-ped or go through the
 * regular encoding procedure.
 *
 * It's responsibility of the caller to fill the AVRects accordingly, and to set
 * the proper AVVideoHintType field.
 *
 * @param out_size if non-NULL, the size in bytes of the resulting data array is
 *                 written here
 *
 * @return newly allocated AVVideoHint struct (must be freed by the caller using
 *         av_free()) on success, NULL on memory allocation failure
 */
func Av_video_hint_alloc(nb_rects uint64,
                                 out_size *uint64) *AVVideoHint {
    return (*AVVideoHint)(unsafe.Pointer(C.av_video_hint_alloc(C.ulonglong(nb_rects), 
        (*C.ulonglong)(unsafe.Pointer(out_size)))))
}

/**
 * Same as av_video_hint_alloc(), except newly-allocated AVVideoHint is attached
 * as side data of type AV_FRAME_DATA_VIDEO_HINT_INFO to frame.
 */
func Av_video_hint_create_side_data(frame *AVFrame,
                                            nb_rects uint64) *AVVideoHint {
    return (*AVVideoHint)(unsafe.Pointer(C.av_video_hint_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)), C.ulonglong(nb_rects))))
}


                                
