package deque

import (
	"fmt"
	"strings"
)

type deque struct {
	items []interface{}
}

func (d *deque) PushFront(e interface{}) {
	d.items = append([]interface{}{e}, d.items...)
}

func (d *deque) PushRear(e interface{}) {
	d.items = append(d.items, e)
}

func (d *deque) PopFront() interface{} {
	x := d.items[0]
	d.items = d.items[1:]
	return x
}

func (d *deque) PopRear() interface{} {
	x := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return x
}

func (d deque) IsEmpty() bool{
	if len(d.items) == 0{
		return true
	}
	return false
}
//TODO: peek function
func (d deque) Peek(direction string) {
	if (len(d.items) > 0) {
		if (strings.ToUpper(direction) == "START") {
			fmt.Println(d.items[0])
		} else if (strings.ToUpper(direction) == "END") {
			fmt.Println(d.items[len(d.items)-1])
		}
	} else {
		fmt.Println(d.items)
	}
}

func (d deque) ItemsCopy() []interface{} {
	itemsCopy := make([]interface{}, len(d.items))
	copy(itemsCopy, d.items)
	return itemsCopy
}