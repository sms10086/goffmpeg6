// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * AES-CTR cipher
 * Copyright (c) 2015 Eran Kornblau <erankor at gmail dot com>
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
//#include "libavutil/attributes.h"
//#include "libavutil/aes_ctr.h"
import "C"
import (
    "unsafe"
)

const AES_CTR_KEY_SIZE = 16
const AES_CTR_IV_SIZE = 8


                        
                        

/**
 * @defgroup lavu_aes_ctr AES-CTR
 * @ingroup lavu_crypto
 * @{
 */

                   

                       




type AVAESCTR struct {
}


/**
 * Allocate an AVAESCTR context.
 */
func Av_aes_ctr_alloc() *AVAESCTR {
    return (*AVAESCTR)(unsafe.Pointer(C.av_aes_ctr_alloc()))
}

/**
 * Initialize an AVAESCTR context.
 *
 * @param a The AVAESCTR context to initialize
 * @param key encryption key, must have a length of AES_CTR_KEY_SIZE
 */
func Av_aes_ctr_init(a *AVAESCTR, key *uint8) int32 {
    return int32(C.av_aes_ctr_init((*C.struct_AVAESCTR)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(key))))
}

/**
 * Release an AVAESCTR context.
 *
 * @param a The AVAESCTR context
 */
func Av_aes_ctr_free(a *AVAESCTR)  {
    C.av_aes_ctr_free((*C.struct_AVAESCTR)(unsafe.Pointer(a)))
}

/**
 * Process a buffer using a previously initialized context.
 *
 * @param a The AVAESCTR context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param size the size of src and dst
 */
func Av_aes_ctr_crypt(a *AVAESCTR, dst *uint8, src *uint8, size int32)  {
    C.av_aes_ctr_crypt((*C.struct_AVAESCTR)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(size))
}

/**
 * Get the current iv
 */
func Av_aes_ctr_get_iv(a *AVAESCTR) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_aes_ctr_get_iv(
        (*C.struct_AVAESCTR)(unsafe.Pointer(a)))))
}

/**
 * Generate a random iv
 */
func Av_aes_ctr_set_random_iv(a *AVAESCTR)  {
    C.av_aes_ctr_set_random_iv((*C.struct_AVAESCTR)(unsafe.Pointer(a)))
}

/**
 * Forcefully change the 8-byte iv
 */
func Av_aes_ctr_set_iv(a *AVAESCTR, iv *uint8)  {
    C.av_aes_ctr_set_iv((*C.struct_AVAESCTR)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(iv)))
}

/**
 * Forcefully change the "full" 16-byte iv, including the counter
 */
func Av_aes_ctr_set_full_iv(a *AVAESCTR, iv *uint8)  {
    C.av_aes_ctr_set_full_iv((*C.struct_AVAESCTR)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(iv)))
}

/**
 * Increment the top 64 bit of the iv (performed after each frame)
 */
func Av_aes_ctr_increment_iv(a *AVAESCTR)  {
    C.av_aes_ctr_increment_iv((*C.struct_AVAESCTR)(unsafe.Pointer(a)))
}

/**
 * @}
 */

                             
