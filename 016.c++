#include <stdio.h>
#include "bigint.h++"

int main( int argc, char **argv )
{
  unsigned int i = 0;
  BigInt product = 1;
  BigInt factor  = 2;

  for ( i=1; i <=1000; i++ )
  {
    product *= factor;
  }

  cout << "2**" << i-1 << " = " << product << endl;
  cout << "Sum of digits = " << product.sumDigits() << endl;
}

