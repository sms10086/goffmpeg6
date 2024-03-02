// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
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
//#include <stdint.h>
//#include <stdlib.h>
//#include "libavcodec/defs.h"
import "C"
import (
    "unsafe"
)

const AV_INPUT_BUFFER_PADDING_SIZE = 64
const AV_EF_CRCCHECK = (1<<0)
const AV_EF_BITSTREAM = (1<<1)
const AV_EF_BUFFER = (1<<2)
const AV_EF_EXPLODE = (1<<3)
const AV_EF_IGNORE_ERR = (1<<15)
const AV_EF_CAREFUL = (1<<16)
const AV_EF_COMPLIANT = (1<<17)
const AV_EF_AGGRESSIVE = (1<<18)
const FF_COMPLIANCE_VERY_STRICT = 2
const FF_COMPLIANCE_STRICT = 1
const FF_COMPLIANCE_NORMAL = 0
const FF_COMPLIANCE_UNOFFICIAL = -1
const FF_COMPLIANCE_EXPERIMENTAL = -2
const AV_PROFILE_UNKNOWN = -99
const AV_PROFILE_RESERVED = -100
const AV_PROFILE_AAC_MAIN = 0
const AV_PROFILE_AAC_LOW = 1
const AV_PROFILE_AAC_SSR = 2
const AV_PROFILE_AAC_LTP = 3
const AV_PROFILE_AAC_HE = 4
const AV_PROFILE_AAC_HE_V2 = 28
const AV_PROFILE_AAC_LD = 22
const AV_PROFILE_AAC_ELD = 38
const AV_PROFILE_MPEG2_AAC_LOW = 128
const AV_PROFILE_MPEG2_AAC_HE = 131
const AV_PROFILE_DNXHD = 0
const AV_PROFILE_DNXHR_LB = 1
const AV_PROFILE_DNXHR_SQ = 2
const AV_PROFILE_DNXHR_HQ = 3
const AV_PROFILE_DNXHR_HQX = 4
const AV_PROFILE_DNXHR_444 = 5
const AV_PROFILE_DTS = 20
const AV_PROFILE_DTS_ES = 30
const AV_PROFILE_DTS_96_24 = 40
const AV_PROFILE_DTS_HD_HRA = 50
const AV_PROFILE_DTS_HD_MA = 60
const AV_PROFILE_DTS_EXPRESS = 70
const AV_PROFILE_DTS_HD_MA_X = 61
const AV_PROFILE_DTS_HD_MA_X_IMAX = 62
const AV_PROFILE_EAC3_DDP_ATMOS = 30
const AV_PROFILE_TRUEHD_ATMOS = 30
const AV_PROFILE_MPEG2_422 = 0
const AV_PROFILE_MPEG2_HIGH = 1
const AV_PROFILE_MPEG2_SS = 2
const AV_PROFILE_MPEG2_SNR_SCALABLE = 3
const AV_PROFILE_MPEG2_MAIN = 4
const AV_PROFILE_MPEG2_SIMPLE = 5
const AV_PROFILE_H264_CONSTRAINED = (1<<9)
const AV_PROFILE_H264_INTRA = (1<<11)
const AV_PROFILE_H264_BASELINE = 66
const AV_PROFILE_H264_CONSTRAINED_BASELINE = (66|AV_PROFILE_H264_CONSTRAINED)
const AV_PROFILE_H264_MAIN = 77
const AV_PROFILE_H264_EXTENDED = 88
const AV_PROFILE_H264_HIGH = 100
const AV_PROFILE_H264_HIGH_10 = 110
const AV_PROFILE_H264_HIGH_10_INTRA = (110|AV_PROFILE_H264_INTRA)
const AV_PROFILE_H264_MULTIVIEW_HIGH = 118
const AV_PROFILE_H264_HIGH_422 = 122
const AV_PROFILE_H264_HIGH_422_INTRA = (122|AV_PROFILE_H264_INTRA)
const AV_PROFILE_H264_STEREO_HIGH = 128
const AV_PROFILE_H264_HIGH_444 = 144
const AV_PROFILE_H264_HIGH_444_PREDICTIVE = 244
const AV_PROFILE_H264_HIGH_444_INTRA = (244|AV_PROFILE_H264_INTRA)
const AV_PROFILE_H264_CAVLC_444 = 44
const AV_PROFILE_VC1_SIMPLE = 0
const AV_PROFILE_VC1_MAIN = 1
const AV_PROFILE_VC1_COMPLEX = 2
const AV_PROFILE_VC1_ADVANCED = 3
const AV_PROFILE_MPEG4_SIMPLE = 0
const AV_PROFILE_MPEG4_SIMPLE_SCALABLE = 1
const AV_PROFILE_MPEG4_CORE = 2
const AV_PROFILE_MPEG4_MAIN = 3
const AV_PROFILE_MPEG4_N_BIT = 4
const AV_PROFILE_MPEG4_SCALABLE_TEXTURE = 5
const AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION = 6
const AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE = 7
const AV_PROFILE_MPEG4_HYBRID = 8
const AV_PROFILE_MPEG4_ADVANCED_REAL_TIME = 9
const AV_PROFILE_MPEG4_CORE_SCALABLE = 10
const AV_PROFILE_MPEG4_ADVANCED_CODING = 11
const AV_PROFILE_MPEG4_ADVANCED_CORE = 12
const AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE = 13
const AV_PROFILE_MPEG4_SIMPLE_STUDIO = 14
const AV_PROFILE_MPEG4_ADVANCED_SIMPLE = 15
const AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0 = 1
const AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1 = 2
const AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = 32768
const AV_PROFILE_JPEG2000_DCINEMA_2K = 3
const AV_PROFILE_JPEG2000_DCINEMA_4K = 4
const AV_PROFILE_VP9_0 = 0
const AV_PROFILE_VP9_1 = 1
const AV_PROFILE_VP9_2 = 2
const AV_PROFILE_VP9_3 = 3
const AV_PROFILE_HEVC_MAIN = 1
const AV_PROFILE_HEVC_MAIN_10 = 2
const AV_PROFILE_HEVC_MAIN_STILL_PICTURE = 3
const AV_PROFILE_HEVC_REXT = 4
const AV_PROFILE_HEVC_SCC = 9
const AV_PROFILE_VVC_MAIN_10 = 1
const AV_PROFILE_VVC_MAIN_10_444 = 33
const AV_PROFILE_AV1_MAIN = 0
const AV_PROFILE_AV1_HIGH = 1
const AV_PROFILE_AV1_PROFESSIONAL = 2
const AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT = 0xc0
const AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = 0xc1
const AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT = 0xc2
const AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS = 0xc3
const AV_PROFILE_MJPEG_JPEG_LS = 0xf7
const AV_PROFILE_SBC_MSBC = 1
const AV_PROFILE_PRORES_PROXY = 0
const AV_PROFILE_PRORES_LT = 1
const AV_PROFILE_PRORES_STANDARD = 2
const AV_PROFILE_PRORES_HQ = 3
const AV_PROFILE_PRORES_4444 = 4
const AV_PROFILE_PRORES_XQ = 5
const AV_PROFILE_ARIB_PROFILE_A = 0
const AV_PROFILE_ARIB_PROFILE_C = 1
const AV_PROFILE_KLVA_SYNC = 0
const AV_PROFILE_KLVA_ASYNC = 1
const AV_PROFILE_EVC_BASELINE = 0
const AV_PROFILE_EVC_MAIN = 1
const AV_LEVEL_UNKNOWN = -99


                      
                      

/**
 * @file
 * @ingroup libavc
 * Misc types and constants that do not belong anywhere else.
 */

                   
                   

/**
 * @ingroup lavc_decoding
 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
 * This is mainly needed because some optimized bitstream readers read
 * 32 or 64 bit at once and could read over the end.<br>
 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
 * MPEG bitstreams could cause overread and segfault.
 */


/**
 * Verify checksums embedded in the bitstream (could be of either encoded or
 * decoded data, depending on the format) and print an error message on mismatch.
 * If AV_EF_EXPLODE is also set, a mismatching checksum will result in the
 * decoder/demuxer returning an error.
 */

///< detect bitstream specification deviations
///< detect improper bitstream length
///< abort decoding on minor error detection

///< ignore errors and continue
///< consider things that violate the spec, are fast to calculate and have not been seen in the wild as errors
///< consider all spec non compliances as errors
///< consider things that a sane encoder/muxer should not do as an error

///< Strictly conform to an older more strict version of the spec or reference software.
///< Strictly conform to all the things in the spec no matter what consequences.

///< Allow unofficial extensions
///< Allow nonstandardized experimental things.











































// 8+1; constraint_set1_flag
// 8+3; constraint_set3_flag


























































































type AVFieldOrder C.enum_AVFieldOrder

/**
 * @ingroup lavc_decoding
 */
type AVDiscard C.enum_AVDiscard

type AVAudioServiceType C.enum_AVAudioServiceType

/**
 * Pan Scan area.
 * This specifies the area which should be displayed.
 * Note there may be multiple such areas for one frame.
 */
type AVPanScan C.struct_AVPanScan

/**
 * This structure describes the bitrate properties of an encoded bitstream. It
 * roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
 * parameters for H.264/HEVC.
 */
type AVCPBProperties C.struct_AVCPBProperties

/**
 * Allocate a CPB properties structure and initialize its fields to default
 * values.
 *
 * @param size if non-NULL, the size of the allocated struct will be written
 *             here. This is useful for embedding it in side data.
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_cpb_properties_alloc(size *uint64) *AVCPBProperties {
    return (*AVCPBProperties)(unsafe.Pointer(C.av_cpb_properties_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * This structure supplies correlation between a packet timestamp and a wall clock
 * production time. The definition follows the Producer Reference Time ('prft')
 * as defined in ISO/IEC 14496-12
 */
type AVProducerReferenceTime C.struct_AVProducerReferenceTime

/**
 * Encode extradata length to a buffer. Used by xiph codecs.
 *
 * @param s buffer to write to; must be at least (v/255+1) bytes long
 * @param v size of extradata in bytes
 * @return number of bytes written to the buffer.
 */
func Av_xiphlacing(s *uint8, v uint32) uint32 {
    return uint32(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), C.uint(v)))
}

                        
