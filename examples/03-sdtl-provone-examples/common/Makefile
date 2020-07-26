RULES_DIR = ../common/rules
SCRIPTS_DIR = ../common/scripts

RESULT_DIR=results

DIRS=$(RESULT_DIR)

all: query

clean:
	rm -rfv $(DIRS)

query: $(RESULT_DIR)
	bash ./run.sh > results/run.txt

$(DIRS):
	mkdir -p $@

