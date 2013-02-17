#include <iostream>
#include <string>
#include <math.h>
#include "bigint.h"

using namespace std;

void AddOneDigit( BigInt x, unsigned int xLength, unsigned int d_digit, unsigned int count, BigInt &sum )
{
  unsigned int tempCount = count;
  unsigned int i = 0;
  unsigned int start = 0;
  BigInt slice = 0;

  // Make room for an extra digit on the end
  x.mulByTen();

  xLength++;

  // Try each possible ending digit
  for ( i=0; i<=9; i++ )
  {
    x[xLength - 1] = i;

    for ( start=0; start<xLength; start++ )
    {
      x.slice( start, xLength - start, slice );
      if ( slice.isDivisibleBy( d_digit ) )
      {
        count++;
        if ( count > 1 ) { break; }
      }
    }
    if ( xLength == d_digit )
    {
      if ( count == 1 )
      {
        sum++;
      }
    } else {
      if ( count <= 1 )
      {
        AddOneDigit( x, xLength, d_digit, count, sum );
      }
    }
    count = tempCount;
  }
}

void TrySequences( BigInt x, unsigned int xLength, unsigned int d_digit, BigInt &sum, unsigned int minLen )
{
  if ( x.containsMultiple( d_digit, 0 ) ) { return; }

  unsigned int count = 0;
  unsigned int start = 0;
  unsigned int len = 0;
  BigInt slice = 0;

  // First work out the (n-1)-digit sequences
  for ( start=0; start < xLength; start++ )
  {
    for ( len = minLen; start + (len - 1) < xLength; len++ )
    {
      x.slice( start, len, slice );
      if ( slice.isDivisibleBy( d_digit ) )
      {
        count++;
        if ( count > 1 ) { break; }
      }
    }
    if ( count > 1 ) { break; }
  }

  // If still a candidate, try n-digit sequences, too.
  if ( xLength == d_digit )
  {
    if ( count == 1 )
    {
      sum++;
    }
  } else {
    if ( count <= 1 )
    {
      AddOneDigit( x, xLength, d_digit, count, sum );
    }
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
  sum = 9;
  cout << "F(" << d_digit << ") = " << sum << endl;

  d_digit = 2;
  // 10 <= x <= 99
  //   if ( x[0] % 2 != 0 && x[1] % 2 != 0 ) { sum++; }
  sum += 4 * 5;
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 29 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 3;
  // 100 <= x <= 999
  for ( x=10; x<=98; x++ )
  {
    TrySequences( x, 2, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 389 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 4;
  // 1000 <= x <= 9999
  for ( x=101; x<=999; x++ )
  {
    TrySequences( x, 3, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 3090 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 5;
  // 10,000 <= x <= 99,999
  //    if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
  sum += pow( 8, 4 ); // All numbers that begin with a 5 and have no other 5's or 0's in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 7186 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 6;
  // 100,000 <= x <= 999,999
  for ( x=1011; x<=9999; x++ )
  {
    TrySequences( x, 4, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 116652 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 7;
  // 1,000,000 <= x <= 9,999,999
  for ( x=101111; x<=999999; x++ )
  {
    TrySequences( x, 6, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 277674 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 8;
  // 10,000,000 <= x <= 99,999,999
  for ( x=1011111; x<=9999999; x++ )
  {
    TrySequences( x, 7, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 13346257 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 9;
  // 100,000,000 <= x <= 999,999,999
  for ( x=10111111; x<=98888888; x++ )
  {
    TrySequences( x, 8, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != 15483217 ) { cout << "FAIL" << endl; exit(1); }

  d_digit = 10;
  // 1,000,000,000 <= x <= 9,999,999,999
  //    if ( !x.containsSequence( 0 ) ) { sum++; }
  BigInt temp = 9;
  sum += temp.power( 10 ); // All of the numbers without the digit zero in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  if ( sum != "3502267618" ) { cout << "FAIL" << endl; exit(1); }

exit(1);
  d_digit = 11;
  // 10,000,000,000 <= x <= 99,999,999,999
  for ( x=10111111; x<=98888888; x++ )
  {
    TrySequences( x, 8, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
}

