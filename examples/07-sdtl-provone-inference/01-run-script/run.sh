#!/usr/bin/env bash


RUNNER='../../common/run_command.sh'

# *****************************************************************************

bash ${RUNNER} COMPUTE "Run the compute script" << END_COMMAND

python3 ../python/compute.py

END_COMMAND


# *****************************************************************************

bash ${RUNNER} "" "Print csv file with final A, B, and C variables" << END_COMMAND

cat outputs/df_updated.csv

END_COMMAND


# *****************************************************************************

bash ${RUNNER} "" "Print csv file with final Fahrenheit, Celsius, and Kelvin variables" << END_COMMAND

cat outputs/temps_updated.csv

END_COMMAND
