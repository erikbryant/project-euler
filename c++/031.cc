//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include <stdio.h>

//
// The array needs to already be sorted in ascending order.
//
unsigned int findCombinations( unsigned int shortfall, unsigned int *array, unsigned int arrayLen )
{
  int i = 0;
  unsigned int combinations = 0;

  if ( shortfall == 0 )
    {
      return 1;
    }

  if ( arrayLen <= 0 )
    {
      return 0;
    }

  for ( i = arrayLen - 1; i >= 0; --i )
    {
      if ( array[i] <= shortfall )
        {
          if ( shortfall - array[i] >= array[i] )
            {
              combinations += findCombinations( shortfall - array[i], array, i + 1 );
            }
          else
            {
              combinations += findCombinations( shortfall - array[i], array, i );
            }
        }
    }

  return combinations;
}


int main( int argc, char **argv )
{
  unsigned int coins[] = { 1, 2, 5, 10, 20, 50, 100, 200 };
  unsigned int coinsCount = sizeof( coins) / sizeof( unsigned int );
  unsigned int combinations = 0;

  combinations = findCombinations( 1, coins, coinsCount );
  printf( "Combinations to make 1: %d\n", combinations );

  combinations = findCombinations( 3, coins, coinsCount );
  printf( "Combinations to make 3: %d\n", combinations );

  combinations = findCombinations( 5, coins, coinsCount );
  printf( "Combinations to make 5: %d\n", combinations );

  combinations = findCombinations( 200, coins, coinsCount );
  printf( "Combinations to make 200: %d\n", combinations );
}
