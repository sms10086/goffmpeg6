/*
 * Audio FIFO
 * Copyright (c) 2012 Justin Ruggles <justin.ruggles@gmail.com>
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
//#include "libavutil/attributes.h"
//#include "libavutil/samplefmt.h"
//#include "libavutil/audio_fifo.h"
import "C"
import (
    "unsafe"
)



/**
 * @file
 * Audio FIFO Buffer
 */

                           
                           

                       
                      

/**
 * @addtogroup lavu_audio
 * @{
 *
 * @defgroup lavu_audiofifo Audio FIFO Buffer
 * @{
 */

/**
 * Context for an Audio FIFO Buffer.
 *
 * - Operates at the sample level rather than the byte level.
 * - Supports multiple channels with either planar or packed sample format.
 * - Automatic reallocation when writing to a full buffer.
 */
type AVAudioFifo C.struct_AVAudioFifo

/**
 * Free an AVAudioFifo.
 *
 * @param af  AVAudioFifo to free
 */
func Av_audio_fifo_free(af *AVAudioFifo)  {
    C.av_audio_fifo_free((*C.AVAudioFifo)(unsafe.Pointer(af)))
}

/**
 * Allocate an AVAudioFifo.
 *
 * @param sample_fmt  sample format
 * @param channels    number of channels
 * @param nb_samples  initial allocation size, in samples
 * @return            newly allocated AVAudioFifo, or NULL on error
 */
func Av_audio_fifo_alloc(sample_fmt AVSampleFormat, channels int32,
                                 nb_samples int32) *AVAudioFifo {
    return (*AVAudioFifo)(unsafe.Pointer(C.av_audio_fifo_alloc(C.enum_AVSampleFormat(sample_fmt), 
        C.int(channels), C.int(nb_samples))))
}

/**
 * Reallocate an AVAudioFifo.
 *
 * @param af          AVAudioFifo to reallocate
 * @param nb_samples  new allocation size, in samples
 * @return            0 if OK, or negative AVERROR code on failure
 */
func Av_audio_fifo_realloc(af *AVAudioFifo, nb_samples int32) int32 {
    return int32(C.av_audio_fifo_realloc((*C.AVAudioFifo)(unsafe.Pointer(af)), 
        C.int(nb_samples)))
}

/**
 * Write data to an AVAudioFifo.
 *
 * The AVAudioFifo will be reallocated automatically if the available space
 * is less than nb_samples.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to write to
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to write
 * @return            number of samples actually written, or negative AVERROR
 *                    code on failure. If successful, the number of samples
 *                    actually written will always be nb_samples.
 */
func Av_audio_fifo_write(af *AVAudioFifo, data *unsafe.Pointer, nb_samples int32) int32 {
    return int32(C.av_audio_fifo_write((*C.AVAudioFifo)(unsafe.Pointer(af)), 
        data, C.int(nb_samples)))
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func Av_audio_fifo_peek(af *AVAudioFifo, data *unsafe.Pointer, nb_samples int32) int32 {
    return int32(C.av_audio_fifo_peek((*C.AVAudioFifo)(unsafe.Pointer(af)), data, 
        C.int(nb_samples)))
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @param offset      offset from current read position
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func Av_audio_fifo_peek_at(af *AVAudioFifo, data *unsafe.Pointer,
                          nb_samples int32, offset int32) int32 {
    return int32(C.av_audio_fifo_peek_at((*C.AVAudioFifo)(unsafe.Pointer(af)), 
        data, C.int(nb_samples), C.int(offset)))
}

/**
 * Read data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to read
 * @return            number of samples actually read, or negative AVERROR code
 *                    on failure. The number of samples actually read will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func Av_audio_fifo_read(af *AVAudioFifo, data *unsafe.Pointer, nb_samples int32) int32 {
    return int32(C.av_audio_fifo_read((*C.AVAudioFifo)(unsafe.Pointer(af)), data, 
        C.int(nb_samples)))
}

/**
 * Drain data from an AVAudioFifo.
 *
 * Removes the data without reading it.
 *
 * @param af          AVAudioFifo to drain
 * @param nb_samples  number of samples to drain
 * @return            0 if OK, or negative AVERROR code on failure
 */
func Av_audio_fifo_drain(af *AVAudioFifo, nb_samples int32) int32 {
    return int32(C.av_audio_fifo_drain((*C.AVAudioFifo)(unsafe.Pointer(af)), 
        C.int(nb_samples)))
}

/**
 * Reset the AVAudioFifo buffer.
 *
 * This empties all data in the buffer.
 *
 * @param af  AVAudioFifo to reset
 */
func Av_audio_fifo_reset(af *AVAudioFifo)  {
    C.av_audio_fifo_reset((*C.AVAudioFifo)(unsafe.Pointer(af)))
}

/**
 * Get the current number of samples in the AVAudioFifo available for reading.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for reading
 */
func Av_audio_fifo_size(af *AVAudioFifo) int32 {
    return int32(C.av_audio_fifo_size((*C.AVAudioFifo)(unsafe.Pointer(af))))
}

/**
 * Get the current number of samples in the AVAudioFifo available for writing.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for writing
 */
func Av_audio_fifo_space(af *AVAudioFifo) int32 {
    return int32(C.av_audio_fifo_space((*C.AVAudioFifo)(unsafe.Pointer(af))))
}

/**
 * @}
 * @}
 */

                                
