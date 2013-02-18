all: test solutions

solutions: lib.h lib.c bigint.h bigint.c 016.c 017.c 020.c 021.c 022.c 023.c 024.c 025.c 029.c 030.c 031.c 040.c 048.c 056.c
	g++ -Wall -O3 -std=c++11 016.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 017.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 020.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 021.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 022.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 023.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 024.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 025.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 029.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 030.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 031.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 040.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 048.c lib.c bigint.c
	g++ -Wall -O3 -std=c++11 056.c lib.c bigint.c
	touch solutions

test: test.c lib.h lib.c bigint.h bigint.c
	g++ -Wall -O3 -std=c++11 test.c lib.c bigint.c -o test
	perf record -- ./test
	mv perf.data ~
