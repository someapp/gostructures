package gostructures

import "math/bits"

type BitVector struct {
	ints   []uint64 // can store 64 bits per int
	length int
}

func NewBitVector(length int) BitVector {
	ints := make([]uint64, ((length-1)/64)+1)

	return BitVector{
		ints: ints,
	}
}

func (bv *BitVector) Add(index uint) {
	inti := index / 64
	mask := uint64(1 << (index % 64))

	bv.ints[inti] = bv.ints[inti] | mask
}

func (bv *BitVector) Contains(index uint) bool {
	inti := index / 64
	mask := uint64(1 << (index % 64))

	return (bv.ints[inti] & mask) != 0
}

func (bv *BitVector) Count() int {
	count := 0
	for _, ints := range bv.ints {
		count += bits.OnesCount64(ints)
	}
	return count
}
