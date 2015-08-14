GO_BUILDOPT := -ldflags '-s -w'

gom:
	go get github.com/mattn/gom
	gom install

run:
	gom run main.go ${ARGS}

fmt:
	gom exec goimports -w *.go

hal: fmt
	gom build $(GO_BUILDOPT) -o bin/hal hal.go

victor: fmt
	gom build $(GO_BUILDOPT) -o bin/victor victor.go

clean:
	rm -f bin/*
