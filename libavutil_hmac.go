// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2012 Martin Storsjo
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
//#include "libavutil/hmac.h"
import "C"
import (
    "unsafe"
)



                     
                     

                   

/**
 * @defgroup lavu_hmac HMAC
 * @ingroup lavu_crypto
 * @{
 */

type AVHMACType C.enum_AVHMACType

type AVHMAC C.struct_AVHMAC

/**
 * Allocate an AVHMAC context.
 * @param type The hash function used for the HMAC.
 */
func Av_hmac_alloc(typex AVHMACType) *AVHMAC {
    return (*AVHMAC)(unsafe.Pointer(C.av_hmac_alloc(C.enum_AVHMACType(typex))))
}

/**
 * Free an AVHMAC context.
 * @param ctx The context to free, may be NULL
 */
func Av_hmac_free(ctx *AVHMAC)  {
    C.av_hmac_free((*C.AVHMAC)(unsafe.Pointer(ctx)))
}

/**
 * Initialize an AVHMAC context with an authentication key.
 * @param ctx    The HMAC context
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 */
func Av_hmac_init(ctx *AVHMAC, key *uint8, keylen uint32)  {
    C.av_hmac_init((*C.AVHMAC)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(key)), C.uint(keylen))
}

/**
 * Hash data with the HMAC.
 * @param ctx  The HMAC context
 * @param data The data to hash
 * @param len  The length of the data, in bytes
 */
func Av_hmac_update(ctx *AVHMAC, data *uint8, len uint32)  {
    C.av_hmac_update((*C.AVHMAC)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(data)), C.uint(len))
}

/**
 * Finish hashing and output the HMAC digest.
 * @param ctx    The HMAC context
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
func Av_hmac_final(ctx *AVHMAC, out *uint8, outlen uint32) int32 {
    return int32(C.av_hmac_final((*C.AVHMAC)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(out)), C.uint(outlen)))
}

/**
 * Hash an array of data with a key.
 * @param ctx    The HMAC context
 * @param data   The data to hash
 * @param len    The length of the data, in bytes
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
func Av_hmac_calc(ctx *AVHMAC, data *uint8, len uint32,
                 key *uint8, keylen uint32,
                 out *uint8, outlen uint32) int32 {
    return int32(C.av_hmac_calc((*C.AVHMAC)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(data)), C.uint(len), 
        (*C.uchar)(unsafe.Pointer(key)), C.uint(keylen), 
        (*C.uchar)(unsafe.Pointer(out)), C.uint(outlen)))
}

/**
 * @}
 */

                          
