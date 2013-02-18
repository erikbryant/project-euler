#include <iostream>
#include <string>
#include <math.h>
#include "bigint.h"

using namespace std;

#define assert( cond, error ) if ( !(cond) ) { cout << "ERROR " << __FILE__ << ":" << __LINE__ << ": " << error << endl; exit(1); }


/*
 *
 * d_digit = 2
 * sum: 1 min: 21
 * sum: 1 max: 89
 *
 * d_digit = 3
 * sum: 1 min: 101
 * sum: 1 max: 988
 *
 * d_digit = 4
 * sum: 1 min: 1011
 * sum: 1 max: 9998
 *
 * d_digit = 5
 * sum: 1 min: 51111
 * sum: 1 max: 59999
 *
 * d_digit = 6
 * sum: 1 min: 101111
 * sum: 1 max: 999986
 *
 * d_digit = 7
 * sum: 1 min: 1011111
 * sum: 1 max: 9999993
 *
 * d_digit = 8
 * sum: 1 min: 10111111
 * sum: 1 max: 99999998
 *
 * d_digit = 9
 * sum: 1 min: 101111111
 * sum: 1 max: 988888888
 *
 * d_digit = 11
 * sum: 1 min: 10134567912
 * sum: 1 max: 99898654321
 *
 * d_digit = 12
 * sum: 1 min: 101111111111
 * sum: 1 max: 999999999990
 *
 * d_digit = 13
 * sum: 1 min: 1011111212174
 * sum: 1 max: 9999998986863
 *
 * d_digit = 14
 * sum: 1 min: 10111111111111
 * sum: 1 max: 99999999999930
 *
 * d_digit = 15
 * sum: 1 min: 101111111111111
 * sum: 1 max: 999999999999980
 *
 * d_digit = 16
 * sum: 1 min: 1011111111111111
 * sum: 1 max: 9999999999999996
 *
 * d_digit = 17
 * sum: 1 min: 10111111111131466
 * sum: 1 max: 99999999999999983
 *
 * d_digit = 18
 * sum: 1 min: 101111111111111111
 * sum: 1 max: 999999999999999980
 *
 * d_digit = 19
 * sum: 1 min: 1011111111113562148
 * sum: 1 max: 9999999999999999981
 *
 */



void AddOneDigit( BigInt x, unsigned int xLength, unsigned int d_digit, unsigned int count, BigInt &sum )
{
  unsigned int tempCount = count;
  unsigned int i = 0;
  unsigned int start = 0;

  // Make room for an extra digit on the end
  x.mulByTen();

  xLength++;

  // Try each possible ending digit
  for ( i=0; i<=9; i++ )
  {
    x[xLength - 1] = i;

    for ( start=0; start<xLength; start++ )
    {
      if ( x.testSliceDivisible( start, xLength - start, d_digit ) )
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

  // First work out the (n-1)-digit sequences
  for ( start=0; start < xLength; start++ )
  {
    for ( len = minLen; start + (len - 1) < xLength; len++ )
    {
      if ( x.testSliceDivisible( start, len, d_digit ) )
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
#if 0
  BigInt x;
  unsigned int d_digit;
  BigInt minX = "10";
  BigInt maxX = "99";
  BigInt sum = 0;

  for ( d_digit=2; d_digit<=19; d_digit++, minX *= 10, minX += 1, maxX *= 10, maxX += 9 )
  {
    if ( d_digit == 10 )
    {
      continue;
    }

    cout << "d_digit = " << d_digit << endl;

    sum = 0;
    for ( x=minX; x<=maxX; x++ )
    {
      TrySequences( x, d_digit, d_digit, sum, 1 );
      if ( sum > 0 )
      {
        cout << "sum: " << sum << " min: " << x << endl;
        break;
      }
    }

    sum = 0;

    for ( x=maxX; x>=minX; x-- )
    {
      TrySequences( x, d_digit, d_digit, sum, 1 );
      if ( sum > 0 )
      {
        cout << "sum: " << sum << " max: " << x << endl;
        break;
      }
    }

    cout << endl;
  }

exit(1);
#else

  BigInt x = 1;
  BigInt minX;
  BigInt maxX;
  BigInt sum = 0;
  unsigned int d_digit = 0;

  d_digit = 1;
  // 1 <= x <= 9
  sum = 9;
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 9, "FAIL" );

  d_digit = 2;
  // 10 <= x <= 99
  //   if ( x[0] % 2 != 0 && x[1] % 2 != 0 ) { sum++; }
  sum += 4 * 5;
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 29, "FAIL" );

  d_digit = 3;
  // 100 <= x <= 999
  minX = "101";
  maxX = "988";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 3, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 389, "FAIL" );

  d_digit = 4;
  // 1000 <= x <= 9999
  minX = "1011";
  maxX = "9998";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 4, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 3090, "FAIL" );

  d_digit = 5;
  // 10,000 <= x <= 99,999
  //    if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
  sum += pow( 8, 4 ); // All numbers that begin with a 5 and have no other 5's or 0's in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 7186, "FAIL" );

  d_digit = 6;
  // 100,000 <= x <= 999,999
  minX = "101111";
  maxX = "999986";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 6, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 116652, "FAIL" );

  d_digit = 7;
  // 1,000,000 <= x <= 9,999,999
  minX = "10";
  maxX = "99";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 2, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 277674, "FAIL" );

  d_digit = 8;
  // 10,000,000 <= x <= 99,999,999
  minX = "1011";
  maxX = "9999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 4, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 13346257, "FAIL" );

  d_digit = 9;
  // 100,000,000 <= x <= 999,999,999
  minX = "101";
  maxX = "988";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 3, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 15483217, "FAIL" );

  d_digit = 10;
  //    if ( !x.containsSequence( 0 ) ) { sum++; }
  BigInt temp = 9;
  sum += temp.power( 10 ); // All of the numbers without the digit zero in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == "3502267618", "FAIL" );

  d_digit = 11;
  minX = "1013456";
  maxX = "9989865";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 7, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == "3573369418", "FAIL" );

exit(1);

  sum = "3573369418";

  d_digit = 12;
  minX = "10111111";
  maxX = "99999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 8, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 13;
  minX = "10111111";
  maxX = "99999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 8, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 14;
  minX = "101111111";
  maxX = "999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 9, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 15;
  minX = "1011111111";
  maxX = "9999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 10, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 16;
  minX = "1011111111";
  maxX = "9999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 10, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 17;
  minX = "1011111111";
  maxX = "9999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 10, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 18;
  minX = "1011111111";
  maxX = "9999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 10, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

  d_digit = 19;
  minX = "1011111111";
  maxX = "9999999999";
  for ( x=minX; x<=maxX; x++ )
  {
    TrySequences( x, 10, d_digit, sum, 1 );
  }
  cout << "F(" << d_digit << ") = " << sum << endl;
//  assert( sum == "", "FAIL" );

#endif
}

