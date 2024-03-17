// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * AVPacket public API
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
//#include "libavutil/attributes.h"
//#include "libavutil/buffer.h"
//#include "libavutil/dict.h"
//#include "libavutil/rational.h"
//#include "libavutil/version.h"
//#include "libavcodec/version_major.h"
//#include "libavcodec/packet.h"
import "C"
import (
    "unsafe"
)

//const AV_PKT_DATA_QUALITY_FACTOR =  AV_PKT_DATA_QUALITY_STATS  
const AV_PKT_FLAG_KEY =      0x0001  
const AV_PKT_FLAG_CORRUPT =  0x0002  
const AV_PKT_FLAG_DISCARD =    0x0004 
const AV_PKT_FLAG_TRUSTED =    0x0008 
const AV_PKT_FLAG_DISPOSABLE =  0x0010 


                        
                        

                   
                   

                                 
                             
                           
                               
                              

                                     

/**
 * @defgroup lavc_packet_side_data AVPacketSideData
 *
 * Types and functions for working with AVPacketSideData.
 * @{
 */
type AVPacketSideDataType int32
const (
    AV_PKT_DATA_PALETTE AVPacketSideDataType = iota
    AV_PKT_DATA_NEW_EXTRADATA
    AV_PKT_DATA_PARAM_CHANGE
    AV_PKT_DATA_H263_MB_INFO
    AV_PKT_DATA_REPLAYGAIN
    AV_PKT_DATA_DISPLAYMATRIX
    AV_PKT_DATA_STEREO3D
    AV_PKT_DATA_AUDIO_SERVICE_TYPE
    AV_PKT_DATA_QUALITY_STATS
    AV_PKT_DATA_FALLBACK_TRACK
    AV_PKT_DATA_CPB_PROPERTIES
    AV_PKT_DATA_SKIP_SAMPLES
    AV_PKT_DATA_JP_DUALMONO
    AV_PKT_DATA_STRINGS_METADATA
    AV_PKT_DATA_SUBTITLE_POSITION
    AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
    AV_PKT_DATA_WEBVTT_IDENTIFIER
    AV_PKT_DATA_WEBVTT_SETTINGS
    AV_PKT_DATA_METADATA_UPDATE
    AV_PKT_DATA_MPEGTS_STREAM_ID
    AV_PKT_DATA_MASTERING_DISPLAY_METADATA
    AV_PKT_DATA_SPHERICAL
    AV_PKT_DATA_CONTENT_LIGHT_LEVEL
    AV_PKT_DATA_A53_CC
    AV_PKT_DATA_ENCRYPTION_INIT_INFO
    AV_PKT_DATA_ENCRYPTION_INFO
    AV_PKT_DATA_AFD
    AV_PKT_DATA_PRFT
    AV_PKT_DATA_ICC_PROFILE
    AV_PKT_DATA_DOVI_CONF
    AV_PKT_DATA_S12M_TIMECODE
    AV_PKT_DATA_DYNAMIC_HDR10_PLUS
    AV_PKT_DATA_NB
)


//DEPRECATED

/**
 * This structure stores auxiliary information for decoding, presenting, or
 * otherwise processing the coded stream. It is typically exported by demuxers
 * and encoders and can be fed to decoders and muxers either in a per packet
 * basis, or as global side data (applying to the entire coded stream).
 *
 * Global side data is handled as follows:
 * - During demuxing, it may be exported through
 *   @ref AVStream.codecpar.side_data "AVStream's codec parameters", which can
 *   then be passed as input to decoders through the
 *   @ref AVCodecContext.coded_side_data "decoder context's side data", for
 *   initialization.
 * - For muxing, it can be fed through @ref AVStream.codecpar.side_data
 *   "AVStream's codec parameters", typically  the output of encoders through
 *   the @ref AVCodecContext.coded_side_data "encoder context's side data", for
 *   initialization.
 *
 * Packet specific side data is handled as follows:
 * - During demuxing, it may be exported through @ref AVPacket.side_data
 *   "AVPacket's side data", which can then be passed as input to decoders.
 * - For muxing, it can be fed through @ref AVPacket.side_data "AVPacket's
 *   side data", typically the output of encoders.
 *
 * Different modules may accept or export different types of side data
 * depending on media type and codec. Refer to @ref AVPacketSideDataType for a
 * list of defined types and where they may be found or used.
 */
type AVPacketSideData struct {
    Data *uint8
    Size uint64
    Type AVPacketSideDataType
}


/**
 * Allocate a new packet side data.
 *
 * @param sd    pointer to an array of side data to which the side data should
 *              be added. *sd may be NULL, in which case the array will be
 *              initialized.
 * @param nb_sd pointer to an integer containing the number of entries in
 *              the array. The integer value will be increased by 1 on success.
 * @param type  side data type
 * @param size  desired side data size
 * @param flags currently unused. Must be zero
 *
 * @return pointer to freshly allocated side data on success, or NULL otherwise.
 */
func Av_packet_side_data_new(psd **AVPacketSideData, pnb_sd *int32,
                                          typex AVPacketSideDataType,
                                          size uint64, flags int32) *AVPacketSideData {
    return (*AVPacketSideData)(unsafe.Pointer(C.av_packet_side_data_new(
        (**C.struct_AVPacketSideData)(unsafe.Pointer(psd)), 
        (*C.int)(unsafe.Pointer(pnb_sd)), C.enum_AVPacketSideDataType(typex), 
        C.ulonglong(size), C.int(flags))))
}

/**
 * Wrap existing data as packet side data.
 *
 * @param sd    pointer to an array of side data to which the side data should
 *              be added. *sd may be NULL, in which case the array will be
 *              initialized
 * @param nb_sd pointer to an integer containing the number of entries in
 *              the array. The integer value will be increased by 1 on success.
 * @param type  side data type
 * @param data  a data array. It must be allocated with the av_malloc() family
 *              of functions. The ownership of the data is transferred to the
 *              side data array on success
 * @param size  size of the data array
 * @param flags currently unused. Must be zero
 *
 * @return pointer to freshly allocated side data on success, or NULL otherwise
 *         On failure, the side data array is unchanged and the data remains
 *         owned by the caller.
 */
func Av_packet_side_data_add(sd **AVPacketSideData, nb_sd *int32,
                                          typex AVPacketSideDataType,
                                          data unsafe.Pointer, size uint64, flags int32) *AVPacketSideData {
    return (*AVPacketSideData)(unsafe.Pointer(C.av_packet_side_data_add(
        (**C.struct_AVPacketSideData)(unsafe.Pointer(sd)), 
        (*C.int)(unsafe.Pointer(nb_sd)), C.enum_AVPacketSideDataType(typex), 
        data, C.ulonglong(size), C.int(flags))))
}

/**
 * Get side information from a side data array.
 *
 * @param sd    the array from which the side data should be fetched
 * @param nb_sd value containing the number of entries in the array.
 * @param type  desired side information type
 *
 * @return pointer to side data if present or NULL otherwise
 */
func Av_packet_side_data_get(sd *AVPacketSideData,
                                                nb_sd int32,
                                                typex AVPacketSideDataType) *AVPacketSideData {
    return (*AVPacketSideData)(unsafe.Pointer(C.av_packet_side_data_get(
        (*C.struct_AVPacketSideData)(unsafe.Pointer(sd)), C.int(nb_sd), 
        C.enum_AVPacketSideDataType(typex))))
}

/**
 * Remove side data of the given type from a side data array.
 *
 * @param sd    the array from which the side data should be removed
 * @param nb_sd pointer to an integer containing the number of entries in
 *              the array. Will be reduced by the amount of entries removed
 *              upon return
 * @param type  side information type
 */
func Av_packet_side_data_remove(sd *AVPacketSideData, nb_sd *int32,
                                typex AVPacketSideDataType)  {
    C.av_packet_side_data_remove((*C.struct_AVPacketSideData)(unsafe.Pointer(sd)), 
        (*C.int)(unsafe.Pointer(nb_sd)), C.enum_AVPacketSideDataType(typex))
}

/**
 * Convenience function to free all the side data stored in an array, and
 * the array itself.
 *
 * @param sd    pointer to array of side data to free. Will be set to NULL
 *              upon return.
 * @param nb_sd pointer to an integer containing the number of entries in
 *              the array. Will be set to 0 upon return.
 */
func Av_packet_side_data_free(sd **AVPacketSideData, nb_sd *int32)  {
    C.av_packet_side_data_free((**C.struct_AVPacketSideData)(unsafe.Pointer(sd)), 
        (*C.int)(unsafe.Pointer(nb_sd)))
}

func Av_packet_side_data_name(typex AVPacketSideDataType) string {
    return C.GoString(C.av_packet_side_data_name(C.enum_AVPacketSideDataType(typex)))
}

/**
 * @}
 */

/**
 * @defgroup lavc_packet AVPacket
 *
 * Types and functions for working with AVPacket.
 * @{
 */

/**
 * This structure stores compressed data. It is typically exported by demuxers
 * and then passed as input to decoders, or received as output from encoders and
 * then passed to muxers.
 *
 * For video, it should typically contain one compressed frame. For audio it may
 * contain several compressed frames. Encoders are allowed to output empty
 * packets, with no compressed data, containing only side data
 * (e.g. to update some stream parameters at the end of encoding).
 *
 * The semantics of data ownership depends on the buf field.
 * If it is set, the packet data is dynamically allocated and is
 * valid indefinitely until a call to av_packet_unref() reduces the
 * reference count to 0.
 *
 * If the buf field is not set av_packet_ref() would make a copy instead
 * of increasing the reference count.
 *
 * The side data is always allocated with av_malloc(), copied by
 * av_packet_ref() and freed by av_packet_unref().
 *
 * sizeof(AVPacket) being a part of the public ABI is deprecated. once
 * av_init_packet() is removed, new packets will only be able to be allocated
 * with av_packet_alloc(), and new fields may be added to the end of the struct
 * with a minor bump.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 * @see av_packet_unref
 */
type AVPacket struct {
    Buf *AVBufferRef
    Pts int64
    Dts int64
    Data *uint8
    Size int32
    Stream_index int32
    Flags int32
    Side_data *AVPacketSideData
    Side_data_elems int32
    Duration int64
    Pos int64
    Opaque unsafe.Pointer
    Opaque_ref *AVBufferRef
    Time_base AVRational
}


                      

type AVPacketList struct {
    Pkt AVPacket
    Next *AVPacketList
}

      

///< The packet contains a keyframe
///< The packet content is corrupted
/**
 * Flag is used to discard packets which are required to maintain valid
 * decoder state but are not required for output and should be dropped
 * after decoding.
 **/

/**
 * The packet comes from a trusted source.
 *
 * Otherwise-unsafe constructs such as arbitrary pointers to data
 * outside the packet may be followed.
 */

/**
 * Flag is used to indicate packets that contain frames that can
 * be discarded by the decoder.  I.e. Non-reference frames.
 */


type AVSideDataParamChangeFlags int32
const (
    AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT AVSideDataParamChangeFlags = 0x0001 + iota
    AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT = 0x0002
    AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE = 0x0004
    AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS = 0x0008
)


/**
 * Allocate an AVPacket and set its fields to default values.  The resulting
 * struct must be freed using av_packet_free().
 *
 * @return An AVPacket filled with default values or NULL on failure.
 *
 * @note this only allocates the AVPacket itself, not the data buffers. Those
 * must be allocated through other means such as av_new_packet.
 *
 * @see av_new_packet
 */
func Av_packet_alloc() *AVPacket {
    return (*AVPacket)(unsafe.Pointer(C.av_packet_alloc()))
}

/**
 * Create a new packet that references the same data as src.
 *
 * This is a shortcut for av_packet_alloc()+av_packet_ref().
 *
 * @return newly created AVPacket on success, NULL on error.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 */
func Av_packet_clone(src *AVPacket) *AVPacket {
    return (*AVPacket)(unsafe.Pointer(C.av_packet_clone(
        (*C.struct_AVPacket)(unsafe.Pointer(src)))))
}

/**
 * Free the packet, if the packet is reference counted, it will be
 * unreferenced first.
 *
 * @param pkt packet to be freed. The pointer will be set to NULL.
 * @note passing NULL is a no-op.
 */
func Av_packet_free(pkt **AVPacket)  {
    C.av_packet_free((**C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

                      
/**
 * Initialize optional fields of a packet with default values.
 *
 * Note, this does not touch the data and size members, which have to be
 * initialized separately.
 *
 * @param pkt packet
 *
 * @see av_packet_alloc
 * @see av_packet_unref
 *
 * @deprecated This function is deprecated. Once it's removed,
               sizeof(AVPacket) will not be a part of the ABI anymore.
 */

func Av_init_packet(pkt *AVPacket)  {
    C.av_init_packet((*C.struct_AVPacket)(unsafe.Pointer(pkt)))
}
      

/**
 * Allocate the payload of a packet and initialize its fields with
 * default values.
 *
 * @param pkt packet
 * @param size wanted payload size
 * @return 0 if OK, AVERROR_xxx otherwise
 */
func Av_new_packet(pkt *AVPacket, size int32) int32 {
    return int32(C.av_new_packet((*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.int(size)))
}

/**
 * Reduce packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param size new size
 */
func Av_shrink_packet(pkt *AVPacket, size int32)  {
    C.av_shrink_packet((*C.struct_AVPacket)(unsafe.Pointer(pkt)), C.int(size))
}

/**
 * Increase packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param grow_by number of bytes by which to increase the size of the packet
 */
func Av_grow_packet(pkt *AVPacket, grow_by int32) int32 {
    return int32(C.av_grow_packet((*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.int(grow_by)))
}

/**
 * Initialize a reference-counted packet from av_malloc()ed data.
 *
 * @param pkt packet to be initialized. This function will set the data, size,
 *        and buf fields, all others are left untouched.
 * @param data Data allocated by av_malloc() to be used as packet data. If this
 *        function returns successfully, the data is owned by the underlying AVBuffer.
 *        The caller may not access the data through other means.
 * @param size size of data in bytes, without the padding. I.e. the full buffer
 *        size is assumed to be size + AV_INPUT_BUFFER_PADDING_SIZE.
 *
 * @return 0 on success, a negative AVERROR on error
 */
func Av_packet_from_data(pkt *AVPacket, data *uint8, size int32) int32 {
    return int32(C.av_packet_from_data((*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        (*C.uchar)(unsafe.Pointer(data)), C.int(size)))
}

/**
 * Allocate new information of a packet.
 *
 * @param pkt packet
 * @param type side information type
 * @param size side information size
 * @return pointer to fresh allocated data or NULL otherwise
 */
func Av_packet_new_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                                 size uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_new_side_data(
        (*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), C.ulonglong(size))))
}

/**
 * Wrap an existing array as a packet side data.
 *
 * @param pkt packet
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             pkt.
 * @param size side information size
 * @return a non-negative number on success, a negative AVERROR code on
 *         failure. On failure, the packet is unchanged and the data remains
 *         owned by the caller.
 */
func Av_packet_add_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                            data *uint8, size uint64) int32 {
    return int32(C.av_packet_add_side_data((*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), (*C.uchar)(unsafe.Pointer(data)), 
        C.ulonglong(size)))
}

/**
 * Shrink the already allocated side data buffer
 *
 * @param pkt packet
 * @param type side information type
 * @param size new side information size
 * @return 0 on success, < 0 on failure
 */
func Av_packet_shrink_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                               size uint64) int32 {
    return int32(C.av_packet_shrink_side_data(
        (*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), C.ulonglong(size)))
}

/**
 * Get side information from packet.
 *
 * @param pkt packet
 * @param type desired side information type
 * @param size If supplied, *size will be set to the size of the side data
 *             or to zero if the desired side data is not present.
 * @return pointer to data if present or NULL otherwise
 */
func Av_packet_get_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                                 size *uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_get_side_data(
        (*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Pack a dictionary for use in side_data.
 *
 * @param dict The dictionary to pack.
 * @param size pointer to store the size of the returned data
 * @return pointer to data if successful, NULL otherwise
 */
func Av_packet_pack_dictionary(dict *AVDictionary, size *uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_pack_dictionary(
        (*C.struct_AVDictionary)(unsafe.Pointer(dict)), 
        (*C.ulonglong)(unsafe.Pointer(size)))))
}
/**
 * Unpack a dictionary from side_data.
 *
 * @param data data from side_data
 * @param size size of the data
 * @param dict the metadata storage dictionary
 * @return 0 on success, < 0 on failure
 */
func Av_packet_unpack_dictionary(data *uint8, size uint64,
                                dict **AVDictionary) int32 {
    return int32(C.av_packet_unpack_dictionary((*C.uchar)(unsafe.Pointer(data)), 
        C.ulonglong(size), (**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

/**
 * Convenience function to free all the side data stored.
 * All the other fields stay untouched.
 *
 * @param pkt packet
 */
func Av_packet_free_side_data(pkt *AVPacket)  {
    C.av_packet_free_side_data((*C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Setup a new reference to the data described by a given packet
 *
 * If src is reference-counted, setup dst as a new reference to the
 * buffer in src. Otherwise allocate a new buffer in dst and copy the
 * data from src into it.
 *
 * All the other fields are copied from src.
 *
 * @see av_packet_unref
 *
 * @param dst Destination packet. Will be completely overwritten.
 * @param src Source packet
 *
 * @return 0 on success, a negative AVERROR on error. On error, dst
 *         will be blank (as if returned by av_packet_alloc()).
 */
func Av_packet_ref(dst *AVPacket, src *AVPacket) int32 {
    return int32(C.av_packet_ref((*C.struct_AVPacket)(unsafe.Pointer(dst)), 
        (*C.struct_AVPacket)(unsafe.Pointer(src))))
}

/**
 * Wipe the packet.
 *
 * Unreference the buffer referenced by the packet and reset the
 * remaining packet fields to their default values.
 *
 * @param pkt The packet to be unreferenced.
 */
func Av_packet_unref(pkt *AVPacket)  {
    C.av_packet_unref((*C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Move every field in src to dst and reset src.
 *
 * @see av_packet_unref
 *
 * @param src Source packet, will be reset
 * @param dst Destination packet
 */
func Av_packet_move_ref(dst *AVPacket, src *AVPacket)  {
    C.av_packet_move_ref((*C.struct_AVPacket)(unsafe.Pointer(dst)), 
        (*C.struct_AVPacket)(unsafe.Pointer(src)))
}

/**
 * Copy only "properties" fields from src to dst.
 *
 * Properties for the purpose of this function are all the fields
 * beside those related to the packet data (buf, data, size)
 *
 * @param dst Destination packet
 * @param src Source packet
 *
 * @return 0 on success AVERROR on failure.
 */
func Av_packet_copy_props(dst *AVPacket, src *AVPacket) int32 {
    return int32(C.av_packet_copy_props((*C.struct_AVPacket)(unsafe.Pointer(dst)), 
        (*C.struct_AVPacket)(unsafe.Pointer(src))))
}

/**
 * Ensure the data described by a given packet is reference counted.
 *
 * @note This function does not ensure that the reference will be writable.
 *       Use av_packet_make_writable instead for that purpose.
 *
 * @see av_packet_ref
 * @see av_packet_make_writable
 *
 * @param pkt packet whose data should be made reference counted.
 *
 * @return 0 on success, a negative AVERROR on error. On failure, the
 *         packet is unchanged.
 */
func Av_packet_make_refcounted(pkt *AVPacket) int32 {
    return int32(C.av_packet_make_refcounted(
        (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Create a writable reference for the data described by a given packet,
 * avoiding data copy if possible.
 *
 * @param pkt Packet whose data should be made writable.
 *
 * @return 0 on success, a negative AVERROR on failure. On failure, the
 *         packet is unchanged.
 */
func Av_packet_make_writable(pkt *AVPacket) int32 {
    return int32(C.av_packet_make_writable((*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Convert valid timing fields (timestamps / durations) in a packet from one
 * timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be
 * ignored.
 *
 * @param pkt packet on which the conversion will be performed
 * @param tb_src source timebase, in which the timing fields in pkt are
 *               expressed
 * @param tb_dst destination timebase, to which the timing fields will be
 *               converted
 */
func Av_packet_rescale_ts(pkt *AVPacket, tb_src AVRational, tb_dst AVRational)  {
    C.av_packet_rescale_ts((*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        *(*C.struct_AVRational)(unsafe.Pointer(&tb_src)), 
        *(*C.struct_AVRational)(unsafe.Pointer(&tb_dst)))
}

/**
 * @}
 */

                          
