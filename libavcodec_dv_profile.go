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

//#cgo pkg-config: libavcodec libavutil
//#include <stdint.h>
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavcodec/dv_profile.h"
import "C"
import (
    "unsafe"
)

const DV_PROFILE_BYTES = 480


                            
                            

                   

                             
                               

/* minimum number of bytes to read from a DV stream in order to
 * determine the profile */
/* 6 DIF blocks */


/*
 * AVDVProfile is used to express the differences between various
 * DV flavors. For now it's primarily used for differentiating
 * 525/60 and 625/50, but the plans are to use it for various
 * DV specs as well (e.g. SMPTE314M vs. IEC 61834).
 */
type AVDVProfile struct {
    Dsf int32
    Video_stype int32
    Frame_size int32
    Difseg_size int32
    N_difchan int32
    Time_base AVRational
    Ltc_divisor int32
    Height int32
    Width int32
    Sar [2]AVRational
    Pix_fmt AVPixelFormat
    Bpm int32
    Block_sizes *uint8
    Audio_stride int32
    Audio_min_samples [3]int32
    Audio_samples_dist [5]int32
    Audio_shuffle* [9]uint8
}


/**
 * Get a DV profile for the provided compressed frame.
 *
 * @param sys the profile used for the previous frame, may be NULL
 * @param frame the compressed data buffer
 * @param buf_size size of the buffer in bytes
 * @return the DV profile for the supplied data or NULL on failure
 */
func Av_dv_frame_profile(sys *AVDVProfile,
                                       frame *uint8, buf_size uint32) *AVDVProfile {
    return (*AVDVProfile)(unsafe.Pointer(C.av_dv_frame_profile(
        (*C.struct_AVDVProfile)(unsafe.Pointer(sys)), 
        (*C.uchar)(unsafe.Pointer(frame)), C.unsigned(buf_size))))
}

/**
 * Get a DV profile for the provided stream parameters.
 */
func Av_dv_codec_profile(width int32, height int32, pix_fmt AVPixelFormat) *AVDVProfile {
    return (*AVDVProfile)(unsafe.Pointer(C.av_dv_codec_profile(C.int(width), 
        C.int(height), C.enum_AVPixelFormat(pix_fmt))))
}

/**
 * Get a DV profile for the provided stream parameters.
 * The frame rate is used as a best-effort parameter.
 */
func Av_dv_codec_profile2(width int32, height int32, pix_fmt AVPixelFormat, frame_rate AVRational) *AVDVProfile {
    return (*AVDVProfile)(unsafe.Pointer(C.av_dv_codec_profile2(C.int(width), 
        C.int(height), C.enum_AVPixelFormat(pix_fmt), 
        *(*C.struct_AVRational)(unsafe.Pointer(&frame_rate)))))
}

                                 
