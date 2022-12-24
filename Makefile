.PHONY: install
install:
	cd ./idp && go mod tidy
	cd ./rp && npm install

.PHONY: run
run:
	cd ./idp && go run ./main.go & cd ./rp && npm run dev
