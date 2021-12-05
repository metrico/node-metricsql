package main

import (
	"C"
	"fmt"
	"log"
	"github.com/VictoriaMetrics/metricsql"
)

//export Parse
func Parse(input *C.char) *C.char {
	inputString := C.GoString(input)
	expr, err := metricsql.Parse(inputString)
        if err != nil {
		log.Fatalf("parse error: %s", err)
	}
	res := fmt.Sprintf("parsed expr: %s\n", expr.AppendString(nil))
	resString := string(res)
	return C.CString(resString)
}

func main() {}
