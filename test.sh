#!/bin/bash

CMD="g++ bigint.c lib.c test.c"
echo $CMD
$CMD
if [[ $? -eq 0 ]]
then
  ./a.out
  if [[ $? -ne 0 ]]
  then
    exit
  fi
else
  exit
fi

CMD="g++ bigint.c lib.c $*"
echo $CMD
$CMD
if [[ $? -eq 0 ]]
then
  time ./a.out
fi
