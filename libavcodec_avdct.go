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

//#cgo pkg-config: libavcodec libavutil
//#include "libavutil/opt.h"
//#include "libavcodec/avdct.h"
import "C"
import (
    "unsafe"
)



                       
                       

                          

/**
 * AVDCT context.
 * @note function pointers can be NULL if the specific features have been
 *       disabled at build time.
 */
type AVDCT struct {
    Av_class *AVClass
    Idct uintptr
    Idct_permutation [64]uint8
    Fdct uintptr
    Dct_algo int32
    Idct_algo int32
    Get_pixels uintptr
    Bits_per_sample int32
    Get_pixels_unaligned uintptr
}


/**
 * Allocates a AVDCT context.
 * This needs to be initialized with avcodec_dct_init() after optionally
 * configuring it with AVOptions.
 *
 * To free it use av_free()
 */
func Avcodec_dct_alloc() *AVDCT {
    return (*AVDCT)(unsafe.Pointer(C.avcodec_dct_alloc()))
}
func Avcodec_dct_init(_var0 *AVDCT) int32 {
    return int32(C.avcodec_dct_init((*C.struct_AVDCT)(unsafe.Pointer(_var0))))
}

func Avcodec_dct_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avcodec_dct_get_class()))
}

                            
