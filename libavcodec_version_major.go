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


//const LIBAVCODEC_VERSION_MAJOR = 60
const FF_API_INIT_PACKET = true
const FF_API_IDCT_NONE = true
const FF_API_SVTAV1_OPTS = true
const FF_API_AYUV_CODECID = true
const FF_API_VT_OUTPUT_CALLBACK = true
const FF_API_AVCODEC_CHROMA_POS = true
const FF_API_VT_HWACCEL_CONTEXT = true
const FF_API_AVCTX_FRAME_NUMBER = true
const FF_API_SLICE_OFFSET = true
const FF_API_SUBFRAMES = true
const FF_API_TICKS_PER_FRAME = true
const FF_API_DROPCHANGED = true
const FF_API_AVFFT = true
const FF_API_FF_PROFILE_LEVEL = true
const FF_CODEC_CRYSTAL_HD = true


                               
                               

/**
 * @file
 * @ingroup libavc
 * Libavcodec version macros.
 */



/**
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 */

















// reminder to remove CrystalHD decoders on next major bump


                                    
