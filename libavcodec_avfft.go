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

//#cgo pkg-config: libavcodec libavutil
//#include "libavutil/attributes.h"
//#include "libavcodec/version_major.h"
//#include "libavcodec/avfft.h"
import "C"
import (
    "unsafe"
)



                       
                       

                                 
                          
                

/**
 * @file
 * @ingroup lavc_fft
 * FFT functions
 */

/**
 * @defgroup lavc_fft FFT functions
 * @ingroup lavc_misc
 *
 * @{
 */

type FFTSample float32

type FFTComplex struct {
    Re FFTSample
    Im FFTSample
}


type FFTContext struct {
}


/**
 * Set up a complex FFT.
 * @param nbits           log2 of the length of the input array
 * @param inverse         if 0 perform the forward transform, if 1 perform the inverse
 * @deprecated use av_tx_init from libavutil/tx.h with a type of AV_TX_FLOAT_FFT
 */

func Av_fft_init(nbits int32, inverse int32) *FFTContext {
    return (*FFTContext)(unsafe.Pointer(C.av_fft_init(C.int(nbits), C.int(inverse))))
}

/**
 * Do the permutation needed BEFORE calling ff_fft_calc().
 * @deprecated without replacement
 */

func Av_fft_permute(s *FFTContext, z *FFTComplex)  {
    C.av_fft_permute((*C.struct_FFTContext)(unsafe.Pointer(s)), 
        (*C.struct_FFTComplex)(unsafe.Pointer(z)))
}

/**
 * Do a complex FFT with the parameters defined in av_fft_init(). The
 * input data must be permuted before. No 1.0/sqrt(n) normalization is done.
 * @deprecated use the av_tx_fn value returned by av_tx_init, which also does permutation
 */

func Av_fft_calc(s *FFTContext, z *FFTComplex)  {
    C.av_fft_calc((*C.struct_FFTContext)(unsafe.Pointer(s)), 
        (*C.struct_FFTComplex)(unsafe.Pointer(z)))
}


func Av_fft_end(s *FFTContext)  {
    C.av_fft_end((*C.struct_FFTContext)(unsafe.Pointer(s)))
}

/**
 * @deprecated use av_tx_init from libavutil/tx.h with a type of AV_TX_FLOAT_MDCT,
 * with a flag of AV_TX_FULL_IMDCT for a replacement to av_imdct_calc.
 */

func Av_mdct_init(nbits int32, inverse int32, scale float64) *FFTContext {
    return (*FFTContext)(unsafe.Pointer(C.av_mdct_init(C.int(nbits), C.int(inverse), 
        C.double(scale))))
}

func Av_imdct_calc(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_imdct_calc((*C.struct_FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}

func Av_imdct_half(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_imdct_half((*C.struct_FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}

func Av_mdct_calc(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_mdct_calc((*C.struct_FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}

func Av_mdct_end(s *FFTContext)  {
    C.av_mdct_end((*C.struct_FFTContext)(unsafe.Pointer(s)))
}

/* Real Discrete Fourier Transform */

type RDFTransformType int32
const (
    DFT_R2C RDFTransformType = iota
    IDFT_C2R
    IDFT_R2C
    DFT_C2R
)


type RDFTContext struct {
}


/**
 * Set up a real FFT.
 * @param nbits           log2 of the length of the input array
 * @param trans           the type of transform
 *
 * @deprecated use av_tx_init from libavutil/tx.h with a type of AV_TX_FLOAT_RDFT
 */

func Av_rdft_init(nbits int32, trans RDFTransformType) *RDFTContext {
    return (*RDFTContext)(unsafe.Pointer(C.av_rdft_init(C.int(nbits), 
        C.enum_RDFTransformType(trans))))
}

func Av_rdft_calc(s *RDFTContext, data *FFTSample)  {
    C.av_rdft_calc((*C.struct_RDFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(data)))
}

func Av_rdft_end(s *RDFTContext)  {
    C.av_rdft_end((*C.struct_RDFTContext)(unsafe.Pointer(s)))
}

/* Discrete Cosine Transform */

type DCTContext struct {
}


type DCTTransformType int32
const (
    DCT_II DCTTransformType = 0 + iota
    DCT_III
    DCT_I
    DST_I
)


/**
 * Set up DCT.
 *
 * @param nbits           size of the input array:
 *                        (1 << nbits)     for DCT-II, DCT-III and DST-I
 *                        (1 << nbits) + 1 for DCT-I
 * @param type            the type of transform
 *
 * @note the first element of the input of DST-I is ignored
 *
 * @deprecated use av_tx_init from libavutil/tx.h with an appropriate type of AV_TX_FLOAT_DCT
 */

func Av_dct_init(nbits int32, typex DCTTransformType) *DCTContext {
    return (*DCTContext)(unsafe.Pointer(C.av_dct_init(C.int(nbits), 
        C.enum_DCTTransformType(typex))))
}

func Av_dct_calc(s *DCTContext, data *FFTSample)  {
    C.av_dct_calc((*C.struct_DCTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(data)))
}

func Av_dct_end (s *DCTContext)  {
    C.av_dct_end((*C.struct_DCTContext)(unsafe.Pointer(s)))
}

/**
 * @}
 */

                         
                            
