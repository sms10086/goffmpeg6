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

import (
    "unsafe"
)



                                  
                                  

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_D3D11VA.
 *
 * The default pool implementation will be fixed-size if initial_pool_size is
 * set (and allocate elements from an array texture). Otherwise it will allocate
 * individual textures. Be aware that decoding requires a single array texture.
 *
 * Using sw_format==AV_PIX_FMT_YUV420P has special semantics, and maps to
 * DXGI_FORMAT_420_OPAQUE. av_hwframe_transfer_data() is not supported for
 * this format. Refer to MSDN for details.
 *
 * av_hwdevice_ctx_create() for this device type supports a key named "debug"
 * for the AVDictionary entry. If this is set to any value, the device creation
 * code will try to load various supported D3D debugging layers.
 */

                  
                   

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVD3D11VADeviceContext struct {
    Device *ID3D11Device
    Device_context *ID3D11DeviceContext
    Video_device *ID3D11VideoDevice
    Video_context *ID3D11VideoContext
    Lock uintptr
    Unlock uintptr
    Lock_ctx unsafe.Pointer
}


/**
 * D3D11 frame descriptor for pool allocation.
 *
 * In user-allocated pools, AVHWFramesContext.pool must return AVBufferRefs
 * with the data pointer pointing at an object of this type describing the
 * planes of the frame.
 *
 * This has no use outside of custom allocation, and AVFrame AVBufferRef do not
 * necessarily point to an instance of this struct.
 */
type AVD3D11FrameDescriptor struct {
    Texture *ID3D11Texture2D
    Index uintptr
}


/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVD3D11VAFramesContext struct {
    Texture *ID3D11Texture2D
    BindFlags uint32
    MiscFlags uint32
    Texture_infos *AVD3D11FrameDescriptor
}


                                       
