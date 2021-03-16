#!/bin/bash

# # avoid error message on ctrl-c
# cleanup() {
#     echo
#     exit 0
# }
# trap cleanup EXIT

# run the service
echo
echo "--------------------------------------------------------------------------"
echo "The Blazegraph service is now running in the REPRO."
echo "Connect to it by navigating in a web browser to http://localhost:9999 "
echo
echo "Terminate the service by pressing the 'q' key in this terminal."
echo "--------------------------------------------------------------------------"
while [ true ] ; do
    read -n 1 key
    if [[ $key = 'q' ]] ; then
        echo
        exit
    else
        echo 
        echo "Type 'q' key to stop the Blazegraph service."
    fi

done
