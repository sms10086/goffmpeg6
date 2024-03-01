// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * AC-3 parser prototypes
 * Copyright (c) 2003 Fabrice Bellard
 * Copyright (c) 2003 Michael Niedermayer
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

//#cgo pkg-config: libavcodec
//#include <stddef.h>
//#include <stdint.h>
//#include "libavcodec/ac3_parser.h"
import "C"
import (
    "unsafe"
)



                            
                            

                   
                   

/**
 * Extract the bitstream ID and the frame size from AC-3 data.
 */
func Av_ac3_parse_header(buf *uint8, size uint64,
                        bitstream_id *uint8, frame_size *uint16) int32 {
    return int32(C.av_ac3_parse_header((*C.uchar)(unsafe.Pointer(buf)), 
        C.ulonglong(size), (*C.uchar)(unsafe.Pointer(bitstream_id)), 
        (*C.ushort)(unsafe.Pointer(frame_size))))
}


                                 
