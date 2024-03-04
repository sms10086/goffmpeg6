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
//#include "libavutil/sha512.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_sha512
 * Public header for SHA-512 implementation.
 */

                       
                       

                   
                   

                       

/**
 * @defgroup lavu_sha512 SHA-512
 * @ingroup lavu_hash
 * SHA-512 (Secure Hash Algorithm) hash function implementations.
 *
 * This module supports the following SHA-2 hash functions:
 *
 * - SHA-512/224: 224 bits
 * - SHA-512/256: 256 bits
 * - SHA-384: 384 bits
 * - SHA-512: 512 bits
 *
 * @see For SHA-1, SHA-256, and variants thereof, see @ref lavu_sha.
 *
 * @{
 */

//extern const int av_sha512_size;

type AVSHA512 C.struct_AVSHA512

/**
 * Allocate an AVSHA512 context.
 */
func Av_sha512_alloc() *AVSHA512 {
    return (*AVSHA512)(unsafe.Pointer(C.av_sha512_alloc()))
}

/**
 * Initialize SHA-2 512 hashing.
 *
 * @param context pointer to the function context (of size av_sha512_size)
 * @param bits    number of bits in digest (224, 256, 384 or 512 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
func Av_sha512_init(context *AVSHA512, bits int32) int32 {
    return int32(C.av_sha512_init((*C.struct_AVSHA512)(unsafe.Pointer(context)), 
        C.int(bits)))
}

/**
 * Update hash value.
 *
 * @param context hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
func Av_sha512_update(context *AVSHA512, data *uint8, len uint64)  {
    C.av_sha512_update((*C.struct_AVSHA512)(unsafe.Pointer(context)), 
        (*C.uchar)(unsafe.Pointer(data)), C.ulonglong(len))
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
func Av_sha512_final(context *AVSHA512, digest *uint8)  {
    C.av_sha512_final((*C.struct_AVSHA512)(unsafe.Pointer(context)), 
        (*C.uchar)(unsafe.Pointer(digest)))
}

/**
 * @}
 */

                            
