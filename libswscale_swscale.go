// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2001-2011 Michael Niedermayer <michaelni@gmx.at>
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

//#cgo pkg-config: libswscale libavutil
//#include <stdint.h>
//#include "libavutil/avutil.h"
//#include "libavutil/frame.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libswscale/version_major.h"
//#include "libswscale/version.h"
//#include "libswscale/swscale.h"
import "C"
import (
    "unsafe"
)

const SWS_FAST_BILINEAR = 1
const SWS_BILINEAR = 2
const SWS_BICUBIC = 4
const SWS_X = 8
const SWS_POINT = 0x10
const SWS_AREA = 0x20
const SWS_BICUBLIN = 0x40
const SWS_GAUSS = 0x80
const SWS_SINC = 0x100
const SWS_LANCZOS = 0x200
const SWS_SPLINE = 0x400
const SWS_SRC_V_CHR_DROP_MASK = 0x30000
const SWS_SRC_V_CHR_DROP_SHIFT = 16
const SWS_PARAM_DEFAULT = 123456
const SWS_PRINT_INFO = 0x1000
const SWS_FULL_CHR_H_INT = 0x2000
const SWS_FULL_CHR_H_INP = 0x4000
const SWS_DIRECT_BGR = 0x8000
const SWS_ACCURATE_RND = 0x40000
const SWS_BITEXACT = 0x80000
const SWS_ERROR_DIFFUSION = 0x800000
const SWS_MAX_REDUCE_CUTOFF = 0.002
const SWS_CS_ITU709 = 1
const SWS_CS_FCC = 4
const SWS_CS_ITU601 = 5
const SWS_CS_ITU624 = 5
const SWS_CS_SMPTE170M = 5
const SWS_CS_SMPTE240M = 7
const SWS_CS_DEFAULT = 5
const SWS_CS_BT2020 = 9


                         
                         

/**
 * @file
 * @ingroup libsws
 * external API header
 */

                   

                             
                            
                          
                             
                          
                        
/* When included as part of the ffmpeg build, only include the major version
 * to avoid unnecessary rebuilds. When included externally, keep including
 * the full version information. */
                    
      

/**
 * @defgroup libsws libswscale
 * Color conversion and scaling library.
 *
 * @{
 *
 * Return the LIBSWSCALE_VERSION_INT constant.
 */
func Swscale_version() uint32 {
    return uint32(C.swscale_version())
}

/**
 * Return the libswscale build-time configuration.
 */
func Swscale_configuration() string {
    return C.GoString(C.swscale_configuration())
}

/**
 * Return the libswscale license.
 */
func Swscale_license() string {
    return C.GoString(C.swscale_license())
}

/* values for the flags, the stuff on the command line is different */



















//the following 3 flags are not completely implemented
//internal chrominance subsampling info

//input subsampling info

















/**
 * Return a pointer to yuv<->rgb coefficients for the given colorspace
 * suitable for sws_setColorspaceDetails().
 *
 * @param colorspace One of the SWS_CS_* macros. If invalid,
 * SWS_CS_DEFAULT is used.
 */
func Sws_getCoefficients(colorspace int32) *int32 {
    return (*int32)(unsafe.Pointer(C.sws_getCoefficients(C.int(colorspace))))
}

// when used for filters they must have an odd number of elements
// coeffs cannot be shared between vectors
type SwsVector C.struct_SwsVector

// vectors can be shared
type SwsFilter C.struct_SwsFilter

type SwsContext C.struct_SwsContext

/**
 * Return a positive value if pix_fmt is a supported input format, 0
 * otherwise.
 */
func Sws_isSupportedInput(pix_fmt AVPixelFormat) int32 {
    return int32(C.sws_isSupportedInput(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * Return a positive value if pix_fmt is a supported output format, 0
 * otherwise.
 */
func Sws_isSupportedOutput(pix_fmt AVPixelFormat) int32 {
    return int32(C.sws_isSupportedOutput(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * @param[in]  pix_fmt the pixel format
 * @return a positive value if an endianness conversion for pix_fmt is
 * supported, 0 otherwise.
 */
func Sws_isSupportedEndiannessConversion(pix_fmt AVPixelFormat) int32 {
    return int32(C.sws_isSupportedEndiannessConversion(
        C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * Allocate an empty SwsContext. This must be filled and passed to
 * sws_init_context(). For filling see AVOptions, options.c and
 * sws_setColorspaceDetails().
 */
func Sws_alloc_context() *SwsContext {
    return (*SwsContext)(unsafe.Pointer(C.sws_alloc_context()))
}

/**
 * Initialize the swscaler context sws_context.
 *
 * @return zero or positive value on success, a negative value on
 * error
 */
func Sws_init_context(sws_context *SwsContext, srcFilter *SwsFilter, dstFilter *SwsFilter) int32 {
    return int32(C.sws_init_context(
        (*C.struct_SwsContext)(unsafe.Pointer(sws_context)), 
        (*C.SwsFilter)(unsafe.Pointer(srcFilter)), 
        (*C.SwsFilter)(unsafe.Pointer(dstFilter))))
}

/**
 * Free the swscaler context swsContext.
 * If swsContext is NULL, then does nothing.
 */
func Sws_freeContext(swsContext *SwsContext)  {
    C.sws_freeContext((*C.struct_SwsContext)(unsafe.Pointer(swsContext)))
}

/**
 * Allocate and return an SwsContext. You need it to perform
 * scaling/conversion operations using sws_scale().
 *
 * @param srcW the width of the source image
 * @param srcH the height of the source image
 * @param srcFormat the source image format
 * @param dstW the width of the destination image
 * @param dstH the height of the destination image
 * @param dstFormat the destination image format
 * @param flags specify which algorithm and options to use for rescaling
 * @param param extra parameters to tune the used scaler
 *              For SWS_BICUBIC param[0] and [1] tune the shape of the basis
 *              function, param[0] tunes f(1) and param[1] f´(1)
 *              For SWS_GAUSS param[0] tunes the exponent and thus cutoff
 *              frequency
 *              For SWS_LANCZOS param[0] tunes the width of the window function
 * @return a pointer to an allocated context, or NULL in case of error
 * @note this function is to be removed after a saner alternative is
 *       written
 */
func Sws_getContext(srcW int32, srcH int32, srcFormat AVPixelFormat,
                                  dstW int32, dstH int32, dstFormat AVPixelFormat,
                                  flags int32, srcFilter *SwsFilter,
                                  dstFilter *SwsFilter, param *float64) *SwsContext {
    return (*SwsContext)(unsafe.Pointer(C.sws_getContext(C.int(srcW), C.int(srcH), 
        C.enum_AVPixelFormat(srcFormat), C.int(dstW), C.int(dstH), 
        C.enum_AVPixelFormat(dstFormat), C.int(flags), 
        (*C.SwsFilter)(unsafe.Pointer(srcFilter)), 
        (*C.SwsFilter)(unsafe.Pointer(dstFilter)), 
        (*C.double)(unsafe.Pointer(param)))))
}

/**
 * Scale the image slice in srcSlice and put the resulting scaled
 * slice in the image in dst. A slice is a sequence of consecutive
 * rows in an image.
 *
 * Slices have to be provided in sequential order, either in
 * top-bottom or bottom-top order. If slices are provided in
 * non-sequential order the behavior of the function is undefined.
 *
 * @param c         the scaling context previously created with
 *                  sws_getContext()
 * @param srcSlice  the array containing the pointers to the planes of
 *                  the source slice
 * @param srcStride the array containing the strides for each plane of
 *                  the source image
 * @param srcSliceY the position in the source image of the slice to
 *                  process, that is the number (counted starting from
 *                  zero) in the image of the first row of the slice
 * @param srcSliceH the height of the source slice, that is the number
 *                  of rows in the slice
 * @param dst       the array containing the pointers to the planes of
 *                  the destination image
 * @param dstStride the array containing the strides for each plane of
 *                  the destination image
 * @return          the height of the output slice
 */
func Sws_scale(c *SwsContext, srcSlice[] *uint8,
              srcStride[] int32, srcSliceY int32, srcSliceH int32,
              dst[] *uint8, dstStride[] int32) int32 {
    return int32(C.sws_scale((*C.struct_SwsContext)(unsafe.Pointer(c)), 
        (**C.uchar)(unsafe.Pointer(&srcSlice[0])), 
        (*C.int)(unsafe.Pointer(&srcStride[0])), C.int(srcSliceY), 
        C.int(srcSliceH), (**C.uchar)(unsafe.Pointer(&dst[0])), 
        (*C.int)(unsafe.Pointer(&dstStride[0]))))
}

/**
 * Scale source data from src and write the output to dst.
 *
 * This is merely a convenience wrapper around
 * - sws_frame_start()
 * - sws_send_slice(0, src->height)
 * - sws_receive_slice(0, dst->height)
 * - sws_frame_end()
 *
 * @param c   The scaling context
 * @param dst The destination frame. See documentation for sws_frame_start() for
 *            more details.
 * @param src The source frame.
 *
 * @return 0 on success, a negative AVERROR code on failure
 */
func Sws_scale_frame(c *SwsContext, dst *AVFrame, src *AVFrame) int32 {
    return int32(C.sws_scale_frame((*C.struct_SwsContext)(unsafe.Pointer(c)), 
        (*C.AVFrame)(unsafe.Pointer(dst)), (*C.AVFrame)(unsafe.Pointer(src))))
}

/**
 * Initialize the scaling process for a given pair of source/destination frames.
 * Must be called before any calls to sws_send_slice() and sws_receive_slice().
 *
 * This function will retain references to src and dst, so they must both use
 * refcounted buffers (if allocated by the caller, in case of dst).
 *
 * @param c   The scaling context
 * @param dst The destination frame.
 *
 *            The data buffers may either be already allocated by the caller or
 *            left clear, in which case they will be allocated by the scaler.
 *            The latter may have performance advantages - e.g. in certain cases
 *            some output planes may be references to input planes, rather than
 *            copies.
 *
 *            Output data will be written into this frame in successful
 *            sws_receive_slice() calls.
 * @param src The source frame. The data buffers must be allocated, but the
 *            frame data does not have to be ready at this point. Data
 *            availability is then signalled by sws_send_slice().
 * @return 0 on success, a negative AVERROR code on failure
 *
 * @see sws_frame_end()
 */
func Sws_frame_start(c *SwsContext, dst *AVFrame, src *AVFrame) int32 {
    return int32(C.sws_frame_start((*C.struct_SwsContext)(unsafe.Pointer(c)), 
        (*C.AVFrame)(unsafe.Pointer(dst)), (*C.AVFrame)(unsafe.Pointer(src))))
}

/**
 * Finish the scaling process for a pair of source/destination frames previously
 * submitted with sws_frame_start(). Must be called after all sws_send_slice()
 * and sws_receive_slice() calls are done, before any new sws_frame_start()
 * calls.
 *
 * @param c   The scaling context
 */
func Sws_frame_end(c *SwsContext)  {
    C.sws_frame_end((*C.struct_SwsContext)(unsafe.Pointer(c)))
}

/**
 * Indicate that a horizontal slice of input data is available in the source
 * frame previously provided to sws_frame_start(). The slices may be provided in
 * any order, but may not overlap. For vertically subsampled pixel formats, the
 * slices must be aligned according to subsampling.
 *
 * @param c   The scaling context
 * @param slice_start first row of the slice
 * @param slice_height number of rows in the slice
 *
 * @return a non-negative number on success, a negative AVERROR code on failure.
 */
func Sws_send_slice(c *SwsContext, slice_start uint32,
                   slice_height uint32) int32 {
    return int32(C.sws_send_slice((*C.struct_SwsContext)(unsafe.Pointer(c)), 
        C.uint(slice_start), C.uint(slice_height)))
}

/**
 * Request a horizontal slice of the output data to be written into the frame
 * previously provided to sws_frame_start().
 *
 * @param c   The scaling context
 * @param slice_start first row of the slice; must be a multiple of
 *                    sws_receive_slice_alignment()
 * @param slice_height number of rows in the slice; must be a multiple of
 *                     sws_receive_slice_alignment(), except for the last slice
 *                     (i.e. when slice_start+slice_height is equal to output
 *                     frame height)
 *
 * @return a non-negative number if the data was successfully written into the output
 *         AVERROR(EAGAIN) if more input data needs to be provided before the
 *                         output can be produced
 *         another negative AVERROR code on other kinds of scaling failure
 */
func Sws_receive_slice(c *SwsContext, slice_start uint32,
                      slice_height uint32) int32 {
    return int32(C.sws_receive_slice((*C.struct_SwsContext)(unsafe.Pointer(c)), 
        C.uint(slice_start), C.uint(slice_height)))
}

/**
 * Get the alignment required for slices
 *
 * @param c   The scaling context
 * @return alignment required for output slices requested with sws_receive_slice().
 *         Slice offsets and sizes passed to sws_receive_slice() must be
 *         multiples of the value returned from this function.
 */
func Sws_receive_slice_alignment(c *SwsContext) uint32 {
    return uint32(C.sws_receive_slice_alignment(
        (*C.struct_SwsContext)(unsafe.Pointer(c))))
}

/**
 * @param c the scaling context
 * @param dstRange flag indicating the while-black range of the output (1=jpeg / 0=mpeg)
 * @param srcRange flag indicating the while-black range of the input (1=jpeg / 0=mpeg)
 * @param table the yuv2rgb coefficients describing the output yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param inv_table the yuv2rgb coefficients describing the input yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param brightness 16.16 fixed point brightness correction
 * @param contrast 16.16 fixed point contrast correction
 * @param saturation 16.16 fixed point saturation correction
 *
 * @return A negative error code on error, non negative otherwise.
 *         If `LIBSWSCALE_VERSION_MAJOR < 7`, returns -1 if not supported.
 */
func Sws_setColorspaceDetails(c *SwsContext, inv_table[4] int32,
                             srcRange int32, table[4] int32, dstRange int32,
                             brightness int32, contrast int32, saturation int32) int32 {
    return int32(C.sws_setColorspaceDetails(
        (*C.struct_SwsContext)(unsafe.Pointer(c)), 
        (*C.int)(unsafe.Pointer(&inv_table[0])), C.int(srcRange), 
        (*C.int)(unsafe.Pointer(&table[0])), C.int(dstRange), C.int(brightness), 
        C.int(contrast), C.int(saturation)))
}

/**
 * @return A negative error code on error, non negative otherwise.
 *         If `LIBSWSCALE_VERSION_MAJOR < 7`, returns -1 if not supported.
 */
func Sws_getColorspaceDetails(c *SwsContext, inv_table **int32,
                             srcRange *int32, table **int32, dstRange *int32,
                             brightness *int32, contrast *int32, saturation *int32) int32 {
    return int32(C.sws_getColorspaceDetails(
        (*C.struct_SwsContext)(unsafe.Pointer(c)), 
        (**C.int)(unsafe.Pointer(inv_table)), (*C.int)(unsafe.Pointer(srcRange)), 
        (**C.int)(unsafe.Pointer(table)), (*C.int)(unsafe.Pointer(dstRange)), 
        (*C.int)(unsafe.Pointer(brightness)), (*C.int)(unsafe.Pointer(contrast)), 
        (*C.int)(unsafe.Pointer(saturation))))
}

/**
 * Allocate and return an uninitialized vector with length coefficients.
 */
func Sws_allocVec(length int32) *SwsVector {
    return (*SwsVector)(unsafe.Pointer(C.sws_allocVec(C.int(length))))
}

/**
 * Return a normalized Gaussian curve used to filter stuff
 * quality = 3 is high quality, lower is lower quality.
 */
func Sws_getGaussianVec(variance float64, quality float64) *SwsVector {
    return (*SwsVector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(variance), C.double(quality))))
}

/**
 * Scale all the coefficients of a by the scalar value.
 */
func Sws_scaleVec(a *SwsVector, scalar float64)  {
    C.sws_scaleVec((*C.SwsVector)(unsafe.Pointer(a)), C.double(scalar))
}

/**
 * Scale all the coefficients of a so that their sum equals height.
 */
func Sws_normalizeVec(a *SwsVector, height float64)  {
    C.sws_normalizeVec((*C.SwsVector)(unsafe.Pointer(a)), C.double(height))
}

func Sws_freeVec(a *SwsVector)  {
    C.sws_freeVec((*C.SwsVector)(unsafe.Pointer(a)))
}

func Sws_getDefaultFilter(lumaGBlur float32, chromaGBlur float32,
                                lumaSharpen float32, chromaSharpen float32,
                                chromaHShift float32, chromaVShift float32,
                                verbose int32) *SwsFilter {
    return (*SwsFilter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lumaGBlur), 
        C.float(chromaGBlur), C.float(lumaSharpen), C.float(chromaSharpen), 
        C.float(chromaHShift), C.float(chromaVShift), C.int(verbose))))
}
func Sws_freeFilter(filter *SwsFilter)  {
    C.sws_freeFilter((*C.SwsFilter)(unsafe.Pointer(filter)))
}

/**
 * Check if context can be reused, otherwise reallocate a new one.
 *
 * If context is NULL, just calls sws_getContext() to get a new
 * context. Otherwise, checks if the parameters are the ones already
 * saved in context. If that is the case, returns the current
 * context. Otherwise, frees context and gets a new context with
 * the new parameters.
 *
 * Be warned that srcFilter and dstFilter are not checked, they
 * are assumed to remain the same.
 */
func Sws_getCachedContext(context *SwsContext,
                                        srcW int32, srcH int32, srcFormat AVPixelFormat,
                                        dstW int32, dstH int32, dstFormat AVPixelFormat,
                                        flags int32, srcFilter *SwsFilter,
                                        dstFilter *SwsFilter, param *float64) *SwsContext {
    return (*SwsContext)(unsafe.Pointer(C.sws_getCachedContext(
        (*C.struct_SwsContext)(unsafe.Pointer(context)), C.int(srcW), 
        C.int(srcH), C.enum_AVPixelFormat(srcFormat), C.int(dstW), C.int(dstH), 
        C.enum_AVPixelFormat(dstFormat), C.int(flags), 
        (*C.SwsFilter)(unsafe.Pointer(srcFilter)), 
        (*C.SwsFilter)(unsafe.Pointer(dstFilter)), 
        (*C.double)(unsafe.Pointer(param)))))
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
 *
 * The output frame will have the same packed format as the palette.
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
func Sws_convertPalette8ToPacked32(src *uint8, dst *uint8, num_pixels int32, palette *uint8)  {
    C.sws_convertPalette8ToPacked32((*C.uchar)(unsafe.Pointer(src)), 
        (*C.uchar)(unsafe.Pointer(dst)), C.int(num_pixels), 
        (*C.uchar)(unsafe.Pointer(palette)))
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
 *
 * With the palette format "ABCD", the destination frame ends up with the format "ABC".
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
func Sws_convertPalette8ToPacked24(src *uint8, dst *uint8, num_pixels int32, palette *uint8)  {
    C.sws_convertPalette8ToPacked24((*C.uchar)(unsafe.Pointer(src)), 
        (*C.uchar)(unsafe.Pointer(dst)), C.int(num_pixels), 
        (*C.uchar)(unsafe.Pointer(palette)))
}

/**
 * Get the AVClass for swsContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Sws_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.sws_get_class()))
}

/**
 * @}
 */

                              
