all: render

render: render.go
	go build -o $@
