#!/usr/bin/env bash

# *****************************************************************************

bash_dot_cell static_dot_file << END_CELL

geist report << '__END_REPORT_TEMPLATE__'

    digraph static_dot_file {
    rankdir=BT
    B -> A
    C -> A
    }

__END_REPORT_TEMPLATE__

END_CELL

