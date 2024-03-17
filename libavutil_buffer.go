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
//#include <stddef.h>
//#include <stdint.h>
//#include "libavutil/buffer.h"
import "C"
import (
    "syscall"
    "unsafe"
)

const AV_BUFFER_FLAG_READONLY =  (1 << 0) 


/**
 * @file
 * @ingroup lavu_buffer
 * refcounted data buffer API
 */

                       
                       

                   
                   

/**
 * @defgroup lavu_buffer AVBuffer
 * @ingroup lavu_data
 *
 * @{
 * AVBuffer is an API for reference-counted data buffers.
 *
 * There are two core objects in this API -- AVBuffer and AVBufferRef. AVBuffer
 * represents the data buffer itself; it is opaque and not meant to be accessed
 * by the caller directly, but only through AVBufferRef. However, the caller may
 * e.g. compare two AVBuffer pointers to check whether two different references
 * are describing the same data buffer. AVBufferRef represents a single
 * reference to an AVBuffer and it is the object that may be manipulated by the
 * caller directly.
 *
 * There are two functions provided for creating a new AVBuffer with a single
 * reference -- av_buffer_alloc() to just allocate a new buffer, and
 * av_buffer_create() to wrap an existing array in an AVBuffer. From an existing
 * reference, additional references may be created with av_buffer_ref().
 * Use av_buffer_unref() to free a reference (this will automatically free the
 * data once all the references are freed).
 *
 * The convention throughout this API and the rest of FFmpeg is such that the
 * buffer is considered writable if there exists only one reference to it (and
 * it has not been marked as read-only). The av_buffer_is_writable() function is
 * provided to check whether this is true and av_buffer_make_writable() will
 * automatically create a new writable buffer when necessary.
 * Of course nothing prevents the calling code from violating this convention,
 * however that is safe only when all the existing references are under its
 * control.
 *
 * @note Referencing and unreferencing the buffers is thread-safe and thus
 * may be done from multiple threads simultaneously without any need for
 * additional locking.
 *
 * @note Two different references to the same buffer can point to different
 * parts of the buffer (i.e. their AVBufferRef.data will not be equal).
 */

/**
 * A reference counted buffer type. It is opaque and is meant to be used through
 * references (AVBufferRef).
 */
type AVBuffer struct {
}


/**
 * A reference to a data buffer.
 *
 * The size of this struct is not a part of the public ABI and it is not meant
 * to be allocated directly.
 */
type AVBufferRef struct {
    Buffer *AVBuffer
    Data *uint8
    Size uint64
}


/**
 * Allocate an AVBuffer of the given size using av_malloc().
 *
 * @return an AVBufferRef of given size or NULL when out of memory
 */
func Av_buffer_alloc(size uint64) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffer_alloc(C.ulonglong(size))))
}

/**
 * Same as av_buffer_alloc(), except the returned buffer will be initialized
 * to zero.
 */
func Av_buffer_allocz(size uint64) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffer_allocz(C.ulonglong(size))))
}

/**
 * Always treat the buffer as read-only, even when it has only one
 * reference.
 */


/**
 * Create an AVBuffer from an existing array.
 *
 * If this function is successful, data is owned by the AVBuffer. The caller may
 * only access data through the returned AVBufferRef and references derived from
 * it.
 * If this function fails, data is left untouched.
 * @param data   data array
 * @param size   size of data in bytes
 * @param free   a callback for freeing this buffer's data
 * @param opaque parameter to be got for processing or passed to free
 * @param flags  a combination of AV_BUFFER_FLAG_*
 *
 * @return an AVBufferRef referring to data on success, NULL on failure.
 */
func Av_buffer_create(data *uint8, size uint64,
                              free func(opaque unsafe.Pointer, data *uint8) ,
                              opaque unsafe.Pointer, flags int32) *AVBufferRef {
    cb2 := syscall.NewCallbackCDecl(free)
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffer_create(
        (*C.uchar)(unsafe.Pointer(data)), C.ulonglong(size), 
        (*[0]byte)(unsafe.Pointer(cb2)), opaque, C.int(flags))))
}

/**
 * Default free callback, which calls av_free() on the buffer data.
 * This function is meant to be passed to av_buffer_create(), not called
 * directly.
 */
func Av_buffer_default_free(opaque unsafe.Pointer, data *uint8)  {
    C.av_buffer_default_free(opaque, (*C.uchar)(unsafe.Pointer(data)))
}

/**
 * Create a new reference to an AVBuffer.
 *
 * @return a new AVBufferRef referring to the same AVBuffer as buf or NULL on
 * failure.
 */
func Av_buffer_ref(buf *AVBufferRef) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffer_ref(
        (*C.struct_AVBufferRef)(unsafe.Pointer(buf)))))
}

/**
 * Free a given reference and automatically free the buffer if there are no more
 * references to it.
 *
 * @param buf the reference to be freed. The pointer is set to NULL on return.
 */
func Av_buffer_unref(buf **AVBufferRef)  {
    C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(buf)))
}

/**
 * @return 1 if the caller may write to the data referred to by buf (which is
 * true if and only if buf is the only reference to the underlying AVBuffer).
 * Return 0 otherwise.
 * A positive answer is valid until av_buffer_ref() is called on buf.
 */
func Av_buffer_is_writable(buf *AVBufferRef) int32 {
    return int32(C.av_buffer_is_writable(
        (*C.struct_AVBufferRef)(unsafe.Pointer(buf))))
}

/**
 * @return the opaque parameter set by av_buffer_create.
 */
func Av_buffer_get_opaque(buf *AVBufferRef) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.av_buffer_get_opaque(
        (*C.struct_AVBufferRef)(unsafe.Pointer(buf)))))
}

func Av_buffer_get_ref_count(buf *AVBufferRef) int32 {
    return int32(C.av_buffer_get_ref_count(
        (*C.struct_AVBufferRef)(unsafe.Pointer(buf))))
}

/**
 * Create a writable reference from a given buffer reference, avoiding data copy
 * if possible.
 *
 * @param buf buffer reference to make writable. On success, buf is either left
 *            untouched, or it is unreferenced and a new writable AVBufferRef is
 *            written in its place. On failure, buf is left untouched.
 * @return 0 on success, a negative AVERROR on failure.
 */
func Av_buffer_make_writable(buf **AVBufferRef) int32 {
    return int32(C.av_buffer_make_writable(
        (**C.struct_AVBufferRef)(unsafe.Pointer(buf))))
}

/**
 * Reallocate a given buffer.
 *
 * @param buf  a buffer reference to reallocate. On success, buf will be
 *             unreferenced and a new reference with the required size will be
 *             written in its place. On failure buf will be left untouched. *buf
 *             may be NULL, then a new buffer is allocated.
 * @param size required new buffer size.
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note the buffer is actually reallocated with av_realloc() only if it was
 * initially allocated through av_buffer_realloc(NULL) and there is only one
 * reference to it (i.e. the one passed to this function). In all other cases
 * a new buffer is allocated and the data is copied.
 */
func Av_buffer_realloc(buf **AVBufferRef, size uint64) int32 {
    return int32(C.av_buffer_realloc((**C.struct_AVBufferRef)(unsafe.Pointer(buf)), 
        C.ulonglong(size)))
}

/**
 * Ensure dst refers to the same data as src.
 *
 * When *dst is already equivalent to src, do nothing. Otherwise unreference dst
 * and replace it with a new reference to src.
 *
 * @param dst Pointer to either a valid buffer reference or NULL. On success,
 *            this will point to a buffer reference equivalent to src. On
 *            failure, dst will be left untouched.
 * @param src A buffer reference to replace dst with. May be NULL, then this
 *            function is equivalent to av_buffer_unref(dst).
 * @return 0 on success
 *         AVERROR(ENOMEM) on memory allocation failure.
 */
func Av_buffer_replace(dst **AVBufferRef, src *AVBufferRef) int32 {
    return int32(C.av_buffer_replace((**C.struct_AVBufferRef)(unsafe.Pointer(dst)), 
        (*C.struct_AVBufferRef)(unsafe.Pointer(src))))
}

/**
 * @}
 */

/**
 * @defgroup lavu_bufferpool AVBufferPool
 * @ingroup lavu_data
 *
 * @{
 * AVBufferPool is an API for a lock-free thread-safe pool of AVBuffers.
 *
 * Frequently allocating and freeing large buffers may be slow. AVBufferPool is
 * meant to solve this in cases when the caller needs a set of buffers of the
 * same size (the most obvious use case being buffers for raw video or audio
 * frames).
 *
 * At the beginning, the user must call av_buffer_pool_init() to create the
 * buffer pool. Then whenever a buffer is needed, call av_buffer_pool_get() to
 * get a reference to a new buffer, similar to av_buffer_alloc(). This new
 * reference works in all aspects the same way as the one created by
 * av_buffer_alloc(). However, when the last reference to this buffer is
 * unreferenced, it is returned to the pool instead of being freed and will be
 * reused for subsequent av_buffer_pool_get() calls.
 *
 * When the caller is done with the pool and no longer needs to allocate any new
 * buffers, av_buffer_pool_uninit() must be called to mark the pool as freeable.
 * Once all the buffers are released, it will automatically be freed.
 *
 * Allocating and releasing buffers with this API is thread-safe as long as
 * either the default alloc callback is used, or the user-supplied one is
 * thread-safe.
 */

/**
 * The buffer pool. This structure is opaque and not meant to be accessed
 * directly. It is allocated with av_buffer_pool_init() and freed with
 * av_buffer_pool_uninit().
 */
type AVBufferPool struct {
}


/**
 * Allocate and initialize a buffer pool.
 *
 * @param size size of each buffer in this pool
 * @param alloc a function that will be used to allocate new buffers when the
 * pool is empty. May be NULL, then the default allocator will be used
 * (av_buffer_alloc()).
 * @return newly created buffer pool on success, NULL on error.
 */
func Av_buffer_pool_init(size uint64, alloc func(size uint64) **AVBufferRef) *AVBufferPool {
    cb1 := syscall.NewCallbackCDecl(alloc)
    return (*AVBufferPool)(unsafe.Pointer(C.av_buffer_pool_init(C.ulonglong(size), 
        (*[0]byte)(unsafe.Pointer(cb1)))))
}

/**
 * Allocate and initialize a buffer pool with a more complex allocator.
 *
 * @param size size of each buffer in this pool
 * @param opaque arbitrary user data used by the allocator
 * @param alloc a function that will be used to allocate new buffers when the
 *              pool is empty. May be NULL, then the default allocator will be
 *              used (av_buffer_alloc()).
 * @param pool_free a function that will be called immediately before the pool
 *                  is freed. I.e. after av_buffer_pool_uninit() is called
 *                  by the caller and all the frames are returned to the pool
 *                  and freed. It is intended to uninitialize the user opaque
 *                  data. May be NULL.
 * @return newly created buffer pool on success, NULL on error.
 */
func Av_buffer_pool_init2(size uint64, opaque unsafe.Pointer,
                                   alloc func(opaque unsafe.Pointer, size uint64) **AVBufferRef,
                                   pool_free func(opaque unsafe.Pointer) ) *AVBufferPool {
    cb2 := syscall.NewCallbackCDecl(alloc)
    cb3 := syscall.NewCallbackCDecl(pool_free)
    return (*AVBufferPool)(unsafe.Pointer(C.av_buffer_pool_init2(C.ulonglong(size), 
        opaque, (*[0]byte)(unsafe.Pointer(cb2)), (*[0]byte)(unsafe.Pointer(cb3)))))
}

/**
 * Mark the pool as being available for freeing. It will actually be freed only
 * once all the allocated buffers associated with the pool are released. Thus it
 * is safe to call this function while some of the allocated buffers are still
 * in use.
 *
 * @param pool pointer to the pool to be freed. It will be set to NULL.
 */
func Av_buffer_pool_uninit(pool **AVBufferPool)  {
    C.av_buffer_pool_uninit((**C.struct_AVBufferPool)(unsafe.Pointer(pool)))
}

/**
 * Allocate a new AVBuffer, reusing an old buffer from the pool when available.
 * This function may be called simultaneously from multiple threads.
 *
 * @return a reference to the new buffer on success, NULL on error.
 */
func Av_buffer_pool_get(pool *AVBufferPool) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_buffer_pool_get(
        (*C.struct_AVBufferPool)(unsafe.Pointer(pool)))))
}

/**
 * Query the original opaque parameter of an allocated buffer in the pool.
 *
 * @param ref a buffer reference to a buffer returned by av_buffer_pool_get.
 * @return the opaque parameter set by the buffer allocator function of the
 *         buffer pool.
 *
 * @note the opaque parameter of ref is used by the buffer pool implementation,
 * therefore you have to use this function to access the original opaque
 * parameter of an allocated buffer.
 */
func Av_buffer_pool_buffer_get_opaque(ref *AVBufferRef) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.av_buffer_pool_buffer_get_opaque(
        (*C.struct_AVBufferRef)(unsafe.Pointer(ref)))))
}

/**
 * @}
 */

                            
