var path = require('path')
var config = {
  entry: path.resolve(__dirname, 'src/js/index.js'),
  output: {
    path: path.resolve(__dirname, 'build/'),
    filename: 'app.js'
  },
  module: {
    loaders: [{
      test: /\.js/,
      exclude: /(node_modules)/,
      loader: 'babel-loader',
      query: {
        cacheDirectory: true,
        plugins: ["transform-runtime"]
      }
    }]
  }

};

module.exports = config;
