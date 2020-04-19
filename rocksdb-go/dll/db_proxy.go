package dll

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lrocksdb-proxy

#include "db_proxy.h"

*/
import "C" // 切勿换行再写这个

import (
	"fmt"
	"unsafe"
)
type Foo struct {
	A int32
	B int32
}
func B() {
	values := make([]*C.Foo, 2)
	for i, _ := range values {
		p := (*C.Foo)(C.malloc(C.size_t(C.sizeof_Foo)))
		values[i] = p
		(*p).a = C.int(1)
		(*p).b = C.int(2)
	}
	val := C.pass_array(&values[0])
	for _, v := range values {
		C.free(unsafe.Pointer(v))
	}
	fmt.Println("B finished", val)
}

func B2() {
	values := make([]*C.CustomString, 2)
	for i, _ := range values {
		p := (*C.CustomString)(C.malloc(C.size_t(C.sizeof_CustomString)))
		values[i] = p
		temp := "1111"
		(*p).str = C.CString(temp)
		(*p).len = C.int(len(temp))
	}
	val := C.pass_str(&values[0], C.int(2))
	for _, v := range values {
		C.free(unsafe.Pointer(v))
	}
	fmt.Println("B finished", val)
}
func B33() int {
	values := make([]*C.CustomString, 2)
	for i, _ := range values {
		values[i] = constructString("1111")
	}
	ptrArr := constructPtrArray(unsafe.Pointer(&values[0]), 2)
	val := C.pass_ptr(ptrArr)
	for _, v := range values {
		destructString(v)
	}
	destructPtrArray(ptrArr)
	//fmt.Println("B3 finished", val)
	return len(values) - (int)(val)
}

func B3() int {
	return B5()
}
type PtrArray struct {
	ptr unsafe.Pointer
	len int32
}
func B4() int {
	cs := constructString("1")
	values := make([]*C.CustomString, 1)
	values[0] = cs
	//ptrArr := constructPtrArrayNew((*unsafe.Pointer)((unsafe.Pointer)(&values[0])), 1)
	ptrArr := C.PtrArray{ptr:(unsafe.Pointer)(&values[0]), len:1}
	val := C.pass_cstringptr((*C.PtrStrNewArray)(unsafe.Pointer(&ptrArr)))
	destructString(cs)
	//C.MyFeee(unsafe.Pointer(ptrArr))
	return len(values) + int(val) - 1
}
func B5() int {
	cs := constructString("1")
	values := make([]*C.CustomString, 1)
	values[0] = cs
	val := C.pass_cstringptr2(&values[0], 1)
	destructString(cs)
	return len(values) + int(val) - 1
}

func constructString(str string) *C.CustomString {
	p := (*C.CustomString)(C.MyMalloc(C.int(C.sizeof_CustomString)))
	(*p).str = nil//C.CString(str)
	(*p).len = C.int(len(str))
	return p
}
func destructString(cstr *C.CustomString)  {
	//C.free(unsafe.Pointer((*cstr).str))
	C.MyFeee(unsafe.Pointer(cstr))
}

func constructPtrArrayNew(ptr *unsafe.Pointer, len int) *C.PtrStrNewArray {
	ptrArr := (*C.PtrStrNewArray)(C.MyMalloc(C.int(C.sizeof_PtrStrNewArray)))
	(*ptrArr).ptr = ptr
	(*ptrArr).len = C.int(len)
	return ptrArr
}
func constructPtrArray(ptr unsafe.Pointer, len int) *C.PtrArray {
	ptrArr := (*C.PtrArray)(C.MyMalloc(C.int(C.sizeof_PtrArray)))
	(*ptrArr).ptr = ptr
	(*ptrArr).len = C.int(len)
	return ptrArr
}
func destructPtrArray(cptrArr *C.PtrArray)  {
	C.MyFeee(unsafe.Pointer(cptrArr))
}

func createNewDB(pathStr string, cfStr... string) unsafe.Pointer {
	cfArr := make([]*C.CustomString, len(cfStr))
	for index, _ := range cfArr {
		cfArr[index] = constructString(cfStr[index])
	}
	temp := make([]*C.CustomString, 1)
	temp[0] = &C.CustomString{str:C.CString(cfStr[0]), len:2}
	//unsafe.Pointer(&cfArr[0])
	//(**C.CustomString)(unsafe.Pointer(&temp[0]))
	ptrCf := (**C.CustomString)(unsafe.Pointer(&temp[0]))
	fmt.Println(ptrCf)
	_ = C.PtrStrArray{ptr:ptrCf, len:C.int(len(cfStr))}
	for i := 0; i < len(cfStr); i++ {
		cfArr[i] = nil //&C.CustomString{str:C.CString(cfStr[i]), len:C.int(len(cfStr[i]))}
	}
	//cfList := make([]*C.StrArray, len(cfStr))
	//cfStrBytes := make([][]byte, len(cfStr))
	//for i := 0; i <= len(cfStrBytes); i++ {
	//	cfStrBytes[i] = ([]byte)(cfStr[i])
	//	cfList[i] = (*C.StrArray)(unsafe.Pointer(&cfStrBytes[i]))
	//}
	path := C.CString(pathStr)
	defer C.free(unsafe.Pointer(path))
	handles := C.PtrArray{ptr:nil, len:0}
	names := C.PtrArray{ptr:nil, len:0}
	C.CreateNewDB(path, ptrCf, C.int(len(cfStr)), (*C.PtrArray)(unsafe.Pointer(&handles)), (*C.PtrArray)(unsafe.Pointer(&names)))
	return nil
}
func createDB(pathStr string) unsafe.Pointer {
	path := C.CString(pathStr)
	defer C.free(unsafe.Pointer(path))
	db := C.CreateDB(path)
	return db
}

func dbPut(db unsafe.Pointer, key string, value string) bool  {
	ckey := C.CustomString{str:C.CString(key), len:C.int(len(key))}
	cvalue := C.CustomString{str:C.CString(value), len:C.int(len(value))}
	C.DBPut(db, &ckey, &cvalue)
	return true
}

func dbGet(db unsafe.Pointer, key string) string {
	ckey := C.CustomString{str:C.CString(key), len:C.int(len(key))}
	data := C.DBGet(db, &ckey)
	myData := (C.DataStringPtr)(data)
	gob := C.GoString(myData.Data)
	return gob
}


func Test1() {
	//C.Test()
	//value := C.GetValue()
	//fmt.Println(value)
	//path := C.CString("temppath")
	//defer C.free(unsafe.Pointer(path))
	//db := C.CreateDB(path)
	//key := C.CString("aaa")
	//defer C.free(unsafe.Pointer(key))
	//keyValue := C.CString("111")
	//defer C.free(unsafe.Pointer(keyValue))
	//C.DBPut(db, key, keyValue)
	//data := C.DBGet(db, key)
	//myData := (C.DataStringPtr)(data)
	//gob := C.GoString(myData.Data)
	////C.ReleaseString(data)
	//fmt.Println("len", myData.Len, gob)
}

func Test2()  {
	db := createDB("test2")
	fmt.Println(dbGet(db, "aaa"))
	dbPut(db, "aaa", "1111")
	fmt.Println(dbGet(db, "aaa"))
}

func Test3()  {
	db := createNewDB("test3", "cf3")
	fmt.Println(dbGet(db, "aaa"))
	dbPut(db, "aaa", "1111")
	fmt.Println(dbGet(db, "aaa"))
}