// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2007 Marco Gerards <marco@gnu.org>
 * Copyright (C) 2009 David Conrad
 * Copyright (C) 2011 Jordi Ortiz
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

//#cgo pkg-config: libavcodec libavutil
//#include <stddef.h>
//#include <stdint.h>
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavcodec/dirac.h"
import "C"
import (
    "unsafe"
)

const MAX_DWT_LEVELS = 5


                       
                       

/**
 * @file
 * Interface to Dirac Decoder/Encoder
 * @author Marco Gerards <marco@gnu.org>
 * @author David Conrad
 * @author Jordi Ortiz
 */

                   
                   

                             
                               

/**
 * The spec limits the number of wavelet decompositions to 4 for both
 * level 1 (VC-2) and 128 (long-gop default).
 * 5 decompositions is the maximum before >16-bit buffers are needed.
 * Schroedinger allows this for DD 9,7 and 13,7 wavelets only, limiting
 * the others to 4 decompositions (or 3 for the fidelity filter).
 *
 * We use this instead of MAX_DECOMPOSITIONS to save some memory.
 */


/**
 * Parse code values:
 *
 * Dirac Specification ->
 * 9.6.1  Table 9.1
 *
 * VC-2 Specification  ->
 * 10.4.1 Table 10.1
 */

type DiracParseCodes C.enum_DiracParseCodes

type DiracVersionInfo C.struct_DiracVersionInfo

type AVDiracSeqHeader C.struct_AVDiracSeqHeader

/**
 * Parse a Dirac sequence header.
 *
 * @param dsh this function will allocate and fill an AVDiracSeqHeader struct
 *            and write it into this pointer. The caller must free it with
 *            av_free().
 * @param buf the data buffer
 * @param buf_size the size of the data buffer in bytes
 * @param log_ctx if non-NULL, this function will log errors here
 * @return 0 on success, a negative AVERROR code on failure
 */
func Av_dirac_parse_sequence_header(dsh **AVDiracSeqHeader,
                                   buf *uint8, buf_size uint64,
                                   log_ctx unsafe.Pointer) int32 {
    return int32(C.av_dirac_parse_sequence_header(
        (**C.AVDiracSeqHeader)(unsafe.Pointer(dsh)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.ulonglong(buf_size), log_ctx))
}

                            
