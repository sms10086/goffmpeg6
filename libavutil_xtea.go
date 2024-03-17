// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * A 32-bit implementation of the XTEA algorithm
 * Copyright (c) 2012 Samuel Pitoiset
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
//#include "libavutil/xtea.h"
import "C"
import (
    "unsafe"
)



                     
                     

                   

/**
 * @file
 * @brief Public header for libavutil XTEA algorithm
 * @defgroup lavu_xtea XTEA
 * @ingroup lavu_crypto
 * @{
 */

type AVXTEA struct {
    Key [16]uint32
}


/**
 * Allocate an AVXTEA context.
 */
func Av_xtea_alloc() *AVXTEA {
    return (*AVXTEA)(unsafe.Pointer(C.av_xtea_alloc()))
}

/**
 * Initialize an AVXTEA context.
 *
 * @param ctx an AVXTEA context
 * @param key a key of 16 bytes used for encryption/decryption,
 *            interpreted as big endian 32 bit numbers
 */
func Av_xtea_init(ctx *AVXTEA, key [16]uint8)  {
    C.av_xtea_init((*C.struct_AVXTEA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(&key[0])))
}

/**
 * Initialize an AVXTEA context.
 *
 * @param ctx an AVXTEA context
 * @param key a key of 16 bytes used for encryption/decryption,
 *            interpreted as little endian 32 bit numbers
 */
func Av_xtea_le_init(ctx *AVXTEA, key [16]uint8)  {
    C.av_xtea_le_init((*C.struct_AVXTEA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(&key[0])))
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context,
 * in big endian format.
 *
 * @param ctx an AVXTEA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_xtea_crypt(ctx *AVXTEA, dst *uint8, src *uint8,
                   count int32, iv *uint8, decrypt int32)  {
    C.av_xtea_crypt((*C.struct_AVXTEA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context,
 * in little endian format.
 *
 * @param ctx an AVXTEA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_xtea_le_crypt(ctx *AVXTEA, dst *uint8, src *uint8,
                      count int32, iv *uint8, decrypt int32)  {
    C.av_xtea_le_crypt((*C.struct_AVXTEA)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @}
 */

                          
