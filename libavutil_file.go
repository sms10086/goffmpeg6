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
//#include <stddef.h>
//#include <stdint.h>
//#include "libavutil/version.h"
//#include "libavutil/attributes.h"
//#include "libavutil/file.h"
import "C"
import (
    "unsafe"
)



                     
                     

                   
                   

                    
                       

/**
 * @file
 * Misc file utilities.
 */

/**
 * Read the file with name filename, and put its content in a newly
 * allocated buffer or map it with mmap() when available.
 * In case of success set *bufptr to the read or mmapped buffer, and
 * *size to the size in bytes of the buffer in *bufptr.
 * Unlike mmap this function succeeds with zero sized files, in this
 * case *bufptr will be set to NULL and *size will be set to 0.
 * The returned buffer must be released with av_file_unmap().
 *
 * @param filename path to the file
 * @param[out] bufptr pointee is set to the mapped or allocated buffer
 * @param[out] size pointee is set to the size in bytes of the buffer
 * @param log_offset loglevel offset used for logging
 * @param log_ctx context used for logging
 * @return a non negative number in case of success, a negative value
 * corresponding to an AVERROR error code in case of failure
 */
func Av_file_map(filename *byte, bufptr **uint8, size *uint64,
                log_offset int32, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_file_map((*C.char)(unsafe.Pointer(filename)), 
        (**C.uchar)(unsafe.Pointer(bufptr)), 
        (*C.ulonglong)(unsafe.Pointer(size)), C.int(log_offset), log_ctx))
}

/**
 * Unmap or free the buffer bufptr created by av_file_map().
 *
 * @param bufptr the buffer previously created with av_file_map()
 * @param size size in bytes of bufptr, must be the same as returned
 * by av_file_map()
 */
func Av_file_unmap(bufptr *uint8, size uint64)  {
    C.av_file_unmap((*C.uchar)(unsafe.Pointer(bufptr)), C.ulonglong(size))
}

                        
   
                                                         
                                                         
                                                                               
                                                                                
                         
                                      
                                                                       
                                                                             
                                                                    
                                                                                    
   
                    
                                                                                    
      

                          
