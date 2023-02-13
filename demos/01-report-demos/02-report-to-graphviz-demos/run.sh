#!/usr/bin/env bash

# *****************************************************************************

bash_dot_cell static_dot_file << 'END_CELL'

geist report << 'END_TEMPLATE'

    digraph static_dot_file {
    rankdir=BT
    B -> A
    C -> A
    }

END_TEMPLATE

END_CELL

