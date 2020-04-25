package dll

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lrocksdb-proxy

#include "db_proxy.h"

*/
import "C" // 切勿换行再写这个

import (
	"fmt"
	"strconv"
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
	//arr := &PtrArray{ptr:unsafe.Pointer(&values[0]), len:2}
	ptrArr := constructPtrArray(unsafe.Pointer(&values[0]), 2)
	//val := C.pass_ptr((*C.PtrArray)(unsafe.Pointer((arr))))
	val := C.pass_ptr(ptrArr, 0)
	for _, v := range values {
		destructString(v)
	}
	destructPtrArray(ptrArr)
	//fmt.Println("B3 finished", val)
	return len(values) - (int)(val)
}

func B3() int {
	return Test5()
}
type PtrArray struct {
	ptr unsafe.Pointer
	len int32
}
type CustomString struct{
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
func B6() int {
	str := C.CString("11")
	values := &PtrArray{ptr:unsafe.Pointer(str), len:1}
	val := C.pass_goptrarr((*C.PtrArray)(unsafe.Pointer(values)))
	C.free(unsafe.Pointer(str))
	return int(val)
}
func B65() int {
	ptr := ToPtrArray([]byte("1111"))
	ptr2 := ToPtrArray([]byte("2222"))
	ptrArray := []*PtrArray{ptr, ptr2}
	values := &PtrArray{ptr:unsafe.Pointer(&ptrArray), len:2}
	val := C.pass_goptrarrarr((*C.PtrArray)(unsafe.Pointer(values)))
	return int(val)
}

func ToPtrArray(data []byte) *PtrArray {
	return (*PtrArray)(unsafe.Pointer(constructPtrArray(unsafe.Pointer(C.CString(string(data))), len(data))))
}
func B7() int {
	ptr := (*PtrArray)(unsafe.Pointer(C.pass_gogetptr()))
	val := ptr.len
	C.MyFeee(ptr.ptr)
	C.MyFeee(unsafe.Pointer(ptr))
	return int(val)
}

//pass
func B8() int {
	values := &PtrArray{ptr:unsafe.Pointer(nil), len:1}
	val := C.pass_goptrarr2((*C.PtrArray)(unsafe.Pointer(values)))
	newptr := (*PtrArray)(unsafe.Pointer(values.ptr))
	fmt.Println(val, values.len, newptr.ptr, newptr.len)
	C.MyFeee(values.ptr)
	return int(val)
}
//pass
func B8getstring() int {
	values := &PtrArray{ptr:unsafe.Pointer(nil), len:1}
	val := C.pass_goptrarr3((*C.PtrArray)(unsafe.Pointer(values)))
	newptr := (*PtrArray)(unsafe.Pointer(values.ptr))
	//str := C.GoStringN((*C.char)(newptr.ptr), C.int(newptr.len))
	//fmt.Println(val, values.len, str, newptr.ptr, newptr.len)
	C.MyFeee((unsafe.Pointer(newptr.ptr)))
	C.MyFeee(values.ptr)
	return int(val)
}

func Debug()  {
	C.SetValue(1)
}
func B8psssstring() int {
	argv := make([]*C.PtrArray, 1)
	str := unsafe.Pointer(C.CString("111"))
	argv[0] = constructPtrArray(str, 1 << 30)
	//argvv := &PtrArray{ptr:unsafe.Pointer(&argv[0]), len:1}
	argvv := constructPtrArray(unsafe.Pointer(&argv[0]), 1)
	val := int(C.pass_string((*C.PtrArray)(unsafe.Pointer(argvv)), argv[0], 0))
	defer C.free(str)
	defer destructPtrArray(argv[0])
	defer destructPtrArray(argvv)
	//defer C.free(unsafe.Pointer(argvv))
	return val
}

func GoString2CStringArr(strList []string) ([]*C.char, *C.int) {
	if len(strList) == 0 {
		return nil, nil
	}

	intList := make([]int32, len(strList))
	argv := make([]*C.char, len(strList))
	for i, s := range strList {
		cs := C.CString(s)
		argv[i] = cs
		intList[i] = int32(len(s))
	}
	return argv, (*C.int)(unsafe.Pointer(&intList[0]))
}
//pass
func B8passstr() int {
	args := []string{"11", "a22"}

	arr, intptr := GoString2CStringArr(args)
	//argv := make([]*C.char, len(args))
	//for i, s := range args {
	//	cs := C.CString(s)
	//	defer C.free(unsafe.Pointer(cs))
	//	argv[i] = cs
	//}
	//sstr := constructPtrArray(unsafe.Pointer(argv[1]), 3)
	//
	//intList := []int32{1, 22}
	sstr := constructPtrArray(unsafe.Pointer(arr[1]), 3)
	val := int(C.pass_string2(&arr[0], intptr, C.int(len(arr)), sstr, 0))
	destructPtrArray(sstr)
	for _, temp := range arr {
		C.free(unsafe.Pointer(temp))
	}
	return val
}
func teststr_direct() int {
	s := "11122"
	//addr := &s
	//hdr := (*reflect.StringHeader)(unsafe.Pointer(addr))
	sb := []byte(s)
	p, n := getBytesCInfo(sb)
	//fmt.Println(s)
	// reflect.StringHeader stores the Data field as a uintptr, not a pointer,
	// so ensure that the string remains reachable until the uintptr is converted.
	//runtime.KeepAlive(addr)
	val := int(C.pass_str_direct(p, C.int(n)))
	return val
}
func constructString(str string) *C.CustomString {
	p := (*C.CustomString)(C.MyMalloc(C.int(C.sizeof_CustomString)))
	(*p).str = C.CString(str)
	(*p).len = C.int(len(str))
	return p
}
func destructString(cstr *C.CustomString)  {
	C.free(unsafe.Pointer((*cstr).str))
	C.MyFeee(unsafe.Pointer(cstr))
}

func constructPtrArrayNew(ptr *unsafe.Pointer, len int) *C.PtrStrNewArray {
	ptrArr := (*C.PtrStrNewArray)(C.MyMalloc(C.sizeof_PtrStrNewArray))
	(*ptrArr).ptr = ptr
	(*ptrArr).len = C.int(len)
	return ptrArr
}
func constructPtrArray(ptr unsafe.Pointer, len int) *C.PtrArray {
	ptrArr := (*C.PtrArray)(C.MyMalloc(C.sizeof_struct_PtrArray))
	ptrArr.ptr = ptr
	ptrArr.len = C.int(len)
	return (*C.PtrArray)(unsafe.Pointer(ptrArr))
}
func destructPtrArray(cptrArr *C.PtrArray)  {
	C.MyFeee(unsafe.Pointer(cptrArr))
}

var dbAgent = RocksDB{talbeMap:make(map[string]*DBTable),}
func createNewDB(pathStr string, cfStr... string) unsafe.Pointer {
	cfarr, cfLenList := GoString2CStringArr(cfStr)
	defer func() {
		for _, temp := range cfarr { C.free(unsafe.Pointer(temp)) }
	}()
	path := C.CString(pathStr)
	defer C.free(unsafe.Pointer(path))

	handles := PtrArray{ptr:unsafe.Pointer(nil), len:0}
	names := PtrArray{ptr:unsafe.Pointer(nil), len:0}
	//val := C.pass_goptrarr2((*C.PtrArray)(unsafe.Pointer(values)))
	db := C.CreateNewDB(path, &cfarr[0], cfLenList, C.int(len(cfarr)), (*C.PtrArray)(unsafe.Pointer(&handles)), (*C.PtrArray)(unsafe.Pointer(&names)))
	//newptr := (*PtrArray)(unsafe.Pointer(values.ptr))
	fmt.Println("CreateNewDB:", handles.len, names.len)
	dbAgent.db = db

	handlesPtr := (*[1<<30]unsafe.Pointer)(handles.ptr)[0:handles.len:handles.len]
	namesPtr := (*[1<<30]unsafe.Pointer)(names.ptr)[0:handles.len:handles.len]
	for i := int32(0); i < handles.len; i++ {
		name := C.GoString((*C.char)(namesPtr[i]))
		fmt.Println("name", name, namesPtr[i])
		var table = &DBTable{handle:handlesPtr[i], name:name}
		dbAgent.talbeHandle = append(dbAgent.talbeHandle, table)
		dbAgent.talbeMap[name] = table
	}
	return db
}

func dbTablePut(db unsafe.Pointer, cf string, key []byte, value []byte) bool {
	keyp, keyn := getBytesCInfo(key)
	valuep, valuen := getBytesCInfo(value)
	C.DBColumnFamilyPut(db, dbAgent.talbeMap[cf].handle, keyp, C.int(keyn), valuep, C.int(valuen))
	return true
}
func dbNewPut(db unsafe.Pointer, cf string, key string, value string) bool  {
	ckey := C.CustomString{str:C.CString(key), len:C.int(len(key))}
	cvalue := C.CustomString{str:C.CString(value), len:C.int(len(value))}
	C.DBColumnFamilyPutOld(db, dbAgent.talbeMap[cf].handle, &ckey, &cvalue)
	return true
}

func dbTableGet(db unsafe.Pointer, cf string, key []byte) []byte {
	keyp, keyn := getBytesCInfo(key)
	myData := (C.DataStringPtr)(C.DBColumnFamilyGet(db, dbAgent.talbeMap[cf].handle, keyp, C.int(keyn)))
	defer C.ReleaseString(unsafe.Pointer(myData))
	return C.GoBytes(unsafe.Pointer(myData.Data), myData.Len)
}
func dbNewGet(db unsafe.Pointer, cf string, key string) string {
	ckey := C.CustomString{str:C.CString(key), len:C.int(len(key))}
	data := C.DBColumnFamilyGetOld(db, dbAgent.talbeMap[cf].handle, &ckey)
	myData := (C.DataStringPtr)(data)
	gob := C.GoStringN(myData.Data, myData.Len)
	return gob
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
	db := createNewDB("test3", "cf3", "cf2")
	fmt.Println(dbGet(db, "aaa"))
	dbPut(db, "aaa", "1111")
	fmt.Println(dbGet(db, "aaa"))

	fmt.Println(dbNewGet(db, "cf3", "aaaa"))
	dbNewPut(db, "cf3", "aaaa", "2222")
	fmt.Println(dbNewGet(db, "cf3", "aaaa"))
}
func LoadDB() {
	createNewDB("test3", "cf3", "cf2")
}
func Test4() int  {
	//fmt.Println(string(dbTableGet(dbAgent.db, "cf3", []byte("aaaa"))))
	dbTablePut(dbAgent.db, "cf3", []byte("aaaa"), []byte("1"))
	return len(string(dbTableGet(dbAgent.db, "cf3", []byte("aaaa"))))
}

var totalData = 100000
func Test5() int  {
	//fmt.Println(string(dbTableGet(dbAgent.db, "cf3", []byte("aaaa"))))
	//ret, _ := strconv.Atoi(string(dbTableGet(dbAgent.db, "cf3", []byte(strconv.Itoa(100)))))
	//fmt.Println("ret", ret)
	//start := time.Now()
	for i := 0; i < 1; i++ {
		dbTablePut(dbAgent.db, "cf3", []byte(strconv.Itoa(i)), []byte(strconv.Itoa(1)))
	}
	ret, _ := strconv.Atoi(string(dbTableGet(dbAgent.db, "cf3", []byte(strconv.Itoa(100)))))
	//fmt.Println("ret", time.Since(start), ret)
	return ret
}
func Test6() int  {
	//fmt.Println(string(dbTableGet(dbAgent.db, "cf3", []byte("aaaa"))))
	//start := time.Now()
	total := 0
	for i := 0; i < 1; i++ {
		temp, _ := strconv.Atoi(string(dbTableGet(dbAgent.db, "cf3", []byte(strconv.Itoa(i)))))
		total += temp
	}
	//fmt.Println("ret", time.Since(start), total)
	return total
}