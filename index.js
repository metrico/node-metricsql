const metricsql = require('./src');
process.argv.splice(0, 2);
var query =  process.argv[0] || `sum(rate(foo{bar="baz"}[5m])) by (x,y)`
try {
  const parsed = JSON.parse(metricsql.parse(query));
  console.log(parsed);
  return parsed;
} catch(e){
  console.log(e); 
}
