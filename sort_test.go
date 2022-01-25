package bitmap

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSortMap(t *testing.T) {
	m := NewSortMap(0xff, 0xffff).(*sortBitMap)
	assert.Equal(t, m.offset, 0xff)
	assert.Equal(t, len(m.store), int(math.Ceil(float64(0xffff-0xff)/float64(intSize))))
	m.Put(0xff)
	assert.Equal(t, m.store[0], 1)
	m.Put(0x10A)
	m.Put(0x15B)
	m.Put(0x13C)
	m.Put(0x11D)
	m.Put(0x1AF)
	assert.Equal(t, m.Sort(), []int{0xff, 0x10A, 0x11D, 0x13C, 0x15B, 0x1AF})
	assert.Equal(t, m.RSort(), []int{0x1AF, 0x15B, 0x13C, 0x11D, 0x10A, 0xff})
	m.Range(func(i int) bool {
		assert.Equal(t, i, 0xff)
		return false
	})
	m.RevRange(func(i int) bool {
		assert.Equal(t, i, 0x1AF)
		return false
	})
}

func TestRevSort(t *testing.T) {
	assert.Equal(t,
		Sort(100, 200, 51, 16, 45, 65, 154, 103, 99),
		[]int{16, 45, 51, 65, 99, 100, 103, 154, 200},
	)
}

func TestSort(t *testing.T) {
	assert.Equal(t,
		RevSort(908, 100, 31, 554, 193, 1003, 101, 200, 20, 431, 183, 19),
		[]int{1003, 908, 554, 431, 200, 193, 183, 101, 100, 31, 20, 19},
	)
}
