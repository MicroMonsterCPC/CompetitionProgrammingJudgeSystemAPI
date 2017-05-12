#!/bin/sh

####################
#args
run_cmd="${1} ${2}"
####################

echo $run_cmd

cat WorkSpace/input.txt | while read line
do
    input=$(echo $line | cut -d "," -f1)
    output=$(echo $line | cut -d "," -f2)
    result=$(echo $input | eval ${run_cmd})
    judge=$(test $result = $output; echo $?) 
    if [ $judge -eq 0 ]; then
        echo "true"
    else
        echo "false"
    fi
done
