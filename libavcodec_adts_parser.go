// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
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

//#cgo pkg-config: libavcodec
//#include <stddef.h>
//#include <stdint.h>
//#include "libavcodec/adts_parser.h"
import "C"
import (
    "unsafe"
)

const AV_AAC_ADTS_HEADER_SIZE = 7


                             
                             

                   
                   



/**
 * Extract the number of samples and frames from AAC data.
 * @param[in]  buf     pointer to AAC data buffer
 * @param[out] samples Pointer to where number of samples is written
 * @param[out] frames  Pointer to where number of frames is written
 * @return Returns 0 on success, error code on failure.
 */
func Av_adts_header_parse(buf *uint8, samples *uint32,
                         frames *uint8) int32 {
    return int32(C.av_adts_header_parse((*C.uchar)(unsafe.Pointer(buf)), 
        (*C.uint)(unsafe.Pointer(samples)), (*C.uchar)(unsafe.Pointer(frames))))
}

                                  
