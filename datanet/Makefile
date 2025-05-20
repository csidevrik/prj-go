.PHONY: run build tidy clean

run:
	go run ./cmd/app

build:
	go build -o datanet ./cmd/app

tidy:
	go mod tidy

clean:
	del datanet.exe 2>nul || true
