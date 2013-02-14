#include <iostream>
#include <string>
#include "bigint.h"

using namespace std;

int main( int argc, char **argv )
{
  string foo = "X";
  unsigned int i = 0;
  unsigned int product = 1;

  for ( i=1; i<1000000; i++ )
  {
    foo += to_string( i );
  }

  for ( i=1; i<=1000000; i*=10 )
  {
    cout << "d(" << i << ") = " << foo[i] << endl;
    product *= foo[i] - '0';
  }

  cout << "Product = " << product << endl;
}
