// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2020 Vacing Fang <vacingfang@tencent.com>
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
//#include <stdint.h>
//#include <stddef.h>
//#include "libavutil/rational.h"
//#include "libavutil/dovi_meta.h"
import "C"
import (
    "unsafe"
)

const AV_DOVI_MAX_PIECES = 8


/**
 * @file
 * DOVI configuration
 */


                          
                          

                   
                   
                     

/*
 * DOVI configuration
 * ref: dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2
        dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2
 * @code
 * uint8_t  dv_version_major, the major version number that the stream complies with
 * uint8_t  dv_version_minor, the minor version number that the stream complies with
 * uint8_t  dv_profile, the Dolby Vision profile
 * uint8_t  dv_level, the Dolby Vision level
 * uint8_t  rpu_present_flag
 * uint8_t  el_present_flag
 * uint8_t  bl_present_flag
 * uint8_t  dv_bl_signal_compatibility_id
 * @endcode
 *
 * @note The struct must be allocated with av_dovi_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVDOVIDecoderConfigurationRecord C.struct_AVDOVIDecoderConfigurationRecord

/**
 * Allocate a AVDOVIDecoderConfigurationRecord structure and initialize its
 * fields to default values.
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_dovi_alloc(size *uint64) *AVDOVIDecoderConfigurationRecord {
    return (*AVDOVIDecoderConfigurationRecord)(unsafe.Pointer(C.av_dovi_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Dolby Vision RPU data header.
 *
 * @note sizeof(AVDOVIRpuDataHeader) is not part of the public ABI.
 */
type AVDOVIRpuDataHeader C.struct_AVDOVIRpuDataHeader

type AVDOVIMappingMethod C.enum_AVDOVIMappingMethod

/**
 * Coefficients of a piece-wise function. The pieces of the function span the
 * value ranges between two adjacent pivot values.
 */

type AVDOVIReshapingCurve C.struct_AVDOVIReshapingCurve

type AVDOVINLQMethod C.enum_AVDOVINLQMethod

/**
 * Coefficients of the non-linear inverse quantization. For the interpretation
 * of these, see ETSI GS CCM 001.
 */
type AVDOVINLQParams C.struct_AVDOVINLQParams

/**
 * Dolby Vision RPU data mapping parameters.
 *
 * @note sizeof(AVDOVIDataMapping) is not part of the public ABI.
 */
type AVDOVIDataMapping C.struct_AVDOVIDataMapping

/**
 * Dolby Vision RPU colorspace metadata parameters.
 *
 * @note sizeof(AVDOVIColorMetadata) is not part of the public ABI.
 */
type AVDOVIColorMetadata C.struct_AVDOVIColorMetadata

/**
 * Combined struct representing a combination of header, mapping and color
 * metadata, for attaching to frames as side data.
 *
 * @note The struct must be allocated with av_dovi_metadata_alloc() and
 *       its size is not a part of the public ABI.
 */

type AVDOVIMetadata C.struct_AVDOVIMetadata

// *av_dovi_get_header(constAVDOVIMetadata*data)

// *av_dovi_get_mapping(constAVDOVIMetadata*data)

// *av_dovi_get_color(constAVDOVIMetadata*data)

/**
 * Allocate an AVDOVIMetadata structure and initialize its
 * fields to default values.
 *
 * @param size If this parameter is non-NULL, the size in bytes of the
 *             allocated struct will be written here on success
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_dovi_metadata_alloc(size *uint64) *AVDOVIMetadata {
    return (*AVDOVIMetadata)(unsafe.Pointer(C.av_dovi_metadata_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

                               
