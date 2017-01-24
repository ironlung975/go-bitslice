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

func (t *BitSlice) Unset(index int) error {
	if index < 0 || index > t.length-1 {
		return fmt.Errorf("Index %d out of range", index)
	}

	dataIndex := index / 64
	offset := uint(index % 64)
	t.data[dataIndex] &^= (uint64(1) << offset)

	return nil
}

func (t *BitSlice) ShiftLeftAndModify(amount int) {
	if amount > 0 {
		add := amount / 64
		if add > 0 {
			extra := make([]uint64, add)
			t.data = append(extra, t.data...)
			t.length += 64 * add
			amount -= 64 * add
		}

		unused := t.unusedSpace()
		carryAmount := amount - unused
		if unused < amount {
			extra := make([]uint64, 1)
			t.data = append(t.data, extra...)
		}

		carry := uint64(0)
		for i := 0; i < len(t.data); i++ {
			shifted := t.data[i] << uint(amount)
			newCarry := uint64(0) | (t.data[i] >> uint(64-unused-carryAmount))
			shifted |= carry
			carry = newCarry
			t.data[i] = shifted
		}

		t.length += amount
	}
}

func (t *BitSlice) ShiftRightAndModify(amount int) {
	if amount > 0 {
		move := amount / 64
		for i := move; i > 0; i-- {
			t.data[i] = t.data[i-1]
			t.length -= 64
		}

		if move > 0 {
			t.data = t.data[:len(t.data)-1-move]
		}
		amount -= move * 64

		unused := t.unusedSpace()
		carryAmount := amount
		if 64-unused < amount {
			carryAmount = amount - (64 - unused)
		}

		carry := (t.data[len(t.data)-1] & createMask(carryAmount)) << uint(64-carryAmount-1)
		t.data[len(t.data)-1] >>= uint(amount)
		for i := len(t.data) - 1; i > 0; i-- {
			newCarry := (t.data[i-1] & createMask(carryAmount)) << uint(64-carryAmount-1)
			t.data[i-1] >>= uint(amount)
			t.data[i-1] |= carry
			carry = newCarry
		}

		if carryAmount != amount {
			t.data = t.data[:len(t.data)-1]
		}

		t.length -= amount
	}
}

func (t *BitSlice) Deepcopy() *BitSlice {
	nbs := NewBitSlice(t.length)
	copy(nbs.data, t.data)
	return nbs
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
