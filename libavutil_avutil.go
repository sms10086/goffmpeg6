// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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

//#cgo pkg-config: libavutil
//#include "libavutil/common.h"
//#include "libavutil/rational.h"
//#include "libavutil/version.h"
//#include "libavutil/macros.h"
//#include "libavutil/mathematics.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/avutil.h"
import "C"
import (
    "unsafe"
)

const FF_LAMBDA_SHIFT = 7
const FF_LAMBDA_SCALE =  (1<<FF_LAMBDA_SHIFT) 
const FF_QP2LAMBDA = 118
const FF_LAMBDA_MAX = 32767
const FF_QUALITY_SCALE =  FF_LAMBDA_SCALE  
const AV_NOPTS_VALUE =           (^int64(0x7fffffffffffffff)) 
const AV_TIME_BASE = 1000000
//const AV_TIME_BASE_Q =           (AVRational){1, AV_TIME_BASE} 
const AV_FOURCC_MAX_STRING_SIZE = 32


                       
                       

/**
 * @file
 * @ingroup lavu
 * Convenience header that includes @ref lavu "libavutil"'s core.
 */

/**
 * @mainpage
 *
 * @section ffmpeg_intro Introduction
 *
 * This document describes the usage of the different libraries
 * provided by FFmpeg.
 *
 * @li @ref libavc "libavcodec" encoding/decoding library
 * @li @ref lavfi "libavfilter" graph-based frame editing library
 * @li @ref libavf "libavformat" I/O and muxing/demuxing library
 * @li @ref lavd "libavdevice" special devices muxing/demuxing library
 * @li @ref lavu "libavutil" common utility library
 * @li @ref lswr "libswresample" audio resampling, format conversion and mixing
 * @li @ref lpp  "libpostproc" post processing library
 * @li @ref libsws "libswscale" color conversion and scaling library
 *
 * @section ffmpeg_versioning Versioning and compatibility
 *
 * Each of the FFmpeg libraries contains a version.h header, which defines a
 * major, minor and micro version number with the
 * <em>LIBRARYNAME_VERSION_{MAJOR,MINOR,MICRO}</em> macros. The major version
 * number is incremented with backward incompatible changes - e.g. removing
 * parts of the public API, reordering public struct members, etc. The minor
 * version number is incremented for backward compatible API changes or major
 * new features - e.g. adding a new public function or a new decoder. The micro
 * version number is incremented for smaller changes that a calling program
 * might still want to check for - e.g. changing behavior in a previously
 * unspecified situation.
 *
 * FFmpeg guarantees backward API and ABI compatibility for each library as long
 * as its major version number is unchanged. This means that no public symbols
 * will be removed or renamed. Types and names of the public struct members and
 * values of public macros and enums will remain the same (unless they were
 * explicitly declared as not part of the public API). Documented behavior will
 * not change.
 *
 * In other words, any correct program that works with a given FFmpeg snapshot
 * should work just as well without any changes with any later snapshot with the
 * same major versions. This applies to both rebuilding the program against new
 * FFmpeg versions or to replacing the dynamic FFmpeg libraries that a program
 * links against.
 *
 * However, new public symbols may be added and new members may be appended to
 * public structs whose size is not part of public ABI (most public structs in
 * FFmpeg). New macros and enum values may be added. Behavior in undocumented
 * situations may change slightly (and be documented). All those are accompanied
 * by an entry in doc/APIchanges and incrementing either the minor or micro
 * version number.
 */

/**
 * @defgroup lavu libavutil
 * Common code shared across all FFmpeg libraries.
 *
 * @note
 * libavutil is designed to be modular. In most cases, in order to use the
 * functions provided by one component of libavutil you must explicitly include
 * the specific header containing that feature. If you are only using
 * media-related components, you could simply include libavutil/avutil.h, which
 * brings in most of the "core" components.
 *
 * @{
 *
 * @defgroup lavu_crypto Crypto and Hashing
 *
 * @{
 * @}
 *
 * @defgroup lavu_math Mathematics
 * @{
 *
 * @}
 *
 * @defgroup lavu_string String Manipulation
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_mem Memory Management
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_data Data Structures
 * @{
 *
 * @}
 *
 * @defgroup lavu_video Video related
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_audio Audio related
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_error Error Codes
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_log Logging Facility
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_misc Other
 *
 * @{
 *
 * @defgroup preproc_misc Preprocessor String Macros
 *
 * @{
 *
 * @}
 *
 * @defgroup version_utils Library Version Macros
 *
 * @{
 *
 * @}
 */


/**
 * @addtogroup lavu_ver
 * @{
 */

/**
 * Return the LIBAVUTIL_VERSION_INT constant.
 */
func Avutil_version() uint32 {
    return uint32(C.avutil_version())
}

/**
 * Return an informative version string. This usually is the actual release
 * version number or a git commit description. This string has no fixed format
 * and can change any time. It should never be parsed by code.
 */
func Av_version_info() string {
    return C.GoString(C.av_version_info())
}

/**
 * Return the libavutil build-time configuration.
 */
func Avutil_configuration() string {
    return C.GoString(C.avutil_configuration())
}

/**
 * Return the libavutil license.
 */
func Avutil_license() string {
    return C.GoString(C.avutil_license())
}

/**
 * @}
 */

/**
 * @addtogroup lavu_media Media Type
 * @brief Media Type
 */

type AVMediaType int32
const (
    AVMEDIA_TYPE_UNKNOWN AVMediaType = -1 + iota
    AVMEDIA_TYPE_VIDEO
    AVMEDIA_TYPE_AUDIO
    AVMEDIA_TYPE_DATA
    AVMEDIA_TYPE_SUBTITLE
    AVMEDIA_TYPE_ATTACHMENT
    AVMEDIA_TYPE_NB
)


/**
 * Return a string describing the media_type enum, NULL if media_type
 * is unknown.
 */
func Av_get_media_type_string(media_type AVMediaType) string {
    return C.GoString(C.av_get_media_type_string(C.enum_AVMediaType(media_type)))
}

/**
 * @defgroup lavu_const Constants
 * @{
 *
 * @defgroup lavu_enc Encoding specific
 *
 * @note those definition should move to avcodec
 * @{
 */



///< factor to convert from H.263 QP to lambda


//FIXME maybe remove

/**
 * @}
 * @defgroup lavu_time Timestamp specific
 *
 * FFmpeg internal timebase and timestamp definitions
 *
 * @{
 */

/**
 * @brief Undefined timestamp value
 *
 * Usually reported by demuxer that work on containers that do not provide
 * either pts or dts.
 */



/**
 * Internal time base represented as integer
 */



/**
 * Internal time base represented as fractional value
 */

                  
                                        
                                                          
     

      

/**
 * @}
 * @}
 * @defgroup lavu_picture Image related
 *
 * AVPicture types, pixel formats and basic image planes manipulation.
 *
 * @{
 */

type AVPictureType int32
const (
    AV_PICTURE_TYPE_NONE AVPictureType = 0 + iota
    AV_PICTURE_TYPE_I
    AV_PICTURE_TYPE_P
    AV_PICTURE_TYPE_B
    AV_PICTURE_TYPE_S
    AV_PICTURE_TYPE_SI
    AV_PICTURE_TYPE_SP
    AV_PICTURE_TYPE_BI
)


/**
 * Return a single letter to describe the given picture type
 * pict_type.
 *
 * @param[in] pict_type the picture type @return a single character
 * representing the picture type, '?' if pict_type is unknown
 */
func Av_get_picture_type_char(pict_type AVPictureType) byte {
    return byte(C.av_get_picture_type_char(C.enum_AVPictureType(pict_type)))
}

/**
 * @}
 */

                   
                     
                    
                   
                        
                
                   

/**
 * Return x default pointer in case p is NULL.
 */
// *av_x_if_null(constvoid*p,constvoid*x)

/**
 * Compute the length of an integer list.
 *
 * @param elsize  size in bytes of each list element (only 1, 2, 4 or 8)
 * @param term    list terminator (usually 0 or -1)
 * @param list    pointer to the list
 * @return  length of the list, in elements, not counting the terminator
 */
func Av_int_list_length_for_size(elsize uint32,
                                     list unsafe.Pointer, term uint64)  uint32 {
    return uint32(C.av_int_list_length_for_size(C.unsigned(elsize), list, 
        C.ulonglong(term)))
}

/**
 * Compute the length of an integer list.
 *
 * @param term  list terminator (usually 0 or -1)
 * @param list  pointer to the list
 * @return  length of the list, in elements, not counting the terminator
 */


                        
/**
 * Open a file using a UTF-8 filename.
 * The API of this function matches POSIX fopen(), errors are returned through
 * errno.
 * @deprecated Avoid using it, as on Windows, the FILE* allocated by this
 *             function may be allocated with a different CRT than the caller
 *             who uses the FILE*. No replacement provided in public API.
 */

func Av_fopen_utf8(path *byte, mode *byte) *C.FILE {
    return (*C.FILE)(unsafe.Pointer(C.av_fopen_utf8((*C.char)(unsafe.Pointer(path)), 
        (*C.char)(unsafe.Pointer(mode)))))
}
      

/**
 * Return the fractional representation of the internal time base.
 */
func Av_get_time_base_q() AVRational {
    _ret := C.av_get_time_base_q()
    return *(*AVRational)(unsafe.Pointer(&_ret))
}





/**
 * Fill the provided buffer with a string containing a FourCC (four-character
 * code) representation.
 *
 * @param buf    a buffer with size in bytes of at least AV_FOURCC_MAX_STRING_SIZE
 * @param fourcc the fourcc to represent
 * @return the buffer in input
 */
func Av_fourcc_make_string(buf *byte, fourcc uint32) string {
    return C.GoString(C.av_fourcc_make_string((*C.char)(unsafe.Pointer(buf)), 
        C.uint(fourcc)))
}

/**
 * @}
 * @}
 */

                            
