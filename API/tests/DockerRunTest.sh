#!/bin/sh

############################
# StartupTest for Ruby
############################

ruby() {
    scriptlang rb ruby:2.3.4-alpine
}
#%%%%%%%%%%%%%%%%%%%%%%%%%%%



############################
# StartupTest for Python
############################
python() {
    scriptlang py python:2-alpine
}
#%%%%%%%%%%%%%%%%%%%%%%%%%%%



############################
# ScriptLang function
############################
scriptlang () {
    cp testfiles/Main.$1 WorkSpace
    lang_name=$(echo $2 | cut -d ":" -f1 )
    docker run -i \
        -v $(pwd)/WorkSpace:/WorkSpace \
        -t  $2 \
        WorkSpace/judge_run.sh $lang_name WorkSpace/Main.$1

    if [ $? -eq 0 ]; then
        echo "Succses to Process there $lang_name"
    else
        echo "Error to Process there $lang_name"
    fi
    rm WorkSpace/Main.$1
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
    python
fi

if [ "$FLG_S" = "TRUE" ]; then
    $VALUE_S
fi
#%%%%%%%%%%%%%%%%%%%%%%%%%%%
