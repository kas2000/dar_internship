package deque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushFront(t *testing.T) {
	val := 1
	d := deque{}
	d.PushFront(val)
	d.PushRear(5)
	d.PushRear(4)
	d.PushRear(3)
	d.PushRear(2)
	items := d.ItemsCopy()
	assert.Equal(t, val, items[0])
}

func TestPushRear(t *testing.T) {
	val := 2
	d := deque{}
	d.PushFront(val)
	d.PushRear(5)
	d.PushRear(4)
	d.PushRear(3)
	expected := 8
	d.PushRear(expected)
	items := d.ItemsCopy()
	assert.Equal(t, expected, items[len(items)-1])
}

func TestPopFront(t *testing.T) {
	d := deque{}
	d.PushFront(2)
	d.PushRear(5)
	d.PushRear(4)
	expected := 11
	d.PushFront(expected)
	d.PushRear(3)
	d.PushRear(8)
	assert.Equal(t, expected, d.PopFront())
}

func TestPopRear(t *testing.T) {
	d := deque{}
	d.PushFront(2)
	d.PushRear(5)
	d.PushRear(4)
	expected := 8
	d.PushFront(11)
	d.PushRear(3)
	d.PushRear(expected)
	assert.Equal(t, expected, d.PopRear())
}