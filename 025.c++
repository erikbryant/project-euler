#include "bigint.h"

#include <iostream>

using std::cout;
using std::endl;

int main( int argc, char **argv )
{
  BigInt a = 1;
  BigInt b = 1;
  BigInt sum = 2;
  unsigned int term = 3;

  while ( sum.length() < 1000 )
  {
    BigInt temp = sum;
    sum += b;
    a = b;
    b = temp ;
    term++;
  }

  cout << "term = " << term << endl;
}
