package bitmap

import (
	"fmt"
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

func ExampleSort() {
	// sort some number by small to large
	result := Sort(100, 30, 50, 90, 101, 300, 400, 5, 14, 99)
	for _, number := range result {
		fmt.Println(number)
	}
	// result:
	// 5
	// 14
	// 30
	// 50
	// 90
	// 99
	// 100
	// 101
	// 300
	// 400
}

func ExampleRevSort() {
	// sort some number by large to small
	result := Sort(100, 30, 50, 90, 101, 300, 400, 5, 14, 99)
	for _, number := range result {
		fmt.Println(number)
	}
	// result:
	// 400
	// 300
	// 101
	// 100
	// 99
	// 90
	// 50
	// 30
	// 14
	// 5
}

func ExampleNewSortMap() {
	// some numbers  100, 10, 95, 5, 131, 260
	numbers := []int{100, 10, 95, 5, 131, 260}
	m := NewSortMap(5, 260)
	for _, number := range numbers {
		m.Put(number)
	}
	// iterator numbers by small to large
	m.Range(func(number int) bool {
		fmt.Println(number)
		return true // return true to continue iterator
	})
	// result:
	// 5
	// 10
	// 95
	// 100
	// 131
	// 260

	// iterator numbers by large to small
	m.RevRange(func(number int) bool {
		fmt.Println(number)
		return true // return true to continue iterator
	})
	// result:
	// 260
	// 131
	// 100
	// 95
	// 10
	// 5
}
