CC       = g++ -Wall -Werror -O3
CC_DEBUG = g++ -Wall -Werror -g -fprofile-arcs -ftest-coverage -pg
LDLIBS   = lib.o graphlib.o bigint.o
C11      = -std=c++11

PROBLEMS = $(basename $(wildcard [0-9][0-9][0-9].c++))

.PHONY: all
all: lib.o libd.o graphlib.o graphlibd.o bigint.o bigintd.o graphlib_test bigint_test $(PROBLEMS)

.PHONY: clean
clean:
	rm -f lib.o libd.o graphlib.o graphlibd.o bigint.o bigintd.o graphlib_test bigint_test perf.data perf.data.old
	rm -f $(PROBLEMS)

015: 015.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

016: 016.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

017: 017.c++ lib.o
	$(CC) $@.c++ lib.o -o $@

018: 018.c++
	$(CC) $@.c++ -o $@

020: 020.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

021: 021.c++ lib.o
	$(CC) $@.c++ lib.o -o $@

022: 022.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

023: 023.c++ lib.o
	$(CC) $@.c++ lib.o -o $@

024: 024.c++ lib.o
	$(CC) $@.c++ lib.o -o $@

025: 025.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

029: 029.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

030: 030.c++
	$(CC) $@.c++ -o $@

031: 031.c++
	$(CC) $@.c++ -o $@

034: 034.c++
	$(CC) $@.c++ -o $@

040: 040.c++ bigint.o lib.o
	$(CC) $(C11) $@.c++ -o $@

043: 043.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

048: 048.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

052: 052.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

053: 053.c++ bigint.o lib.o
	$(CC) bigint.o lib.o $@.c++ -o $@

055: 055.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

056: 056.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

059: 059.c++ bigint.o lib.o
	$(CC) bigint.o lib.o $@.c++ -o $@

067: 067.c++
	$(CC) $@.c++ -o $@

074: 074.c++
	$(CC) $@.c++ -o $@

079: 079.c++
	$(CC) $@.c++ -o $@

232: 232.c++
	$(CC) $@.c++ -o $@

413: 413.c++ bigint.o
	$(CC) $@.c++ bigint.o -o $@

graphlib_test: graphlibd.o graphlib_test.c++
	$(CC_DEBUG) $+ -o $@
	./$@
	gprof $@ gmon.out > $@.gprof

bigint_test: bigintd.o bigint_test.c++
	$(CC_DEBUG) $+ -o $@
	./$@
	gprof $@ gmon.out > $@.gprof

bigint.o: bigint.h++ bigint.c++
	$(CC) -c bigint.c++ -o $@

bigintd.o: bigint.h++ bigint.c++
	$(CC_DEBUG) -c bigint.c++ -o $@

graphlib.o: graphlib.h++ graphlib.c++
	$(CC) -c graphlib.c++ -o $@

graphlibd.o: graphlib.h++ graphlib.c++
	$(CC_DEBUG) -c graphlib.c++ -o $@

lib.o: lib.h++ lib.c++
	$(CC) -c lib.c++ -o $@

libd.o: lib.h++ lib.c++
	$(CC_DEBUG) -c lib.c++ -o $@
