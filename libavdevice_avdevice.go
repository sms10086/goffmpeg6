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

//#cgo pkg-config: libavutil libavdevice libavformat
//#include "libavdevice/version_major.h"
//#include "libavdevice/version.h"
//#include "libavutil/log.h"
//#include "libavutil/opt.h"
//#include "libavutil/dict.h"
//#include "libavformat/avformat.h"
//#include "libavdevice/avdevice.h"
import "C"
import (
    "unsafe"
)



                           
                           

                          
                        
/* When included as part of the ffmpeg build, only include the major version
 * to avoid unnecessary rebuilds. When included externally, keep including
 * the full version information. */
                    
      

/**
 * @file
 * @ingroup lavd
 * Main libavdevice API header
 */

/**
 * @defgroup lavd libavdevice
 * Special devices muxing/demuxing library.
 *
 * Libavdevice is a complementary library to @ref libavf "libavformat". It
 * provides various "special" platform-specific muxers and demuxers, e.g. for
 * grabbing devices, audio capture and playback etc. As a consequence, the
 * (de)muxers in libavdevice are of the AVFMT_NOFILE type (they use their own
 * I/O functions). The filename passed to avformat_open_input() often does not
 * refer to an actually existing file, but has some special device-specific
 * meaning - e.g. for xcbgrab it is the display name.
 *
 * To use libavdevice, simply call avdevice_register_all() to register all
 * compiled muxers and demuxers. They all use standard libavformat API.
 *
 * @{
 */

                          
                          
                           
                                 

/**
 * Return the LIBAVDEVICE_VERSION_INT constant.
 */
func Avdevice_version() uint32 {
    return uint32(C.avdevice_version())
}

/**
 * Return the libavdevice build-time configuration.
 */
func Avdevice_configuration() string {
    return C.GoString(C.avdevice_configuration())
}

/**
 * Return the libavdevice license.
 */
func Avdevice_license() string {
    return C.GoString(C.avdevice_license())
}

/**
 * Initialize libavdevice and register all the input and output devices.
 */
func Avdevice_register_all()  {
    C.avdevice_register_all()
}

/**
 * Audio input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
func Av_input_audio_device_next(d *AVInputFormat) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_input_audio_device_next(
        (*C.struct_AVInputFormat)(unsafe.Pointer(d)))))
}

/**
 * Video input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
func Av_input_video_device_next(d *AVInputFormat) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_input_video_device_next(
        (*C.struct_AVInputFormat)(unsafe.Pointer(d)))))
}

/**
 * Audio output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
func Av_output_audio_device_next(d *AVOutputFormat) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_output_audio_device_next(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(d)))))
}

/**
 * Video output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
func Av_output_video_device_next(d *AVOutputFormat) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_output_video_device_next(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(d)))))
}

type AVDeviceRect struct {
    X int32
    Y int32
    Width int32
    Height int32
}


/**
 * Message types used by avdevice_app_to_dev_control_message().
 */
type AVAppToDevMessageType int32


/**
 * Message types used by avdevice_dev_to_app_control_message().
 */
type AVDevToAppMessageType int32


/**
 * Send control message from application to device.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Exact type depends on message type.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when device doesn't implement handler of the message.
 */
func Avdevice_app_to_dev_control_message(s *AVFormatContext,
                                        typex AVAppToDevMessageType,
                                        data unsafe.Pointer, data_size uint64) int32 {
    return int32(C.avdevice_app_to_dev_control_message(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.enum_AVAppToDevMessageType(typex), data, C.ulonglong(data_size)))
}

/**
 * Send control message from device to application.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Can be NULL.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when application doesn't implement handler of the message.
 */
func Avdevice_dev_to_app_control_message(s *AVFormatContext,
                                        typex AVDevToAppMessageType,
                                        data unsafe.Pointer, data_size uint64) int32 {
    return int32(C.avdevice_dev_to_app_control_message(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.enum_AVDevToAppMessageType(typex), data, C.ulonglong(data_size)))
}

/**
 * Structure describes basic parameters of the device.
 */
type AVDeviceInfo struct {
    Device_name *byte
    Device_description *byte
    Media_types *AVMediaType
    Nb_media_types int32
}


/**
 * List of devices.
 */
type AVDeviceInfoList struct {
    Devices **AVDeviceInfo
    Nb_devices int32
    Default_device int32
}


/**
 * List devices.
 *
 * Returns available device names and their parameters.
 *
 * @note: Some devices may accept system-dependent device names that cannot be
 *        autodetected. The list returned by this function cannot be assumed to
 *        be always completed.
 *
 * @param s                device context.
 * @param[out] device_list list of autodetected devices.
 * @return count of autodetected devices, negative on error.
 */
func Avdevice_list_devices(s *AVFormatContext, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_devices(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}

/**
 * Convenient function to free result of avdevice_list_devices().
 *
 * @param device_list device list to be freed.
 */
func Avdevice_free_list_devices(device_list **AVDeviceInfoList)  {
    C.avdevice_free_list_devices(
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list)))
}

/**
 * List devices.
 *
 * Returns available device names and their parameters.
 * These are convinient wrappers for avdevice_list_devices().
 * Device context is allocated and deallocated internally.
 *
 * @param device           device format. May be NULL if device name is set.
 * @param device_name      device name. May be NULL if device format is set.
 * @param device_options   An AVDictionary filled with device-private options. May be NULL.
 *                         The same options must be passed later to avformat_write_header() for output
 *                         devices or avformat_open_input() for input devices, or at any other place
 *                         that affects device-private options.
 * @param[out] device_list list of autodetected devices
 * @return count of autodetected devices, negative on error.
 * @note device argument takes precedence over device_name when both are set.
 */
func Avdevice_list_input_sources(device *AVInputFormat, device_name *byte,
                                device_options *AVDictionary, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_input_sources(
        (*C.struct_AVInputFormat)(unsafe.Pointer(device)), 
        (*C.char)(unsafe.Pointer(device_name)), 
        (*C.struct_AVDictionary)(unsafe.Pointer(device_options)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}
func Avdevice_list_output_sinks(device *AVOutputFormat, device_name *byte,
                               device_options *AVDictionary, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_output_sinks(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(device)), 
        (*C.char)(unsafe.Pointer(device_name)), 
        (*C.struct_AVDictionary)(unsafe.Pointer(device_options)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}

/**
 * @}
 */

                                
