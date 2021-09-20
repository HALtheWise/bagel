// Code generated by capnpc-go. DO NOT EDIT.

package graph

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type CacheEntry struct{ capnp.Struct }

// CacheEntry_TypeID is the unique identifier for the type CacheEntry.
const CacheEntry_TypeID = 0x90bfaf0383e76e8a

func NewCacheEntry(s *capnp.Segment) (CacheEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return CacheEntry{st}, err
}

func NewRootCacheEntry(s *capnp.Segment) (CacheEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return CacheEntry{st}, err
}

func ReadRootCacheEntry(msg *capnp.Message) (CacheEntry, error) {
	root, err := msg.RootPtr()
	return CacheEntry{root.Struct()}, err
}

func (s CacheEntry) String() string {
	str, _ := text.Marshal(0x90bfaf0383e76e8a, s.Struct)
	return str
}

func (s CacheEntry) Id() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s CacheEntry) SetId(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

func (s CacheEntry) MaybeDirty() bool {
	return s.Struct.Bit(32)
}

func (s CacheEntry) SetMaybeDirty(v bool) {
	s.Struct.SetBit(32, v)
}

func (s CacheEntry) ChangedAt() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s CacheEntry) SetChangedAt(v int32) {
	s.Struct.SetUint32(8, uint32(v))
}

func (s CacheEntry) Dependencies() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Int32List{List: p.List()}, err
}

func (s CacheEntry) HasDependencies() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CacheEntry) SetDependencies(v capnp.Int32List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewDependencies sets the dependencies field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s CacheEntry) NewDependencies(n int32) (capnp.Int32List, error) {
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s CacheEntry) Dependents() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.Int32List{List: p.List()}, err
}

func (s CacheEntry) HasDependents() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s CacheEntry) SetDependents(v capnp.Int32List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewDependents sets the dependents field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s CacheEntry) NewDependents(n int32) (capnp.Int32List, error) {
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s CacheEntry) ArgsHash() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s CacheEntry) SetArgsHash(v int64) {
	s.Struct.SetUint64(16, uint64(v))
}

func (s CacheEntry) Result() (capnp.Pointer, error) {
	return s.Struct.Pointer(3)
}

func (s CacheEntry) HasResult() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s CacheEntry) ResultPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(3)
}

func (s CacheEntry) SetResult(v capnp.Pointer) error {
	return s.Struct.SetPointer(3, v)
}

func (s CacheEntry) SetResultPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(3, v)
}

func (s CacheEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s CacheEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CacheEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CacheEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

// CacheEntry_List is a list of CacheEntry.
type CacheEntry_List struct{ capnp.List }

// NewCacheEntry creates a new list of CacheEntry.
func NewCacheEntry_List(s *capnp.Segment, sz int32) (CacheEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4}, sz)
	return CacheEntry_List{l}, err
}

func (s CacheEntry_List) At(i int) CacheEntry { return CacheEntry{s.List.Struct(i)} }

func (s CacheEntry_List) Set(i int, v CacheEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s CacheEntry_List) String() string {
	str, _ := text.MarshalList(0x90bfaf0383e76e8a, s.List)
	return str
}

// CacheEntry_Promise is a wrapper for a CacheEntry promised by a client call.
type CacheEntry_Promise struct{ *capnp.Pipeline }

func (p CacheEntry_Promise) Struct() (CacheEntry, error) {
	s, err := p.Pipeline.Struct()
	return CacheEntry{s}, err
}

func (p CacheEntry_Promise) Result() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(3)
}

const schema_85d3acc39d94e0f8 = "x\xdad\xd01k\x14Q\x14\x05\xe0s\xee}\xb3\xdb" +
	"\x84\xac/;\x82J\x8a(\x0aja\xb0M\xa3\xa2\x82" +
	"X\xe5e\xba4\xf2\x9cy\xecLH\x1e\xc3\xccXl" +
	"-6V\x16\x96\xfa\x0f\x04\x7f\x82\x85\x85\x8d\x8d\xa8`" +
	"\xa1\xb0\x8d\x18\xc1\x1f`7\xf2\x14A\xb0\xb9\\\xbes" +
	"\xaas\xe2\xd9u\xb1\xd9=\xc0\x99l2>\x8e\xdf\x1e" +
	"\xea\xcbWO\xe0\xceQ\xc7\x9f\xab\xa7\xcf_\xbfx\xff" +
	"\x08\x99\x99\x02WO\x1e\xd0^J\xcf\x857\x04\xc7&" +
	"\x0e\xa1\x8b\xfe\xd0lW\xa5/\xeb\xb0\xbd\xe8|[\xff" +
	"\xb9WJ\xdf\xc6v\xe7f\xf2\xdb\xb38t\xcb]\xd2" +
	"m\xa9\x01\x0c\x01\xfb\xee\x0c\xe0\xde*\xdd'\xa1\xe5V" +
	"\xce\x84\x1f\xf7\x01\xf7A\xe9VB+\x92S\x00\xfbe" +
	"\x0fp\x9f\x95\xeeX\xc8iN\x05\xec\xd7\xcb\x80[)" +
	"\xdd\x0f\xa1U\xe64\x80\xfd~\x00\xb8cea(\xb4" +
	"Frf\xc0\x9c\xdc\x07\xf6\xa8,6\x13g\x92s\x02" +
	"\xccO\xf3.P\x9cJ~>\xf9DsN\x81\xf9Y" +
	"\xee\x00\xc5f\xf2\x8b\x14jS\xd1@h\xc0\xf1\xc8/" +
	"\xef\x87[M\x07\x1d\x96$\x84i\x83\xb2\xf6q\x11\xaa" +
	"\x1b\xe0\xf0\xb78\x8b\xfe(p\x0d\xc25p\xacB\x1b" +
	"b\x15\"fe\x13z\xae\x83\xbb\xca\xdf\xd5\xf5\x7fS" +
	"\x1d\xfe\xcb|\xb7\xe8\xef\xf8\xbe\x06\xc0\x0c\xc2\x0c\xbc\xd6" +
	"\x85\xfe\xc1\xe1\xc0\x0d\x087\xc0_\x01\x00\x00\xff\xff\xce" +
	"\xe6Yg"

func init() {
	schemas.Register(schema_85d3acc39d94e0f8,
		0x90bfaf0383e76e8a)
}
