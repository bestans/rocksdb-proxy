package dll

import "C"
import (
	"reflect"
	"unsafe"
)

type DBTable struct {
	handle unsafe.Pointer
	name string
}
type RocksDB struct {
	db unsafe.Pointer
	talbeHandle []*DBTable
	talbeMap map[string]*DBTable
}

func getBytesCInfo(b []byte) (*C.char, int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return (*C.char)(unsafe.Pointer(hdr.Data)), hdr.Len
}
