// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2014 Tim Walker <tdskywalker@gmail.com>
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
//#include "libavutil/downmix_info.h"
import "C"
import (
    "unsafe"
)



                             
                             

                  

/**
 * @file
 * audio downmix medatata
 */

/**
 * @addtogroup lavu_audio
 * @{
 */

/**
 * @defgroup downmix_info Audio downmix metadata
 * @{
 */

/**
 * Possible downmix types.
 */
type AVDownmixType C.enum_AVDownmixType

/**
 * This structure describes optional metadata relevant to a downmix procedure.
 *
 * All fields are set by the decoder to the value indicated in the audio
 * bitstream (if present), or to a "sane" default otherwise.
 */
type AVDownmixInfo C.struct_AVDownmixInfo

/**
 * Get a frame's AV_FRAME_DATA_DOWNMIX_INFO side data for editing.
 *
 * If the side data is absent, it is created and added to the frame.
 *
 * @param frame the frame for which the side data is to be obtained or created
 *
 * @return the AVDownmixInfo structure to be edited by the caller, or NULL if
 *         the structure cannot be allocated.
 */
func Av_downmix_info_update_side_data(frame *AVFrame) *AVDownmixInfo {
    return (*AVDownmixInfo)(unsafe.Pointer(C.av_downmix_info_update_side_data(
        (*C.AVFrame)(unsafe.Pointer(frame)))))
}

/**
 * @}
 */

/**
 * @}
 */

                                  
