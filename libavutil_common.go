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
//#include <errno.h>
//#include <inttypes.h>
//#include <limits.h>
//#include <math.h>
//#include <stdint.h>
//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
//#include "libavutil/attributes.h"
//#include "libavutil/error.h"
//#include "libavutil/macros.h"
//#include "libavutil/mem.h"
//#include "libavutil/common.h"
import "C"



/**
 * @file
 * common internal and external API header
 */

                       
                       

                                                                                  
                                                                        
      

                  
                     
                   
                 
                   
                  
                   
                   

                       
                  
                   

//rounded division & shift

/* assume b>0 */

/* Fast a/(1<<b) rounded toward +inf. Assume a>=0 and b>=0 */

/* Backwards compat. */





/**
 * Absolute value, Note, INT_MIN / INT64_MIN result in undefined behavior as they
 * are not representable as absolute values of their type. This is the same
 * as with *abs()
 * @see FFNABS()
 */



/**
 * Negative Absolute value.
 * this works for all integers of all types.
 * As with many macros, this evaluates its argument twice, it thus must not have
 * a sideeffect, that is FFNABS(x++) has undefined behavior.
 */


/**
 * Unsigned Absolute value.
 * This takes the absolute value of a signed int and returns it as a unsigned.
 * This also works with INT_MIN which would otherwise not be representable
 * As with many macros, this evaluates its argument twice.
 */



/* misc math functions */

                       
                      
                       
      

                    

      
               

      
                 

      
                     

      
                    

      
                      

      
                     

      
                      

      
                     

      
                      

      
                     

      
                    

      
                     

      
                    

      
                     

      
                    

      
                    

      
                

      
                

      
                   

      
                     

      
                 

      

               
 func Av_log2(v uint32) int32 {
    return int32(C.av_log2(C.unsigned(v)))
}
      

                     
 func Av_log2_16bit(v uint32) int32 {
    return int32(C.av_log2_16bit(C.unsigned(v)))
}
      

/**
 * Clip a signed integer value into the amin-amax range.
 * @param a value to clip
 * @param amin minimum value of the clip range
 * @param amax maximum value of the clip range
 * @return clipped value
 */
// av_clip_c(inta,intamin,intamax)

/**
 * Clip a signed 64bit integer value into the amin-amax range.
 * @param a value to clip
 * @param amin minimum value of the clip range
 * @param amax maximum value of the clip range
 * @return clipped value
 */
// av_clip64_c(int64_ta,int64_tamin,int64_tamax)

/**
 * Clip a signed integer value into the 0-255 range.
 * @param a value to clip
 * @return clipped value
 */
// av_clip_uint8_c(inta)

/**
 * Clip a signed integer value into the -128,127 range.
 * @param a value to clip
 * @return clipped value
 */
// av_clip_int8_c(inta)

/**
 * Clip a signed integer value into the 0-65535 range.
 * @param a value to clip
 * @return clipped value
 */
// av_clip_uint16_c(inta)

/**
 * Clip a signed integer value into the -32768,32767 range.
 * @param a value to clip
 * @return clipped value
 */
// av_clip_int16_c(inta)

/**
 * Clip a signed 64-bit integer value into the -2147483648,2147483647 range.
 * @param a value to clip
 * @return clipped value
 */
// av_clipl_int32_c(int64_ta)

/**
 * Clip a signed integer into the -(2^p),(2^p-1) range.
 * @param  a value to clip
 * @param  p bit position to clip at
 * @return clipped value
 */
// av_clip_intp2_c(inta,intp)

/**
 * Clip a signed integer to an unsigned power of two range.
 * @param  a value to clip
 * @param  p bit position to clip at
 * @return clipped value
 */
// av_clip_uintp2_c(inta,intp)

/**
 * Clear high bits from an unsigned integer starting with specific bit position
 * @param  a value to clip
 * @param  p bit position to clip at
 * @return clipped value
 */
// av_mod_uintp2_c(unsigneda,unsignedp)

/**
 * Add two signed 32-bit values with saturation.
 *
 * @param  a one value
 * @param  b another value
 * @return sum with signed saturation
 */
// av_sat_add32_c(inta,intb)

/**
 * Add a doubled value to another value with saturation at both stages.
 *
 * @param  a first value
 * @param  b value doubled and added to a
 * @return sum sat(a + sat(2*b)) with signed saturation
 */
// av_sat_dadd32_c(inta,intb)

/**
 * Subtract two signed 32-bit values with saturation.
 *
 * @param  a one value
 * @param  b another value
 * @return difference with signed saturation
 */
// av_sat_sub32_c(inta,intb)

/**
 * Subtract a doubled value from another value with saturation at both stages.
 *
 * @param  a first value
 * @param  b value doubled and subtracted from a
 * @return difference sat(a - sat(2*b)) with signed saturation
 */
// av_sat_dsub32_c(inta,intb)

/**
 * Add two signed 64-bit values with saturation.
 *
 * @param  a one value
 * @param  b another value
 * @return sum with signed saturation
 */
// av_sat_add64_c(int64_ta,int64_tb)

/**
 * Subtract two signed 64-bit values with saturation.
 *
 * @param  a one value
 * @param  b another value
 * @return difference with signed saturation
 */
// av_sat_sub64_c(int64_ta,int64_tb)

/**
 * Clip a float value into the amin-amax range.
 * If a is nan or -inf amin will be returned.
 * If a is +inf amax will be returned.
 * @param a value to clip
 * @param amin minimum value of the clip range
 * @param amax maximum value of the clip range
 * @return clipped value
 */
// av_clipf_c(floata,floatamin,floatamax)

/**
 * Clip a double value into the amin-amax range.
 * If a is nan or -inf amin will be returned.
 * If a is +inf amax will be returned.
 * @param a value to clip
 * @param amin minimum value of the clip range
 * @param amax maximum value of the clip range
 * @return clipped value
 */
// av_clipd_c(doublea,doubleamin,doubleamax)

/** Compute ceil(log2(x)).
 * @param x value used to compute ceil(log2(x))
 * @return computed ceiling of log2(x)
 */
// av_ceil_log2_c(intx)

/**
 * Count number of bits set to one in x
 * @param x value to count bits of
 * @return the number of bits set to one in x
 */
// av_popcount_c(uint32_tx)

/**
 * Count number of bits set to one in x
 * @param x value to count bits of
 * @return the number of bits set to one in x
 */
// av_popcount64_c(uint64_tx)

// av_parity_c(uint32_tv)

/**
 * Convert a UTF-8 character (up to 4 bytes) to its 32-bit UCS-4 encoded form.
 *
 * @param val      Output value, must be an lvalue of type uint32_t.
 * @param GET_BYTE Expression reading one byte from the input.
 *                 Evaluated up to 7 times (4 for the currently
 *                 assigned Unicode range).  With a memory buffer
 *                 input, this could be *ptr++, or if you want to make sure
 *                 that *ptr stops at the end of a NULL terminated string then
 *                 *ptr ? *ptr++ : 0
 * @param ERROR    Expression to be evaluated on invalid input,
 *                 typically a goto statement.
 *
 * @warning ERROR should not contain a loop control statement which
 * could interact with the internal while loop, and should force an
 * exit from the macro code (e.g. through a goto or a return) in order
 * to prevent undefined results.
 */


/**
 * Convert a UTF-16 character (2 or 4 bytes) to its 32-bit UCS-4 encoded form.
 *
 * @param val       Output value, must be an lvalue of type uint32_t.
 * @param GET_16BIT Expression returning two bytes of UTF-16 data converted
 *                  to native byte order.  Evaluated one or two times.
 * @param ERROR     Expression to be evaluated on invalid input,
 *                  typically a goto statement.
 */


/**
 * @def PUT_UTF8(val, tmp, PUT_BYTE)
 * Convert a 32-bit Unicode character to its UTF-8 encoded form (up to 4 bytes long).
 * @param val is an input-only argument and should be of type uint32_t. It holds
 * a UCS-4 encoded Unicode character that is to be converted to UTF-8. If
 * val is given as a function it is executed only once.
 * @param tmp is a temporary variable and should be of type uint8_t. It
 * represents an intermediate value during conversion that is to be
 * output by PUT_BYTE.
 * @param PUT_BYTE writes the converted UTF-8 bytes to any proper destination.
 * It could be a function or a statement, and uses tmp as the input byte.
 * For example, PUT_BYTE could be "*output++ = tmp;" PUT_BYTE will be
 * executed up to 4 times for values in the valid UTF-8 range and up to
 * 7 times in the general case, depending on the length of the converted
 * Unicode character.
 */


/**
 * @def PUT_UTF16(val, tmp, PUT_16BIT)
 * Convert a 32-bit Unicode character to its UTF-16 encoded form (2 or 4 bytes).
 * @param val is an input-only argument and should be of type uint32_t. It holds
 * a UCS-4 encoded Unicode character that is to be converted to UTF-16. If
 * val is given as a function it is executed only once.
 * @param tmp is a temporary variable and should be of type uint16_t. It
 * represents an intermediate value during conversion that is to be
 * output by PUT_16BIT.
 * @param PUT_16BIT writes the converted UTF-16 data to any proper destination
 * in desired endianness. It could be a function or a statement, and uses tmp
 * as the input byte.  For example, PUT_BYTE could be "*output++ = tmp;"
 * PUT_BYTE will be executed 1 or 2 times depending on input character.
 */




                

                       
                         
                             

                            
