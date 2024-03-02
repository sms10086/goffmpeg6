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
//#include <stdint.h>
//#include "libavcodec/vorbis_parser.h"
import "C"
import (
    "unsafe"
)

const VORBIS_FLAG_HEADER = 0x00000001
const VORBIS_FLAG_COMMENT = 0x00000002
const VORBIS_FLAG_SETUP = 0x00000004


/**
 * @file
 * A public API for Vorbis parsing
 *
 * Determines the duration for each packet.
 */

                               
                               

                   

type AVVorbisParseContext C.struct_AVVorbisParseContext

/**
 * Allocate and initialize the Vorbis parser using headers in the extradata.
 */
func Av_vorbis_parse_init(extradata *uint8,
                                           extradata_size int32) *AVVorbisParseContext {
    return (*AVVorbisParseContext)(unsafe.Pointer(C.av_vorbis_parse_init(
        (*C.uchar)(unsafe.Pointer(extradata)), C.int(extradata_size))))
}

/**
 * Free the parser and everything associated with it.
 */
func Av_vorbis_parse_free(s **AVVorbisParseContext)  {
    C.av_vorbis_parse_free((**C.AVVorbisParseContext)(unsafe.Pointer(s)))
}





/**
 * Get the duration for a Vorbis packet.
 *
 * If @p flags is @c NULL,
 * special frames are considered invalid.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 * @param flags    flags for special frames
 */
func Av_vorbis_parse_frame_flags(s *AVVorbisParseContext, buf *uint8,
                                buf_size int32, flags *int32) int32 {
    return int32(C.av_vorbis_parse_frame_flags(
        (*C.AVVorbisParseContext)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size), 
        (*C.int)(unsafe.Pointer(flags))))
}

/**
 * Get the duration for a Vorbis packet.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 */
func Av_vorbis_parse_frame(s *AVVorbisParseContext, buf *uint8,
                          buf_size int32) int32 {
    return int32(C.av_vorbis_parse_frame(
        (*C.AVVorbisParseContext)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size)))
}

func Av_vorbis_parse_reset(s *AVVorbisParseContext)  {
    C.av_vorbis_parse_reset((*C.AVVorbisParseContext)(unsafe.Pointer(s)))
}

                                    
