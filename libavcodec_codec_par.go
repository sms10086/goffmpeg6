// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Codec parameters public API
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
//#include <stdint.h>
//#include "libavutil/avutil.h"
//#include "libavutil/channel_layout.h"
//#include "libavutil/rational.h"
//#include "libavutil/pixfmt.h"
//#include "libavcodec/codec_id.h"
//#include "libavcodec/defs.h"
//#include "libavcodec/packet.h"
//#include "libavcodec/codec_par.h"
import "C"
import (
    "unsafe"
)



                           
                           

                   

                             
                                     
                               
                             

                     
                 
                   

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * This struct describes the properties of an encoded stream.
 *
 * sizeof(AVCodecParameters) is not a part of the public ABI, this struct must
 * be allocated with avcodec_parameters_alloc() and freed with
 * avcodec_parameters_free().
 */
type AVCodecParameters struct {
    Codec_type AVMediaType
    Codec_id AVCodecID
    Codec_tag uint32
    Extradata *uint8
    Extradata_size int32
    Format int32
    Bit_rate int64
    Bits_per_coded_sample int32
    Bits_per_raw_sample int32
    Profile int32
    Level int32
    Width int32
    Height int32
    Sample_aspect_ratio AVRational
    Field_order AVFieldOrder
    Color_range AVColorRange
    Color_primaries AVColorPrimaries
    Color_trc AVColorTransferCharacteristic
    Color_space AVColorSpace
    Chroma_location AVChromaLocation
    Video_delay int32
    Channel_layout uint64
    Channels int32
    Sample_rate int32
    Block_align int32
    Frame_size int32
    Initial_padding int32
    Trailing_padding int32
    Seek_preroll int32
    Ch_layout AVChannelLayout
    Framerate AVRational
    Coded_side_data *AVPacketSideData
    Nb_coded_side_data int32
}


/**
 * Allocate a new AVCodecParameters and set its fields to default values
 * (unknown/invalid/0). The returned struct must be freed with
 * avcodec_parameters_free().
 */
func Avcodec_parameters_alloc() *AVCodecParameters {
    return (*AVCodecParameters)(unsafe.Pointer(C.avcodec_parameters_alloc()))
}

/**
 * Free an AVCodecParameters instance and everything associated with it and
 * write NULL to the supplied pointer.
 */
func Avcodec_parameters_free(par **AVCodecParameters)  {
    C.avcodec_parameters_free((**C.struct_AVCodecParameters)(unsafe.Pointer(par)))
}

/**
 * Copy the contents of src to dst. Any allocated fields in dst are freed and
 * replaced with newly allocated duplicates of the corresponding fields in src.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func Avcodec_parameters_copy(dst *AVCodecParameters, src *AVCodecParameters) int32 {
    return int32(C.avcodec_parameters_copy(
        (*C.struct_AVCodecParameters)(unsafe.Pointer(dst)), 
        (*C.struct_AVCodecParameters)(unsafe.Pointer(src))))
}

/**
 * This function is the same as av_get_audio_frame_duration(), except it works
 * with AVCodecParameters instead of an AVCodecContext.
 */
func Av_get_audio_frame_duration2(par *AVCodecParameters, frame_bytes int32) int32 {
    return int32(C.av_get_audio_frame_duration2(
        (*C.struct_AVCodecParameters)(unsafe.Pointer(par)), C.int(frame_bytes)))
}

/**
 * @}
 */

                             
