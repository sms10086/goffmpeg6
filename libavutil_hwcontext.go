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
//#include "libavutil/buffer.h"
//#include "libavutil/frame.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/hwcontext.h"
import "C"
import (
    "unsafe"
)



                          
                          

                   
                  
                
                   

type AVHWDeviceType C.enum_AVHWDeviceType

type AVHWDeviceInternal C.struct_AVHWDeviceInternal

/**
 * This struct aggregates all the (hardware/vendor-specific) "high-level" state,
 * i.e. state that is not tied to a concrete processing configuration.
 * E.g., in an API that supports hardware-accelerated encoding and decoding,
 * this struct will (if possible) wrap the state that is common to both encoding
 * and decoding and from which specific instances of encoders or decoders can be
 * derived.
 *
 * This struct is reference-counted with the AVBuffer mechanism. The
 * av_hwdevice_ctx_alloc() constructor yields a reference, whose data field
 * points to the actual AVHWDeviceContext. Further objects derived from
 * AVHWDeviceContext (such as AVHWFramesContext, describing a frame pool with
 * specific properties) will hold an internal reference to it. After all the
 * references are released, the AVHWDeviceContext itself will be freed,
 * optionally invoking a user-specified callback for uninitializing the hardware
 * state.
 */
type AVHWDeviceContext C.struct_AVHWDeviceContext

type AVHWFramesInternal C.struct_AVHWFramesInternal

/**
 * This struct describes a set or pool of "hardware" frames (i.e. those with
 * data not located in normal system memory). All the frames in the pool are
 * assumed to be allocated in the same way and interchangeable.
 *
 * This struct is reference-counted with the AVBuffer mechanism and tied to a
 * given AVHWDeviceContext instance. The av_hwframe_ctx_alloc() constructor
 * yields a reference, whose data field points to the actual AVHWFramesContext
 * struct.
 */
type AVHWFramesContext C.struct_AVHWFramesContext

/**
 * Look up an AVHWDeviceType by name.
 *
 * @param name String name of the device type (case-insensitive).
 * @return The type from enum AVHWDeviceType, or AV_HWDEVICE_TYPE_NONE if
 *         not found.
 */
func Av_hwdevice_find_type_by_name(name *byte) AVHWDeviceType {
    return AVHWDeviceType(C.av_hwdevice_find_type_by_name(
        (*C.char)(unsafe.Pointer(name))))
}

/** Get the string name of an AVHWDeviceType.
 *
 * @param type Type from enum AVHWDeviceType.
 * @return Pointer to a static string containing the name, or NULL if the type
 *         is not valid.
 */
func Av_hwdevice_get_type_name(typex AVHWDeviceType) string {
    return C.GoString(C.av_hwdevice_get_type_name(C.enum_AVHWDeviceType(typex)))
}

/**
 * Iterate over supported device types.
 *
 * @param prev AV_HWDEVICE_TYPE_NONE initially, then the previous type
 *             returned by this function in subsequent iterations.
 * @return The next usable device type from enum AVHWDeviceType, or
 *         AV_HWDEVICE_TYPE_NONE if there are no more.
 */
func Av_hwdevice_iterate_types(prev AVHWDeviceType) AVHWDeviceType {
    return AVHWDeviceType(C.av_hwdevice_iterate_types(
        C.enum_AVHWDeviceType(prev)))
}

/**
 * Allocate an AVHWDeviceContext for a given hardware type.
 *
 * @param type the type of the hardware device to allocate.
 * @return a reference to the newly created AVHWDeviceContext on success or NULL
 *         on failure.
 */
func Av_hwdevice_ctx_alloc(typex AVHWDeviceType) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_hwdevice_ctx_alloc(C.enum_AVHWDeviceType(typex))))
}

/**
 * Finalize the device context before use. This function must be called after
 * the context is filled with all the required information and before it is
 * used in any way.
 *
 * @param ref a reference to the AVHWDeviceContext
 * @return 0 on success, a negative AVERROR code on failure
 */
func Av_hwdevice_ctx_init(ref *AVBufferRef) int32 {
    return int32(C.av_hwdevice_ctx_init((*C.AVBufferRef)(unsafe.Pointer(ref))))
}

/**
 * Open a device of the specified type and create an AVHWDeviceContext for it.
 *
 * This is a convenience function intended to cover the simple cases. Callers
 * who need to fine-tune device creation/management should open the device
 * manually and then wrap it in an AVHWDeviceContext using
 * av_hwdevice_ctx_alloc()/av_hwdevice_ctx_init().
 *
 * The returned context is already initialized and ready for use, the caller
 * should not call av_hwdevice_ctx_init() on it. The user_opaque/free fields of
 * the created AVHWDeviceContext are set by this function and should not be
 * touched by the caller.
 *
 * @param device_ctx On success, a reference to the newly-created device context
 *                   will be written here. The reference is owned by the caller
 *                   and must be released with av_buffer_unref() when no longer
 *                   needed. On failure, NULL will be written to this pointer.
 * @param type The type of the device to create.
 * @param device A type-specific string identifying the device to open.
 * @param opts A dictionary of additional (type-specific) options to use in
 *             opening the device. The dictionary remains owned by the caller.
 * @param flags currently unused
 *
 * @return 0 on success, a negative AVERROR code on failure.
 */
func Av_hwdevice_ctx_create(device_ctx **AVBufferRef, typex AVHWDeviceType,
                           device *byte, opts *AVDictionary, flags int32) int32 {
    return int32(C.av_hwdevice_ctx_create(
        (**C.AVBufferRef)(unsafe.Pointer(device_ctx)), 
        C.enum_AVHWDeviceType(typex), (*C.char)(unsafe.Pointer(device)), 
        (*C.AVDictionary)(unsafe.Pointer(opts)), C.int(flags)))
}

/**
 * Create a new device of the specified type from an existing device.
 *
 * If the source device is a device of the target type or was originally
 * derived from such a device (possibly through one or more intermediate
 * devices of other types), then this will return a reference to the
 * existing device of the same type as is requested.
 *
 * Otherwise, it will attempt to derive a new device from the given source
 * device.  If direct derivation to the new type is not implemented, it will
 * attempt the same derivation from each ancestor of the source device in
 * turn looking for an implemented derivation method.
 *
 * @param dst_ctx On success, a reference to the newly-created
 *                AVHWDeviceContext.
 * @param type    The type of the new device to create.
 * @param src_ctx A reference to an existing AVHWDeviceContext which will be
 *                used to create the new device.
 * @param flags   Currently unused; should be set to zero.
 * @return        Zero on success, a negative AVERROR code on failure.
 */
func Av_hwdevice_ctx_create_derived(dst_ctx **AVBufferRef,
                                   typex AVHWDeviceType,
                                   src_ctx *AVBufferRef, flags int32) int32 {
    return int32(C.av_hwdevice_ctx_create_derived(
        (**C.AVBufferRef)(unsafe.Pointer(dst_ctx)), C.enum_AVHWDeviceType(typex), 
        (*C.AVBufferRef)(unsafe.Pointer(src_ctx)), C.int(flags)))
}

/**
 * Create a new device of the specified type from an existing device.
 *
 * This function performs the same action as av_hwdevice_ctx_create_derived,
 * however, it is able to set options for the new device to be derived.
 *
 * @param dst_ctx On success, a reference to the newly-created
 *                AVHWDeviceContext.
 * @param type    The type of the new device to create.
 * @param src_ctx A reference to an existing AVHWDeviceContext which will be
 *                used to create the new device.
 * @param options Options for the new device to create, same format as in
 *                av_hwdevice_ctx_create.
 * @param flags   Currently unused; should be set to zero.
 * @return        Zero on success, a negative AVERROR code on failure.
 */
func Av_hwdevice_ctx_create_derived_opts(dst_ctx **AVBufferRef,
                                        typex AVHWDeviceType,
                                        src_ctx *AVBufferRef,
                                        options *AVDictionary, flags int32) int32 {
    return int32(C.av_hwdevice_ctx_create_derived_opts(
        (**C.AVBufferRef)(unsafe.Pointer(dst_ctx)), C.enum_AVHWDeviceType(typex), 
        (*C.AVBufferRef)(unsafe.Pointer(src_ctx)), 
        (*C.AVDictionary)(unsafe.Pointer(options)), C.int(flags)))
}

/**
 * Allocate an AVHWFramesContext tied to a given device context.
 *
 * @param device_ctx a reference to a AVHWDeviceContext. This function will make
 *                   a new reference for internal use, the one passed to the
 *                   function remains owned by the caller.
 * @return a reference to the newly created AVHWFramesContext on success or NULL
 *         on failure.
 */
func Av_hwframe_ctx_alloc(device_ctx *AVBufferRef) *AVBufferRef {
    return (*AVBufferRef)(unsafe.Pointer(C.av_hwframe_ctx_alloc(
        (*C.AVBufferRef)(unsafe.Pointer(device_ctx)))))
}

/**
 * Finalize the context before use. This function must be called after the
 * context is filled with all the required information and before it is attached
 * to any frames.
 *
 * @param ref a reference to the AVHWFramesContext
 * @return 0 on success, a negative AVERROR code on failure
 */
func Av_hwframe_ctx_init(ref *AVBufferRef) int32 {
    return int32(C.av_hwframe_ctx_init((*C.AVBufferRef)(unsafe.Pointer(ref))))
}

/**
 * Allocate a new frame attached to the given AVHWFramesContext.
 *
 * @param hwframe_ctx a reference to an AVHWFramesContext
 * @param frame an empty (freshly allocated or unreffed) frame to be filled with
 *              newly allocated buffers.
 * @param flags currently unused, should be set to zero
 * @return 0 on success, a negative AVERROR code on failure
 */
func Av_hwframe_get_buffer(hwframe_ctx *AVBufferRef, frame *AVFrame, flags int32) int32 {
    return int32(C.av_hwframe_get_buffer(
        (*C.AVBufferRef)(unsafe.Pointer(hwframe_ctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * Copy data to or from a hw surface. At least one of dst/src must have an
 * AVHWFramesContext attached.
 *
 * If src has an AVHWFramesContext attached, then the format of dst (if set)
 * must use one of the formats returned by av_hwframe_transfer_get_formats(src,
 * AV_HWFRAME_TRANSFER_DIRECTION_FROM).
 * If dst has an AVHWFramesContext attached, then the format of src must use one
 * of the formats returned by av_hwframe_transfer_get_formats(dst,
 * AV_HWFRAME_TRANSFER_DIRECTION_TO)
 *
 * dst may be "clean" (i.e. with data/buf pointers unset), in which case the
 * data buffers will be allocated by this function using av_frame_get_buffer().
 * If dst->format is set, then this format will be used, otherwise (when
 * dst->format is AV_PIX_FMT_NONE) the first acceptable format will be chosen.
 *
 * The two frames must have matching allocated dimensions (i.e. equal to
 * AVHWFramesContext.width/height), since not all device types support
 * transferring a sub-rectangle of the whole surface. The display dimensions
 * (i.e. AVFrame.width/height) may be smaller than the allocated dimensions, but
 * also have to be equal for both frames. When the display dimensions are
 * smaller than the allocated dimensions, the content of the padding in the
 * destination frame is unspecified.
 *
 * @param dst the destination frame. dst is not touched on failure.
 * @param src the source frame.
 * @param flags currently unused, should be set to zero
 * @return 0 on success, a negative AVERROR error code on failure.
 */
func Av_hwframe_transfer_data(dst *AVFrame, src *AVFrame, flags int32) int32 {
    return int32(C.av_hwframe_transfer_data((*C.AVFrame)(unsafe.Pointer(dst)), 
        (*C.AVFrame)(unsafe.Pointer(src)), C.int(flags)))
}

type AVHWFrameTransferDirection C.enum_AVHWFrameTransferDirection

/**
 * Get a list of possible source or target formats usable in
 * av_hwframe_transfer_data().
 *
 * @param hwframe_ctx the frame context to obtain the information for
 * @param dir the direction of the transfer
 * @param formats the pointer to the output format list will be written here.
 *                The list is terminated with AV_PIX_FMT_NONE and must be freed
 *                by the caller when no longer needed using av_free().
 *                If this function returns successfully, the format list will
 *                have at least one item (not counting the terminator).
 *                On failure, the contents of this pointer are unspecified.
 * @param flags currently unused, should be set to zero
 * @return 0 on success, a negative AVERROR code on failure.
 */
func Av_hwframe_transfer_get_formats(hwframe_ctx *AVBufferRef,
                                    dir AVHWFrameTransferDirection,
                                    formats **AVPixelFormat, flags int32) int32 {
    return int32(C.av_hwframe_transfer_get_formats(
        (*C.AVBufferRef)(unsafe.Pointer(hwframe_ctx)), 
        C.enum_AVHWFrameTransferDirection(dir), 
        (**C.enum_AVPixelFormat)(unsafe.Pointer(formats)), C.int(flags)))
}


/**
 * This struct describes the constraints on hardware frames attached to
 * a given device with a hardware-specific configuration.  This is returned
 * by av_hwdevice_get_hwframe_constraints() and must be freed by
 * av_hwframe_constraints_free() after use.
 */
type AVHWFramesConstraints C.struct_AVHWFramesConstraints

/**
 * Allocate a HW-specific configuration structure for a given HW device.
 * After use, the user must free all members as required by the specific
 * hardware structure being used, then free the structure itself with
 * av_free().
 *
 * @param device_ctx a reference to the associated AVHWDeviceContext.
 * @return The newly created HW-specific configuration structure on
 *         success or NULL on failure.
 */
func Av_hwdevice_hwconfig_alloc(device_ctx *AVBufferRef) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.av_hwdevice_hwconfig_alloc(
        (*C.AVBufferRef)(unsafe.Pointer(device_ctx)))))
}

/**
 * Get the constraints on HW frames given a device and the HW-specific
 * configuration to be used with that device.  If no HW-specific
 * configuration is provided, returns the maximum possible capabilities
 * of the device.
 *
 * @param ref a reference to the associated AVHWDeviceContext.
 * @param hwconfig a filled HW-specific configuration structure, or NULL
 *        to return the maximum possible capabilities of the device.
 * @return AVHWFramesConstraints structure describing the constraints
 *         on the device, or NULL if not available.
 */
func Av_hwdevice_get_hwframe_constraints(ref *AVBufferRef,
                                                           hwconfig unsafe.Pointer) *AVHWFramesConstraints {
    return (*AVHWFramesConstraints)(unsafe.Pointer(C.av_hwdevice_get_hwframe_constraints(
        (*C.AVBufferRef)(unsafe.Pointer(ref)), hwconfig)))
}

/**
 * Free an AVHWFrameConstraints structure.
 *
 * @param constraints The (filled or unfilled) AVHWFrameConstraints structure.
 */
func Av_hwframe_constraints_free(constraints **AVHWFramesConstraints)  {
    C.av_hwframe_constraints_free(
        (**C.AVHWFramesConstraints)(unsafe.Pointer(constraints)))
}


/**
 * Flags to apply to frame mappings.
 */


/**
 * Map a hardware frame.
 *
 * This has a number of different possible effects, depending on the format
 * and origin of the src and dst frames.  On input, src should be a usable
 * frame with valid buffers and dst should be blank (typically as just created
 * by av_frame_alloc()).  src should have an associated hwframe context, and
 * dst may optionally have a format and associated hwframe context.
 *
 * If src was created by mapping a frame from the hwframe context of dst,
 * then this function undoes the mapping - dst is replaced by a reference to
 * the frame that src was originally mapped from.
 *
 * If both src and dst have an associated hwframe context, then this function
 * attempts to map the src frame from its hardware context to that of dst and
 * then fill dst with appropriate data to be usable there.  This will only be
 * possible if the hwframe contexts and associated devices are compatible -
 * given compatible devices, av_hwframe_ctx_create_derived() can be used to
 * create a hwframe context for dst in which mapping should be possible.
 *
 * If src has a hwframe context but dst does not, then the src frame is
 * mapped to normal memory and should thereafter be usable as a normal frame.
 * If the format is set on dst, then the mapping will attempt to create dst
 * with that format and fail if it is not possible.  If format is unset (is
 * AV_PIX_FMT_NONE) then dst will be mapped with whatever the most appropriate
 * format to use is (probably the sw_format of the src hwframe context).
 *
 * A return value of AVERROR(ENOSYS) indicates that the mapping is not
 * possible with the given arguments and hwframe setup, while other return
 * values indicate that it failed somehow.
 *
 * On failure, the destination frame will be left blank, except for the
 * hw_frames_ctx/format fields thay may have been set by the caller - those will
 * be preserved as they were.
 *
 * @param dst Destination frame, to contain the mapping.
 * @param src Source frame, to be mapped.
 * @param flags Some combination of AV_HWFRAME_MAP_* flags.
 * @return Zero on success, negative AVERROR code on failure.
 */
func Av_hwframe_map(dst *AVFrame, src *AVFrame, flags int32) int32 {
    return int32(C.av_hwframe_map((*C.AVFrame)(unsafe.Pointer(dst)), 
        (*C.AVFrame)(unsafe.Pointer(src)), C.int(flags)))
}


/**
 * Create and initialise an AVHWFramesContext as a mapping of another existing
 * AVHWFramesContext on a different device.
 *
 * av_hwframe_ctx_init() should not be called after this.
 *
 * @param derived_frame_ctx  On success, a reference to the newly created
 *                           AVHWFramesContext.
 * @param format             The AVPixelFormat for the derived context.
 * @param derived_device_ctx A reference to the device to create the new
 *                           AVHWFramesContext on.
 * @param source_frame_ctx   A reference to an existing AVHWFramesContext
 *                           which will be mapped to the derived context.
 * @param flags  Some combination of AV_HWFRAME_MAP_* flags, defining the
 *               mapping parameters to apply to frames which are allocated
 *               in the derived device.
 * @return       Zero on success, negative AVERROR code on failure.
 */
func Av_hwframe_ctx_create_derived(derived_frame_ctx **AVBufferRef,
                                  format AVPixelFormat,
                                  derived_device_ctx *AVBufferRef,
                                  source_frame_ctx *AVBufferRef,
                                  flags int32) int32 {
    return int32(C.av_hwframe_ctx_create_derived(
        (**C.AVBufferRef)(unsafe.Pointer(derived_frame_ctx)), 
        C.enum_AVPixelFormat(format), 
        (*C.AVBufferRef)(unsafe.Pointer(derived_device_ctx)), 
        (*C.AVBufferRef)(unsafe.Pointer(source_frame_ctx)), C.int(flags)))
}

                               
