#include <iostream>
#include <string>
#include <math.h>
#include "bigint.h"

using namespace std;

void AddOneDigit( BigInt x, unsigned int d_digit, unsigned int count, BigInt &sum )
{
  unsigned int tempCount = count;
  unsigned int i = 0;
  unsigned int start = 0;
  BigInt slice = 0;

  // Make room for an extra digit on the end
  x *= 10;

  for ( i=0; i<=9; i++ )
  {
    for ( start=0; start<d_digit; start++ )
    {
      x.slice( start, d_digit - start, slice );
      if ( slice.isDivisibleBy( d_digit ) )
      {
        count++;
        if ( count > 1 ) { break; }
      }
    }
    if ( count == 1 ) { sum++; }
    count = tempCount;
    // Update the digit on the end
    x += 1;
  }
}

void TrySequences( BigInt x, unsigned int d_digit, BigInt &sum )
{
  unsigned int count = 0;
  unsigned int start = 0;
  unsigned int len = 0;
  BigInt slice = 0;

  // First work out the (n-1)-digit sequences
  while ( start < (d_digit - 1) )
  {
    len = 1;
    while ( start + (len - 1) < (d_digit - 1) )
    {
      x.slice( start, len, slice );
      if ( slice.isDivisibleBy( d_digit ) )
      {
        count++;
        if ( count > 1 ) { break; }
      }
      len++;
    }
    if ( count > 1 ) { break; }
    start++;
  }

  // If still a candidate, try n-digit sequences, too.
  if ( count <= 1 )
  {
    AddOneDigit( x, d_digit, count, sum );
  }
}

int main( int argc, char **argv )
{
  BigInt x = 1;
  BigInt sum = 0;
  BigInt max = "10";
  max = max.power( 1 );
  unsigned int d_digit = 0;

  d_digit = 1;
  // 1 <= x <= 9
  sum += 9;
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 2;
  // 10 <= x <= 99
  //   if ( x[0] % 2 != 0 && x[1] % 2 != 0 ) { sum++; }
  sum += 4 * 5;
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 3;
  // 100 <= x <= 999
  for ( x=10; x<=99; x++ )
  {
    TrySequences( x, d_digit, sum );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 4;
  // 1000 <= x <= 9999
  for ( x=100; x<=999; x++ )
  {
    TrySequences( x, d_digit, sum );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 5;
  // 10,000 <= x <= 99,999
  //    if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
  sum += pow( 8, 4 ); // All numbers that begin with a 5 and have no other 5's or 0's in them
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 6;
  // 100,000 <= x <= 999,999
  for ( x=10000; x<=99999; x++ )
  {
    if ( x.countSequence( d_digit ) <= 1 )
    {
      TrySequences( x, d_digit, sum );
    }
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 7;
  // 1,000,000 <= x <= 9,999,999
  for ( x=100000; x<=999999; x++ )
  {
    if ( x.countSequence( d_digit ) <= 1 )
    {
      TrySequences( x, d_digit, sum );
    }
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 8;
  // 10,000,000 <= x <= 99,999,999
  for ( x=1000000; x<=9999999; x++ )
  {
    if ( x.countSequence( d_digit ) <= 1 )
    {
      TrySequences( x, d_digit, sum );
    }
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 9;
  // 100,000,000 <= x <= 999,999,999
  for ( x=10000000; x<=99999999; x++ )
  {
    if ( x.countSequence( d_digit ) <= 1 )
    {
      TrySequences( x, d_digit, sum );
    }
  }
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 10;
  // 1,000,000,000 <= x <= 9,999,999,999
  //    if ( !x.containsSequence( 0 ) ) { sum++; }
  BigInt temp = 9;
  sum += temp.power( 10 ); // All of the numbers without the digit zero in them
  cout << "F(" << d_digit << ") = " << sum << endl;
}






