// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * RC4 encryption/decryption/pseudo-random number generator
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
//#include "libavutil/rc4.h"
import "C"
import (
    "unsafe"
)



                    
                    

                   

/**
 * @defgroup lavu_rc4 RC4
 * @ingroup lavu_crypto
 * @{
 */

type AVRC4 C.struct_AVRC4

/**
 * Allocate an AVRC4 context.
 */
func Av_rc4_alloc() *AVRC4 {
    return (*AVRC4)(unsafe.Pointer(C.av_rc4_alloc()))
}

/**
 * @brief Initializes an AVRC4 context.
 *
 * @param d pointer to the AVRC4 context
 * @param key buffer containig the key
 * @param key_bits must be a multiple of 8
 * @param decrypt 0 for encryption, 1 for decryption, currently has no effect
 * @return zero on success, negative value otherwise
 */
func Av_rc4_init(d *AVRC4, key *uint8, key_bits int32, decrypt int32) int32 {
    return int32(C.av_rc4_init((*C.struct_AVRC4)(unsafe.Pointer(d)), 
        (*C.uchar)(unsafe.Pointer(key)), C.int(key_bits), C.int(decrypt)))
}

/**
 * @brief Encrypts / decrypts using the RC4 algorithm.
 *
 * @param d pointer to the AVRC4 context
 * @param count number of bytes
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst, may be NULL
 * @param iv not (yet) used for RC4, should be NULL
 * @param decrypt 0 for encryption, 1 for decryption, not (yet) used
 */
func Av_rc4_crypt(d *AVRC4, dst *uint8, src *uint8, count int32, iv *uint8, decrypt int32)  {
    C.av_rc4_crypt((*C.struct_AVRC4)(unsafe.Pointer(d)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @}
 */

                         
