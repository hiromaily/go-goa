#------------------------------------------------------------------------------
# Initial Setting
#------------------------------------------------------------------------------
.PHONY: setup
setup:
	./scripts/setup.sh

.PHONY: setup-msw
setup-msw:
	# setup worker
	npx msw init ./public --save
	mkdir src/mocks
	touch src/mocks/handlers.ts
	touch src/mocks/browser.ts
	touch src/mocks/server.ts
	touch src/mocks/worker.ts


#------------------------------------------------------------------------------
# Maintenance
#------------------------------------------------------------------------------
.PHONY: check
check:
	npx npm-check-updates

.PHONY: update
update:
	npm update


#------------------------------------------------------------------------------
# Development
#------------------------------------------------------------------------------
.PHONY: dev
dev:
	npm run dev

.PHONY: lint
lint:
	npm run lint-all

.PHONY: test
test:
	npm run test


#------------------------------------------------------------------------------
# Deploy
#------------------------------------------------------------------------------
# This output is generated inside the .next folder
.PHONY: build
build:
	npm run build

# .PHONY: deploy
# deploy:
# 	npm run export

.PHONY: publish
publish:
	ngrok http 3000