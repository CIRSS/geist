RULES_DIR = ../common/rules
SCRIPTS_DIR = ../common/scripts

RUN_SCRIPT_FILE=./run.sh
RUN_OUTPUT_FILE=./run.txt
RUN_SCRATCH_DIR=.scratch

DIRS=$(RUN_SCRATCH_DIR)

all: run

clean:
	rm -fv ${RUN_OUTPUT_FILE}
	rm -rfv $(RUN_SCRATCH_DIR)

run: $(RUN_SCRATCH_DIR)
	bash ${RUN_SCRIPT_FILE} > ${RUN_OUTPUT_FILE}

$(DIRS):
	mkdir -p $@

