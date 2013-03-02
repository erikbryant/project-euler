#include <iostream>
#include <string>
#include <math.h>
#include <stdlib.h>
#include "bigint.h++"

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
 * d_digit = 10
 * There are no 1-child numbers with 10 digits.
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

//#define MASKS

unsigned char Masks_02[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 1, 1 },  // b
};

unsigned char Masks_03[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 2, 1 },  // c
};

unsigned char Masks_04[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 0, 4 },  // abcd
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 1, 3 },  // bcd
  { 2, 1 },  // c
  { 2, 2 },  // cd
  { 3, 1 },  // d
};

unsigned char Masks_05[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 0, 4 },  // abcd
  { 0, 5 },  // abcde
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 1, 3 },  // bcd
  { 1, 4 },  // bcde
  { 2, 1 },  // c
  { 2, 2 },  // cd
  { 2, 3 },  // cde
  { 3, 1 },  // d
  { 3, 2 },  // de
  { 4, 1 },  // e
};

unsigned char Masks_06[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 0, 4 },  // abcd
  { 0, 5 },  // abcde
  { 0, 6 },  // abcdef
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 1, 3 },  // bcd
  { 1, 4 },  // bcde
  { 1, 5 },  // bcdef
  { 2, 1 },  // c
  { 2, 2 },  // cd
  { 2, 3 },  // cde
  { 2, 4 },  // cdef
  { 3, 1 },  // d
  { 3, 2 },  // de
  { 3, 3 },  // def
  { 4, 1 },  // e
  { 4, 2 },  // ef
  { 5, 1 },  // f
};

unsigned char Masks_07[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 0, 4 },  // abcd
  { 0, 5 },  // abcde
  { 0, 6 },  // abcdef
  { 0, 7 },  // abcdefg
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 1, 3 },  // bcd
  { 1, 4 },  // bcde
  { 1, 5 },  // bcdef
  { 1, 6 },  // bcdefg
  { 2, 1 },  // c
  { 2, 2 },  // cd
  { 2, 3 },  // cde
  { 2, 4 },  // cdef
  { 2, 5 },  // cdefg
  { 3, 1 },  // d
  { 3, 2 },  // de
  { 3, 3 },  // def
  { 3, 4 },  // defg
  { 4, 1 },  // e
  { 4, 2 },  // ef
  { 4, 3 },  // efg
  { 5, 1 },  // f
  { 5, 2 },  // fg
  { 6, 1 },  // g
};

unsigned char Masks_08[][2] = {
  { 0, 1 },  // a
  { 0, 2 },  // ab
  { 0, 3 },  // abc
  { 0, 4 },  // abcd
  { 0, 5 },  // abcde
  { 0, 6 },  // abcdef
  { 0, 7 },  // abcdefg
  { 0, 8 },  // abcdefgh
  { 1, 1 },  // b
  { 1, 2 },  // bc
  { 1, 3 },  // bcd
  { 1, 4 },  // bcde
  { 1, 5 },  // bcdef
  { 1, 6 },  // bcdefg
  { 1, 7 },  // bcdefgh
  { 2, 1 },  // c
  { 2, 2 },  // cd
  { 2, 3 },  // cde
  { 2, 4 },  // cdef
  { 2, 5 },  // cdefg
  { 2, 6 },  // cdefgh
  { 3, 1 },  // d
  { 3, 2 },  // de
  { 3, 3 },  // def
  { 3, 4 },  // defg
  { 3, 5 },  // defgh
  { 4, 1 },  // e
  { 4, 2 },  // ef
  { 4, 3 },  // efg
  { 4, 4 },  // efgh
  { 5, 1 },  // f
  { 5, 2 },  // fg
  { 5, 3 },  // fgh
  { 6, 1 },  // g
  { 6, 2 },  // gh
  { 7, 1 },  // h
};

//#define LOG

void AddOneDigit(
  BigInt &x,
  unsigned int xLength,
  unsigned int d_digit,
  unsigned int count,
  BigInt &sum
)
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

  // Remove that extra digit we put on the end
  x.divByTen();
}

void TrySequences(
  unsigned int d_digit,
  unsigned int xLength,
  const unsigned char mask[][2],
  unsigned int maskCount,
  const char *min,
  const char *max,
  BigInt &sum
)
{
  BigInt x    = min;
  BigInt maxX = max;
  unsigned int i = 0;

#ifdef LOG
  unsigned int hits[maskCount];
  BigInt iterations = 0;
#endif

  for ( ; x <= maxX; x++ )
  {
    if ( x.containsMultiple( d_digit, 0 ) ) { continue; }
    unsigned int count = 0;

#ifdef LOG
    iterations++;
#endif

    // First work out the (n-1)-digit sequences
    for ( i=0; i<maskCount; i++ )
    {
      if ( x.testSliceDivisible( mask[i][0], mask[i][1], d_digit ) )
      {
#ifdef LOG
        hits[i]++;
#endif
        count++;
        if ( count > 1 ) { break; }
      }
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

  cout << "F(" << d_digit << ") = " << sum << endl;
#ifdef LOG
  cout << "  Iterations: " << iterations << endl;
  for ( i=0; i<maskCount; i++ )
  {
    cout << "  { " << (int)mask[i][0] << ", " << (int)mask[i][1] << " },   // " << hits[i] << endl;
  }
#endif
}

int main( int argc, char **argv )
{
  BigInt sum = 0;
  unsigned int d_digit = 0;

  d_digit = 1;
  sum = 9;
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 9, "FAIL" );

  d_digit = 2;
  //   if ( x[0] % 2 != 0 && x[1] % 2 != 0 ) { sum++; }
  sum += 4 * 5;
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 29, "FAIL" );

  d_digit = 3;
  //   Number contains one of: { 0, 3, 6, 9 }
  //     and number does not sum to 3
  //   Number does NOT contain one of: { 0, 3, 6, 9 }
  //     and number DOES sum to 3
  sum += (3*6*6-18)+(6*4*6-18)+(6*6*4-18) + 6*3*1;
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 389, "FAIL" );

  d_digit = 4;
  TrySequences( d_digit, 4, Masks_04, sizeof(Masks_04) >> 1, "1011", "9998", sum );
  assert( sum == 3090, "FAIL" );

  d_digit = 5;
  //    if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
  sum += pow( 8, 4 ); // All numbers that begin with a 5 and have no other 5's or 0's in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 7186, "FAIL" );

  d_digit = 6;
  TrySequences( d_digit, 6, Masks_06, sizeof(Masks_06) >> 1, "101111", "999986", sum );
  assert( sum == 116652, "FAIL" );

  d_digit = 7;
  TrySequences( d_digit, 2, Masks_02, sizeof(Masks_02) >> 1, "10", "99", sum );
  assert( sum == 277674, "FAIL" );

  d_digit = 8;
  TrySequences( d_digit, 4, Masks_04, sizeof(Masks_04) >> 1, "1011", "9999", sum );
  assert( sum == 13346257, "FAIL" );

  d_digit = 9;
  TrySequences( d_digit, 3, Masks_03, sizeof(Masks_03) >> 1, "101", "988", sum );
  assert( sum == 15483217, "FAIL" );

  d_digit = 10;
  //    if ( !x.containsSequence( 0 ) ) { sum++; }
  BigInt temp = 9;
  sum += 0; // All of the numbers without the digit zero in them
  cout << "F(" << d_digit << ") = " << sum << endl;
  assert( sum == 15483217, "FAIL" );

  d_digit = 11;
  TrySequences( d_digit, 7, Masks_07, sizeof(Masks_07) >> 1, "1013456", "9989865", sum );
  assert( sum == "3573369418", "FAIL" );

  sum = "3573369418";

#if 0
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

