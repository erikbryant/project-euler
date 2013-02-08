#include <stdio.h>
#include "lib.h"
#include "bigint.h"

int main( int argc, char **argv )
{
  unsigned int i = 0;
  BigInt sum = 0;

  for ( i=1; i <=1000; i++ )
  {
    BigInt x = i;

    sum += x.power( x );
  }

  cout << "sum = " << sum << endl;
}
