// Code generated by capnpc-go. DO NOT EDIT.

package books

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Book struct{ capnp.Struct }

// Book_TypeID is the unique identifier for the type Book.
const Book_TypeID = 0x8100cc88d7d4d47c

func NewBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func NewRootBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func ReadRootBook(msg *capnp.Message) (Book, error) {
	root, err := msg.RootPtr()
	return Book{root.Struct()}, err
}

func (s Book) String() string {
	str, _ := text.Marshal(0x8100cc88d7d4d47c, s.Struct)
	return str
}

func (s Book) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Book) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Book) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Book) SetTitle(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Book) PageCount() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Book) SetPageCount(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// Book_List is a list of Book.
type Book_List struct{ capnp.List }

// NewBook creates a new list of Book.
func NewBook_List(s *capnp.Segment, sz int32) (Book_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Book_List{l}, err
}

func (s Book_List) At(i int) Book { return Book{s.List.Struct(i)} }

func (s Book_List) Set(i int, v Book) error { return s.List.SetStruct(i, v.Struct) }

func (s Book_List) String() string {
	str, _ := text.MarshalList(0x8100cc88d7d4d47c, s.List)
	return str
}

// Book_Promise is a wrapper for a Book promised by a client call.
type Book_Promise struct{ *capnp.Pipeline }

func (p Book_Promise) Struct() (Book, error) {
	s, err := p.Pipeline.Struct()
	return Book{s}, err
}

const schema_85d3acc39d94e0f8 = "x\xda\x120u`\x12d\x8dg`\x08dae\xdb" +
	"_s\xe5\xca\xf5\x8e3\x8d\x81J\x8c\x8c\xff\x7f<\x98" +
	"2\xf7\xf0\x9a\xcb\xad\x0c\xac\x8c\xec\x0c\x0c\x86\xa2V\x8c" +
	"\x82\xaa\xec\x0c\x0c\x82\x8a\xe5\x0c\x8c\xff3\xf3JR\x8b" +
	"\xf2\x12s\x98\xf5S\x92\x13\x933R\xf5\x93\xf2\xf3\xb3" +
	"\x8b!\xa4^rbA\x9e}\x81\x95S~~v\x00" +
	"#c \x073\x0b\x03\x03\x0b#\x03\x83\xa0\xa6\x11\x03" +
	"C\xa0\x0a3c\xa0\x01\x13##\xa3\x08#HL7" +
	"\x88\x81!P\x87\x991\xd0\x82\x89Q\xbe$\xb3$'" +
	"\x95\x91\x87\x81\x89\x91\x87\x81\xf1\x7fAbz\xaas~" +
	"i\x1e\x03c\x09#\x0b\x03\x13#\x0b\x03# \x00\x00" +
	"\xff\xff\xf0#*O"

func init() {
	schemas.Register(schema_85d3acc39d94e0f8,
		0x8100cc88d7d4d47c)
}
