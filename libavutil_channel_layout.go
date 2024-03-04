// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
 * Copyright (c) 2008 Peter Ross
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
//#include <stdlib.h>
//#include "libavutil/version.h"
//#include "libavutil/attributes.h"
//#include "libavutil/channel_layout.h"
import "C"
import (
    "unsafe"
)

//const AV_CH_FRONT_LEFT = (1ULL << AV_CHAN_FRONT_LEFT           )
//const AV_CH_FRONT_RIGHT = (1ULL << AV_CHAN_FRONT_RIGHT          )
//const AV_CH_FRONT_CENTER = (1ULL << AV_CHAN_FRONT_CENTER         )
//const AV_CH_LOW_FREQUENCY = (1ULL << AV_CHAN_LOW_FREQUENCY        )
//const AV_CH_BACK_LEFT = (1ULL << AV_CHAN_BACK_LEFT            )
//const AV_CH_BACK_RIGHT = (1ULL << AV_CHAN_BACK_RIGHT           )
//const AV_CH_FRONT_LEFT_OF_CENTER = (1ULL << AV_CHAN_FRONT_LEFT_OF_CENTER )
//const AV_CH_FRONT_RIGHT_OF_CENTER = (1ULL << AV_CHAN_FRONT_RIGHT_OF_CENTER)
//const AV_CH_BACK_CENTER = (1ULL << AV_CHAN_BACK_CENTER          )
//const AV_CH_SIDE_LEFT = (1ULL << AV_CHAN_SIDE_LEFT            )
//const AV_CH_SIDE_RIGHT = (1ULL << AV_CHAN_SIDE_RIGHT           )
//const AV_CH_TOP_CENTER = (1ULL << AV_CHAN_TOP_CENTER           )
//const AV_CH_TOP_FRONT_LEFT = (1ULL << AV_CHAN_TOP_FRONT_LEFT       )
//const AV_CH_TOP_FRONT_CENTER = (1ULL << AV_CHAN_TOP_FRONT_CENTER     )
//const AV_CH_TOP_FRONT_RIGHT = (1ULL << AV_CHAN_TOP_FRONT_RIGHT      )
//const AV_CH_TOP_BACK_LEFT = (1ULL << AV_CHAN_TOP_BACK_LEFT        )
//const AV_CH_TOP_BACK_CENTER = (1ULL << AV_CHAN_TOP_BACK_CENTER      )
//const AV_CH_TOP_BACK_RIGHT = (1ULL << AV_CHAN_TOP_BACK_RIGHT       )
//const AV_CH_STEREO_LEFT = (1ULL << AV_CHAN_STEREO_LEFT          )
//const AV_CH_STEREO_RIGHT = (1ULL << AV_CHAN_STEREO_RIGHT         )
//const AV_CH_WIDE_LEFT = (1ULL << AV_CHAN_WIDE_LEFT            )
//const AV_CH_WIDE_RIGHT = (1ULL << AV_CHAN_WIDE_RIGHT           )
//const AV_CH_SURROUND_DIRECT_LEFT = (1ULL << AV_CHAN_SURROUND_DIRECT_LEFT )
//const AV_CH_SURROUND_DIRECT_RIGHT = (1ULL << AV_CHAN_SURROUND_DIRECT_RIGHT)
//const AV_CH_LOW_FREQUENCY_2 = (1ULL << AV_CHAN_LOW_FREQUENCY_2      )
//const AV_CH_TOP_SIDE_LEFT = (1ULL << AV_CHAN_TOP_SIDE_LEFT        )
//const AV_CH_TOP_SIDE_RIGHT = (1ULL << AV_CHAN_TOP_SIDE_RIGHT       )
//const AV_CH_BOTTOM_FRONT_CENTER = (1ULL << AV_CHAN_BOTTOM_FRONT_CENTER  )
//const AV_CH_BOTTOM_FRONT_LEFT = (1ULL << AV_CHAN_BOTTOM_FRONT_LEFT    )
//const AV_CH_BOTTOM_FRONT_RIGHT = (1ULL << AV_CHAN_BOTTOM_FRONT_RIGHT   )
//const AV_CH_LAYOUT_MONO = (AV_CH_FRONT_CENTER)
//const AV_CH_LAYOUT_STEREO = (AV_CH_FRONT_LEFT|AV_CH_FRONT_RIGHT)
//const AV_CH_LAYOUT_2POINT1 = (AV_CH_LAYOUT_STEREO|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_2_1 = (AV_CH_LAYOUT_STEREO|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_SURROUND = (AV_CH_LAYOUT_STEREO|AV_CH_FRONT_CENTER)
//const AV_CH_LAYOUT_3POINT1 = (AV_CH_LAYOUT_SURROUND|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_4POINT0 = (AV_CH_LAYOUT_SURROUND|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_4POINT1 = (AV_CH_LAYOUT_4POINT0|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_2_2 = (AV_CH_LAYOUT_STEREO|AV_CH_SIDE_LEFT|AV_CH_SIDE_RIGHT)
//const AV_CH_LAYOUT_QUAD = (AV_CH_LAYOUT_STEREO|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT)
//const AV_CH_LAYOUT_5POINT0 = (AV_CH_LAYOUT_SURROUND|AV_CH_SIDE_LEFT|AV_CH_SIDE_RIGHT)
//const AV_CH_LAYOUT_5POINT1 = (AV_CH_LAYOUT_5POINT0|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_5POINT0_BACK = (AV_CH_LAYOUT_SURROUND|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT)
//const AV_CH_LAYOUT_5POINT1_BACK = (AV_CH_LAYOUT_5POINT0_BACK|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_6POINT0 = (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_6POINT0_FRONT = (AV_CH_LAYOUT_2_2|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER)
//const AV_CH_LAYOUT_HEXAGONAL = (AV_CH_LAYOUT_5POINT0_BACK|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_3POINT1POINT2 = (AV_CH_LAYOUT_3POINT1|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT)
//const AV_CH_LAYOUT_6POINT1 = (AV_CH_LAYOUT_5POINT1|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_6POINT1_BACK = (AV_CH_LAYOUT_5POINT1_BACK|AV_CH_BACK_CENTER)
//const AV_CH_LAYOUT_6POINT1_FRONT = (AV_CH_LAYOUT_6POINT0_FRONT|AV_CH_LOW_FREQUENCY)
//const AV_CH_LAYOUT_7POINT0 = (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT)
//const AV_CH_LAYOUT_7POINT0_FRONT = (AV_CH_LAYOUT_5POINT0|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER)
//const AV_CH_LAYOUT_7POINT1 = (AV_CH_LAYOUT_5POINT1|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT)
//const AV_CH_LAYOUT_7POINT1_WIDE = (AV_CH_LAYOUT_5POINT1|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER)
//const AV_CH_LAYOUT_7POINT1_WIDE_BACK = (AV_CH_LAYOUT_5POINT1_BACK|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER)
//const AV_CH_LAYOUT_5POINT1POINT2_BACK = (AV_CH_LAYOUT_5POINT1_BACK|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT)
//const AV_CH_LAYOUT_OCTAGONAL = (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_LEFT|AV_CH_BACK_CENTER|AV_CH_BACK_RIGHT)
//const AV_CH_LAYOUT_CUBE = (AV_CH_LAYOUT_QUAD|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT|AV_CH_TOP_BACK_LEFT|AV_CH_TOP_BACK_RIGHT)
//const AV_CH_LAYOUT_5POINT1POINT4_BACK = (AV_CH_LAYOUT_5POINT1POINT2_BACK|AV_CH_TOP_BACK_LEFT|AV_CH_TOP_BACK_RIGHT)
//const AV_CH_LAYOUT_7POINT1POINT2 = (AV_CH_LAYOUT_7POINT1|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT)
//const AV_CH_LAYOUT_7POINT1POINT4_BACK = (AV_CH_LAYOUT_7POINT1POINT2|AV_CH_TOP_BACK_LEFT|AV_CH_TOP_BACK_RIGHT)
//const AV_CH_LAYOUT_HEXADECAGONAL = (AV_CH_LAYOUT_OCTAGONAL|AV_CH_WIDE_LEFT|AV_CH_WIDE_RIGHT|AV_CH_TOP_BACK_LEFT|AV_CH_TOP_BACK_RIGHT|AV_CH_TOP_BACK_CENTER|AV_CH_TOP_FRONT_CENTER|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT)
//const AV_CH_LAYOUT_STEREO_DOWNMIX = (AV_CH_STEREO_LEFT|AV_CH_STEREO_RIGHT)
//const AV_CH_LAYOUT_22POINT2 = (AV_CH_LAYOUT_7POINT1POINT4_BACK|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER|AV_CH_BACK_CENTER|AV_CH_LOW_FREQUENCY_2|AV_CH_TOP_FRONT_CENTER|AV_CH_TOP_CENTER|AV_CH_TOP_SIDE_LEFT|AV_CH_TOP_SIDE_RIGHT|AV_CH_TOP_BACK_CENTER|AV_CH_BOTTOM_FRONT_CENTER|AV_CH_BOTTOM_FRONT_LEFT|AV_CH_BOTTOM_FRONT_RIGHT)
//const AV_CH_LAYOUT_7POINT1_TOP_BACK = AV_CH_LAYOUT_5POINT1POINT2_BACK
//const AV_CHANNEL_LAYOUT_MONO = AV_CHANNEL_LAYOUT_MASK(1,  AV_CH_LAYOUT_MONO)
//const AV_CHANNEL_LAYOUT_STEREO = AV_CHANNEL_LAYOUT_MASK(2,  AV_CH_LAYOUT_STEREO)
//const AV_CHANNEL_LAYOUT_2POINT1 = AV_CHANNEL_LAYOUT_MASK(3,  AV_CH_LAYOUT_2POINT1)
//const AV_CHANNEL_LAYOUT_2_1 = AV_CHANNEL_LAYOUT_MASK(3,  AV_CH_LAYOUT_2_1)
//const AV_CHANNEL_LAYOUT_SURROUND = AV_CHANNEL_LAYOUT_MASK(3,  AV_CH_LAYOUT_SURROUND)
//const AV_CHANNEL_LAYOUT_3POINT1 = AV_CHANNEL_LAYOUT_MASK(4,  AV_CH_LAYOUT_3POINT1)
//const AV_CHANNEL_LAYOUT_4POINT0 = AV_CHANNEL_LAYOUT_MASK(4,  AV_CH_LAYOUT_4POINT0)
//const AV_CHANNEL_LAYOUT_4POINT1 = AV_CHANNEL_LAYOUT_MASK(5,  AV_CH_LAYOUT_4POINT1)
//const AV_CHANNEL_LAYOUT_2_2 = AV_CHANNEL_LAYOUT_MASK(4,  AV_CH_LAYOUT_2_2)
//const AV_CHANNEL_LAYOUT_QUAD = AV_CHANNEL_LAYOUT_MASK(4,  AV_CH_LAYOUT_QUAD)
//const AV_CHANNEL_LAYOUT_5POINT0 = AV_CHANNEL_LAYOUT_MASK(5,  AV_CH_LAYOUT_5POINT0)
//const AV_CHANNEL_LAYOUT_5POINT1 = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_5POINT1)
//const AV_CHANNEL_LAYOUT_5POINT0_BACK = AV_CHANNEL_LAYOUT_MASK(5,  AV_CH_LAYOUT_5POINT0_BACK)
//const AV_CHANNEL_LAYOUT_5POINT1_BACK = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_5POINT1_BACK)
//const AV_CHANNEL_LAYOUT_6POINT0 = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_6POINT0)
//const AV_CHANNEL_LAYOUT_6POINT0_FRONT = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_6POINT0_FRONT)
//const AV_CHANNEL_LAYOUT_3POINT1POINT2 = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_3POINT1POINT2)
//const AV_CHANNEL_LAYOUT_HEXAGONAL = AV_CHANNEL_LAYOUT_MASK(6,  AV_CH_LAYOUT_HEXAGONAL)
//const AV_CHANNEL_LAYOUT_6POINT1 = AV_CHANNEL_LAYOUT_MASK(7,  AV_CH_LAYOUT_6POINT1)
//const AV_CHANNEL_LAYOUT_6POINT1_BACK = AV_CHANNEL_LAYOUT_MASK(7,  AV_CH_LAYOUT_6POINT1_BACK)
//const AV_CHANNEL_LAYOUT_6POINT1_FRONT = AV_CHANNEL_LAYOUT_MASK(7,  AV_CH_LAYOUT_6POINT1_FRONT)
//const AV_CHANNEL_LAYOUT_7POINT0 = AV_CHANNEL_LAYOUT_MASK(7,  AV_CH_LAYOUT_7POINT0)
//const AV_CHANNEL_LAYOUT_7POINT0_FRONT = AV_CHANNEL_LAYOUT_MASK(7,  AV_CH_LAYOUT_7POINT0_FRONT)
//const AV_CHANNEL_LAYOUT_7POINT1 = AV_CHANNEL_LAYOUT_MASK(8,  AV_CH_LAYOUT_7POINT1)
//const AV_CHANNEL_LAYOUT_7POINT1_WIDE = AV_CHANNEL_LAYOUT_MASK(8,  AV_CH_LAYOUT_7POINT1_WIDE)
//const AV_CHANNEL_LAYOUT_7POINT1_WIDE_BACK = AV_CHANNEL_LAYOUT_MASK(8,  AV_CH_LAYOUT_7POINT1_WIDE_BACK)
//const AV_CHANNEL_LAYOUT_5POINT1POINT2_BACK = AV_CHANNEL_LAYOUT_MASK(8, AV_CH_LAYOUT_5POINT1POINT2_BACK)
//const AV_CHANNEL_LAYOUT_OCTAGONAL = AV_CHANNEL_LAYOUT_MASK(8,  AV_CH_LAYOUT_OCTAGONAL)
//const AV_CHANNEL_LAYOUT_CUBE = AV_CHANNEL_LAYOUT_MASK(8,  AV_CH_LAYOUT_CUBE)
//const AV_CHANNEL_LAYOUT_5POINT1POINT4_BACK = AV_CHANNEL_LAYOUT_MASK(10, AV_CH_LAYOUT_5POINT1POINT4_BACK)
//const AV_CHANNEL_LAYOUT_7POINT1POINT2 = AV_CHANNEL_LAYOUT_MASK(10, AV_CH_LAYOUT_7POINT1POINT2)
//const AV_CHANNEL_LAYOUT_7POINT1POINT4_BACK = AV_CHANNEL_LAYOUT_MASK(12, AV_CH_LAYOUT_7POINT1POINT4_BACK)
//const AV_CHANNEL_LAYOUT_HEXADECAGONAL = AV_CHANNEL_LAYOUT_MASK(16, AV_CH_LAYOUT_HEXADECAGONAL)
//const AV_CHANNEL_LAYOUT_STEREO_DOWNMIX = AV_CHANNEL_LAYOUT_MASK(2,  AV_CH_LAYOUT_STEREO_DOWNMIX)
//const AV_CHANNEL_LAYOUT_22POINT2 = AV_CHANNEL_LAYOUT_MASK(24, AV_CH_LAYOUT_22POINT2)
//const AV_CHANNEL_LAYOUT_7POINT1_TOP_BACK = AV_CHANNEL_LAYOUT_5POINT1POINT2_BACK
//const AV_CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER = {   AV_CHANNEL_ORDER_AMBISONIC,          4,          { 0 },          NULL }


                               
                               

                   
                   

                    
                       

/**
 * @file
 * @ingroup lavu_audio_channels
 * Public libavutil channel layout APIs header.
 */


/**
 * @defgroup lavu_audio_channels Audio channels
 * @ingroup lavu_audio
 *
 * Audio channel layout utility functions
 *
 * @{
 */

type AVChannel C.enum_AVChannel

type AVChannelOrder C.enum_AVChannelOrder


/**
 * @defgroup channel_masks Audio channel masks
 *
 * A channel layout is a 64-bits integer with a bit set for every channel.
 * The number of bits set must be equal to the number of channels.
 * The value 0 means that the channel layout is not known.
 * @note this data structure is not powerful enough to handle channels
 * combinations that have the same channel multiple times, such as
 * dual-mono.
 *
 * @{
 */































                             
                                                                     
                                                                              
                                         
                                                                    
                               
      
                                                          
      

/**
 * @}
 * @defgroup channel_mask_c Audio channel layouts
 * @{
 * */






































type AVMatrixEncoding C.enum_AVMatrixEncoding

/**
 * @}
 */

/**
 * An AVChannelCustom defines a single channel within a custom order layout
 *
 * Unlike most structures in FFmpeg, sizeof(AVChannelCustom) is a part of the
 * public ABI.
 *
 * No new fields may be added to it without a major version bump.
 */
type AVChannelCustom C.struct_AVChannelCustom

/**
 * An AVChannelLayout holds information about the channel layout of audio data.
 *
 * A channel layout here is defined as a set of channels ordered in a specific
 * way (unless the channel order is AV_CHANNEL_ORDER_UNSPEC, in which case an
 * AVChannelLayout carries only the channel count).
 * All orders may be treated as if they were AV_CHANNEL_ORDER_UNSPEC by
 * ignoring everything but the channel count, as long as av_channel_layout_check()
 * considers they are valid.
 *
 * Unlike most structures in FFmpeg, sizeof(AVChannelLayout) is a part of the
 * public ABI and may be used by the caller. E.g. it may be allocated on stack
 * or embedded in caller-defined structs.
 *
 * AVChannelLayout can be initialized as follows:
 * - default initialization with {0}, followed by setting all used fields
 *   correctly;
 * - by assigning one of the predefined AV_CHANNEL_LAYOUT_* initializers;
 * - with a constructor function, such as av_channel_layout_default(),
 *   av_channel_layout_from_mask() or av_channel_layout_from_string().
 *
 * The channel layout must be unitialized with av_channel_layout_uninit()
 *
 * Copying an AVChannelLayout via assigning is forbidden,
 * av_channel_layout_copy() must be used instead (and its return value should
 * be checked)
 *
 * No new fields may be added to it without a major version bump, except for
 * new elements of the union fitting in sizeof(uint64_t).
 */
type AVChannelLayout C.struct_AVChannelLayout

/**
 * Macro to define native channel layouts
 *
 * @note This doesn't use designated initializers for compatibility with C++ 17 and older.
 */


/**
 * @name Common pre-defined channel layouts
 * @{
 */







































/** @} */



                             
   
                             
     
   

   
                                                                           
  
                                                         
                           
                                                                       
                                                         
                                                                         
                                                       
                                                                
                                                                 
                                    
                                                                      
                     
  
                                                   
  
                                                  
   
                    
                                                 

   
                                                                                  
  
                                                                               
                                         
  
                                                                   
                                                                    
                                                  
  
                                                              
                                                  
   
                    
                                                                                                 

   
                                            
                                                                 
  
                                                               
                                              
                                        
                                              
                                               
   
                    
                                                                                                     

   
                                                               
                                               
   
                    
                                                                                             

   
                                                       
                                              
   
                    
                                                               

   
                                                                
  
                                              
   
                    
                                                       

   
                                                
  
                                              
                                                                               
                                            
  
                                                                            
                    
  
                                                         
   
                    
                                                                
                                                          

   
                                                          
                                                         
   
                    
                                                                               

   
                                   
  
                                                  
  
                                    
   
                    
                                                  

   
                                          
  
                                                         
                                                         
                                           
   
                    
                                                         

   
                                                       
  
                                                               
                                          
                                         
                                    
                                            
                                               
   
                    
                                                                    
                                                      
   
     
   
      

/**
 * Get a human readable string in an abbreviated form describing a given channel.
 * This is the inverse function of @ref av_channel_from_string().
 *
 * @param buf pre-allocated buffer where to put the generated string
 * @param buf_size size in bytes of the buffer.
 * @param channel the AVChannel whose name to get
 * @return amount of bytes needed to hold the output string, or a negative AVERROR
 *         on failure. If the returned value is bigger than buf_size, then the
 *         string was truncated.
 */
func Av_channel_name(buf *byte, buf_size uint64, channel AVChannel) int32 {
    return int32(C.av_channel_name((*C.char)(unsafe.Pointer(buf)), 
        C.ulonglong(buf_size), C.enum_AVChannel(channel)))
}

/**
 * bprint variant of av_channel_name().
 *
 * @note the string will be appended to the bprint buffer.
 */
func Av_channel_name_bprint(bp *AVBPrint, channel_id AVChannel)  {
    C.av_channel_name_bprint((*C.struct_AVBPrint)(unsafe.Pointer(bp)), 
        C.enum_AVChannel(channel_id))
}

/**
 * Get a human readable string describing a given channel.
 *
 * @param buf pre-allocated buffer where to put the generated string
 * @param buf_size size in bytes of the buffer.
 * @param channel the AVChannel whose description to get
 * @return amount of bytes needed to hold the output string, or a negative AVERROR
 *         on failure. If the returned value is bigger than buf_size, then the
 *         string was truncated.
 */
func Av_channel_description(buf *byte, buf_size uint64, channel AVChannel) int32 {
    return int32(C.av_channel_description((*C.char)(unsafe.Pointer(buf)), 
        C.ulonglong(buf_size), C.enum_AVChannel(channel)))
}

/**
 * bprint variant of av_channel_description().
 *
 * @note the string will be appended to the bprint buffer.
 */
func Av_channel_description_bprint(bp *AVBPrint, channel_id AVChannel)  {
    C.av_channel_description_bprint((*C.struct_AVBPrint)(unsafe.Pointer(bp)), 
        C.enum_AVChannel(channel_id))
}

/**
 * This is the inverse function of @ref av_channel_name().
 *
 * @return the channel with the given name
 *         AV_CHAN_NONE when name does not identify a known channel
 */
func Av_channel_from_string(name *byte) AVChannel {
    return AVChannel(C.av_channel_from_string((*C.char)(unsafe.Pointer(name))))
}

/**
 * Initialize a native channel layout from a bitmask indicating which channels
 * are present.
 *
 * @param channel_layout the layout structure to be initialized
 * @param mask bitmask describing the channel layout
 *
 * @return 0 on success
 *         AVERROR(EINVAL) for invalid mask values
 */
func Av_channel_layout_from_mask(channel_layout *AVChannelLayout, mask uint64) int32 {
    return int32(C.av_channel_layout_from_mask(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), C.ulonglong(mask)))
}

/**
 * Initialize a channel layout from a given string description.
 * The input string can be represented by:
 *  - the formal channel layout name (returned by av_channel_layout_describe())
 *  - single or multiple channel names (returned by av_channel_name(), eg. "FL",
 *    or concatenated with "+", each optionally containing a custom name after
 *    a "@", eg. "FL@Left+FR@Right+LFE")
 *  - a decimal or hexadecimal value of a native channel layout (eg. "4" or "0x4")
 *  - the number of channels with default layout (eg. "4c")
 *  - the number of unordered channels (eg. "4C" or "4 channels")
 *  - the ambisonic order followed by optional non-diegetic channels (eg.
 *    "ambisonic 2+stereo")
 *
 * @param channel_layout input channel layout
 * @param str string describing the channel layout
 * @return 0 channel layout was detected, AVERROR_INVALIDATATA otherwise
 */
func Av_channel_layout_from_string(channel_layout *AVChannelLayout,
                                  str *byte) int32 {
    return int32(C.av_channel_layout_from_string(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        (*C.char)(unsafe.Pointer(str))))
}

/**
 * Get the default channel layout for a given number of channels.
 *
 * @param ch_layout the layout structure to be initialized
 * @param nb_channels number of channels
 */
func Av_channel_layout_default(ch_layout *AVChannelLayout, nb_channels int32)  {
    C.av_channel_layout_default((*C.AVChannelLayout)(unsafe.Pointer(ch_layout)), 
        C.int(nb_channels))
}

/**
 * Iterate over all standard channel layouts.
 *
 * @param opaque a pointer where libavutil will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the standard channel layout or NULL when the iteration is
 *         finished
 */
func Av_channel_layout_standard(opaque *unsafe.Pointer) *AVChannelLayout {
    return (*AVChannelLayout)(unsafe.Pointer(C.av_channel_layout_standard(opaque)))
}

/**
 * Free any allocated data in the channel layout and reset the channel
 * count to 0.
 *
 * @param channel_layout the layout structure to be uninitialized
 */
func Av_channel_layout_uninit(channel_layout *AVChannelLayout)  {
    C.av_channel_layout_uninit(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)))
}

/**
 * Make a copy of a channel layout. This differs from just assigning src to dst
 * in that it allocates and copies the map for AV_CHANNEL_ORDER_CUSTOM.
 *
 * @note the destination channel_layout will be always uninitialized before copy.
 *
 * @param dst destination channel layout
 * @param src source channel layout
 * @return 0 on success, a negative AVERROR on error.
 */
func Av_channel_layout_copy(dst *AVChannelLayout, src *AVChannelLayout) int32 {
    return int32(C.av_channel_layout_copy(
        (*C.AVChannelLayout)(unsafe.Pointer(dst)), 
        (*C.AVChannelLayout)(unsafe.Pointer(src))))
}

/**
 * Get a human-readable string describing the channel layout properties.
 * The string will be in the same format that is accepted by
 * @ref av_channel_layout_from_string(), allowing to rebuild the same
 * channel layout, except for opaque pointers.
 *
 * @param channel_layout channel layout to be described
 * @param buf pre-allocated buffer where to put the generated string
 * @param buf_size size in bytes of the buffer.
 * @return amount of bytes needed to hold the output string, or a negative AVERROR
 *         on failure. If the returned value is bigger than buf_size, then the
 *         string was truncated.
 */
func Av_channel_layout_describe(channel_layout *AVChannelLayout,
                               buf *byte, buf_size uint64) int32 {
    return int32(C.av_channel_layout_describe(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        (*C.char)(unsafe.Pointer(buf)), C.ulonglong(buf_size)))
}

/**
 * bprint variant of av_channel_layout_describe().
 *
 * @note the string will be appended to the bprint buffer.
 * @return 0 on success, or a negative AVERROR value on failure.
 */
func Av_channel_layout_describe_bprint(channel_layout *AVChannelLayout,
                                      bp *AVBPrint) int32 {
    return int32(C.av_channel_layout_describe_bprint(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        (*C.struct_AVBPrint)(unsafe.Pointer(bp))))
}

/**
 * Get the channel with the given index in a channel layout.
 *
 * @param channel_layout input channel layout
 * @param idx index of the channel
 * @return channel with the index idx in channel_layout on success or
 *         AV_CHAN_NONE on failure (if idx is not valid or the channel order is
 *         unspecified)
 */
func
Av_channel_layout_channel_from_index(channel_layout *AVChannelLayout, idx uint32) AVChannel {
    return AVChannel(C.av_channel_layout_channel_from_index(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), C.uint(idx)))
}

/**
 * Get the index of a given channel in a channel layout. In case multiple
 * channels are found, only the first match will be returned.
 *
 * @param channel_layout input channel layout
 * @param channel the channel whose index to obtain
 * @return index of channel in channel_layout on success or a negative number if
 *         channel is not present in channel_layout.
 */
func Av_channel_layout_index_from_channel(channel_layout *AVChannelLayout,
                                         channel AVChannel) int32 {
    return int32(C.av_channel_layout_index_from_channel(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        C.enum_AVChannel(channel)))
}

/**
 * Get the index in a channel layout of a channel described by the given string.
 * In case multiple channels are found, only the first match will be returned.
 *
 * This function accepts channel names in the same format as
 * @ref av_channel_from_string().
 *
 * @param channel_layout input channel layout
 * @param name string describing the channel whose index to obtain
 * @return a channel index described by the given string, or a negative AVERROR
 *         value.
 */
func Av_channel_layout_index_from_string(channel_layout *AVChannelLayout,
                                        name *byte) int32 {
    return int32(C.av_channel_layout_index_from_string(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        (*C.char)(unsafe.Pointer(name))))
}

/**
 * Get a channel described by the given string.
 *
 * This function accepts channel names in the same format as
 * @ref av_channel_from_string().
 *
 * @param channel_layout input channel layout
 * @param name string describing the channel to obtain
 * @return a channel described by the given string in channel_layout on success
 *         or AV_CHAN_NONE on failure (if the string is not valid or the channel
 *         order is unspecified)
 */
func
Av_channel_layout_channel_from_string(channel_layout *AVChannelLayout,
                                      name *byte) AVChannel {
    return AVChannel(C.av_channel_layout_channel_from_string(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), 
        (*C.char)(unsafe.Pointer(name))))
}

/**
 * Find out what channels from a given set are present in a channel layout,
 * without regard for their positions.
 *
 * @param channel_layout input channel layout
 * @param mask a combination of AV_CH_* representing a set of channels
 * @return a bitfield representing all the channels from mask that are present
 *         in channel_layout
 */
func Av_channel_layout_subset(channel_layout *AVChannelLayout,
                                  mask uint64) uint64 {
    return uint64(C.av_channel_layout_subset(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout)), C.ulonglong(mask)))
}

/**
 * Check whether a channel layout is valid, i.e. can possibly describe audio
 * data.
 *
 * @param channel_layout input channel layout
 * @return 1 if channel_layout is valid, 0 otherwise.
 */
func Av_channel_layout_check(channel_layout *AVChannelLayout) int32 {
    return int32(C.av_channel_layout_check(
        (*C.AVChannelLayout)(unsafe.Pointer(channel_layout))))
}

/**
 * Check whether two channel layouts are semantically the same, i.e. the same
 * channels are present on the same positions in both.
 *
 * If one of the channel layouts is AV_CHANNEL_ORDER_UNSPEC, while the other is
 * not, they are considered to be unequal. If both are AV_CHANNEL_ORDER_UNSPEC,
 * they are considered equal iff the channel counts are the same in both.
 *
 * @param chl input channel layout
 * @param chl1 input channel layout
 * @return 0 if chl and chl1 are equal, 1 if they are not equal. A negative
 *         AVERROR code if one or both are invalid.
 */
func Av_channel_layout_compare(chl *AVChannelLayout, chl1 *AVChannelLayout) int32 {
    return int32(C.av_channel_layout_compare(
        (*C.AVChannelLayout)(unsafe.Pointer(chl)), 
        (*C.AVChannelLayout)(unsafe.Pointer(chl1))))
}

/**
 * @}
 */

                                    
