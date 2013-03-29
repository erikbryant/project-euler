#
# Copyright Erik Bryant (erikbryantology@gmail.com)
# Gnu All-Permissive http://www.gnu.org/philosophy/license-list.html#GNUAllPermissive
#

C        = gcc -Wall -Werror -O3
CC       = g++ -Wall -Werror -Weffc++ -O3
CC_DEBUG = g++ -Wall -Werror -Weffc++ -D_GLIBCXX_DEBUG -g -fprofile-arcs -ftest-coverage -pg
C11      = -std=c++11
CPPCHECK = ../cppcheck-1.58/cppcheck

PROBLEMS = $(basename $(wildcard [0-9][0-9][0-9].c++)) 413

.PHONY: all
all: lib.o libd.o bigint.o bigintd.o graphlib_test bigint_test $(PROBLEMS)

.PHONY: clean
clean:
	rm -f lib.o libd.o bigint.o bigintd.o graphlib_test bigint_test
	rm -f perf.data perf.data.old
	rm -f *.gcov *.gcda *.gcno *.gprof \#*# gmon.out
	rm -f a.out $(PROBLEMS)

015: 015.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

016: 016.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

017: 017.c++ lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

018: 018.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

019: 019.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

020: 020.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

021: 021.c++ lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

022: 022.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

023: 023.c++ lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

024: 024.c++ lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

025: 025.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

028: 028.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

029: 029.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

030: 030.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

031: 031.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

034: 034.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

040: 040.c++ bigint.o lib.o
	$(CPPCHECK) $@.c++
	$(CC) $(C11) $^ -o $@

043: 043.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

048: 048.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

052: 052.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

053: 053.c++ bigint.o lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

055: 055.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

056: 056.c++ bigint.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

059: 059.c++ bigint.o lib.o
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

067: 067.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

074: 074.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

079: 079.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

081: 081.c++ graphlib.h
	$(CPPCHECK) $@.c++
	$(CC) $(C11) $@.c++ -o $@

082: 082.c++ graphlib.h
	$(CPPCHECK) $@.c++
	$(CC) $(C11) $@.c++ -o $@

083: 083.c++ graphlib.h
	$(CPPCHECK) $@.c++
	$(CC) $(C11) $@.c++ -o $@

089: 089.c++
	$(CPPCHECK) $@.c++
	$(CC) $(C11) $@.c++ -o $@

107: 107.c++ graphlib.h
	$(CPPCHECK) $@.c++
	$(CC_DEBUG) $@.c++ -o $@

232: 232.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

413: 413.c
	$(C) $^ -o $@

419: 419.c++
	$(CPPCHECK) $@.c++
	$(CC) $^ -o $@

graphlib_test: graphlib.h graphlib_test.cc
	$(CPPCHECK) graphlib_test.cc
	$(CC_DEBUG) $@.cc -o $@
	./$@
	gprof $@ gmon.out > $@.gprof
	gcov $@ > /dev/null

bigint_test: bigintd.o bigint_test.cc
	$(CPPCHECK) bigint_test.cc
	$(CC_DEBUG) $^ -o $@
	./$@
	gprof $@ gmon.out > $@.gprof
	gcov bigintd > /dev/null

bigint.o: bigint.h bigint.cc
	$(CPPCHECK) bigint.cc
	$(CC) -c bigint.cc -o $@

bigintd.o: bigint.h bigint.cc
	$(CC_DEBUG) -DDO_VALIDATION -c bigint.cc -o $@

lib.o: lib.h lib.cc
	$(CPPCHECK) lib.cc
	$(CC) -c lib.cc -o $@

libd.o: lib.h lib.cc
	$(CC_DEBUG) -c lib.cc -o $@
