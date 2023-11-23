package id

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"reflect"
	"strconv"
	"unsafe"
)

// GenID generates an ID according to the raw material.
func GenID(raw string) string {
	if raw == "" {
		return ""
	}
	sh := &reflect.SliceHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(&raw)).Data,
		Len:  len(raw),
		Cap:  len(raw),
	}
	p := *(*[]byte)(unsafe.Pointer(sh))

	res := crc32.ChecksumIEEE(p)
	return fmt.Sprintf("%x", res)
}

// ID is the type of the id field used for any entities
type ID uint64

// String indicates how to convert ID to a string.
func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON is the way to encode ID to JSON string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(id), 10))
}

// UnmarshalJSON is the way to decode ID from JSON string.
func (id *ID) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		*id = ID(u)
	default:
		panic("unknown type")
	}
	return nil
}

// IDGenerator is an interface for generating IDs.
type IDGenerator interface {
	// NextID generates an ID.
	NextID() ID
}
