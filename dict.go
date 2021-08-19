package python

// #include <Python.h>
import "C"

type Dict C.PyObject

func NewDict() *Dict {
	return (*Dict)(togo(C.PyDict_New()))
}

func (dict *Dict) C() CPyObject {
	return (CPyObject)(dict)
}

func (dict *Dict) Object() *Object {
	return (*Object)(dict)
}

func (dict *Dict) SetItem(key interface{}, value interface{}) {
	C.PyDict_SetItem(dict.C(), toC(key), toC(value))
}

func (dict *Dict) GetItem(key interface{}) *Object {
	return (*Object)(C.PyDict_GetItem(dict.C(), toC(key)))
}
