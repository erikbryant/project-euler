#!/bin/bash

make
if [[ $? -ne 0 ]]
then
  exit
fi

CMD="g++ -Wall -O3 -std=c++11 bigint.c lib.c $*"
echo $CMD
$CMD
if [[ $? -eq 0 ]]
then
  time perf record -- ./a.out
  mv perf.data ~
fi
