#include <stdio.h>

#include "lib.h"

char isSummable( unsigned int n, unsigned int *array, unsigned int arrayCount )
{
  unsigned int i = 0;
  unsigned int difference = 0;

  for ( i=0; i<arrayCount && n > array[i]; i++ )
  {
    difference = n - array[i];
    if ( arrayContains( difference, array, arrayCount ) )
    {
      return 1;
    }
  }

  return 0;
}

int main( int argc, char **argv )
{
  const unsigned int MAX = 28122;
  unsigned int i = 0;
  unsigned int abundants[MAX];
  unsigned int abundantsCount = 0;
  unsigned int total = 0;

  abundantsCount = findAbundants( MAX, abundants );

  for ( i=0; i<=MAX; i++ )
  {
    if ( !isSummable( i, abundants, abundantsCount ) )
    {
      total += i;
    }
  }
  printf( "Total: %d\n", total );
}

