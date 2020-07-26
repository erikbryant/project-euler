#include <stdio.h>
#include "lib.h"

int main( int argc, char **argv )
{
  unsigned int i = 0;
  unsigned int sum1 = 0;
  unsigned int sum2 = 0;
  unsigned int total = 0;

  for ( i=0; i<=10000; i++ )
  {
    sum1 = sumOfDivisors( i );
    sum2 = sumOfDivisors( sum1 );
    if ( i == sum2 && i < sum1 )
    {
      printf( "%d <==> %d\n", i, sum1 );
      total += sum1 + sum2;
    }
  }

  printf( "Total: %d\n", total );
}

