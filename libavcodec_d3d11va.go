// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Direct3D11 HW acceleration
 *
 * copyright (c) 2009 Laurent Aimar
 * copyright (c) 2015 Steve Lhomme
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

//#cgo pkg-config: libavcodec
//#include <stdint.h>
//#include <d3d11.h>
//#include "libavcodec/d3d11va.h"
import "C"
import (
    "unsafe"
)

const FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG = 1
const FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO = 2


                         
                         

/**
 * @file
 * @ingroup lavc_codec_hwaccel_d3d11va
 * Public libavcodec D3D11VA header.
 */

                                                   
                   
                           
      

                   
                  

/**
 * @defgroup lavc_codec_hwaccel_d3d11va Direct3D11
 * @ingroup lavc_codec_hwaccel
 *
 * @{
 */

///< Work around for Direct3D11 and old UVD/UVD+ ATI video cards
///< Work around for Direct3D11 and old Intel GPUs with ClearVideo interface

/**
 * This structure is used to provides the necessary configurations and data
 * to the Direct3D11 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 *
 * Use av_d3d11va_alloc_context() exclusively to allocate an AVD3D11VAContext.
 */
type AVD3D11VAContext struct {
    Decoder *ID3D11VideoDecoder
    Video_context *ID3D11VideoContext
    Cfg *D3D11_VIDEO_DECODER_CONFIG
    Surface_count uint32
    Surface **ID3D11VideoDecoderOutputView
    Workaround uint64
    Report_id uint32
    Context_mutex HANDLE
}


/**
 * Allocate an AVD3D11VAContext.
 *
 * @return Newly-allocated AVD3D11VAContext or NULL on failure.
 */
func Av_d3d11va_alloc_context() *AVD3D11VAContext {
    return (*AVD3D11VAContext)(unsafe.Pointer(C.av_d3d11va_alloc_context()))
}

/**
 * @}
 */

                              
