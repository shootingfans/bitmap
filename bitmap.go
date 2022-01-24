package bitmap

import (
	"math"
)

type bitmap struct {
	store  []int
	length int
}

func (b *bitmap) indexAndOffset(i int) (int, int) {
	return i / intSize, i % intSize
}

func (b *bitmap) Exists(i int) bool {
	d, o := b.indexAndOffset(i)
	return b.store[d]&(1<<o) > 0
}

func (b *bitmap) Store(i int) bool {
	if b.Exists(i) {
		return false
	}
	d, o := b.indexAndOffset(i)
	b.store[d] = b.store[d] | (1 << o)
	b.length++
	return true
}

func (b *bitmap) Remove(i int) {
	if !b.Exists(i) {
		return
	}
	d, o := b.indexAndOffset(i)
	b.store[d] = b.store[d] & ^(1 << o)
	b.length--
}

func (b *bitmap) Len() int {
	return b.length
}

func (b *bitmap) Cap() int {
	return cap(b.store) * intSize
}

func (b *bitmap) Range(f func(num int) bool) {
	var m, n int
	for m < b.length && n < len(b.store) {
		for i := 0; i < intSize; i++ {
			if b.store[n]&(1<<i) > 0 {
				m++
				if !f(n*intSize + i) {
					return
				}
			}
		}
		n++
	}
}

func (b *bitmap) RevRange(f func(num int) bool) {
	var m, n int
	n = len(b.store) - 1
	for m < b.length && n >= 0 {
		for i := intSize - 1; i >= 0; i-- {
			if b.store[n]&(1<<i) > 0 {
				m++
				if !f(n*intSize + i) {
					return
				}
			}
		}
		n--
	}
}

// Map is interface of a bitmap
type Map interface {
	// Exists returns whether the number is in the map
	Exists(int) bool
	// Store is store the number to map
	// if the number already in map will return false otherwise true
	Store(int) bool
	// Remove is remove the number from map
	Remove(int)
	// Len is return the number length
	Len() int
	// Cap is return the capacity of the map, it's the maximum number of the map
	Cap() int
	// Range is through the map from small to large
	Range(func(num int) bool)
	// RevRange is through the map from large to small
	RevRange(func(num int) bool)
}

const intSize = 32 << (^uint(0) >> 63)

// NewMap create a bitmap
// This bitmap is not  goroutine safety
func NewMap(capacity int) Map {
	return &bitmap{
		store:  make([]int, int(math.Ceil(float64(capacity)/float64(intSize)))),
		length: 0,
	}
}
