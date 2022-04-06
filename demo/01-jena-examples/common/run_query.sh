#!/usr/bin/env bash

# save command line arguments
data_file=$1
query_id=$2
query_description=$3

if [[ -n REPRO_DEMO_TMP_DIRNAME ]] ; then
    tmp_dir=${REPRO_DEMO_TMP_DIRNAME}
    mkdir -p ${tmp_dir}
else
    tmp_dir=$(mktemp --tmpdir=/tmp -d bash_dev_XXXXXXX)
fi

# define names of query and query result files
query_file=${REPRO_DEMO_TMP_DIRNAME}/${query_id}.rq
result_file=${REPRO_DEMO_TMP_DIRNAME}/${query_id}.txt

# copy query from stdin to the query file
printf "# ${query_description}\n\n" > ${query_file}

IFS=''; while read line
do
    printf "$line\n" >> ${query_file}
done

# execute the query on the given data file and save results
java -cp "${JENA_CLASSPATH}" arq.sparql --syntax arq --data ${data_file} --query ${query_file} > ${result_file}

# print the query results
echo
echo "**************************** QUERY ${query_id} **********************************"
echo
cat ${query_file}
echo
echo "--------------------- RESULTS FOR QUERY ${query_id}------------------------------"
echo
cat ${result_file}
echo
