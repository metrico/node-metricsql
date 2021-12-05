const metricsql = require('./src');

process.argv.splice(0, 2);
var query =  process.argv[0] || `sum(rate(foo{bar="baz"}[5m])) by (x,y)`

console.log(
  metricsql.parse(query)
);
