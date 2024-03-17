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





                                
                                

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_DXVA2.
 *
 * Only fixed-size pools are supported.
 *
 * For user-allocated pools, AVHWFramesContext.pool must return AVBufferRefs
 * with the data pointer set to a pointer to IDirect3DSurface9.
 */

                 
                     

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVDXVA2DeviceContext struct {
    Devmgr *IDirect3DDeviceManager9
}


/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVDXVA2FramesContext struct {
    Surface_type uint32
    Surfaces **IDirect3DSurface9
    Nb_surfaces int32
    Decoder_to_release *IDirectXVideoDecoder
}


                                     
