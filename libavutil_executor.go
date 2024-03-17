// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2023 Nuo Mi
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
//#include "libavutil/executor.h"
import "C"
import (
    "unsafe"
)



                         
                         

type AVExecutor struct {
}

// type AVTask

type AVTask struct {
    Next *AVTask
}


type AVTaskCallbacks struct {
    User_data unsafe.Pointer
    Local_context_size int32
    Priority_higher uintptr
    Ready uintptr
    Run uintptr
}


/**
 * Alloc executor
 * @param callbacks callback structure for executor
 * @param thread_count worker thread number
 * @return return the executor
 */
func Av_executor_alloc(callbacks *AVTaskCallbacks, thread_count int32) *AVExecutor {
    return (*AVExecutor)(unsafe.Pointer(C.av_executor_alloc(
        (*C.struct_AVTaskCallbacks)(unsafe.Pointer(callbacks)), 
        C.int(thread_count))))
}

/**
 * Free executor
 * @param e  pointer to executor
 */
func Av_executor_free(e **AVExecutor)  {
    C.av_executor_free((**C.struct_AVExecutor)(unsafe.Pointer(e)))
}

/**
 * Add task to executor
 * @param e pointer to executor
 * @param t pointer to task. If NULL, it will wakeup one work thread
 */
func Av_executor_execute(e *AVExecutor, t *AVTask)  {
    C.av_executor_execute((*C.struct_AVExecutor)(unsafe.Pointer(e)), 
        (*C.struct_AVTask)(unsafe.Pointer(t)))
}

                          
