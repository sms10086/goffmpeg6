// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Codec descriptors public API
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
//#include "libavcodec/codec_id.h"
//#include "libavcodec/codec_desc.h"
import "C"
import (
    "unsafe"
)

const AV_CODEC_PROP_INTRA_ONLY =     (1 << 0) 
const AV_CODEC_PROP_LOSSY =          (1 << 1) 
const AV_CODEC_PROP_LOSSLESS =       (1 << 2) 
const AV_CODEC_PROP_REORDER =        (1 << 3) 
const AV_CODEC_PROP_FIELDS =         (1 << 4) 
const AV_CODEC_PROP_BITMAP_SUB =     (1 << 16) 
const AV_CODEC_PROP_TEXT_SUB =       (1 << 17) 


                            
                            

                             

                     

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * This struct describes the properties of a single codec described by an
 * AVCodecID.
 * @see avcodec_descriptor_get()
 */
type AVCodecDescriptor struct {
    Id AVCodecID
    Type AVMediaType
    Name *byte
    Long_name *byte
    Props int32
    Mime_types **byte
    Profiles *AVProfile
}


/**
 * Codec uses only intra compression.
 * Video and audio codecs only.
 */

/**
 * Codec supports lossy compression. Audio and video codecs only.
 * @note a codec may support both lossy and lossless
 * compression modes
 */

/**
 * Codec supports lossless compression. Audio and video codecs only.
 */

/**
 * Codec supports frame reordering. That is, the coded order (the order in which
 * the encoded packets are output by the encoders / stored / input to the
 * decoders) may be different from the presentation order of the corresponding
 * frames.
 *
 * For codecs that do not have this property set, PTS and DTS should always be
 * equal.
 */


/**
 * Video codec supports separate coding of fields in interlaced frames.
 */


/**
 * Subtitle codec is bitmap based
 * Decoded AVSubtitle data can be read from the AVSubtitleRect->pict field.
 */

/**
 * Subtitle codec is text based.
 * Decoded AVSubtitle data can be read from the AVSubtitleRect->ass field.
 */


/**
 * @return descriptor for given codec ID or NULL if no descriptor exists.
 */
func Avcodec_descriptor_get(id AVCodecID) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_get(
        C.enum_AVCodecID(id))))
}

/**
 * Iterate over all codec descriptors known to libavcodec.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
func Avcodec_descriptor_next(prev *AVCodecDescriptor) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_next(
        (*C.struct_AVCodecDescriptor)(unsafe.Pointer(prev)))))
}

/**
 * @return codec descriptor with the given name or NULL if no such descriptor
 *         exists.
 */
func Avcodec_descriptor_get_by_name(name *byte) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_get_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * @}
 */

                              
