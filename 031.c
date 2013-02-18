#include <stdio.h>

#include "lib.h"

//
// The array needs to already be sorted in ascending order.
//
char findCombinations( unsigned int target, unsigned int *array, unsigned int count )
{
  unsigned int remainder = target;
  unsigned int i = 0;

  if ( count == 0 )
  {
    return 0;
  }

  for ( i=count-1; i >= 0; i-- )
  {
    while ( array[i] <= remainder )
    {
      remainder -= array[i];
      if ( remainder == 0 || findCombinations( remainder, array, i+1 ) )
      {
        printf( "%d ", array[i] );
        return 1;
      }
    }
    count--;
  }

  return 0;
}


int main( int argc, char **argv )
{
  unsigned int coins[] = { 1, 2, 5, 10, 20, 50, 100, 200 };
  unsigned int coinsCount = sizeof( coins) / sizeof( unsigned int );

  if ( findCombinations( 240, coins, coinsCount ) )
  {
    printf( "\n" );
  }
}
