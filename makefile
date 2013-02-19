CC=g++ -Wall -O3
C11=-std=c++11

all: test 413

clean:
	rm -f 413 test bigint.o lib.o

413: 413.c bigint.o lib.o
	$(CC) bigint.o lib.o $@.c -o $@

test: bigint.o test.c
	$(CC) bigint.o test.c -o $@
	./test

bigint.o: bigint.h bigint.c
	$(CC) -c bigint.c -o $@

lib.o: lib.h lib.c
	$(CC) -c lib.c -o $@
