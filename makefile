all: test solutions

solutions: lib.h lib.c bigint.h bigint.c 016.c 020.c 021.c 022.c 023.c 030.c 031.c 048.c 056.c
	g++ 016.c lib.c bigint.c
	g++ 020.c lib.c bigint.c
	g++ 021.c lib.c bigint.c
	g++ 022.c lib.c bigint.c
	g++ 023.c lib.c bigint.c
	g++ 030.c lib.c bigint.c
	g++ 031.c lib.c bigint.c
	g++ 048.c lib.c bigint.c
	g++ 056.c lib.c bigint.c
	touch solutions

test: test.c lib.h lib.c bigint.h bigint.c
	g++ test.c lib.c bigint.c -o test
	./test
