#include "bigint.h"

int main( int argc, char **argv )
{
  unsigned int i = 0;
  unsigned int j = 0;
  unsigned int max_i = 0;
  unsigned int max_j = 0;
  BigInt max_sum = 0;

  for ( i=2; i<100; i++ )
  {
    BigInt base = i;
    for ( j=1; j<100; j++ )
    {
      BigInt exponent = j;
      BigInt result = base.power( exponent );
      BigInt sum = result.sumDigits();
      if ( sum >= max_sum )
      {
        max_sum = sum;
        max_i = i;
        max_j = j;
      }
//      cout << i << "**" << j << " = " << result << ", sum = " << sum << endl;
    }
  }

  cout << "Max sum: " << max_sum << " found at " << max_i << "**" << max_j << endl;
}
