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
//#include <stdarg.h>
//#include "libavutil/attributes.h"
//#include "libavutil/version.h"
//#include "libavutil/log.h"
import "C"
import (
    "unsafe"
)

const AV_LOG_QUIET = -8
const AV_LOG_PANIC = 0
const AV_LOG_FATAL = 8
const AV_LOG_ERROR = 16
const AV_LOG_WARNING = 24
const AV_LOG_INFO = 32
const AV_LOG_VERBOSE = 40
const AV_LOG_DEBUG = 48
const AV_LOG_TRACE = 56
const AV_LOG_MAX_OFFSET = 64
const AV_LOG_SKIP_REPEATED = 1
const AV_LOG_PRINT_LEVEL = 2


                    
                    

                   
                       
                    

type AVClassCategory int32
const (
    AV_CLASS_CATEGORY_NA AVClassCategory = 0 + iota
    AV_CLASS_CATEGORY_INPUT
    AV_CLASS_CATEGORY_OUTPUT
    AV_CLASS_CATEGORY_MUXER
    AV_CLASS_CATEGORY_DEMUXER
    AV_CLASS_CATEGORY_ENCODER
    AV_CLASS_CATEGORY_DECODER
    AV_CLASS_CATEGORY_FILTER
    AV_CLASS_CATEGORY_BITSTREAM_FILTER
    AV_CLASS_CATEGORY_SWSCALER
    AV_CLASS_CATEGORY_SWRESAMPLER
    AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT = 40
    AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT = 40 + iota - 11
    AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT
    AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT
    AV_CLASS_CATEGORY_DEVICE_OUTPUT
    AV_CLASS_CATEGORY_DEVICE_INPUT
    AV_CLASS_CATEGORY_NB
)






// type AVOptionRanges

/**
 * Describe the class of an AVClass context structure. That is an
 * arbitrary struct of which the first field is a pointer to an
 * AVClass struct (e.g. AVCodecContext, AVFormatContext etc.).
 */
type AVClass struct {
    Class_name *byte
    Item_name uintptr
    Option *AVOption
    Version int32
    Log_level_offset_offset int32
    Parent_log_context_offset int32
    Category AVClassCategory
    Get_category uintptr
    Query_ranges uintptr
    Child_next uintptr
    Child_class_iterate uintptr
}


/**
 * @addtogroup lavu_log
 *
 * @{
 *
 * @defgroup lavu_log_constants Logging Constants
 *
 * @{
 */

/**
 * Print no output.
 */


/**
 * Something went really wrong and we will crash now.
 */


/**
 * Something went wrong and recovery is not possible.
 * For example, no header was found for a format which depends
 * on headers or an illegal combination of parameters is used.
 */


/**
 * Something went wrong and cannot losslessly be recovered.
 * However, not all future data is affected.
 */


/**
 * Something somehow does not look correct. This may or may not
 * lead to problems. An example would be the use of '-vstrict -2'.
 */


/**
 * Standard information.
 */


/**
 * Detailed information.
 */


/**
 * Stuff which is only useful for libav* developers.
 */


/**
 * Extremely verbose debugging, useful for libav* development.
 */




/**
 * @}
 */

/**
 * Sets additional colors for extended debugging sessions.
 * @code
   av_log(ctx, AV_LOG_DEBUG|AV_LOG_C(134), "Message in purple\n");
   @endcode
 * Requires 256color terminal support. Uses outside debugging is not
 * recommended.
 */


/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 */
//void av_log(void *avcl, int level, const char *fmt, ...) av_printf_format(3, 4);

/**
 * Send the specified message to the log once with the initial_level and then with
 * the subsequent_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param initial_level importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant" for the first occurance.
 * @param subsequent_level importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant" after the first occurance.
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param state a variable to keep trak of if a message has already been printed
 *        this must be initialized to 0 before the first use. The same state
 *        must not be accessed by 2 Threads simultaneously.
 */
//void av_log_once(void* avcl, int initial_level, int subsequent_level, int *state, const char *fmt, ...) av_printf_format(5, 6);


/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
// not supported: av_vlog

/**
 * Get the current log level
 *
 * @see lavu_log_constants
 *
 * @return Current log level
 */
func Av_log_get_level() int32 {
    return int32(C.av_log_get_level())
}

/**
 * Set the log level
 *
 * @see lavu_log_constants
 *
 * @param level Logging level
 */
func Av_log_set_level(level int32)  {
    C.av_log_set_level(C.int(level))
}

/**
 * Set the logging callback
 *
 * @note The callback must be thread safe, even if the application does not use
 *       threads itself as some codecs are multithreaded.
 *
 * @see av_log_default_callback
 *
 * @param callback A logging function with a compatible signature.
 */
// not supported: av_log_set_callback

/**
 * Default logging callback
 *
 * It prints the message to stderr, optionally colorizing it.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
// not supported: av_log_default_callback

/**
 * Return the context name
 *
 * @param  ctx The AVClass context
 *
 * @return The AVClass class_name
 */
func Av_default_item_name(ctx unsafe.Pointer) string {
    return C.GoString(C.av_default_item_name(ctx))
}
func Av_default_get_category(ptr unsafe.Pointer) AVClassCategory {
    return AVClassCategory(C.av_default_get_category(ptr))
}

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line
 * @param line_size     size of the buffer
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 */
// not supported: av_log_format_line

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line;
 *                      may be NULL if line_size is 0
 * @param line_size     size of the buffer; at most line_size-1 characters will
 *                      be written to the buffer, plus one null terminator
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 * @return Returns a negative value if an error occurred, otherwise returns
 *         the number of characters that would have been written for a
 *         sufficiently large buffer, not including the terminating null
 *         character. If the return value is not less than line_size, it means
 *         that the log message was truncated to fit the buffer.
 */
// not supported: av_log_format_line2

/**
 * Skip repeated messages, this requires the user app to use av_log() instead of
 * (f)printf as the 2 would otherwise interfere and lead to
 * "Last message repeated x times" messages below (f)printf messages with some
 * bad luck.
 * Also to receive the last, "last repeated" line if any, the user app must
 * call av_log(NULL, AV_LOG_QUIET, "%s", ""); at the end
 */


/**
 * Include the log severity in messages originating from codecs.
 *
 * Results in messages such as:
 * [rawvideo @ 0xDEADBEEF] [error] encode did not produce valid pts
 */


func Av_log_set_flags(arg int32)  {
    C.av_log_set_flags(C.int(arg))
}
func Av_log_get_flags() int32 {
    return int32(C.av_log_get_flags())
}

/**
 * @}
 */

                         
