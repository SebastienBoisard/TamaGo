BINARY_NAME=TamaGo
TERMINAL="terminator -e"

all: build

build:
	go build -o $(BINARY_NAME) -v main/tamago.go

run:
	$(TERMINAL) './$(BINARY_NAME) back'
	$(TERMINAL) './$(BINARY_NAME) front'
