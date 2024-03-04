// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2006 Smartjog S.A.S, Baptiste Coudurier <baptiste.coudurier@gmail.com>
 * Copyright (c) 2011-2012 Smartjog S.A.S, Clément Bœsch <clement.boesch@smartjog.com>
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
//#include "libavutil/rational.h"
//#include "libavutil/timecode.h"
import "C"
import (
    "unsafe"
)

const AV_TIMECODE_STR_SIZE = 23


/**
 * @file
 * Timecode helpers header
 */

                         
                         

                   
                     



type AVTimecodeFlag C.enum_AVTimecodeFlag

type AVTimecode C.struct_AVTimecode

/**
 * Adjust frame number for NTSC drop frame time code.
 *
 * @param framenum frame number to adjust
 * @param fps      frame per second, multiples of 30
 * @return         adjusted frame number
 * @warning        adjustment is only valid for multiples of NTSC 29.97
 */
func Av_timecode_adjust_ntsc_framenum2(framenum int32, fps int32) int32 {
    return int32(C.av_timecode_adjust_ntsc_framenum2(C.int(framenum), C.int(fps)))
}

/**
 * Convert frame number to SMPTE 12M binary representation.
 *
 * @param tc       timecode data correctly initialized
 * @param framenum frame number
 * @return         the SMPTE binary representation
 *
 * See SMPTE ST 314M-2005 Sec 4.4.2.2.1 "Time code pack (TC)"
 * the format description as follows:
 * bits 0-5:   hours, in BCD(6bits)
 * bits 6:     BGF1
 * bits 7:     BGF2 (NTSC) or FIELD (PAL)
 * bits 8-14:  minutes, in BCD(7bits)
 * bits 15:    BGF0 (NTSC) or BGF2 (PAL)
 * bits 16-22: seconds, in BCD(7bits)
 * bits 23:    FIELD (NTSC) or BGF0 (PAL)
 * bits 24-29: frames, in BCD(6bits)
 * bits 30:    drop  frame flag (0: non drop,    1: drop)
 * bits 31:    color frame flag (0: unsync mode, 1: sync mode)
 * @note BCD numbers (6 or 7 bits): 4 or 5 lower bits for units, 2 higher bits for tens.
 * @note Frame number adjustment is automatically done in case of drop timecode,
 *       you do NOT have to call av_timecode_adjust_ntsc_framenum2().
 * @note The frame number is relative to tc->start.
 * @note Color frame (CF) and binary group flags (BGF) bits are set to zero.
 */
func Av_timecode_get_smpte_from_framenum(tc *AVTimecode, framenum int32) uint32 {
    return uint32(C.av_timecode_get_smpte_from_framenum(
        (*C.AVTimecode)(unsafe.Pointer(tc)), C.int(framenum)))
}

/**
 * Convert sei info to SMPTE 12M binary representation.
 *
 * @param rate     frame rate in rational form
 * @param drop     drop flag
 * @param hh       hour
 * @param mm       minute
 * @param ss       second
 * @param ff       frame number
 * @return         the SMPTE binary representation
 */
func Av_timecode_get_smpte(rate AVRational, drop int32, hh int32, mm int32, ss int32, ff int32) uint32 {
    return uint32(C.av_timecode_get_smpte(C.AVRational(rate), C.int(drop), 
        C.int(hh), C.int(mm), C.int(ss), C.int(ff)))
}

/**
 * Load timecode string in buf.
 *
 * @param buf      destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tc       timecode data correctly initialized
 * @param framenum frame number
 * @return         the buf parameter
 *
 * @note Timecode representation can be a negative timecode and have more than
 *       24 hours, but will only be honored if the flags are correctly set.
 * @note The frame number is relative to tc->start.
 */
func Av_timecode_make_string(tc *AVTimecode, buf *byte, framenum int32) string {
    return C.GoString(C.av_timecode_make_string((*C.AVTimecode)(unsafe.Pointer(tc)), 
        (*C.char)(unsafe.Pointer(buf)), C.int(framenum)))
}

/**
 * Get the timecode string from the SMPTE timecode format.
 *
 * In contrast to av_timecode_make_smpte_tc_string this function supports 50/60
 * fps timecodes by using the field bit.
 *
 * @param buf        destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param rate       frame rate of the timecode
 * @param tcsmpte    the 32-bit SMPTE timecode
 * @param prevent_df prevent the use of a drop flag when it is known the DF bit
 *                   is arbitrary
 * @param skip_field prevent the use of a field flag when it is known the field
 *                   bit is arbitrary (e.g. because it is used as PC flag)
 * @return           the buf parameter
 */
func Av_timecode_make_smpte_tc_string2(buf *byte, rate AVRational, tcsmpte uint32, prevent_df int32, skip_field int32) string {
    return C.GoString(C.av_timecode_make_smpte_tc_string2(
        (*C.char)(unsafe.Pointer(buf)), C.AVRational(rate), C.uint(tcsmpte), 
        C.int(prevent_df), C.int(skip_field)))
}

/**
 * Get the timecode string from the SMPTE timecode format.
 *
 * @param buf        destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tcsmpte    the 32-bit SMPTE timecode
 * @param prevent_df prevent the use of a drop flag when it is known the DF bit
 *                   is arbitrary
 * @return           the buf parameter
 */
func Av_timecode_make_smpte_tc_string(buf *byte, tcsmpte uint32, prevent_df int32) string {
    return C.GoString(C.av_timecode_make_smpte_tc_string(
        (*C.char)(unsafe.Pointer(buf)), C.uint(tcsmpte), C.int(prevent_df)))
}

/**
 * Get the timecode string from the 25-bit timecode format (MPEG GOP format).
 *
 * @param buf     destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tc25bit the 25-bits timecode
 * @return        the buf parameter
 */
func Av_timecode_make_mpeg_tc_string(buf *byte, tc25bit uint32) string {
    return C.GoString(C.av_timecode_make_mpeg_tc_string(
        (*C.char)(unsafe.Pointer(buf)), C.uint(tc25bit)))
}

/**
 * Init a timecode struct with the passed parameters.
 *
 * @param log_ctx     a pointer to an arbitrary struct of which the first field
 *                    is a pointer to an AVClass struct (used for av_log)
 * @param tc          pointer to an allocated AVTimecode
 * @param rate        frame rate in rational form
 * @param flags       miscellaneous flags such as drop frame, +24 hours, ...
 *                    (see AVTimecodeFlag)
 * @param frame_start the first frame number
 * @return            0 on success, AVERROR otherwise
 */
func Av_timecode_init(tc *AVTimecode, rate AVRational, flags int32, frame_start int32, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_timecode_init((*C.AVTimecode)(unsafe.Pointer(tc)), 
        C.AVRational(rate), C.int(flags), C.int(frame_start), log_ctx))
}

/**
 * Init a timecode struct from the passed timecode components.
 *
 * @param log_ctx     a pointer to an arbitrary struct of which the first field
 *                    is a pointer to an AVClass struct (used for av_log)
 * @param tc          pointer to an allocated AVTimecode
 * @param rate        frame rate in rational form
 * @param flags       miscellaneous flags such as drop frame, +24 hours, ...
 *                    (see AVTimecodeFlag)
 * @param hh          hours
 * @param mm          minutes
 * @param ss          seconds
 * @param ff          frames
 * @return            0 on success, AVERROR otherwise
 */
func Av_timecode_init_from_components(tc *AVTimecode, rate AVRational, flags int32, hh int32, mm int32, ss int32, ff int32, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_timecode_init_from_components(
        (*C.AVTimecode)(unsafe.Pointer(tc)), C.AVRational(rate), C.int(flags), 
        C.int(hh), C.int(mm), C.int(ss), C.int(ff), log_ctx))
}

/**
 * Parse timecode representation (hh:mm:ss[:;.]ff).
 *
 * @param log_ctx a pointer to an arbitrary struct of which the first field is a
 *                pointer to an AVClass struct (used for av_log).
 * @param tc      pointer to an allocated AVTimecode
 * @param rate    frame rate in rational form
 * @param str     timecode string which will determine the frame start
 * @return        0 on success, AVERROR otherwise
 */
func Av_timecode_init_from_string(tc *AVTimecode, rate AVRational, str *byte, log_ctx unsafe.Pointer) int32 {
    return int32(C.av_timecode_init_from_string(
        (*C.AVTimecode)(unsafe.Pointer(tc)), C.AVRational(rate), 
        (*C.char)(unsafe.Pointer(str)), log_ctx))
}

/**
 * Check if the timecode feature is available for the given frame rate
 *
 * @return 0 if supported, <0 otherwise
 */
func Av_timecode_check_frame_rate(rate AVRational) int32 {
    return int32(C.av_timecode_check_frame_rate(C.AVRational(rate)))
}

                              
