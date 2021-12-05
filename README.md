<img src="https://user-images.githubusercontent.com/1423657/144750632-c129d650-a898-4436-a65a-a5d4519c42d1.png" width=100 />

# node-metricsql

> Native node binding for the [MetricsQL library](https://github.com/VictoriaMetrics/metricsql)


### Build Module
```console
make
```

### Test Module
```console
node index.js 'sum(rate(foo{bar="baz"}[5m]))'
```
```json
{"Name":"sum","Args":[{"Name":"rate","Args":[{"Expr":{"LabelFilters":[{"Label":"__name__","Value":"foo","IsNegative":false,"IsRegexp":false},{"Label":"bar","Value":"baz","IsNegative":false,"IsRegexp":false}]},"Window":{},"Offset":null,"Step":null,"InheritStep":false}]}],"Modifier":{"Op":"by","Args":["x","y"]},"Limit":0}
```

#### Todo
- [x] go binding
- [x] function mapping
- [x] format conversion

