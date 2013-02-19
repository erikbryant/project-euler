CC=g++ -Wall -O3
C11=-std=c++11

PROBLEMS=016 017 020 021 022 023 024 025 029 030 031 040 048 056 413

all: test $(PROBLEMS)

clean:
	rm -f 413 test bigint.o lib.o $(PROBLEMS) perf.data perf.data.old

$(PROBLEMS): $(addsuffix .c,$*) bigint.o lib.o
	$(CC) $(C11) $@.c bigint.o lib.o -o $@

test: bigint.o test.c
	$(CC) bigint.o test.c -o $@
	./test

bigint.o: bigint.h bigint.c
	$(CC) -c bigint.c -o $@

lib.o: lib.h lib.c
	$(CC) -c lib.c -o $@
