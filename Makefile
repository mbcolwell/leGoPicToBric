MAIN = picToBric
BIN = bin/
CMD = cmd/
ARGS = 


# all: build run
all: build

build:
	go build -o $(BIN)$(MAIN) $(CMD)$(MAIN).go

run:
	$(PIPE) ./$(BIN)$(MAIN) $(ARGS)

clean:
	rm -f $(BIN)$(MAIN) output.log
