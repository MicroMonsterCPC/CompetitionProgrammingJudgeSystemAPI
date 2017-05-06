#!/bin/sh
################################
run_cmd="${1} ${2}"
imageName=$3
################################

echo "---------- Starting Docker ----------"

 docker run \
     -v $(pwd)/Judge/WorkSpace:/WorkSpace \
     -t $imageName WorkSpace/judge_run.sh $run_cmd
