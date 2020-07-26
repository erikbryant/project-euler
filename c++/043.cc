#include <iostream>
#include "bigint.h"

using namespace std;

int main( int argc, char *argv[] )
{
  BigInt i;
  BigInt min = "1023456789";
  BigInt max = "9876543210";
  BigInt sum = 0;

  for ( i = min; i <= max; ++i )
  {
    if ( i.isPandigital( 0, 9 ) && 
         i.testSliceDivisible( 1, 3,  2 ) &&
         i.testSliceDivisible( 2, 3,  3 ) &&
         i.testSliceDivisible( 3, 3,  5 ) &&
         i.testSliceDivisible( 4, 3,  7 ) &&
         i.testSliceDivisible( 5, 3, 11 ) &&
         i.testSliceDivisible( 6, 3, 13 ) &&
         i.testSliceDivisible( 7, 3, 17 ) )
    {
      cout << "Found:  " << i << endl;
      sum += i;
    }
  }

  cout << "Sum = " << sum << endl;
}
