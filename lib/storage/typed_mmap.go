package storage

import (
	"log"
	"os"
	"reflect"
	"unsafe"
)

type index_type interface {
	~uint16 | ~uint32
}

type typedData[T comparable, I index_type] struct {
	file *os.File
	d    []T
}

func (t typedData[T, I]) Get(i I) *T {
	return &t.d[i]
}

func MmapTyped[T comparable, I index_type](path string) typedData[T, I] {
	// Check if T is a primitive type
	var t T
	if !isPlainOldData(reflect.TypeOf(t)) {
		log.Fatalf("Type %T is not plain old data", t)
	}

	expected_size := ARRAY_SIZE * int(unsafe.Sizeof(t))

	// Create the file if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
		os.Truncate(path, int64(expected_size))
	}

	// Map the file as a byte slice
	data := mmap(path)

	if len(data.d) != expected_size {
		log.Fatalf("File %s is %d bytes, expected %d bytes", path, len(data.d), expected_size)
	}

	// (unsafe) Convert the byte slice to a typed slice
	slice := unsafe.Slice((*T)(unsafe.Pointer(&data.d[0])), ARRAY_SIZE)

	return typedData[T, I]{data.f, slice}
}

func isPlainOldData(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if !isPlainOldData(t.Field(i).Type) {
				return false
			}
		}
		return true
	case reflect.Array:
		return isPlainOldData(t.Elem())
	default:
		return false
	}
}
