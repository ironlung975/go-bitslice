package byteslice

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
	ceilSize := size / 8
	if size > ceilSize*8 {
		ceilSize++
	}
	return &BitSlice{
		data:   make([]byte, ceilSize),
		length: size,
	}
}

func (t *BitSlice) Get(index int) (int, error) {
	if index > t.length-1 {
		return 0, fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 8
	offset := uint(index % 8)
	if (t.data[dataIndex] & (byte(1) << offset)) >= byte(1) {
		return 1, nil
	}

	return 0, nil
}

func (t *BitSlice) Set(index int) error {
	if index < 0 || index > t.length-1 {
		return fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 8
	offset := uint(index % 8)
	t.data[dataIndex] |= (byte(1) << offset)

	return nil
}

func (t *BitSlice) Or(bslice *BitSlice) {
	if len(bslice.data) > len(t.data) {
		diff := len(bslice.data) - len(t.data)
		pre := make([]byte, diff)
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

	dataIndex := index / 8
	offset := uint(index % 8)
	t.data[dataIndex] &^= (byte(1) << offset)

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
	if t.length%8 != 0 {
		unused = 8 - t.length%8
	}
	return unused
}

func (t *BitSlice) shiftLeft(num int) {
	// Shifts the bitslice over for up to 8 places (1 byte) left
	unused := t.unusedSpace()

	// Expand the data in bytes if necessary
	shiftedBytes := 0
	if num > unused {
		shiftedBytes = (num-unused)/8 + 1
		bs := make([]byte, shiftedBytes)
		t.data = append(t.data, bs...)
	}

	t.length += num
	newUnused := t.unusedSpace()
	if newUnused > unused {
		carry := byte(0)
		for i, _ := range t.data {
			old := t.data[i]
			t.data[i] >>= uint(newUnused - unused)
			t.data[i] |= (carry << uint(8-newUnused-unused))
			carry = old & (byte(255) >> uint(8-newUnused-unused))
		}
	} else if unused > newUnused {
		carry := byte(0)
		for i := len(t.data) - 1; i >= 0; i-- {
			old := t.data[i]
			t.data[i] <<= uint(unused - newUnused)
			t.data[i] |= (carry >> uint(8-unused-newUnused))
			carry = old & (byte(255) >> uint(8-unused-newUnused))
		}
	}
}

func (t *BitSlice) upperBits(num int) byte {
	// Returns the most significant bits (up to a byte) from a bitslice
	b := byte(0)
	if num > 8 {
		return b
	}

	bitsLeft := num
	numBitsUpperByte := t.length % 8
	byteNum := t.length / 8
	for bitsLeft > 0 {
		if bitsLeft >= numBitsUpperByte {
			b |= t.data[byteNum]
			bitsLeft -= numBitsUpperByte
			byteNum--
			if byteNum < 0 {
				bitsLeft = 0
			}
			b <<= uint(bitsLeft)
		} else {
			b |= t.data[byteNum] >> uint(8-bitsLeft)
			bitsLeft = 0
		}
	}
	return b
}
