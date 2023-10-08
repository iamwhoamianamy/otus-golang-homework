package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("one item, push front", func(t *testing.T) {
		l := NewList()

		expected := 10
		l.PushFront(expected)

		require.Equal(t, expected, l.Front().Value.(int))
		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Front(), l.Back())
	})

	t.Run("one item, push back", func(t *testing.T) {
		l := NewList()

		expected := 10
		l.PushBack(expected)

		require.Equal(t, expected, l.Front().Value.(int))
		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Front(), l.Back())
	})

	t.Run("remove single item", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		toRemove := l.Front()
		l.Remove(toRemove)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
		require.Nil(t, toRemove.Next)
		require.Nil(t, toRemove.Prev)
	})

	t.Run("remove one of two items", func(t *testing.T) {
		l := NewList()

		l.PushBack(1)
		l.PushBack(2)

		require.Equal(t, 2, l.Len())

		toRemove := l.Front()
		l.Remove(toRemove)

		require.Equal(t, 1, l.Len())
		require.Nil(t, toRemove.Next)
		require.Nil(t, toRemove.Prev)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
