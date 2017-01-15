package bitslice

import (
	"fmt"
)

const (
	MAXUINT = ^uint64(0)
)

type BitSlice struct {
	data   []uint64
	length int
}

func NewEmptyBitSlice() *BitSlice {
	return &BitSlice{
		data:   []uint64{},
		length: 0,
	}
}

func NewBitSlice(size int) *BitSlice {
	ceilSize := size / 64
	if size > ceilSize*64 {
		ceilSize++
	}
	return &BitSlice{
		data:   make([]uint64, ceilSize),
		length: size,
	}
}

func (t *BitSlice) Get(index int) (int, error) {
	if index > t.length-1 {
		return 0, fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 64
	offset := uint(index % 64)
	if (t.data[dataIndex] & (uint64(1) << offset)) >= uint64(1) {
		return 1, nil
	}

	return 0, nil
}

func (t *BitSlice) Set(index int) error {
	if index < 0 || index > t.length-1 {
		return fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 64
	offset := uint(index % 64)
	t.data[dataIndex] |= (uint64(1) << offset)

	return nil
}

func (t *BitSlice) Or(bslice *BitSlice) {
	if len(bslice.data) > len(t.data) {
		diff := len(bslice.data) - len(t.data)
		pre := make([]uint64, diff)
		t.data = append(pre, t.data...)
	}

	for i, b := range bslice.data {
		t.data[i] |= b
	}
}

func (t *BitSlice) And(bslice *BitSlice) {
	// Note: overflow bits on the larger of the two slices
	// will be zeroed out
	if len(bslice.data) > len(t.data) {
		pre := bslice.data[len(t.data):]
		t.data = append(pre, t.data...)
	}

	for i, b := range bslice.data {
		t.data[i] &= b
	}
}

func (t *BitSlice) Unset(index int) error {
	if index < 0 || index > t.length-1 {
		return fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 64
	offset := uint(index % 64)
	t.data[dataIndex] &^= (uint64(1) << offset)

	return nil
}

func (t *BitSlice) Append(bslice *BitSlice) {
	// TODO: FILL THIS IN
	// Resize
	//newDataSize := (t.length + bslice.length) / 8
	unused := 0
	if t.length%8 != 0 {
		unused = 8 - t.length%8
	}
	fmt.Println(unused)
}

func (t *BitSlice) unusedSpace() int {
	unused := 0
	if t.length%64 != 0 {
		unused = 64 - t.length%64
	}
	return unused
}

func (t *BitSlice) shiftLeft(num int) {
	// Shifts the bitslice over for up to 64 places (1 uint64) left
	unused := t.unusedSpace()

	// Expand the data in uint64 if necessary
	shifted := 0
	if num > unused {
		shifted = (num-unused)/64 + 1
		bs := make([]uint64, shifted)
		t.data = append(t.data, bs...)
	}

	t.length += num
	newUnused := t.unusedSpace()
	if newUnused > unused {
		carry := uint64(0)
		for i, _ := range t.data {
			old := t.data[i]
			t.data[i] >>= uint(newUnused - unused)
			t.data[i] |= (carry << uint(64-newUnused-unused))
			carry = old & (MAXUINT >> uint(64-newUnused-unused))
		}
	} else if unused > newUnused {
		carry := uint64(0)
		for i := len(t.data) - 1; i >= 0; i-- {
			old := t.data[i]
			t.data[i] <<= uint(unused - newUnused)
			t.data[i] |= (carry >> uint(64-unused-newUnused))
			carry = old & (MAXUINT >> uint(64-unused-newUnused))
		}
	}
}

func (t *BitSlice) upperBits(num int) uint64 {
	// Returns the most significant bits (up to 64) from a bitslice
	b := uint64(0)
	if num > 64 {
		return b
	}

	bitsLeft := num
	numBitsUpper := t.length % 64
	dataNum := t.length / 64
	for bitsLeft > 0 {
		if bitsLeft >= numBitsUpper {
			b |= t.data[dataNum]
			bitsLeft -= numBitsUpper
			dataNum--
			if dataNum < 0 {
				bitsLeft = 0
			}
			b <<= uint(bitsLeft)
		} else {
			b |= t.data[dataNum] >> uint(8-bitsLeft)
			bitsLeft = 0
		}
	}
	return b
}
