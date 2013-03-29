#include "bigint.h"

#include <iostream>

using std::cout;
using std::endl;

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
