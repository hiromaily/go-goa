require('babel-polyfill');
const path = require('path');
const webpack = require("webpack");

const DEBUG = !process.argv.includes('--env.production');

const plugins = [
    new webpack.optimize.OccurrenceOrderPlugin()
];

if(!DEBUG){
    plugins.push(
        new webpack.optimize.UglifyJsPlugin({ compress: { screw_ie8: false, warnings: false } }),
        new webpack.optimize.AggressiveMergingPlugin()
    );
}

module.exports = {
  entry: path.join(__dirname, 'src/app.js'),
  output: {
    path: path.join(__dirname, 'public/js'),
    filename: 'bundle.js'
  },
  externals: {
    "react": "React",
    "react-dom": "ReactDOM",
    "jquery": "jQuery"
  },
  devServer: {
    contentBase: path.join(__dirname, 'public'),
    port: 3000,
    inline: true
  },
  plugins: plugins,
  devtool: DEBUG ? 'cheap-module-eval-source-map' : false,
  module: {
    rules: [
      {
        test: /\.jsx$/,
        use: 'babel-loader',
        exclude: /node_modules/
      }
    ]
  }
};
