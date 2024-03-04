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
//#include "libavutil/rational.h"
//#include "libavutil/avassert.h"
//#include "libavutil/frame.h"
//#include "libavutil/detection_bbox.h"
import "C"
import (
    "unsafe"
)

const AV_DETECTION_BBOX_LABEL_NAME_MAX_SIZE = 64
const AV_NUM_DETECTION_BBOX_CLASSIFY = 4


                               
                               

                     
                     
                  

type AVDetectionBBox C.struct_AVDetectionBBox

type AVDetectionBBoxHeader C.struct_AVDetectionBBoxHeader

/*
 * Get the bounding box at the specified {@code idx}. Must be between 0 and nb_bboxes.
 */
// *av_get_detection_bbox(constAVDetectionBBoxHeader*header,unsignedintidx)

/**
 * Allocates memory for AVDetectionBBoxHeader, plus an array of {@code nb_bboxes}
 * AVDetectionBBox, and initializes the variables.
 * Can be freed with a normal av_free() call.
 *
 * @param nb_bboxes number of AVDetectionBBox structures to allocate
 * @param out_size if non-NULL, the size in bytes of the resulting data array is
 * written here.
 */
func Av_detection_bbox_alloc(nb_bboxes uint32, out_size *uint64) *AVDetectionBBoxHeader {
    return (*AVDetectionBBoxHeader)(unsafe.Pointer(C.av_detection_bbox_alloc(C.uint(nb_bboxes), 
        (*C.ulonglong)(unsafe.Pointer(out_size)))))
}

/**
 * Allocates memory for AVDetectionBBoxHeader, plus an array of {@code nb_bboxes}
 * AVDetectionBBox, in the given AVFrame {@code frame} as AVFrameSideData of type
 * AV_FRAME_DATA_DETECTION_BBOXES and initializes the variables.
 */
func Av_detection_bbox_create_side_data(frame *AVFrame, nb_bboxes uint32) *AVDetectionBBoxHeader {
    return (*AVDetectionBBoxHeader)(unsafe.Pointer(C.av_detection_bbox_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)), C.uint(nb_bboxes))))
}
      
