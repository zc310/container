package list

import (
	"bytes"
	"fmt"
)

// List very useful general purpose list container
type List struct {
	elements []interface{}
	size     int
	capacity int
}

// New returns an initialized list.
func New() *List {
	return new(List)
}

// Count The number of items in the list
func (p List) Count() int {
	return p.size
}

// Empty If an empty list
func (p List) Empty() bool {
	return p.size == 0
}

// Add Add an item to the list
func (p *List) Add(v interface{}) int {
	if p.size == p.capacity {
		p.grow()
	}
	p.elements[p.size] = v
	p.size++
	return p.size - 1
}

// Get Gets an item from the list by its list positiont
func (p List) Get(i int) (interface{}, bool) {
	if !p.inRange(i) {
		return nil, false
	}

	return p.elements[i], true
}

// First Gets the first item in the list
func (p *List) First() (interface{}, bool) {
	return p.Get(0)
}

// Clear Removes all list items, setting the Count to 0
func (p *List) Clear() {
	p.elements = make([]interface{}, 0)
	p.size = 0
	p.capacity = 0
}

// Delete Removes an item from the list by its list position
func (p *List) Delete(i int) {
	if !p.inRange(i) {
		return
	}
	p.elements = append(p.elements[:i], p.elements[i+1:]...)
	p.size--
	p.shrink()
}

// Last Gets the last item in the list
func (p *List) Last() (interface{}, bool) {
	return p.Get(p.size - 1)
}

// IndexOf Gives the list position of a specified object in the list
func (p *List) IndexOf(v interface{}) int {
	var i int
	for i < p.size {
		if p.elements[i] == v {
			return i
		}
		i++
	}
	return -1
}

// Insert Inserts a new item into the list at a given index position
func (p *List) Insert(i int, v interface{}) {
	p.elements = append(p.elements, 0)
	copy(p.elements[i+1:], p.elements[i:])
	p.elements[i] = v
	p.size++
}

// Swap Moves an item to a new list position
func (p *List) Swap(a, b int) {
	if p.inRange(a) && p.inRange(b) {
		p.elements[a], p.elements[b] = p.elements[b], p.elements[a]
	}
}

// Put Set a new item into the list at a given index position
func (p *List) Put(i int, v interface{}) bool {
	if !p.inRange(i) {
		return false
	}
	p.elements[i] = v
	return true
}
func (p *List) inRange(index int) bool {
	return index >= 0 && index < p.size
}

// Remove Removes an item from the list by its object
func (p *List) Remove(v interface{}) int {
	i := p.IndexOf(v)
	if i >= 0 {
		p.Delete(i)
		p.shrink()
	}
	return i
}

func (p *List) grow() {
	var delta int
	if p.capacity > 64 {
		delta = p.capacity / 4
	} else {
		if p.capacity > 8 {
			delta = 16
		} else {
			delta = 4
		}
	}
	p.Grow(p.capacity + delta)
}

// Grow Used to set the size (number object pointers) of the list
func (p *List) Grow(n int) {
	newSlice := make([]interface{}, n)
	copy(newSlice, p.elements)
	p.elements = newSlice
	p.capacity = n
}
func (p *List) shrink() {
	if p.size <= int(float32(len(p.elements))*0.2) {
		p.Grow(p.size)
	}
}
func (p *List) String() string {
	var b bytes.Buffer
	if p.size == 0 {
		return ""
	}
	b.WriteString(fmt.Sprintf("%v", p.elements[0]))

	for i := 1; i < p.size; i++ {
		b.WriteString(fmt.Sprintf(",%v", p.elements[i]))
	}
	return b.String()
}
