#include <iostream>

using namespace std;

unsigned int factorial( unsigned int n )
{
  unsigned int result = 1;

  while ( n > 1 )
  {
    result *= n;
    n--;
  }

  return result;
}

unsigned int countDigits( unsigned int n )
{
  unsigned int result = 0;

  while ( n > 0 )
  {
    result++;
    n = n / 10;
  }

  return result;
}

bool isDigitFactorial( unsigned int n )
{
  unsigned int i = n;
  unsigned int sum = 0;

  while ( i > 0 )
  {
    sum += factorial( i % 10 );
    i = i / 10;
  }

  return sum == n;
}

int main( int argc, char *argv[] )
{
  unsigned int i = 0;
  unsigned int max = 0;
  unsigned int sum = 0;

  for ( i=1; i<100; i++ )
  {
    if ( countDigits( i * factorial( 9 ) ) < i )
    {
      cout << "Cross-over at i = " << i << endl;
      max = (i - 1) * factorial( 9 );
      break;
    }
  }

  cout << "Max number to inspect: " << max << endl;

  // We are told to skip 1 and 2
  for ( i=3; i<=max; i++ )
  {
    if ( isDigitFactorial( i ) )
    {
      cout << "i = " << i << endl;
      sum += i;
    }
  }

  cout << "Sum = " << sum << endl;
}
