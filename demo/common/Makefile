RULES_DIR = ../common/rules
SCRIPTS_DIR = ../common/scripts

RUN_SCRIPT_FILE=./run.sh
RUN_OUTPUT_FILE=./run.txt
RUN_OUTPUTS_DIR=./outputs
RUN_SCRATCH_DIR=./.scratch

DIRS=${RUN_OUTPUTS_DIR} ${RUN_SCRATCH_DIR}

all: run

clean:
	rm -fv ${RUN_OUTPUT_FILE}
	rm -rfv ${DIRS}

run: $(DIRS)
	bash ${RUN_SCRIPT_FILE} > ${RUN_OUTPUT_FILE}

$(DIRS):
	mkdir -p $@

