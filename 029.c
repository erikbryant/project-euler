#include <iostream>
#include <set>
#include "bigint.h"
#include "lib.h"

using namespace std;

int main( int argc, char **argv )
{
  BigInt a = 0;
  BigInt b = 1;
  BigInt c = 0;
  std::set<BigInt> myset;

  for ( a=2; a<=100; a++ )
  {
    for ( b=2; b<=100; b++ )
    {
      c = a.power( b );
      myset.insert( c );
    }
  }

  cout << "Power set has " << myset.size() << " unique members." << endl;
}
