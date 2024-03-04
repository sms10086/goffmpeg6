// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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
//#include "libavutil/avconfig.h"
//#include "libavutil/version.h"
//#include "libavutil/pixfmt.h"
import "C"

const AVPALETTE_SIZE = 1024
const AVPALETTE_COUNT = 256
//const AV_PIX_FMT_RGB32 = AV_PIX_FMT_NE(ARGB, BGRA)
//const AV_PIX_FMT_RGB32_1 = AV_PIX_FMT_NE(RGBA, ABGR)
//const AV_PIX_FMT_BGR32 = AV_PIX_FMT_NE(ABGR, RGBA)
//const AV_PIX_FMT_BGR32_1 = AV_PIX_FMT_NE(BGRA, ARGB)
//const AV_PIX_FMT_0RGB32 = AV_PIX_FMT_NE(0RGB, BGR0)
//const AV_PIX_FMT_0BGR32 = AV_PIX_FMT_NE(0BGR, RGB0)
//const AV_PIX_FMT_GRAY9 = AV_PIX_FMT_NE(GRAY9BE,  GRAY9LE)
//const AV_PIX_FMT_GRAY10 = AV_PIX_FMT_NE(GRAY10BE, GRAY10LE)
//const AV_PIX_FMT_GRAY12 = AV_PIX_FMT_NE(GRAY12BE, GRAY12LE)
//const AV_PIX_FMT_GRAY14 = AV_PIX_FMT_NE(GRAY14BE, GRAY14LE)
//const AV_PIX_FMT_GRAY16 = AV_PIX_FMT_NE(GRAY16BE, GRAY16LE)
//const AV_PIX_FMT_YA16 = AV_PIX_FMT_NE(YA16BE,   YA16LE)
//const AV_PIX_FMT_RGB48 = AV_PIX_FMT_NE(RGB48BE,  RGB48LE)
//const AV_PIX_FMT_RGB565 = AV_PIX_FMT_NE(RGB565BE, RGB565LE)
//const AV_PIX_FMT_RGB555 = AV_PIX_FMT_NE(RGB555BE, RGB555LE)
//const AV_PIX_FMT_RGB444 = AV_PIX_FMT_NE(RGB444BE, RGB444LE)
//const AV_PIX_FMT_RGBA64 = AV_PIX_FMT_NE(RGBA64BE, RGBA64LE)
//const AV_PIX_FMT_BGR48 = AV_PIX_FMT_NE(BGR48BE,  BGR48LE)
//const AV_PIX_FMT_BGR565 = AV_PIX_FMT_NE(BGR565BE, BGR565LE)
//const AV_PIX_FMT_BGR555 = AV_PIX_FMT_NE(BGR555BE, BGR555LE)
//const AV_PIX_FMT_BGR444 = AV_PIX_FMT_NE(BGR444BE, BGR444LE)
//const AV_PIX_FMT_BGRA64 = AV_PIX_FMT_NE(BGRA64BE, BGRA64LE)
//const AV_PIX_FMT_YUV420P9 = AV_PIX_FMT_NE(YUV420P9BE , YUV420P9LE)
//const AV_PIX_FMT_YUV422P9 = AV_PIX_FMT_NE(YUV422P9BE , YUV422P9LE)
//const AV_PIX_FMT_YUV444P9 = AV_PIX_FMT_NE(YUV444P9BE , YUV444P9LE)
//const AV_PIX_FMT_YUV420P10 = AV_PIX_FMT_NE(YUV420P10BE, YUV420P10LE)
//const AV_PIX_FMT_YUV422P10 = AV_PIX_FMT_NE(YUV422P10BE, YUV422P10LE)
//const AV_PIX_FMT_YUV440P10 = AV_PIX_FMT_NE(YUV440P10BE, YUV440P10LE)
//const AV_PIX_FMT_YUV444P10 = AV_PIX_FMT_NE(YUV444P10BE, YUV444P10LE)
//const AV_PIX_FMT_YUV420P12 = AV_PIX_FMT_NE(YUV420P12BE, YUV420P12LE)
//const AV_PIX_FMT_YUV422P12 = AV_PIX_FMT_NE(YUV422P12BE, YUV422P12LE)
//const AV_PIX_FMT_YUV440P12 = AV_PIX_FMT_NE(YUV440P12BE, YUV440P12LE)
//const AV_PIX_FMT_YUV444P12 = AV_PIX_FMT_NE(YUV444P12BE, YUV444P12LE)
//const AV_PIX_FMT_YUV420P14 = AV_PIX_FMT_NE(YUV420P14BE, YUV420P14LE)
//const AV_PIX_FMT_YUV422P14 = AV_PIX_FMT_NE(YUV422P14BE, YUV422P14LE)
//const AV_PIX_FMT_YUV444P14 = AV_PIX_FMT_NE(YUV444P14BE, YUV444P14LE)
//const AV_PIX_FMT_YUV420P16 = AV_PIX_FMT_NE(YUV420P16BE, YUV420P16LE)
//const AV_PIX_FMT_YUV422P16 = AV_PIX_FMT_NE(YUV422P16BE, YUV422P16LE)
//const AV_PIX_FMT_YUV444P16 = AV_PIX_FMT_NE(YUV444P16BE, YUV444P16LE)
//const AV_PIX_FMT_GBRP9 = AV_PIX_FMT_NE(GBRP9BE ,    GBRP9LE)
//const AV_PIX_FMT_GBRP10 = AV_PIX_FMT_NE(GBRP10BE,    GBRP10LE)
//const AV_PIX_FMT_GBRP12 = AV_PIX_FMT_NE(GBRP12BE,    GBRP12LE)
//const AV_PIX_FMT_GBRP14 = AV_PIX_FMT_NE(GBRP14BE,    GBRP14LE)
//const AV_PIX_FMT_GBRP16 = AV_PIX_FMT_NE(GBRP16BE,    GBRP16LE)
//const AV_PIX_FMT_GBRAP10 = AV_PIX_FMT_NE(GBRAP10BE,   GBRAP10LE)
//const AV_PIX_FMT_GBRAP12 = AV_PIX_FMT_NE(GBRAP12BE,   GBRAP12LE)
//const AV_PIX_FMT_GBRAP14 = AV_PIX_FMT_NE(GBRAP14BE,   GBRAP14LE)
//const AV_PIX_FMT_GBRAP16 = AV_PIX_FMT_NE(GBRAP16BE,   GBRAP16LE)
//const AV_PIX_FMT_BAYER_BGGR16 = AV_PIX_FMT_NE(BAYER_BGGR16BE,    BAYER_BGGR16LE)
//const AV_PIX_FMT_BAYER_RGGB16 = AV_PIX_FMT_NE(BAYER_RGGB16BE,    BAYER_RGGB16LE)
//const AV_PIX_FMT_BAYER_GBRG16 = AV_PIX_FMT_NE(BAYER_GBRG16BE,    BAYER_GBRG16LE)
//const AV_PIX_FMT_BAYER_GRBG16 = AV_PIX_FMT_NE(BAYER_GRBG16BE,    BAYER_GRBG16LE)
//const AV_PIX_FMT_GBRPF32 = AV_PIX_FMT_NE(GBRPF32BE,  GBRPF32LE)
//const AV_PIX_FMT_GBRAPF32 = AV_PIX_FMT_NE(GBRAPF32BE, GBRAPF32LE)
//const AV_PIX_FMT_GRAYF32 = AV_PIX_FMT_NE(GRAYF32BE, GRAYF32LE)
//const AV_PIX_FMT_YUVA420P9 = AV_PIX_FMT_NE(YUVA420P9BE , YUVA420P9LE)
//const AV_PIX_FMT_YUVA422P9 = AV_PIX_FMT_NE(YUVA422P9BE , YUVA422P9LE)
//const AV_PIX_FMT_YUVA444P9 = AV_PIX_FMT_NE(YUVA444P9BE , YUVA444P9LE)
//const AV_PIX_FMT_YUVA420P10 = AV_PIX_FMT_NE(YUVA420P10BE, YUVA420P10LE)
//const AV_PIX_FMT_YUVA422P10 = AV_PIX_FMT_NE(YUVA422P10BE, YUVA422P10LE)
//const AV_PIX_FMT_YUVA444P10 = AV_PIX_FMT_NE(YUVA444P10BE, YUVA444P10LE)
//const AV_PIX_FMT_YUVA422P12 = AV_PIX_FMT_NE(YUVA422P12BE, YUVA422P12LE)
//const AV_PIX_FMT_YUVA444P12 = AV_PIX_FMT_NE(YUVA444P12BE, YUVA444P12LE)
//const AV_PIX_FMT_YUVA420P16 = AV_PIX_FMT_NE(YUVA420P16BE, YUVA420P16LE)
//const AV_PIX_FMT_YUVA422P16 = AV_PIX_FMT_NE(YUVA422P16BE, YUVA422P16LE)
//const AV_PIX_FMT_YUVA444P16 = AV_PIX_FMT_NE(YUVA444P16BE, YUVA444P16LE)
//const AV_PIX_FMT_XYZ12 = AV_PIX_FMT_NE(XYZ12BE, XYZ12LE)
//const AV_PIX_FMT_NV20 = AV_PIX_FMT_NE(NV20BE,  NV20LE)
//const AV_PIX_FMT_AYUV64 = AV_PIX_FMT_NE(AYUV64BE, AYUV64LE)
//const AV_PIX_FMT_P010 = AV_PIX_FMT_NE(P010BE,  P010LE)
//const AV_PIX_FMT_P012 = AV_PIX_FMT_NE(P012BE,  P012LE)
//const AV_PIX_FMT_P016 = AV_PIX_FMT_NE(P016BE,  P016LE)
//const AV_PIX_FMT_Y210 = AV_PIX_FMT_NE(Y210BE,  Y210LE)
//const AV_PIX_FMT_Y212 = AV_PIX_FMT_NE(Y212BE,  Y212LE)
//const AV_PIX_FMT_XV30 = AV_PIX_FMT_NE(XV30BE,  XV30LE)
//const AV_PIX_FMT_XV36 = AV_PIX_FMT_NE(XV36BE,  XV36LE)
//const AV_PIX_FMT_X2RGB10 = AV_PIX_FMT_NE(X2RGB10BE, X2RGB10LE)
//const AV_PIX_FMT_X2BGR10 = AV_PIX_FMT_NE(X2BGR10BE, X2BGR10LE)
//const AV_PIX_FMT_P210 = AV_PIX_FMT_NE(P210BE, P210LE)
//const AV_PIX_FMT_P410 = AV_PIX_FMT_NE(P410BE, P410LE)
//const AV_PIX_FMT_P212 = AV_PIX_FMT_NE(P212BE, P212LE)
//const AV_PIX_FMT_P412 = AV_PIX_FMT_NE(P412BE, P412LE)
//const AV_PIX_FMT_P216 = AV_PIX_FMT_NE(P216BE, P216LE)
//const AV_PIX_FMT_P416 = AV_PIX_FMT_NE(P416BE, P416LE)
//const AV_PIX_FMT_RGBAF16 = AV_PIX_FMT_NE(RGBAF16BE, RGBAF16LE)
//const AV_PIX_FMT_RGBF32 = AV_PIX_FMT_NE(RGBF32BE, RGBF32LE)
//const AV_PIX_FMT_RGBAF32 = AV_PIX_FMT_NE(RGBAF32BE, RGBAF32LE)


                       
                       

/**
 * @file
 * pixel format definitions
 */

                               
                    




/**
 * Pixel format.
 *
 * @note
 * AV_PIX_FMT_RGB32 is handled in an endian-specific manner. An RGBA
 * color is put together as:
 *  (A << 24) | (R << 16) | (G << 8) | B
 * This is stored as BGRA on little-endian CPU architectures and ARGB on
 * big-endian CPUs.
 *
 * @note
 * If the resolution is not a multiple of the chroma subsampling factor
 * then the chroma plane resolution must be rounded up.
 *
 * @par
 * When the pixel format is palettized RGB32 (AV_PIX_FMT_PAL8), the palettized
 * image data is stored in AVFrame.data[0]. The palette is transported in
 * AVFrame.data[1], is 1024 bytes long (256 4-byte entries) and is
 * formatted the same as in AV_PIX_FMT_RGB32 described above (i.e., it is
 * also endian-specific). Note also that the individual RGB32 palette
 * components stored in AVFrame.data[1] should be in the range 0..255.
 * This is important as many custom PAL8 video codecs that were designed
 * to run on the IBM VGA graphics adapter use 6-bit palette components.
 *
 * @par
 * For all the 8 bits per pixel formats, an RGB32 palette is in data[1] like
 * for pal8. This palette is filled in automatically by the function
 * allocating the picture.
 */
type AVPixelFormat C.enum_AVPixelFormat

                     
                                                
     

      





































































































/**
  * Chromaticity coordinates of the source primaries.
  * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.1 and ITU-T H.273.
  */
type AVColorPrimaries C.enum_AVColorPrimaries

/**
 * Color Transfer Characteristic.
 * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.2.
 */
type AVColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

/**
 * YUV colorspace type.
 * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.3.
 */
type AVColorSpace C.enum_AVColorSpace

/**
 * Visual content value range.
 *
 * These values are based on definitions that can be found in multiple
 * specifications, such as ITU-T BT.709 (3.4 - Quantization of RGB, luminance
 * and colour-difference signals), ITU-T BT.2020 (Table 5 - Digital
 * Representation) as well as ITU-T BT.2100 (Table 9 - Digital 10- and 12-bit
 * integer representation). At the time of writing, the BT.2100 one is
 * recommended, as it also defines the full range representation.
 *
 * Common definitions:
 *   - For RGB and luma planes such as Y in YCbCr and I in ICtCp,
 *     'E' is the original value in range of 0.0 to 1.0.
 *   - For chroma planes such as Cb,Cr and Ct,Cp, 'E' is the original
 *     value in range of -0.5 to 0.5.
 *   - 'n' is the output bit depth.
 *   - For additional definitions such as rounding and clipping to valid n
 *     bit unsigned integer range, please refer to BT.2100 (Table 9).
 */
type AVColorRange C.enum_AVColorRange

/**
 * Location of chroma samples.
 *
 * Illustration showing the location of the first (top left) chroma sample of the
 * image, the left shows only luma, the right
 * shows the location of the chroma sample, the 2 could be imagined to overlay
 * each other but are drawn separately due to limitations of ASCII
 *
 *                1st 2nd       1st 2nd horizontal luma sample positions
 *                 v   v         v   v
 *                 ______        ______
 *1st luma line > |X   X ...    |3 4 X ...     X are luma samples,
 *                |             |1 2           1-6 are possible chroma positions
 *2nd luma line > |X   X ...    |5 6 X ...     0 is undefined/unknown position
 */
type AVChromaLocation C.enum_AVChromaLocation

                            
