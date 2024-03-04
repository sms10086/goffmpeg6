// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public License
 * as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with FFmpeg; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

package goffmpeg6

//#cgo pkg-config: libavutil
//#include "libavutil/threadmessage.h"
import "C"
import (
    "syscall"
    "unsafe"
)



                              
                              

type AVThreadMessageQueue C.struct_AVThreadMessageQueue

type AVThreadMessageFlags C.enum_AVThreadMessageFlags

/**
 * Allocate a new message queue.
 *
 * @param mq      pointer to the message queue
 * @param nelem   maximum number of elements in the queue
 * @param elsize  size of each element in the queue
 * @return  >=0 for success; <0 for error, in particular AVERROR(ENOSYS) if
 *          lavu was built without thread support
 */
func Av_thread_message_queue_alloc(mq **AVThreadMessageQueue,
                                  nelem uint32,
                                  elsize uint32) int32 {
    return int32(C.av_thread_message_queue_alloc(
        (**C.AVThreadMessageQueue)(unsafe.Pointer(mq)), C.unsigned(nelem), 
        C.unsigned(elsize)))
}

/**
 * Free a message queue.
 *
 * The message queue must no longer be in use by another thread.
 */
func Av_thread_message_queue_free(mq **AVThreadMessageQueue)  {
    C.av_thread_message_queue_free(
        (**C.AVThreadMessageQueue)(unsafe.Pointer(mq)))
}

/**
 * Send a message on the queue.
 */
func Av_thread_message_queue_send(mq *AVThreadMessageQueue,
                                 msg unsafe.Pointer,
                                 flags uint32) int32 {
    return int32(C.av_thread_message_queue_send(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq)), msg, C.unsigned(flags)))
}

/**
 * Receive a message from the queue.
 */
func Av_thread_message_queue_recv(mq *AVThreadMessageQueue,
                                 msg unsafe.Pointer,
                                 flags uint32) int32 {
    return int32(C.av_thread_message_queue_recv(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq)), msg, C.unsigned(flags)))
}

/**
 * Set the sending error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_send() will
 * return it immediately. Conventional values, such as AVERROR_EOF or
 * AVERROR(EAGAIN), can be used to cause the sending thread to stop or
 * suspend its operation.
 */
func Av_thread_message_queue_set_err_send(mq *AVThreadMessageQueue,
                                          err int32)  {
    C.av_thread_message_queue_set_err_send(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq)), C.int(err))
}

/**
 * Set the receiving error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_recv() will
 * return it immediately when there are no longer available messages.
 * Conventional values, such as AVERROR_EOF or AVERROR(EAGAIN), can be used
 * to cause the receiving thread to stop or suspend its operation.
 */
func Av_thread_message_queue_set_err_recv(mq *AVThreadMessageQueue,
                                          err int32)  {
    C.av_thread_message_queue_set_err_recv(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq)), C.int(err))
}

/**
 * Set the optional free message callback function which will be called if an
 * operation is removing messages from the queue.
 */
func Av_thread_message_queue_set_free_func(mq *AVThreadMessageQueue,
                                           free_func func(msg unsafe.Pointer) )  {
    cb1 := syscall.NewCallbackCDecl(free_func)
    C.av_thread_message_queue_set_free_func(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq)), 
        (*[0]byte)(unsafe.Pointer(cb1)))
}

/**
 * Return the current number of messages in the queue.
 *
 * @return the current number of messages or AVERROR(ENOSYS) if lavu was built
 *         without thread support
 */
func Av_thread_message_queue_nb_elems(mq *AVThreadMessageQueue) int32 {
    return int32(C.av_thread_message_queue_nb_elems(
        (*C.AVThreadMessageQueue)(unsafe.Pointer(mq))))
}

/**
 * Flush the message queue
 *
 * This function is mostly equivalent to reading and free-ing every message
 * except that it will be done in a single operation (no lock/unlock between
 * reads).
 */
func Av_thread_message_flush(mq *AVThreadMessageQueue)  {
    C.av_thread_message_flush((*C.AVThreadMessageQueue)(unsafe.Pointer(mq)))
}

                                   
