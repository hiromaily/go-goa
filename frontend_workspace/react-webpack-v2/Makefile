init:
	npm init -y
	npm i babel-core babel-loader babel-preset-es2015 webpack -D
	npm i babel-preset-react webpack-dev-server -D
	#npm i browserify browserify-shim -D
	npm i react react-dom -S
	npm i babel-polyfill -D

	mkdir src public
	touch .babelrc webpack.config.js


	cat <<_END_ > .babelrc
{
  "presets": [
	"es2015", "react"
  ]
}_END_


	cat <<_END_ > webpack.config.js
const path = require('path');

module.exports = {
  entry: path.join(__dirname, 'src/app.js'),
  output: {
	path: path.join(__dirname, 'public'),
	filename: 'bundle.js'
  },
  devServer: {
	contentBase: path.join(__dirname, 'public'),
	port: 3000,
	inline: true
  },
  module: {
	rules: [
	  {
		test: /\.js$/,
		use: 'babel-loader',
		exclude: /node_modules/
	  }
	]
  }
};
_END_


	cat <<_END_ > src/app.js
import React from 'react';
import {render} from 'react-dom';

class App extends React.Component {
  render () {
    return <p>Hello React!</p>;
  }
}

render(<App/>, document.getElementById('app'));
_END_


	cat <<_END_ > public/index.html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>React</title>
  </head>
  <body>
    <div id="app"></div>
    <script src="bundle.js"></script>
  </body>
</html>
_END_


###############################################################################
# Operation
###############################################################################
build:
	./node_modules/.bin/webpack
	#npm run build

server:
	./node_modules/.bin/webpack-dev-server
	#npm start

open:
    http://localhost:3000/

production:
    npm run prod

###############################################################################
# Riot
###############################################################################
riot:
    npm install riot

###############################################################################
# npm
###############################################################################
install:
	npm install

outdated:
	npm outdated

update:
	npm update

updatepkg:
	npm update --save babel-loader
	npm update --save browserify
	npm update --save flux
	npm update --save gulp-uglify
	npm update --save history
	npm update --save react-router
