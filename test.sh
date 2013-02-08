#!/bin/bash

make
if [[ $? -ne 0 ]]
then
  exit
fi

CMD="g++ bigint.c lib.c $*"
echo $CMD
$CMD
if [[ $? -eq 0 ]]
then
  time ./a.out
fi
