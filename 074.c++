#include <iostream>
#include <set>

using namespace std;

unsigned int factorial( unsigned int n )
{
  switch (n)
    {
    case 0:
      return 1;
      break;
    case 1:
      return 1;
      break;
    case 2:
      return 2 * 1;
      break;
    case 3:
      return 3 * 2 * 1;
      break;
    case 4:
      return 4 * 3 * 2 * 1;
      break;
    case 5:
      return 5 * 4 * 3 * 2 * 1;
      break;
    case 6:
      return 6 * 5 * 4 * 3 * 2 * 1;
      break;
    case 7:
      return 7 * 6 * 5 * 4 * 3 * 2 * 1;
      break;
    case 8:
      return 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1;
      break;
    case 9:
      return 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1;
      break;
    }

  return 0;
}

unsigned int factorialSumDigits( unsigned int n )
{
  unsigned int sum = 0;

  while ( n > 0 )
    {
      sum += factorial( n % 10 );
      n /= 10;
    }

  return sum;
}

unsigned int factorialChain( unsigned int n )
{
  unsigned int sum = n;
  unsigned int count = 1;
  std::set<unsigned int> links;

  do {
    links.insert( sum );
    sum = factorialSumDigits( sum );
    if ( links.count( sum ) != 0 || count > 60 )
      {
	// We have just entered a loop
	return count;
      }
    count++;
  } while ( sum != n );

  return count;
}

int main( int argc, char *argv[] )
{
  unsigned int i = 0;
  unsigned int count = 0;

  cout << "factorialChain( 69 )  = " << factorialChain( 69 ) << endl;
  cout << "factorialChain( 78 )  = " << factorialChain( 78 ) << endl;
  cout << "factorialChain( 540 ) = " << factorialChain( 540 ) << endl;

  for ( i=1; i<1000000; i++ )
    {
      unsigned int length = 0;
      length = factorialChain( i );
      if ( length == 60 )
	{
	  count++;
	}
    }

  cout << "Count: " << count << endl;
}
