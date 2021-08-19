package python

// #include <Python.h>
import "C"

type Float C.PyObject

func (float *Float) Object() *Object {
	return (*Object)(float)
}

func (float *Float) Float64() float64 {
	return float64(C.PyFloat_AsDouble(toc((*Object)(float))))
}

func PyFloat(num float64) *Float {
	return (*Float)(C.PyFloat_FromDouble(C.double(num)))
}
