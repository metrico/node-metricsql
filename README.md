<img src="https://user-images.githubusercontent.com/1423657/144750632-c129d650-a898-4436-a65a-a5d4519c42d1.png" width=100 />

# node-metricsql

> Native node binding for [MetricsQL Parser](https://github.com/VictoriaMetrics/metricsql)


### Build Module
```console
make
```

### Usage
```javascript
const metricsql = require('node-metricsql');
const parsed = metricsql.parse(promql);
```

### Test Module
```console
node example.js 'sum(rate(foo{bar="baz"}[5m]))'
```
```json
{
   "Name":"sum",
   "Args":[
      {
         "Name":"rate",
         "Args":[
            {
               "Expr":{
                  "LabelFilters":[
                     {
                        "IsNegative":false,
                        "IsRegexp":false,
                        "Label":"__name__",
                        "Value":"foo"
                     },
                     {
                        "IsNegative":false,
                        "IsRegexp":false,
                        "Label":"bar",
                        "Value":"baz"
                     }
                  ]
               },
               "Window":"5m",
               "Offset":"",
               "Step":"",
               "InheritStep":false
            }
         ]
      }
   ],
   "Modifier":{
      "Args":[
         "x",
         "y"
      ],
      "Op":"by"
   }
}
```

#### Todo
- [x] go binding
- [x] function mapping
- [x] format conversion

