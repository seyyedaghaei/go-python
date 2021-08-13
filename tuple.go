package pyhton

// #include<Python.h>
import "C"

type Tuple Object

func NewTuple(size int) *Tuple {
	return (*Tuple)(togo(C.PyTuple_New(C.ssize_t(size))))
}

func (tuple *Tuple) SetItem(index int, item *Object) int {
	return int(C.PyTuple_SetItem(toc((*Object)(tuple)), C.ssize_t(index), toc(item)))
}
