// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * AVOptions
 * copyright (c) 2005 Michael Niedermayer <michaelni@gmx.at>
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
//#include "libavutil/rational.h"
//#include "libavutil/avutil.h"
//#include "libavutil/channel_layout.h"
//#include "libavutil/dict.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/samplefmt.h"
//#include "libavutil/opt.h"
import "C"
import (
    "unsafe"
)

const AV_OPT_FLAG_ENCODING_PARAM = 1
const AV_OPT_FLAG_DECODING_PARAM = 2
const AV_OPT_FLAG_AUDIO_PARAM = 8
const AV_OPT_FLAG_VIDEO_PARAM = 16
const AV_OPT_FLAG_SUBTITLE_PARAM = 32
const AV_OPT_FLAG_EXPORT = 64
const AV_OPT_FLAG_READONLY = 128
const AV_OPT_FLAG_BSF_PARAM = (1<<8)
const AV_OPT_FLAG_RUNTIME_PARAM = (1<<15)
const AV_OPT_FLAG_FILTERING_PARAM = (1<<16)
const AV_OPT_FLAG_DEPRECATED = (1<<17)
const AV_OPT_FLAG_CHILD_CONSTS = (1<<18)
const AV_OPT_SEARCH_CHILDREN = (1 << 0)
const AV_OPT_SEARCH_FAKE_OBJ = (1 << 1)
const AV_OPT_ALLOW_NULL = (1 << 2)
const AV_OPT_MULTI_COMPONENT_RANGE = (1 << 12)
const AV_OPT_SERIALIZE_SKIP_DEFAULTS = 0x00000001
const AV_OPT_SERIALIZE_OPT_FLAGS_EXACT = 0x00000002


                    
                    

/**
 * @file
 * AVOptions
 */

                     
                   
                           
                 
                
                   
                      

/**
 * @defgroup avoptions AVOptions
 * @ingroup lavu_data
 * @{
 * AVOptions provide a generic system to declare options on arbitrary structs
 * ("objects"). An option can have a help text, a type and a range of possible
 * values. Options may then be enumerated, read and written to.
 *
 * @section avoptions_implement Implementing AVOptions
 * This section describes how to add AVOptions capabilities to a struct.
 *
 * All AVOptions-related information is stored in an AVClass. Therefore
 * the first member of the struct should be a pointer to an AVClass describing it.
 * The option field of the AVClass must be set to a NULL-terminated static array
 * of AVOptions. Each AVOption must have a non-empty name, a type, a default
 * value and for number-type AVOptions also a range of allowed values. It must
 * also declare an offset in bytes from the start of the struct, where the field
 * associated with this AVOption is located. Other fields in the AVOption struct
 * should also be set when applicable, but are not required.
 *
 * The following example illustrates an AVOptions-enabled struct:
 * @code
 * typedef struct test_struct {
 *     const AVClass *class;
 *     int      int_opt;
 *     char    *str_opt;
 *     uint8_t *bin_opt;
 *     int      bin_len;
 * } test_struct;
 *
 * static const AVOption test_options[] = {
 *   { "test_int", "This is a test option of int type.", offsetof(test_struct, int_opt),
 *     AV_OPT_TYPE_INT, { .i64 = -1 }, INT_MIN, INT_MAX },
 *   { "test_str", "This is a test option of string type.", offsetof(test_struct, str_opt),
 *     AV_OPT_TYPE_STRING },
 *   { "test_bin", "This is a test option of binary type.", offsetof(test_struct, bin_opt),
 *     AV_OPT_TYPE_BINARY },
 *   { NULL },
 * };
 *
 * static const AVClass test_class = {
 *     .class_name = "test class",
 *     .item_name  = av_default_item_name,
 *     .option     = test_options,
 *     .version    = LIBAVUTIL_VERSION_INT,
 * };
 * @endcode
 *
 * Next, when allocating your struct, you must ensure that the AVClass pointer
 * is set to the correct value. Then, av_opt_set_defaults() can be called to
 * initialize defaults. After that the struct is ready to be used with the
 * AVOptions API.
 *
 * When cleaning up, you may use the av_opt_free() function to automatically
 * free all the allocated string and binary options.
 *
 * Continuing with the above example:
 *
 * @code
 * test_struct *alloc_test_struct(void)
 * {
 *     test_struct *ret = av_mallocz(sizeof(*ret));
 *     ret->class = &test_class;
 *     av_opt_set_defaults(ret);
 *     return ret;
 * }
 * void free_test_struct(test_struct **foo)
 * {
 *     av_opt_free(*foo);
 *     av_freep(foo);
 * }
 * @endcode
 *
 * @subsection avoptions_implement_nesting Nesting
 *      It may happen that an AVOptions-enabled struct contains another
 *      AVOptions-enabled struct as a member (e.g. AVCodecContext in
 *      libavcodec exports generic options, while its priv_data field exports
 *      codec-specific options). In such a case, it is possible to set up the
 *      parent struct to export a child's options. To do that, simply
 *      implement AVClass.child_next() and AVClass.child_class_iterate() in the
 *      parent struct's AVClass.
 *      Assuming that the test_struct from above now also contains a
 *      child_struct field:
 *
 *      @code
 *      typedef struct child_struct {
 *          AVClass *class;
 *          int flags_opt;
 *      } child_struct;
 *      static const AVOption child_opts[] = {
 *          { "test_flags", "This is a test option of flags type.",
 *            offsetof(child_struct, flags_opt), AV_OPT_TYPE_FLAGS, { .i64 = 0 }, INT_MIN, INT_MAX },
 *          { NULL },
 *      };
 *      static const AVClass child_class = {
 *          .class_name = "child class",
 *          .item_name  = av_default_item_name,
 *          .option     = child_opts,
 *          .version    = LIBAVUTIL_VERSION_INT,
 *      };
 *
 *      void *child_next(void *obj, void *prev)
 *      {
 *          test_struct *t = obj;
 *          if (!prev && t->child_struct)
 *              return t->child_struct;
 *          return NULL
 *      }
 *      const AVClass child_class_iterate(void **iter)
 *      {
 *          const AVClass *c = *iter ? NULL : &child_class;
 *          *iter = (void*)(uintptr_t)c;
 *          return c;
 *      }
 *      @endcode
 *      Putting child_next() and child_class_iterate() as defined above into
 *      test_class will now make child_struct's options accessible through
 *      test_struct (again, proper setup as described above needs to be done on
 *      child_struct right after it is created).
 *
 *      From the above example it might not be clear why both child_next()
 *      and child_class_iterate() are needed. The distinction is that child_next()
 *      iterates over actually existing objects, while child_class_iterate()
 *      iterates over all possible child classes. E.g. if an AVCodecContext
 *      was initialized to use a codec which has private options, then its
 *      child_next() will return AVCodecContext.priv_data and finish
 *      iterating. OTOH child_class_iterate() on AVCodecContext.av_class will
 *      iterate over all available codecs with private options.
 *
 * @subsection avoptions_implement_named_constants Named constants
 *      It is possible to create named constants for options. Simply set the unit
 *      field of the option the constants should apply to a string and
 *      create the constants themselves as options of type AV_OPT_TYPE_CONST
 *      with their unit field set to the same string.
 *      Their default_val field should contain the value of the named
 *      constant.
 *      For example, to add some named constants for the test_flags option
 *      above, put the following into the child_opts array:
 *      @code
 *      { "test_flags", "This is a test option of flags type.",
 *        offsetof(child_struct, flags_opt), AV_OPT_TYPE_FLAGS, { .i64 = 0 }, INT_MIN, INT_MAX, "test_unit" },
 *      { "flag1", "This is a flag with value 16", 0, AV_OPT_TYPE_CONST, { .i64 = 16 }, 0, 0, "test_unit" },
 *      @endcode
 *
 * @section avoptions_use Using AVOptions
 * This section deals with accessing options in an AVOptions-enabled struct.
 * Such structs in FFmpeg are e.g. AVCodecContext in libavcodec or
 * AVFormatContext in libavformat.
 *
 * @subsection avoptions_use_examine Examining AVOptions
 * The basic functions for examining options are av_opt_next(), which iterates
 * over all options defined for one object, and av_opt_find(), which searches
 * for an option with the given name.
 *
 * The situation is more complicated with nesting. An AVOptions-enabled struct
 * may have AVOptions-enabled children. Passing the AV_OPT_SEARCH_CHILDREN flag
 * to av_opt_find() will make the function search children recursively.
 *
 * For enumerating there are basically two cases. The first is when you want to
 * get all options that may potentially exist on the struct and its children
 * (e.g.  when constructing documentation). In that case you should call
 * av_opt_child_class_iterate() recursively on the parent struct's AVClass.  The
 * second case is when you have an already initialized struct with all its
 * children and you want to get all options that can be actually written or read
 * from it. In that case you should call av_opt_child_next() recursively (and
 * av_opt_next() on each result).
 *
 * @subsection avoptions_use_get_set Reading and writing AVOptions
 * When setting options, you often have a string read directly from the
 * user. In such a case, simply passing it to av_opt_set() is enough. For
 * non-string type options, av_opt_set() will parse the string according to the
 * option type.
 *
 * Similarly av_opt_get() will read any option type and convert it to a string
 * which will be returned. Do not forget that the string is allocated, so you
 * have to free it with av_free().
 *
 * In some cases it may be more convenient to put all options into an
 * AVDictionary and call av_opt_set_dict() on it. A specific case of this
 * are the format/codec open functions in lavf/lavc which take a dictionary
 * filled with option as a parameter. This makes it possible to set some options
 * that cannot be set otherwise, since e.g. the input file format is not known
 * before the file is actually opened.
 */

type AVOptionType C.enum_AVOptionType

/**
 * AVOption
 */
type AVOption C.struct_AVOption

/**
 * A single allowed range of values, or a single allowed value.
 */
type AVOptionRange C.struct_AVOptionRange

/**
 * List of AVOptionRange structs.
 */


/**
 * Show the obj options.
 *
 * @param req_flags requested flags for the options to show. Show only the
 * options for which it is opt->flags & req_flags.
 * @param rej_flags rejected flags for the options to show. Show only the
 * options for which it is !(opt->flags & req_flags).
 * @param av_log_obj log context to use for showing the options
 */
func Av_opt_show2(obj unsafe.Pointer, av_log_obj unsafe.Pointer, req_flags int32, rej_flags int32) int32 {
    return int32(C.av_opt_show2(obj, av_log_obj, C.int(req_flags), 
        C.int(rej_flags)))
}

/**
 * Set the values of all AVOption fields to their default values.
 *
 * @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
 */
func Av_opt_set_defaults(s unsafe.Pointer)  {
    C.av_opt_set_defaults(s)
}

/**
 * Set the values of all AVOption fields to their default values. Only these
 * AVOption fields for which (opt->flags & mask) == flags will have their
 * default applied to s.
 *
 * @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
 * @param mask combination of AV_OPT_FLAG_*
 * @param flags combination of AV_OPT_FLAG_*
 */
func Av_opt_set_defaults2(s unsafe.Pointer, mask int32, flags int32)  {
    C.av_opt_set_defaults2(s, C.int(mask), C.int(flags))
}

/**
 * Parse the key/value pairs list in opts. For each key/value pair
 * found, stores the value in the field in ctx that is named like the
 * key. ctx must be an AVClass context, storing is done using
 * AVOptions.
 *
 * @param opts options string to parse, may be NULL
 * @param key_val_sep a 0-terminated list of characters used to
 * separate key from value
 * @param pairs_sep a 0-terminated list of characters used to separate
 * two pairs from each other
 * @return the number of successfully set key/value pairs, or a negative
 * value corresponding to an AVERROR code in case of error:
 * AVERROR(EINVAL) if opts cannot be parsed,
 * the error code issued by av_opt_set() if a key/value pair
 * cannot be set
 */
func Av_set_options_string(ctx unsafe.Pointer, opts *byte,
                          key_val_sep *byte, pairs_sep *byte) int32 {
    return int32(C.av_set_options_string(ctx, (*C.char)(unsafe.Pointer(opts)), 
        (*C.char)(unsafe.Pointer(key_val_sep)), 
        (*C.char)(unsafe.Pointer(pairs_sep))))
}

/**
 * Parse the key-value pairs list in opts. For each key=value pair found,
 * set the value of the corresponding option in ctx.
 *
 * @param ctx          the AVClass object to set options on
 * @param opts         the options string, key-value pairs separated by a
 *                     delimiter
 * @param shorthand    a NULL-terminated array of options names for shorthand
 *                     notation: if the first field in opts has no key part,
 *                     the key is taken from the first element of shorthand;
 *                     then again for the second, etc., until either opts is
 *                     finished, shorthand is finished or a named option is
 *                     found; after that, all options must be named
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value, for example '='
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other, for example ':' or ','
 * @return  the number of successfully set key=value pairs, or a negative
 *          value corresponding to an AVERROR code in case of error:
 *          AVERROR(EINVAL) if opts cannot be parsed,
 *          the error code issued by av_set_string3() if a key/value pair
 *          cannot be set
 *
 * Options names must use only the following characters: a-z A-Z 0-9 - . / _
 * Separators must use characters distinct from option names and from each
 * other.
 */
func Av_opt_set_from_string(ctx unsafe.Pointer, opts *byte,
                           shorthand **byte,
                           key_val_sep *byte, pairs_sep *byte) int32 {
    return int32(C.av_opt_set_from_string(ctx, (*C.char)(unsafe.Pointer(opts)), 
        (**C.char)(unsafe.Pointer(shorthand)), 
        (*C.char)(unsafe.Pointer(key_val_sep)), 
        (*C.char)(unsafe.Pointer(pairs_sep))))
}
/**
 * Free all allocated objects in obj.
 */
func Av_opt_free(obj unsafe.Pointer)  {
    C.av_opt_free(obj)
}

/**
 * Check whether a particular flag is set in a flags field.
 *
 * @param field_name the name of the flag field option
 * @param flag_name the name of the flag to check
 * @return non-zero if the flag is set, zero if the flag isn't set,
 *         isn't of the right type, or the flags field doesn't exist.
 */
func Av_opt_flag_is_set(obj unsafe.Pointer, field_name *byte, flag_name *byte) int32 {
    return int32(C.av_opt_flag_is_set(obj, (*C.char)(unsafe.Pointer(field_name)), 
        (*C.char)(unsafe.Pointer(flag_name))))
}

/**
 * Set all the options from a given dictionary on an object.
 *
 * @param obj a struct whose first element is a pointer to AVClass
 * @param options options to process. This dictionary will be freed and replaced
 *                by a new one containing all options not found in obj.
 *                Of course this new dictionary needs to be freed by caller
 *                with av_dict_free().
 *
 * @return 0 on success, a negative AVERROR if some option was found in obj,
 *         but could not be set.
 *
 * @see av_dict_copy()
 */
func Av_opt_set_dict(obj unsafe.Pointer, options **AVDictionary) int32 {
    return int32(C.av_opt_set_dict(obj, 
        (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}


/**
 * Set all the options from a given dictionary on an object.
 *
 * @param obj a struct whose first element is a pointer to AVClass
 * @param options options to process. This dictionary will be freed and replaced
 *                by a new one containing all options not found in obj.
 *                Of course this new dictionary needs to be freed by caller
 *                with av_dict_free().
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 *
 * @return 0 on success, a negative AVERROR if some option was found in obj,
 *         but could not be set.
 *
 * @see av_dict_copy()
 */
func Av_opt_set_dict2(obj unsafe.Pointer, options **AVDictionary, search_flags int32) int32 {
    return int32(C.av_opt_set_dict2(obj, 
        (**C.struct_AVDictionary)(unsafe.Pointer(options)), C.int(search_flags)))
}

/**
 * Extract a key-value pair from the beginning of a string.
 *
 * @param ropts        pointer to the options string, will be updated to
 *                     point to the rest of the string (one of the pairs_sep
 *                     or the final NUL)
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value, for example '='
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other, for example ':' or ','
 * @param flags        flags; see the AV_OPT_FLAG_* values below
 * @param rkey         parsed key; must be freed using av_free()
 * @param rval         parsed value; must be freed using av_free()
 *
 * @return  >=0 for success, or a negative value corresponding to an
 *          AVERROR code in case of error; in particular:
 *          AVERROR(EINVAL) if no key is present
 *
 */
func Av_opt_get_key_value(ropts **byte,
                         key_val_sep *byte, pairs_sep *byte,
                         flags uint32,
                         rkey **byte, rval **byte) int32 {
    return int32(C.av_opt_get_key_value((**C.char)(unsafe.Pointer(ropts)), 
        (*C.char)(unsafe.Pointer(key_val_sep)), 
        (*C.char)(unsafe.Pointer(pairs_sep)), C.unsigned(flags), 
        (**C.char)(unsafe.Pointer(rkey)), (**C.char)(unsafe.Pointer(rval))))
}



/**
 * @defgroup opt_eval_funcs Evaluating option strings
 * @{
 * This group of functions can be used to evaluate option strings
 * and get numbers out of them. They do the same thing as av_opt_set(),
 * except the result is written into the caller-supplied pointer.
 *
 * @param obj a struct whose first element is a pointer to AVClass.
 * @param o an option for which the string is to be evaluated.
 * @param val string to be evaluated.
 * @param *_out value of the string will be written here.
 *
 * @return 0 on success, a negative number on failure.
 */
func Av_opt_eval_flags (obj unsafe.Pointer, o *AVOption, val *byte, flags_out *int32) int32 {
    return int32(C.av_opt_eval_flags(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.int)(unsafe.Pointer(flags_out))))
}
func Av_opt_eval_int   (obj unsafe.Pointer, o *AVOption, val *byte, int_out *int32) int32 {
    return int32(C.av_opt_eval_int(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.int)(unsafe.Pointer(int_out))))
}
func Av_opt_eval_int64 (obj unsafe.Pointer, o *AVOption, val *byte, int64_out *int64) int32 {
    return int32(C.av_opt_eval_int64(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.longlong)(unsafe.Pointer(int64_out))))
}
func Av_opt_eval_float (obj unsafe.Pointer, o *AVOption, val *byte, float_out *float32) int32 {
    return int32(C.av_opt_eval_float(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.float)(unsafe.Pointer(float_out))))
}
func Av_opt_eval_double(obj unsafe.Pointer, o *AVOption, val *byte, double_out *float64) int32 {
    return int32(C.av_opt_eval_double(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.double)(unsafe.Pointer(double_out))))
}
func Av_opt_eval_q     (obj unsafe.Pointer, o *AVOption, val *byte, q_out *AVRational) int32 {
    return int32(C.av_opt_eval_q(obj, (*C.AVOption)(unsafe.Pointer(o)), 
        (*C.char)(unsafe.Pointer(val)), (*C.AVRational)(unsafe.Pointer(q_out))))
}
/**
 * @}
 */

/**< Search in possible children of the
                                               given object first. */
/**
 *  The obj passed to av_opt_find() is fake -- only a double pointer to AVClass
 *  instead of a required pointer to a struct containing AVClass. This is
 *  useful for searching for options without needing to allocate the corresponding
 *  object.
 */


/**
 *  In av_opt_get, return NULL if the option has a pointer type and is set to NULL,
 *  rather than returning an empty string.
 */


/**
 *  Allows av_opt_query_ranges and av_opt_query_ranges_default to return more than
 *  one component for certain option types.
 *  @see AVOptionRanges for details.
 */


/**
 * Look for an option in an object. Consider only options which
 * have all the specified flags set.
 *
 * @param[in] obj A pointer to a struct whose first element is a
 *                pointer to an AVClass.
 *                Alternatively a double pointer to an AVClass, if
 *                AV_OPT_SEARCH_FAKE_OBJ search flag is set.
 * @param[in] name The name of the option to look for.
 * @param[in] unit When searching for named constants, name of the unit
 *                 it belongs to.
 * @param opt_flags Find only options with all the specified flags set (AV_OPT_FLAG).
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 *
 * @return A pointer to the option found, or NULL if no option
 *         was found.
 *
 * @note Options found with AV_OPT_SEARCH_CHILDREN flag may not be settable
 * directly with av_opt_set(). Use special calls which take an options
 * AVDictionary (e.g. avformat_open_input()) to set options found with this
 * flag.
 */
func Av_opt_find(obj unsafe.Pointer, name *byte, unit *byte,
                            opt_flags int32, search_flags int32) *AVOption {
    return (*AVOption)(unsafe.Pointer(C.av_opt_find(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.char)(unsafe.Pointer(unit)), C.int(opt_flags), C.int(search_flags))))
}

/**
 * Look for an option in an object. Consider only options which
 * have all the specified flags set.
 *
 * @param[in] obj A pointer to a struct whose first element is a
 *                pointer to an AVClass.
 *                Alternatively a double pointer to an AVClass, if
 *                AV_OPT_SEARCH_FAKE_OBJ search flag is set.
 * @param[in] name The name of the option to look for.
 * @param[in] unit When searching for named constants, name of the unit
 *                 it belongs to.
 * @param opt_flags Find only options with all the specified flags set (AV_OPT_FLAG).
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 * @param[out] target_obj if non-NULL, an object to which the option belongs will be
 * written here. It may be different from obj if AV_OPT_SEARCH_CHILDREN is present
 * in search_flags. This parameter is ignored if search_flags contain
 * AV_OPT_SEARCH_FAKE_OBJ.
 *
 * @return A pointer to the option found, or NULL if no option
 *         was found.
 */
func Av_opt_find2(obj unsafe.Pointer, name *byte, unit *byte,
                             opt_flags int32, search_flags int32, target_obj *unsafe.Pointer) *AVOption {
    return (*AVOption)(unsafe.Pointer(C.av_opt_find2(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.char)(unsafe.Pointer(unit)), C.int(opt_flags), C.int(search_flags), 
        target_obj)))
}

/**
 * Iterate over all AVOptions belonging to obj.
 *
 * @param obj an AVOptions-enabled struct or a double pointer to an
 *            AVClass describing it.
 * @param prev result of the previous call to av_opt_next() on this object
 *             or NULL
 * @return next AVOption or NULL
 */
func Av_opt_next(obj unsafe.Pointer, prev *AVOption) *AVOption {
    return (*AVOption)(unsafe.Pointer(C.av_opt_next(obj, (*C.AVOption)(unsafe.Pointer(prev)))))
}

/**
 * Iterate over AVOptions-enabled children of obj.
 *
 * @param prev result of a previous call to this function or NULL
 * @return next AVOptions-enabled child or NULL
 */
func Av_opt_child_next(obj unsafe.Pointer, prev unsafe.Pointer) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.av_opt_child_next(obj, prev)))
}

/**
 * Iterate over potential AVOptions-enabled children of parent.
 *
 * @param iter a pointer where iteration state is stored.
 * @return AVClass corresponding to next potential child or NULL
 */
func Av_opt_child_class_iterate(parent *AVClass, iter *unsafe.Pointer) *AVClass {
    return (*AVClass)(unsafe.Pointer(C.av_opt_child_class_iterate(
        (*C.AVClass)(unsafe.Pointer(parent)), iter)))
}

/**
 * @defgroup opt_set_funcs Option setting functions
 * @{
 * Those functions set the field of obj with the given name to value.
 *
 * @param[in] obj A struct whose first element is a pointer to an AVClass.
 * @param[in] name the name of the field to set
 * @param[in] val The value to set. In case of av_opt_set() if the field is not
 * of a string type, then the given string is parsed.
 * SI postfixes and some named scalars are supported.
 * If the field is of a numeric type, it has to be a numeric or named
 * scalar. Behavior with more than one scalar and +- infix operators
 * is undefined.
 * If the field is of a flags type, it has to be a sequence of numeric
 * scalars or named flags separated by '+' or '-'. Prefixing a flag
 * with '+' causes it to be set without affecting the other flags;
 * similarly, '-' unsets a flag.
 * If the field is of a dictionary type, it has to be a ':' separated list of
 * key=value parameters. Values containing ':' special characters must be
 * escaped.
 * @param search_flags flags passed to av_opt_find2. I.e. if AV_OPT_SEARCH_CHILDREN
 * is passed here, then the option may be set on a child of obj.
 *
 * @return 0 if the value has been set, or an AVERROR code in case of
 * error:
 * AVERROR_OPTION_NOT_FOUND if no matching option exists
 * AVERROR(ERANGE) if the value is out of range
 * AVERROR(EINVAL) if the value is not valid
 */
func Av_opt_set         (obj unsafe.Pointer, name *byte, val *byte, search_flags int32) int32 {
    return int32(C.av_opt_set(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.char)(unsafe.Pointer(val)), C.int(search_flags)))
}
func Av_opt_set_int     (obj unsafe.Pointer, name *byte, val int64, search_flags int32) int32 {
    return int32(C.av_opt_set_int(obj, (*C.char)(unsafe.Pointer(name)), 
        C.longlong(val), C.int(search_flags)))
}
func Av_opt_set_double  (obj unsafe.Pointer, name *byte, val float64, search_flags int32) int32 {
    return int32(C.av_opt_set_double(obj, (*C.char)(unsafe.Pointer(name)), 
        C.double(val), C.int(search_flags)))
}
func Av_opt_set_q       (obj unsafe.Pointer, name *byte, val AVRational, search_flags int32) int32 {
    return int32(C.av_opt_set_q(obj, (*C.char)(unsafe.Pointer(name)), 
        C.AVRational(val), C.int(search_flags)))
}
func Av_opt_set_bin     (obj unsafe.Pointer, name *byte, val *uint8, size int32, search_flags int32) int32 {
    return int32(C.av_opt_set_bin(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.uchar)(unsafe.Pointer(val)), C.int(size), C.int(search_flags)))
}
func Av_opt_set_image_size(obj unsafe.Pointer, name *byte, w int32, h int32, search_flags int32) int32 {
    return int32(C.av_opt_set_image_size(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(w), C.int(h), C.int(search_flags)))
}
func Av_opt_set_pixel_fmt (obj unsafe.Pointer, name *byte, fmt AVPixelFormat, search_flags int32) int32 {
    return int32(C.av_opt_set_pixel_fmt(obj, (*C.char)(unsafe.Pointer(name)), 
        C.enum_AVPixelFormat(fmt), C.int(search_flags)))
}
func Av_opt_set_sample_fmt(obj unsafe.Pointer, name *byte, fmt AVSampleFormat, search_flags int32) int32 {
    return int32(C.av_opt_set_sample_fmt(obj, (*C.char)(unsafe.Pointer(name)), 
        C.enum_AVSampleFormat(fmt), C.int(search_flags)))
}
func Av_opt_set_video_rate(obj unsafe.Pointer, name *byte, val AVRational, search_flags int32) int32 {
    return int32(C.av_opt_set_video_rate(obj, (*C.char)(unsafe.Pointer(name)), 
        C.AVRational(val), C.int(search_flags)))
}
                             
                    
                                                                                                
      
func Av_opt_set_chlayout(obj unsafe.Pointer, name *byte, layout *AVChannelLayout, search_flags int32) int32 {
    return int32(C.av_opt_set_chlayout(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.AVChannelLayout)(unsafe.Pointer(layout)), C.int(search_flags)))
}
/**
 * @note Any old dictionary present is discarded and replaced with a copy of the new one. The
 * caller still owns val is and responsible for freeing it.
 */
func Av_opt_set_dict_val(obj unsafe.Pointer, name *byte, val *AVDictionary, search_flags int32) int32 {
    return int32(C.av_opt_set_dict_val(obj, (*C.char)(unsafe.Pointer(name)), 
        (*C.AVDictionary)(unsafe.Pointer(val)), C.int(search_flags)))
}

/**
 * Set a binary option to an integer list.
 *
 * @param obj    AVClass object to set options on
 * @param name   name of the binary option
 * @param val    pointer to an integer list (must have the correct type with
 *               regard to the contents of the list)
 * @param term   list terminator (usually 0 or -1)
 * @param flags  search flags
 */


/**
 * @}
 */

/**
 * @defgroup opt_get_funcs Option getting functions
 * @{
 * Those functions get a value of the option with the given name from an object.
 *
 * @param[in] obj a struct whose first element is a pointer to an AVClass.
 * @param[in] name name of the option to get.
 * @param[in] search_flags flags passed to av_opt_find2. I.e. if AV_OPT_SEARCH_CHILDREN
 * is passed here, then the option may be found in a child of obj.
 * @param[out] out_val value of the option will be written here
 * @return >=0 on success, a negative error code otherwise
 */
/**
 * @note the returned string will be av_malloc()ed and must be av_free()ed by the caller
 *
 * @note if AV_OPT_ALLOW_NULL is set in search_flags in av_opt_get, and the
 * option is of type AV_OPT_TYPE_STRING, AV_OPT_TYPE_BINARY or AV_OPT_TYPE_DICT
 * and is set to NULL, *out_val will be set to NULL instead of an allocated
 * empty string.
 */
func Av_opt_get         (obj unsafe.Pointer, name *byte, search_flags int32, out_val **uint8) int32 {
    return int32(C.av_opt_get(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (**C.uchar)(unsafe.Pointer(out_val))))
}
func Av_opt_get_int     (obj unsafe.Pointer, name *byte, search_flags int32, out_val *int64) int32 {
    return int32(C.av_opt_get_int(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.longlong)(unsafe.Pointer(out_val))))
}
func Av_opt_get_double  (obj unsafe.Pointer, name *byte, search_flags int32, out_val *float64) int32 {
    return int32(C.av_opt_get_double(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.double)(unsafe.Pointer(out_val))))
}
func Av_opt_get_q       (obj unsafe.Pointer, name *byte, search_flags int32, out_val *AVRational) int32 {
    return int32(C.av_opt_get_q(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.AVRational)(unsafe.Pointer(out_val))))
}
func Av_opt_get_image_size(obj unsafe.Pointer, name *byte, search_flags int32, w_out *int32, h_out *int32) int32 {
    return int32(C.av_opt_get_image_size(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.int)(unsafe.Pointer(w_out)), 
        (*C.int)(unsafe.Pointer(h_out))))
}
func Av_opt_get_pixel_fmt (obj unsafe.Pointer, name *byte, search_flags int32, out_fmt *AVPixelFormat) int32 {
    return int32(C.av_opt_get_pixel_fmt(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.enum_AVPixelFormat)(unsafe.Pointer(out_fmt))))
}
func Av_opt_get_sample_fmt(obj unsafe.Pointer, name *byte, search_flags int32, out_fmt *AVSampleFormat) int32 {
    return int32(C.av_opt_get_sample_fmt(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.enum_AVSampleFormat)(unsafe.Pointer(out_fmt))))
}
func Av_opt_get_video_rate(obj unsafe.Pointer, name *byte, search_flags int32, out_val *AVRational) int32 {
    return int32(C.av_opt_get_video_rate(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.AVRational)(unsafe.Pointer(out_val))))
}
                             
                    
                                                                                                 
      
func Av_opt_get_chlayout(obj unsafe.Pointer, name *byte, search_flags int32, layout *AVChannelLayout) int32 {
    return int32(C.av_opt_get_chlayout(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (*C.AVChannelLayout)(unsafe.Pointer(layout))))
}
/**
 * @param[out] out_val The returned dictionary is a copy of the actual value and must
 * be freed with av_dict_free() by the caller
 */
func Av_opt_get_dict_val(obj unsafe.Pointer, name *byte, search_flags int32, out_val **AVDictionary) int32 {
    return int32(C.av_opt_get_dict_val(obj, (*C.char)(unsafe.Pointer(name)), 
        C.int(search_flags), (**C.AVDictionary)(unsafe.Pointer(out_val))))
}
/**
 * @}
 */
/**
 * Gets a pointer to the requested field in a struct.
 * This function allows accessing a struct even when its fields are moved or
 * renamed since the application making the access has been compiled,
 *
 * @returns a pointer to the field, it can be cast to the correct type and read
 *          or written to.
 */
func Av_opt_ptr(avclass *AVClass, obj unsafe.Pointer, name *byte) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.av_opt_ptr((*C.AVClass)(unsafe.Pointer(avclass)), 
        obj, (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Free an AVOptionRanges struct and set it to NULL.
 */
func Av_opt_freep_ranges(ranges **AVOptionRanges)  {
    C.av_opt_freep_ranges((**C.AVOptionRanges)(unsafe.Pointer(ranges)))
}

/**
 * Get a list of allowed ranges for the given option.
 *
 * The returned list may depend on other fields in obj like for example profile.
 *
 * @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
 *              AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
 *              AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges
 *
 * The result must be freed with av_opt_freep_ranges.
 *
 * @return number of compontents returned on success, a negative errro code otherwise
 */
func Av_opt_query_ranges(_var0 **AVOptionRanges, obj unsafe.Pointer, key *byte, flags int32) int32 {
    return int32(C.av_opt_query_ranges(
        (**C.AVOptionRanges)(unsafe.Pointer(_var0)), obj, 
        (*C.char)(unsafe.Pointer(key)), C.int(flags)))
}

/**
 * Copy options from src object into dest object.
 *
 * The underlying AVClass of both src and dest must coincide. The guarantee
 * below does not apply if this is not fulfilled.
 *
 * Options that require memory allocation (e.g. string or binary) are malloc'ed in dest object.
 * Original memory allocated for such options is freed unless both src and dest options points to the same memory.
 *
 * Even on error it is guaranteed that allocated options from src and dest
 * no longer alias each other afterwards; in particular calling av_opt_free()
 * on both src and dest is safe afterwards if dest has been memdup'ed from src.
 *
 * @param dest Object to copy from
 * @param src  Object to copy into
 * @return 0 on success, negative on error
 */
func Av_opt_copy(dest unsafe.Pointer, src unsafe.Pointer) int32 {
    return int32(C.av_opt_copy(dest, src))
}

/**
 * Get a default list of allowed ranges for the given option.
 *
 * This list is constructed without using the AVClass.query_ranges() callback
 * and can be used as fallback from within the callback.
 *
 * @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
 *              AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
 *              AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges
 *
 * The result must be freed with av_opt_free_ranges.
 *
 * @return number of compontents returned on success, a negative errro code otherwise
 */
func Av_opt_query_ranges_default(_var0 **AVOptionRanges, obj unsafe.Pointer, key *byte, flags int32) int32 {
    return int32(C.av_opt_query_ranges_default(
        (**C.AVOptionRanges)(unsafe.Pointer(_var0)), obj, 
        (*C.char)(unsafe.Pointer(key)), C.int(flags)))
}

/**
 * Check if given option is set to its default value.
 *
 * Options o must belong to the obj. This function must not be called to check child's options state.
 * @see av_opt_is_set_to_default_by_name().
 *
 * @param obj  AVClass object to check option on
 * @param o    option to be checked
 * @return     >0 when option is set to its default,
 *              0 when option is not set its default,
 *             <0 on error
 */
func Av_opt_is_set_to_default(obj unsafe.Pointer, o *AVOption) int32 {
    return int32(C.av_opt_is_set_to_default(obj, 
        (*C.AVOption)(unsafe.Pointer(o))))
}

/**
 * Check if given option is set to its default value.
 *
 * @param obj          AVClass object to check option on
 * @param name         option name
 * @param search_flags combination of AV_OPT_SEARCH_*
 * @return             >0 when option is set to its default,
 *                     0 when option is not set its default,
 *                     <0 on error
 */
func Av_opt_is_set_to_default_by_name(obj unsafe.Pointer, name *byte, search_flags int32) int32 {
    return int32(C.av_opt_is_set_to_default_by_name(obj, 
        (*C.char)(unsafe.Pointer(name)), C.int(search_flags)))
}


///< Serialize options that are not set to default values only.
///< Serialize options that exactly match opt_flags only.

/**
 * Serialize object's options.
 *
 * Create a string containing object's serialized options.
 * Such string may be passed back to av_opt_set_from_string() in order to restore option values.
 * A key/value or pairs separator occurring in the serialized value or
 * name string are escaped through the av_escape() function.
 *
 * @param[in]  obj           AVClass object to serialize
 * @param[in]  opt_flags     serialize options with all the specified flags set (AV_OPT_FLAG)
 * @param[in]  flags         combination of AV_OPT_SERIALIZE_* flags
 * @param[out] buffer        Pointer to buffer that will be allocated with string containg serialized options.
 *                           Buffer must be freed by the caller when is no longer needed.
 * @param[in]  key_val_sep   character used to separate key from value
 * @param[in]  pairs_sep     character used to separate two pairs from each other
 * @return                   >= 0 on success, negative on error
 * @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 */
func Av_opt_serialize(obj unsafe.Pointer, opt_flags int32, flags int32, buffer **byte,
                     key_val_sep byte, pairs_sep byte) int32 {
    return int32(C.av_opt_serialize(obj, C.int(opt_flags), C.int(flags), 
        (**C.char)(unsafe.Pointer(buffer)), C.char(key_val_sep), 
        C.char(pairs_sep)))
}
/**
 * @}
 */

                         
