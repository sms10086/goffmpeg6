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
//#include "libavutil/pixdesc.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavutil/imgutils.h"
import "C"
import (
    "unsafe"
)



                         
                         

/**
 * @file
 * misc image utilities
 *
 * @addtogroup lavu_picture
 * @{
 */

                   
                   
                    
                   
                     

/**
 * Compute the max pixel step for each plane of an image with a
 * format described by pixdesc.
 *
 * The pixel step is the distance in bytes between the first byte of
 * the group of bytes which describe a pixel component and the first
 * byte of the successive group in the same plane for the same
 * component.
 *
 * @param max_pixsteps an array which is filled with the max pixel step
 * for each plane. Since a plane may contain different pixel
 * components, the computed max_pixsteps[plane] is relative to the
 * component in the plane with the max pixel step.
 * @param max_pixstep_comps an array which is filled with the component
 * for each plane which has the max pixel step. May be NULL.
 * @param pixdesc the AVPixFmtDescriptor for the image, describing its format
 */
func Av_image_fill_max_pixsteps(max_pixsteps [4]int32, max_pixstep_comps [4]int32,
                                pixdesc *AVPixFmtDescriptor)  {
    C.av_image_fill_max_pixsteps((*C.int)(unsafe.Pointer(&max_pixsteps[0])), 
        (*C.int)(unsafe.Pointer(&max_pixstep_comps[0])), 
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(pixdesc)))
}

/**
 * Compute the size of an image line with format pix_fmt and width
 * width for the plane plane.
 *
 * @return the computed size in bytes
 */
func Av_image_get_linesize(pix_fmt AVPixelFormat, width int32, plane int32) int32 {
    return int32(C.av_image_get_linesize(C.enum_AVPixelFormat(pix_fmt), C.int(width), 
        C.int(plane)))
}

/**
 * Fill plane linesizes for an image with pixel format pix_fmt and
 * width width.
 *
 * @param linesizes array to be filled with the linesize for each plane
 * @param pix_fmt the AVPixelFormat of the image
 * @param width width of the image in pixels
 * @return >= 0 in case of success, a negative error code otherwise
 */
func Av_image_fill_linesizes(linesizes [4]int32, pix_fmt AVPixelFormat, width int32) int32 {
    return int32(C.av_image_fill_linesizes((*C.int)(unsafe.Pointer(&linesizes[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.int(width)))
}

/**
 * Fill plane sizes for an image with pixel format pix_fmt and height height.
 *
 * @param size the array to be filled with the size of each image plane
 * @param pix_fmt the AVPixelFormat of the image
 * @param height height of the image in pixels
 * @param linesizes the array containing the linesize for each
 *        plane, should be filled by av_image_fill_linesizes()
 * @return >= 0 in case of success, a negative error code otherwise
 *
 * @note The linesize parameters have the type ptrdiff_t here, while they are
 *       int for av_image_fill_linesizes().
 */
func Av_image_fill_plane_sizes(size [4]uint64, pix_fmt AVPixelFormat,
                              height int32, linesizes [4]int32) int32 {
    return int32(C.av_image_fill_plane_sizes(
        (*C.ulonglong)(unsafe.Pointer(&size[0])), C.enum_AVPixelFormat(pix_fmt), 
        C.int(height), (*C.ptrdiff_t)(unsafe.Pointer(&linesizes[0]))))
}

/**
 * Fill plane data pointers for an image with pixel format pix_fmt and
 * height height.
 *
 * @param data pointers array to be filled with the pointer for each image plane
 * @param pix_fmt the AVPixelFormat of the image
 * @param height height of the image in pixels
 * @param ptr the pointer to a buffer which will contain the image
 * @param linesizes the array containing the linesize for each
 * plane, should be filled by av_image_fill_linesizes()
 * @return the size in bytes required for the image buffer, a negative
 * error code in case of failure
 */
func Av_image_fill_pointers(data [4]*uint8, pix_fmt AVPixelFormat, height int32,
                           ptr *uint8, linesizes [4]int32) int32 {
    return int32(C.av_image_fill_pointers((**C.uchar)(unsafe.Pointer(&data[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.int(height), 
        (*C.uchar)(unsafe.Pointer(ptr)), (*C.int)(unsafe.Pointer(&linesizes[0]))))
}

/**
 * Allocate an image with size w and h and pixel format pix_fmt, and
 * fill pointers and linesizes accordingly.
 * The allocated image buffer has to be freed by using
 * av_freep(&pointers[0]).
 *
 * @param pointers array to be filled with the pointer for each image plane
 * @param linesizes the array filled with the linesize for each plane
 * @param w width of the image in pixels
 * @param h height of the image in pixels
 * @param pix_fmt the AVPixelFormat of the image
 * @param align the value to use for buffer size alignment
 * @return the size in bytes required for the image buffer, a negative
 * error code in case of failure
 */
func Av_image_alloc(pointers [4]*uint8, linesizes [4]int32,
                   w int32, h int32, pix_fmt AVPixelFormat, align int32) int32 {
    return int32(C.av_image_alloc((**C.uchar)(unsafe.Pointer(&pointers[0])), 
        (*C.int)(unsafe.Pointer(&linesizes[0])), C.int(w), C.int(h), 
        C.enum_AVPixelFormat(pix_fmt), C.int(align)))
}

/**
 * Copy image plane from src to dst.
 * That is, copy "height" number of lines of "bytewidth" bytes each.
 * The first byte of each successive line is separated by *_linesize
 * bytes.
 *
 * bytewidth must be contained by both absolute values of dst_linesize
 * and src_linesize, otherwise the function behavior is undefined.
 *
 * @param dst          destination plane to copy to
 * @param dst_linesize linesize for the image plane in dst
 * @param src          source plane to copy from
 * @param src_linesize linesize for the image plane in src
 * @param height       height (number of lines) of the plane
 */
func Av_image_copy_plane(dst *uint8, dst_linesize int32,
                         src *uint8, src_linesize int32,
                         bytewidth int32, height int32)  {
    C.av_image_copy_plane((*C.uchar)(unsafe.Pointer(dst)), C.int(dst_linesize), 
        (*C.uchar)(unsafe.Pointer(src)), C.int(src_linesize), C.int(bytewidth), 
        C.int(height))
}

/**
 * Copy image data located in uncacheable (e.g. GPU mapped) memory. Where
 * available, this function will use special functionality for reading from such
 * memory, which may result in greatly improved performance compared to plain
 * av_image_copy_plane().
 *
 * bytewidth must be contained by both absolute values of dst_linesize
 * and src_linesize, otherwise the function behavior is undefined.
 *
 * @note The linesize parameters have the type ptrdiff_t here, while they are
 *       int for av_image_copy_plane().
 * @note On x86, the linesizes currently need to be aligned to the cacheline
 *       size (i.e. 64) to get improved performance.
 */
func Av_image_copy_plane_uc_from(dst *uint8, dst_linesize int32,
                                 src *uint8, src_linesize int32,
                                 bytewidth int32, height int32)  {
    C.av_image_copy_plane_uc_from((*C.uchar)(unsafe.Pointer(dst)), 
        C.ptrdiff_t(dst_linesize), (*C.uchar)(unsafe.Pointer(src)), 
        C.ptrdiff_t(src_linesize), C.ptrdiff_t(bytewidth), C.int(height))
}

/**
 * Copy image in src_data to dst_data.
 *
 * @param dst_data      destination image data buffer to copy to
 * @param dst_linesizes linesizes for the image in dst_data
 * @param src_data      source image data buffer to copy from
 * @param src_linesizes linesizes for the image in src_data
 * @param pix_fmt       the AVPixelFormat of the image
 * @param width         width of the image in pixels
 * @param height        height of the image in pixels
 */
func Av_image_copy(dst_data [4]*uint8, dst_linesizes [4]int32,
                   src_data [4]*uint8, src_linesizes [4]int32,
                   pix_fmt AVPixelFormat, width int32, height int32)  {
    C.av_image_copy((**C.uchar)(unsafe.Pointer(&dst_data[0])), 
        (*C.int)(unsafe.Pointer(&dst_linesizes[0])), 
        (**C.uchar)(unsafe.Pointer(&src_data[0])), 
        (*C.int)(unsafe.Pointer(&src_linesizes[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.int(width), C.int(height))
}

/**
 * Wrapper around av_image_copy() to workaround the limitation
 * that the conversion from uint8_t * const * to const uint8_t * const *
 * is not performed automatically in C.
 * @see av_image_copy()
 */
// av_image_copy2(uint8_t*constdst_data[4],constintdst_linesizes[4],uint8_t*constsrc_data[4],constintsrc_linesizes[4],enumAVPixelFormatpix_fmt,intwidth,intheight)

/**
 * Copy image data located in uncacheable (e.g. GPU mapped) memory. Where
 * available, this function will use special functionality for reading from such
 * memory, which may result in greatly improved performance compared to plain
 * av_image_copy().
 *
 * The data pointers and the linesizes must be aligned to the maximum required
 * by the CPU architecture.
 *
 * @note The linesize parameters have the type ptrdiff_t here, while they are
 *       int for av_image_copy().
 * @note On x86, the linesizes currently need to be aligned to the cacheline
 *       size (i.e. 64) to get improved performance.
 */
func Av_image_copy_uc_from(dst_data [4]*uint8,       dst_linesizes [4]int32,
                           src_data [4]*uint8, src_linesizes [4]int32,
                           pix_fmt AVPixelFormat, width int32, height int32)  {
    C.av_image_copy_uc_from((**C.uchar)(unsafe.Pointer(&dst_data[0])), 
        (*C.ptrdiff_t)(unsafe.Pointer(&dst_linesizes[0])), 
        (**C.uchar)(unsafe.Pointer(&src_data[0])), 
        (*C.ptrdiff_t)(unsafe.Pointer(&src_linesizes[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.int(width), C.int(height))
}

/**
 * Setup the data pointers and linesizes based on the specified image
 * parameters and the provided array.
 *
 * The fields of the given image are filled in by using the src
 * address which points to the image data buffer. Depending on the
 * specified pixel format, one or multiple image data pointers and
 * line sizes will be set.  If a planar format is specified, several
 * pointers will be set pointing to the different picture planes and
 * the line sizes of the different planes will be stored in the
 * lines_sizes array. Call with src == NULL to get the required
 * size for the src buffer.
 *
 * To allocate the buffer and fill in the dst_data and dst_linesize in
 * one call, use av_image_alloc().
 *
 * @param dst_data      data pointers to be filled in
 * @param dst_linesize  linesizes for the image in dst_data to be filled in
 * @param src           buffer which will contain or contains the actual image data, can be NULL
 * @param pix_fmt       the pixel format of the image
 * @param width         the width of the image in pixels
 * @param height        the height of the image in pixels
 * @param align         the value used in src for linesize alignment
 * @return the size in bytes required for src, a negative error code
 * in case of failure
 */
func Av_image_fill_arrays(dst_data [4]*uint8, dst_linesize [4]int32,
                         src *uint8,
                         pix_fmt AVPixelFormat, width int32, height int32, align int32) int32 {
    return int32(C.av_image_fill_arrays((**C.uchar)(unsafe.Pointer(&dst_data[0])), 
        (*C.int)(unsafe.Pointer(&dst_linesize[0])), 
        (*C.uchar)(unsafe.Pointer(src)), C.enum_AVPixelFormat(pix_fmt), 
        C.int(width), C.int(height), C.int(align)))
}

/**
 * Return the size in bytes of the amount of data required to store an
 * image with the given parameters.
 *
 * @param pix_fmt  the pixel format of the image
 * @param width    the width of the image in pixels
 * @param height   the height of the image in pixels
 * @param align    the assumed linesize alignment
 * @return the buffer size in bytes, a negative error code in case of failure
 */
func Av_image_get_buffer_size(pix_fmt AVPixelFormat, width int32, height int32, align int32) int32 {
    return int32(C.av_image_get_buffer_size(C.enum_AVPixelFormat(pix_fmt), 
        C.int(width), C.int(height), C.int(align)))
}

/**
 * Copy image data from an image into a buffer.
 *
 * av_image_get_buffer_size() can be used to compute the required size
 * for the buffer to fill.
 *
 * @param dst           a buffer into which picture data will be copied
 * @param dst_size      the size in bytes of dst
 * @param src_data      pointers containing the source image data
 * @param src_linesize  linesizes for the image in src_data
 * @param pix_fmt       the pixel format of the source image
 * @param width         the width of the source image in pixels
 * @param height        the height of the source image in pixels
 * @param align         the assumed linesize alignment for dst
 * @return the number of bytes written to dst, or a negative value
 * (error code) on error
 */
func Av_image_copy_to_buffer(dst *uint8, dst_size int32,
                            src_data [4]*uint8, src_linesize [4]int32,
                            pix_fmt AVPixelFormat, width int32, height int32, align int32) int32 {
    return int32(C.av_image_copy_to_buffer((*C.uchar)(unsafe.Pointer(dst)), 
        C.int(dst_size), (**C.uchar)(unsafe.Pointer(&src_data[0])), 
        (*C.int)(unsafe.Pointer(&src_linesize[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.int(width), C.int(height), C.int(align)))
}

/**
 * Check if the given dimension of an image is valid, meaning that all
 * bytes of the image can be addressed with a signed int.
 *
 * @param w the width of the picture
 * @param h the height of the picture
 * @param log_offset the offset to sum to the log level for logging with log_ctx
 * @param log_ctx the parent logging context, it may be NULL
 * @return >= 0 if valid, a negative error code otherwise
 */
func Av_image_check_size(w uint32, h uint32, log_offset int32, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_image_check_size(C.uint(w), C.uint(h), C.int(log_offset), 
        log_ctx))
}

/**
 * Check if the given dimension of an image is valid, meaning that all
 * bytes of a plane of an image with the specified pix_fmt can be addressed
 * with a signed int.
 *
 * @param w the width of the picture
 * @param h the height of the picture
 * @param max_pixels the maximum number of pixels the user wants to accept
 * @param pix_fmt the pixel format, can be AV_PIX_FMT_NONE if unknown.
 * @param log_offset the offset to sum to the log level for logging with log_ctx
 * @param log_ctx the parent logging context, it may be NULL
 * @return >= 0 if valid, a negative error code otherwise
 */
func Av_image_check_size2(w uint32, h uint32, max_pixels int64, pix_fmt AVPixelFormat, log_offset int32, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_image_check_size2(C.uint(w), C.uint(h), C.longlong(max_pixels), 
        C.enum_AVPixelFormat(pix_fmt), C.int(log_offset), log_ctx))
}

/**
 * Check if the given sample aspect ratio of an image is valid.
 *
 * It is considered invalid if the denominator is 0 or if applying the ratio
 * to the image size would make the smaller dimension less than 1. If the
 * sar numerator is 0, it is considered unknown and will return as valid.
 *
 * @param w width of the image
 * @param h height of the image
 * @param sar sample aspect ratio of the image
 * @return 0 if valid, a negative AVERROR code otherwise
 */
func Av_image_check_sar(w uint32, h uint32, sar AVRational) int32 {
    return int32(C.av_image_check_sar(C.uint(w), C.uint(h), 
        *(*C.struct_AVRational)(unsafe.Pointer(&sar))))
}

/**
 * Overwrite the image data with black. This is suitable for filling a
 * sub-rectangle of an image, meaning the padding between the right most pixel
 * and the left most pixel on the next line will not be overwritten. For some
 * formats, the image size might be rounded up due to inherent alignment.
 *
 * If the pixel format has alpha, the alpha is cleared to opaque.
 *
 * This can return an error if the pixel format is not supported. Normally, all
 * non-hwaccel pixel formats should be supported.
 *
 * Passing NULL for dst_data is allowed. Then the function returns whether the
 * operation would have succeeded. (It can return an error if the pix_fmt is
 * not supported.)
 *
 * @param dst_data      data pointers to destination image
 * @param dst_linesize  linesizes for the destination image
 * @param pix_fmt       the pixel format of the image
 * @param range         the color range of the image (important for colorspaces such as YUV)
 * @param width         the width of the image in pixels
 * @param height        the height of the image in pixels
 * @return 0 if the image data was cleared, a negative AVERROR code otherwise
 */
func Av_image_fill_black(dst_data [4]*uint8, dst_linesize [4]int32,
                        pix_fmt AVPixelFormat, rangex AVColorRange,
                        width int32, height int32) int32 {
    return int32(C.av_image_fill_black((**C.uchar)(unsafe.Pointer(&dst_data[0])), 
        (*C.ptrdiff_t)(unsafe.Pointer(&dst_linesize[0])), 
        C.enum_AVPixelFormat(pix_fmt), C.enum_AVColorRange(rangex), C.int(width), 
        C.int(height)))
}

/**
 * @}
 */


                              
