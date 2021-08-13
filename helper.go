package python

// #include <Python.h>
//PyObject * BuildValue(const char * name, int num) {
//    return Py_BuildValue(name, num);
//}
import "C"

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *Object {
	return (*Object)(cobject)
}

func toc(object *Object) *C.PyObject {
	return (*C.PyObject)(object)
}

func String(str string) *Object {
	return togo(C.PyUnicode_FromString(C.CString(str)))
}

func Int(num int) *Object {
	return togo(C.BuildValue(C.CString("i"), C.int(num)))
}
