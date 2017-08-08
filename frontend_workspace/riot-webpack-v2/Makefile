init:
	npm init -y
	npm i babel-core babel-loader babel-preset-es2015 webpack -D
	npm i babel-preset-es2015-riot riotjs-loader riot-tag-loader webpack-dev-server -D
	npm i babel-polyfill -D
	npm i riot riot-route -S

	mkdir -p src/tag public
	touch .babelrc webpack.config.js


###############################################################################
# Operation
###############################################################################
build:
	npm run build

server:
	npm run start

open:
	http://localhost:3000/

deployjs:
	cp public/js/bundle.js ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/js/resume.js
	cp public/json/liketech.json ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/json/
	cp public/json/disliketech.json ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/json/
	cp public/json/workhistory.json ${HOME}/work/go/src/github.com/hiromaily/go-goa/public/json/


###############################################################################
# Riot
###############################################################################
riot:
	npm install riot

