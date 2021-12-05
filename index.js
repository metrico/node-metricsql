const metricsql = require('./src');

console.log(metricsql);
console.log(
  metricsql.parse(`sum(rate(foo{bar="baz"}[5m])) by (x,y)`)
);
