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

//#cgo pkg-config: libavutil libavformat
//#include <stdint.h>
//#include <stdio.h>
//#include "libavutil/attributes.h"
//#include "libavutil/dict.h"
//#include "libavutil/log.h"
//#include "libavformat/version_major.h"
//#include "libavformat/avio.h"
import "C"
import (
    "syscall"
    "unsafe"
)

const AVIO_SEEKABLE_NORMAL =  (1 << 0) 
const AVIO_SEEKABLE_TIME =    (1 << 1) 
const AVSEEK_SIZE =  0x10000 
const AVSEEK_FORCE =  0x20000 
const AVIO_FLAG_READ = 1
const AVIO_FLAG_WRITE = 2
const AVIO_FLAG_READ_WRITE =  (AVIO_FLAG_READ|AVIO_FLAG_WRITE)   
const AVIO_FLAG_NONBLOCK = 8
const AVIO_FLAG_DIRECT =  0x8000 

                       
                       

/**
 * @file
 * @ingroup lavf_io
 * Buffered I/O operations
 */

                   
                  

                                 
                           
                          

                                      

/**
 * Seeking works like for a local file.
 */


/**
 * Seeking by timestamp with avio_seek_time() is possible.
 */


/**
 * Callback for checking whether to abort blocking functions.
 * AVERROR_EXIT is returned in this case by the interrupted
 * function. During blocking operations, callback is called with
 * opaque as parameter. If the callback returns 1, the
 * blocking operation will be aborted.
 *
 * No members can be added to this struct without a major bump, if
 * new elements have been added after this struct in AVFormatContext
 * or AVIOContext.
 */
type AVIOInterruptCB struct {
    Callback uintptr
    Opaque unsafe.Pointer
}


/**
 * Directory entry types.
 */
type AVIODirEntryType int32
const (
    AVIO_ENTRY_UNKNOWN AVIODirEntryType = iota
    AVIO_ENTRY_BLOCK_DEVICE
    AVIO_ENTRY_CHARACTER_DEVICE
    AVIO_ENTRY_DIRECTORY
    AVIO_ENTRY_NAMED_PIPE
    AVIO_ENTRY_SYMBOLIC_LINK
    AVIO_ENTRY_SOCKET
    AVIO_ENTRY_FILE
    AVIO_ENTRY_SERVER
    AVIO_ENTRY_SHARE
    AVIO_ENTRY_WORKGROUP
)


/**
 * Describes single entry of the directory.
 *
 * Only name and type fields are guaranteed be set.
 * Rest of fields are protocol or/and platform dependent and might be unknown.
 */
type AVIODirEntry struct {
    Name *byte
    Type int32
    Utf8 int32
    Size int64
    Modification_timestamp int64
    Access_timestamp int64
    Status_change_timestamp int64
    User_id int64
    Group_id int64
    Filemode int64
}


                         
type AVIODirContext struct {
    Url_context *URLContext
}

     
                                             
      

/**
 * Different data types that can be returned via the AVIO
 * write_data_type callback.
 */
type AVIODataMarkerType int32
const (
    AVIO_DATA_MARKER_HEADER AVIODataMarkerType = iota
    AVIO_DATA_MARKER_SYNC_POINT
    AVIO_DATA_MARKER_BOUNDARY_POINT
    AVIO_DATA_MARKER_UNKNOWN
    AVIO_DATA_MARKER_TRAILER
    AVIO_DATA_MARKER_FLUSH_POINT
)


/**
 * Bytestream IO Context.
 * New public fields can be added with minor version bumps.
 * Removal, reordering and changes to existing public fields require
 * a major version bump.
 * sizeof(AVIOContext) must not be used outside libav*.
 *
 * @note None of the function pointers in AVIOContext should be called
 *       directly, they should only be set by the client application
 *       when implementing custom I/O. Normally these are set to the
 *       function pointers specified in avio_alloc_context()
 */
type AVIOContext struct {
    Av_class *AVClass
    Buffer *uint8
    Buffer_size int32
    Buf_ptr *uint8
    Buf_end *uint8
    Opaque unsafe.Pointer
    Read_packet uintptr
    Write_packet uintptr
    Seek uintptr
    Pos int64
    Eof_reached int32
    Error int32
    Write_flag int32
    Max_packet_size int32
    Min_packet_size int32
    Checksum uint32
    Checksum_ptr *uint8
    Update_checksum uintptr
    Read_pause uintptr
    Read_seek uintptr
    Seekable int32
    Direct int32
    Protocol_whitelist *byte
    Protocol_blacklist *byte
    Write_data_type uintptr
    Ignore_boundary_point int32
    Buf_ptr_max *uint8
    Bytes_read int64
    Bytes_written int64
}


/**
 * Return the name of the protocol that will handle the passed URL.
 *
 * NULL is returned if no protocol could be found for the given URL.
 *
 * @return Name of the protocol or NULL.
 */
func Avio_find_protocol_name(url *byte) string {
    return C.GoString(C.avio_find_protocol_name((*C.char)(unsafe.Pointer(url))))
}

/**
 * Return AVIO_FLAG_* access flags corresponding to the access permissions
 * of the resource in url, or a negative value corresponding to an
 * AVERROR code in case of failure. The returned access flags are
 * masked by the value in flags.
 *
 * @note This function is intrinsically unsafe, in the sense that the
 * checked resource may change its existence or permission status from
 * one call to another. Thus you should not trust the returned value,
 * unless you are sure that no other processes are accessing the
 * checked resource.
 */
func Avio_check(url *byte, flags int32) int32 {
    return int32(C.avio_check((*C.char)(unsafe.Pointer(url)), C.int(flags)))
}

/**
 * Open directory for reading.
 *
 * @param s       directory read context. Pointer to a NULL pointer must be passed.
 * @param url     directory to be listed.
 * @param options A dictionary filled with protocol-private options. On return
 *                this parameter will be destroyed and replaced with a dictionary
 *                containing options that were not found. May be NULL.
 * @return >=0 on success or negative on error.
 */
func Avio_open_dir(s **AVIODirContext, url *byte, options **AVDictionary) int32 {
    return int32(C.avio_open_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(url)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Get next directory entry.
 *
 * Returned entry must be freed with avio_free_directory_entry(). In particular
 * it may outlive AVIODirContext.
 *
 * @param s         directory read context.
 * @param[out] next next entry or NULL when no more entries.
 * @return >=0 on success or negative on error. End of list is not considered an
 *             error.
 */
func Avio_read_dir(s *AVIODirContext, next **AVIODirEntry) int32 {
    return int32(C.avio_read_dir((*C.struct_AVIODirContext)(unsafe.Pointer(s)), 
        (**C.struct_AVIODirEntry)(unsafe.Pointer(next))))
}

/**
 * Close directory.
 *
 * @note Entries created using avio_read_dir() are not deleted and must be
 * freeded with avio_free_directory_entry().
 *
 * @param s         directory read context.
 * @return >=0 on success or negative on error.
 */
func Avio_close_dir(s **AVIODirContext) int32 {
    return int32(C.avio_close_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s))))
}

/**
 * Free entry allocated by avio_read_dir().
 *
 * @param entry entry to be freed.
 */
func Avio_free_directory_entry(entry **AVIODirEntry)  {
    C.avio_free_directory_entry((**C.struct_AVIODirEntry)(unsafe.Pointer(entry)))
}

/**
 * Allocate and initialize an AVIOContext for buffered I/O. It must be later
 * freed with avio_context_free().
 *
 * @param buffer Memory block for input/output operations via AVIOContext.
 *        The buffer must be allocated with av_malloc() and friends.
 *        It may be freed and replaced with a new buffer by libavformat.
 *        AVIOContext.buffer holds the buffer currently in use,
 *        which must be later freed with av_free().
 * @param buffer_size The buffer size is very important for performance.
 *        For protocols with fixed blocksize it should be set to this blocksize.
 *        For others a typical size is a cache page, e.g. 4kb.
 * @param write_flag Set to 1 if the buffer should be writable, 0 otherwise.
 * @param opaque An opaque pointer to user-specific data.
 * @param read_packet  A function for refilling the buffer, may be NULL.
 *                     For stream protocols, must never return 0 but rather
 *                     a proper AVERROR code.
 * @param write_packet A function for writing the buffer contents, may be NULL.
 *        The function may not change the input buffers content.
 * @param seek A function for seeking to specified byte position, may be NULL.
 *
 * @return Allocated AVIOContext or NULL on failure.
 */
func Avio_alloc_context(
                  buffer *uint8,
                  buffer_size int32,
                  write_flag int32,
                  opaque unsafe.Pointer,
                  read_packet func(opaque unsafe.Pointer, buf *uint8, buf_size int32) int32,
                              
                  write_packet func(opaque unsafe.Pointer, buf *uint8, buf_size int32) int32,
     
                                                                                      
      
                  seek func(opaque unsafe.Pointer, offset int64, whence int32) int64) *AVIOContext {
    cb4 := syscall.NewCallbackCDecl(read_packet)
    cb5 := syscall.NewCallbackCDecl(write_packet)
    cb6 := syscall.NewCallbackCDecl(seek)
    return (*AVIOContext)(unsafe.Pointer(C.avio_alloc_context(
        (*C.uchar)(unsafe.Pointer(buffer)), C.int(buffer_size), 
        C.int(write_flag), opaque, (*[0]byte)(unsafe.Pointer(cb4)), 
        (*[0]byte)(unsafe.Pointer(cb5)), (*[0]byte)(unsafe.Pointer(cb6)))))
}

/**
 * Free the supplied IO context and everything associated with it.
 *
 * @param s Double pointer to the IO context. This function will write NULL
 * into s.
 */
func Avio_context_free(s **AVIOContext)  {
    C.avio_context_free((**C.struct_AVIOContext)(unsafe.Pointer(s)))
}

func Avio_w8(s *AVIOContext, b int32)  {
    C.avio_w8((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.int(b))
}
func Avio_write(s *AVIOContext, buf *uint8, size int32)  {
    C.avio_write((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(size))
}
func Avio_wl64(s *AVIOContext, val uint64)  {
    C.avio_wl64((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.ulonglong(val))
}
func Avio_wb64(s *AVIOContext, val uint64)  {
    C.avio_wb64((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.ulonglong(val))
}
func Avio_wl32(s *AVIOContext, val uint32)  {
    C.avio_wl32((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}
func Avio_wb32(s *AVIOContext, val uint32)  {
    C.avio_wb32((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}
func Avio_wl24(s *AVIOContext, val uint32)  {
    C.avio_wl24((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}
func Avio_wb24(s *AVIOContext, val uint32)  {
    C.avio_wb24((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}
func Avio_wl16(s *AVIOContext, val uint32)  {
    C.avio_wl16((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}
func Avio_wb16(s *AVIOContext, val uint32)  {
    C.avio_wb16((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.uint(val))
}

/**
 * Write a NULL-terminated string.
 * @return number of bytes written.
 */
func Avio_put_str(s *AVIOContext, str *byte) int32 {
    return int32(C.avio_put_str((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(str))))
}

/**
 * Convert an UTF-8 string to UTF-16LE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
func Avio_put_str16le(s *AVIOContext, str *byte) int32 {
    return int32(C.avio_put_str16le((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(str))))
}

/**
 * Convert an UTF-8 string to UTF-16BE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
func Avio_put_str16be(s *AVIOContext, str *byte) int32 {
    return int32(C.avio_put_str16be((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(str))))
}

/**
 * Mark the written bytestream as a specific type.
 *
 * Zero-length ranges are omitted from the output.
 *
 * @param s    the AVIOContext
 * @param time the stream time the current bytestream pos corresponds to
 *             (in AV_TIME_BASE units), or AV_NOPTS_VALUE if unknown or not
 *             applicable
 * @param type the kind of data written starting at the current pos
 */
func Avio_write_marker(s *AVIOContext, time int64, typex AVIODataMarkerType)  {
    C.avio_write_marker((*C.struct_AVIOContext)(unsafe.Pointer(s)), C.longlong(time), 
        C.enum_AVIODataMarkerType(typex))
}

/**
 * ORing this as the "whence" parameter to a seek function causes it to
 * return the filesize without seeking anywhere. Supporting this is optional.
 * If it is not supported then the seek function will return <0.
 */


/**
 * Passing this flag as the "whence" parameter to a seek function causes it to
 * seek by any means (like reopening and linear reading) or other normally unreasonable
 * means that can be extremely slow.
 * This may be ignored by the seek code.
 */


/**
 * fseek() equivalent for AVIOContext.
 * @return new position or AVERROR.
 */
func Avio_seek(s *AVIOContext, offset int64, whence int32) int64 {
    return int64(C.avio_seek((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        C.longlong(offset), C.int(whence)))
}

/**
 * Skip given number of bytes forward
 * @return new position or AVERROR.
 */
func Avio_skip(s *AVIOContext, offset int64) int64 {
    return int64(C.avio_skip((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        C.longlong(offset)))
}

/**
 * ftell() equivalent for AVIOContext.
 * @return position or AVERROR.
 */
// avio_tell(AVIOContext*s)

/**
 * Get the filesize.
 * @return filesize or AVERROR
 */
func Avio_size(s *AVIOContext) int64 {
    return int64(C.avio_size((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Similar to feof() but also returns nonzero on read errors.
 * @return non zero if and only if at end of file or a read error happened when reading.
 */
func Avio_feof(s *AVIOContext) int32 {
    return int32(C.avio_feof((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Writes a formatted string to the context taking a va_list.
 * @return number of bytes written, < 0 on error.
 */
// not supported: avio_vprintf

/**
 * Writes a formatted string to the context.
 * @return number of bytes written, < 0 on error.
 */
//int avio_printf(AVIOContext *s, const char *fmt, ...) av_printf_format(2, 3);

/**
 * Write a NULL terminated array of strings to the context.
 * Usually you don't need to use this function directly but its macro wrapper,
 * avio_print.
 */
func Avio_print_string_array(s *AVIOContext, strings []*byte)  {
    C.avio_print_string_array((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (**C.char)(unsafe.Pointer(&strings[0])))
}

/**
 * Write strings (const char *) to the context.
 * This is a convenience macro around avio_print_string_array and it
 * automatically creates the string array from the variable argument list.
 * For simple string concatenations this function is more performant than using
 * avio_printf since it does not need a temporary buffer.
 */


/**
 * Force flushing of buffered data.
 *
 * For write streams, force the buffered data to be immediately written to the output,
 * without to wait to fill the internal buffer.
 *
 * For read streams, discard all currently buffered data, and advance the
 * reported file position to that of the underlying stream. This does not
 * read new data, and does not perform any seeks.
 */
func Avio_flush(s *AVIOContext)  {
    C.avio_flush((*C.struct_AVIOContext)(unsafe.Pointer(s)))
}

/**
 * Read size bytes from AVIOContext into buf.
 * @return number of bytes read or AVERROR
 */
func Avio_read(s *AVIOContext, buf *uint8, size int32) int32 {
    return int32(C.avio_read((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(size)))
}

/**
 * Read size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
 * to read fewer bytes than requested. The missing bytes can be read in the next
 * call. This always tries to read at least 1 byte.
 * Useful to reduce latency in certain cases.
 * @return number of bytes read or AVERROR
 */
func Avio_read_partial(s *AVIOContext, buf *uint8, size int32) int32 {
    return int32(C.avio_read_partial((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(size)))
}

/**
 * @name Functions for reading from AVIOContext
 * @{
 *
 * @note return 0 if EOF, so you cannot use it if EOF handling is
 *       necessary
 */
func          Avio_r8  (s *AVIOContext) int32 {
    return int32(C.avio_r8((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rl16(s *AVIOContext) uint32 {
    return uint32(C.avio_rl16((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rl24(s *AVIOContext) uint32 {
    return uint32(C.avio_rl24((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rl32(s *AVIOContext) uint32 {
    return uint32(C.avio_rl32((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func     Avio_rl64(s *AVIOContext) uint64 {
    return uint64(C.avio_rl64((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rb16(s *AVIOContext) uint32 {
    return uint32(C.avio_rb16((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rb24(s *AVIOContext) uint32 {
    return uint32(C.avio_rb24((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func Avio_rb32(s *AVIOContext) uint32 {
    return uint32(C.avio_rb32((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
func     Avio_rb64(s *AVIOContext) uint64 {
    return uint64(C.avio_rb64((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}
/**
 * @}
 */

/**
 * Read a string from pb into buf. The reading will terminate when either
 * a NULL character was encountered, maxlen bytes have been read, or nothing
 * more can be read from pb. The result is guaranteed to be NULL-terminated, it
 * will be truncated if buf is too small.
 * Note that the string is not interpreted or validated in any way, it
 * might get truncated in the middle of a sequence for multi-byte encodings.
 *
 * @return number of bytes read (is always <= maxlen).
 * If reading ends on EOF or error, the return value will be one more than
 * bytes actually read.
 */
func Avio_get_str(pb *AVIOContext, maxlen int32, buf *byte, buflen int32) int32 {
    return int32(C.avio_get_str((*C.struct_AVIOContext)(unsafe.Pointer(pb)), 
        C.int(maxlen), (*C.char)(unsafe.Pointer(buf)), C.int(buflen)))
}

/**
 * Read a UTF-16 string from pb and convert it to UTF-8.
 * The reading will terminate when either a null or invalid character was
 * encountered or maxlen bytes have been read.
 * @return number of bytes read (is always <= maxlen)
 */
func Avio_get_str16le(pb *AVIOContext, maxlen int32, buf *byte, buflen int32) int32 {
    return int32(C.avio_get_str16le((*C.struct_AVIOContext)(unsafe.Pointer(pb)), 
        C.int(maxlen), (*C.char)(unsafe.Pointer(buf)), C.int(buflen)))
}
func Avio_get_str16be(pb *AVIOContext, maxlen int32, buf *byte, buflen int32) int32 {
    return int32(C.avio_get_str16be((*C.struct_AVIOContext)(unsafe.Pointer(pb)), 
        C.int(maxlen), (*C.char)(unsafe.Pointer(buf)), C.int(buflen)))
}


/**
 * @name URL open modes
 * The flags argument to avio_open must be one of the following
 * constants, optionally ORed with other flags.
 * @{
 */
/**< read-only */
/**< write-only */
/**< read-write pseudo flag */
/**
 * @}
 */

/**
 * Use non-blocking mode.
 * If this flag is set, operations on the context will return
 * AVERROR(EAGAIN) if they can not be performed immediately.
 * If this flag is not set, operations on the context will never return
 * AVERROR(EAGAIN).
 * Note that this flag does not affect the opening/connecting of the
 * context. Connecting a protocol will always block if necessary (e.g. on
 * network protocols) but never hang (e.g. on busy devices).
 * Warning: non-blocking protocols is work-in-progress; this flag may be
 * silently ignored.
 */


/**
 * Use direct mode.
 * avio_read and avio_write should if possible be satisfied directly
 * instead of going through a buffer, and avio_seek will always
 * call the underlying seek function directly.
 */


/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
func Avio_open(s **AVIOContext, url *byte, flags int32) int32 {
    return int32(C.avio_open((**C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(url)), C.int(flags)))
}

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @param int_cb an interrupt callback to be used at the protocols level
 * @param options  A dictionary filled with protocol-private options. On return
 * this parameter will be destroyed and replaced with a dict containing options
 * that were not found. May be NULL.
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
func Avio_open2(s **AVIOContext, url *byte, flags int32,
               int_cb *AVIOInterruptCB, options **AVDictionary) int32 {
    return int32(C.avio_open2((**C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (*C.char)(unsafe.Pointer(url)), C.int(flags), 
        (*C.struct_AVIOInterruptCB)(unsafe.Pointer(int_cb)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Close the resource accessed by the AVIOContext s and free it.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_closep
 */
func Avio_close(s *AVIOContext) int32 {
    return int32(C.avio_close((*C.struct_AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Close the resource accessed by the AVIOContext *s, free it
 * and set the pointer pointing to it to NULL.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_close
 */
func Avio_closep(s **AVIOContext) int32 {
    return int32(C.avio_closep((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}


/**
 * Open a write only memory stream.
 *
 * @param s new IO context
 * @return zero if no error.
 */
func Avio_open_dyn_buf(s **AVIOContext) int32 {
    return int32(C.avio_open_dyn_buf((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Return the written size and a pointer to the buffer.
 * The AVIOContext stream is left intact.
 * The buffer must NOT be freed.
 * No padding is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
func Avio_get_dyn_buf(s *AVIOContext, pbuffer **uint8) int32 {
    return int32(C.avio_get_dyn_buf((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (**C.uchar)(unsafe.Pointer(pbuffer))))
}

/**
 * Return the written size and a pointer to the buffer. The buffer
 * must be freed with av_free().
 * Padding of AV_INPUT_BUFFER_PADDING_SIZE is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
func Avio_close_dyn_buf(s *AVIOContext, pbuffer **uint8) int32 {
    return int32(C.avio_close_dyn_buf((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (**C.uchar)(unsafe.Pointer(pbuffer))))
}

/**
 * Iterate through names of available protocols.
 *
 * @param opaque A private pointer representing current protocol.
 *        It must be a pointer to NULL on first iteration and will
 *        be updated by successive calls to avio_enum_protocols.
 * @param output If set to 1, iterate over output protocols,
 *               otherwise over input protocols.
 *
 * @return A static string containing the name of current protocol or NULL
 */
func Avio_enum_protocols(opaque *unsafe.Pointer, output int32) string {
    return C.GoString(C.avio_enum_protocols(opaque, C.int(output)))
}

/**
 * Get AVClass by names of available protocols.
 *
 * @return A AVClass of input protocol name or NULL
 */
func Avio_protocol_get_class(name *byte) *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avio_protocol_get_class(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Pause and resume playing - only meaningful if using a network streaming
 * protocol (e.g. MMS).
 *
 * @param h     IO context from which to call the read_pause function pointer
 * @param pause 1 for pause, 0 for resume
 */
func     Avio_pause(h *AVIOContext, pause int32) int32 {
    return int32(C.avio_pause((*C.struct_AVIOContext)(unsafe.Pointer(h)), 
        C.int(pause)))
}

/**
 * Seek to a given timestamp relative to some component stream.
 * Only meaningful if using a network streaming protocol (e.g. MMS.).
 *
 * @param h IO context from which to call the seek function pointers
 * @param stream_index The stream index that the timestamp is relative to.
 *        If stream_index is (-1) the timestamp should be in AV_TIME_BASE
 *        units from the beginning of the presentation.
 *        If a stream_index >= 0 is used and the protocol does not support
 *        seeking based on component streams, the call will fail.
 * @param timestamp timestamp in AVStream.time_base units
 *        or if there is no stream specified then in AV_TIME_BASE units.
 * @param flags Optional combination of AVSEEK_FLAG_BACKWARD, AVSEEK_FLAG_BYTE
 *        and AVSEEK_FLAG_ANY. The protocol may silently ignore
 *        AVSEEK_FLAG_BACKWARD and AVSEEK_FLAG_ANY, but AVSEEK_FLAG_BYTE will
 *        fail if used and not supported.
 * @return >= 0 on success
 * @see AVInputFormat::read_seek
 */
func Avio_seek_time(h *AVIOContext, stream_index int32,
                       timestamp int64, flags int32) int64 {
    return int64(C.avio_seek_time((*C.struct_AVIOContext)(unsafe.Pointer(h)), 
        C.int(stream_index), C.longlong(timestamp), C.int(flags)))
}

/* Avoid a warning. The header can not be included because it breaks c++. */
// type AVBPrint

/**
 * Read contents of h into print buffer, up to max_size bytes, or up to EOF.
 *
 * @return 0 for success (max_size bytes read or EOF reached), negative error
 * code otherwise
 */
func Avio_read_to_bprint(h *AVIOContext, pb *AVBPrint, max_size uint64) int32 {
    return int32(C.avio_read_to_bprint((*C.struct_AVIOContext)(unsafe.Pointer(h)), 
        (*C.struct_AVBPrint)(unsafe.Pointer(pb)), C.ulonglong(max_size)))
}

/**
 * Accept and allocate a client context on a server context.
 * @param  s the server context
 * @param  c the client context, must be unallocated
 * @return   >= 0 on success or a negative value corresponding
 *           to an AVERROR on failure
 */
func Avio_accept(s *AVIOContext, c **AVIOContext) int32 {
    return int32(C.avio_accept((*C.struct_AVIOContext)(unsafe.Pointer(s)), 
        (**C.struct_AVIOContext)(unsafe.Pointer(c))))
}

/**
 * Perform one step of the protocol handshake to accept a new client.
 * This function must be called on a client returned by avio_accept() before
 * using it as a read/write context.
 * It is separate from avio_accept() because it may block.
 * A step of the handshake is defined by places where the application may
 * decide to change the proceedings.
 * For example, on a protocol with a request header and a reply header, each
 * one can constitute a step because the application may use the parameters
 * from the request to change parameters in the reply; or each individual
 * chunk of the request can constitute a step.
 * If the handshake is already finished, avio_handshake() does nothing and
 * returns 0 immediately.
 *
 * @param  c the client context to perform the handshake on
 * @return   0   on a complete and successful handshake
 *           > 0 if the handshake progressed, but is not complete
 *           < 0 for an AVERROR code
 */
func Avio_handshake(c *AVIOContext) int32 {
    return int32(C.avio_handshake((*C.struct_AVIOContext)(unsafe.Pointer(c))))
}
                            
