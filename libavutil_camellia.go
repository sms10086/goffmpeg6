// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * An implementation of the CAMELLIA algorithm as mentioned in RFC3713
 * Copyright (c) 2014 Supraja Meedinti
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
//#include "libavutil/camellia.h"
import "C"
import (
    "unsafe"
)



                         
                         

                   


/**
  * @file
  * @brief Public header for libavutil CAMELLIA algorithm
  * @defgroup lavu_camellia CAMELLIA
  * @ingroup lavu_crypto
  * @{
  */

//extern const int av_camellia_size;

type AVCAMELLIA struct {
}


/**
  * Allocate an AVCAMELLIA context
  * To free the struct: av_free(ptr)
  */
func Av_camellia_alloc() *AVCAMELLIA {
    return (*AVCAMELLIA)(unsafe.Pointer(C.av_camellia_alloc()))
}

/**
  * Initialize an AVCAMELLIA context.
  *
  * @param ctx an AVCAMELLIA context
  * @param key a key of 16, 24, 32 bytes used for encryption/decryption
  * @param key_bits number of keybits: possible are 128, 192, 256
 */
func Av_camellia_init(ctx *AVCAMELLIA, key *uint8, key_bits int32) int32 {
    return int32(C.av_camellia_init((*C.struct_AVCAMELLIA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(key)), C.int(key_bits)))
}

/**
  * Encrypt or decrypt a buffer using a previously initialized context
  *
  * @param ctx an AVCAMELLIA context
  * @param dst destination array, can be equal to src
  * @param src source array, can be equal to dst
  * @param count number of 16 byte blocks
  * @param iv initialization vector for CBC mode, NULL for ECB mode
  * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_camellia_crypt(ctx *AVCAMELLIA, dst *uint8, src *uint8, count int32, iv *uint8, decrypt int32)  {
    C.av_camellia_crypt((*C.struct_AVCAMELLIA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @}
 */
                              
