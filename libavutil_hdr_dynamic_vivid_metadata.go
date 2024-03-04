// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2021 Limin Wang <lance.lmwang at gmail.com>
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
//#include "libavutil/frame.h"
//#include "libavutil/rational.h"
//#include "libavutil/hdr_dynamic_vivid_metadata.h"
import "C"
import (
    "unsafe"
)



                                           
                                           

                  
                     

/**
 * HDR Vivid three spline params.
 */
type AVHDRVivid3SplineParams C.struct_AVHDRVivid3SplineParams

/**
 * Color tone mapping parameters at a processing window in a dynamic metadata for
 * CUVA 005.1:2021.
 */
type AVHDRVividColorToneMappingParams C.struct_AVHDRVividColorToneMappingParams


/**
 * Color transform parameters at a processing window in a dynamic metadata for
 * CUVA 005.1:2021.
 */
type AVHDRVividColorTransformParams C.struct_AVHDRVividColorTransformParams

/**
 * This struct represents dynamic metadata for color volume transform -
 * CUVA 005.1:2021 standard
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with
 * av_dynamic_hdr_vivid_alloc() and its size is not a part of
 * the public ABI.
 */
type AVDynamicHDRVivid C.struct_AVDynamicHDRVivid

/**
 * Allocate an AVDynamicHDRVivid structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVDynamicHDRVivid filled with default values or NULL
 *         on failure.
 */
func Av_dynamic_hdr_vivid_alloc(size *uint64) *AVDynamicHDRVivid {
    return (*AVDynamicHDRVivid)(unsafe.Pointer(C.av_dynamic_hdr_vivid_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Allocate a complete AVDynamicHDRVivid and add it to the frame.
 * @param frame The frame which side data is added to.
 *
 * @return The AVDynamicHDRVivid structure to be filled by caller or NULL
 *         on failure.
 */
func Av_dynamic_hdr_vivid_create_side_data(frame *AVFrame) *AVDynamicHDRVivid {
    return (*AVDynamicHDRVivid)(unsafe.Pointer(C.av_dynamic_hdr_vivid_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)))))
}

                                                
