// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2000, 2001, 2002 Fabrice Bellard
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
//#include <stddef.h>
//#include "libavutil/cpu.h"
import "C"
import (
    "unsafe"
)

const AV_CPU_FLAG_FORCE = 0x80000000
const AV_CPU_FLAG_MMX = 0x0001
const AV_CPU_FLAG_MMXEXT = 0x0002
const AV_CPU_FLAG_MMX2 = 0x0002
const AV_CPU_FLAG_3DNOW = 0x0004
const AV_CPU_FLAG_SSE = 0x0008
const AV_CPU_FLAG_SSE2 = 0x0010
const AV_CPU_FLAG_SSE2SLOW = 0x40000000
const AV_CPU_FLAG_3DNOWEXT = 0x0020
const AV_CPU_FLAG_SSE3 = 0x0040
const AV_CPU_FLAG_SSE3SLOW = 0x20000000
const AV_CPU_FLAG_SSSE3 = 0x0080
const AV_CPU_FLAG_SSSE3SLOW = 0x4000000
const AV_CPU_FLAG_ATOM = 0x10000000
const AV_CPU_FLAG_SSE4 = 0x0100
const AV_CPU_FLAG_SSE42 = 0x0200
const AV_CPU_FLAG_AESNI = 0x80000
const AV_CPU_FLAG_AVX = 0x4000
const AV_CPU_FLAG_AVXSLOW = 0x8000000
const AV_CPU_FLAG_XOP = 0x0400
const AV_CPU_FLAG_FMA4 = 0x0800
const AV_CPU_FLAG_CMOV = 0x1000
const AV_CPU_FLAG_AVX2 = 0x8000
const AV_CPU_FLAG_FMA3 = 0x10000
const AV_CPU_FLAG_BMI1 = 0x20000
const AV_CPU_FLAG_BMI2 = 0x40000
const AV_CPU_FLAG_AVX512 = 0x100000
const AV_CPU_FLAG_AVX512ICL = 0x200000
const AV_CPU_FLAG_SLOW_GATHER = 0x2000000
const AV_CPU_FLAG_ALTIVEC = 0x0001
const AV_CPU_FLAG_VSX = 0x0002
const AV_CPU_FLAG_POWER8 = 0x0004
const AV_CPU_FLAG_ARMV5TE = (1 << 0)
const AV_CPU_FLAG_ARMV6 = (1 << 1)
const AV_CPU_FLAG_ARMV6T2 = (1 << 2)
const AV_CPU_FLAG_VFP = (1 << 3)
const AV_CPU_FLAG_VFPV3 = (1 << 4)
const AV_CPU_FLAG_NEON = (1 << 5)
const AV_CPU_FLAG_ARMV8 = (1 << 6)
const AV_CPU_FLAG_VFP_VM = (1 << 7)
const AV_CPU_FLAG_DOTPROD = (1 << 8)
const AV_CPU_FLAG_I8MM = (1 << 9)
const AV_CPU_FLAG_SETEND = (1 <<16)
const AV_CPU_FLAG_MMI = (1 << 0)
const AV_CPU_FLAG_MSA = (1 << 1)
const AV_CPU_FLAG_LSX = (1 << 0)
const AV_CPU_FLAG_LASX = (1 << 1)
const AV_CPU_FLAG_RVI = (1 << 0)
const AV_CPU_FLAG_RVF = (1 << 1)
const AV_CPU_FLAG_RVD = (1 << 2)
const AV_CPU_FLAG_RVV_I32 = (1 << 3)
const AV_CPU_FLAG_RVV_F32 = (1 << 4)
const AV_CPU_FLAG_RVV_I64 = (1 << 5)
const AV_CPU_FLAG_RVV_F64 = (1 << 6)
const AV_CPU_FLAG_RVB_BASIC = (1 << 7)
const AV_CPU_FLAG_RVB_ADDR = (1 << 8)


                    
                    

                   

/* force usage of selected flags (OR) */

    /* lower 16 bits - CPU features */
///< standard MMX
///< SSE integer functions or AMD MMX ext
///< SSE integer functions or AMD MMX ext
///< AMD 3DNOW
///< SSE functions
///< PIV SSE2 functions
///< SSE2 supported, but usually not faster
                                        ///< than regular MMX/SSE (e.g. Core1)
///< AMD 3DNowExt
///< Prescott SSE3 functions
///< SSE3 supported, but usually not faster
                                        ///< than regular MMX/SSE (e.g. Core1)
///< Conroe SSSE3 functions
///< SSSE3 supported, but usually not faster
///< Atom processor, some SSSE3 instructions are slower
///< Penryn SSE4.1 functions
///< Nehalem SSE4.2 functions
///< Advanced Encryption Standard functions
///< AVX functions: requires OS support even if YMM registers aren't used
///< AVX supported, but slow when using YMM registers (e.g. Bulldozer)
///< Bulldozer XOP functions
///< Bulldozer FMA4 functions
///< supports cmov instruction
///< AVX2 functions: requires OS support even if YMM registers aren't used
///< Haswell FMA3 functions
///< Bit Manipulation Instruction Set 1
///< Bit Manipulation Instruction Set 2
///< AVX-512 functions: requires OS support even if YMM/ZMM registers aren't used
///< F/CD/BW/DQ/VL/VNNI/IFMA/VBMI/VBMI2/VPOPCNTDQ/BITALG/GFNI/VAES/VPCLMULQDQ
///< CPU has slow gathers.

///< standard
///< ISA 2.06
///< ISA 2.07








///< VFPv2 vector mode, deprecated in ARMv7-A and unavailable in various CPUs implementations







//Loongarch SIMD extension.



// RISC-V extensions
///< I (full GPR bank)
///< F (single precision FP)
///< D (double precision FP)
///< Vectors of 8/16/32-bit int's */
///< Vectors of float's */
///< Vectors of 64-bit int's */
///< Vectors of double's
///< Basic bit-manipulations
///< Address bit-manipulations

/**
 * Return the flags which specify extensions supported by the CPU.
 * The returned value is affected by av_force_cpu_flags() if that was used
 * before. So av_get_cpu_flags() can easily be used in an application to
 * detect the enabled cpu flags.
 */
func Av_get_cpu_flags() int32 {
    return int32(C.av_get_cpu_flags())
}

/**
 * Disables cpu detection and forces the specified flags.
 * -1 is a special case that disables forcing of specific flags.
 */
func Av_force_cpu_flags(flags int32)  {
    C.av_force_cpu_flags(C.int(flags))
}

/**
 * Parse CPU caps from a string and update the given AV_CPU_* flags based on that.
 *
 * @return negative on error.
 */
func Av_parse_cpu_caps(flags *uint32, s *byte) int32 {
    return int32(C.av_parse_cpu_caps((*C.unsigned)(unsafe.Pointer(flags)), 
        (*C.char)(unsafe.Pointer(s))))
}

/**
 * @return the number of logical CPU cores present.
 */
func Av_cpu_count() int32 {
    return int32(C.av_cpu_count())
}

/**
 * Overrides cpu count detection and forces the specified count.
 * Count < 1 disables forcing of specific count.
 */
func Av_cpu_force_count(count int32)  {
    C.av_cpu_force_count(C.int(count))
}

/**
 * Get the maximum data alignment that may be required by FFmpeg.
 *
 * Note that this is affected by the build configuration and the CPU flags mask,
 * so e.g. if the CPU supports AVX, but libavutil has been built with
 * --disable-avx or the AV_CPU_FLAG_AVX flag has been disabled through
 *  av_set_cpu_flags_mask(), then this function will behave as if AVX is not
 *  present.
 */
func Av_cpu_max_align() uint64 {
    return uint64(C.av_cpu_max_align())
}

                         
