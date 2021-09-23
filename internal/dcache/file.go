package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bales/internal/dcache/graph"
	capnp "zombiezen.com/go/capnproto2"
)

func main() {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	entry, err := graph.NewCacheEntry(seg)
	if err != nil {
		panic(err)
	}
	entry.SetArgsHash(420)

	fmt.Println(len(seg.Data()), seg.Data())
	// Write the message to stdout.
	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		panic(err)
	}
}
