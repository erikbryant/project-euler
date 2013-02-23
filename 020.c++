#include <stdio.h>
#include "lib.h++"
#include "bigint.h++"

int main( int argc, char **argv )
{
  unsigned int i = 0;
  BigInt product = 1;

  for ( i=1; i <=100; i++ )
  {
    BigInt factor = i;

    product *= factor;
  }

  cout << "100! = " << product << endl;
  cout << "Sum of digits = " << product.sumDigits() << endl;
}

