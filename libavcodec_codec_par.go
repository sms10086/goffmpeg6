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
    C.avcodec_parameters_free((**C.AVCodecParameters)(unsafe.Pointer(par)))
}

/**
 * Copy the contents of src to dst. Any allocated fields in dst are freed and
 * replaced with newly allocated duplicates of the corresponding fields in src.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func Avcodec_parameters_copy(dst *AVCodecParameters, src *AVCodecParameters) int32 {
    return int32(C.avcodec_parameters_copy(
        (*C.AVCodecParameters)(unsafe.Pointer(dst)), 
        (*C.AVCodecParameters)(unsafe.Pointer(src))))
}

/**
 * This function is the same as av_get_audio_frame_duration(), except it works
 * with AVCodecParameters instead of an AVCodecContext.
 */
func Av_get_audio_frame_duration2(par *AVCodecParameters, frame_bytes int32) int32 {
    return int32(C.av_get_audio_frame_duration2(
        (*C.AVCodecParameters)(unsafe.Pointer(par)), C.int(frame_bytes)))
}

/**
 * @}
 */

                             
