// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Lagged Fibonacci PRNG
 * Copyright (c) 2008 Michael Niedermayer
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
//#include "libavutil/lfg.h"
import "C"
import (
    "unsafe"
)



                    
                    

                   

/**
 * Context structure for the Lagged Fibonacci PRNG.
 * The exact layout, types and content of this struct may change and should
 * not be accessed directly. Only its `sizeof()` is guaranteed to stay the same
 * to allow easy instanciation.
 */
type AVLFG C.struct_AVLFG

func Av_lfg_init(c *AVLFG, seed uint32)  {
    C.av_lfg_init((*C.AVLFG)(unsafe.Pointer(c)), C.uint(seed))
}

/**
 * Seed the state of the ALFG using binary data.
 *
 * @return 0 on success, negative value (AVERROR) on failure.
 */
func Av_lfg_init_from_data(c *AVLFG, data *uint8, length uint32) int32 {
    return int32(C.av_lfg_init_from_data((*C.AVLFG)(unsafe.Pointer(c)), 
        (*C.uchar)(unsafe.Pointer(data)), C.uint(length)))
}

/**
 * Get the next random unsigned 32-bit number using an ALFG.
 *
 * Please also consider a simple LCG like state= state*1664525+1013904223,
 * it may be good enough and faster for your specific use case.
 */
// av_lfg_get(AVLFG*c)

/**
 * Get the next random unsigned 32-bit number using a MLFG.
 *
 * Please also consider av_lfg_get() above, it is faster.
 */
// av_mlfg_get(AVLFG*c)

/**
 * Get the next two numbers generated by a Box-Muller Gaussian
 * generator using the random numbers issued by lfg.
 *
 * @param lfg pointer to the contex structure
 * @param out array where the two generated numbers are placed
 */
func Av_bmg_get(lfg *AVLFG, out[2] float64)  {
    C.av_bmg_get((*C.AVLFG)(unsafe.Pointer(lfg)), 
        (*C.double)(unsafe.Pointer(&out[0])))
}

                         
