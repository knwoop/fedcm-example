.PHONY: run
run:
	cd ./idp && go run ./main.go & cd ./client && npm run dev
