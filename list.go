package python

// #include<Python.h>
import "C"

type List Object

func NewList(size int) *List {
	return (*List)(togo(C.PyList_New(C.ssize_t(size))))
}

func (list *List) C() CPyObject {
	return (CPyObject)(list)
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

func (list *List) GetSlice(start int, stop int) *Object {
	return togo(C.PyList_GetSlice(list.C(), C.ssize_t(start), C.ssize_t(stop)))
}

func (list *List) Append(item interface{}) int {
	return int(C.PyList_Append(list.C(), toC(item)))
}

func (list *List) AsTuple(item interface{}) *Tuple {
	return (*Tuple)(C.PyList_AsTuple(list.C()))
}
