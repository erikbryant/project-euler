//
// Copyright Erik Bryant (erikbryantology@gmail.com)
// GPLv2 http://www.gnu.org/licenses/gpl-2.0.html
//

#include <iostream>

using std::cout;
using std::endl;

unsigned int sumDiagonals( unsigned int width )
{
  unsigned int sum = 0;
  unsigned int i = 0;
  unsigned int corner = 0;

  if ( width % 2 == 0 )
    {
      // Even-sized matrices don't have these sorts of diagonals
      return 0;
    }

  // Account for the single, center cell
  sum += 1;

  corner = 1;

  // Account for the rest of the diagonals
  for ( i = 3; i <= width; i += 2 )
    {
      // Advance to bottom right corner
      corner += i - 1;
      sum += corner;
      // Advance to bottom left corner
      corner += i - 1;
      sum += corner;
      // Advance to top left corner
      corner += i - 1;
      sum += corner;
      // Advance to top right corner
      corner += i - 1;
      sum += corner;
    }

  return sum;
}

int main( int argc, char *argv[] )
{
  unsigned int width = 1;
  cout << "Sum of diagonals for a " << width << "x" << width << " matrix: " << sumDiagonals( width ) << endl;
  width = 3;
  cout << "Sum of diagonals for a " << width << "x" << width << " matrix: " << sumDiagonals( width ) << endl;
  width = 5;
  cout << "Sum of diagonals for a " << width << "x" << width << " matrix: " << sumDiagonals( width ) << endl;
  width = 7;
  cout << "Sum of diagonals for a " << width << "x" << width << " matrix: " << sumDiagonals( width ) << endl;
  width = 1001;
  cout << "Sum of diagonals for a " << width << "x" << width << " matrix: " << sumDiagonals( width ) << endl;
}
