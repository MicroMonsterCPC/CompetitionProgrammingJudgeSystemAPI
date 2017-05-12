#!/bin/sh

../Judge/docker_container_start.sh ruby WorkSpace/Main.rb

if [ $? -eq 0 ]; then
    echo "Succses to Process"
else
    echo "Error to Process"
fi

