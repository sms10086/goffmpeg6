// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * DES encryption/decryption
 * Copyright (c) 2007 Reimar Doeffinger
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
//#include "libavutil/des.h"
import "C"
import (
    "unsafe"
)



                    
                    

                   

/**
 * @defgroup lavu_des DES
 * @ingroup lavu_crypto
 * @{
 */

type AVDES C.struct_AVDES

/**
 * Allocate an AVDES context.
 */
func Av_des_alloc() *AVDES {
    return (*AVDES)(unsafe.Pointer(C.av_des_alloc()))
}

/**
 * @brief Initializes an AVDES context.
 *
 * @param d pointer to a AVDES structure to initialize
 * @param key pointer to the key to use
 * @param key_bits must be 64 or 192
 * @param decrypt 0 for encryption/CBC-MAC, 1 for decryption
 * @return zero on success, negative value otherwise
 */
func Av_des_init(d *AVDES, key *uint8, key_bits int32, decrypt int32) int32 {
    return int32(C.av_des_init((*C.struct_AVDES)(unsafe.Pointer(d)), 
        (*C.uchar)(unsafe.Pointer(key)), C.int(key_bits), C.int(decrypt)))
}

/**
 * @brief Encrypts / decrypts using the DES algorithm.
 *
 * @param d pointer to the AVDES structure
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used,
 *           must be 8-byte aligned
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_des_crypt(d *AVDES, dst *uint8, src *uint8, count int32, iv *uint8, decrypt int32)  {
    C.av_des_crypt((*C.struct_AVDES)(unsafe.Pointer(d)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @brief Calculates CBC-MAC using the DES algorithm.
 *
 * @param d pointer to the AVDES structure
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 * @param count number of 8 byte blocks
 */
func Av_des_mac(d *AVDES, dst *uint8, src *uint8, count int32)  {
    C.av_des_mac((*C.struct_AVDES)(unsafe.Pointer(d)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count))
}

/**
 * @}
 */

                         
