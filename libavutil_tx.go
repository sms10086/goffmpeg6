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
//#include <stdint.h>
//#include <stddef.h>
//#include "libavutil/tx.h"
import "C"
import (
    "unsafe"
)



                   
                   

                   
                   

type AVTXContext struct {
}


type AVComplexFloat struct {
    Re float32
    Im float32
}


type AVComplexDouble struct {
    Re float64
    Im float64
}


type AVComplexInt32 struct {
    Re int32
    Im int32
}


type AVTXType int32
const (
    AV_TX_FLOAT_FFT AVTXType = 0 + iota
    AV_TX_DOUBLE_FFT = 2
    AV_TX_INT32_FFT = 4
    AV_TX_FLOAT_MDCT = 1
    AV_TX_DOUBLE_MDCT = 3
    AV_TX_INT32_MDCT = 5
    AV_TX_FLOAT_RDFT = 6
    AV_TX_DOUBLE_RDFT = 7
    AV_TX_INT32_RDFT = 8
    AV_TX_FLOAT_DCT = 9
    AV_TX_DOUBLE_DCT = 10
    AV_TX_INT32_DCT = 11
    AV_TX_FLOAT_DCT_I = 12
    AV_TX_DOUBLE_DCT_I = 13
    AV_TX_INT32_DCT_I = 14
    AV_TX_FLOAT_DST_I = 15
    AV_TX_DOUBLE_DST_I = 16
    AV_TX_INT32_DST_I = 17
    AV_TX_NB = 17 + iota - 17
)


/**
 * Function pointer to a function to perform the transform.
 *
 * @note Using a different context than the one allocated during av_tx_init()
 * is not allowed.
 *
 * @param s the transform context
 * @param out the output array
 * @param in the input array
 * @param stride the input or output stride in bytes
 *
 * The out and in arrays must be aligned to the maximum required by the CPU
 * architecture unless the AV_TX_UNALIGNED flag was set in av_tx_init().
 * The stride must follow the constraints the transform type has specified.
 */
type av_tx_fn C.av_tx_fn

/**
 * Flags for av_tx_init()
 */
type AVTXFlags int32
const (
    AV_TX_INPLACE AVTXFlags = 1<<0 + iota
    AV_TX_UNALIGNED = 1<<1
    AV_TX_FULL_IMDCT = 1<<2
    AV_TX_REAL_TO_REAL = 1<<3
    AV_TX_REAL_TO_IMAGINARY = 1<<4
)


/**
 * Initialize a transform context with the given configuration
 * (i)MDCTs with an odd length are currently not supported.
 *
 * @param ctx the context to allocate, will be NULL on error
 * @param tx pointer to the transform function pointer to set
 * @param type type the type of transform
 * @param inv whether to do an inverse or a forward transform
 * @param len the size of the transform in samples
 * @param scale pointer to the value to scale the output if supported by type
 * @param flags a bitmask of AVTXFlags or 0
 *
 * @return 0 on success, negative error code on failure
 */
func Av_tx_init(ctx **AVTXContext, tx *av_tx_fn, typex AVTXType,
               inv int32, len int32, scale unsafe.Pointer, flags uint64) int32 {
    return int32(C.av_tx_init((**C.struct_AVTXContext)(unsafe.Pointer(ctx)), 
        (*C.av_tx_fn)(unsafe.Pointer(tx)), C.enum_AVTXType(typex), C.int(inv), 
        C.int(len), scale, C.ulonglong(flags)))
}

/**
 * Frees a context and sets *ctx to NULL, does nothing when *ctx == NULL.
 */
func Av_tx_uninit(ctx **AVTXContext)  {
    C.av_tx_uninit((**C.struct_AVTXContext)(unsafe.Pointer(ctx)))
}

                        
