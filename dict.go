package pyhton

// #include <Python.h>
import "C"

type Dict C.PyObject

func NewDict() *Dict {
	return (*Dict)(togo(C.PyDict_New()))
}

func (dict *Dict) SetItem(key string, value interface{}) {
	C.PyDict_SetItemString(toc((*Object)(dict)), C.CString(key), toc(toObj(value)))
}