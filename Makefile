build:
	go build -gcflags="all=-N -l" -o ./bin/blog-agg ./cmd/blog-agg 

run:
	./bin/blog-agg

all: build run