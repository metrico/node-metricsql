try {
  var binding = require('node-gyp-build')(__dirname)
  module.exports = binding
} catch (err) {
  module.exports = require('./build/Release/addon.node');
}
