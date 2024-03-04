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
//#include <stdint.h>
//#include <stddef.h>
//#include "libavutil/attributes.h"
//#include "libavutil/crc.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * @ingroup lavu_crc32
 * Public header for CRC hash function implementation.
 */

                    
                    

                   
                   
                       

/**
 * @defgroup lavu_crc32 CRC
 * @ingroup lavu_hash
 * CRC (Cyclic Redundancy Check) hash function implementation.
 *
 * This module supports numerous CRC polynomials, in addition to the most
 * widely used CRC-32-IEEE. See @ref AVCRCId for a list of available
 * polynomials.
 *
 * @{
 */

type AVCRC uint32

type AVCRCId int

/**
 * Initialize a CRC table.
 * @param ctx must be an array of size sizeof(AVCRC)*257 or sizeof(AVCRC)*1024
 * @param le If 1, the lowest bit represents the coefficient for the highest
 *           exponent of the corresponding polynomial (both for poly and
 *           actual CRC).
 *           If 0, you must swap the CRC parameter and the result of av_crc
 *           if you need the standard representation (can be simplified in
 *           most cases to e.g. bswap16):
 *           av_bswap32(crc << (32-bits))
 * @param bits number of bits for the CRC
 * @param poly generator polynomial without the x**bits coefficient, in the
 *             representation as specified by le
 * @param ctx_size size of ctx in bytes
 * @return <0 on failure
 */
func Av_crc_init(ctx *AVCRC, le int32, bits int32, poly uint32, ctx_size int32) int32 {
    return int32(C.av_crc_init((*C.AVCRC)(unsafe.Pointer(ctx)), C.int(le), 
        C.int(bits), C.uint(poly), C.int(ctx_size)))
}

/**
 * Get an initialized standard CRC table.
 * @param crc_id ID of a standard CRC
 * @return a pointer to the CRC table or NULL on failure
 */
func Av_crc_get_table(crc_id AVCRCId) *AVCRC {
    return (*AVCRC)(unsafe.Pointer(C.av_crc_get_table(C.AVCRCId(crc_id))))
}

/**
 * Calculate the CRC of a block.
 * @param ctx initialized AVCRC array (see av_crc_init())
 * @param crc CRC of previous blocks if any or initial value for CRC
 * @param buffer buffer whose CRC to calculate
 * @param length length of the buffer
 * @return CRC updated with the data from the given block
 *
 * @see av_crc_init() "le" parameter
 */
func Av_crc(ctx *AVCRC, crc uint32,
                buffer *uint8, length uint64)  uint32 {
    return uint32(C.av_crc((*C.AVCRC)(unsafe.Pointer(ctx)), C.uint(crc), 
        (*C.uchar)(unsafe.Pointer(buffer)), C.ulonglong(length)))
}

/**
 * @}
 */

                         
