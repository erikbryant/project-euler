#include <iostream>
#include "bigint.h"

using namespace std;

BigInt countRoutesRightAndDown( int width, int height )
{
  if ( width == 1 )
    {
      return BigInt( height + 1 );
    }

  if ( height == 1 )
    {
      return BigInt( width + 1 );
    }

  if ( width == 2 && height == 2 )
    {
      return BigInt( 6 );
    }

  if ( width == 3 && height == 3 )
    {
      return BigInt( 20 );
    }

  if ( width == height )
    {
      return countRoutesRightAndDown( width - 1, height ) * 2;
    }
  else
    {
      return countRoutesRightAndDown( width - 1, height ) +
	countRoutesRightAndDown( width, height - 1 );
    }
}

int main( int argc, char *argv[] )
{
  cout << "Routes through 1x1 grid (should be 2): " << countRoutesRightAndDown( 1, 1 ) << endl;

  cout << "Routes through 2x2 grid (should be 6): " << countRoutesRightAndDown( 2, 2 ) << endl;

  /*
  int i = 0;

  for ( i=3; i<=19; i++ )
    {
      cout << "Routes through " << i << "x" << i << " grid: " << countRoutesRightAndDown( i, i ) << endl;
    }
  */

  cout << "Routes through 20x20 grid: " << countRoutesRightAndDown( 20, 20 ) << endl;
}
