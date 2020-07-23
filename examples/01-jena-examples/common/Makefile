RULES_DIR = ../common/rules
SCRIPTS_DIR = ../common/scripts

QUERIES_DIR=queries

DIRS=$(QUERIES_DIR)

all: query

clean:
	rm -rfv $(DIRS)

query: $(QUERIES_DIR)
	bash ./run_queries.sh > queries/run_queries.txt

$(DIRS):
	mkdir -p $@

