// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2018 Mohammad Izadi <moh.izadi at gmail.com>
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
//#include "libavutil/hdr_dynamic_metadata.h"
import "C"
import (
    "unsafe"
)

const AV_HDR_PLUS_MAX_PAYLOAD_SIZE = 907


                                     
                                     

                  
                     

/**
 * Option for overlapping elliptical pixel selectors in an image.
 */
type AVHDRPlusOverlapProcessOption C.enum_AVHDRPlusOverlapProcessOption

/**
 * Represents the percentile at a specific percentage in
 * a distribution.
 */
type AVHDRPlusPercentile C.struct_AVHDRPlusPercentile

/**
 * Color transform parameters at a processing window in a dynamic metadata for
 * SMPTE 2094-40.
 */
type AVHDRPlusColorTransformParams C.struct_AVHDRPlusColorTransformParams

/**
 * This struct represents dynamic metadata for color volume transform -
 * application 4 of SMPTE 2094-40:2016 standard.
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with
 * av_dynamic_hdr_plus_alloc() and its size is not a part of
 * the public ABI.
 */
type AVDynamicHDRPlus C.struct_AVDynamicHDRPlus

/**
 * Allocate an AVDynamicHDRPlus structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVDynamicHDRPlus filled with default values or NULL
 *         on failure.
 */
func Av_dynamic_hdr_plus_alloc(size *uint64) *AVDynamicHDRPlus {
    return (*AVDynamicHDRPlus)(unsafe.Pointer(C.av_dynamic_hdr_plus_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Allocate a complete AVDynamicHDRPlus and add it to the frame.
 * @param frame The frame which side data is added to.
 *
 * @return The AVDynamicHDRPlus structure to be filled by caller or NULL
 *         on failure.
 */
func Av_dynamic_hdr_plus_create_side_data(frame *AVFrame) *AVDynamicHDRPlus {
    return (*AVDynamicHDRPlus)(unsafe.Pointer(C.av_dynamic_hdr_plus_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)))))
}

/**
 * Parse the user data registered ITU-T T.35 to AVbuffer (AVDynamicHDRPlus).
 * The T.35 buffer must begin with the application mode, skipping the
 * country code, terminal provider codes, and application identifier.
 * @param s A pointer containing the decoded AVDynamicHDRPlus structure.
 * @param data The byte array containing the raw ITU-T T.35 data.
 * @param size Size of the data array in bytes.
 *
 * @return >= 0 on success. Otherwise, returns the appropriate AVERROR.
 */
func Av_dynamic_hdr_plus_from_t35(s *AVDynamicHDRPlus, data *uint8,
                                 size uint64) int32 {
    return int32(C.av_dynamic_hdr_plus_from_t35(
        (*C.AVDynamicHDRPlus)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(data)), C.ulonglong(size)))
}



/**
 * Serialize dynamic HDR10+ metadata to a user data registered ITU-T T.35 buffer,
 * excluding the first 48 bytes of the header, and beginning with the application mode.
 * @param s A pointer containing the decoded AVDynamicHDRPlus structure.
 * @param data[in,out] A pointer to pointer to a byte buffer to be filled with the
 *                     serialized metadata.
 *                     If *data is NULL, a buffer be will be allocated and a pointer to
 *                     it stored in its place. The caller assumes ownership of the buffer.
 *                     May be NULL, in which case the function will only store the
 *                     required buffer size in *size.
 * @param size[in,out] A pointer to a size to be set to the returned buffer's size.
 *                     If *data is not NULL, *size must contain the size of the input
 *                     buffer. May be NULL only if *data is NULL.
 *
 * @return >= 0 on success. Otherwise, returns the appropriate AVERROR.
 */
func Av_dynamic_hdr_plus_to_t35(s *AVDynamicHDRPlus, data **uint8, size *uint64) int32 {
    return int32(C.av_dynamic_hdr_plus_to_t35(
        (*C.AVDynamicHDRPlus)(unsafe.Pointer(s)), 
        (**C.uchar)(unsafe.Pointer(data)), (*C.ulonglong)(unsafe.Pointer(size))))
}

                                          
