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



                                  
                                  

                  

type AVFilmGrainParamsType C.enum_AVFilmGrainParamsType

/**
 * This structure describes how to handle film grain synthesis for AOM codecs.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainAOMParams C.struct_AVFilmGrainAOMParams

/**
 * This structure describes how to handle film grain synthesis for codecs using
 * the ITU-T H.274 Versatile suplemental enhancement information message.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainH274Params C.struct_AVFilmGrainH274Params

/**
 * This structure describes how to handle film grain synthesis in video
 * for specific codecs. Must be present on every frame where film grain is
 * meant to be synthesised for correct presentation.
 *
 * @note The struct must be allocated with av_film_grain_params_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVFilmGrainParams C.struct_AVFilmGrainParams

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
        (*C.AVFrame)(unsafe.Pointer(frame)))))
}

                                       
