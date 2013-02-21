CC       = g++ -Wall -O3
LDLIBS   = bigint.o lib.o
C11      = -std=c++11

PROBLEMS = $(basename $(wildcard [0-9][0-9][0-9].c))

all: lib.o bigint.o test $(PROBLEMS)

.PHONY: clean
clean:
	rm -f test bigint.o lib.o perf.data perf.data.old
	rm -f $(PROBLEMS)

040: 040.c bigint.o lib.o
	$(CC) $(C11) $@.c -o $@

053: 053.c bigint.o lib.o
	$(CC) bigint.o lib.o $@.c -o $@

test: bigint.o test.c
	$(CC) $+ -o $@
	./$@

bigint.o: bigint.h bigint.c
	$(CC) -c bigint.c -o $@

lib.o: lib.h lib.c
	$(CC) -c lib.c -o $@
