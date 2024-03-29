default: build

compile-go:
	go build -buildmode=c-shared -o metricsql.so src/golang/metricsql.go
	strip metricsql.so

build: compile-go
	npm run build
