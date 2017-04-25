#!/bin/bash

####################
#args
run_cmd=("${1} ${2}")
id=$3
####################
touch data 

cat Judge/Questions/$id/input.txt | while read line
do
    data=(`echo $line | tr -s ',' ' '`)
    result=$(echo ${data[0]} | eval ${run_cmd})
    judge=$(test $result = ${data[1]}; echo $?) 
    if [ $judge -eq 0 ]; then
        echo "true" >> data
    else
        echo "false" >> data
    fi
done
