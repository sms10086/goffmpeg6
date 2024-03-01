// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2001 Fabrice Bellard
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
//#include "libavutil/samplefmt.h"
//#include "libavutil/attributes.h"
//#include "libavutil/avutil.h"
//#include "libavutil/buffer.h"
//#include "libavutil/channel_layout.h"
//#include "libavutil/dict.h"
//#include "libavutil/frame.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavcodec/codec.h"
//#include "libavcodec/codec_id.h"
//#include "libavcodec/defs.h"
//#include "libavcodec/packet.h"
//#include "libavcodec/version_major.h"
//#include "libavcodec/version.h"
//#include "libavcodec/codec_desc.h"
//#include "libavcodec/codec_par.h"
//#include "libavcodec/avcodec.h"
import "C"
import (
    "syscall"
    "unsafe"
)

const AV_INPUT_BUFFER_MIN_SIZE = 16384
const AV_CODEC_FLAG_UNALIGNED = (1 <<  0)
const AV_CODEC_FLAG_QSCALE = (1 <<  1)
const AV_CODEC_FLAG_4MV = (1 <<  2)
const AV_CODEC_FLAG_OUTPUT_CORRUPT = (1 <<  3)
const AV_CODEC_FLAG_QPEL = (1 <<  4)
const AV_CODEC_FLAG_RECON_FRAME = (1 <<  6)
const AV_CODEC_FLAG_COPY_OPAQUE = (1 <<  7)
const AV_CODEC_FLAG_FRAME_DURATION = (1 <<  8)
const AV_CODEC_FLAG_PASS1 = (1 <<  9)
const AV_CODEC_FLAG_PASS2 = (1 << 10)
const AV_CODEC_FLAG_LOOP_FILTER = (1 << 11)
const AV_CODEC_FLAG_GRAY = (1 << 13)
const AV_CODEC_FLAG_PSNR = (1 << 15)
const AV_CODEC_FLAG_INTERLACED_DCT = (1 << 18)
const AV_CODEC_FLAG_LOW_DELAY = (1 << 19)
const AV_CODEC_FLAG_GLOBAL_HEADER = (1 << 22)
const AV_CODEC_FLAG_BITEXACT = (1 << 23)
const AV_CODEC_FLAG_AC_PRED = (1 << 24)
const AV_CODEC_FLAG_INTERLACED_ME = (1 << 29)
const AV_CODEC_FLAG_CLOSED_GOP = (uint32(1) << 31)
const AV_CODEC_FLAG2_FAST = (1 <<  0)
const AV_CODEC_FLAG2_NO_OUTPUT = (1 <<  2)
const AV_CODEC_FLAG2_LOCAL_HEADER = (1 <<  3)
const AV_CODEC_FLAG2_CHUNKS = (1 << 15)
const AV_CODEC_FLAG2_IGNORE_CROP = (1 << 16)
const AV_CODEC_FLAG2_SHOW_ALL = (1 << 22)
const AV_CODEC_FLAG2_EXPORT_MVS = (1 << 28)
const AV_CODEC_FLAG2_SKIP_MANUAL = (1 << 29)
const AV_CODEC_FLAG2_RO_FLUSH_NOOP = (1 << 30)
const AV_CODEC_FLAG2_ICC_PROFILES = (uint32(1) << 31)
const AV_CODEC_EXPORT_DATA_MVS = (1 << 0)
const AV_CODEC_EXPORT_DATA_PRFT = (1 << 1)
const AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS = (1 << 2)
const AV_CODEC_EXPORT_DATA_FILM_GRAIN = (1 << 3)
const AV_GET_BUFFER_FLAG_REF = (1 << 0)
const AV_GET_ENCODE_BUFFER_FLAG_REF = (1 << 0)
const FF_COMPRESSION_DEFAULT = -1
const FF_CMP_SAD = 0
const FF_CMP_SSE = 1
const FF_CMP_SATD = 2
const FF_CMP_DCT = 3
const FF_CMP_PSNR = 4
const FF_CMP_BIT = 5
const FF_CMP_RD = 6
const FF_CMP_ZERO = 7
const FF_CMP_VSAD = 8
const FF_CMP_VSSE = 9
const FF_CMP_NSSE = 10
const FF_CMP_W53 = 11
const FF_CMP_W97 = 12
const FF_CMP_DCTMAX = 13
const FF_CMP_DCT264 = 14
const FF_CMP_MEDIAN_SAD = 15
const FF_CMP_CHROMA = 256
const SLICE_FLAG_CODED_ORDER = 0x0001
const SLICE_FLAG_ALLOW_FIELD = 0x0002
const SLICE_FLAG_ALLOW_PLANE = 0x0004
const FF_MB_DECISION_SIMPLE = 0
const FF_MB_DECISION_BITS = 1
const FF_MB_DECISION_RD = 2
const FF_BUG_AUTODETECT = 1
const FF_BUG_XVID_ILACE = 4
const FF_BUG_UMP4 = 8
const FF_BUG_NO_PADDING = 16
const FF_BUG_AMV = 32
const FF_BUG_QPEL_CHROMA = 64
const FF_BUG_STD_QPEL = 128
const FF_BUG_QPEL_CHROMA2 = 256
const FF_BUG_DIRECT_BLOCKSIZE = 512
const FF_BUG_EDGE = 1024
const FF_BUG_HPEL_CHROMA = 2048
const FF_BUG_DC_CLIP = 4096
const FF_BUG_MS = 8192
const FF_BUG_TRUNCATED = 16384
const FF_BUG_IEDGE = 32768
const FF_EC_GUESS_MVS = 1
const FF_EC_DEBLOCK = 2
const FF_EC_FAVOR_INTER = 256
const FF_DEBUG_PICT_INFO = 1
const FF_DEBUG_RC = 2
const FF_DEBUG_BITSTREAM = 4
const FF_DEBUG_MB_TYPE = 8
const FF_DEBUG_QP = 16
const FF_DEBUG_DCT_COEFF = 0x00000040
const FF_DEBUG_SKIP = 0x00000080
const FF_DEBUG_STARTCODE = 0x00000100
const FF_DEBUG_ER = 0x00000400
const FF_DEBUG_MMCO = 0x00000800
const FF_DEBUG_BUGS = 0x00001000
const FF_DEBUG_BUFFERS = 0x00008000
const FF_DEBUG_THREADS = 0x00010000
const FF_DEBUG_GREEN_MD = 0x00800000
const FF_DEBUG_NOMC = 0x01000000
const FF_DCT_AUTO = 0
const FF_DCT_FASTINT = 1
const FF_DCT_INT = 2
const FF_DCT_MMX = 3
const FF_DCT_ALTIVEC = 5
const FF_DCT_FAAN = 6
const FF_IDCT_AUTO = 0
const FF_IDCT_INT = 1
const FF_IDCT_SIMPLE = 2
const FF_IDCT_SIMPLEMMX = 3
const FF_IDCT_ARM = 7
const FF_IDCT_ALTIVEC = 8
const FF_IDCT_SIMPLEARM = 10
const FF_IDCT_XVID = 14
const FF_IDCT_SIMPLEARMV5TE = 16
const FF_IDCT_SIMPLEARMV6 = 17
const FF_IDCT_FAAN = 20
const FF_IDCT_SIMPLENEON = 22
const FF_IDCT_SIMPLEAUTO = 128
const FF_THREAD_FRAME = 1
const FF_THREAD_SLICE = 2
const FF_SUB_CHARENC_MODE_DO_NOTHING = -1
const FF_SUB_CHARENC_MODE_AUTOMATIC = 0
const FF_SUB_CHARENC_MODE_PRE_DECODER = 1
const FF_SUB_CHARENC_MODE_IGNORE = 2
const FF_CODEC_PROPERTY_LOSSLESS = 0x00000001
const FF_CODEC_PROPERTY_CLOSED_CAPTIONS = 0x00000002
const FF_CODEC_PROPERTY_FILM_GRAIN = 0x00000004
const AV_HWACCEL_CODEC_CAP_EXPERIMENTAL = 0x0200
const AV_HWACCEL_FLAG_IGNORE_LEVEL = (1 << 0)
const AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH = (1 << 1)
const AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = (1 << 2)
const AV_HWACCEL_FLAG_UNSAFE_OUTPUT = (1 << 3)
const AV_SUBTITLE_FLAG_FORCED = 0x00000001
const AV_PARSER_PTS_NB = 4
const PARSER_FLAG_COMPLETE_FRAMES = 0x0001
const PARSER_FLAG_ONCE = 0x0002
const PARSER_FLAG_FETCHED_OFFSET = 0x0004
const PARSER_FLAG_USE_CODEC_TS = 0x1000


                         
                         

/**
 * @file
 * @ingroup libavc
 * Libavcodec external API header
 */

                                
                                 
                             
                             
                                     
                           
                            
                          
                             
                               

                  
                     
                 
                   
                          
                        
/* When included as part of the ffmpeg build, only include the major version
 * to avoid unnecessary rebuilds. When included externally, keep including
 * the full version information. */
                    

                       
                      
      

type AVCodecParameters C.struct_AVCodecParameters

/**
 * @defgroup libavc libavcodec
 * Encoding/Decoding Library
 *
 * @{
 *
 * @defgroup lavc_decoding Decoding
 * @{
 * @}
 *
 * @defgroup lavc_encoding Encoding
 * @{
 * @}
 *
 * @defgroup lavc_codec Codecs
 * @{
 * @defgroup lavc_codec_native Native Codecs
 * @{
 * @}
 * @defgroup lavc_codec_wrappers External library wrappers
 * @{
 * @}
 * @defgroup lavc_codec_hwaccel Hardware Accelerators bridge
 * @{
 * @}
 * @}
 * @defgroup lavc_internal Internal
 * @{
 * @}
 * @}
 */

/**
 * @ingroup libavc
 * @defgroup lavc_encdec send/receive encoding and decoding API overview
 * @{
 *
 * The avcodec_send_packet()/avcodec_receive_frame()/avcodec_send_frame()/
 * avcodec_receive_packet() functions provide an encode/decode API, which
 * decouples input and output.
 *
 * The API is very similar for encoding/decoding and audio/video, and works as
 * follows:
 * - Set up and open the AVCodecContext as usual.
 * - Send valid input:
 *   - For decoding, call avcodec_send_packet() to give the decoder raw
 *     compressed data in an AVPacket.
 *   - For encoding, call avcodec_send_frame() to give the encoder an AVFrame
 *     containing uncompressed audio or video.
 *
 *   In both cases, it is recommended that AVPackets and AVFrames are
 *   refcounted, or libavcodec might have to copy the input data. (libavformat
 *   always returns refcounted AVPackets, and av_frame_get_buffer() allocates
 *   refcounted AVFrames.)
 * - Receive output in a loop. Periodically call one of the avcodec_receive_*()
 *   functions and process their output:
 *   - For decoding, call avcodec_receive_frame(). On success, it will return
 *     an AVFrame containing uncompressed audio or video data.
 *   - For encoding, call avcodec_receive_packet(). On success, it will return
 *     an AVPacket with a compressed frame.
 *
 *   Repeat this call until it returns AVERROR(EAGAIN) or an error. The
 *   AVERROR(EAGAIN) return value means that new input data is required to
 *   return new output. In this case, continue with sending input. For each
 *   input frame/packet, the codec will typically return 1 output frame/packet,
 *   but it can also be 0 or more than 1.
 *
 * At the beginning of decoding or encoding, the codec might accept multiple
 * input frames/packets without returning a frame, until its internal buffers
 * are filled. This situation is handled transparently if you follow the steps
 * outlined above.
 *
 * In theory, sending input can result in EAGAIN - this should happen only if
 * not all output was received. You can use this to structure alternative decode
 * or encode loops other than the one suggested above. For example, you could
 * try sending new input on each iteration, and try to receive output if that
 * returns EAGAIN.
 *
 * End of stream situations. These require "flushing" (aka draining) the codec,
 * as the codec might buffer multiple frames or packets internally for
 * performance or out of necessity (consider B-frames).
 * This is handled as follows:
 * - Instead of valid input, send NULL to the avcodec_send_packet() (decoding)
 *   or avcodec_send_frame() (encoding) functions. This will enter draining
 *   mode.
 * - Call avcodec_receive_frame() (decoding) or avcodec_receive_packet()
 *   (encoding) in a loop until AVERROR_EOF is returned. The functions will
 *   not return AVERROR(EAGAIN), unless you forgot to enter draining mode.
 * - Before decoding can be resumed again, the codec has to be reset with
 *   avcodec_flush_buffers().
 *
 * Using the API as outlined above is highly recommended. But it is also
 * possible to call functions outside of this rigid schema. For example, you can
 * call avcodec_send_packet() repeatedly without calling
 * avcodec_receive_frame(). In this case, avcodec_send_packet() will succeed
 * until the codec's internal buffer has been filled up (which is typically of
 * size 1 per output frame, after initial input), and then reject input with
 * AVERROR(EAGAIN). Once it starts rejecting input, you have no choice but to
 * read at least some output.
 *
 * Not all codecs will follow a rigid and predictable dataflow; the only
 * guarantee is that an AVERROR(EAGAIN) return value on a send/receive call on
 * one end implies that a receive/send call on the other end will succeed, or
 * at least will not fail with AVERROR(EAGAIN). In general, no codec will
 * permit unlimited buffering of input or output.
 *
 * A codec is not allowed to return AVERROR(EAGAIN) for both sending and receiving. This
 * would be an invalid state, which could put the codec user into an endless
 * loop. The API has no concept of time either: it cannot happen that trying to
 * do avcodec_send_packet() results in AVERROR(EAGAIN), but a repeated call 1 second
 * later accepts the packet (with no other receive/flush API calls involved).
 * The API is a strict state machine, and the passage of time is not supposed
 * to influence it. Some timing-dependent behavior might still be deemed
 * acceptable in certain cases. But it must never result in both send/receive
 * returning EAGAIN at the same time at any point. It must also absolutely be
 * avoided that the current state is "unstable" and can "flip-flop" between
 * the send/receive APIs allowing progress. For example, it's not allowed that
 * the codec randomly decides that it actually wants to consume a packet now
 * instead of returning a frame, after it just returned AVERROR(EAGAIN) on an
 * avcodec_send_packet() call.
 * @}
 */

/**
 * @defgroup lavc_core Core functions/structures.
 * @ingroup libavc
 *
 * Basic definitions, functions for querying libavcodec capabilities,
 * allocating core structures, etc.
 * @{
 */

/**
 * @ingroup lavc_encoding
 * minimum encoding buffer size
 * Used to avoid some checks during header writing.
 */


/**
 * @ingroup lavc_encoding
 */
type RcOverride C.struct_RcOverride

/* encoding support
   These flags can be passed in AVCodecContext.flags before initialization.
   Note: Not everything is supported yet.
*/

/**
 * Allow decoders to produce frames with data planes that are not aligned
 * to CPU requirements (e.g. due to cropping).
 */

/**
 * Use fixed qscale.
 */

/**
 * 4 MV per MB allowed / advanced prediction for H.263.
 */

/**
 * Output even those frames that might be corrupted.
 */

/**
 * Use qpel MC.
 */

                      
   
                                                         
                           
  
                                                                            
   
                                               
      
/**
 * Request the encoder to output reconstructed frames, i.e.\ frames that would
 * be produced by decoding the encoded bistream. These frames may be retrieved
 * by calling avcodec_receive_frame() immediately after a successful call to
 * avcodec_receive_packet().
 *
 * Should only be used with encoders flagged with the
 * @ref AV_CODEC_CAP_ENCODER_RECON_FRAME capability.
 *
 * @note
 * Each reconstructed frame returned by the encoder corresponds to the last
 * encoded packet, i.e. the frames are returned in coded order rather than
 * presentation order.
 *
 * @note
 * Frame parameters (like pixel format or dimensions) do not have to match the
 * AVCodecContext values. Make sure to use the values from the returned frame.
 */

/**
 * @par decoding
 * Request the decoder to propagate each packet's AVPacket.opaque and
 * AVPacket.opaque_ref to its corresponding output AVFrame.
 *
 * @par encoding:
 * Request the encoder to propagate each frame's AVFrame.opaque and
 * AVFrame.opaque_ref values to its corresponding output AVPacket.
 *
 * @par
 * May only be set on encoders that have the
 * @ref AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE capability flag.
 *
 * @note
 * While in typical cases one input frame produces exactly one output packet
 * (perhaps after a delay), in general the mapping of frames to packets is
 * M-to-N, so
 * - Any number of input frames may be associated with any given output packet.
 *   This includes zero - e.g. some encoders may output packets that carry only
 *   metadata about the whole stream.
 * - A given input frame may be associated with any number of output packets.
 *   Again this includes zero - e.g. some encoders may drop frames under certain
 *   conditions.
 * .
 * This implies that when using this flag, the caller must NOT assume that
 * - a given input frame's opaques will necessarily appear on some output packet;
 * - every output packet will have some non-NULL opaque value.
 * .
 * When an output packet contains multiple frames, the opaque values will be
 * taken from the first of those.
 *
 * @note
 * The converse holds for decoders, with frames and packets switched.
 */

/**
 * Signal to the encoder that the values of AVFrame.duration are valid and
 * should be used (typically for transferring them to output packets).
 *
 * If this flag is not set, frame durations are ignored.
 */

/**
 * Use internal 2pass ratecontrol in first pass mode.
 */

/**
 * Use internal 2pass ratecontrol in second pass mode.
 */

/**
 * loop filter.
 */

/**
 * Only decode/encode grayscale.
 */

/**
 * error[?] variables will be set during encoding.
 */

/**
 * Use interlaced DCT.
 */

/**
 * Force low delay.
 */

/**
 * Place global headers in extradata instead of every keyframe.
 */

/**
 * Use only bitexact stuff (except (I)DCT).
 */

/* Fx : Flag for H.263+ extra options */
/**
 * H.263 advanced intra coding / MPEG-4 AC prediction
 */

/**
 * interlaced motion estimation
 */



/**
 * Allow non spec compliant speedup tricks.
 */

/**
 * Skip bitstream encoding.
 */

/**
 * Place global headers at every keyframe instead of in extradata.
 */


/**
 * Input bitstream might be truncated at a packet boundaries
 * instead of only at frame boundaries.
 */

/**
 * Discard cropping information from SPS.
 */


/**
 * Show all frames before the first keyframe
 */

/**
 * Export motion vectors through frame side data
 */

/**
 * Do not skip samples and export skip information as frame side data
 */

/**
 * Do not reset ASS ReadOrder field on flush (subtitles decoding)
 */

/**
 * Generate/parse ICC profiles on encode/decode, as appropriate for the type of
 * file. No effect on codecs which cannot contain embedded ICC profiles, or
 * when compiled without support for lcms2.
 */


/* Exported side data.
   These flags can be passed in AVCodecContext.export_side_data before initialization.
*/
/**
 * Export motion vectors through frame side data
 */

/**
 * Export encoder Producer Reference Time through packet side data
 */

/**
 * Decoding only.
 * Export the AVVideoEncParams structure through frame side data.
 */

/**
 * Decoding only.
 * Do not apply film grain, export it instead.
 */


/**
 * The decoder will keep a reference to the frame and may reuse it later.
 */


/**
 * The encoder will keep a reference to the packet and may reuse it later.
 */


/**
 * main external API structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * You can use AVOptions (av_opt* / av_set/get*()) to access these fields from user
 * applications.
 * The name string for AVOptions options matches the associated command line
 * parameter name and can be found in libavcodec/options_table.h
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 * sizeof(AVCodecContext) must not be used outside libav*.
 */
type AVCodecContext C.struct_AVCodecContext

/**
 * @defgroup lavc_hwaccel AVHWAccel
 *
 * @note  Nothing in this structure should be accessed by the user.  At some
 *        point in future it will not be externally visible at all.
 *
 * @{
 */
type AVHWAccel C.struct_AVHWAccel

/**
 * HWAccel is experimental and is thus avoided in favor of non experimental
 * codecs
 */


/**
 * Hardware acceleration should be used for decoding even if the codec level
 * used is unknown or higher than the maximum supported level reported by the
 * hardware driver.
 *
 * It's generally a good idea to pass this flag unless you have a specific
 * reason not to, as hardware tends to under-report supported levels.
 */


/**
 * Hardware acceleration can output YUV pixel formats with a different chroma
 * sampling than 4:2:0 and/or other than 8 bits per component.
 */


/**
 * Hardware acceleration should still be attempted for decoding when the
 * codec profile does not match the reported capabilities of the hardware.
 *
 * For example, this can be used to try to decode baseline profile H.264
 * streams in hardware - it will often succeed, because many streams marked
 * as baseline profile actually conform to constrained baseline profile.
 *
 * @warning If the stream is actually not supported then the behaviour is
 *          undefined, and may include returning entirely incorrect output
 *          while indicating success.
 */


/**
 * Some hardware decoders (namely nvdec) can either output direct decoder
 * surfaces, or make an on-device copy and return said copy.
 * There is a hard limit on how many decoder surfaces there can be, and it
 * cannot be accurately guessed ahead of time.
 * For some processing chains, this can be okay, but others will run into the
 * limit and in turn produce very confusing errors that require fine tuning of
 * more or less obscure options by the user, or in extreme cases cannot be
 * resolved at all without inserting an avfilter that forces a copy.
 *
 * Thus, the hwaccel will by default make a copy for safety and resilience.
 * If a users really wants to minimize the amount of copies, they can set this
 * flag and ensure their processing chain does not exhaust the surface pool.
 */


/**
 * @}
 */

type AVSubtitleType C.enum_AVSubtitleType



type AVSubtitleRect C.struct_AVSubtitleRect

type AVSubtitle C.struct_AVSubtitle

/**
 * Return the LIBAVCODEC_VERSION_INT constant.
 */
func Avcodec_version() uint32 {
    return uint32(C.avcodec_version())
}

/**
 * Return the libavcodec build-time configuration.
 */
func Avcodec_configuration() string {
    return C.GoString(C.avcodec_configuration())
}

/**
 * Return the libavcodec license.
 */
func Avcodec_license() string {
    return C.GoString(C.avcodec_license())
}

/**
 * Allocate an AVCodecContext and set its fields to default values. The
 * resulting struct should be freed with avcodec_free_context().
 *
 * @param codec if non-NULL, allocate private data and initialize defaults
 *              for the given codec. It is illegal to then call avcodec_open2()
 *              with a different codec.
 *              If NULL, then the codec-specific defaults won't be initialized,
 *              which may result in suboptimal default settings (this is
 *              important mainly for encoders, e.g. libx264).
 *
 * @return An AVCodecContext filled with default values or NULL on failure.
 */
func Avcodec_alloc_context3(codec *AVCodec) *AVCodecContext {
    return (*AVCodecContext)(unsafe.Pointer(C.avcodec_alloc_context3(
        (*C.AVCodec)(unsafe.Pointer(codec)))))
}

/**
 * Free the codec context and everything associated with it and write NULL to
 * the provided pointer.
 */
func Avcodec_free_context(avctx **AVCodecContext)  {
    C.avcodec_free_context((**C.AVCodecContext)(unsafe.Pointer(avctx)))
}

/**
 * Get the AVClass for AVCodecContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Avcodec_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avcodec_get_class()))
}

/**
 * Get the AVClass for AVSubtitleRect. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Avcodec_get_subtitle_rect_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avcodec_get_subtitle_rect_class()))
}

/**
 * Fill the parameters struct based on the values from the supplied codec
 * context. Any allocated fields in par are freed and replaced with duplicates
 * of the corresponding fields in codec.
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
func Avcodec_parameters_from_context(par *AVCodecParameters,
                                    codec *AVCodecContext) int32 {
    return int32(C.avcodec_parameters_from_context(
        (*C.struct_AVCodecParameters)(unsafe.Pointer(par)), 
        (*C.AVCodecContext)(unsafe.Pointer(codec))))
}

/**
 * Fill the codec context based on the values from the supplied codec
 * parameters. Any allocated fields in codec that have a corresponding field in
 * par are freed and replaced with duplicates of the corresponding field in par.
 * Fields in codec that do not have a counterpart in par are not touched.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func Avcodec_parameters_to_context(codec *AVCodecContext,
                                  par *AVCodecParameters) int32 {
    return int32(C.avcodec_parameters_to_context(
        (*C.AVCodecContext)(unsafe.Pointer(codec)), 
        (*C.struct_AVCodecParameters)(unsafe.Pointer(par))))
}

/**
 * Initialize the AVCodecContext to use the given AVCodec. Prior to using this
 * function the context has to be allocated with avcodec_alloc_context3().
 *
 * The functions avcodec_find_decoder_by_name(), avcodec_find_encoder_by_name(),
 * avcodec_find_decoder() and avcodec_find_encoder() provide an easy way for
 * retrieving a codec.
 *
 * Depending on the codec, you might need to set options in the codec context
 * also for decoding (e.g. width, height, or the pixel or audio sample format in
 * the case the information is not available in the bitstream, as when decoding
 * raw audio or video).
 *
 * Options in the codec context can be set either by setting them in the options
 * AVDictionary, or by setting the values in the context itself, directly or by
 * using the av_opt_set() API before calling this function.
 *
 * Example:
 * @code
 * av_dict_set(&opts, "b", "2.5M", 0);
 * codec = avcodec_find_decoder(AV_CODEC_ID_H264);
 * if (!codec)
 *     exit(1);
 *
 * context = avcodec_alloc_context3(codec);
 *
 * if (avcodec_open2(context, codec, opts) < 0)
 *     exit(1);
 * @endcode
 *
 * In the case AVCodecParameters are available (e.g. when demuxing a stream
 * using libavformat, and accessing the AVStream contained in the demuxer), the
 * codec parameters can be copied to the codec context using
 * avcodec_parameters_to_context(), as in the following example:
 *
 * @code
 * AVStream *stream = ...;
 * context = avcodec_alloc_context3(codec);
 * if (avcodec_parameters_to_context(context, stream->codecpar) < 0)
 *     exit(1);
 * if (avcodec_open2(context, codec, NULL) < 0)
 *     exit(1);
 * @endcode
 *
 * @note Always call this function before using decoding routines (such as
 * @ref avcodec_receive_frame()).
 *
 * @param avctx The context to initialize.
 * @param codec The codec to open this context for. If a non-NULL codec has been
 *              previously passed to avcodec_alloc_context3() or
 *              for this context, then this parameter MUST be either NULL or
 *              equal to the previously passed codec.
 * @param options A dictionary filled with AVCodecContext and codec-private
 *                options, which are set on top of the options already set in
 *                avctx, can be NULL. On return this object will be filled with
 *                options that were not found in the avctx codec context.
 *
 * @return zero on success, a negative value on error
 * @see avcodec_alloc_context3(), avcodec_find_decoder(), avcodec_find_encoder(),
 *      av_dict_set(), av_opt_set(), av_opt_find(), avcodec_parameters_to_context()
 */
func Avcodec_open2(avctx *AVCodecContext, codec *AVCodec, options **AVDictionary) int32 {
    return int32(C.avcodec_open2((*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVCodec)(unsafe.Pointer(codec)), 
        (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Close a given AVCodecContext and free all the data associated with it
 * (but not the AVCodecContext itself).
 *
 * Calling this function on an AVCodecContext that hasn't been opened will free
 * the codec-specific data allocated in avcodec_alloc_context3() with a non-NULL
 * codec. Subsequent calls will do nothing.
 *
 * @note Do not use this function. Use avcodec_free_context() to destroy a
 * codec context (either open or closed). Opening and closing a codec context
 * multiple times is not supported anymore -- use multiple codec contexts
 * instead.
 */
func Avcodec_close(avctx *AVCodecContext) int32 {
    return int32(C.avcodec_close((*C.AVCodecContext)(unsafe.Pointer(avctx))))
}

/**
 * Free all allocated data in the given subtitle struct.
 *
 * @param sub AVSubtitle to free.
 */
func Avsubtitle_free(sub *AVSubtitle)  {
    C.avsubtitle_free((*C.AVSubtitle)(unsafe.Pointer(sub)))
}

/**
 * @}
 */

/**
 * @addtogroup lavc_decoding
 * @{
 */

/**
 * The default callback for AVCodecContext.get_buffer2(). It is made public so
 * it can be called by custom get_buffer2() implementations for decoders without
 * AV_CODEC_CAP_DR1 set.
 */
func Avcodec_default_get_buffer2(s *AVCodecContext, frame *AVFrame, flags int32) int32 {
    return int32(C.avcodec_default_get_buffer2(
        (*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * The default callback for AVCodecContext.get_encode_buffer(). It is made public so
 * it can be called by custom get_encode_buffer() implementations for encoders without
 * AV_CODEC_CAP_DR1 set.
 */
func Avcodec_default_get_encode_buffer(s *AVCodecContext, pkt *AVPacket, flags int32) int32 {
    return int32(C.avcodec_default_get_encode_buffer(
        (*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.AVPacket)(unsafe.Pointer(pkt)), C.int(flags)))
}

/**
 * Modify width and height values so that they will result in a memory
 * buffer that is acceptable for the codec if you do not use any horizontal
 * padding.
 *
 * May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
func Avcodec_align_dimensions(s *AVCodecContext, width *int32, height *int32)  {
    C.avcodec_align_dimensions((*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)))
}

/**
 * Modify width and height values so that they will result in a memory
 * buffer that is acceptable for the codec if you also ensure that all
 * line sizes are a multiple of the respective linesize_align[i].
 *
 * May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
func Avcodec_align_dimensions2(s *AVCodecContext, width *int32, height *int32,
                               linesize_align[AV_NUM_DATA_POINTERS] int32)  {
    C.avcodec_align_dimensions2((*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)), 
        (*C.int)(unsafe.Pointer(&linesize_align[0])))
}

                                
   
                                                            
  
                                                                            
                                                                             
  
                                                 
                                                 
                                                            
   
                     
                                                                                

   
                                                            
  
                                                                            
                                                                             
  
                                                 
                                                 
                                                            
   
                     
                                                                     
      

/**
 * Decode a subtitle message.
 * Return a negative value on error, otherwise return the number of bytes used.
 * If no subtitle could be decompressed, got_sub_ptr is zero.
 * Otherwise, the subtitle is stored in *sub.
 * Note that AV_CODEC_CAP_DR1 is not available for subtitle codecs. This is for
 * simplicity, because the performance difference is expected to be negligible
 * and reusing a get_buffer written for video codecs would probably perform badly
 * due to a potentially very different allocation pattern.
 *
 * Some decoders (those marked with AV_CODEC_CAP_DELAY) have a delay between input
 * and output. This means that for some packets they will not immediately
 * produce decoded output and need to be flushed at the end of decoding to get
 * all the decoded data. Flushing is done by calling this function with packets
 * with avpkt->data set to NULL and avpkt->size set to 0 until it stops
 * returning subtitles. It is safe to flush even those decoders that are not
 * marked with AV_CODEC_CAP_DELAY, then no subtitles will be returned.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 * before packets may be fed to the decoder.
 *
 * @param avctx the codec context
 * @param[out] sub The preallocated AVSubtitle in which the decoded subtitle will be stored,
 *                 must be freed with avsubtitle_free if *got_sub_ptr is set.
 * @param[in,out] got_sub_ptr Zero if no subtitle could be decompressed, otherwise, it is nonzero.
 * @param[in] avpkt The input AVPacket containing the input buffer.
 */
func Avcodec_decode_subtitle2(avctx *AVCodecContext, sub *AVSubtitle,
                             got_sub_ptr *int32, avpkt *AVPacket) int32 {
    return int32(C.avcodec_decode_subtitle2(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVSubtitle)(unsafe.Pointer(sub)), 
        (*C.int)(unsafe.Pointer(got_sub_ptr)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Supply raw packet data as input to a decoder.
 *
 * Internally, this call will copy relevant AVCodecContext fields, which can
 * influence decoding per-packet, and apply them when the packet is actually
 * decoded. (For example AVCodecContext.skip_frame, which might direct the
 * decoder to drop the frame contained by the packet sent with this function.)
 *
 * @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
 *          larger than the actual read bytes because some optimized bitstream
 *          readers read 32 or 64 bits at once and could read over the end.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 *       before packets may be fed to the decoder.
 *
 * @param avctx codec context
 * @param[in] avpkt The input AVPacket. Usually, this will be a single video
 *                  frame, or several complete audio frames.
 *                  Ownership of the packet remains with the caller, and the
 *                  decoder will not write to the packet. The decoder may create
 *                  a reference to the packet data (or copy it if the packet is
 *                  not reference-counted).
 *                  Unlike with older APIs, the packet is always fully consumed,
 *                  and if it contains multiple frames (e.g. some audio codecs),
 *                  will require you to call avcodec_receive_frame() multiple
 *                  times afterwards before you can send a new packet.
 *                  It can be NULL (or an AVPacket with data set to NULL and
 *                  size set to 0); in this case, it is considered a flush
 *                  packet, which signals the end of the stream. Sending the
 *                  first flush packet will return success. Subsequent ones are
 *                  unnecessary and will return AVERROR_EOF. If the decoder
 *                  still has frames buffered, it will return them after sending
 *                  a flush packet.
 *
 * @retval 0                 success
 * @retval AVERROR(EAGAIN)   input is not accepted in the current state - user
 *                           must read output with avcodec_receive_frame() (once
 *                           all output is read, the packet should be resent,
 *                           and the call will not fail with EAGAIN).
 * @retval AVERROR_EOF       the decoder has been flushed, and no new packets can be
 *                           sent to it (also returned if more than 1 flush
 *                           packet is sent)
 * @retval AVERROR(EINVAL)   codec not opened, it is an encoder, or requires flush
 * @retval AVERROR(ENOMEM)   failed to add packet to internal queue, or similar
 * @retval "another negative error code" legitimate decoding errors
 */
func Avcodec_send_packet(avctx *AVCodecContext, avpkt *AVPacket) int32 {
    return int32(C.avcodec_send_packet(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Return decoded output data from a decoder or encoder (when the
 * @ref AV_CODEC_FLAG_RECON_FRAME flag is used).
 *
 * @param avctx codec context
 * @param frame This will be set to a reference-counted video or audio
 *              frame (depending on the decoder type) allocated by the
 *              codec. Note that the function will always call
 *              av_frame_unref(frame) before doing anything else.
 *
 * @retval 0                success, a frame was returned
 * @retval AVERROR(EAGAIN)  output is not available in this state - user must
 *                          try to send new input
 * @retval AVERROR_EOF      the codec has been fully flushed, and there will be
 *                          no more output frames
 * @retval AVERROR(EINVAL)  codec not opened, or it is an encoder without the
 *                          @ref AV_CODEC_FLAG_RECON_FRAME flag enabled
 * @retval "other negative error code" legitimate decoding errors
 */
func Avcodec_receive_frame(avctx *AVCodecContext, frame *AVFrame) int32 {
    return int32(C.avcodec_receive_frame(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Supply a raw video or audio frame to the encoder. Use avcodec_receive_packet()
 * to retrieve buffered output packets.
 *
 * @param avctx     codec context
 * @param[in] frame AVFrame containing the raw audio or video frame to be encoded.
 *                  Ownership of the frame remains with the caller, and the
 *                  encoder will not write to the frame. The encoder may create
 *                  a reference to the frame data (or copy it if the frame is
 *                  not reference-counted).
 *                  It can be NULL, in which case it is considered a flush
 *                  packet.  This signals the end of the stream. If the encoder
 *                  still has packets buffered, it will return them after this
 *                  call. Once flushing mode has been entered, additional flush
 *                  packets are ignored, and sending frames will return
 *                  AVERROR_EOF.
 *
 *                  For audio:
 *                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
 *                  can have any number of samples.
 *                  If it is not set, frame->nb_samples must be equal to
 *                  avctx->frame_size for all frames except the last.
 *                  The final frame may be smaller than avctx->frame_size.
 * @retval 0                 success
 * @retval AVERROR(EAGAIN)   input is not accepted in the current state - user must
 *                           read output with avcodec_receive_packet() (once all
 *                           output is read, the packet should be resent, and the
 *                           call will not fail with EAGAIN).
 * @retval AVERROR_EOF       the encoder has been flushed, and no new frames can
 *                           be sent to it
 * @retval AVERROR(EINVAL)   codec not opened, it is a decoder, or requires flush
 * @retval AVERROR(ENOMEM)   failed to add packet to internal queue, or similar
 * @retval "another negative error code" legitimate encoding errors
 */
func Avcodec_send_frame(avctx *AVCodecContext, frame *AVFrame) int32 {
    return int32(C.avcodec_send_frame((*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Read encoded data from the encoder.
 *
 * @param avctx codec context
 * @param avpkt This will be set to a reference-counted packet allocated by the
 *              encoder. Note that the function will always call
 *              av_packet_unref(avpkt) before doing anything else.
 * @retval 0               success
 * @retval AVERROR(EAGAIN) output is not available in the current state - user must
 *                         try to send input
 * @retval AVERROR_EOF     the encoder has been fully flushed, and there will be no
 *                         more output packets
 * @retval AVERROR(EINVAL) codec not opened, or it is a decoder
 * @retval "another negative error code" legitimate encoding errors
 */
func Avcodec_receive_packet(avctx *AVCodecContext, avpkt *AVPacket) int32 {
    return int32(C.avcodec_receive_packet(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Create and return a AVHWFramesContext with values adequate for hardware
 * decoding. This is meant to get called from the get_format callback, and is
 * a helper for preparing a AVHWFramesContext for AVCodecContext.hw_frames_ctx.
 * This API is for decoding with certain hardware acceleration modes/APIs only.
 *
 * The returned AVHWFramesContext is not initialized. The caller must do this
 * with av_hwframe_ctx_init().
 *
 * Calling this function is not a requirement, but makes it simpler to avoid
 * codec or hardware API specific details when manually allocating frames.
 *
 * Alternatively to this, an API user can set AVCodecContext.hw_device_ctx,
 * which sets up AVCodecContext.hw_frames_ctx fully automatically, and makes
 * it unnecessary to call this function or having to care about
 * AVHWFramesContext initialization at all.
 *
 * There are a number of requirements for calling this function:
 *
 * - It must be called from get_format with the same avctx parameter that was
 *   passed to get_format. Calling it outside of get_format is not allowed, and
 *   can trigger undefined behavior.
 * - The function is not always supported (see description of return values).
 *   Even if this function returns successfully, hwaccel initialization could
 *   fail later. (The degree to which implementations check whether the stream
 *   is actually supported varies. Some do this check only after the user's
 *   get_format callback returns.)
 * - The hw_pix_fmt must be one of the choices suggested by get_format. If the
 *   user decides to use a AVHWFramesContext prepared with this API function,
 *   the user must return the same hw_pix_fmt from get_format.
 * - The device_ref passed to this function must support the given hw_pix_fmt.
 * - After calling this API function, it is the user's responsibility to
 *   initialize the AVHWFramesContext (returned by the out_frames_ref parameter),
 *   and to set AVCodecContext.hw_frames_ctx to it. If done, this must be done
 *   before returning from get_format (this is implied by the normal
 *   AVCodecContext.hw_frames_ctx API rules).
 * - The AVHWFramesContext parameters may change every time time get_format is
 *   called. Also, AVCodecContext.hw_frames_ctx is reset before get_format. So
 *   you are inherently required to go through this process again on every
 *   get_format call.
 * - It is perfectly possible to call this function without actually using
 *   the resulting AVHWFramesContext. One use-case might be trying to reuse a
 *   previously initialized AVHWFramesContext, and calling this API function
 *   only to test whether the required frame parameters have changed.
 * - Fields that use dynamically allocated values of any kind must not be set
 *   by the user unless setting them is explicitly allowed by the documentation.
 *   If the user sets AVHWFramesContext.free and AVHWFramesContext.user_opaque,
 *   the new free callback must call the potentially set previous free callback.
 *   This API call may set any dynamically allocated fields, including the free
 *   callback.
 *
 * The function will set at least the following fields on AVHWFramesContext
 * (potentially more, depending on hwaccel API):
 *
 * - All fields set by av_hwframe_ctx_alloc().
 * - Set the format field to hw_pix_fmt.
 * - Set the sw_format field to the most suited and most versatile format. (An
 *   implication is that this will prefer generic formats over opaque formats
 *   with arbitrary restrictions, if possible.)
 * - Set the width/height fields to the coded frame size, rounded up to the
 *   API-specific minimum alignment.
 * - Only _if_ the hwaccel requires a pre-allocated pool: set the initial_pool_size
 *   field to the number of maximum reference surfaces possible with the codec,
 *   plus 1 surface for the user to work (meaning the user can safely reference
 *   at most 1 decoded surface at a time), plus additional buffering introduced
 *   by frame threading. If the hwaccel does not require pre-allocation, the
 *   field is left to 0, and the decoder will allocate new surfaces on demand
 *   during decoding.
 * - Possibly AVHWFramesContext.hwctx fields, depending on the underlying
 *   hardware API.
 *
 * Essentially, out_frames_ref returns the same as av_hwframe_ctx_alloc(), but
 * with basic frame parameters set.
 *
 * The function is stateless, and does not change the AVCodecContext or the
 * device_ref AVHWDeviceContext.
 *
 * @param avctx The context which is currently calling get_format, and which
 *              implicitly contains all state needed for filling the returned
 *              AVHWFramesContext properly.
 * @param device_ref A reference to the AVHWDeviceContext describing the device
 *                   which will be used by the hardware decoder.
 * @param hw_pix_fmt The hwaccel format you are going to return from get_format.
 * @param out_frames_ref On success, set to a reference to an _uninitialized_
 *                       AVHWFramesContext, created from the given device_ref.
 *                       Fields will be set to values required for decoding.
 *                       Not changed if an error is returned.
 * @return zero on success, a negative value on error. The following error codes
 *         have special semantics:
 *      AVERROR(ENOENT): the decoder does not support this functionality. Setup
 *                       is always manual, or it is a decoder which does not
 *                       support setting AVCodecContext.hw_frames_ctx at all,
 *                       or it is a software format.
 *      AVERROR(EINVAL): it is known that hardware decoding is not supported for
 *                       this configuration, or the device_ref is not supported
 *                       for the hwaccel referenced by hw_pix_fmt.
 */
func Avcodec_get_hw_frames_parameters(avctx *AVCodecContext,
                                     device_ref *AVBufferRef,
                                     hw_pix_fmt AVPixelFormat,
                                     out_frames_ref **AVBufferRef) int32 {
    return int32(C.avcodec_get_hw_frames_parameters(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVBufferRef)(unsafe.Pointer(device_ref)), 
        C.enum_AVPixelFormat(hw_pix_fmt), 
        (**C.AVBufferRef)(unsafe.Pointer(out_frames_ref))))
}



/**
 * @defgroup lavc_parsing Frame parsing
 * @{
 */

type AVPictureStructure C.enum_AVPictureStructure

type AVCodecParserContext C.struct_AVCodecParserContext

type AVCodecParser C.struct_AVCodecParser

/**
 * Iterate over all registered codec parsers.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered codec parser or NULL when the iteration is
 *         finished
 */
func Av_parser_iterate(opaque *unsafe.Pointer) *AVCodecParser {
    return (*AVCodecParser)(unsafe.Pointer(C.av_parser_iterate(opaque)))
}

func Av_parser_init(codec_id int32) *AVCodecParserContext {
    return (*AVCodecParserContext)(unsafe.Pointer(C.av_parser_init(C.int(codec_id))))
}

/**
 * Parse a packet.
 *
 * @param s             parser context.
 * @param avctx         codec context.
 * @param poutbuf       set to pointer to parsed buffer or NULL if not yet finished.
 * @param poutbuf_size  set to size of parsed buffer or zero if not yet finished.
 * @param buf           input buffer.
 * @param buf_size      buffer size in bytes without the padding. I.e. the full buffer
                        size is assumed to be buf_size + AV_INPUT_BUFFER_PADDING_SIZE.
                        To signal EOF, this should be 0 (so that the last frame
                        can be output).
 * @param pts           input presentation timestamp.
 * @param dts           input decoding timestamp.
 * @param pos           input byte position in stream.
 * @return the number of bytes of the input bitstream used.
 *
 * Example:
 * @code
 *   while(in_len){
 *       len = av_parser_parse2(myparser, AVCodecContext, &data, &size,
 *                                        in_data, in_len,
 *                                        pts, dts, pos);
 *       in_data += len;
 *       in_len  -= len;
 *
 *       if(size)
 *          decode_frame(data, size);
 *   }
 * @endcode
 */
func Av_parser_parse2(s *AVCodecParserContext,
                     avctx *AVCodecContext,
                     poutbuf **uint8, poutbuf_size *int32,
                     buf *uint8, buf_size int32,
                     pts int64, dts int64,
                     pos int64) int32 {
    return int32(C.av_parser_parse2((*C.AVCodecParserContext)(unsafe.Pointer(s)), 
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (**C.uchar)(unsafe.Pointer(poutbuf)), 
        (*C.int)(unsafe.Pointer(poutbuf_size)), (*C.uchar)(unsafe.Pointer(buf)), 
        C.int(buf_size), C.longlong(pts), C.longlong(dts), C.longlong(pos)))
}

func Av_parser_close(s *AVCodecParserContext)  {
    C.av_parser_close((*C.AVCodecParserContext)(unsafe.Pointer(s)))
}

/**
 * @}
 * @}
 */

/**
 * @addtogroup lavc_encoding
 * @{
 */

func Avcodec_encode_subtitle(avctx *AVCodecContext, buf *uint8, buf_size int32,
                            sub *AVSubtitle) int32 {
    return int32(C.avcodec_encode_subtitle(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size), 
        (*C.AVSubtitle)(unsafe.Pointer(sub))))
}


/**
 * @}
 */

/**
 * @defgroup lavc_misc Utility functions
 * @ingroup libavc
 *
 * Miscellaneous utility functions related to both encoding and decoding
 * (or neither).
 * @{
 */

/**
 * @defgroup lavc_misc_pixfmt Pixel formats
 *
 * Functions for working with pixel formats.
 * @{
 */

/**
 * Return a value representing the fourCC code associated to the
 * pixel format pix_fmt, or 0 if no associated fourCC code can be
 * found.
 */
func Avcodec_pix_fmt_to_codec_tag(pix_fmt AVPixelFormat) uint32 {
    return uint32(C.avcodec_pix_fmt_to_codec_tag(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * Find the best pixel format to convert to given a certain source pixel
 * format.  When converting from one pixel format to another, information loss
 * may occur.  For example, when converting from RGB24 to GRAY, the color
 * information will be lost. Similarly, other losses occur when converting from
 * some formats to other formats. avcodec_find_best_pix_fmt_of_2() searches which of
 * the given pixel formats should be used to suffer the least amount of loss.
 * The pixel formats from which it chooses one, are determined by the
 * pix_fmt_list parameter.
 *
 *
 * @param[in] pix_fmt_list AV_PIX_FMT_NONE terminated array of pixel formats to choose from
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @param[out] loss_ptr Combination of flags informing you what kind of losses will occur.
 * @return The best pixel format to convert to or -1 if none was found.
 */
func Avcodec_find_best_pix_fmt_of_list(pix_fmt_list *AVPixelFormat,
                                            src_pix_fmt AVPixelFormat,
                                            has_alpha int32, loss_ptr *int32) AVPixelFormat {
    return AVPixelFormat(C.avcodec_find_best_pix_fmt_of_list(
        (*C.enum_AVPixelFormat)(unsafe.Pointer(pix_fmt_list)), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha), 
        (*C.int)(unsafe.Pointer(loss_ptr))))
}

func Avcodec_default_get_format(s *AVCodecContext, fmt *AVPixelFormat) AVPixelFormat {
    return AVPixelFormat(C.avcodec_default_get_format(
        (*C.struct_AVCodecContext)(unsafe.Pointer(s)), 
        (*C.enum_AVPixelFormat)(unsafe.Pointer(fmt))))
}

/**
 * @}
 */

func Avcodec_string(buf *byte, buf_size int32, enc *AVCodecContext, encode int32)  {
    C.avcodec_string((*C.char)(unsafe.Pointer(buf)), C.int(buf_size), 
        (*C.AVCodecContext)(unsafe.Pointer(enc)), C.int(encode))
}

func Avcodec_default_execute(c *AVCodecContext, funcx func(c2 *AVCodecContext, arg2 unsafe.Pointer) int32,arg unsafe.Pointer, ret *int32, count int32, size int32) int32 {
    cb1 := syscall.NewCallbackCDecl(funcx)
    return int32(C.avcodec_default_execute(
        (*C.AVCodecContext)(unsafe.Pointer(c)), (*[0]byte)(unsafe.Pointer(cb1)), 
        arg, (*C.int)(unsafe.Pointer(ret)), C.int(count), C.int(size)))
}
func Avcodec_default_execute2(c *AVCodecContext, funcx func(c2 *AVCodecContext, arg2 unsafe.Pointer, _var2 int32, _var3 int32) int32,arg unsafe.Pointer, ret *int32, count int32) int32 {
    cb1 := syscall.NewCallbackCDecl(funcx)
    return int32(C.avcodec_default_execute2(
        (*C.AVCodecContext)(unsafe.Pointer(c)), (*[0]byte)(unsafe.Pointer(cb1)), 
        arg, (*C.int)(unsafe.Pointer(ret)), C.int(count)))
}
//FIXME func typedef

/**
 * Fill AVFrame audio data and linesize pointers.
 *
 * The buffer buf must be a preallocated buffer with a size big enough
 * to contain the specified samples amount. The filled AVFrame data
 * pointers will point to this buffer.
 *
 * AVFrame extended_data channel pointers are allocated if necessary for
 * planar audio.
 *
 * @param frame       the AVFrame
 *                    frame->nb_samples must be set prior to calling the
 *                    function. This function fills in frame->data,
 *                    frame->extended_data, frame->linesize[0].
 * @param nb_channels channel count
 * @param sample_fmt  sample format
 * @param buf         buffer to use for frame data
 * @param buf_size    size of buffer
 * @param align       plane size sample alignment (0 = default)
 * @return            >=0 on success, negative error code on failure
 * @todo return the size in bytes required to store the samples in
 * case of success, at the next libavutil bump
 */
func Avcodec_fill_audio_frame(frame *AVFrame, nb_channels int32,
                             sample_fmt AVSampleFormat, buf *uint8,
                             buf_size int32, align int32) int32 {
    return int32(C.avcodec_fill_audio_frame((*C.AVFrame)(unsafe.Pointer(frame)), 
        C.int(nb_channels), C.enum_AVSampleFormat(sample_fmt), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size), C.int(align)))
}

/**
 * Reset the internal codec state / flush internal buffers. Should be called
 * e.g. when seeking or when switching to a different stream.
 *
 * @note for decoders, this function just releases any references the decoder
 * might keep internally, but the caller's references remain valid.
 *
 * @note for encoders, this function will only do something if the encoder
 * declares support for AV_CODEC_CAP_ENCODER_FLUSH. When called, the encoder
 * will drain any remaining packets, and can then be re-used for a different
 * stream (as opposed to sending a null frame which will leave the encoder
 * in a permanent EOF state after draining). This can be desirable if the
 * cost of tearing down and replacing the encoder instance is high.
 */
func Avcodec_flush_buffers(avctx *AVCodecContext)  {
    C.avcodec_flush_buffers((*C.AVCodecContext)(unsafe.Pointer(avctx)))
}

/**
 * Return audio frame duration.
 *
 * @param avctx        codec context
 * @param frame_bytes  size of the frame, or 0 if unknown
 * @return             frame duration, in samples, if known. 0 if not able to
 *                     determine.
 */
func Av_get_audio_frame_duration(avctx *AVCodecContext, frame_bytes int32) int32 {
    return int32(C.av_get_audio_frame_duration(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), C.int(frame_bytes)))
}

/* memory */

/**
 * Same behaviour av_fast_malloc but the buffer has additional
 * AV_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
 *
 * In addition the whole buffer will initially and after resizes
 * be 0-initialized so that no uninitialized data will ever appear.
 */
func Av_fast_padded_malloc(ptr unsafe.Pointer, size *uint32, min_size uint64)  {
    C.av_fast_padded_malloc(ptr, (*C.uint)(unsafe.Pointer(size)), 
        C.ulonglong(min_size))
}

/**
 * Same behaviour av_fast_padded_malloc except that buffer will always
 * be 0-initialized after call.
 */
func Av_fast_padded_mallocz(ptr unsafe.Pointer, size *uint32, min_size uint64)  {
    C.av_fast_padded_mallocz(ptr, (*C.uint)(unsafe.Pointer(size)), 
        C.ulonglong(min_size))
}

/**
 * @return a positive value if s is open (i.e. avcodec_open2() was called on it
 * with no corresponding avcodec_close()), 0 otherwise.
 */
func Avcodec_is_open(s *AVCodecContext) int32 {
    return int32(C.avcodec_is_open((*C.AVCodecContext)(unsafe.Pointer(s))))
}

/**
 * @}
 */

                              
