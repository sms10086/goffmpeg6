// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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
//#include "libavutil/md5.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_md5
 * Public header for MD5 hash function implementation.
 */

                    
                    

                   
                   

                       

/**
 * @defgroup lavu_md5 MD5
 * @ingroup lavu_hash
 * MD5 hash function implementation.
 *
 * @{
 */

//extern const int av_md5_size;

type AVMD5 C.struct_AVMD5

/**
 * Allocate an AVMD5 context.
 */
func Av_md5_alloc() *AVMD5 {
    return (*AVMD5)(unsafe.Pointer(C.av_md5_alloc()))
}

/**
 * Initialize MD5 hashing.
 *
 * @param ctx pointer to the function context (of size av_md5_size)
 */
func Av_md5_init(ctx *AVMD5)  {
    C.av_md5_init((*C.struct_AVMD5)(unsafe.Pointer(ctx)))
}

/**
 * Update hash value.
 *
 * @param ctx hash function context
 * @param src input data to update hash with
 * @param len input data length
 */
func Av_md5_update(ctx *AVMD5, src *uint8, len uint64)  {
    C.av_md5_update((*C.struct_AVMD5)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(src)), C.ulonglong(len))
}

/**
 * Finish hashing and output digest value.
 *
 * @param ctx hash function context
 * @param dst buffer where output digest value is stored
 */
func Av_md5_final(ctx *AVMD5, dst *uint8)  {
    C.av_md5_final((*C.struct_AVMD5)(unsafe.Pointer(ctx)), 
        (*C.uchar)(unsafe.Pointer(dst)))
}

/**
 * Hash an array of data.
 *
 * @param dst The output buffer to write the digest into
 * @param src The data to hash
 * @param len The length of the data, in bytes
 */
func Av_md5_sum(dst *uint8, src *uint8, len uint64)  {
    C.av_md5_sum((*C.uchar)(unsafe.Pointer(dst)), 
        (*C.uchar)(unsafe.Pointer(src)), C.ulonglong(len))
}

/**
 * @}
 */

                         
