package storage

import (
	"log"
	"os"
	"syscall"
)

// An mmapData is mmap'ed read-only data from a file.
type mmapData struct {
	f *os.File
	d []byte
}

// mmap maps the given file into memory.
func mmap(file string) mmapData {
	f, err := os.OpenFile(file, os.O_RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}
	return mmapFile(f)
}

func mmapFile(f *os.File) mmapData {
	st, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := st.Size()
	if int64(int(size+4095)) != size+4095 {
		log.Fatalf("%s: too large for mmap", f.Name())
	}
	n := int(size)
	if n == 0 {
		return mmapData{f, nil}
	}
	data, err := syscall.Mmap(int(f.Fd()), 0, (n+4095)&^4095, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("mmap %s: %v", f.Name(), err)
	}
	return mmapData{f, data[:n]}
}
