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

//#cgo pkg-config: libavcodec libavutil libavformat
//#include <stdio.h>  
//#include "libavcodec/codec_par.h"
//#include "libavcodec/defs.h"
//#include "libavcodec/packet.h"
//#include "libavutil/dict.h"
//#include "libavutil/log.h"
//#include "libavformat/avio.h"
//#include "libavformat/version_major.h"
//#include "libavformat/version.h"
//#include "libavutil/frame.h"
//#include "libavcodec/codec.h"
//#include "libavformat/avformat.h"
import "C"
import (
    "unsafe"
)

const AVPROBE_SCORE_RETRY = 25
const AVPROBE_SCORE_STREAM_RETRY = 24
const AVPROBE_SCORE_EXTENSION = 50
const AVPROBE_SCORE_MIME = 75
const AVPROBE_SCORE_MAX = 100
const AVPROBE_PADDING_SIZE = 32
const AVFMT_NOFILE =         0x0001 
const AVFMT_NEEDNUMBER =     0x0002  
const AVFMT_EXPERIMENTAL =   0x0004 
const AVFMT_SHOW_IDS =       0x0008  
const AVFMT_GLOBALHEADER =   0x0040  
const AVFMT_NOTIMESTAMPS =   0x0080  
const AVFMT_GENERIC_INDEX =  0x0100  
const AVFMT_TS_DISCONT =     0x0200  
const AVFMT_VARIABLE_FPS =   0x0400  
const AVFMT_NODIMENSIONS =   0x0800  
const AVFMT_NOSTREAMS =      0x1000  
const AVFMT_NOBINSEARCH =    0x2000  
const AVFMT_NOGENSEARCH =    0x4000  
const AVFMT_NO_BYTE_SEEK =   0x8000  
const AVFMT_ALLOW_FLUSH =   0x10000  
const AVFMT_TS_NONSTRICT =  0x20000  
const AVFMT_TS_NEGATIVE =   0x40000  
const AVFMT_SEEK_TO_PTS =    0x4000000  
const AVINDEX_KEYFRAME =  0x0001 
const AVINDEX_DISCARD_FRAME =   0x0002     
const AV_DISPOSITION_DEFAULT =               (1 << 0) 
const AV_DISPOSITION_DUB =                   (1 << 1) 
const AV_DISPOSITION_ORIGINAL =              (1 << 2) 
const AV_DISPOSITION_COMMENT =               (1 << 3) 
const AV_DISPOSITION_LYRICS =                (1 << 4) 
const AV_DISPOSITION_KARAOKE =               (1 << 5) 
const AV_DISPOSITION_FORCED =                (1 << 6) 
const AV_DISPOSITION_HEARING_IMPAIRED =      (1 << 7) 
const AV_DISPOSITION_VISUAL_IMPAIRED =       (1 << 8) 
const AV_DISPOSITION_CLEAN_EFFECTS =         (1 << 9) 
const AV_DISPOSITION_ATTACHED_PIC =          (1 << 10) 
const AV_DISPOSITION_TIMED_THUMBNAILS =      (1 << 11) 
const AV_DISPOSITION_NON_DIEGETIC =          (1 << 12) 
const AV_DISPOSITION_CAPTIONS =              (1 << 16) 
const AV_DISPOSITION_DESCRIPTIONS =          (1 << 17) 
const AV_DISPOSITION_METADATA =              (1 << 18) 
const AV_DISPOSITION_DEPENDENT =             (1 << 19) 
const AV_DISPOSITION_STILL_IMAGE =           (1 << 20) 
const AV_PTS_WRAP_IGNORE = 0
const AV_PTS_WRAP_ADD_OFFSET = 1
const AV_PTS_WRAP_SUB_OFFSET = -1
const AVSTREAM_EVENT_FLAG_METADATA_UPDATED =  0x0001 
const AVSTREAM_EVENT_FLAG_NEW_PACKETS =  (1 << 1) 
const AV_PROGRAM_RUNNING = 1
const AVFMTCTX_NOHEADER =       0x0001  
const AVFMTCTX_UNSEEKABLE =     0x0002  
const AVFMT_FLAG_GENPTS =        0x0001  
const AVFMT_FLAG_IGNIDX =        0x0002  
const AVFMT_FLAG_NONBLOCK =      0x0004  
const AVFMT_FLAG_IGNDTS =        0x0008  
const AVFMT_FLAG_NOFILLIN =      0x0010  
const AVFMT_FLAG_NOPARSE =       0x0020  
const AVFMT_FLAG_NOBUFFER =      0x0040  
const AVFMT_FLAG_CUSTOM_IO =     0x0080  
const AVFMT_FLAG_DISCARD_CORRUPT =   0x0100  
const AVFMT_FLAG_FLUSH_PACKETS =     0x0200  
const AVFMT_FLAG_BITEXACT =          0x0400 
const AVFMT_FLAG_SORT_DTS =     0x10000  
const AVFMT_FLAG_FAST_SEEK =    0x80000  
const AVFMT_FLAG_SHORTEST =    0x100000  
const AVFMT_FLAG_AUTO_BSF =    0x200000  
const FF_FDEBUG_TS =         0x0001 
const AVFMT_EVENT_FLAG_METADATA_UPDATED =  0x0001 
const AVFMT_AVOID_NEG_TS_AUTO = -1
const AVFMT_AVOID_NEG_TS_DISABLED = 0
const AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE = 1
const AVFMT_AVOID_NEG_TS_MAKE_ZERO = 2
const AVSEEK_FLAG_BACKWARD = 1
const AVSEEK_FLAG_BYTE = 2
const AVSEEK_FLAG_ANY = 4
const AVSEEK_FLAG_FRAME = 8
const AVSTREAM_INIT_IN_WRITE_HEADER = 0
const AVSTREAM_INIT_IN_INIT_OUTPUT = 1
const AV_FRAME_FILENAME_FLAGS_MULTIPLE = 1


                           
                           

/**
 * @file
 * @ingroup libavf
 * Main libavformat public API header
 */

/**
 * @defgroup libavf libavformat
 * I/O and Muxing/Demuxing Library
 *
 * Libavformat (lavf) is a library for dealing with various media container
 * formats. Its main two purposes are demuxing - i.e. splitting a media file
 * into component streams, and the reverse process of muxing - writing supplied
 * data in a specified container format. It also has an @ref lavf_io
 * "I/O module" which supports a number of protocols for accessing the data (e.g.
 * file, tcp, http and others).
 * Unless you are absolutely sure you won't use libavformat's network
 * capabilities, you should also call avformat_network_init().
 *
 * A supported input format is described by an AVInputFormat struct, conversely
 * an output format is described by AVOutputFormat. You can iterate over all
 * input/output formats using the  av_demuxer_iterate / av_muxer_iterate() functions.
 * The protocols layer is not part of the public API, so you can only get the names
 * of supported protocols with the avio_enum_protocols() function.
 *
 * Main lavf structure used for both muxing and demuxing is AVFormatContext,
 * which exports all information about the file being read or written. As with
 * most Libavformat structures, its size is not part of public ABI, so it cannot be
 * allocated on stack or directly with av_malloc(). To create an
 * AVFormatContext, use avformat_alloc_context() (some functions, like
 * avformat_open_input() might do that for you).
 *
 * Most importantly an AVFormatContext contains:
 * @li the @ref AVFormatContext.iformat "input" or @ref AVFormatContext.oformat
 * "output" format. It is either autodetected or set by user for input;
 * always set by user for output.
 * @li an @ref AVFormatContext.streams "array" of AVStreams, which describe all
 * elementary streams stored in the file. AVStreams are typically referred to
 * using their index in this array.
 * @li an @ref AVFormatContext.pb "I/O context". It is either opened by lavf or
 * set by user for input, always set by user for output (unless you are dealing
 * with an AVFMT_NOFILE format).
 *
 * @section lavf_options Passing options to (de)muxers
 * It is possible to configure lavf muxers and demuxers using the @ref avoptions
 * mechanism. Generic (format-independent) libavformat options are provided by
 * AVFormatContext, they can be examined from a user program by calling
 * av_opt_next() / av_opt_find() on an allocated AVFormatContext (or its AVClass
 * from avformat_get_class()). Private (format-specific) options are provided by
 * AVFormatContext.priv_data if and only if AVInputFormat.priv_class /
 * AVOutputFormat.priv_class of the corresponding format struct is non-NULL.
 * Further options may be provided by the @ref AVFormatContext.pb "I/O context",
 * if its AVClass is non-NULL, and the protocols layer. See the discussion on
 * nesting in @ref avoptions documentation to learn how to access those.
 *
 * @section urls
 * URL strings in libavformat are made of a scheme/protocol, a ':', and a
 * scheme specific string. URLs without a scheme and ':' used for local files
 * are supported but deprecated. "file:" should be used for local files.
 *
 * It is important that the scheme string is not taken from untrusted
 * sources without checks.
 *
 * Note that some schemes/protocols are quite powerful, allowing access to
 * both local and remote files, parts of them, concatenations of them, local
 * audio and video devices and so on.
 *
 * @{
 *
 * @defgroup lavf_decoding Demuxing
 * @{
 * Demuxers read a media file and split it into chunks of data (@em packets). A
 * @ref AVPacket "packet" contains one or more encoded frames which belongs to a
 * single elementary stream. In the lavf API this process is represented by the
 * avformat_open_input() function for opening a file, av_read_frame() for
 * reading a single packet and finally avformat_close_input(), which does the
 * cleanup.
 *
 * @section lavf_decoding_open Opening a media file
 * The minimum information required to open a file is its URL, which
 * is passed to avformat_open_input(), as in the following code:
 * @code
 * const char    *url = "file:in.mp3";
 * AVFormatContext *s = NULL;
 * int ret = avformat_open_input(&s, url, NULL, NULL);
 * if (ret < 0)
 *     abort();
 * @endcode
 * The above code attempts to allocate an AVFormatContext, open the
 * specified file (autodetecting the format) and read the header, exporting the
 * information stored there into s. Some formats do not have a header or do not
 * store enough information there, so it is recommended that you call the
 * avformat_find_stream_info() function which tries to read and decode a few
 * frames to find missing information.
 *
 * In some cases you might want to preallocate an AVFormatContext yourself with
 * avformat_alloc_context() and do some tweaking on it before passing it to
 * avformat_open_input(). One such case is when you want to use custom functions
 * for reading input data instead of lavf internal I/O layer.
 * To do that, create your own AVIOContext with avio_alloc_context(), passing
 * your reading callbacks to it. Then set the @em pb field of your
 * AVFormatContext to newly created AVIOContext.
 *
 * Since the format of the opened file is in general not known until after
 * avformat_open_input() has returned, it is not possible to set demuxer private
 * options on a preallocated context. Instead, the options should be passed to
 * avformat_open_input() wrapped in an AVDictionary:
 * @code
 * AVDictionary *options = NULL;
 * av_dict_set(&options, "video_size", "640x480", 0);
 * av_dict_set(&options, "pixel_format", "rgb24", 0);
 *
 * if (avformat_open_input(&s, url, NULL, &options) < 0)
 *     abort();
 * av_dict_free(&options);
 * @endcode
 * This code passes the private options 'video_size' and 'pixel_format' to the
 * demuxer. They would be necessary for e.g. the rawvideo demuxer, since it
 * cannot know how to interpret raw video data otherwise. If the format turns
 * out to be something different than raw video, those options will not be
 * recognized by the demuxer and therefore will not be applied. Such unrecognized
 * options are then returned in the options dictionary (recognized options are
 * consumed). The calling program can handle such unrecognized options as it
 * wishes, e.g.
 * @code
 * AVDictionaryEntry *e;
 * if (e = av_dict_get(options, "", NULL, AV_DICT_IGNORE_SUFFIX)) {
 *     fprintf(stderr, "Option %s not recognized by the demuxer.\n", e->key);
 *     abort();
 * }
 * @endcode
 *
 * After you have finished reading the file, you must close it with
 * avformat_close_input(). It will free everything associated with the file.
 *
 * @section lavf_decoding_read Reading from an opened file
 * Reading data from an opened AVFormatContext is done by repeatedly calling
 * av_read_frame() on it. Each call, if successful, will return an AVPacket
 * containing encoded data for one AVStream, identified by
 * AVPacket.stream_index. This packet may be passed straight into the libavcodec
 * decoding functions avcodec_send_packet() or avcodec_decode_subtitle2() if the
 * caller wishes to decode the data.
 *
 * AVPacket.pts, AVPacket.dts and AVPacket.duration timing information will be
 * set if known. They may also be unset (i.e. AV_NOPTS_VALUE for
 * pts/dts, 0 for duration) if the stream does not provide them. The timing
 * information will be in AVStream.time_base units, i.e. it has to be
 * multiplied by the timebase to convert them to seconds.
 *
 * A packet returned by av_read_frame() is always reference-counted,
 * i.e. AVPacket.buf is set and the user may keep it indefinitely.
 * The packet must be freed with av_packet_unref() when it is no
 * longer needed.
 *
 * @section lavf_decoding_seek Seeking
 * @}
 *
 * @defgroup lavf_encoding Muxing
 * @{
 * Muxers take encoded data in the form of @ref AVPacket "AVPackets" and write
 * it into files or other output bytestreams in the specified container format.
 *
 * The main API functions for muxing are avformat_write_header() for writing the
 * file header, av_write_frame() / av_interleaved_write_frame() for writing the
 * packets and av_write_trailer() for finalizing the file.
 *
 * At the beginning of the muxing process, the caller must first call
 * avformat_alloc_context() to create a muxing context. The caller then sets up
 * the muxer by filling the various fields in this context:
 *
 * - The @ref AVFormatContext.oformat "oformat" field must be set to select the
 *   muxer that will be used.
 * - Unless the format is of the AVFMT_NOFILE type, the @ref AVFormatContext.pb
 *   "pb" field must be set to an opened IO context, either returned from
 *   avio_open2() or a custom one.
 * - Unless the format is of the AVFMT_NOSTREAMS type, at least one stream must
 *   be created with the avformat_new_stream() function. The caller should fill
 *   the @ref AVStream.codecpar "stream codec parameters" information, such as the
 *   codec @ref AVCodecParameters.codec_type "type", @ref AVCodecParameters.codec_id
 *   "id" and other parameters (e.g. width / height, the pixel or sample format,
 *   etc.) as known. The @ref AVStream.time_base "stream timebase" should
 *   be set to the timebase that the caller desires to use for this stream (note
 *   that the timebase actually used by the muxer can be different, as will be
 *   described later).
 * - It is advised to manually initialize only the relevant fields in
 *   AVCodecParameters, rather than using @ref avcodec_parameters_copy() during
 *   remuxing: there is no guarantee that the codec context values remain valid
 *   for both input and output format contexts.
 * - The caller may fill in additional information, such as @ref
 *   AVFormatContext.metadata "global" or @ref AVStream.metadata "per-stream"
 *   metadata, @ref AVFormatContext.chapters "chapters", @ref
 *   AVFormatContext.programs "programs", etc. as described in the
 *   AVFormatContext documentation. Whether such information will actually be
 *   stored in the output depends on what the container format and the muxer
 *   support.
 *
 * When the muxing context is fully set up, the caller must call
 * avformat_write_header() to initialize the muxer internals and write the file
 * header. Whether anything actually is written to the IO context at this step
 * depends on the muxer, but this function must always be called. Any muxer
 * private options must be passed in the options parameter to this function.
 *
 * The data is then sent to the muxer by repeatedly calling av_write_frame() or
 * av_interleaved_write_frame() (consult those functions' documentation for
 * discussion on the difference between them; only one of them may be used with
 * a single muxing context, they should not be mixed). Do note that the timing
 * information on the packets sent to the muxer must be in the corresponding
 * AVStream's timebase. That timebase is set by the muxer (in the
 * avformat_write_header() step) and may be different from the timebase
 * requested by the caller.
 *
 * Once all the data has been written, the caller must call av_write_trailer()
 * to flush any buffered packets and finalize the output file, then close the IO
 * context (if any) and finally free the muxing context with
 * avformat_free_context().
 * @}
 *
 * @defgroup lavf_io I/O Read/Write
 * @{
 * @section lavf_io_dirlist Directory listing
 * The directory listing API makes it possible to list files on remote servers.
 *
 * Some of possible use cases:
 * - an "open file" dialog to choose files from a remote location,
 * - a recursive media finder providing a player with an ability to play all
 * files from a given directory.
 *
 * @subsection lavf_io_dirlist_open Opening a directory
 * At first, a directory needs to be opened by calling avio_open_dir()
 * supplied with a URL and, optionally, ::AVDictionary containing
 * protocol-specific parameters. The function returns zero or positive
 * integer and allocates AVIODirContext on success.
 *
 * @code
 * AVIODirContext *ctx = NULL;
 * if (avio_open_dir(&ctx, "smb://example.com/some_dir", NULL) < 0) {
 *     fprintf(stderr, "Cannot open directory.\n");
 *     abort();
 * }
 * @endcode
 *
 * This code tries to open a sample directory using smb protocol without
 * any additional parameters.
 *
 * @subsection lavf_io_dirlist_read Reading entries
 * Each directory's entry (i.e. file, another directory, anything else
 * within ::AVIODirEntryType) is represented by AVIODirEntry.
 * Reading consecutive entries from an opened AVIODirContext is done by
 * repeatedly calling avio_read_dir() on it. Each call returns zero or
 * positive integer if successful. Reading can be stopped right after the
 * NULL entry has been read -- it means there are no entries left to be
 * read. The following code reads all entries from a directory associated
 * with ctx and prints their names to standard output.
 * @code
 * AVIODirEntry *entry = NULL;
 * for (;;) {
 *     if (avio_read_dir(ctx, &entry) < 0) {
 *         fprintf(stderr, "Cannot list directory.\n");
 *         abort();
 *     }
 *     if (!entry)
 *         break;
 *     printf("%s\n", entry->name);
 *     avio_free_directory_entry(&entry);
 * }
 * @endcode
 * @}
 *
 * @defgroup lavf_codec Demuxers
 * @{
 * @defgroup lavf_codec_native Native Demuxers
 * @{
 * @}
 * @defgroup lavf_codec_wrappers External library wrappers
 * @{
 * @}
 * @}
 * @defgroup lavf_protos I/O Protocols
 * @{
 * @}
 * @defgroup lavf_internal Internal
 * @{
 * @}
 * @}
 */

                              

                                 
                            
                              

                           
                          

                 
                                      
                        
/* When included as part of the ffmpeg build, only include the major version
 * to avoid unnecessary rebuilds. When included externally, keep including
 * the full version information. */
                                

                            
                             
      

// type AVFormatContext
// type AVFrame
// type AVDeviceInfoList

/**
 * @defgroup metadata_api Public Metadata API
 * @{
 * @ingroup libavf
 * The metadata API allows libavformat to export metadata tags to a client
 * application when demuxing. Conversely it allows a client application to
 * set metadata when muxing.
 *
 * Metadata is exported or set as pairs of key/value strings in the 'metadata'
 * fields of the AVFormatContext, AVStream, AVChapter and AVProgram structs
 * using the @ref lavu_dict "AVDictionary" API. Like all strings in FFmpeg,
 * metadata is assumed to be UTF-8 encoded Unicode. Note that metadata
 * exported by demuxers isn't checked to be valid UTF-8 in most cases.
 *
 * Important concepts to keep in mind:
 * -  Keys are unique; there can never be 2 tags with the same key. This is
 *    also meant semantically, i.e., a demuxer should not knowingly produce
 *    several keys that are literally different but semantically identical.
 *    E.g., key=Author5, key=Author6. In this example, all authors must be
 *    placed in the same tag.
 * -  Metadata is flat, not hierarchical; there are no subtags. If you
 *    want to store, e.g., the email address of the child of producer Alice
 *    and actor Bob, that could have key=alice_and_bobs_childs_email_address.
 * -  Several modifiers can be applied to the tag name. This is done by
 *    appending a dash character ('-') and the modifier name in the order
 *    they appear in the list below -- e.g. foo-eng-sort, not foo-sort-eng.
 *    -  language -- a tag whose value is localized for a particular language
 *       is appended with the ISO 639-2/B 3-letter language code.
 *       For example: Author-ger=Michael, Author-eng=Mike
 *       The original/default language is in the unqualified "Author" tag.
 *       A demuxer should set a default if it sets any translated tag.
 *    -  sorting  -- a modified version of a tag that should be used for
 *       sorting will have '-sort' appended. E.g. artist="The Beatles",
 *       artist-sort="Beatles, The".
 * - Some protocols and demuxers support metadata updates. After a successful
 *   call to av_read_frame(), AVFormatContext.event_flags or AVStream.event_flags
 *   will be updated to indicate if metadata changed. In order to detect metadata
 *   changes on a stream, you need to loop through all streams in the AVFormatContext
 *   and check their individual event_flags.
 *
 * -  Demuxers attempt to export metadata in a generic format, however tags
 *    with no generic equivalents are left as they are stored in the container.
 *    Follows a list of generic tag names:
 *
 @verbatim
 album        -- name of the set this work belongs to
 album_artist -- main creator of the set/album, if different from artist.
                 e.g. "Various Artists" for compilation albums.
 artist       -- main creator of the work
 comment      -- any additional description of the file.
 composer     -- who composed the work, if different from artist.
 copyright    -- name of copyright holder.
 creation_time-- date when the file was created, preferably in ISO 8601.
 date         -- date when the work was created, preferably in ISO 8601.
 disc         -- number of a subset, e.g. disc in a multi-disc collection.
 encoder      -- name/settings of the software/hardware that produced the file.
 encoded_by   -- person/group who created the file.
 filename     -- original name of the file.
 genre        -- <self-evident>.
 language     -- main language in which the work is performed, preferably
                 in ISO 639-2 format. Multiple languages can be specified by
                 separating them with commas.
 performer    -- artist who performed the work, if different from artist.
                 E.g for "Also sprach Zarathustra", artist would be "Richard
                 Strauss" and performer "London Philharmonic Orchestra".
 publisher    -- name of the label/publisher.
 service_name     -- name of the service in broadcasting (channel name).
 service_provider -- name of the service provider in broadcasting.
 title        -- name of the work.
 track        -- number of this work in the set, can be in form current/total.
 variant_bitrate -- the total bitrate of the bitrate variant that the current stream is part of
 @endverbatim
 *
 * Look in the examples section for an application example how to use the Metadata API.
 *
 * @}
 */

/* packet functions */


/**
 * Allocate and read the payload of a packet and initialize its
 * fields with default values.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size desired payload size
 * @return >0 (read size) if OK, AVERROR_xxx otherwise
 */
func Av_get_packet(s *AVIOContext, pkt *AVPacket, size int32) int32 {
    return int32(C.av_get_packet((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.struct_AVPacket)(unsafe.Pointer(pkt)), C.int(size)))
}


/**
 * Read data and append it to the current content of the AVPacket.
 * If pkt->size is 0 this is identical to av_get_packet.
 * Note that this uses av_grow_packet and thus involves a realloc
 * which is inefficient. Thus this function should only be used
 * when there is no reasonable way to know (an upper bound of)
 * the final size.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size amount of data to read
 * @return >0 (read size) if OK, AVERROR_xxx otherwise, previous data
 *         will not be lost even if an error occurs.
 */
func Av_append_packet(s *AVIOContext, pkt *AVPacket, size int32) int32 {
    return int32(C.av_append_packet((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.struct_AVPacket)(unsafe.Pointer(pkt)), C.int(size)))
}

/*************************************************/
/* input/output formats */

type AVCodecTag struct {
}


/**
 * This structure contains the data a format has to probe a file.
 */
type AVProbeData struct {
    Filename *byte
    Buf *uint8
    Buf_size int32
    Mime_type *byte
}





///< score for file extension
///< score for file mime type
///< maximum score

///< extra allocated bytes at the end of the probe buffer

/// Demuxer will use avio_open, no opened file should be provided by the caller.

/**< Needs '%d' in filename. */
/**
 * The muxer/demuxer is experimental and should be used with caution.
 *
 * - demuxers: will not be selected automatically by probing, must be specified
 *             explicitly.
 */

/**< Show format stream IDs numbers. */
/**< Format wants global header. */
/**< Format does not need / have any timestamps. */
/**< Use generic index building code. */
/**< Format allows timestamp discontinuities. Note, muxers always require valid (monotone) timestamps */
/**< Format allows variable fps. */
/**< Format does not need width/height */
/**< Format does not require any streams */
/**< Format does not allow to fall back on binary search via read_timestamp */
/**< Format does not allow to fall back on generic search */
/**< Format does not allow seeking by bytes */
                      
/**< @deprecated: Just send a NULL packet if you want to flush a muxer. */
      
/**< Format does not require strictly
                                        increasing timestamps, but they must
                                        still be monotonic */
/**< Format allows muxing negative
                                        timestamps. If not set the timestamp
                                        will be shifted in av_write_frame and
                                        av_interleaved_write_frame so they
                                        start from 0.
                                        The user or muxer can override this through
                                        AVFormatContext.avoid_negative_ts
                                        */

/**< Seeking is based on PTS */

/**
 * @addtogroup lavf_encoding
 * @{
 */
type AVOutputFormat struct {
    Name *byte
    Long_name *byte
    Mime_type *byte
    Extensions *byte
    Audio_codec AVCodecID
    Video_codec AVCodecID
    Subtitle_codec AVCodecID
    Flags int32
    Codec_tag **AVCodecTag
    Priv_class *AVClass
}

/**
 * @}
 */

/**
 * @addtogroup lavf_decoding
 * @{
 */
type AVInputFormat struct {
    Name *byte
    Long_name *byte
    Flags int32
    Extensions *byte
    Codec_tag **AVCodecTag
    Priv_class *AVClass
    Mime_type *byte
    Raw_codec_id int32
    Priv_data_size int32
    Flags_internal int32
    Read_probe uintptr
    Read_header uintptr
    Read_packet uintptr
    Read_close uintptr
    Read_seek uintptr
    Read_timestamp uintptr
    Read_play uintptr
    Read_pause uintptr
    Read_seek2 uintptr
    Get_device_list uintptr
}

/**
 * @}
 */

type AVStreamParseType int32
const (
    AVSTREAM_PARSE_NONE AVStreamParseType = iota
    AVSTREAM_PARSE_FULL
    AVSTREAM_PARSE_HEADERS
    AVSTREAM_PARSE_TIMESTAMPS
    AVSTREAM_PARSE_FULL_ONCE
    AVSTREAM_PARSE_FULL_RAW
)


// type AVIndexEntry

/**
 * The stream should be chosen by default among other streams of the same type,
 * unless the user has explicitly specified otherwise.
 */

/**
 * The stream is not in original language.
 *
 * @note AV_DISPOSITION_ORIGINAL is the inverse of this disposition. At most
 *       one of them should be set in properly tagged streams.
 * @note This disposition may apply to any stream type, not just audio.
 */

/**
 * The stream is in original language.
 *
 * @see the notes for AV_DISPOSITION_DUB
 */

/**
 * The stream is a commentary track.
 */

/**
 * The stream contains song lyrics.
 */

/**
 * The stream contains karaoke audio.
 */


/**
 * Track should be used during playback by default.
 * Useful for subtitle track that should be displayed
 * even when user did not explicitly ask for subtitles.
 */

/**
 * The stream is intended for hearing impaired audiences.
 */

/**
 * The stream is intended for visually impaired audiences.
 */

/**
 * The audio stream contains music and sound effects without voice.
 */

/**
 * The stream is stored in the file as an attached picture/"cover art" (e.g.
 * APIC frame in ID3v2). The first (usually only) packet associated with it
 * will be returned among the first few packets read from the file unless
 * seeking takes place. It can also be accessed at any time in
 * AVStream.attached_pic.
 */

/**
 * The stream is sparse, and contains thumbnail images, often corresponding
 * to chapter markers. Only ever used with AV_DISPOSITION_ATTACHED_PIC.
 */


/**
 * The stream is intended to be mixed with a spatial audio track. For example,
 * it could be used for narration or stereo music, and may remain unchanged by
 * listener head rotation.
 */


/**
 * The subtitle stream contains captions, providing a transcription and possibly
 * a translation of audio. Typically intended for hearing-impaired audiences.
 */

/**
 * The subtitle stream contains a textual description of the video content.
 * Typically intended for visually-impaired audiences or for the cases where the
 * video cannot be seen.
 */

/**
 * The subtitle stream contains time-aligned metadata that is not intended to be
 * directly presented to the user.
 */

/**
 * The audio stream is intended to be mixed with another stream before
 * presentation.
 * Corresponds to mix_type=0 in mpegts.
 */

/**
 * The video stream contains still images.
 */


/**
 * @return The AV_DISPOSITION_* flag corresponding to disp or a negative error
 *         code if disp does not correspond to a known stream disposition.
 */
func Av_disposition_from_string(disp *byte) int32 {
    return int32(C.av_disposition_from_string((*C.char)(unsafe.Pointer(disp))))
}

/**
 * @param disposition a combination of AV_DISPOSITION_* values
 * @return The string description corresponding to the lowest set bit in
 *         disposition. NULL when the lowest set bit does not correspond
 *         to a known disposition or when disposition is 0.
 */
func Av_disposition_to_string(disposition int32) string {
    return C.GoString(C.av_disposition_to_string(C.int(disposition)))
}

/**
 * Options for behavior on timestamp wrap detection.
 */
///< ignore the wrap
///< add the format specific offset on wrap detection
///< subtract the format specific offset on wrap detection

/**
 * Stream structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVStream) must not be used outside libav*.
 */
type AVStream struct {
    Av_class *AVClass
    Index int32
    Id int32
    Codecpar *AVCodecParameters
    Priv_data unsafe.Pointer
    Time_base AVRational
    Start_time int64
    Duration int64
    Nb_frames int64
    Disposition int32
    Discard AVDiscard
    Sample_aspect_ratio AVRational
    Metadata *AVDictionary
    Avg_frame_rate AVRational
    Attached_pic AVPacket
    Side_data *AVPacketSideData
    Nb_side_data int32
    Event_flags int32
    R_frame_rate AVRational
    Pts_wrap_bits int32
}


func Av_stream_get_parser(s *AVStream) *AVCodecParserContext {
    return (*AVCodecParserContext)(unsafe.Pointer(C.av_stream_get_parser(
        (*C.struct_AVStream)(unsafe.Pointer(s)))))
}

                      
/**
 * Returns the pts of the last muxed packet + its duration
 *
 * the retuned value is undefined when used with a demuxer.
 */

func    Av_stream_get_end_pts(st *AVStream) int64 {
    return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(unsafe.Pointer(st))))
}
      



/**
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVProgram) must not be used outside libav*.
 */
type AVProgram struct {
    Id int32
    Flags int32
    Discard AVDiscard
    Stream_index *uint32
    Nb_stream_indexes uint32
    Metadata *AVDictionary
    Program_num int32
    Pmt_pid int32
    Pcr_pid int32
    Pmt_version int32
    Start_time int64
    End_time int64
    Pts_wrap_reference int64
    Pts_wrap_behavior int32
}


/**< signal that no header is present
                                         (streams are added dynamically) */
/**< signal that the stream is definitely
                                         not seekable, and attempts to call the
                                         seek function will fail. For some
                                         network protocols (e.g. HLS), this can
                                         change dynamically at runtime. */

type AVChapter struct {
    Id int64
    Time_base AVRational
    Start int64
    End int64
    Metadata *AVDictionary
}



/**
 * Callback used by devices to communicate with application.
 */
type av_format_control_message C.av_format_control_message

type AVOpenCallback C.AVOpenCallback

/**
 * The duration of a video can be estimated through various ways, and this enum can be used
 * to know how the duration was estimated.
 */
type AVDurationEstimationMethod int32
const (
    AVFMT_DURATION_FROM_PTS AVDurationEstimationMethod = iota
    AVFMT_DURATION_FROM_STREAM
    AVFMT_DURATION_FROM_BITRATE
)


/**
 * Format I/O context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVFormatContext) must not be used outside libav*, use
 * avformat_alloc_context() to create an AVFormatContext.
 *
 * Fields can be accessed through AVOptions (av_opt*),
 * the name string used matches the associated command line parameter name and
 * can be found in libavformat/options_table.h.
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 */
type AVFormatContext struct {
    Av_class *AVClass
    Iformat *AVInputFormat
    Oformat *AVOutputFormat
    Priv_data unsafe.Pointer
    Pb *AVIOContext
    Ctx_flags int32
    Nb_streams uint32
    Streams **AVStream
    Url *byte
    Start_time int64
    Duration int64
    Bit_rate int64
    Packet_size uint32
    Max_delay int32
    Flags int32
    Probesize int64
    Max_analyze_duration int64
    Key *uint8
    Keylen int32
    Nb_programs uint32
    Programs **AVProgram
    Video_codec_id AVCodecID
    Audio_codec_id AVCodecID
    Subtitle_codec_id AVCodecID
    Max_index_size uint32
    Max_picture_buffer uint32
    Nb_chapters uint32
    Chapters **AVChapter
    Metadata *AVDictionary
    Start_time_realtime int64
    Fps_probe_size int32
    Error_recognition int32
    Interrupt_callback AVIOInterruptCB
    Debug int32
    Max_interleave_delta int64
    Strict_std_compliance int32
    Event_flags int32
    Max_ts_probe int32
    Avoid_negative_ts int32
    Ts_id int32
    Audio_preload int32
    Max_chunk_duration int32
    Max_chunk_size int32
    Use_wallclock_as_timestamps int32
    Avio_flags int32
    Duration_estimation_method AVDurationEstimationMethod
    Skip_initial_bytes int64
    Correct_ts_overflow uint32
    Seek2any int32
    Flush_packets int32
    Probe_score int32
    Format_probesize int32
    Codec_whitelist *byte
    Format_whitelist *byte
    Io_repositioned int32
    Video_codec *AVCodec
    Audio_codec *AVCodec
    Subtitle_codec *AVCodec
    Data_codec *AVCodec
    Metadata_header_padding int32
    Opaque unsafe.Pointer
    Control_message_cb av_format_control_message
    Output_ts_offset int64
    Dump_separator *uint8
    Data_codec_id AVCodecID
    Protocol_whitelist *byte
    Io_open uintptr
    Io_close uintptr
    Protocol_blacklist *byte
    Max_streams int32
    Skip_estimate_duration_from_pts int32
    Max_probe_packets int32
    Io_close2 uintptr
}


/**
 * This function will cause global side data to be injected in the next packet
 * of each stream as well as after any subsequent seek.
 *
 * @note global side data is always available in every AVStream's
 *       @ref AVCodecParameters.coded_side_data "codecpar side data" array, and
 *       in a @ref AVCodecContext.coded_side_data "decoder's side data" array if
 *       initialized with said stream's codecpar.
 * @see av_packet_side_data_get()
 */
func Av_format_inject_global_side_data(s *AVFormatContext)  {
    C.av_format_inject_global_side_data(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

/**
 * Returns the method used to set ctx->duration.
 *
 * @return AVFMT_DURATION_FROM_PTS, AVFMT_DURATION_FROM_STREAM, or AVFMT_DURATION_FROM_BITRATE.
 */
func Av_fmt_ctx_get_duration_estimation_method(ctx *AVFormatContext) AVDurationEstimationMethod {
    return AVDurationEstimationMethod(C.av_fmt_ctx_get_duration_estimation_method(
        (*C.struct_AVFormatContext)(unsafe.Pointer(ctx))))
}

/**
 * @defgroup lavf_core Core functions
 * @ingroup libavf
 *
 * Functions for querying libavformat capabilities, allocating core structures,
 * etc.
 * @{
 */

/**
 * Return the LIBAVFORMAT_VERSION_INT constant.
 */
func Avformat_version() uint32 {
    return uint32(C.avformat_version())
}

/**
 * Return the libavformat build-time configuration.
 */
func Avformat_configuration() string {
    return C.GoString(C.avformat_configuration())
}

/**
 * Return the libavformat license.
 */
func Avformat_license() string {
    return C.GoString(C.avformat_license())
}

/**
 * Do global initialization of network libraries. This is optional,
 * and not recommended anymore.
 *
 * This functions only exists to work around thread-safety issues
 * with older GnuTLS or OpenSSL libraries. If libavformat is linked
 * to newer versions of those libraries, or if you do not use them,
 * calling this function is unnecessary. Otherwise, you need to call
 * this function before any other threads using them are started.
 *
 * This function will be deprecated once support for older GnuTLS and
 * OpenSSL libraries is removed, and this function has no purpose
 * anymore.
 */
func Avformat_network_init() int32 {
    return int32(C.avformat_network_init())
}

/**
 * Undo the initialization done by avformat_network_init. Call it only
 * once for each time you called avformat_network_init.
 */
func Avformat_network_deinit() int32 {
    return int32(C.avformat_network_deinit())
}

/**
 * Iterate over all registered muxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered muxer or NULL when the iteration is
 *         finished
 */
func Av_muxer_iterate(opaque *unsafe.Pointer) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_muxer_iterate(opaque)))
}

/**
 * Iterate over all registered demuxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state.
 *               Must point to NULL to start the iteration.
 *
 * @return the next registered demuxer or NULL when the iteration is
 *         finished
 */
func Av_demuxer_iterate(opaque *unsafe.Pointer) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_demuxer_iterate(opaque)))
}

/**
 * Allocate an AVFormatContext.
 * avformat_free_context() can be used to free the context and everything
 * allocated by the framework within it.
 */
func Avformat_alloc_context() *AVFormatContext {
    return (*AVFormatContext)(unsafe.Pointer(C.avformat_alloc_context()))
}

/**
 * Free an AVFormatContext and all its streams.
 * @param s context to free
 */
func Avformat_free_context(s *AVFormatContext)  {
    C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

/**
 * Get the AVClass for AVFormatContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Avformat_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avformat_get_class()))
}

/**
 * Get the AVClass for AVStream. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Av_stream_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.av_stream_get_class()))
}

/**
 * Add a new stream to a media file.
 *
 * When demuxing, it is called by the demuxer in read_header(). If the
 * flag AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also
 * be called in read_packet().
 *
 * When muxing, should be called by the user before avformat_write_header().
 *
 * User is required to call avformat_free_context() to clean up the allocation
 * by avformat_new_stream().
 *
 * @param s media file handle
 * @param c unused, does nothing
 *
 * @return newly created stream or NULL on error.
 */
func Avformat_new_stream(s *AVFormatContext, c *AVCodec) *AVStream {
    return (*AVStream)(unsafe.Pointer(C.avformat_new_stream(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (*C.struct_AVCodec)(unsafe.Pointer(c)))))
}

                             
/**
 * Wrap an existing array as stream side data.
 *
 * @param st   stream
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             st.
 * @param size side information size
 *
 * @return zero on success, a negative AVERROR code on failure. On failure,
 *         the stream is unchanged and the data remains owned by the caller.
 * @deprecated use av_packet_side_data_add() with the stream's
 *             @ref AVCodecParameters.coded_side_data "codecpar side data"
 */

func Av_stream_add_side_data(st *AVStream, typex AVPacketSideDataType,
                            data *uint8, size uint64) int32 {
    return int32(C.av_stream_add_side_data((*C.struct_AVStream)(unsafe.Pointer(st)), 
        C.enum_AVPacketSideDataType(typex), (*C.uchar)(unsafe.Pointer(data)), 
        C.ulonglong(size)))
}

/**
 * Allocate new information from stream.
 *
 * @param stream stream
 * @param type   desired side information type
 * @param size   side information size
 *
 * @return pointer to fresh allocated data or NULL otherwise
 * @deprecated use av_packet_side_data_new() with the stream's
 *             @ref AVCodecParameters.coded_side_data "codecpar side data"
 */

func Av_stream_new_side_data(stream *AVStream,
                                 typex AVPacketSideDataType, size uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_stream_new_side_data(
        (*C.struct_AVStream)(unsafe.Pointer(stream)), 
        C.enum_AVPacketSideDataType(typex), C.ulonglong(size))))
}
/**
 * Get side information from stream.
 *
 * @param stream stream
 * @param type   desired side information type
 * @param size   If supplied, *size will be set to the size of the side data
 *               or to zero if the desired side data is not present.
 *
 * @return pointer to data if present or NULL otherwise
 * @deprecated use av_packet_side_data_get() with the stream's
 *             @ref AVCodecParameters.coded_side_data "codecpar side data"
 */

func Av_stream_get_side_data(stream *AVStream,
                                 typex AVPacketSideDataType, size *uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_stream_get_side_data(
        (*C.struct_AVStream)(unsafe.Pointer(stream)), 
        C.enum_AVPacketSideDataType(typex), (*C.ulonglong)(unsafe.Pointer(size)))))
}
      

func Av_new_program(s *AVFormatContext, id int32) *AVProgram {
    return (*AVProgram)(unsafe.Pointer(C.av_new_program(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(id))))
}

/**
 * @}
 */


/**
 * Allocate an AVFormatContext for an output format.
 * avformat_free_context() can be used to free the context and
 * everything allocated by the framework within it.
 *
 * @param ctx           pointee is set to the created format context,
 *                      or to NULL in case of failure
 * @param oformat       format to use for allocating the context, if NULL
 *                      format_name and filename are used instead
 * @param format_name   the name of output format to use for allocating the
 *                      context, if NULL filename is used instead
 * @param filename      the name of the filename to use for allocating the
 *                      context, may be NULL
 *
 * @return  >= 0 in case of success, a negative AVERROR code in case of
 *          failure
 */
func Avformat_alloc_output_context2(ctx **AVFormatContext, oformat *AVOutputFormat,
                                   format_name *byte, filename *byte) int32 {
    return int32(C.avformat_alloc_output_context2(
        (**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVOutputFormat)(unsafe.Pointer(oformat)), 
        (*C.char)(unsafe.Pointer(format_name)), 
        (*C.char)(unsafe.Pointer(filename))))
}

/**
 * @addtogroup lavf_decoding
 * @{
 */

/**
 * Find AVInputFormat based on the short name of the input format.
 */
func Av_find_input_format(short_name *byte) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_find_input_format(
        (*C.char)(unsafe.Pointer(short_name)))))
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 */
func Av_probe_input_format(pd *AVProbeData, is_opened int32) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_probe_input_format(
        (*C.struct_AVProbeData)(unsafe.Pointer(pd)), C.int(is_opened))))
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_max A probe score larger that this is required to accept a
 *                  detection, the variable is set to the actual detection
 *                  score afterwards.
 *                  If the score is <= AVPROBE_SCORE_MAX / 4 it is recommended
 *                  to retry with a larger probe buffer.
 */
func Av_probe_input_format2(pd *AVProbeData,
                                            is_opened int32, score_max *int32) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_probe_input_format2(
        (*C.struct_AVProbeData)(unsafe.Pointer(pd)), C.int(is_opened), 
        (*C.int)(unsafe.Pointer(score_max)))))
}

/**
 * Guess the file format.
 *
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_ret The score of the best detection.
 */
func Av_probe_input_format3(pd *AVProbeData,
                                            is_opened int32, score_ret *int32) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_probe_input_format3(
        (*C.struct_AVProbeData)(unsafe.Pointer(pd)), C.int(is_opened), 
        (*C.int)(unsafe.Pointer(score_ret)))))
}

/**
 * Probe a bytestream to determine the input format. Each time a probe returns
 * with a score that is too low, the probe buffer size is increased and another
 * attempt is made. When the maximum probe size is reached, the input format
 * with the highest score is returned.
 *
 * @param pb             the bytestream to probe
 * @param fmt            the input format is put here
 * @param url            the url of the stream
 * @param logctx         the log context
 * @param offset         the offset within the bytestream to probe from
 * @param max_probe_size the maximum probe buffer size (zero for default)
 *
 * @return the score in case of success, a negative value corresponding to an
 *         the maximal score is AVPROBE_SCORE_MAX
 *         AVERROR code otherwise
 */
func Av_probe_input_buffer2(pb *AVIOContext, fmt **AVInputFormat,
                           url *byte, logctx unsafe.Pointer,
                           offset uint32, max_probe_size uint32) int32 {
    return int32(C.av_probe_input_buffer2(
        (*C.struct_AVIOContext)(unsafe.Pointer(pb)), 
        (**C.struct_AVInputFormat)(unsafe.Pointer(fmt)), 
        (*C.char)(unsafe.Pointer(url)), logctx, C.uint(offset), 
        C.uint(max_probe_size)))
}

/**
 * Like av_probe_input_buffer2() but returns 0 on success
 */
func Av_probe_input_buffer(pb *AVIOContext, fmt **AVInputFormat,
                          url *byte, logctx unsafe.Pointer,
                          offset uint32, max_probe_size uint32) int32 {
    return int32(C.av_probe_input_buffer((*C.struct_AVIOContext)(unsafe.Pointer(pb)), 
        (**C.struct_AVInputFormat)(unsafe.Pointer(fmt)), 
        (*C.char)(unsafe.Pointer(url)), logctx, C.uint(offset), 
        C.uint(max_probe_size)))
}

/**
 * Open an input stream and read the header. The codecs are not opened.
 * The stream must be closed with avformat_close_input().
 *
 * @param ps       Pointer to user-supplied AVFormatContext (allocated by
 *                 avformat_alloc_context). May be a pointer to NULL, in
 *                 which case an AVFormatContext is allocated by this
 *                 function and written into ps.
 *                 Note that a user-supplied AVFormatContext will be freed
 *                 on failure.
 * @param url      URL of the stream to open.
 * @param fmt      If non-NULL, this parameter forces a specific input format.
 *                 Otherwise the format is autodetected.
 * @param options  A dictionary filled with AVFormatContext and demuxer-private
 *                 options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note If you want to use custom IO, preallocate the format context and set its pb field.
 */
func Avformat_open_input(ps **AVFormatContext, url *byte,
                        fmt *AVInputFormat, options **AVDictionary) int32 {
    return int32(C.avformat_open_input(
        (**C.struct_AVFormatContext)(unsafe.Pointer(ps)), 
        (*C.char)(unsafe.Pointer(url)), 
        (*C.struct_AVInputFormat)(unsafe.Pointer(fmt)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Read packets of a media file to get stream information. This
 * is useful for file formats with no headers such as MPEG. This
 * function also computes the real framerate in case of MPEG-2 repeat
 * frame mode.
 * The logical file position is not changed by this function;
 * examined packets may be buffered for later processing.
 *
 * @param ic media file handle
 * @param options  If non-NULL, an ic.nb_streams long array of pointers to
 *                 dictionaries, where i-th member contains options for
 *                 codec corresponding to i-th stream.
 *                 On return each dictionary will be filled with options that were not found.
 * @return >=0 if OK, AVERROR_xxx on error
 *
 * @note this function isn't guaranteed to open all the codecs, so
 *       options being non-empty at return is a perfectly normal behavior.
 *
 * @todo Let the user decide somehow what information is needed so that
 *       we do not waste time getting stuff the user does not need.
 */
func Avformat_find_stream_info(ic *AVFormatContext, options **AVDictionary) int32 {
    return int32(C.avformat_find_stream_info(
        (*C.struct_AVFormatContext)(unsafe.Pointer(ic)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Find the programs which belong to a given stream.
 *
 * @param ic    media file handle
 * @param last  the last found program, the search will start after this
 *              program, or from the beginning if it is NULL
 * @param s     stream index
 *
 * @return the next program which belongs to s, NULL if no program is found or
 *         the last program is not among the programs of ic.
 */
func Av_find_program_from_stream(ic *AVFormatContext, last *AVProgram, s int32) *AVProgram {
    return (*AVProgram)(unsafe.Pointer(C.av_find_program_from_stream(
        (*C.struct_AVFormatContext)(unsafe.Pointer(ic)), 
        (*C.struct_AVProgram)(unsafe.Pointer(last)), C.int(s))))
}

func Av_program_add_stream_index(ac *AVFormatContext, progid int32, idx uint32)  {
    C.av_program_add_stream_index((*C.struct_AVFormatContext)(unsafe.Pointer(ac)), 
        C.int(progid), C.uint(idx))
}

/**
 * Find the "best" stream in the file.
 * The best stream is determined according to various heuristics as the most
 * likely to be what the user expects.
 * If the decoder parameter is non-NULL, av_find_best_stream will find the
 * default decoder for the stream's codec; streams for which no decoder can
 * be found are ignored.
 *
 * @param ic                media file handle
 * @param type              stream type: video, audio, subtitles, etc.
 * @param wanted_stream_nb  user-requested stream number,
 *                          or -1 for automatic selection
 * @param related_stream    try to find a stream related (eg. in the same
 *                          program) to this one, or -1 if none
 * @param decoder_ret       if non-NULL, returns the decoder for the
 *                          selected stream
 * @param flags             flags; none are currently defined
 *
 * @return  the non-negative stream number in case of success,
 *          AVERROR_STREAM_NOT_FOUND if no stream with the requested type
 *          could be found,
 *          AVERROR_DECODER_NOT_FOUND if streams were found but no decoder
 *
 * @note  If av_find_best_stream returns successfully and decoder_ret is not
 *        NULL, then *decoder_ret is guaranteed to be set to a valid AVCodec.
 */
func Av_find_best_stream(ic *AVFormatContext,
                        typex AVMediaType,
                        wanted_stream_nb int32,
                        related_stream int32,
                        decoder_ret **AVCodec,
                        flags int32) int32 {
    return int32(C.av_find_best_stream(
        (*C.struct_AVFormatContext)(unsafe.Pointer(ic)), 
        C.enum_AVMediaType(typex), C.int(wanted_stream_nb), 
        C.int(related_stream), (**C.struct_AVCodec)(unsafe.Pointer(decoder_ret)), 
        C.int(flags)))
}

/**
 * Return the next frame of a stream.
 * This function returns what is stored in the file, and does not validate
 * that what is there are valid frames for the decoder. It will split what is
 * stored in the file into frames and return one for each call. It will not
 * omit invalid data between valid frames so as to give the decoder the maximum
 * information possible for decoding.
 *
 * On success, the returned packet is reference-counted (pkt->buf is set) and
 * valid indefinitely. The packet must be freed with av_packet_unref() when
 * it is no longer needed. For video, the packet contains exactly one frame.
 * For audio, it contains an integer number of frames if each frame has
 * a known fixed size (e.g. PCM or ADPCM data). If the audio frames have
 * a variable size (e.g. MPEG audio), then it contains one frame.
 *
 * pkt->pts, pkt->dts and pkt->duration are always set to correct
 * values in AVStream.time_base units (and guessed if the format cannot
 * provide them). pkt->pts can be AV_NOPTS_VALUE if the video format
 * has B-frames, so it is better to rely on pkt->dts if you do not
 * decompress the payload.
 *
 * @return 0 if OK, < 0 on error or end of file. On error, pkt will be blank
 *         (as if it came from av_packet_alloc()).
 *
 * @note pkt will be initialized, so it may be uninitialized, but it must not
 *       contain data that needs to be freed.
 */
func Av_read_frame(s *AVFormatContext, pkt *AVPacket) int32 {
    return int32(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Seek to the keyframe at timestamp.
 * 'timestamp' in 'stream_index'.
 *
 * @param s            media file handle
 * @param stream_index If stream_index is (-1), a default stream is selected,
 *                     and timestamp is automatically converted from
 *                     AV_TIME_BASE units to the stream specific time_base.
 * @param timestamp    Timestamp in AVStream.time_base units or, if no stream
 *                     is specified, in AV_TIME_BASE units.
 * @param flags        flags which select direction and seeking mode
 *
 * @return >= 0 on success
 */
func Av_seek_frame(s *AVFormatContext, stream_index int32, timestamp int64,
                  flags int32) int32 {
    return int32(C.av_seek_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.int(stream_index), C.longlong(timestamp), C.int(flags)))
}

/**
 * Seek to timestamp ts.
 * Seeking will be done so that the point from which all active streams
 * can be presented successfully will be closest to ts and within min/max_ts.
 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
 *
 * If flags contain AVSEEK_FLAG_BYTE, then all timestamps are in bytes and
 * are the file position (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_FRAME, then all timestamps are in frames
 * in the stream with stream_index (this may not be supported by all demuxers).
 * Otherwise all timestamps are in units of the stream selected by stream_index
 * or if stream_index is -1, in AV_TIME_BASE units.
 * If flags contain AVSEEK_FLAG_ANY, then non-keyframes are treated as
 * keyframes (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_BACKWARD, it is ignored.
 *
 * @param s            media file handle
 * @param stream_index index of the stream which is used as time base reference
 * @param min_ts       smallest acceptable timestamp
 * @param ts           target timestamp
 * @param max_ts       largest acceptable timestamp
 * @param flags        flags
 * @return >=0 on success, error code otherwise
 *
 * @note This is part of the new seek API which is still under construction.
 */
func Avformat_seek_file(s *AVFormatContext, stream_index int32, min_ts int64, ts int64, max_ts int64, flags int32) int32 {
    return int32(C.avformat_seek_file((*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.int(stream_index), C.longlong(min_ts), C.longlong(ts), 
        C.longlong(max_ts), C.int(flags)))
}

/**
 * Discard all internally buffered data. This can be useful when dealing with
 * discontinuities in the byte stream. Generally works only with formats that
 * can resync. This includes headerless formats like MPEG-TS/TS but should also
 * work with NUT, Ogg and in a limited way AVI for example.
 *
 * The set of streams, the detected duration, stream parameters and codecs do
 * not change when calling this function. If you want a complete reset, it's
 * better to open a new AVFormatContext.
 *
 * This does not flush the AVIOContext (s->pb). If necessary, call
 * avio_flush(s->pb) before calling this function.
 *
 * @param s media file handle
 * @return >=0 on success, error code otherwise
 */
func Avformat_flush(s *AVFormatContext) int32 {
    return int32(C.avformat_flush((*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

/**
 * Start playing a network-based stream (e.g. RTSP stream) at the
 * current position.
 */
func Av_read_play(s *AVFormatContext) int32 {
    return int32(C.av_read_play((*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

/**
 * Pause a network-based stream (e.g. RTSP stream).
 *
 * Use av_read_play() to resume it.
 */
func Av_read_pause(s *AVFormatContext) int32 {
    return int32(C.av_read_pause((*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

/**
 * Close an opened input AVFormatContext. Free it and all its contents
 * and set *s to NULL.
 */
func Avformat_close_input(s **AVFormatContext)  {
    C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(s)))
}
/**
 * @}
 */

///< seek backward
///< seeking based on position in bytes
///< seek to any frame, even non-keyframes
///< seeking based on frame number

/**
 * @addtogroup lavf_encoding
 * @{
 */

///< stream parameters initialized in avformat_write_header
///< stream parameters initialized in avformat_init_output

/**
 * Allocate the stream private data and write the stream header to
 * an output media file.
 *
 * @param s        Media file handle, must be allocated with
 *                 avformat_alloc_context().
 *                 Its \ref AVFormatContext.oformat "oformat" field must be set
 *                 to the desired output format;
 *                 Its \ref AVFormatContext.pb "pb" field must be set to an
 *                 already opened ::AVIOContext.
 * @param options  An ::AVDictionary filled with AVFormatContext and
 *                 muxer-private options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec had not already been
 *                                       fully initialized in avformat_init_output().
 * @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec had already been fully
 *                                       initialized in avformat_init_output().
 * @retval AVERROR                       A negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
 */
func Avformat_write_header(s *AVFormatContext, options **AVDictionary) int32 {
    return int32(C.avformat_write_header(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Allocate the stream private data and initialize the codec, but do not write the header.
 * May optionally be used before avformat_write_header() to initialize stream parameters
 * before actually writing the header.
 * If using this function, do not pass the same options to avformat_write_header().
 *
 * @param s        Media file handle, must be allocated with
 *                 avformat_alloc_context().
 *                 Its \ref AVFormatContext.oformat "oformat" field must be set
 *                 to the desired output format;
 *                 Its \ref AVFormatContext.pb "pb" field must be set to an
 *                 already opened ::AVIOContext.
 * @param options  An ::AVDictionary filled with AVFormatContext and
 *                 muxer-private options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec requires
 *                                       avformat_write_header to fully initialize.
 * @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec has been fully
 *                                       initialized.
 * @retval AVERROR                       Anegative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_write_header.
 */
func Avformat_init_output(s *AVFormatContext, options **AVDictionary) int32 {
    return int32(C.avformat_init_output(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Write a packet to an output media file.
 *
 * This function passes the packet directly to the muxer, without any buffering
 * or reordering. The caller is responsible for correctly interleaving the
 * packets if the format requires it. Callers that want libavformat to handle
 * the interleaving should call av_interleaved_write_frame() instead of this
 * function.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written. Note that unlike
 *            av_interleaved_write_frame(), this function does not take
 *            ownership of the packet passed to it (though some muxers may make
 *            an internal reference to the input packet).
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), in
 *            order to immediately flush data buffered within the muxer, for
 *            muxers that buffer up data internally before writing it to the
 *            output.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets passed to this function must be strictly
 *            increasing when compared in their respective timebases (unless the
 *            output format is flagged with the AVFMT_TS_NONSTRICT, then they
 *            merely have to be nondecreasing).  @ref AVPacket.duration
 *            "duration") should also be set if known.
 * @return < 0 on error, = 0 if OK, 1 if flushed and there is no more data to flush
 *
 * @see av_interleaved_write_frame()
 */
func Av_write_frame(s *AVFormatContext, pkt *AVPacket) int32 {
    return int32(C.av_write_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Write a packet to an output media file ensuring correct interleaving.
 *
 * This function will buffer the packets internally as needed to make sure the
 * packets in the output file are properly interleaved, usually ordered by
 * increasing dts. Callers doing their own interleaving should call
 * av_write_frame() instead of this function.
 *
 * Using this function instead of av_write_frame() can give muxers advance
 * knowledge of future packets, improving e.g. the behaviour of the mp4
 * muxer for VFR content in fragmenting mode.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written.
 *            <br>
 *            If the packet is reference-counted, this function will take
 *            ownership of this reference and unreference it later when it sees
 *            fit. If the packet is not reference-counted, libavformat will
 *            make a copy.
 *            The returned packet will be blank (as if returned from
 *            av_packet_alloc()), even on error.
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), to
 *            flush the interleaving queues.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets in one stream must be strictly
 *            increasing (unless the output format is flagged with the
 *            AVFMT_TS_NONSTRICT, then they merely have to be nondecreasing).
 *            @ref AVPacket.duration "duration" should also be set if known.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_write_frame(), AVFormatContext.max_interleave_delta
 */
func Av_interleaved_write_frame(s *AVFormatContext, pkt *AVPacket) int32 {
    return int32(C.av_interleaved_write_frame(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Write an uncoded frame to an output media file.
 *
 * The frame must be correctly interleaved according to the container
 * specification; if not, av_interleaved_write_uncoded_frame() must be used.
 *
 * See av_interleaved_write_uncoded_frame() for details.
 */
func Av_write_uncoded_frame(s *AVFormatContext, stream_index int32,
                           frame *AVFrame) int32 {
    return int32(C.av_write_uncoded_frame(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(stream_index), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Write an uncoded frame to an output media file.
 *
 * If the muxer supports it, this function makes it possible to write an AVFrame
 * structure directly, without encoding it into a packet.
 * It is mostly useful for devices and similar special muxers that use raw
 * video or PCM data and will not serialize it into a byte stream.
 *
 * To test whether it is possible to use it with a given muxer and stream,
 * use av_write_uncoded_frame_query().
 *
 * The caller gives up ownership of the frame and must not access it
 * afterwards.
 *
 * @return  >=0 for success, a negative code on error
 */
func Av_interleaved_write_uncoded_frame(s *AVFormatContext, stream_index int32,
                                       frame *AVFrame) int32 {
    return int32(C.av_interleaved_write_uncoded_frame(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(stream_index), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Test whether a muxer supports uncoded frame.
 *
 * @return  >=0 if an uncoded frame can be written to that muxer and stream,
 *          <0 if not
 */
func Av_write_uncoded_frame_query(s *AVFormatContext, stream_index int32) int32 {
    return int32(C.av_write_uncoded_frame_query(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(stream_index)))
}

/**
 * Write the stream trailer to an output media file and free the
 * file private data.
 *
 * May only be called after a successful call to avformat_write_header.
 *
 * @param s media file handle
 * @return 0 if OK, AVERROR_xxx on error
 */
func Av_write_trailer(s *AVFormatContext) int32 {
    return int32(C.av_write_trailer((*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

/**
 * Return the output format in the list of registered output formats
 * which best matches the provided parameters, or return NULL if
 * there is no match.
 *
 * @param short_name if non-NULL checks if short_name matches with the
 *                   names of the registered formats
 * @param filename   if non-NULL checks if filename terminates with the
 *                   extensions of the registered formats
 * @param mime_type  if non-NULL checks if mime_type matches with the
 *                   MIME type of the registered formats
 */
func Av_guess_format(short_name *byte,
                                      filename *byte,
                                      mime_type *byte) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_guess_format(
        (*C.char)(unsafe.Pointer(short_name)), 
        (*C.char)(unsafe.Pointer(filename)), 
        (*C.char)(unsafe.Pointer(mime_type)))))
}

/**
 * Guess the codec ID based upon muxer and filename.
 */
func Av_guess_codec(fmt *AVOutputFormat, short_name *byte,
                              filename *byte, mime_type *byte,
                              typex AVMediaType) AVCodecID {
    return AVCodecID(C.av_guess_codec(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(fmt)), 
        (*C.char)(unsafe.Pointer(short_name)), 
        (*C.char)(unsafe.Pointer(filename)), 
        (*C.char)(unsafe.Pointer(mime_type)), C.enum_AVMediaType(typex)))
}

/**
 * Get timing information for the data currently output.
 * The exact meaning of "currently output" depends on the format.
 * It is mostly relevant for devices that have an internal buffer and/or
 * work in real time.
 * @param s          media file handle
 * @param stream     stream in the media file
 * @param[out] dts   DTS of the last packet output for the stream, in stream
 *                   time_base units
 * @param[out] wall  absolute time when that packet whas output,
 *                   in microsecond
 * @retval  0               Success
 * @retval  AVERROR(ENOSYS) The format does not support it
 *
 * @note Some formats or devices may not allow to measure dts and wall
 *       atomically.
 */
func Av_get_output_timestamp(s *AVFormatContext, stream int32,
                            dts *int64, wall *int64) int32 {
    return int32(C.av_get_output_timestamp(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(stream), 
        (*C.longlong)(unsafe.Pointer(dts)), (*C.longlong)(unsafe.Pointer(wall))))
}


/**
 * @}
 */


/**
 * @defgroup lavf_misc Utility functions
 * @ingroup libavf
 * @{
 *
 * Miscellaneous utility functions related to both muxing and demuxing
 * (or neither).
 */

/**
 * Send a nice hexadecimal dump of a buffer to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump_log, av_pkt_dump2, av_pkt_dump_log2
 */
func Av_hex_dump(f *C.FILE, buf *uint8, size int32)  {
    C.av_hex_dump(f, (*C.uchar)(unsafe.Pointer(buf)), C.int(size))
}

/**
 * Send a nice hexadecimal dump of a buffer to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump, av_pkt_dump2, av_pkt_dump_log2
 */
func Av_hex_dump_log(avcl unsafe.Pointer, level int32, buf *uint8, size int32)  {
    C.av_hex_dump_log(avcl, C.int(level), (*C.uchar)(unsafe.Pointer(buf)), 
        C.int(size))
}

/**
 * Send a nice dump of a packet to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
func Av_pkt_dump2(f *C.FILE, pkt *AVPacket, dump_payload int32, st *AVStream)  {
    C.av_pkt_dump2(f, (*C.struct_AVPacket)(unsafe.Pointer(pkt)), C.int(dump_payload), 
        (*C.struct_AVStream)(unsafe.Pointer(st)))
}


/**
 * Send a nice dump of a packet to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
func Av_pkt_dump_log2(avcl unsafe.Pointer, level int32, pkt *AVPacket, dump_payload int32,
                      st *AVStream)  {
    C.av_pkt_dump_log2(avcl, C.int(level), (*C.struct_AVPacket)(unsafe.Pointer(pkt)), 
        C.int(dump_payload), (*C.struct_AVStream)(unsafe.Pointer(st)))
}

/**
 * Get the AVCodecID for the given codec tag tag.
 * If no codec id is found returns AV_CODEC_ID_NONE.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param tag  codec tag to match to a codec ID
 */
func Av_codec_get_id(tags **AVCodecTag, tag uint32) AVCodecID {
    return AVCodecID(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), 
        C.uint(tag)))
}

/**
 * Get the codec tag for the given codec id id.
 * If no codec tag is found returns 0.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id   codec ID to match to a codec tag
 */
func Av_codec_get_tag(tags **AVCodecTag, id AVCodecID) uint32 {
    return uint32(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), 
        C.enum_AVCodecID(id)))
}

/**
 * Get the codec tag for the given codec id.
 *
 * @param tags list of supported codec_id - codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id codec id that should be searched for in the list
 * @param tag A pointer to the found tag
 * @return 0 if id was not found in tags, > 0 if it was found
 */
func Av_codec_get_tag2(tags **AVCodecTag, id AVCodecID,
                      tag *uint32) int32 {
    return int32(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), 
        C.enum_AVCodecID(id), (*C.uint)(unsafe.Pointer(tag))))
}

func Av_find_default_stream_index(s *AVFormatContext) int32 {
    return int32(C.av_find_default_stream_index(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

/**
 * Get the index for a specific timestamp.
 *
 * @param st        stream that the timestamp belongs to
 * @param timestamp timestamp to retrieve the index for
 * @param flags if AVSEEK_FLAG_BACKWARD then the returned index will correspond
 *                 to the timestamp which is <= the requested one, if backward
 *                 is 0, then it will be >=
 *              if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise
 * @return < 0 if no such timestamp could be found
 */
func Av_index_search_timestamp(st *AVStream, timestamp int64, flags int32) int32 {
    return int32(C.av_index_search_timestamp(
        (*C.struct_AVStream)(unsafe.Pointer(st)), C.longlong(timestamp), 
        C.int(flags)))
}

/**
 * Get the index entry count for the given AVStream.
 *
 * @param st stream
 * @return the number of index entries in the stream
 */
func Avformat_index_get_entries_count(st *AVStream) int32 {
    return int32(C.avformat_index_get_entries_count(
        (*C.struct_AVStream)(unsafe.Pointer(st))))
}

/**
 * Get the AVIndexEntry corresponding to the given index.
 *
 * @param st          Stream containing the requested AVIndexEntry.
 * @param idx         The desired index.
 * @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.
 *
 * @note The pointer returned by this function is only guaranteed to be valid
 *       until any function that takes the stream or the parent AVFormatContext
 *       as input argument is called.
 */
func Avformat_index_get_entry(st *AVStream, idx int32) *AVIndexEntry {
    return (*AVIndexEntry)(unsafe.Pointer(C.avformat_index_get_entry(
        (*C.struct_AVStream)(unsafe.Pointer(st)), C.int(idx))))
}

/**
 * Get the AVIndexEntry corresponding to the given timestamp.
 *
 * @param st          Stream containing the requested AVIndexEntry.
 * @param wanted_timestamp   Timestamp to retrieve the index entry for.
 * @param flags       If AVSEEK_FLAG_BACKWARD then the returned entry will correspond
 *                    to the timestamp which is <= the requested one, if backward
 *                    is 0, then it will be >=
 *                    if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise.
 * @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.
 *
 * @note The pointer returned by this function is only guaranteed to be valid
 *       until any function that takes the stream or the parent AVFormatContext
 *       as input argument is called.
 */
func Avformat_index_get_entry_from_timestamp(st *AVStream,
                                                            wanted_timestamp int64,
                                                            flags int32) *AVIndexEntry {
    return (*AVIndexEntry)(unsafe.Pointer(C.avformat_index_get_entry_from_timestamp(
        (*C.struct_AVStream)(unsafe.Pointer(st)), C.longlong(wanted_timestamp), 
        C.int(flags))))
}
/**
 * Add an index entry into a sorted list. Update the entry if the list
 * already contains it.
 *
 * @param timestamp timestamp in the time base of the given stream
 */
func Av_add_index_entry(st *AVStream, pos int64, timestamp int64,
                       size int32, distance int32, flags int32) int32 {
    return int32(C.av_add_index_entry((*C.struct_AVStream)(unsafe.Pointer(st)), 
        C.longlong(pos), C.longlong(timestamp), C.int(size), C.int(distance), 
        C.int(flags)))
}


/**
 * Split a URL string into components.
 *
 * The pointers to buffers for storing individual components may be null,
 * in order to ignore that component. Buffers for components not found are
 * set to empty strings. If the port is not found, it is set to a negative
 * value.
 *
 * @param proto the buffer for the protocol
 * @param proto_size the size of the proto buffer
 * @param authorization the buffer for the authorization
 * @param authorization_size the size of the authorization buffer
 * @param hostname the buffer for the host name
 * @param hostname_size the size of the hostname buffer
 * @param port_ptr a pointer to store the port number in
 * @param path the buffer for the path
 * @param path_size the size of the path buffer
 * @param url the URL to split
 */
func Av_url_split(proto *byte,         proto_size int32,
                  authorization *byte, authorization_size int32,
                  hostname *byte,      hostname_size int32,
                  port_ptr *int32,
                  path *byte,          path_size int32,
                  url *byte)  {
    C.av_url_split((*C.char)(unsafe.Pointer(proto)), C.int(proto_size), 
        (*C.char)(unsafe.Pointer(authorization)), C.int(authorization_size), 
        (*C.char)(unsafe.Pointer(hostname)), C.int(hostname_size), 
        (*C.int)(unsafe.Pointer(port_ptr)), (*C.char)(unsafe.Pointer(path)), 
        C.int(path_size), (*C.char)(unsafe.Pointer(url)))
}


/**
 * Print detailed information about the input or output format, such as
 * duration, bitrate, streams, container, programs, metadata, side data,
 * codec and time base.
 *
 * @param ic        the context to analyze
 * @param index     index of the stream to dump information about
 * @param url       the URL to print, such as source or destination file
 * @param is_output Select whether the specified context is an input(0) or output(1)
 */
func Av_dump_format(ic *AVFormatContext,
                    index int32,
                    url *byte,
                    is_output int32)  {
    C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(ic)), C.int(index), 
        (*C.char)(unsafe.Pointer(url)), C.int(is_output))
}


///< Allow multiple %d

/**
 * Return in 'buf' the path with '%d' replaced by a number.
 *
 * Also handles the '%0nd' format where 'n' is the total number
 * of digits and '%%'.
 *
 * @param buf destination buffer
 * @param buf_size destination buffer size
 * @param path numbered sequence string
 * @param number frame number
 * @param flags AV_FRAME_FILENAME_FLAGS_*
 * @return 0 if OK, -1 on format error
 */
func Av_get_frame_filename2(buf *byte, buf_size int32,
                          path *byte, number int32, flags int32) int32 {
    return int32(C.av_get_frame_filename2((*C.char)(unsafe.Pointer(buf)), 
        C.int(buf_size), (*C.char)(unsafe.Pointer(path)), C.int(number), 
        C.int(flags)))
}

func Av_get_frame_filename(buf *byte, buf_size int32,
                          path *byte, number int32) int32 {
    return int32(C.av_get_frame_filename((*C.char)(unsafe.Pointer(buf)), 
        C.int(buf_size), (*C.char)(unsafe.Pointer(path)), C.int(number)))
}

/**
 * Check whether filename actually is a numbered sequence generator.
 *
 * @param filename possible numbered sequence string
 * @return 1 if a valid numbered sequence string, 0 otherwise
 */
func Av_filename_number_test(filename *byte) int32 {
    return int32(C.av_filename_number_test((*C.char)(unsafe.Pointer(filename))))
}

/**
 * Generate an SDP for an RTP session.
 *
 * Note, this overwrites the id values of AVStreams in the muxer contexts
 * for getting unique dynamic payload types.
 *
 * @param ac array of AVFormatContexts describing the RTP streams. If the
 *           array is composed by only one context, such context can contain
 *           multiple AVStreams (one AVStream per RTP stream). Otherwise,
 *           all the contexts in the array (an AVCodecContext per RTP stream)
 *           must contain only one AVStream.
 * @param n_files number of AVCodecContexts contained in ac
 * @param buf buffer where the SDP will be stored (must be allocated by
 *            the caller)
 * @param size the size of the buffer
 * @return 0 if OK, AVERROR_xxx on error
 */
func Av_sdp_create(ac []*AVFormatContext, n_files int32, buf *byte, size int32) int32 {
    return int32(C.av_sdp_create(
        (**C.struct_AVFormatContext)(unsafe.Pointer(&ac[0])), C.int(n_files), 
        (*C.char)(unsafe.Pointer(buf)), C.int(size)))
}

/**
 * Return a positive value if the given filename has one of the given
 * extensions, 0 otherwise.
 *
 * @param filename   file name to check against the given extensions
 * @param extensions a comma-separated list of filename extensions
 */
func Av_match_ext(filename *byte, extensions *byte) int32 {
    return int32(C.av_match_ext((*C.char)(unsafe.Pointer(filename)), 
        (*C.char)(unsafe.Pointer(extensions))))
}

/**
 * Test if the given container can store a codec.
 *
 * @param ofmt           container to check for compatibility
 * @param codec_id       codec to potentially store in container
 * @param std_compliance standards compliance level, one of FF_COMPLIANCE_*
 *
 * @return 1 if codec with ID codec_id can be stored in ofmt, 0 if it cannot.
 *         A negative number if this information is not available.
 */
func Avformat_query_codec(ofmt *AVOutputFormat, codec_id AVCodecID,
                         std_compliance int32) int32 {
    return int32(C.avformat_query_codec(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(ofmt)), 
        C.enum_AVCodecID(codec_id), C.int(std_compliance)))
}

/**
 * @defgroup riff_fourcc RIFF FourCCs
 * @{
 * Get the tables mapping RIFF FourCCs to libavcodec AVCodecIDs. The tables are
 * meant to be passed to av_codec_get_id()/av_codec_get_tag() as in the
 * following code:
 * @code
 * uint32_t tag = MKTAG('H', '2', '6', '4');
 * const struct AVCodecTag *table[] = { avformat_get_riff_video_tags(), 0 };
 * enum AVCodecID id = av_codec_get_id(table, tag);
 * @endcode
 */
/**
 * @return the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
 */
func Avformat_get_riff_video_tags() *AVCodecTag {
    return (*AVCodecTag)(unsafe.Pointer(C.avformat_get_riff_video_tags()))
}
/**
 * @return the table mapping RIFF FourCCs for audio to AVCodecID.
 */
func Avformat_get_riff_audio_tags() *AVCodecTag {
    return (*AVCodecTag)(unsafe.Pointer(C.avformat_get_riff_audio_tags()))
}
/**
 * @return the table mapping MOV FourCCs for video to libavcodec AVCodecID.
 */
func Avformat_get_mov_video_tags() *AVCodecTag {
    return (*AVCodecTag)(unsafe.Pointer(C.avformat_get_mov_video_tags()))
}
/**
 * @return the table mapping MOV FourCCs for audio to AVCodecID.
 */
func Avformat_get_mov_audio_tags() *AVCodecTag {
    return (*AVCodecTag)(unsafe.Pointer(C.avformat_get_mov_audio_tags()))
}

/**
 * @}
 */

/**
 * Guess the sample aspect ratio of a frame, based on both the stream and the
 * frame aspect ratio.
 *
 * Since the frame aspect ratio is set by the codec but the stream aspect ratio
 * is set by the demuxer, these two may not be equal. This function tries to
 * return the value that you should use if you would like to display the frame.
 *
 * Basic logic is to use the stream aspect ratio if it is set to something sane
 * otherwise use the frame aspect ratio. This way a container setting, which is
 * usually easy to modify can override the coded value in the frames.
 *
 * @param format the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame with the aspect ratio to be determined
 * @return the guessed (valid) sample_aspect_ratio, 0/1 if no idea
 */
func Av_guess_sample_aspect_ratio(format *AVFormatContext, stream *AVStream,
                                        frame *AVFrame) AVRational {
    _ret := C.av_guess_sample_aspect_ratio(
        (*C.struct_AVFormatContext)(unsafe.Pointer(format)), 
        (*C.struct_AVStream)(unsafe.Pointer(stream)), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}

/**
 * Guess the frame rate, based on both the container and codec information.
 *
 * @param ctx the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame for which the frame rate should be determined, may be NULL
 * @return the guessed (valid) frame rate, 0/1 if no idea
 */
func Av_guess_frame_rate(ctx *AVFormatContext, stream *AVStream,
                               frame *AVFrame) AVRational {
    _ret := C.av_guess_frame_rate((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), 
        (*C.struct_AVStream)(unsafe.Pointer(stream)), 
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}

/**
 * Check if the stream st contained in s is matched by the stream specifier
 * spec.
 *
 * See the "stream specifiers" chapter in the documentation for the syntax
 * of spec.
 *
 * @return  >0 if st is matched by spec;
 *          0  if st is not matched by spec;
 *          AVERROR code if spec is invalid
 *
 * @note  A stream specifier can match several streams in the format.
 */
func Avformat_match_stream_specifier(s *AVFormatContext, st *AVStream,
                                    spec *byte) int32 {
    return int32(C.avformat_match_stream_specifier(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (*C.struct_AVStream)(unsafe.Pointer(st)), 
        (*C.char)(unsafe.Pointer(spec))))
}

func Avformat_queue_attached_pictures(s *AVFormatContext) int32 {
    return int32(C.avformat_queue_attached_pictures(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s))))
}

type AVTimebaseSource int32
const (
    AVFMT_TBCF_AUTO AVTimebaseSource = -1 + iota
    AVFMT_TBCF_DECODER
    AVFMT_TBCF_DEMUXER
    AVFMT_TBCF_R_FRAMERATE
)


/**
 * Transfer internal timing information from one stream to another.
 *
 * This function is useful when doing stream copy.
 *
 * @param ofmt     target output format for ost
 * @param ost      output stream which needs timings copy and adjustments
 * @param ist      reference input stream to copy timings from
 * @param copy_tb  define from where the stream codec timebase needs to be imported
 */
func Avformat_transfer_internal_stream_timing_info(ofmt *AVOutputFormat,
                                                  ost *AVStream, ist *AVStream,
                                                  copy_tb AVTimebaseSource) int32 {
    return int32(C.avformat_transfer_internal_stream_timing_info(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(ofmt)), 
        (*C.struct_AVStream)(unsafe.Pointer(ost)), 
        (*C.struct_AVStream)(unsafe.Pointer(ist)), 
        C.enum_AVTimebaseSource(copy_tb)))
}

/**
 * Get the internal codec timebase from a stream.
 *
 * @param st  input stream to extract the timebase from
 */
func Av_stream_get_codec_timebase(st *AVStream) AVRational {
    _ret := C.av_stream_get_codec_timebase((*C.struct_AVStream)(unsafe.Pointer(st)))
    return *(*AVRational)(unsafe.Pointer(&_ret))
}

/**
 * @}
 */

                                
