/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: verushash.i

package VH

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef _gostring_ swig_type_2;

#cgo LDFLAGS: -L${SRCDIR}/build -l:libverushash.a -lsodium

typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
extern void _wrap_Swig_free_VH_4119d1d66918a908(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_VH_4119d1d66918a908(swig_intgo arg1);
extern swig_type_1 _wrap_cdata_VH_4119d1d66918a908(intgo _swig_args, uintptr_t arg1, swig_intgo arg2);
extern void _wrap_memmove_VH_4119d1d66918a908(uintptr_t arg1, swig_type_2 arg2, swig_intgo arg3);
extern void _wrap_Verushash_initialized_set_VH_4119d1d66918a908(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_Verushash_initialized_get_VH_4119d1d66918a908(uintptr_t arg1);
extern void _wrap_Verushash_initialize_VH_4119d1d66918a908(uintptr_t arg1);
extern void _wrap_Verushash_verushash_VH_4119d1d66918a908(uintptr_t arg1, swig_type_3 arg2, swig_intgo arg3, uintptr_t arg4);
extern void _wrap_Verushash_verushash_v2_VH_4119d1d66918a908(uintptr_t arg1, swig_type_4 arg2, swig_intgo arg3, uintptr_t arg4);
extern void _wrap_Verushash_verushash_v2b_VH_4119d1d66918a908(uintptr_t arg1, swig_type_5 arg2, swig_intgo arg3, uintptr_t arg4);
extern void _wrap_Verushash_verushash_v2b1_VH_4119d1d66918a908(uintptr_t arg1, swig_type_6 arg2, swig_intgo arg3, uintptr_t arg4);
extern void _wrap_Verushash_verushash_v2b2_VH_4119d1d66918a908(uintptr_t arg1, swig_type_7 arg2, uintptr_t arg3);
extern uintptr_t _wrap_new_Verushash_VH_4119d1d66918a908(void);
extern void _wrap_delete_Verushash_VH_4119d1d66918a908(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_VH_4119d1d66918a908(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Cdata(arg1 uintptr, _swig_args ...interface{}) (_swig_ret []byte) {
	var arg2 int
	if len(_swig_args) > 0 {
		arg2 = _swig_args[0].(int)
	}
	var swig_r uint64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (uint64)(C._wrap_cdata_VH_4119d1d66918a908(C.swig_intgo(len(_swig_args)), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	var swig_r_1 []byte

  {
    type swigcdata struct { size int; data uintptr }
    p := (*swigcdata)(unsafe.Pointer(uintptr(swig_r)))
    if p == nil || p.data == 0 {
      swig_r_1 = nil
    } else {
      b := make([]byte, p.size)
      a := (*[0x7fffffff]byte)(unsafe.Pointer(p.data))[:p.size]
      copy(b, a)
      Swig_free(p.data)
      Swig_free(uintptr(unsafe.Pointer(p)))
      swig_r_1 = b
    }
  }

	return swig_r_1
}

func Memmove(arg1 uintptr, arg2 string, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_memmove_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

type SwigcptrVerushash uintptr

func (p SwigcptrVerushash) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVerushash) SwigIsVerushash() {
}

func (arg1 SwigcptrVerushash) SetInitialized(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Verushash_initialized_set_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrVerushash) GetInitialized() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Verushash_initialized_get_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVerushash) Initialize() {
	_swig_i_0 := arg1
	C._wrap_Verushash_initialize_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVerushash) Verushash(arg2 string, arg3 int, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_Verushash_verushash_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrVerushash) Verushash_v2(arg2 string, arg3 int, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_Verushash_verushash_v2_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrVerushash) Verushash_v2b(arg2 string, arg3 int, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_Verushash_verushash_v2b_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrVerushash) Verushash_v2b1(arg2 string, arg3 int, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_Verushash_verushash_v2b1_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrVerushash) Verushash_v2b2(arg2 string, arg3 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Verushash_verushash_v2b2_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func NewVerushash() (_swig_ret Verushash) {
	var swig_r Verushash
	swig_r = (Verushash)(SwigcptrVerushash(C._wrap_new_Verushash_VH_4119d1d66918a908()))
	return swig_r
}

func DeleteVerushash(arg1 Verushash) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Verushash_VH_4119d1d66918a908(C.uintptr_t(_swig_i_0))
}

type Verushash interface {
	Swigcptr() uintptr
	SwigIsVerushash()
	SetInitialized(arg2 bool)
	GetInitialized() (_swig_ret bool)
	Initialize()
	Verushash(arg2 string, arg3 int, arg4 uintptr)
	Verushash_v2(arg2 string, arg3 int, arg4 uintptr)
	Verushash_v2b(arg2 string, arg3 int, arg4 uintptr)
	Verushash_v2b1(arg2 string, arg3 int, arg4 uintptr)
	Verushash_v2b2(arg2 string, arg3 uintptr)
}


