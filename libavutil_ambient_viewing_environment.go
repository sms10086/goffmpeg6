// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2023 Jan Ekström <jeebjp@gmail.com>
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
//#include "libavutil/frame.h"
//#include "libavutil/rational.h"
//#include "libavutil/ambient_viewing_environment.h"
import "C"
import (
    "unsafe"
)



                                            
                                            

                   
                  
                     

/**
 * Ambient viewing environment metadata as defined by H.274. The values are
 * saved in AVRationals so that they keep their exactness, while allowing for
 * easy access to a double value with f.ex. av_q2d.
 *
 * @note sizeof(AVAmbientViewingEnvironment) is not part of the public ABI, and
 *       it must be allocated using av_ambient_viewing_environment_alloc.
 */
type AVAmbientViewingEnvironment struct {
    Ambient_illuminance AVRational
    Ambient_light_x AVRational
    Ambient_light_y AVRational
}


/**
 * Allocate an AVAmbientViewingEnvironment structure.
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_ambient_viewing_environment_alloc(size *uint64) *AVAmbientViewingEnvironment {
    return (*AVAmbientViewingEnvironment)(unsafe.Pointer(C.av_ambient_viewing_environment_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Allocate and add an AVAmbientViewingEnvironment structure to an existing
 * AVFrame as side data.
 *
 * @return the newly allocated struct, or NULL on failure
 */
func Av_ambient_viewing_environment_create_side_data(frame *AVFrame) *AVAmbientViewingEnvironment {
    return (*AVAmbientViewingEnvironment)(unsafe.Pointer(C.av_ambient_viewing_environment_create_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

                                                 
