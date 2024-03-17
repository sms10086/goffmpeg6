// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2003 Ivan Kalvachev
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


const AV_XVMC_ID =                     0x1DC711C0   


                      
                      

/**
 * @file
 * @ingroup lavc_codec_hwaccel_xvmc
 * Public libavcodec XvMC header.
 */

                                                                                             

                                

                                 
                    

/**
 * @defgroup lavc_codec_hwaccel_xvmc XvMC
 * @ingroup lavc_codec_hwaccel
 *
 * @{
 */

/**< special value to ensure that regular pixel routines haven't corrupted the struct
                                                       the number is 1337 speak for the letters IDCT MCo (motion compensation) */

type xvmc_pix_fmt struct {
    Xvmc_id int32
    Data_blocks *int16
    Mv_blocks *XvMCMacroBlock
    Allocated_mv_blocks int32
    Allocated_data_blocks int32
    Idct int32
    Unsigned_intra int32
    P_surface *XvMCSurface
    P_past_surface *XvMCSurface
    P_future_surface *XvMCSurface
    Picture_structure uint32
    Flags uint32
    Start_mv_blocks_num int32
    Filled_mv_blocks_num int32
    Next_free_data_block_num int32
}


/**
 * @}
 */

                           
