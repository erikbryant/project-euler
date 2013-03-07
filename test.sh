#!/bin/bash
#
# Copyright Erik Bryant (erikbryantology@gmail.com)
# Gnu All-Permissive http://www.gnu.org/philosophy/license-list.html#GNUAllPermissive
#

make
if [[ $? -ne 0 ]]
then
  exit 1
fi

PROG=$1
shift

echo "time perf record -- ./$PROG $*"
time perf record -- ./$PROG $*
mv perf.data ~
