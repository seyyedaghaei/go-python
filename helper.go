package python

// #include <Python.h>
import "C"

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *Object {
	return (*Object)(cobject)
}

func toc(object *Object) *C.PyObject {
	return (*C.PyObject)(object)
}

func PyString(str string) *Object {
	return togo(C.PyUnicode_FromString(C.CString(str)))
}
