#!/bin/sh

############################
# StartupTest for Ruby
############################

ruby() {
    docker run -i \
        -v $(pwd)/WorkSpace:/WorkSpace \
        -t ruby:2.3.4-alpine \
        WorkSpace/judge_run.sh ruby WorkSpace/Main.rb

    if [ $? -eq 0 ]; then
        echo "Succses to Process there Ruby"
    else
        echo "Error to Process there Ruby"
    fi
}
############################

$1;
