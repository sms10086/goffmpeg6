// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2015 Kevin Wheatley <kevin.j.wheatley@gmail.com>
 * Copyright (c) 2016 Ronald S. Bultje <rsbultje@gmail.com>
 * Copyright (c) 2023 Leo Izen <leo.izen@gmail.com>
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
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavutil/csp.h"
import "C"
import (
    "unsafe"
)



                    
                    

                   
                     

/**
 * @file
 * Colorspace value utility functions for libavutil.
 * @ingroup lavu_math_csp
 * @author Ronald S. Bultje <rsbultje@gmail.com>
 * @author Leo Izen <leo.izen@gmail.com>
 * @author Kevin Wheatley <kevin.j.wheatley@gmail.com>
 */

/**
 * @defgroup lavu_math_csp Colorspace Utility
 * @ingroup lavu_math
 * @{
 */

/**
 * Struct containing luma coefficients to be used for RGB to YUV/YCoCg, or similar
 * calculations.
 */
type AVLumaCoefficients struct {
    Cr AVRational
    Cg AVRational
    Cb AVRational
}


/**
 * Struct containing chromaticity x and y values for the standard CIE 1931
 * chromaticity definition.
 */
type AVCIExy struct {
    X AVRational
    Y AVRational
}


/**
 * Struct defining the red, green, and blue primary locations in terms of CIE
 * 1931 chromaticity x and y.
 */
type AVPrimaryCoefficients struct {
    R AVCIExy
    G AVCIExy
    B AVCIExy
}


/**
 * Struct defining white point location in terms of CIE 1931 chromaticity x
 * and y.
 */
type AVWhitepointCoefficients AVCIExy

/**
 * Struct that contains both white point location and primaries location, providing
 * the complete description of a color gamut.
 */
type AVColorPrimariesDesc struct {
    Wp AVWhitepointCoefficients
    Prim AVPrimaryCoefficients
}


/**
 * Function pointer representing a double -> double transfer function that performs
 * an EOTF transfer inversion. This function outputs linear light.
 */
type av_csp_trc_function C.av_csp_trc_function

/**
 * Retrieves the Luma coefficients necessary to construct a conversion matrix
 * from an enum constant describing the colorspace.
 * @param csp An enum constant indicating YUV or similar colorspace.
 * @return The Luma coefficients associated with that colorspace, or NULL
 *     if the constant is unknown to libavutil.
 */
func Av_csp_luma_coeffs_from_avcsp(csp AVColorSpace) *AVLumaCoefficients {
    return (*AVLumaCoefficients)(unsafe.Pointer(C.av_csp_luma_coeffs_from_avcsp(
        C.enum_AVColorSpace(csp))))
}

/**
 * Retrieves a complete gamut description from an enum constant describing the
 * color primaries.
 * @param prm An enum constant indicating primaries
 * @return A description of the colorspace gamut associated with that enum
 *     constant, or NULL if the constant is unknown to libavutil.
 */
func Av_csp_primaries_desc_from_id(prm AVColorPrimaries) *AVColorPrimariesDesc {
    return (*AVColorPrimariesDesc)(unsafe.Pointer(C.av_csp_primaries_desc_from_id(
        C.enum_AVColorPrimaries(prm))))
}

/**
 * Detects which enum AVColorPrimaries constant corresponds to the given complete
 * gamut description.
 * @see enum AVColorPrimaries
 * @param prm A description of the colorspace gamut
 * @return The enum constant associated with this gamut, or
 *     AVCOL_PRI_UNSPECIFIED if no clear match can be idenitified.
 */
func Av_csp_primaries_id_from_desc(prm *AVColorPrimariesDesc) AVColorPrimaries {
    return AVColorPrimaries(C.av_csp_primaries_id_from_desc(
        (*C.struct_AVColorPrimariesDesc)(unsafe.Pointer(prm))))
}

/**
 * Determine a suitable 'gamma' value to match the supplied
 * AVColorTransferCharacteristic.
 *
 * See Apple Technical Note TN2257 (https://developer.apple.com/library/mac/technotes/tn2257/_index.html)
 *
 * This function returns the gamma exponent for the OETF. For example, sRGB is approximated
 * by gamma 2.2, not by gamma 0.45455.
 *
 * @return Will return an approximation to the simple gamma function matching
 *         the supplied Transfer Characteristic, Will return 0.0 for any
 *         we cannot reasonably match against.
 */
func Av_csp_approximate_trc_gamma(trc AVColorTransferCharacteristic) float64 {
    return float64(C.av_csp_approximate_trc_gamma(
        C.enum_AVColorTransferCharacteristic(trc)))
}

/**
 * Determine the function needed to apply the given
 * AVColorTransferCharacteristic to linear input.
 *
 * The function returned should expect a nominal domain and range of [0.0-1.0]
 * values outside of this range maybe valid depending on the chosen
 * characteristic function.
 *
 * @return Will return pointer to the function matching the
 *         supplied Transfer Characteristic. If unspecified will
 *         return NULL:
 */
func Av_csp_trc_func_from_id(trc AVColorTransferCharacteristic) av_csp_trc_function {
    return av_csp_trc_function(C.av_csp_trc_func_from_id(
        C.enum_AVColorTransferCharacteristic(trc)))
}

/**
 * @}
 */

                         
