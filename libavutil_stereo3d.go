// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2013 Vittorio Giovara <vittorio.giovara@gmail.com>
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
//#include "libavutil/frame.h"
//#include "libavutil/stereo3d.h"
import "C"
import (
    "unsafe"
)

const AV_STEREO3D_FLAG_INVERT = (1 << 0)


/**
 * @file
 * @ingroup lavu_video_stereo3d
 * Stereoscopic video
 */

                         
                         

                   

                  

/**
 * @defgroup lavu_video_stereo3d Stereo3D types and functions
 * @ingroup lavu_video
 *
 * A stereoscopic video file consists in multiple views embedded in a single
 * frame, usually describing two views of a scene. This file describes all
 * possible codec-independent view arrangements.
 *
 * @{
 */

/**
 * List of possible 3D Types
 */
type AVStereo3DType C.enum_AVStereo3DType

/**
 * List of possible view types.
 */
type AVStereo3DView C.enum_AVStereo3DView

/**
 * Inverted views, Right/Bottom represents the left view.
 */


/**
 * Stereo 3D type: this structure describes how two videos are packed
 * within a single video surface, with additional information as needed.
 *
 * @note The struct must be allocated with av_stereo3d_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVStereo3D C.struct_AVStereo3D

/**
 * Allocate an AVStereo3D structure and set its fields to default values.
 * The resulting struct can be freed using av_freep().
 *
 * @return An AVStereo3D filled with default values or NULL on failure.
 */
func Av_stereo3d_alloc() *AVStereo3D {
    return (*AVStereo3D)(unsafe.Pointer(C.av_stereo3d_alloc()))
}

/**
 * Allocate a complete AVFrameSideData and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVStereo3D structure to be filled by caller.
 */
func Av_stereo3d_create_side_data(frame *AVFrame) *AVStereo3D {
    return (*AVStereo3D)(unsafe.Pointer(C.av_stereo3d_create_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)))))
}

/**
 * Provide a human-readable name of a given stereo3d type.
 *
 * @param type The input stereo3d type value.
 *
 * @return The name of the stereo3d value, or "unknown".
 */
func Av_stereo3d_type_name(typex uint32) string {
    return C.GoString(C.av_stereo3d_type_name(C.uint(typex)))
}

/**
 * Get the AVStereo3DType form a human-readable name.
 *
 * @param name The input string.
 *
 * @return The AVStereo3DType value, or -1 if not found.
 */
func Av_stereo3d_from_name(name *byte) int32 {
    return int32(C.av_stereo3d_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * @}
 */

                              
