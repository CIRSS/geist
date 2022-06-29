#!/usr/bin/env bash

# *****************************************************************************

bash_cell COMPUTE << END_COMMAND

# Run the compute script

python3 ../python/compute.py

END_COMMAND


# *****************************************************************************

bash_cell PRINT1 << END_COMMAND

# Print csv file with final A, B, and C variables

cat products/df_updated.csv

END_COMMAND


# *****************************************************************************

bash_cell PRINT2 << END_COMMAND

# Print csv file with final Fahrenheit, Celsius, and Kelvin variables

cat products/temps_updated.csv

END_COMMAND
