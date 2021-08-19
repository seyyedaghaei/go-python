package python

// #include <Python.h>
//PyObject * BuildValue(const char * name, int num) {
//    return Py_BuildValue(name, num);
//}
import "C"

type Int C.PyObject

func (i *Int) C() CPyObject {
	return (CPyObject)(i)
}

func (i *Int) Object() *Object {
	return (*Object)(i)
}

func (i *Int) Int() int {
	return int(C.PyFloat_AsDouble(i.C()))
}

func PyInt(num int) *Int {
	return (*Int)(C.BuildValue(C.CString("i"), C.int(num)))
}
