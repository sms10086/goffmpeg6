// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2022 Pierre-Anthony Lemieux <pal@palemieux.com>
 *                    Zane van Iperen <zane@zanevaniperen.com>
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
//#include <stdint.h>
//#include <string.h>
//#include "libavutil/uuid.h"
import "C"
import (
    "unsafe"
)

const AV_PRI_UUID = "%02hhx%02hhx%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx%02hhx%02hhx%02hhx%02hhx"
const AV_PRI_URN_UUID = "urn:uuid:%02hhx%02hhx%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx-%02hhx%02hhx%02hhx%02hhx%02hhx%02hhx"
const AV_UUID_LEN = 16


/**
 * @file
 * UUID parsing and serialization utilities.
 * The library treats the UUID as an opaque sequence of 16 unsigned bytes,
 * i.e. ignoring the internal layout of the UUID, which depends on the type
 * of the UUID.
 *
 * @author Pierre-Anthony Lemieux <pal@palemieux.com>
 * @author Zane van Iperen <zane@zanevaniperen.com>
 */

                     
                     

                   
                   





/* AV_UUID_ARG() is used together with AV_PRI_UUID() or AV_PRI_URN_UUID
 * to print UUIDs, e.g.
 * av_log(NULL, AV_LOG_DEBUG, "UUID: " AV_PRI_UUID, AV_UUID_ARG(uuid));
 */




/* Binary representation of a UUID */
type AVUUID [AV_UUID_LEN]uint8

/**
 * Parses a string representation of a UUID formatted according to IETF RFC 4122
 * into an AVUUID. The parsing is case-insensitive. The string must be 37
 * characters long, including the terminating NUL character.
 *
 * Example string representation: "2fceebd0-7017-433d-bafb-d073a7116696"
 *
 * @param[in]  in  String representation of a UUID,
 *                 e.g. 2fceebd0-7017-433d-bafb-d073a7116696
 * @param[out] uu  AVUUID
 * @return         A non-zero value in case of an error.
 */
func Av_uuid_parse(in *byte, uu AVUUID) int32 {
    return int32(C.av_uuid_parse((*C.char)(unsafe.Pointer(in)), 
        (*C.uchar)(unsafe.Pointer(&uu[0]))))
}

/**
 * Parses a URN representation of a UUID, as specified at IETF RFC 4122,
 * into an AVUUID. The parsing is case-insensitive. The string must be 46
 * characters long, including the terminating NUL character.
 *
 * Example string representation: "urn:uuid:2fceebd0-7017-433d-bafb-d073a7116696"
 *
 * @param[in]  in  URN UUID
 * @param[out] uu  AVUUID
 * @return         A non-zero value in case of an error.
 */
func Av_uuid_urn_parse(in *byte, uu AVUUID) int32 {
    return int32(C.av_uuid_urn_parse((*C.char)(unsafe.Pointer(in)), 
        (*C.uchar)(unsafe.Pointer(&uu[0]))))
}

/**
 * Parses a string representation of a UUID formatted according to IETF RFC 4122
 * into an AVUUID. The parsing is case-insensitive.
 *
 * @param[in]  in_start Pointer to the first character of the string representation
 * @param[in]  in_end   Pointer to the character after the last character of the
 *                      string representation. That memory location is never
 *                      accessed. It is an error if `in_end - in_start != 36`.
 * @param[out] uu       AVUUID
 * @return              A non-zero value in case of an error.
 */
func Av_uuid_parse_range(in_start *byte, in_end *byte, uu AVUUID) int32 {
    return int32(C.av_uuid_parse_range((*C.char)(unsafe.Pointer(in_start)), 
        (*C.char)(unsafe.Pointer(in_end)), (*C.uchar)(unsafe.Pointer(&uu[0]))))
}

/**
 * Serializes a AVUUID into a string representation according to IETF RFC 4122.
 * The string is lowercase and always 37 characters long, including the
 * terminating NUL character.
 *
 * @param[in]  uu  AVUUID
 * @param[out] out Pointer to an array of no less than 37 characters.
 */
func Av_uuid_unparse(uu AVUUID, out *byte)  {
    C.av_uuid_unparse((*C.uchar)(unsafe.Pointer(&uu[0])), 
        (*C.char)(unsafe.Pointer(out)))
}

/**
 * Compares two UUIDs for equality.
 *
 * @param[in]  uu1  AVUUID
 * @param[in]  uu2  AVUUID
 * @return          Nonzero if uu1 and uu2 are identical, 0 otherwise
 */
// av_uuid_equal(constAVUUIDuu1,constAVUUIDuu2)

/**
 * Copies the bytes of src into dest.
 *
 * @param[out]  dest  AVUUID
 * @param[in]   src   AVUUID
 */
// av_uuid_copy(AVUUIDdest,constAVUUIDsrc)

/**
 * Sets a UUID to the nil UUID, i.e. a UUID with have all
 * its 128 bits set to zero.
 *
 * @param[in,out]  uu  UUID to be set to the nil UUID
 */
// av_uuid_nil(AVUUIDuu)

                          
