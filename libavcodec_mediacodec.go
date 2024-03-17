// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Android MediaCodec public API
 *
 * Copyright (c) 2016 Matthieu Bouron <matthieu.bouron stupeflix.com>
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

//#cgo pkg-config: libavcodec
//#include "libavcodec/avcodec.h"
//#include "libavcodec/mediacodec.h"
import "C"
import (
    "unsafe"
)



                            
                            

                               

/**
 * This structure holds a reference to a android/view/Surface object that will
 * be used as output by the decoder.
 *
 */
type AVMediaCodecContext struct {
    Surface unsafe.Pointer
}


/**
 * Allocate and initialize a MediaCodec context.
 *
 * When decoding with MediaCodec is finished, the caller must free the
 * MediaCodec context with av_mediacodec_default_free.
 *
 * @return a pointer to a newly allocated AVMediaCodecContext on success, NULL otherwise
 */
func Av_mediacodec_alloc_context() *AVMediaCodecContext {
    return (*AVMediaCodecContext)(unsafe.Pointer(C.av_mediacodec_alloc_context()))
}

/**
 * Convenience function that sets up the MediaCodec context.
 *
 * @param avctx codec context
 * @param ctx MediaCodec context to initialize
 * @param surface reference to an android/view/Surface
 * @return 0 on success, < 0 otherwise
 */
func Av_mediacodec_default_init(avctx *AVCodecContext, ctx *AVMediaCodecContext, surface unsafe.Pointer) int32 {
    return int32(C.av_mediacodec_default_init(
        (*C.struct_AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.struct_AVMediaCodecContext)(unsafe.Pointer(ctx)), surface))
}

/**
 * This function must be called to free the MediaCodec context initialized with
 * av_mediacodec_default_init().
 *
 * @param avctx codec context
 */
func Av_mediacodec_default_free(avctx *AVCodecContext)  {
    C.av_mediacodec_default_free((*C.struct_AVCodecContext)(unsafe.Pointer(avctx)))
}

/**
 * Opaque structure representing a MediaCodec buffer to render.
 */
type AVMediaCodecBuffer struct {
}


/**
 * Release a MediaCodec buffer and render it to the surface that is associated
 * with the decoder. This function should only be called once on a given
 * buffer, once released the underlying buffer returns to the codec, thus
 * subsequent calls to this function will have no effect.
 *
 * @param buffer the buffer to render
 * @param render 1 to release and render the buffer to the surface or 0 to
 * discard the buffer
 * @return 0 on success, < 0 otherwise
 */
func Av_mediacodec_release_buffer(buffer *AVMediaCodecBuffer, render int32) int32 {
    return int32(C.av_mediacodec_release_buffer(
        (*C.AVMediaCodecBuffer)(unsafe.Pointer(buffer)), C.int(render)))
}

/**
 * Release a MediaCodec buffer and render it at the given time to the surface
 * that is associated with the decoder. The timestamp must be within one second
 * of the current `java/lang/System#nanoTime()` (which is implemented using
 * `CLOCK_MONOTONIC` on Android). See the Android MediaCodec documentation
 * of [`android/media/MediaCodec#releaseOutputBuffer(int,long)`][0] for more details.
 *
 * @param buffer the buffer to render
 * @param time timestamp in nanoseconds of when to render the buffer
 * @return 0 on success, < 0 otherwise
 *
 * [0]: https://developer.android.com/reference/android/media/MediaCodec#releaseOutputBuffer(int,%20long)
 */
func Av_mediacodec_render_buffer_at_time(buffer *AVMediaCodecBuffer, time int64) int32 {
    return int32(C.av_mediacodec_render_buffer_at_time(
        (*C.AVMediaCodecBuffer)(unsafe.Pointer(buffer)), C.longlong(time)))
}

                                 
