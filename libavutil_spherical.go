// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2016 Vittorio Giovara <vittorio.giovara@gmail.com>
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
//#include "libavutil/spherical.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_video_spherical
 * Spherical video
 */

                          
                          

                   
                   

/**
 * @defgroup lavu_video_spherical Spherical video mapping
 * @ingroup lavu_video
 *
 * A spherical video file contains surfaces that need to be mapped onto a
 * sphere. Depending on how the frame was converted, a different distortion
 * transformation or surface recomposition function needs to be applied before
 * the video should be mapped and displayed.
 * @{
 */

/**
 * Projection of the video surface(s) on a sphere.
 */
type AVSphericalProjection C.enum_AVSphericalProjection

/**
 * This structure describes how to handle spherical videos, outlining
 * information about projection, initial layout, and any other view modifier.
 *
 * @note The struct must be allocated with av_spherical_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVSphericalMapping C.struct_AVSphericalMapping

/**
 * Allocate a AVSphericalVideo structure and initialize its fields to default
 * values.
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_spherical_alloc(size *uint64) *AVSphericalMapping {
    return (*AVSphericalMapping)(unsafe.Pointer(C.av_spherical_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Convert the @ref bounding fields from an AVSphericalVideo
 * from 0.32 fixed point to pixels.
 *
 * @param map    The AVSphericalVideo map to read bound values from.
 * @param width  Width of the current frame or stream.
 * @param height Height of the current frame or stream.
 * @param left   Pixels from the left edge.
 * @param top    Pixels from the top edge.
 * @param right  Pixels from the right edge.
 * @param bottom Pixels from the bottom edge.
 */
func Av_spherical_tile_bounds(mapx *AVSphericalMapping,
                              width uint64, height uint64,
                              left *uint64, top *uint64,
                              right *uint64, bottom *uint64)  {
    C.av_spherical_tile_bounds((*C.AVSphericalMapping)(unsafe.Pointer(mapx)), 
        C.ulonglong(width), C.ulonglong(height), 
        (*C.ulonglong)(unsafe.Pointer(left)), 
        (*C.ulonglong)(unsafe.Pointer(top)), 
        (*C.ulonglong)(unsafe.Pointer(right)), 
        (*C.ulonglong)(unsafe.Pointer(bottom)))
}

/**
 * Provide a human-readable name of a given AVSphericalProjection.
 *
 * @param projection The input AVSphericalProjection.
 *
 * @return The name of the AVSphericalProjection, or "unknown".
 */
func Av_spherical_projection_name(projection AVSphericalProjection) string {
    return C.GoString(C.av_spherical_projection_name(
        C.enum_AVSphericalProjection(projection)))
}

/**
 * Get the AVSphericalProjection form a human-readable name.
 *
 * @param name The input string.
 *
 * @return The AVSphericalProjection value, or -1 if not found.
 */
func Av_spherical_from_name(name *byte) int32 {
    return int32(C.av_spherical_from_name((*C.char)(unsafe.Pointer(name))))
}
/**
 * @}
 */

                               
