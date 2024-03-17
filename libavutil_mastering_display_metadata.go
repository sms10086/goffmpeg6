// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2016 Neil Birkbeck <neil.birkbeck@gmail.com>
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
//#include "libavutil/frame.h"
//#include "libavutil/rational.h"
//#include "libavutil/mastering_display_metadata.h"
import "C"
import (
    "unsafe"
)



                                           
                                           

                  
                     


/**
 * Mastering display metadata capable of representing the color volume of
 * the display used to master the content (SMPTE 2086:2014).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_mastering_display_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVMasteringDisplayMetadata struct {
    Display_primaries[3] [2]AVRational
    White_point [2]AVRational
    Min_luminance AVRational
    Max_luminance AVRational
    Has_primaries int32
    Has_luminance int32
}


/**
 * Allocate an AVMasteringDisplayMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVMasteringDisplayMetadata filled with default values or NULL
 *         on failure.
 */
func Av_mastering_display_metadata_alloc() *AVMasteringDisplayMetadata {
    return (*AVMasteringDisplayMetadata)(unsafe.Pointer(C.av_mastering_display_metadata_alloc()))
}

/**
 * Allocate a complete AVMasteringDisplayMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVMasteringDisplayMetadata structure to be filled by caller.
 */
func Av_mastering_display_metadata_create_side_data(frame *AVFrame) *AVMasteringDisplayMetadata {
    return (*AVMasteringDisplayMetadata)(unsafe.Pointer(C.av_mastering_display_metadata_create_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

/**
 * Content light level needed by to transmit HDR over HDMI (CTA-861.3).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_content_light_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVContentLightMetadata struct {
    MaxCLL uint32
    MaxFALL uint32
}


/**
 * Allocate an AVContentLightMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVContentLightMetadata filled with default values or NULL
 *         on failure.
 */
func Av_content_light_metadata_alloc(size *uint64) *AVContentLightMetadata {
    return (*AVContentLightMetadata)(unsafe.Pointer(C.av_content_light_metadata_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * Allocate a complete AVContentLightMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVContentLightMetadata structure to be filled by caller.
 */
func Av_content_light_metadata_create_side_data(frame *AVFrame) *AVContentLightMetadata {
    return (*AVContentLightMetadata)(unsafe.Pointer(C.av_content_light_metadata_create_side_data(
        (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

                                                
