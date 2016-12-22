package bitslice

import (
	"fmt"
)

type BitSlice struct {
	data   []byte
	length int
}

func NewEmptyBitSlice() *BitSlice {
	return &BitSlice{
		data:   []byte{},
		length: 0,
	}
}

func NewBitSlice(size int) *BitSlice {
	return &BitSlice{
		data:   make([]byte, size/8),
		length: size,
	}
}

func (t *BitSlice) Get(index int) (int, error) {
	if index > t.size {
		return 0, fmt.Errorf("Index %d out of range")
	}

	dataIndex := index / 8
	offset := index % 8
	if (t.data[dataIndex] & (byte(1) << offset)) > byte(1) {
		return 1, nil
	}

	return 0, nil
}

func (t *BitSlice) Append(bslice *BitSlice) {
	// TODO: FILL THIS IN
}
