RULES_DIR = ../common/rules
SCRIPTS_DIR = ../common/scripts

SCRATCH_DIR=.scratch

DIRS=$(SCRATCH_DIR)

all: run

clean:
	rm -rfv $(DIRS)

run: $(SCRATCH_DIR)
	bash ./run.sh > run.txt

$(DIRS):
	mkdir -p $@

