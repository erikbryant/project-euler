#!/bin/bash

make
if [[ $? -ne 0 ]]
then
  exit 1
fi

echo "time perf record -- ./$1"
time perf record -- ./$1
mv perf.data ~
