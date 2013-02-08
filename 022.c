#include <iostream>
#include <string>
#include "bigint.h"

using namespace std;

unsigned int score( string s )
{
  unsigned int i = 0;
  unsigned int sum = 0;

  for ( i=0; s[i] != '\0'; i++ )
  {
    sum += toupper( s[i] ) - 'A' + 1;
  }

  return sum;
}

int main( int argc, char **argv )
{
  string temp = "";
  BigInt sum = 0;
  unsigned int i = 1;

  while ( getline( cin, temp ) )
  {
    cout << temp << endl;
    sum += score( temp ) * i;
    i++;
  }

  cout << "i: " << i << " sum: " << sum << endl;
}
