package main

import (
	"C"
//	"fmt"
	"log"
        "encoding/json"
	"github.com/VictoriaMetrics/metricsql"
)

//export Parse
func Parse(input *C.char) *C.char {
	inputString := C.GoString(input)
	expr, err := metricsql.Parse(inputString)
        if err != nil {
		log.Fatalf("Parsing error: %s", err)
	}

	/* debug output */
	/*
	ae := expr.(*metricsql.AggrFuncExpr)
	fmt.Printf("aggr func: name=%s, arg=%s, modifier=%s\n", ae.Name, ae.Args[0].AppendString(nil), ae.Modifier.AppendString(nil))

	fe := ae.Args[0].(*metricsql.FuncExpr)
	fmt.Printf("func: name=%s, arg=%s\n", fe.Name, fe.Args[0].AppendString(nil))

	re := fe.Args[0].(*metricsql.RollupExpr)
	fmt.Printf("rollup: expr=%s, window=%s\n", re.Expr.AppendString(nil), re.Window.AppendString(nil))

	me := re.Expr.(*metricsql.MetricExpr)
	fmt.Printf("metric: labelFilter1=%s, labelFilter2=%s", me.LabelFilters[0].AppendString(nil), me.LabelFilters[1].AppendString(nil))
	*/

	j, err := json.Marshal(expr)
    	if err != nil {
    	    log.Fatalf("JSON error: %s", err.Error())
    	}
	resString := string(j)
	return C.CString(resString)
}

func main() {}
