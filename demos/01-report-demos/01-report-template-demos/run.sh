#!/usr/bin/env bash

# *****************************************************************************

bash_cell static_template << 'END_CELL'

geist report << 'END_TEMPLATE'

    Materials report
    ================
    42 items are made of cotton

END_TEMPLATE

END_CELL


# # *****************************************************************************

bash_cell template_with_variable << 'END_CELL'

geist report << 'END_TEMPLATE'

    {{ $CottonItemCount := 42 }}

    Materials report
    ================
    {{ $CottonItemCount }} items are made of cotton

END_TEMPLATE

END_CELL


# *****************************************************************************

bash_cell printf_with_constant << 'END_CELL'

geist report << 'END_TEMPLATE'

    Materials report
    ================
    {{ printf "%d items are made of cotton" 42 }}

END_TEMPLATE

END_CELL


# *****************************************************************************

bash_cell printf_with_variable << 'END_CELL'

geist report << 'END_TEMPLATE'

    {{ $CottonItemCount := 42 }}

    Materials report
    ================
    {{ printf "%d items are made of cotton" $CottonItemCount }}

END_TEMPLATE

END_CELL

