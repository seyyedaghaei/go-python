package pyhton

// #cgo pkg-config: python3
// #cgo CFLAGS : -I./ -I/usr/include/python3.9
// #cgo LDFLAGS: -lpython3.9 -lcrypt -lpthread -ldl  -lutil -lm -lm
//#include <Python.h>
import "C"
import "os"

/* These definitions must match corresponding definitions in graminit.h. */
const (
	singleInput   = 256
	fileInput     = 257
	evalInput     = 258
	funcTypeInput = 345
	/* This doesn't need to match anything */
	fstringInput = 800
)

func Initialize() {
	C.Py_Initialize()
	dir, _ := os.Getwd()
	C.PyRun_SimpleString(C.CString("import sys\nsys.path.append(\"" + dir + "\")"))
}

func Finalize() {
	C.Py_Finalize()
}

func PrintError() {
	C.PyErr_Print()
}

func ImportModule(module string) *Object {
	return togo(C.PyImport_ImportModule(C.CString(module)))
}

type PyGILState_STATE C.PyGILState_STATE

func GILStateEnsure() PyGILState_STATE {
	return PyGILState_STATE(C.PyGILState_Ensure())
}

func GILStateRelease(state PyGILState_STATE) {
	C.PyGILState_Release(C.PyGILState_STATE(state))
}

func Run(code string, globals *Dict, locals *Dict) *Object {
	return togo(C.PyRun_StringFlags(C.CString(code), C.int(fileInput), toc((*Object)(globals)), toc((*Object)(locals)), nil))
}
