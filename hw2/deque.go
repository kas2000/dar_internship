package main

import _"fmt"

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
	x := d.items[0]
	d.items = d.items[:len(d.items)-1]
	return x
}

func (d *deque) IsEmpty() bool{
	if len(d.items) == 0{
		return true
	}
	return false
}
//TODO: peek function
func (d deque) Peek() interface{} {
	return 1
}

func main() {

}