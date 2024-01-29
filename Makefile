build:
	go build -o ./bin/blog-agg ./cmd/blog-agg

run:
	./bin/blog-agg

all: build run