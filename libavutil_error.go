// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
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
//#include <errno.h>
//#include <stddef.h>
//#include "libavutil/macros.h"
//#include "libavutil/error.h"
import "C"
import (
    "unsafe"
)

//const AVERROR_BSF_NOT_FOUND =       FFERRTAG(0xF8,'B','S','F')  
//const AVERROR_BUG =                 FFERRTAG( 'B','U','G','!')  
//const AVERROR_BUFFER_TOO_SMALL =    FFERRTAG( 'B','U','F','S')  
//const AVERROR_DECODER_NOT_FOUND =   FFERRTAG(0xF8,'D','E','C')  
//const AVERROR_DEMUXER_NOT_FOUND =   FFERRTAG(0xF8,'D','E','M')  
//const AVERROR_ENCODER_NOT_FOUND =   FFERRTAG(0xF8,'E','N','C')  
//const AVERROR_EOF =                 FFERRTAG( 'E','O','F',' ')  
//const AVERROR_EXIT =                FFERRTAG( 'E','X','I','T')  
//const AVERROR_EXTERNAL =            FFERRTAG( 'E','X','T',' ')  
//const AVERROR_FILTER_NOT_FOUND =    FFERRTAG(0xF8,'F','I','L')  
//const AVERROR_INVALIDDATA =         FFERRTAG( 'I','N','D','A')  
//const AVERROR_MUXER_NOT_FOUND =     FFERRTAG(0xF8,'M','U','X')  
//const AVERROR_OPTION_NOT_FOUND =    FFERRTAG(0xF8,'O','P','T')  
//const AVERROR_PATCHWELCOME =        FFERRTAG( 'P','A','W','E')  
//const AVERROR_PROTOCOL_NOT_FOUND =  FFERRTAG(0xF8,'P','R','O')  
//const AVERROR_STREAM_NOT_FOUND =    FFERRTAG(0xF8,'S','T','R')  
//const AVERROR_BUG2 =                FFERRTAG( 'B','U','G',' ') 
//const AVERROR_UNKNOWN =             FFERRTAG( 'U','N','K','N')  
const AVERROR_EXPERIMENTAL =        (-0x2bb2afa8)  
const AVERROR_INPUT_CHANGED =       (-0x636e6701)  
const AVERROR_OUTPUT_CHANGED =      (-0x636e6702)  
//const AVERROR_HTTP_BAD_REQUEST =    FFERRTAG(0xF8,'4','0','0') 
//const AVERROR_HTTP_UNAUTHORIZED =   FFERRTAG(0xF8,'4','0','1') 
//const AVERROR_HTTP_FORBIDDEN =      FFERRTAG(0xF8,'4','0','3') 
//const AVERROR_HTTP_NOT_FOUND =      FFERRTAG(0xF8,'4','0','4') 
//const AVERROR_HTTP_OTHER_4XX =      FFERRTAG(0xF8,'4','X','X') 
//const AVERROR_HTTP_SERVER_ERROR =   FFERRTAG(0xF8,'5','X','X') 
const AV_ERROR_MAX_STRING_SIZE = 64


/**
 * @file
 * error code definitions
 */

                      
                      

                  
                   

                   

/**
 * @addtogroup lavu_error
 *
 * @{
 */


/* error handling */
            
                                                                                                                         
                                                                                                       
     
/* Some platforms have E* and errno already negated. */


      



///< Bitstream filter not found
///< Internal bug, also see AVERROR_BUG2
///< Buffer too small
///< Decoder not found
///< Demuxer not found
///< Encoder not found
///< End of file
///< Immediate exit was requested; the called function should not be restarted
///< Generic error in an external library
///< Filter not found
///< Invalid data found when processing input
///< Muxer not found
///< Option not found
///< Not yet implemented in FFmpeg, patches welcome
///< Protocol not found

///< Stream not found
/**
 * This is semantically identical to AVERROR_BUG
 * it has been introduced in Libav after our AVERROR_BUG and with a modified value.
 */

///< Unknown error, typically from an external library
///< Requested feature is flagged experimental. Set strict_std_compliance if you really want to use it.
///< Input changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_OUTPUT_CHANGED)
///< Output changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_INPUT_CHANGED)
/* HTTP & RTSP errors */









/**
 * Put a description of the AVERROR code errnum in errbuf.
 * In case of failure the global variable errno is set to indicate the
 * error. Even in case of failure av_strerror() will print a generic
 * error message indicating the errnum provided to errbuf.
 *
 * @param errnum      error code to describe
 * @param errbuf      buffer to which description is written
 * @param errbuf_size the size in bytes of errbuf
 * @return 0 on success, a negative value if a description for errnum
 * cannot be found
 */
func Av_strerror(errnum int32, errbuf *byte, errbuf_size uint64) int32 {
    return int32(C.av_strerror(C.int(errnum), (*C.char)(unsafe.Pointer(errbuf)), 
        C.ulonglong(errbuf_size)))
}

/**
 * Fill the provided buffer with a string containing an error string
 * corresponding to the AVERROR code errnum.
 *
 * @param errbuf         a buffer
 * @param errbuf_size    size in bytes of errbuf
 * @param errnum         error code to describe
 * @return the buffer in input, filled with the error description
 * @see av_strerror()
 */
// *av_make_error_string(char*errbuf,size_terrbuf_size,interrnum)

/**
 * Convenience macro, the return value should be used only directly in
 * function arguments but never stand-alone.
 */


/**
 * @}
 */

                           
