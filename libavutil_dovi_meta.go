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
type AVDOVIDecoderConfigurationRecord struct {
    Dv_version_major uint8
    Dv_version_minor uint8
    Dv_profile uint8
    Dv_level uint8
    Rpu_present_flag uint8
    El_present_flag uint8
    Bl_present_flag uint8
    Dv_bl_signal_compatibility_id uint8
}


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
type AVDOVIRpuDataHeader struct {
    Rpu_type uint8
    Rpu_format uint16
    Vdr_rpu_profile uint8
    Vdr_rpu_level uint8
    Chroma_resampling_explicit_filter_flag uint8
    Coef_data_type uint8
    Coef_log2_denom uint8
    Vdr_rpu_normalized_idc uint8
    Bl_video_full_range_flag uint8
    Bl_bit_depth uint8
    El_bit_depth uint8
    Vdr_bit_depth uint8
    Spatial_resampling_filter_flag uint8
    El_spatial_resampling_filter_flag uint8
    Disable_residual_flag uint8
}


type AVDOVIMappingMethod int32
const (
    AV_DOVI_MAPPING_POLYNOMIAL AVDOVIMappingMethod = 0 + iota
    AV_DOVI_MAPPING_MMR = 1
)


/**
 * Coefficients of a piece-wise function. The pieces of the function span the
 * value ranges between two adjacent pivot values.
 */

type AVDOVIReshapingCurve struct {
    Num_pivots uint8
    Pivots [AV_DOVI_MAX_PIECES+1]uint16
    Mapping_idc [AV_DOVI_MAX_PIECES]AVDOVIMappingMethod
    Poly_order [AV_DOVI_MAX_PIECES]uint8
    Poly_coef[AV_DOVI_MAX_PIECES] [3]int64
    Mmr_order [AV_DOVI_MAX_PIECES]uint8
    Mmr_constant [AV_DOVI_MAX_PIECES]int64
    Mmr_coef[AV_DOVI_MAX_PIECES][3] [7]int64
}


type AVDOVINLQMethod int32
const (
    AV_DOVI_NLQ_NONE AVDOVINLQMethod = -1 + iota
    AV_DOVI_NLQ_LINEAR_DZ = 0
)


/**
 * Coefficients of the non-linear inverse quantization. For the interpretation
 * of these, see ETSI GS CCM 001.
 */
type AVDOVINLQParams struct {
    Nlq_offset uint16
    Vdr_in_max uint64
    Linear_deadzone_slope uint64
    Linear_deadzone_threshold uint64
}


/**
 * Dolby Vision RPU data mapping parameters.
 *
 * @note sizeof(AVDOVIDataMapping) is not part of the public ABI.
 */
type AVDOVIDataMapping struct {
    Vdr_rpu_id uint8
    Mapping_color_space uint8
    Mapping_chroma_format_idc uint8
    Curves [3]AVDOVIReshapingCurve
    Nlq_method_idc AVDOVINLQMethod
    Num_x_partitions uint32
    Num_y_partitions uint32
    Nlq [3]AVDOVINLQParams
}


/**
 * Dolby Vision RPU colorspace metadata parameters.
 *
 * @note sizeof(AVDOVIColorMetadata) is not part of the public ABI.
 */
type AVDOVIColorMetadata struct {
    Dm_metadata_id uint8
    Scene_refresh_flag uint8
    Ycc_to_rgb_matrix [9]AVRational
    Ycc_to_rgb_offset [3]AVRational
    Rgb_to_lms_matrix [9]AVRational
    Signal_eotf uint16
    Signal_eotf_param0 uint16
    Signal_eotf_param1 uint16
    Signal_eotf_param2 uint32
    Signal_bit_depth uint8
    Signal_color_space uint8
    Signal_chroma_format uint8
    Signal_full_range_flag uint8
    Source_min_pq uint16
    Source_max_pq uint16
    Source_diagonal uint16
}


/**
 * Combined struct representing a combination of header, mapping and color
 * metadata, for attaching to frames as side data.
 *
 * @note The struct must be allocated with av_dovi_metadata_alloc() and
 *       its size is not a part of the public ABI.
 */

type AVDOVIMetadata struct {
    Header_offset uint64
    Mapping_offset uint64
    Color_offset uint64
}


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

                               
