// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * DXVA2 HW acceleration
 *
 * copyright (c) 2009 Laurent Aimar
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


//const FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG = 1
//const FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO = 2


                       
                       

/**
 * @file
 * @ingroup lavc_codec_hwaccel_dxva2
 * Public libavcodec DXVA2 header.
 */

                                                   
                   
                           
      

                   
                 
                     

/**
 * @defgroup lavc_codec_hwaccel_dxva2 DXVA2
 * @ingroup lavc_codec_hwaccel
 *
 * @{
 */

///< Work around for DXVA2 and old UVD/UVD+ ATI video cards
///< Work around for DXVA2 and old Intel GPUs with ClearVideo interface

/**
 * This structure is used to provides the necessary configurations and data
 * to the DXVA2 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 */
type dxva_context struct {
    Decoder *IDirectXVideoDecoder
    Cfg *DXVA2_ConfigPictureDecode
    Surface_count uint32
    Surface *LPDIRECT3DSURFACE9
    Workaround uint64
    Report_id uint32
}


/**
 * @}
 */

                            
