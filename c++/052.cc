#include <iostream>
#include "bigint.h"

using namespace std;

int main( int argc, char *argv[] )
{
  unsigned int digits = 0;
  unsigned int i = 0;;
  BigInt candidate = 0;
  BigInt tester    = 0;

  while ( 1 )
  {
    bool found = true;

    digits = candidate.uniqueDigits();
    tester = candidate;
    for ( i=2; i<=6; i++ )
    {
      tester += candidate;
      if ( tester.uniqueDigits() != digits )
      {
        found = false;
        break;
      }
    }

    if ( found )
    {
      cout << "Found: " << candidate << endl;
      cout << "   2x: " << 2 * candidate << endl;
      cout << "   3x: " << 3 * candidate << endl;
      cout << "   4x: " << 4 * candidate << endl;
      cout << "   5x: " << 5 * candidate << endl;
      cout << "   6x: " << 6 * candidate << endl;
    }
    candidate++;
  }
}
