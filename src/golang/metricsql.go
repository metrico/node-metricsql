package main

import (
	"C"
	//	"fmt"
	"encoding/json"
	"log"

	metricsql "github.com/VictoriaMetrics/metricsql"
)
import (
	"strings"
)

type RollupExprCopy struct {
	Name string
	Args []struct {
		Name string
		Args []struct {
			Expr        map[string]interface{}
			Window      string
			Offset      string
			Step        string
			InheritStep bool
		}
	}
	Modifier map[string]interface{}
}

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
		fmt.Printf("metric: labelFilter1=%s, labelFilter2=%s\n", me.LabelFilters[0].AppendString(nil), me.LabelFilters[1].AppendString(nil))
	*/

	j, err := json.Marshal(expr)
	if err != nil {
		log.Fatalf("JSON error: %s", err.Error())
	}

	//Small bug - Windows is object {} - need replace it to "string"
	//or unmarshal will show an error
	j = []byte(strings.ReplaceAll(string(j), "\"Window\":{}", "\"Window\":\"\""))

	rollExp := RollupExprCopy{}
	/* we do unmarshal to our structure */
	err = json.Unmarshal(j, &rollExp)
	if err != nil {
		log.Fatalf("JSON unmarshal error: %s\n", err.Error())
	}

	// here we set the window
	if ae, ok := expr.(*metricsql.AggrFuncExpr); ok {
		for mIndex, _ := range rollExp.Args {
			if fe, ok := ae.Args[mIndex].(*metricsql.FuncExpr); ok {
				for subIndex, _ := range rollExp.Args[mIndex].Args {
					if re, ok := fe.Args[subIndex].(*metricsql.RollupExpr); ok {
						rollExp.Args[mIndex].Args[subIndex].Window = string(re.Window.AppendString(nil))
					}
				}
			}
		}
	}

	j, err = json.Marshal(rollExp)
	if err != nil {
		log.Fatalf("JSON error2: %s\n", err.Error())
	}
	return C.CString(string(j))
}

func main() {}
