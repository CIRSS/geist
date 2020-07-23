#!/usr/bin/env bash

# save command line arguments
data_file=$1
query_id=$2
query_description=$3

# define names of query and query result files
query_file=queries/${query_id}.rq
result_file=queries/${query_id}.txt

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
