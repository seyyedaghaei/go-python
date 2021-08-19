package python

// #include <Python.h>
import "C"
import (
	"reflect"
)

//togo converts a CPyObject to a *PyObject
func togo(cobject CPyObject) *Object {
	return (*Object)(cobject)
}

func toObj(obj interface{}) *Object {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Array:
		size := value.Len()
		tuple := NewTuple(size)
		for i := 0; i < size; i++ {
			tuple.SetItem(i, value.Index(i).Interface())
		}
		return tuple.Object()
	case reflect.Slice:
		size := value.Len()
		list := NewList(size)
		for i := 0; i < size; i++ {
			list.SetItem(i, value.Index(i).Interface())
		}
		return list.Object()
	case reflect.Map:
		dict := NewDict()
		keys := value.MapKeys()
		for _, key := range keys {
			dict.SetItem(key.Interface(), value.MapIndex(key).Interface())
		}
		return dict.Object()
	}
	switch o := obj.(type) {
	case PyObject:
		return o.Object()
	case int:
		return PyInt(o).Object()
	case string:
		return PyString(o).Object()
	case float64:
		return PyFloat(o).Object()
	}
	return nil
}

func toC(obj interface{}) CPyObject {
	return toObj(obj).C()
}
