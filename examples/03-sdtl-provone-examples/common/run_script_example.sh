#!/usr/bin/env bash

# save command line arguments
# data_file=$1
script_id=$1
script_description=$2

# define name of result file
script_file=results/${script_id}.sh
result_file=results/${script_id}.txt

# copy query from stdin to the query file
printf "# ${script_description}\n" > ${script_file}
IFS=''; while read line
do
    printf "$line\n" >> ${script_file}
done

# execute the script on the given data file and save results
bash ${script_file} > ${result_file}

# print the script results
echo
echo "**************************** EXAMPLE ${script_id} *********************************"
echo
cat ${script_file}
echo "---------------------------- ${script_id} OUTPUTS ---------------------------------"
echo
cat ${result_file}
