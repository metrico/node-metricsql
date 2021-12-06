const metricsql = require('./');
process.argv.splice(0, 2);
var query =  process.argv[0] || `sum(rate(foo{bar="baz"}[5m])) by (x,y)`
try {
  const parsed = metricsql.parse(query);
  console.log(parsed);
} catch(e){
  console.log(e); 
}
