// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2007 Michael Niedermayer <michaelni@gmx.at>
 * Copyright (C) 2013 James Almer <jamrial@gmail.com>
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
//#include "libavutil/attributes.h"
//#include "libavutil/ripemd.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_ripemd
 * Public header for RIPEMD hash function implementation.
 */

                       
                       

                   
                   

                       

/**
 * @defgroup lavu_ripemd RIPEMD
 * @ingroup lavu_hash
 * RIPEMD hash function implementation.
 *
 * @{
 */

//extern const int av_ripemd_size;

type AVRIPEMD C.struct_AVRIPEMD

/**
 * Allocate an AVRIPEMD context.
 */
func Av_ripemd_alloc() *AVRIPEMD {
    return (*AVRIPEMD)(unsafe.Pointer(C.av_ripemd_alloc()))
}

/**
 * Initialize RIPEMD hashing.
 *
 * @param context pointer to the function context (of size av_ripemd_size)
 * @param bits    number of bits in digest (128, 160, 256 or 320 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
func Av_ripemd_init(context *AVRIPEMD, bits int32) int32 {
    return int32(C.av_ripemd_init((*C.struct_AVRIPEMD)(unsafe.Pointer(context)), 
        C.int(bits)))
}

/**
 * Update hash value.
 *
 * @param context hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
func Av_ripemd_update(context *AVRIPEMD, data *uint8, len uint64)  {
    C.av_ripemd_update((*C.struct_AVRIPEMD)(unsafe.Pointer(context)), 
        (*C.uchar)(unsafe.Pointer(data)), C.ulonglong(len))
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
func Av_ripemd_final(context *AVRIPEMD, digest *uint8)  {
    C.av_ripemd_final((*C.struct_AVRIPEMD)(unsafe.Pointer(context)), 
        (*C.uchar)(unsafe.Pointer(digest)))
}

/**
 * @}
 */

                            
