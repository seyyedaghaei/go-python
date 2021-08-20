package python

// #include<Python.h>
import "C"

type List Object

func NewList(size int) *List {
	return (*List)(C.PyList_New(C.ssize_t(size)))
}

func (list *List) C() CPyObject {
	return list.Object().C()
}

func (list *List) Object() *Object {
	return (*Object)(list)
}

func (list *List) SetItem(index int, item interface{}) int {
	return int(C.PyList_SetItem(list.C(), C.ssize_t(index), toC(item)))
}

func (list *List) GetItem(index int) *Object {
	return togo(C.PyList_GetItem(list.C(), C.ssize_t(index)))
}

func (list *List) GetSlice(start int, stop int) *List {
	return (*List)(C.PyList_GetSlice(list.C(), C.ssize_t(start), C.ssize_t(stop)))
}

func (list *List) Insert(index int, item interface{}) int {
	return int(C.PyList_Insert(list.C(), C.ssize_t(index), toC(item)))
}

func (list *List) Reverse() int {
	return int(C.PyList_Reverse(list.C()))
}

func (list *List) Sort() int {
	return int(C.PyList_Sort(list.C()))
}

func (list *List) Append(item interface{}) int {
	return int(C.PyList_Append(list.C(), toC(item)))
}

func (list *List) AsTuple() *Tuple {
	return (*Tuple)(C.PyList_AsTuple(list.C()))
}

func (list *List) Size() int {
	return int(C.PyList_Size(list.C()))
}

func (list *List) String() string {
	return list.Object().String()
}

func (list *List) Interface() interface{} {
	return list.Object().Interface()
}

func (list *List) Slice() []interface{} {
	size := list.Size()
	slice := make([]interface{}, size)
	for i := 0; i < size; i++ {
		slice[i] = list.GetItem(i).Interface()
	}
	return slice
}
