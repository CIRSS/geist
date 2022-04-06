#!/usr/bin/env bash

# *****************************************************************************

run_cell COMPUTE "Run the compute script" << END_COMMAND

python3 ../python/compute.py

END_COMMAND


# *****************************************************************************

run_cell PRINT1 "Print csv file with final A, B, and C variables" << END_COMMAND

cat products/df_updated.csv

END_COMMAND


# *****************************************************************************

run_cell PRINT2 "Print csv file with final Fahrenheit, Celsius, and Kelvin variables" << END_COMMAND

cat products/temps_updated.csv

END_COMMAND
