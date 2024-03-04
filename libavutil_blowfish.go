// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Blowfish algorithm
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
//#include "libavutil/blowfish.h"
import "C"
import (
    "unsafe"
)

const AV_BF_ROUNDS = 16


                         
                         

                   

/**
 * @defgroup lavu_blowfish Blowfish
 * @ingroup lavu_crypto
 * @{
 */



type AVBlowfish C.struct_AVBlowfish

/**
 * Allocate an AVBlowfish context.
 */
func Av_blowfish_alloc() *AVBlowfish {
    return (*AVBlowfish)(unsafe.Pointer(C.av_blowfish_alloc()))
}

/**
 * Initialize an AVBlowfish context.
 *
 * @param ctx an AVBlowfish context
 * @param key a key
 * @param key_len length of the key
 */
func Av_blowfish_init(ctx *AVBlowfish, key *uint8, key_len int32)  {
    C.av_blowfish_init((*C.struct_AVBlowfish)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(key)), C.int(key_len))
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param xl left four bytes halves of input to be encrypted
 * @param xr right four bytes halves of input to be encrypted
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_blowfish_crypt_ecb(ctx *AVBlowfish, xl *uint32, xr *uint32,
                           decrypt int32)  {
    C.av_blowfish_crypt_ecb((*C.struct_AVBlowfish)(unsafe.Pointer(ctx)), 
        (*C.uint)(unsafe.Pointer(xl)), (*C.uint)(unsafe.Pointer(xr)), 
        C.int(decrypt))
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_blowfish_crypt(ctx *AVBlowfish, dst *uint8, src *uint8,
                       count int32, iv *uint8, decrypt int32)  {
    C.av_blowfish_crypt((*C.struct_AVBlowfish)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @}
 */

                              
