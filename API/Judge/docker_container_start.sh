################################
imageName="ruby:2.3.4-alpine" # imageName=$1
run_cmd=("${1} ${2}")
id=$3
################################

docker run --name test \
    -i -v $(pwd)/Judge/WorkSpace:/WorkSpace \
    -t $imageName /WorkSpace/judge_run.sh $run_cmd $id
