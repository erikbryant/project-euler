all: test solutions

solutions:
	g++ 016.c lib.c bigint.c
	g++ 020.c lib.c bigint.c
	g++ 021.c lib.c bigint.c
	g++ 022.c lib.c bigint.c
	g++ 023.c lib.c bigint.c
	g++ 030.c lib.c bigint.c
	g++ 031.c lib.c bigint.c
	g++ 048.c lib.c bigint.c
	g++ 056.c lib.c bigint.c

test:
	g++ test.c lib.c bigint.c
	./a.out
