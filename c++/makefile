#
# Copyright Erik Bryant (erikbryantology@gmail.com)
#

C        = gcc -Wall -Werror -O3
CC       = g++ -Wall -Werror -Weffc++ -O3
CC_DEBUG = g++ -Wall -Werror -Weffc++ -D_GLIBCXX_DEBUG -g -fprofile-arcs -ftest-coverage -pg
C11      = -std=c++11
CPPCHECK = ../cppcheck-1.58/cppcheck

PROBLEMS = $(basename $(wildcard [0-9][0-9][0-9].cc)) 413

.PHONY: all
all: lib.o libd.o bigint.o bigintd.o graphlib_test bigint_test $(PROBLEMS)

.PHONY: clean
clean:
	rm -f lib.o libd.o bigint.o bigintd.o graphlib_test bigint_test
	rm -f perf.data perf.data.old
	rm -f *.gcov *.gcda *.gcno *.gprof \#*# gmon.out
	rm -f a.out $(PROBLEMS)

015: 015.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

016: 016.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

017: 017.cc lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

018: 018.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

019: 019.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

020: 020.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

021: 021.cc lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

022: 022.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

023: 023.cc lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

024: 024.cc lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

025: 025.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

028: 028.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

029: 029.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

030: 030.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

031: 031.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

034: 034.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

040: 040.cc bigint.o lib.o
	$(CPPCHECK) $@.cc
	$(CC) $(C11) $^ -o $@

043: 043.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

048: 048.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

052: 052.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

053: 053.cc bigint.o lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

055: 055.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

056: 056.cc bigint.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

059: 059.cc bigint.o lib.o
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

067: 067.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

074: 074.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

079: 079.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

081: 081.cc graphlib.h
	$(CPPCHECK) $@.cc
	$(CC) $(C11) $@.cc -o $@

082: 082.cc graphlib.h
	$(CPPCHECK) $@.cc
	$(CC) $(C11) $@.cc -o $@

083: 083.cc graphlib.h
	$(CPPCHECK) $@.cc
	$(CC) $(C11) $@.cc -o $@

089: 089.cc
	$(CPPCHECK) $@.cc
	$(CC) $(C11) $@.cc -o $@

107: 107.cc graphlib.h
	$(CPPCHECK) $@.cc
	$(CC_DEBUG) $@.cc -o $@

232: 232.cc
	$(CPPCHECK) $@.cc
	$(CC) $^ -o $@

413: 413.c
	$(C) $^ -o $@

419: 419.cc
	$(CPPCHECK) $@.cc
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
