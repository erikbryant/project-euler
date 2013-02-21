#include <iostream>
#include "bigint.h"

using namespace std;

int main( int argc, char *argv[] )
{
  unsigned int lychrel = 0;
  BigInt candidate = 0;
  BigInt max = 10000;
  BigInt result = 0;

  for ( ; candidate<max; candidate++ )
  {
    unsigned int attempts = 0;
    result = candidate;
    do
    {
      result += result.reverse();
      attempts++;
      if ( attempts > 50 )
      {
        cout << "lychrel : " << candidate << endl;
        lychrel++;
        break;
      }
    } while ( !result.isPalindrome() );
  }

  cout << "Lychrel count = " << lychrel << endl;
}
