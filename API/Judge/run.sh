#!/bin/bash

##########################
#テスト用
cmd="ruby main.rb"
id="D001"
##########################

cat Questions/$id/input.txt | while read line
do
    data=(`echo $line | tr -s ',' ' '`)
    result=$(echo ${data[0]} | eval ${cmd})
    judge=$(test $result = ${data[1]}; echo $?) 
    if [ $judge -eq 0 ]; then
        echo "true" >> data
    else
        echo "false" >> data
    fi
done
