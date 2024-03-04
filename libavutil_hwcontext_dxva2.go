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
//#include <d3d9.h>
//#include <dxva2api.h>
//#include "libavutil/hwcontext_dxva2.h"
import "C"




                                
                                

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
type AVDXVA2DeviceContext C.struct_AVDXVA2DeviceContext

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVDXVA2FramesContext C.struct_AVDXVA2FramesContext

                                     
