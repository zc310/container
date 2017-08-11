package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var list List
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	assert.Equal(t, "0,1,2,3,4,5,6,7,8,9", list.String())

	list.Swap(0, 9)

	assert.Equal(t, list.Count(), 10)
	v, _ := list.First()
	assert.Equal(t, v.(int), 9)
	v, _ = list.Last()
	assert.Equal(t, v.(int), 0)
	list.Swap(0, 9)

	list.Delete(0)
	assert.Equal(t, list.Count(), 9)
	v, _ = list.First()
	assert.Equal(t, v.(int), 1)
	v, _ = list.Get(2)
	assert.Equal(t, v.(int), 3)
	assert.Equal(t, list.IndexOf(7), 6)
	ok := list.Put(0, 10)
	assert.Equal(t, ok, true)
	v, _ = list.Get(0)
	assert.Equal(t, 10, v)
	list.Remove(1)
	assert.Equal(t, list.Count(), 9)
	assert.Equal(t, list.IndexOf(9), 8)
	assert.Equal(t, list.Remove(9), 8)
	assert.Equal(t, list.IndexOf(9), -1)

	list.Clear()
	assert.Equal(t, list.Empty(), true)

}
func BenchmarkAdd(b *testing.B) {
	var list List
	for i := 0; i < 4782969; i++ {
		list.Add(i)
	}
}
func BenchmarkSetCapacity(b *testing.B) {
	var list List
	list.SetCapacity(4782969)
	for i := 0; i < 4782969; i++ {
		list.Add(i)
	}
}
