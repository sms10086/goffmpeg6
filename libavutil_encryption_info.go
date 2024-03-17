// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/**
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
//#include "libavutil/encryption_info.h"
import "C"
import (
    "unsafe"
)



                                
                                

                   
                   

type AVSubsampleEncryptionInfo struct {
    Bytes_of_clear_data uint32
    Bytes_of_protected_data uint32
}


/**
 * This describes encryption info for a packet.  This contains frame-specific
 * info for how to decrypt the packet before passing it to the decoder.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInfo struct {
    Scheme uint32
    Crypt_byte_block uint32
    Skip_byte_block uint32
    Key_id *uint8
    Key_id_size uint32
    Iv *uint8
    Iv_size uint32
    Subsamples *AVSubsampleEncryptionInfo
    Subsample_count uint32
}


/**
 * This describes info used to initialize an encryption key system.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInitInfo struct {
    System_id *uint8
    System_id_size uint32
    Key_ids **uint8
    Num_key_ids uint32
    Key_id_size uint32
    Data *uint8
    Data_size uint32
    Next *AVEncryptionInitInfo
}


/**
 * Allocates an AVEncryptionInfo structure and sub-pointers to hold the given
 * number of subsamples.  This will allocate pointers for the key ID, IV,
 * and subsample entries, set the size members, and zero-initialize the rest.
 *
 * @param subsample_count The number of subsamples.
 * @param key_id_size The number of bytes in the key ID, should be 16.
 * @param iv_size The number of bytes in the IV, should be 16.
 *
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
func Av_encryption_info_alloc(subsample_count uint32, key_id_size uint32, iv_size uint32) *AVEncryptionInfo {
    return (*AVEncryptionInfo)(unsafe.Pointer(C.av_encryption_info_alloc(
        C.uint(subsample_count), C.uint(key_id_size), C.uint(iv_size))))
}

/**
 * Allocates an AVEncryptionInfo structure with a copy of the given data.
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
func Av_encryption_info_clone(info *AVEncryptionInfo) *AVEncryptionInfo {
    return (*AVEncryptionInfo)(unsafe.Pointer(C.av_encryption_info_clone(
        (*C.struct_AVEncryptionInfo)(unsafe.Pointer(info)))))
}

/**
 * Frees the given encryption info object.  This MUST NOT be used to free the
 * side-data data pointer, that should use normal side-data methods.
 */
func Av_encryption_info_free(info *AVEncryptionInfo)  {
    C.av_encryption_info_free((*C.struct_AVEncryptionInfo)(unsafe.Pointer(info)))
}

/**
 * Creates a copy of the AVEncryptionInfo that is contained in the given side
 * data.  The resulting object should be passed to av_encryption_info_free()
 * when done.
 *
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
func Av_encryption_info_get_side_data(side_data *uint8, side_data_size uint64) *AVEncryptionInfo {
    return (*AVEncryptionInfo)(unsafe.Pointer(C.av_encryption_info_get_side_data(
        (*C.uchar)(unsafe.Pointer(side_data)), C.ulonglong(side_data_size))))
}

/**
 * Allocates and initializes side data that holds a copy of the given encryption
 * info.  The resulting pointer should be either freed using av_free or given
 * to av_packet_add_side_data().
 *
 * @return The new side-data pointer, or NULL.
 */
func Av_encryption_info_add_side_data(
      info *AVEncryptionInfo, side_data_size *uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_encryption_info_add_side_data(
        (*C.struct_AVEncryptionInfo)(unsafe.Pointer(info)), 
        (*C.ulonglong)(unsafe.Pointer(side_data_size)))))
}


/**
 * Allocates an AVEncryptionInitInfo structure and sub-pointers to hold the
 * given sizes.  This will allocate pointers and set all the fields.
 *
 * @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
func Av_encryption_init_info_alloc(
    system_id_size uint32, num_key_ids uint32, key_id_size uint32, data_size uint32) *AVEncryptionInitInfo {
    return (*AVEncryptionInitInfo)(unsafe.Pointer(C.av_encryption_init_info_alloc(
        C.uint(system_id_size), C.uint(num_key_ids), C.uint(key_id_size), 
        C.uint(data_size))))
}

/**
 * Frees the given encryption init info object.  This MUST NOT be used to free
 * the side-data data pointer, that should use normal side-data methods.
 */
func Av_encryption_init_info_free(info *AVEncryptionInitInfo)  {
    C.av_encryption_init_info_free(
        (*C.struct_AVEncryptionInitInfo)(unsafe.Pointer(info)))
}

/**
 * Creates a copy of the AVEncryptionInitInfo that is contained in the given
 * side data.  The resulting object should be passed to
 * av_encryption_init_info_free() when done.
 *
 * @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
func Av_encryption_init_info_get_side_data(
    side_data *uint8, side_data_size uint64) *AVEncryptionInitInfo {
    return (*AVEncryptionInitInfo)(unsafe.Pointer(C.av_encryption_init_info_get_side_data(
        (*C.uchar)(unsafe.Pointer(side_data)), C.ulonglong(side_data_size))))
}

/**
 * Allocates and initializes side data that holds a copy of the given encryption
 * init info.  The resulting pointer should be either freed using av_free or
 * given to av_packet_add_side_data().
 *
 * @return The new side-data pointer, or NULL.
 */
func Av_encryption_init_info_add_side_data(
    info *AVEncryptionInitInfo, side_data_size *uint64) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_encryption_init_info_add_side_data(
        (*C.struct_AVEncryptionInitInfo)(unsafe.Pointer(info)), 
        (*C.ulonglong)(unsafe.Pointer(side_data_size)))))
}

                                     
