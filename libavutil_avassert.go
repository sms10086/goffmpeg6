// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2010 Michael Niedermayer <michaelni@gmx.at>
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
//#include <stdlib.h>
//#include "libavutil/log.h"
//#include "libavutil/macros.h"
//#include "libavutil/avassert.h"
import "C"



/**
 * @file
 * simple assert() macros that are a bit more flexible than ISO C assert().
 * @author Michael Niedermayer <michaelni@gmx.at>
 */

                         
                         

                   
                       
                      
      
                
                   

/**
 * assert() equivalent, that is always enabled.
 */



/**
 * assert() equivalent, that does not lie in speed critical code.
 * These asserts() thus can be enabled without fearing speed loss.
 */
                                             
                                         
     

      


/**
 * assert() equivalent, that does lie in speed critical code.
 */
                                             
                                         
                                         
     


      

/**
 * Assert that floating point operations can be executed.
 *
 * This will av_assert0() that the cpu is not in MMX state on X86
 */
func Av_assert0_fpu()  {
    C.av_assert0_fpu()
}

                              
