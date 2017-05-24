#!/bin/sh

####################
#args
run_cmd="${1} ${2}"
####################

echo "---------- Starting Judge ----------"

cat WorkSpace/input.txt | while read line
do
    input=$(echo $line | cut -d "," -f1)
    output=$(echo $line | cut -d "," -f2)
    result=$(echo $input | eval ${run_cmd})
    judge=$(test $result = $output; echo $?) 
    if [ $judge -eq 0 ]; then
        echo "true" >> WorkSpace/data
    else
        echo "false" >> WorkSpace/data
    fi
done
