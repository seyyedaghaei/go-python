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

func (dict *Dict) SetItem(key string, value interface{}) {
	C.PyDict_SetItemString(dict.C(), C.CString(key), toC(value))
}
