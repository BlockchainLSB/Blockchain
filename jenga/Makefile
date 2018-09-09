CC=go
ROOT=$(shell pwd)
CMDDIR=cmd/pms
CMDNAME=pms
CONF=etc/config.yml
BIN=bin

all: clean pms run

pms:
	cd $(CMDDIR) && go build -o $(CMDNAME)
	cd $(ROOT)
	rm -rf $(BIN)
	mkdir $(BIN)
	cp -rf $(CMDDIR)/$(CMDNAME) public $(CONF) $(BIN)/
	rm -rf $(CMDDIR)/$@

clean:
	rm -rf $(BIN)

run:
	cd $(BIN) && ./$(CMDNAME)
