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
//#include "libavutil/frame.h"
//#include "libavutil/film_grain_params.h"
import "C"
import (
    "unsafe"
)



                                  
                                  

                  

type AVFilmGrainParamsType int32
const (
    AV_FILM_GRAIN_PARAMS_NONE AVFilmGrainParamsType = 0 + iota
    AV_FILM_GRAIN_PARAMS_AV1
    AV_FILM_GRAIN_PARAMS_H274
)


/**
 * This structure describes how to handle film grain synthesis for AOM codecs.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainAOMParams struct {
    Num_y_points int32
    Y_points[14] [2]uint8
    Chroma_scaling_from_luma int32
    Num_uv_points [2]int32
    Uv_points[2][10] [2]uint8
    Scaling_shift int32
    Ar_coeff_lag int32
    Ar_coeffs_y [24]int8
    Ar_coeffs_uv[2] [25]int8
    Ar_coeff_shift int32
    Grain_scale_shift int32
    Uv_mult [2]int32
    Uv_mult_luma [2]int32
    Uv_offset [2]int32
    Overlap_flag int32
    Limit_output_range int32
}


/**
 * This structure describes how to handle film grain synthesis for codecs using
 * the ITU-T H.274 Versatile suplemental enhancement information message.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainH274Params struct {
    Model_id int32
    Bit_depth_luma int32
    Bit_depth_chroma int32
    Color_range AVColorRange
    Color_primaries AVColorPrimaries
    Color_trc AVColorTransferCharacteristic
    Color_space AVColorSpace
    Blending_mode_id int32
    Log2_scale_factor int32
    Component_model_present [3]int32
    Num_intensity_intervals [3]uint16
    Num_model_values [3]uint8
    Intensity_interval_lower_bound[3] [256]uint8
    Intensity_interval_upper_bound[3] [256]uint8
    Comp_model_value[3][256] [6]int16
}


/**
 * This structure describes how to handle film grain synthesis in video
 * for specific codecs. Must be present on every frame where film grain is
 * meant to be synthesised for correct presentation.
 *
 * @note The struct must be allocated with av_film_grain_params_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVFilmGrainParams struct {
    Type AVFilmGrainParamsType
    Seed uint64
    Codec AVFilmGrainH274Params
}


/**
 * Allocate an AVFilmGrainParams structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 * If size is not NULL it will be set to the number of bytes allocated.
 *
 * @return An AVFilmGrainParams filled with default values or NULL
 *         on failure.
 */
func Av_film_grain_params_alloc(size *uint64) *AVFilmGrainParams {
    return (*AVFilmGrainParams)(unsafe.Pointer(C.av_film_grain_params_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Allocate a complete AVFilmGrainParams and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVFilmGrainParams structure to be filled by caller.
 */
func Av_film_grain_params_create_side_data(frame *AVFrame) *AVFilmGrainParams {
    return (*AVFilmGrainParams)(unsafe.Pointer(C.av_film_grain_params_create_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

                                       
