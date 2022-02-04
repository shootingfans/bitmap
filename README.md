[![Test](https://github.com/shootingfans/bitmap/actions/workflows/go.yml/badge.svg)](https://github.com/shootingfans/bitmap/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/shootingfans/bitmap/branch/main/graph/badge.svg?token=SYGDITZ3JT)](https://codecov.io/gh/shootingfans/bitmap)
[![GoDoc](https://godoc.org/github.com/shootingfans/bitmap?status.png)](https://godoc.org/github.com/shootingfans/bitmap)
[![GoReport](https://goreportcard.com/badge/github.com/shootingfans/bitmap)](https://goreportcard.com/report/github.com/shootingfans/bitmap)

# bitmap

A bitmap struct realize by golang.

## introduction

The bitmap structure is use bit to store **unrepeated** numbers.
In this package bitmap will use `[]int` store numbers, it can be used in 32bit or 64bit environment.

## interface

### Map

The Map interface is used to store numbers and find whether exists.

### SortMap

The SortMap interface is used to sort numbers.