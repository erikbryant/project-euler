#include <iostream>
#include <stdlib.h>
#include "bigint.h"

using namespace std;

BigInt factorial( unsigned int n, unsigned int limit=0 )
{
  BigInt result = 1;

  while ( n > limit )
  {
    result *= n;
    n--;
  }

  return result;
}

BigInt combinatorial( unsigned int n, unsigned int r )
{
//
//  nCr = n! / ( r! * (n - r)! )
//
  if ( r > (n - r) )
  {
    return factorial( n, r ) / factorial( n - r );
  } else {
    return factorial( n, (n - r) ) / factorial( r );
  }
}

int main( int argc, char *argv[] )
{
  unsigned int n = 0;
  unsigned int r = 0;
  unsigned int count = 0;
  unsigned int max = 100;
  BigInt nCr = 0;

  for ( n=1; n<=max; n++ )
  {
    nCr = n;
    for ( r=2; r<=n/2; r++ )
    {
      nCr *= n - (r - 1);
      nCr = nCr / r;
      if ( nCr > 1000000 )
      {
        // nCr is mirrored around n/2
        // values increase from r=1 to r=n/2
        cout << n << "C" << r << " & " << n << "C" << n-r << " " << nCr << endl;
        count += (n - r) - (r - 1);
        break;
      }
    }
  }

  cout << "Count: " << count << endl;
}
