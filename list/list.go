package list

import (
	"bytes"
	"fmt"
)

// List very useful general purpose list container
type List struct {
	elements []interface{}
}

// New returns an initialized list.
func New() *List {
	return new(List)
}

// Count The number of items in the list
func (p List) Count() int {
	return len(p.elements)
}

// Empty If an empty list
func (p List) Empty() bool {
	return p.Count() == 0
}

// Add Add an item to the list
func (p *List) Add(v interface{}) int {
	p.elements = append(p.elements, v)
	return len(p.elements) - 1
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
}

// Delete Removes an item from the list by its list position
func (p *List) Delete(i int) {
	if !p.inRange(i) {
		return
	}
	p.elements = append(p.elements[:i], p.elements[i+1:]...)
}

// Last Gets the last item in the list
func (p *List) Last() (interface{}, bool) {
	return p.Get(p.Count() - 1)
}

// IndexOf Gives the list position of a specified object in the list
func (p *List) IndexOf(v interface{}) int {
	var i int
	for i < p.Count() {
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
	return index >= 0 && index < p.Count()
}

// Remove Removes an item from the list by its object
func (p *List) Remove(v interface{}) int {
	i := p.IndexOf(v)
	if i >= 0 {
		p.Delete(i)
	}
	return i
}

// Grow Used to set the size (number object pointers) of the list
func (p *List) Grow(n int) {
	newSlice := make([]interface{}, len(p.elements), n)
	copy(newSlice, p.elements)
	p.elements = newSlice
}

func (p *List) String() string {
	var b bytes.Buffer
	if p.Empty() {
		return ""
	}
	b.WriteString(fmt.Sprintf("%v", p.elements[0]))

	for i := 1; i < p.Count(); i++ {
		b.WriteString(fmt.Sprintf(",%v", p.elements[i]))
	}
	return b.String()
}
