// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2007 Michael Niedermayer <michaelni@gmx.at>
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
//#include "libavutil/sha.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_sha
 * Public header for SHA-1 & SHA-256 hash function implementations.
 */

                    
                    

                   
                   

                       

/**
 * @defgroup lavu_sha SHA
 * @ingroup lavu_hash
 * SHA-1 and SHA-256 (Secure Hash Algorithm) hash function implementations.
 *
 * This module supports the following SHA hash functions:
 *
 * - SHA-1: 160 bits
 * - SHA-224: 224 bits, as a variant of SHA-2
 * - SHA-256: 256 bits, as a variant of SHA-2
 *
 * @see For SHA-384, SHA-512, and variants thereof, see @ref lavu_sha512.
 *
 * @{
 */

//extern const int av_sha_size;

type AVSHA C.struct_AVSHA

/**
 * Allocate an AVSHA context.
 */
func Av_sha_alloc() *AVSHA {
    return (*AVSHA)(unsafe.Pointer(C.av_sha_alloc()))
}

/**
 * Initialize SHA-1 or SHA-2 hashing.
 *
 * @param context pointer to the function context (of size av_sha_size)
 * @param bits    number of bits in digest (SHA-1 - 160 bits, SHA-2 224 or 256 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
func Av_sha_init(context *AVSHA, bits int32) int32 {
    return int32(C.av_sha_init((*C.struct_AVSHA)(unsafe.Pointer(context)), 
        C.int(bits)))
}

/**
 * Update hash value.
 *
 * @param ctx     hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
func Av_sha_update(ctx *AVSHA, data *uint8, len uint64)  {
    C.av_sha_update((*C.struct_AVSHA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(data)), C.ulonglong(len))
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
func Av_sha_final(context *AVSHA, digest *uint8)  {
    C.av_sha_final((*C.struct_AVSHA)(unsafe.Pointer(context)), 
        (*C.uchar)(unsafe.Pointer(digest)))
}

/**
 * @}
 */

                         
