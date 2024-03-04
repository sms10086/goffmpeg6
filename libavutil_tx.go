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



                   
                   

                   
                   

type AVTXContext C.struct_AVTXContext

type AVComplexFloat C.struct_AVComplexFloat

type AVComplexDouble C.struct_AVComplexDouble

type AVComplexInt32 C.struct_AVComplexInt32

type AVTXType C.enum_AVTXType

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
type AVTXFlags C.enum_AVTXFlags

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
    return int32(C.av_tx_init((**C.AVTXContext)(unsafe.Pointer(ctx)), 
        (*C.av_tx_fn)(unsafe.Pointer(tx)), C.enum_AVTXType(typex), C.int(inv), 
        C.int(len), scale, C.ulonglong(flags)))
}

/**
 * Frees a context and sets *ctx to NULL, does nothing when *ctx == NULL.
 */
func Av_tx_uninit(ctx **AVTXContext)  {
    C.av_tx_uninit((**C.AVTXContext)(unsafe.Pointer(ctx)))
}

                        
