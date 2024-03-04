// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2012 Nicolas George
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
//#include "libavutil/avstring.h"
//#include "libavutil/bprint.h"
import "C"
import (
    "unsafe"
)

const AV_BPRINT_SIZE_UNLIMITED = (^uint32(0))
const AV_BPRINT_SIZE_AUTOMATIC = 1
const AV_BPRINT_SIZE_COUNT_ONLY = 0


/**
 * @file
 * @ingroup lavu_avbprint
 * AVBPrint public header
 */

                       
                       

                   

                       
                     

/**
 * @defgroup lavu_avbprint AVBPrint
 * @ingroup lavu_data
 *
 * A buffer to print data progressively
 * @{
 */

/**
 * Define a structure with extra padding to a fixed size
 * This helps ensuring binary compatibility with future versions.
 */



/**
 * Buffer to print data progressively
 *
 * The string buffer grows as necessary and is always 0-terminated.
 * The content of the string is never accessed, and thus is
 * encoding-agnostic and can even hold binary data.
 *
 * Small buffers are kept in the structure itself, and thus require no
 * memory allocation at all (unless the contents of the buffer is needed
 * after the structure goes out of scope). This is almost as lightweight as
 * declaring a local `char buf[512]`.
 *
 * The length of the string can go beyond the allocated size: the buffer is
 * then truncated, but the functions still keep account of the actual total
 * length.
 *
 * In other words, AVBPrint.len can be greater than AVBPrint.size and records
 * the total length of what would have been to the buffer if there had been
 * enough memory.
 *
 * Append operations do not need to be tested for failure: if a memory
 * allocation fails, data stop being appended to the buffer, but the length
 * is still updated. This situation can be tested with
 * av_bprint_is_complete().
 *
 * The AVBPrint.size_max field determines several possible behaviours:
 * - `size_max = -1` (= `UINT_MAX`) or any large value will let the buffer be
 *   reallocated as necessary, with an amortized linear cost.
 * - `size_max = 0` prevents writing anything to the buffer: only the total
 *   length is computed. The write operations can then possibly be repeated in
 *   a buffer with exactly the necessary size
 *   (using `size_init = size_max = len + 1`).
 * - `size_max = 1` is automatically replaced by the exact size available in the
 *   structure itself, thus ensuring no dynamic memory allocation. The
 *   internal buffer is large enough to hold a reasonable paragraph of text,
 *   such as the current paragraph.
 */

//FF_PAD_STRUCTURE(AVBPrint, 1024,


/**
 * @name Max size special values
 * Convenience macros for special values for av_bprint_init() size_max
 * parameter.
 * @{
 */

/**
 * Buffer will be reallocated as necessary, with an amortized linear cost.
 */

/**
 * Use the exact size available in the AVBPrint structure itself.
 *
 * Thus ensuring no dynamic memory allocation. The internal buffer is large
 * enough to hold a reasonable paragraph of text, such as the current paragraph.
 */

/**
 * Do not write anything to the buffer, only calculate the total length.
 *
 * The write operations can then possibly be repeated in a buffer with
 * exactly the necessary size (using `size_init = size_max = AVBPrint.len + 1`).
 */

/** @} */

/**
 * Init a print buffer.
 *
 * @param buf        buffer to init
 * @param size_init  initial size (including the final 0)
 * @param size_max   maximum size;
 *                   - `0` means do not write anything, just count the length
 *                   - `1` is replaced by the maximum value for automatic storage
 *                       any large value means that the internal buffer will be
 *                       reallocated as needed up to that limit
 *                   - `-1` is converted to `UINT_MAX`, the largest limit possible.
 *                   Check also `AV_BPRINT_SIZE_*` macros.
 */
func Av_bprint_init(buf *AVBPrint, size_init uint32, size_max uint32)  {
    C.av_bprint_init((*C.AVBPrint)(unsafe.Pointer(buf)), C.unsigned(size_init), 
        C.unsigned(size_max))
}

/**
 * Init a print buffer using a pre-existing buffer.
 *
 * The buffer will not be reallocated.
 * In case size equals zero, the AVBPrint will be initialized to use
 * the internal buffer as if using AV_BPRINT_SIZE_COUNT_ONLY with
 * av_bprint_init().
 *
 * @param buf     buffer structure to init
 * @param buffer  byte buffer to use for the string data
 * @param size    size of buffer
 */
func Av_bprint_init_for_buffer(buf *AVBPrint, buffer *byte, size uint32)  {
    C.av_bprint_init_for_buffer((*C.AVBPrint)(unsafe.Pointer(buf)), 
        (*C.char)(unsafe.Pointer(buffer)), C.unsigned(size))
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_bprintf(AVBPrint *buf, const char *fmt, ...) av_printf_format(2, 3);

/**
 * Append a formatted string to a print buffer.
 */
// not supported: av_vbprintf

/**
 * Append char c n times to a print buffer.
 */
func Av_bprint_chars(buf *AVBPrint, c byte, n uint32)  {
    C.av_bprint_chars((*C.AVBPrint)(unsafe.Pointer(buf)), C.char(c), 
        C.unsigned(n))
}

/**
 * Append data to a print buffer.
 *
 * param buf  bprint buffer to use
 * param data pointer to data
 * param size size of data
 */
func Av_bprint_append_data(buf *AVBPrint, data *byte, size uint32)  {
    C.av_bprint_append_data((*C.AVBPrint)(unsafe.Pointer(buf)), 
        (*C.char)(unsafe.Pointer(data)), C.unsigned(size))
}

type tm C.struct_tm
/**
 * Append a formatted date and time to a print buffer.
 *
 * param buf  bprint buffer to use
 * param fmt  date and time format string, see strftime()
 * param tm   broken-down time structure to translate
 *
 * @note due to poor design of the standard strftime function, it may
 * produce poor results if the format string expands to a very long text and
 * the bprint buffer is near the limit stated by the size_max option.
 */
func Av_bprint_strftime(buf *AVBPrint, fmt *byte, tm *tm)  {
    C.av_bprint_strftime((*C.AVBPrint)(unsafe.Pointer(buf)), 
        (*C.char)(unsafe.Pointer(fmt)), (*C.struct_tm)(unsafe.Pointer(tm)))
}

/**
 * Allocate bytes in the buffer for external use.
 *
 * @param[in]  buf          buffer structure
 * @param[in]  size         required size
 * @param[out] mem          pointer to the memory area
 * @param[out] actual_size  size of the memory area after allocation;
 *                          can be larger or smaller than size
 */
func Av_bprint_get_buffer(buf *AVBPrint, size uint32,
                          mem **uint8, actual_size *uint32)  {
    C.av_bprint_get_buffer((*C.AVBPrint)(unsafe.Pointer(buf)), C.unsigned(size), 
        (**C.uchar)(unsafe.Pointer(mem)), 
        (*C.unsigned)(unsafe.Pointer(actual_size)))
}

/**
 * Reset the string to "" but keep internal allocated data.
 */
func Av_bprint_clear(buf *AVBPrint)  {
    C.av_bprint_clear((*C.AVBPrint)(unsafe.Pointer(buf)))
}

/**
 * Test if the print buffer is complete (not truncated).
 *
 * It may have been truncated due to a memory allocation failure
 * or the size_max limit (compare size and size_max if necessary).
 */
// av_bprint_is_complete(constAVBPrint*buf)

/**
 * Finalize a print buffer.
 *
 * The print buffer can no longer be used afterwards,
 * but the len and size fields are still valid.
 *
 * @arg[out] ret_str  if not NULL, used to return a permanent copy of the
 *                    buffer contents, or NULL if memory allocation fails;
 *                    if NULL, the buffer is discarded and freed
 * @return  0 for success or error code (probably AVERROR(ENOMEM))
 */
func Av_bprint_finalize(buf *AVBPrint, ret_str **byte) int32 {
    return int32(C.av_bprint_finalize((*C.AVBPrint)(unsafe.Pointer(buf)), 
        (**C.char)(unsafe.Pointer(ret_str))))
}

/**
 * Escape the content in src and append it to dstbuf.
 *
 * @param dstbuf        already inited destination bprint buffer
 * @param src           string containing the text to escape
 * @param special_chars string containing the special characters which
 *                      need to be escaped, can be NULL
 * @param mode          escape mode to employ, see AV_ESCAPE_MODE_* macros.
 *                      Any unknown value for mode will be considered equivalent to
 *                      AV_ESCAPE_MODE_BACKSLASH, but this behaviour can change without
 *                      notice.
 * @param flags         flags which control how to escape, see AV_ESCAPE_FLAG_* macros
 */
func Av_bprint_escape(dstbuf *AVBPrint, src *byte, special_chars *byte,
                      mode AVEscapeMode, flags int32)  {
    C.av_bprint_escape((*C.AVBPrint)(unsafe.Pointer(dstbuf)), 
        (*C.char)(unsafe.Pointer(src)), (*C.char)(unsafe.Pointer(special_chars)), 
        C.enum_AVEscapeMode(mode), C.int(flags))
}

/** @} */

                            
