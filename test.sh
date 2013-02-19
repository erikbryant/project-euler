#!/bin/bash

make
if [[ $? -ne 0 ]]
then
  exit 1
fi

echo "time perf record -- ./413"
time perf record -- ./413
mv perf.data ~
