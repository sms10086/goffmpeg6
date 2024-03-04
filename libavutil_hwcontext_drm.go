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
//#include "libavutil/hwcontext_drm.h"
import "C"



                              
                              

                   
                   

/**
 * @file
 * API-specific header for AV_HWDEVICE_TYPE_DRM.
 *
 * Internal frame allocation is not currently supported - all frames
 * must be allocated by the user.  Thus AVHWFramesContext is always
 * NULL, though this may change if support for frame allocation is
 * added in future.
 */



/**
 * DRM object descriptor.
 *
 * Describes a single DRM object, addressing it as a PRIME file
 * descriptor.
 */
type AVDRMObjectDescriptor C.struct_AVDRMObjectDescriptor

/**
 * DRM plane descriptor.
 *
 * Describes a single plane of a layer, which is contained within
 * a single object.
 */
type AVDRMPlaneDescriptor C.struct_AVDRMPlaneDescriptor

/**
 * DRM layer descriptor.
 *
 * Describes a single layer within a frame.  This has the structure
 * defined by its format, and will contain one or more planes.
 */
type AVDRMLayerDescriptor C.struct_AVDRMLayerDescriptor

/**
 * DRM frame descriptor.
 *
 * This is used as the data pointer for AV_PIX_FMT_DRM_PRIME frames.
 * It is also used by user-allocated frame pools - allocating in
 * AVHWFramesContext.pool must return AVBufferRefs which contain
 * an object of this type.
 *
 * The fields of this structure should be set such it can be
 * imported directly by EGL using the EGL_EXT_image_dma_buf_import
 * and EGL_EXT_image_dma_buf_import_modifiers extensions.
 * (Note that the exact layout of a particular format may vary between
 * platforms - we only specify that the same platform should be able
 * to import it.)
 *
 * The total number of planes must not exceed AV_DRM_MAX_PLANES, and
 * the order of the planes by increasing layer index followed by
 * increasing plane index must be the same as the order which would
 * be used for the data pointers in the equivalent software format.
 */
type AVDRMFrameDescriptor C.struct_AVDRMFrameDescriptor

/**
 * DRM device.
 *
 * Allocated as AVHWDeviceContext.hwctx.
 */
type AVDRMDeviceContext C.struct_AVDRMDeviceContext

                                   
