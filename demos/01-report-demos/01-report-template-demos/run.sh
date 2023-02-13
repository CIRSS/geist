#!/usr/bin/env bash

# *****************************************************************************

bash_cell static_template << END_CELL

geist report << '__END_REPORT_TEMPLATE__'

    Materials report
    ================
    42 items are made of cotton

__END_REPORT_TEMPLATE__

END_CELL


# # *****************************************************************************

bash_cell template_with_variable << 'END_CELL'

geist report << '__END_REPORT_TEMPLATE__'

    {{ $CottonItemCount := 42 }}

    Materials report
    ================
    {{ $CottonItemCount }} items are made of cotton

__END_REPORT_TEMPLATE__

END_CELL


# *****************************************************************************

bash_cell printf_with_constant << END_CELL

geist report << '__END_REPORT_TEMPLATE__'

    Materials report
    ================
    {{ printf "%d items are made of cotton" 42 }}

__END_REPORT_TEMPLATE__

END_CELL


# *****************************************************************************

bash_cell printf_with_variable << 'END_CELL'

geist report << '__END_REPORT_TEMPLATE__'

    {{ $CottonItemCount := 42 }}

    Materials report
    ================
    {{ printf "%d items are made of cotton" $CottonItemCount }}

__END_REPORT_TEMPLATE__

END_CELL

