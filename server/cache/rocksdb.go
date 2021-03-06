package cache

// #include "rocksdb/c.h"
// #cgo CFLAGS: -I/Users/yrc/go/src/rcache/rocksdb/include
// #cgo LDFLAGS: -L/Users/yrc/go/src/rcache/rocksdb -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"

type rocksdbCache struct {
	db *C.rocksdb_t
	ro *C.rocksdb_readoptions_t
	wo *C.rocksdb_writeoptions_t
	e  *C.char
	ch chan *pair
}
