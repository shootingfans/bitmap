package bitmap

import "math"

// SortMap is fast sort bitmap
type SortMap interface {
	// Put is put number to bitmap
	Put(i int)
	// Sort is return all number in bitmap by small to large
	Sort() []int
	// RSort is return all number in bitmap by large to small
	RSort() []int
	// Range is iterate the number in bitmap by small to large
	// if handler function return false the iterator will stop
	Range(f func(int) bool)
	// RevRange is iterate the number in bitmap by large to small
	// if handler function return false the iterator will stop
	RevRange(f func(int) bool)
}

type sortBitMap struct {
	store  []int
	offset int
	count  int
}

func (s *sortBitMap) indexAndOffset(i int) (int, int) {
	return (i - s.offset) / intSize, (i - s.offset) % intSize
}

func (s *sortBitMap) Put(i int) {
	i, o := s.indexAndOffset(i)
	s.store[i] = s.store[i] | (1 << o)
	s.count++
}

func (s *sortBitMap) Sort() []int {
	res := make([]int, 0, s.count)
	s.Range(func(i int) bool {
		res = append(res, i)
		return true
	})
	return res
}

func (s *sortBitMap) RSort() []int {
	res := make([]int, 0, s.count)
	s.RevRange(func(i int) bool {
		res = append(res, i)
		return true
	})
	return res
}

func (s *sortBitMap) Range(f func(int) bool) {
	var m, n int
	for m < s.count && n < len(s.store) {
		for i := 0; i < intSize; i++ {
			if s.store[n]&(1<<i) > 0 {
				m++
				if !f(s.offset + n*intSize + i) {
					return
				}
			}
		}
		n++
	}
}

func (s *sortBitMap) RevRange(f func(int) bool) {
	var m, n int
	n = len(s.store) - 1
	for m < s.count && n >= 0 {
		for i := intSize - 1; i >= 0; i-- {
			if s.store[n]&(1<<i) > 0 {
				m++
				if !f(s.offset + n*intSize + i) {
					return
				}
			}
		}
		n--
	}
}

// NewSortMap create sortmap by min and max number
func NewSortMap(min, max int) SortMap {
	return &sortBitMap{
		store:  make([]int, int(math.Ceil(float64(max-min)/float64(intSize)))),
		offset: min,
	}
}

func buildSortMap(number int, numbers ...int) SortMap {
	min, max := number, number
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	m := NewSortMap(min, max)
	m.Put(number)
	for i := 0; i < len(numbers); i++ {
		m.Put(numbers[i])
	}
	return m
}

func Sort(number int, numbers ...int) []int {
	return buildSortMap(number, numbers...).Sort()
}

func RevSort(number int, numbers ...int) []int {
	return buildSortMap(number, numbers...).RSort()
}
