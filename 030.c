#include <stdio.h>
#include <math.h>

#include "lib.h"

unsigned int Precompute[10];

void precompute( unsigned int exponent )
{
  unsigned int i = 0;

  for ( i=0; i <=9; i++ )
  {
    Precompute[i] = power( i, exponent );
  }
}

unsigned int sumOfPowers( unsigned int n, unsigned int exponent )
{
  unsigned int digit = 0;
  unsigned int sum = 0;

  while ( n > 0 )
  {
    digit = n % 10;
    sum += Precompute[digit];
    n = n / 10;
  }

  return sum;
}


int main( int argc, char **argv )
{
  precompute( 5 );

  unsigned int i = 0;
  unsigned int sum = 0;
  unsigned int total = 0;

  //
  // 9**5 = 59,049
  // 7 digits of nines adds to a 6-digit number
  // Thus, no sums of powers can be 7 digits or
  // longer. Test only up through 6-digits.
  //
  for ( i=10; i<=999999; i++ )
  {
    sum = sumOfPowers( i, 5 );
    if ( i == sum )
    {
      printf( "%d\n", i );
      total += sum;
    }
  }

  printf( "Total: %d\n", total );
}
