// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2014 Vittorio Giovara <vittorio.giovara@gmail.com>
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
//#include "libavutil/display.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_video_display
 * Display matrix
 */

                        
                        

                   

/**
 * @defgroup lavu_video_display Display transformation matrix functions
 * @ingroup lavu_video
 *
 * The display transformation matrix specifies an affine transformation that
 * should be applied to video frames for correct presentation. It is compatible
 * with the matrices stored in the ISO/IEC 14496-12 container format.
 *
 * The data is a 3x3 matrix represented as a 9-element array:
 *
 * @code{.unparsed}
 *                                  | a b u |
 *   (a, b, u, c, d, v, x, y, w) -> | c d v |
 *                                  | x y w |
 * @endcode
 *
 * All numbers are stored in native endianness, as 16.16 fixed-point values,
 * except for u, v and w, which are stored as 2.30 fixed-point values.
 *
 * The transformation maps a point (p, q) in the source (pre-transformation)
 * frame to the point (p', q') in the destination (post-transformation) frame as
 * follows:
 *
 * @code{.unparsed}
 *               | a b u |
 *   (p, q, 1) . | c d v | = z * (p', q', 1)
 *               | x y w |
 * @endcode
 *
 * The transformation can also be more explicitly written in components as
 * follows:
 *
 * @code{.unparsed}
 *   p' = (a * p + c * q + x) / z;
 *   q' = (b * p + d * q + y) / z;
 *   z  =  u * p + v * q + w
 * @endcode
 *
 * @{
 */

/**
 * Extract the rotation component of the transformation matrix.
 *
 * @param matrix the transformation matrix
 * @return the angle (in degrees) by which the transformation rotates the frame
 *         counterclockwise. The angle will be in range [-180.0, 180.0],
 *         or NaN if the matrix is singular.
 *
 * @note floating point numbers are inherently inexact, so callers are
 *       recommended to round the return value to nearest integer before use.
 */
func Av_display_rotation_get(matrix[9] int32) float64 {
    return float64(C.av_display_rotation_get(
        (*C.int)(unsafe.Pointer(&matrix[0]))))
}

/**
 * Initialize a transformation matrix describing a pure clockwise
 * rotation by the specified angle (in degrees).
 *
 * @param[out] matrix a transformation matrix (will be fully overwritten
 *                    by this function)
 * @param angle rotation angle in degrees.
 */
func Av_display_rotation_set(matrix[9] int32, angle float64)  {
    C.av_display_rotation_set((*C.int)(unsafe.Pointer(&matrix[0])), 
        C.double(angle))
}

/**
 * Flip the input matrix horizontally and/or vertically.
 *
 * @param[in,out] matrix a transformation matrix
 * @param hflip whether the matrix should be flipped horizontally
 * @param vflip whether the matrix should be flipped vertically
 */
func Av_display_matrix_flip(matrix[9] int32, hflip int32, vflip int32)  {
    C.av_display_matrix_flip((*C.int)(unsafe.Pointer(&matrix[0])), C.int(hflip), 
        C.int(vflip))
}

/**
 * @}
 */

                             