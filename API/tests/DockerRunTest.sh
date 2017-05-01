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
#%%%%%%%%%%%%%%%%%%%%%%%%%%%

############################
# Select Tester 
############################
CMDNAME=`basename $0`

while getopts as: OPT
do
  case $OPT in
    "a" ) FLG_A="TRUE" ;;
    "s" ) FLG_S="TRUE" ; VALUE_S="$OPTARG" ;;
      * ) echo "Usage: $CMDNAME [-a] [-s VALUE]" 1>&2
          exit 1 ;;
  esac
done

if [ "$FLG_A" = "TRUE" ]; then
    ruby
fi

if [ "$FLG_S" = "TRUE" ]; then
  $VALUE_S
fi
#%%%%%%%%%%%%%%%%%%%%%%%%%%%
