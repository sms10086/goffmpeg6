// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * AVCodec public API
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
//#include "libavutil/hwcontext.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavutil/samplefmt.h"
//#include "libavcodec/codec_id.h"
//#include "libavcodec/version_major.h"
//#include "libavcodec/codec.h"
import "C"
import (
    "unsafe"
)

const AV_CODEC_CAP_DRAW_HORIZ_BAND =      (1 <<  0) 
const AV_CODEC_CAP_DR1 =                  (1 <<  1) 
const AV_CODEC_CAP_DELAY =                (1 <<  5) 
const AV_CODEC_CAP_SMALL_LAST_FRAME =     (1 <<  6) 
const AV_CODEC_CAP_SUBFRAMES =            (1 <<  8) 
const AV_CODEC_CAP_EXPERIMENTAL =         (1 <<  9) 
const AV_CODEC_CAP_CHANNEL_CONF =         (1 << 10) 
const AV_CODEC_CAP_FRAME_THREADS =        (1 << 12) 
const AV_CODEC_CAP_SLICE_THREADS =        (1 << 13) 
const AV_CODEC_CAP_PARAM_CHANGE =         (1 << 14) 
const AV_CODEC_CAP_OTHER_THREADS =        (1 << 15) 
const AV_CODEC_CAP_VARIABLE_FRAME_SIZE =  (1 << 16) 
const AV_CODEC_CAP_AVOID_PROBING =        (1 << 17) 
const AV_CODEC_CAP_HARDWARE =             (1 << 18) 
const AV_CODEC_CAP_HYBRID =               (1 << 19) 
const AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE =  (1 << 20) 
const AV_CODEC_CAP_ENCODER_FLUSH =    (1 << 21) 
const AV_CODEC_CAP_ENCODER_RECON_FRAME =  (1 << 22) 


                       
                       

                   

                             
                                
                          
                             
                               
                                

                                
                                     

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * Decoder can use draw_horiz_band callback.
 */

/**
 * Codec uses get_buffer() or get_encode_buffer() for allocating buffers and
 * supports custom allocators.
 * If not set, it might not use get_buffer() or get_encode_buffer() at all, or
 * use operations that assume the buffer was allocated by
 * avcodec_default_get_buffer2 or avcodec_default_get_encode_buffer.
 */

/**
 * Encoder or decoder requires flushing with NULL input at the end in order to
 * give the complete and correct output.
 *
 * NOTE: If this flag is not set, the codec is guaranteed to never be fed with
 *       with NULL data. The user can still send NULL data to the public encode
 *       or decode function, but libavcodec will not pass it along to the codec
 *       unless this flag is set.
 *
 * Decoders:
 * The decoder has a non-zero delay and needs to be fed with avpkt->data=NULL,
 * avpkt->size=0 at the end to get the delayed data until the decoder no longer
 * returns frames.
 *
 * Encoders:
 * The encoder needs to be fed with NULL data at the end of encoding until the
 * encoder no longer returns data.
 *
 * NOTE: For encoders implementing the AVCodec.encode2() function, setting this
 *       flag also means that the encoder must set the pts and duration for
 *       each output packet. If this flag is not set, the pts and duration will
 *       be determined by libavcodec from the input frame.
 */

/**
 * Codec can be fed a final frame with a smaller size.
 * This can be used to prevent truncation of the last audio samples.
 */


                    
/**
 * Codec can output multiple frames per AVPacket
 * Normally demuxers return one frame at a time, demuxers which do not do
 * are connected to a parser to split what they return into proper frames.
 * This flag is reserved to the very rare category of codecs which have a
 * bitstream that cannot be split into frames without timeconsuming
 * operations like full decoding. Demuxers carrying such bitstreams thus
 * may return multiple frames in a packet. This has many disadvantages like
 * prohibiting stream copy in many cases thus it should only be considered
 * as a last resort.
 */

      

/**
 * Codec is experimental and is thus avoided in favor of non experimental
 * encoders
 */

/**
 * Codec should fill in channel configuration and samplerate instead of container
 */

/**
 * Codec supports frame-level multithreading.
 */

/**
 * Codec supports slice-based (or partition-based) multithreading.
 */

/**
 * Codec supports changed parameters at any point.
 */

/**
 * Codec supports multithreading through a method other than slice- or
 * frame-level multithreading. Typically this marks wrappers around
 * multithreading-capable external libraries.
 */

/**
 * Audio encoder supports receiving a different number of samples in each call.
 */

/**
 * Decoder is not a preferred choice for probing.
 * This indicates that the decoder is not a good choice for probing.
 * It could for example be an expensive to spin up hardware decoder,
 * or it could simply not provide a lot of useful information about
 * the stream.
 * A decoder marked with this flag should only be used as last resort
 * choice for probing.
 */


/**
 * Codec is backed by a hardware implementation. Typically used to
 * identify a non-hwaccel hardware decoder. For information about hwaccels, use
 * avcodec_get_hw_config() instead.
 */


/**
 * Codec is potentially backed by a hardware implementation, but not
 * necessarily. This is used instead of AV_CODEC_CAP_HARDWARE, if the
 * implementation provides some sort of internal fallback.
 */


/**
 * This encoder can reorder user opaque values from input AVFrames and return
 * them with corresponding output packets.
 * @see AV_CODEC_FLAG_COPY_OPAQUE
 */


/**
 * This encoder can be flushed using avcodec_flush_buffers(). If this flag is
 * not set, the encoder must be closed and reopened to ensure that no frames
 * remain pending.
 */


/**
 * The encoder is able to output reconstructed frame data, i.e. raw frames that
 * would be produced by decoding the encoded bitstream.
 *
 * Reconstructed frame output is enabled by the AV_CODEC_FLAG_RECON_FRAME flag.
 */


/**
 * AVProfile.
 */
type AVProfile struct {
    Profile int32
    Name *byte
}


/**
 * AVCodec.
 */
type AVCodec struct {
    Name *byte
    Long_name *byte
    Type AVMediaType
    Id AVCodecID
    Capabilities int32
    Max_lowres uint8
    Supported_framerates *AVRational
    Pix_fmts *AVPixelFormat
    Supported_samplerates *int32
    Sample_fmts *AVSampleFormat
    Channel_layouts *uint64
    Priv_class *AVClass
    Profiles *AVProfile
    Wrapper_name *byte
    Ch_layouts *AVChannelLayout
}


/**
 * Iterate over all registered codecs.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered codec or NULL when the iteration is
 *         finished
 */
func Av_codec_iterate(opaque *unsafe.Pointer) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.av_codec_iterate(opaque)))
}

/**
 * Find a registered decoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
func Avcodec_find_decoder(id AVCodecID) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_decoder(C.enum_AVCodecID(id))))
}

/**
 * Find a registered decoder with the specified name.
 *
 * @param name name of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
func Avcodec_find_decoder_by_name(name *byte) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_decoder_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Find a registered encoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
func Avcodec_find_encoder(id AVCodecID) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_encoder(C.enum_AVCodecID(id))))
}

/**
 * Find a registered encoder with the specified name.
 *
 * @param name name of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
func Avcodec_find_encoder_by_name(name *byte) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_encoder_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}
/**
 * @return a non-zero number if codec is an encoder, zero otherwise
 */
func Av_codec_is_encoder(codec *AVCodec) int32 {
    return int32(C.av_codec_is_encoder((*C.struct_AVCodec)(unsafe.Pointer(codec))))
}

/**
 * @return a non-zero number if codec is a decoder, zero otherwise
 */
func Av_codec_is_decoder(codec *AVCodec) int32 {
    return int32(C.av_codec_is_decoder((*C.struct_AVCodec)(unsafe.Pointer(codec))))
}

/**
 * Return a name for the specified profile, if available.
 *
 * @param codec the codec that is searched for the given profile
 * @param profile the profile value for which a name is requested
 * @return A name for the profile if found, NULL otherwise.
 */
func Av_get_profile_name(codec *AVCodec, profile int32) string {
    return C.GoString(C.av_get_profile_name(
        (*C.struct_AVCodec)(unsafe.Pointer(codec)), C.int(profile)))
}

const (
    AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX  = 0x01 + iota
    AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX = 0x02
    AV_CODEC_HW_CONFIG_METHOD_INTERNAL = 0x04
    AV_CODEC_HW_CONFIG_METHOD_AD_HOC = 0x08
)


type AVCodecHWConfig struct {
    Pix_fmt AVPixelFormat
    Methods int32
    Device_type AVHWDeviceType
}


/**
 * Retrieve supported hardware configurations for a codec.
 *
 * Values of index from zero to some maximum return the indexed configuration
 * descriptor; all other values return NULL.  If the codec does not support
 * any hardware configurations then it will always return NULL.
 */
func Avcodec_get_hw_config(codec *AVCodec, index int32) *AVCodecHWConfig {
    return (*AVCodecHWConfig)(unsafe.Pointer(C.avcodec_get_hw_config(
        (*C.struct_AVCodec)(unsafe.Pointer(codec)), C.int(index))))
}

/**
 * @}
 */

                            
