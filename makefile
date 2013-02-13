all: test solutions

solutions: lib.h lib.c bigint.h bigint.c 016.c 020.c 021.c 022.c 023.c 030.c 031.c 048.c 056.c
	g++ -O3 016.c lib.c bigint.c
	g++ -O3 020.c lib.c bigint.c
	g++ -O3 021.c lib.c bigint.c
	g++ -O3 022.c lib.c bigint.c
	g++ -O3 023.c lib.c bigint.c
	g++ -O3 030.c lib.c bigint.c
	g++ -O3 031.c lib.c bigint.c
	g++ -O3 048.c lib.c bigint.c
	g++ -O3 056.c lib.c bigint.c
	touch solutions

test: test.c lib.h lib.c bigint.h bigint.c
	g++ -O3 test.c lib.c bigint.c -o test
	perf record -- ./test
	mv perf.data ~
