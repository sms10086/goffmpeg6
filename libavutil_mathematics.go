// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2005-2012 Michael Niedermayer <michaelni@gmx.at>
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
//#include <stdint.h>
//#include <math.h>
//#include "libavutil/attributes.h"
//#include "libavutil/rational.h"
//#include "libavutil/intfloat.h"
//#include "libavutil/mathematics.h"
import "C"
import (
    "unsafe"
)

const M_E = 2.7182818284590452354
const M_Ef = 2.7182818284590452354
const M_LN2 = 0.69314718055994530942
const M_LN2f = 0.69314718055994530942
const M_LN10 = 2.30258509299404568402
const M_LN10f = 2.30258509299404568402
const M_LOG2_10 = 3.32192809488736234787
const M_LOG2_10f = 3.32192809488736234787
const M_PHI = 1.61803398874989484820
const M_PHIf = 1.61803398874989484820
const M_PI = 3.14159265358979323846
const M_PIf = 3.14159265358979323846
const M_PI_2 = 1.57079632679489661923
const M_PI_2f = 1.57079632679489661923
const M_PI_4 = 0.78539816339744830962
const M_PI_4f = 0.78539816339744830962
const M_1_PI = 0.31830988618379067154
const M_1_PIf = 0.31830988618379067154
const M_2_PI = 0.63661977236758134308
const M_2_PIf = 0.63661977236758134308
const M_2_SQRTPI = 1.12837916709551257390
const M_2_SQRTPIf = 1.12837916709551257390
const M_SQRT1_2 = 0.70710678118654752440
const M_SQRT1_2f = 0.70710678118654752440
const M_SQRT2 = 1.41421356237309504880
const M_SQRT2f = 1.41421356237309504880
//const NAN = av_int2float(0x7fc00000)
//const INFINITY = av_int2float(0x7f800000)


/**
 * @file
 * @addtogroup lavu_math
 * Mathematical utilities for working with timestamp and time base.
 */

                            
                            

                   
                 
                       
                     
                     

           
/* e */
      
            
/* e */
      
             
/* log_e 2 */
      
              
/* log_e 2 */
      
              
/* log_e 10 */
      
               
/* log_e 10 */
      
                 
/* log_2 10 */
      
                  
/* log_2 10 */
      
             
/* phi / golden ratio */
      
              
/* phi / golden ratio */
      
            
/* pi */
      
             
/* pi */
      
              
/* pi/2 */
      
               
/* pi/2 */
      
              
/* pi/4 */
      
               
/* pi/4 */
      
              
/* 1/pi */
      
               
/* 1/pi */
      
              
/* 2/pi */
      
               
/* 2/pi */
      
                  
/* 2/sqrt(pi) */
      
                   
/* 2/sqrt(pi) */
      
                 
/* 1/sqrt(2) */
      
                  
/* 1/sqrt(2) */
      
               
/* sqrt(2) */
      
                
/* sqrt(2) */
      
           

      
                

      

/**
 * @addtogroup lavu_math
 *
 * @{
 */

/**
 * Rounding methods.
 */
type AVRounding C.enum_AVRounding

/**
 * Compute the greatest common divisor of two integer operands.
 *
 * @param a Operand
 * @param b Operand
 * @return GCD of a and b up to sign; if a >= 0 and b >= 0, return value is >= 0;
 * if a == 0 and b == 0, returns 0.
 */
func  Av_gcd(a int64, b int64) int64 {
    return int64(C.av_gcd(C.longlong(a), C.longlong(b)))
}

/**
 * Rescale a 64-bit integer with rounding to nearest.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow.
 *
 * This function is equivalent to av_rescale_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale_rnd(), av_rescale_q(), av_rescale_q_rnd()
 */
func Av_rescale(a int64, b int64, c int64)  int64 {
    return int64(C.av_rescale(C.longlong(a), C.longlong(b), C.longlong(c)))
}

/**
 * Rescale a 64-bit integer with specified rounding.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow, and does not support different rounding methods.
 * If the result is not representable then INT64_MIN is returned.
 *
 * @see av_rescale(), av_rescale_q(), av_rescale_q_rnd()
 */
func Av_rescale_rnd(a int64, b int64, c int64, rnd AVRounding)  int64 {
    return int64(C.av_rescale_rnd(C.longlong(a), C.longlong(b), C.longlong(c), 
        C.enum_AVRounding(rnd)))
}

/**
 * Rescale a 64-bit integer by 2 rational numbers.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * This function is equivalent to av_rescale_q_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q_rnd()
 */
func Av_rescale_q(a int64, bq AVRational, cq AVRational)  int64 {
    return int64(C.av_rescale_q(C.longlong(a), C.AVRational(bq), 
        C.AVRational(cq)))
}

/**
 * Rescale a 64-bit integer by 2 rational numbers with specified rounding.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q()
 */
func Av_rescale_q_rnd(a int64, bq AVRational, cq AVRational,
                         rnd AVRounding)  int64 {
    return int64(C.av_rescale_q_rnd(C.longlong(a), C.AVRational(bq), 
        C.AVRational(cq), C.enum_AVRounding(rnd)))
}

/**
 * Compare two timestamps each in its own time base.
 *
 * @return One of the following values:
 *         - -1 if `ts_a` is before `ts_b`
 *         - 1 if `ts_a` is after `ts_b`
 *         - 0 if they represent the same position
 *
 * @warning
 * The result of the function is undefined if one of the timestamps is outside
 * the `int64_t` range when represented in the other's timebase.
 */
func Av_compare_ts(ts_a int64, tb_a AVRational, ts_b int64, tb_b AVRational) int32 {
    return int32(C.av_compare_ts(C.longlong(ts_a), C.AVRational(tb_a), 
        C.longlong(ts_b), C.AVRational(tb_b)))
}

/**
 * Compare the remainders of two integer operands divided by a common divisor.
 *
 * In other words, compare the least significant `log2(mod)` bits of integers
 * `a` and `b`.
 *
 * @code{.c}
 * av_compare_mod(0x11, 0x02, 0x10) < 0 // since 0x11 % 0x10  (0x1) < 0x02 % 0x10  (0x2)
 * av_compare_mod(0x11, 0x02, 0x20) > 0 // since 0x11 % 0x20 (0x11) > 0x02 % 0x20 (0x02)
 * @endcode
 *
 * @param a Operand
 * @param b Operand
 * @param mod Divisor; must be a power of 2
 * @return
 *         - a negative value if `a % mod < b % mod`
 *         - a positive value if `a % mod > b % mod`
 *         - zero             if `a % mod == b % mod`
 */
func Av_compare_mod(a uint64, b uint64, mod uint64) int64 {
    return int64(C.av_compare_mod(C.ulonglong(a), C.ulonglong(b), 
        C.ulonglong(mod)))
}

/**
 * Rescale a timestamp while preserving known durations.
 *
 * This function is designed to be called per audio packet to scale the input
 * timestamp to a different time base. Compared to a simple av_rescale_q()
 * call, this function is robust against possible inconsistent frame durations.
 *
 * The `last` parameter is a state variable that must be preserved for all
 * subsequent calls for the same stream. For the first call, `*last` should be
 * initialized to #AV_NOPTS_VALUE.
 *
 * @param[in]     in_tb    Input time base
 * @param[in]     in_ts    Input timestamp
 * @param[in]     fs_tb    Duration time base; typically this is finer-grained
 *                         (greater) than `in_tb` and `out_tb`
 * @param[in]     duration Duration till the next call to this function (i.e.
 *                         duration of the current packet/frame)
 * @param[in,out] last     Pointer to a timestamp expressed in terms of
 *                         `fs_tb`, acting as a state variable
 * @param[in]     out_tb   Output timebase
 * @return        Timestamp expressed in terms of `out_tb`
 *
 * @note In the context of this function, "duration" is in term of samples, not
 *       seconds.
 */
func Av_rescale_delta(in_tb AVRational, in_ts int64,  fs_tb AVRational, duration int32, last *int64, out_tb AVRational) int64 {
    return int64(C.av_rescale_delta(C.AVRational(in_tb), C.longlong(in_ts), 
        C.AVRational(fs_tb), C.int(duration), 
        (*C.longlong)(unsafe.Pointer(last)), C.AVRational(out_tb)))
}

/**
 * Add a value to a timestamp.
 *
 * This function guarantees that when the same value is repeatly added that
 * no accumulation of rounding errors occurs.
 *
 * @param[in] ts     Input timestamp
 * @param[in] ts_tb  Input timestamp time base
 * @param[in] inc    Value to be added
 * @param[in] inc_tb Time base of `inc`
 */
func Av_add_stable(ts_tb AVRational, ts int64, inc_tb AVRational, inc int64) int64 {
    return int64(C.av_add_stable(C.AVRational(ts_tb), C.longlong(ts), 
        C.AVRational(inc_tb), C.longlong(inc)))
}

/**
 * 0th order modified bessel function of the first kind.
 */
func Av_bessel_i0(x float64) float64 {
    return float64(C.av_bessel_i0(C.double(x)))
}

/**
 * @}
 */

                                 
