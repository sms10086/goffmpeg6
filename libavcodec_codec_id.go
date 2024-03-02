// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Codec IDs
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
//#include "libavutil/avutil.h"
//#include "libavutil/samplefmt.h"
//#include "libavcodec/version_major.h"
//#include "libavcodec/codec_id.h"
import "C"

//const AV_CODEC_ID_IFF_BYTERUN1 = AV_CODEC_ID_IFF_ILBM
//const AV_CODEC_ID_H265 = AV_CODEC_ID_HEVC
//const AV_CODEC_ID_H266 = AV_CODEC_ID_VVC


                          
                          

                             
                                

                          

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * Identify the syntax and semantics of the bitstream.
 * The principle is roughly:
 * Two decoders with the same ID can decode the same streams.
 * Two encoders with the same ID can encode compatible streams.
 * There may be slight deviations from the principle due to implementation
 * details.
 *
 * If you add a codec ID to this list, add it so that
 * 1. no value of an existing codec ID changes (that would break ABI),
 * 2. it is as close as possible to similar codecs
 *
 * After adding new codec IDs, do not forget to add an entry to the codec
 * descriptor list and bump libavcodec minor version.
 */
type AVCodecID C.enum_AVCodecID

/**
 * Get the type of the given codec.
 */
func Avcodec_get_type(codec_id AVCodecID) AVMediaType {
    return AVMediaType(C.avcodec_get_type(C.enum_AVCodecID(codec_id)))
}

/**
 * Get the name of a codec.
 * @return  a static string identifying the codec; never NULL
 */
func Avcodec_get_name(id AVCodecID) string {
    return C.GoString(C.avcodec_get_name(C.enum_AVCodecID(id)))
}

/**
 * Return codec bits per sample.
 *
 * @param[in] codec_id the codec
 * @return Number of bits per sample or zero if unknown for the given codec.
 */
func Av_get_bits_per_sample(codec_id AVCodecID) int32 {
    return int32(C.av_get_bits_per_sample(C.enum_AVCodecID(codec_id)))
}

/**
 * Return codec bits per sample.
 * Only return non-zero if the bits per sample is exactly correct, not an
 * approximation.
 *
 * @param[in] codec_id the codec
 * @return Number of bits per sample or zero if unknown for the given codec.
 */
func Av_get_exact_bits_per_sample(codec_id AVCodecID) int32 {
    return int32(C.av_get_exact_bits_per_sample(C.enum_AVCodecID(codec_id)))
}

/**
 * Return a name for the specified profile, if available.
 *
 * @param codec_id the ID of the codec to which the requested profile belongs
 * @param profile the profile value for which a name is requested
 * @return A name for the profile if found, NULL otherwise.
 *
 * @note unlike av_get_profile_name(), which searches a list of profiles
 *       supported by a specific decoder or encoder implementation, this
 *       function searches the list of profiles from the AVCodecDescriptor
 */
func Avcodec_profile_name(codec_id AVCodecID, profile int32) string {
    return C.GoString(C.avcodec_profile_name(C.enum_AVCodecID(codec_id), 
        C.int(profile)))
}

/**
 * Return the PCM codec associated with a sample format.
 * @param be  endianness, 0 for little, 1 for big,
 *            -1 (or anything else) for native
 * @return  AV_CODEC_ID_PCM_* or AV_CODEC_ID_NONE
 */
func Av_get_pcm_codec(fmt AVSampleFormat, be int32) AVCodecID {
    return AVCodecID(C.av_get_pcm_codec(C.enum_AVSampleFormat(fmt), C.int(be)))
}

/**
 * @}
 */

                            
