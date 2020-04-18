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

func createNewDB(pathStr string, cfStr... string) unsafe.Pointer {
	cfArr := make([]*C.CustomString, len(cfStr))
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