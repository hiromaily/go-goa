init:
	npm init -y
	npm i babel-core babel-loader babel-preset-es2015 webpack -D
	npm i babel-preset-es2015-riot riotjs-loader riot-tag-loader webpack-dev-server -D
	npm i babel-polyfill -D
	npm i riot riot-route -S

	mkdir -p src/tag public
	touch .babelrc webpack.config.js


	cat <<_END_ > .babelrc
{
  "presets": [
	"es2015-riot"
  ]
}
_END_

	cat <<_END_ > webpack.config.js
const path = require('path');

module.exports = {
  entry: './src/app.js',
  output: {
	 path: __dirname + '/src/js/',
	 filename: 'bundle.js'
  },
  module: {
	 rules: [
		{
		   test: /\.tag$/,
		   enforce: 'pre',
		   exclude: /node_modules/,
		   use: [
			  {
				 loader: 'riot-tag-loader',
				 options: {
					template: 'pug',   //template if needed
					debug: true
				 }
			  }
		   ]
		},
		{
		   test: /\.js|\.tag$/,
		   enforce: 'post',
		   exclude: /node_modules/,
		   use: [
			  {
				 loader: 'babel-loader',
				 options: {
					presets: `es2015-riot`
				 }
			  }
		   ]
		}
	 ]
  },
  resolve: {
	 extensions: ['.js', '.tag']
  },
  plugins: [
	 new webpack.ProvidePlugin({ riot: 'riot' })
  ]
};
_END_


	cat <<_END_ > public/index.html
<!DOCTYPE html>
<html>
  <head>
	<meta charset="utf-8">
	<title>sample</title>
  </head>
  <body>
	<app></app>
	<script src="src/js/bundle.js"></script>
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
