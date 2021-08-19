package python

// #include<Python.h>
import "C"

type Tuple Object

func NewTuple(size int) *Tuple {
	return (*Tuple)(togo(C.PyTuple_New(C.ssize_t(size))))
}

func (tuple *Tuple) C() CPyObject {
	return (CPyObject)(tuple)
}

func (tuple *Tuple) Object() *Object {
	return (*Object)(tuple)
}

func (tuple *Tuple) SetItem(index int, item interface{}) int {
	return int(C.PyTuple_SetItem(tuple.C(), C.ssize_t(index), toC(item)))
}

func (tuple *Tuple) GetItem(index int) *Object {
	return togo(C.PyTuple_GetItem(tuple.C(), C.ssize_t(index)))
}

func (tuple *Tuple) GetSlice(start int, stop int) *Object {
	return togo(C.PyTuple_GetSlice(tuple.C(), C.ssize_t(start), C.ssize_t(stop)))
}

func (tuple *Tuple) Size(index int, item interface{}) int {
	return int(C.PyTuple_Size(tuple.C()))
}
