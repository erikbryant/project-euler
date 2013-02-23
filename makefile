CC       = g++ -Wall -O3
LDLIBS   = lib.o graphlib.o bigint.o
C11      = -std=c++11

PROBLEMS = $(basename $(wildcard [0-9][0-9][0-9].c))

all: lib.o graphlib.o bigint.o graphlib_test bigint_test $(PROBLEMS)

.PHONY: clean
clean:
	rm -f lib.o graphlib.o bigint.o graphlib_test bigint_test perf.data perf.data.old
	rm -f $(PROBLEMS)

015: 015.c bigint.o
	$(CC) $@.c bigint.o -o $@

018: 018.c
	$(CC) $@.c -o $@

040: 040.c bigint.o lib.o
	$(CC) $(C11) $@.c -o $@

053: 053.c bigint.o lib.o
	$(CC) bigint.o lib.o $@.c -o $@

059: 059.c bigint.o lib.o
	$(CC) bigint.o lib.o $@.c -o $@

067: 067.c
	$(CC) $@.c -o $@

graphlib_test: graphlib.o graphlib_test.c
	$(CC) $+ -o $@
	./$@

bigint_test: bigint.o bigint_test.c
	$(CC) $+ -o $@
	./$@

bigint.o: bigint.h bigint.c
	$(CC) -c bigint.c -o $@

graphlib.o: graphlib.h graphlib.c
	$(CC) -c graphlib.c -o $@

lib.o: lib.h lib.c
	$(CC) -c lib.c -o $@
