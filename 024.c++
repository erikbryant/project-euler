#include <stdio.h>
#include "lib.h"

//
// For some reason mathmaticians define
// factorial( 0 ) to be 1.
//
unsigned int factorial( unsigned int f )
{
  unsigned int result = 1;

  while ( f > 0 )
  {
    result *= f;
    f--;
  }

  return result;
}

int main( int argc, char **argv )
{
  unsigned int array[] = { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
  //
  // Iterate 1 time to get the 2nd value.
  // Iterate 999,999 times to get the 1,000,000th value.
  //
  unsigned int target = 999999;
  unsigned int i=0;
  unsigned int f=0;

  while ( target > 0 )
  {
    for ( i=10; i>=1; i-- )
    {
      f = factorial( i );
      if ( f <= target )
      {
        printf( "%d\n", i );
        array[9-i]++;
        target -= f;
        break;
      }
    }
  }

  printArray( array, 10 );
}
