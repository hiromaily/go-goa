init:
	npm init -y
	npm i babel-core babel-loader babel-preset-es2015 webpack -D
	npm i babel-preset-es2015-riot riotjs-loader riot-tag-loader webpack-dev-server -D
	npm i babel-polyfill -D
	npm i riot riot-route -S

	mkdir -p src/tag public
	touch .babelrc webpack.config.js

initlink:
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/assets ./public/assets
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/components ./public/components
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/css ./public/css
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/img ./public/img
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/js ./public/js
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/json ./public/json
	ln -s ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/themes ./public/themes

###############################################################################
# Operation
###############################################################################
build:
	npm run build

server:
	npm run start

start: build server

open:
	http://localhost:3000/

deployjs:
	cp public/resume.bundle.js ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/js/
	cp public/admin.bundle.js ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/js/


###############################################################################
# Riot
###############################################################################
riot:
	npm install riot

