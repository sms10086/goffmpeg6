// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2007 Michael Niedermayer <michaelni@gmx.at>
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
//#include "libavutil/aes.h"
import "C"
import (
    "unsafe"
)



                    
                    

                   

                       

/**
 * @defgroup lavu_aes AES
 * @ingroup lavu_crypto
 * @{
 */

//extern const int av_aes_size;

type AVAES C.struct_AVAES

/**
 * Allocate an AVAES context.
 */
func Av_aes_alloc() *AVAES {
    return (*AVAES)(unsafe.Pointer(C.av_aes_alloc()))
}

/**
 * Initialize an AVAES context.
 *
 * @param a The AVAES context
 * @param key Pointer to the key
 * @param key_bits 128, 192 or 256
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_aes_init(a *AVAES, key *uint8, key_bits int32, decrypt int32) int32 {
    return int32(C.av_aes_init((*C.struct_AVAES)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(key)), C.int(key_bits), C.int(decrypt)))
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param a The AVAES context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 16 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
func Av_aes_crypt(a *AVAES, dst *uint8, src *uint8, count int32, iv *uint8, decrypt int32)  {
    C.av_aes_crypt((*C.struct_AVAES)(unsafe.Pointer(a)), 
        (*C.uchar)(unsafe.Pointer(dst)), (*C.uchar)(unsafe.Pointer(src)), 
        C.int(count), (*C.uchar)(unsafe.Pointer(iv)), C.int(decrypt))
}

/**
 * @}
 */

                         
