.PHONY: install
install:
	cd ./idp && go mod tidy
	cd ./client && npm install

.PHONY: run
run:
	cd ./idp && go run ./main.go & cd ./client && npm run dev
