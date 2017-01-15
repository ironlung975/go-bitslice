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

func (s *BitSliceSuite) TestOr(c *C) {
	// Test working on equally sized bitslices
	{
		bs1 := NewBitSlice(64)
		err := bs1.Set(2)
		c.Assert(err, IsNil)
		err = bs1.Set(6)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(64)
		err = bs2.Set(5)
		c.Assert(err, IsNil)

		bs1.Or(bs2)
		c.Assert(bs1.data[0], Equals, uint64(100))
	}

	// Test working on unequally sized bitslices represented by same number of buffers
	{
		bs1 := NewBitSlice(128)
		err := bs1.Set(70)
		c.Assert(err, IsNil)
		err = bs1.Set(67)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(80)
		err = bs2.Set(2)
		c.Assert(err, IsNil)
		err = bs2.Set(64)
		c.Assert(err, IsNil)

		bs1.Or(bs2)
		c.Assert(bs1.data[1], Equals, uint64(73))
		c.Assert(bs1.data[0], Equals, uint64(4))
	}

	// Test that it cuts off correctly
}
