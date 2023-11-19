MAIN = picToBric
BIN = bin/
CMD = cmd/
PIPE = 
ARGS = internal/test.jpg


all: build run

build:
	go build -o $(BIN)$(MAIN) $(CMD)$(MAIN).go

run:
	$(PIPE) ./$(BIN)$(MAIN) $(ARGS)

clean:
	rm -f $(BIN)$(MAIN) output.log
