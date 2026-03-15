NAME=libship

build:
	@mkdir -p bin
	go build -o bin/$(NAME) ./cmd/$(NAME)/main.go

run: build
	@./bin/$(NAME)

clean:
	@rm -rf bin
