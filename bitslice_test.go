package bitslice

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type BitSliceSuite struct{}

var _ = Suite(&BitSliceSuite{})

func (s *BitSliceSuite) TestNewEmptyBitSlice(c *C) {
	nebs := NewEmptyBitSlice()
	c.Assert(nebs.length, Equals, 0)
	c.Assert(nebs.data, HasLen, 0)
}

func (s *BitSliceSuite) TestNewBitSlice(c *C) {
	nbs := NewBitSlice(64)
	c.Assert(nbs.length, Equals, 64)
	c.Assert(nbs.data, HasLen, 1)

	nbs1 := NewBitSlice(96)
	c.Assert(nbs1.length, Equals, 96)
	c.Assert(nbs1.data, HasLen, 2)

	nbs2 := NewBitSlice(128)
	c.Assert(nbs2.length, Equals, 128)
	c.Assert(nbs2.data, HasLen, 2)
}

func (s *BitSliceSuite) TestUnusedSpace(c *C) {
	bs := NewBitSlice(65)
	c.Assert(bs.unusedSpace(), Equals, 63)

	bs1 := NewBitSlice(3)
	c.Assert(bs1.unusedSpace(), Equals, 61)

	bs2 := NewBitSlice(64)
	c.Assert(bs2.unusedSpace(), Equals, 0)

	bs3 := NewEmptyBitSlice()
	c.Assert(bs3.unusedSpace(), Equals, 0)
}

func (s *BitSliceSuite) TestDeepcopy(c *C) {
	bs := NewBitSlice(128)
	err := bs.Set(0)
	c.Assert(err, IsNil)
	err = bs.Set(1)
	c.Assert(err, IsNil)
	err = bs.Set(66)
	c.Assert(err, IsNil)

	new := bs.Deepcopy()
	c.Assert(new.length, Equals, bs.length)
	c.Assert(new.data[1], Equals, bs.data[1])
	c.Assert(new.data[0], Equals, bs.data[0])
	c.Assert(new, Not(Equals), bs)
}

func (s *BitSliceSuite) TestSet(c *C) {
	bs := NewBitSlice(96)
	err := bs.Set(96)
	c.Assert(err, ErrorMatches, "Index 96 out of range")

	err = bs.Set(95)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, uint64(2147483648))

	err = bs.Set(2)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, uint64(4))
}

func (s *BitSliceSuite) TestUnset(c *C) {
	bs := NewBitSlice(96)

	err := bs.Set(95)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, uint64(2147483648))
	err = bs.Unset(95)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, uint64(0))

	err = bs.Set(2)
	c.Assert(err, IsNil)
	err = bs.Set(3)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, uint64(12))
	err = bs.Unset(2)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, uint64(8))
}

func (s *BitSliceSuite) TestShiftLeftAndModify(c *C) {
	bs := NewBitSlice(64)
	err := bs.Set(0)
	c.Assert(err, IsNil)

	// Nothing changed
	bs.ShiftLeftAndModify(0)
	c.Assert(bs.data[0], Equals, uint64(1))
	c.Assert(bs.length, Equals, 64)
	c.Assert(len(bs.data), Equals, 1)

	bs.ShiftLeftAndModify(1)
	c.Assert(bs.data[0], Equals, uint64(2))
	c.Assert(bs.length, Equals, 65)
	c.Assert(len(bs.data), Equals, 2)

	bs.ShiftLeftAndModify(63)
	c.Assert(bs.data[1], Equals, uint64(1))
	c.Assert(bs.data[0], Equals, uint64(0))
	c.Assert(bs.length, Equals, 128)
	c.Assert(len(bs.data), Equals, 2)
}

func (s *BitSliceSuite) TestShiftRightAndModify(c *C) {
	bs := NewBitSlice(96)
	err := bs.Set(95)
	c.Assert(err, IsNil)

	// Nothing changed
	bs.ShiftRightAndModify(0)
	c.Assert(len(bs.data), Equals, 2)
	c.Assert(bs.data[1], Equals, uint64(2147483648))
	c.Assert(bs.data[0], Equals, uint64(0))
	c.Assert(bs.length, Equals, 96)

	bs.ShiftRightAndModify(31)
	c.Assert(len(bs.data), Equals, 2)
	c.Assert(bs.data[1], Equals, uint64(1))
	c.Assert(bs.data[0], Equals, uint64(0))
	c.Assert(bs.length, Equals, 65)

	bs.ShiftRightAndModify(63)
	c.Assert(len(bs.data), Equals, 1)
	c.Assert(bs.data[0], Equals, uint64(2))
	c.Assert(bs.length, Equals, 2)
}

func (s *BitSliceSuite) TestShiftLeft(c *C) {
	// TODO: Create this
}

func (s *BitSliceSuite) TestShiftRight(c *C) {
	// TODO: Create this
}

func (s *BitSliceSuite) AppendRight(c *C) {
	// TODO: Create this
}

func (s *BitSliceSuite) AppendLeft(c *C) {
	// TODO: Create this
}
