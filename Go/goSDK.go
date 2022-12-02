package main

import "C"

//export GoCrossCompileExample
func GoCrossCompileExample(stringValue *C.char, intValue C.int, resultSize *C.int, error *C.char) *C.char {

	//Converting c string to go string
	goString := C.GoString(stringValue)

	//Converting
	goInt := C.GoInt(intValue)
	println("Received String: ", goString)
	println("Received Int: ", goInt)
	newGoString := "Go" + goString
	errorString := "go error"
	*error = C.CString(errorString)
	*resultSize = C.int(len(newGoString))

	result := C.CString(newGoString)
	return result
}

func main() {}