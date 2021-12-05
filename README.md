<img src="https://user-images.githubusercontent.com/1423657/144750632-c129d650-a898-4436-a65a-a5d4519c42d1.png" width=100 />

# node-metricsql

> Native node binding for the [MetricsQL library](https://github.com/VictoriaMetrics/metricsql)


### Build Module
```
make
```

### Test Module
```
node index.js 'sum(rate(foo{bar="baz"}[5m]))'
```
```
aggr func: name=sum, arg=rate(foo{bar="baz"}[5m]), modifier=by (x, y)
func: name=rate, arg=foo{bar="baz"}[5m]
rollup: expr=foo{bar="baz"}, window=5m
metric: labelFilter1=__name__="foo", labelFilter2=bar="baz"parsed expr: sum(rate(foo{bar="baz"}[5m])) by (x, y)
```

#### Todo
- [x] go binding
- [ ] function mapping
- [ ] format conversion

