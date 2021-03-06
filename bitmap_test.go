package bitmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	m := NewMap(1024)
	assert.Equal(t, m.Cap(), 1024)
	assert.Equal(t, m.Len(), 0)
	assert.False(t, m.Exists(5))
	assert.True(t, m.Store(5))
	assert.True(t, m.Exists(5))
	assert.False(t, m.Store(5))
	assert.Equal(t, m.Len(), 1)
	assert.True(t, m.Store(100))
	assert.Equal(t, m.Len(), 2)
	assert.True(t, m.Store(53))
	assert.Equal(t, m.Len(), 3)
	var s2b, b2s []int
	m.Range(func(num int) bool {
		s2b = append(s2b, num)
		return true
	})
	m.RevRange(func(num int) bool {
		b2s = append(b2s, num)
		return true
	})
	assert.Equal(t, s2b, []int{5, 53, 100})
	assert.Equal(t, b2s, []int{100, 53, 5})
	m.Store(105)
	m.Store(1)
	m.Store(98)
	var b2s1, s2b1 []int
	m.Range(func(num int) bool {
		s2b1 = append(s2b1, num)
		return len(s2b1) <= 4
	})
	m.RevRange(func(num int) bool {
		b2s1 = append(b2s1, num)
		return len(b2s1) <= 2
	})
	assert.Equal(t, s2b1, []int{1, 5, 53, 98, 100})
	assert.Equal(t, b2s1, []int{105, 100, 98})
	assert.Equal(t, m.Len(), 6)
	m.Remove(53)
	assert.Equal(t, m.Len(), 5)
	assert.False(t, m.Exists(53))
	s2b1 = []int{}
	m.Range(func(num int) bool {
		s2b1 = append(s2b1, num)
		return len(s2b1) <= 4
	})
	assert.Equal(t, s2b1, []int{1, 5, 98, 100, 105})
	m.Remove(111)
	assert.Equal(t, m.Len(), 5)
}

func BenchmarkNewMap(b *testing.B) {
	m := NewMap(0xffff)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Store(i % 0xffff)
	}
}

func ExampleNewMap() {
	m := NewMap(0xff) // store number 0 ~ 0xff
	m.Exists(100)     // this will return false
	m.Store(100)      // store 100 to map
	m.Exists(100)     // this will return true
	m.Remove(100)     // remove 100 from map
	m.Exists(100)     // this will return false
}

func ExampleBitmap_Range() {
	m := NewMap(0xff) // store number 0 ~ 0xff
	// store some numbers
	m.Store(50)
	m.Store(60)
	m.Store(10)
	m.Store(100)
	m.Store(34)
	// iterator the numbers by small to large
	m.Range(func(num int) bool {
		fmt.Println(num)
		return true // return true to continue iterator
	})
	// result:
	// 10
	// 34
	// 50
	// 60
	// 100
}
