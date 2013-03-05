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

unsigned char Masks_10[][2] = {
  { 0, 1 },  // a
  { 1, 1 },  // b
  { 0, 2 },  // ab
  { 2, 1 },  // c
  { 1, 2 },  // bc
  { 0, 3 },  // abc
  { 3, 1 },  // d
  { 2, 2 },  // cd
  { 1, 3 },  // bcd
  { 0, 4 },  // abcd
  { 4, 1 },  // e
  { 3, 2 },  // de
  { 2, 3 },  // cde
  { 1, 4 },  // bcde
  { 0, 5 },  // abcde
  { 5, 1 },  // f
  { 4, 2 },  // ef
  { 3, 3 },  // def
  { 2, 4 },  // cdef
  { 1, 5 },  // bcdef
  { 0, 6 },  // abcdef
  { 6, 1 },  // g
  { 5, 2 },  // fg
  { 4, 3 },  // efg
  { 3, 4 },  // defg
  { 2, 5 },  // cdefg
  { 1, 6 },  // bcdefg
  { 0, 7 },  // abcdefg
  { 7, 1 },  // h
  { 6, 2 },  // gh
  { 5, 3 },  // fgh
  { 4, 4 },  // efgh
  { 3, 5 },  // defgh
  { 2, 6 },  // cdefgh
  { 1, 7 },  // bcdefgh
  { 0, 8 },  // abcdefgh
  { 8, 1 },  // i
  { 7, 2 },  // hi
  { 6, 3 },  // ghi
  { 5, 4 },  // fghi
  { 4, 5 },  // efghi
  { 3, 6 },  // defghi
  { 2, 7 },  // cdefghi
  { 1, 8 },  // bcdefghi
  { 0, 9 },  // abcdefghi
  { 9, 1 },  // j
  { 8, 2 },  // ij
  { 7, 3 },  // hij
  { 6, 4 },  // ghij
  { 5, 5 },  // fghij
  { 4, 6 },  // efghij
  { 3, 7 },  // defghij
  { 2, 8 },  // cdefghij
  { 1, 9 },  // bcdefghij
  { 0, 10 }, // abcdefghij
};

unsigned char Masks_11[][2] = {
  { 0, 1 },  // a
  { 1, 1 },  // b
  { 0, 2 },  // ab
  { 2, 1 },  // c
  { 1, 2 },  // bc
  { 0, 3 },  // abc
  { 3, 1 },  // d
  { 2, 2 },  // cd
  { 1, 3 },  // bcd
  { 0, 4 },  // abcd
  { 4, 1 },  // e
  { 3, 2 },  // de
  { 2, 3 },  // cde
  { 1, 4 },  // bcde
  { 0, 5 },  // abcde
  { 5, 1 },  // f
  { 4, 2 },  // ef
  { 3, 3 },  // def
  { 2, 4 },  // cdef
  { 1, 5 },  // bcdef
  { 0, 6 },  // abcdef
  { 6, 1 },  // g
  { 5, 2 },  // fg
  { 4, 3 },  // efg
  { 3, 4 },  // defg
  { 2, 5 },  // cdefg
  { 1, 6 },  // bcdefg
  { 0, 7 },  // abcdefg
  { 7, 1 },  // h
  { 6, 2 },  // gh
  { 5, 3 },  // fgh
  { 4, 4 },  // efgh
  { 3, 5 },  // defgh
  { 2, 6 },  // cdefgh
  { 1, 7 },  // bcdefgh
  { 0, 8 },  // abcdefgh
  { 8, 1 },  // i
  { 7, 2 },  // hi
  { 6, 3 },  // ghi
  { 5, 4 },  // fghi
  { 4, 5 },  // efghi
  { 3, 6 },  // defghi
  { 2, 7 },  // cdefghi
  { 1, 8 },  // bcdefghi
  { 0, 9 },  // abcdefghi
  { 9, 1 },  // j
  { 8, 2 },  // ij
  { 7, 3 },  // hij
  { 6, 4 },  // ghij
  { 5, 5 },  // fghij
  { 4, 6 },  // efghij
  { 3, 7 },  // defghij
  { 2, 8 },  // cdefghij
  { 1, 9 },  // bcdefghij
  { 0, 10 },  // abcdefghij
  { 10, 1 },  // k
  { 9, 2 },  // jk
  { 8, 3 },  // ijk
  { 7, 4 },  // hijk
  { 6, 5 },  // ghijk
  { 5, 6 },  // fghijk
  { 4, 7 },  // efghijk
  { 3, 8 },  // defghijk
  { 2, 9 },  // cdefghijk
  { 1, 10 },  // bcdefghijk
  { 0, 11 },  // abcdefghijk
};

typedef struct
{
  int startDigits;
  void *mask;
  int maskSize;
  const char *start;
  const char *end;
  unsigned long int expectedCount;
} f_params;

f_params FP[] =
  {
    // 0
    { 0, NULL, 0, "", "",  0 },

    // 1
    { 1, NULL, 0, "", "",  9 },

    // 2
    { 2, NULL, 0, "", "", 20 },

    // 3
    // Number contains one of: { 0, 3, 6, 9 }
    //   and number does not sum to 3
    // Number does NOT contain one of: { 0, 3, 6, 9 }
    //   and number DOES sum to 3
    { 3, NULL, 0, "", "", (3*6*6-18)+(6*4*6-18)+(6*6*4-18) + 6*3*1 },

    // 4
    { 4, Masks_04, sizeof(Masks_04) / 2, "1011", "9998", 2701 },

    // 5
    // if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
    // I.E., All numbers that begin with a 5 and have no other 5's or 0's in them
    { 5, NULL, 0, "", "", (unsigned long int) pow( 8, 4 ) },

    // 6
    { 6, Masks_06, sizeof(Masks_06) / 2, "101111", "999986", 109466 },

    // 7
    { 2, Masks_02, sizeof(Masks_02) / 2, "10", "99", 161022 },

    // 8
    { 4, Masks_04, sizeof(Masks_04) / 2, "1011", "9999", 13068583 },

    // 9
    { 3, Masks_03, sizeof(Masks_03) / 2, "101", "988", 2136960 },

    // 10
#ifdef HARD_WAY
    { 7, Masks_07, sizeof(Masks_07) / 2, "1000000", "9999999", 0 },
#else
    { 0, NULL, 0, "", "", 0 },
#endif

    // 11
    { 7, Masks_07, sizeof(Masks_07) / 2, "1013456", "9989865", 71101800 },   // 3m21
    //    { 11, Masks_11, sizeof(Masks_11) / 2, "10134567912", "99898654321", 71101800 },   // > 16m

    // 12
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 13
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 14
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 15
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 16
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 17
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 18
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },

    // 19
    { 8, Masks_08, sizeof(Masks_08) / 2, "10111112", "99999989", 55121700430 },
  };

//#define LOG

unsigned int AddOneDigit(
  BigInt &x,
  unsigned int xLength,
  unsigned int d_digit,
  unsigned int count
)
{
  unsigned int tempCount = count;
  unsigned int i = 0;
  unsigned int start = 0;
  unsigned int sum = 0;

  // Make room for an extra digit on the end
  x.mulByTen();

  xLength++;

  // Try each possible ending digit unless we
  // have already found a child, in which case
  // we can skip digits that are zero because
  // they would add another child
  if ( count > 0 )
    {
      i = 1;
    }
  for ( ; i<=9; i++ )
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
        sum += AddOneDigit( x, xLength, d_digit, count );
      }
    }
    count = tempCount;
  }

  // Remove that extra digit we put on the end
  x.divByTen();

  return sum;
}

unsigned int TrySequences(
  unsigned int d_digit,
  unsigned int xLength,
  const unsigned char mask[][2],
  unsigned int maskCount,
  const char *min,
  const char *max
)
{
  BigInt x    = min;
  BigInt maxX = max;
  unsigned int i = 0;
  unsigned int sum = 0;

#ifdef LOG
  unsigned int hits[maskCount];
  BigInt iterations = 0;
#endif

  for ( ; x <= maxX; x++ )
  {
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
        sum += AddOneDigit( x, xLength, d_digit, count );
      }
    }
  }

#ifdef LOG
  cout << "  Iterations: " << iterations << endl;
  for ( i=0; i<maskCount; i++ )
  {
    cout << "  { " << (int)mask[i][0] << ", " << (int)mask[i][1] << " },   // " << hits[i] << endl;
  }
#endif

  return sum;
}

int main( int argc, char **argv )
{
  unsigned int d_count = 0;
  unsigned int sum = 0;
  unsigned int d_digit = 0;
  unsigned int low = 1;
  unsigned int high = 19;

  if ( argc > 1 )
    {
      low = atoi( argv[1] );
    }
  if ( argc > 2 )
    {
      high = atoi( argv[2] );
    }

  for ( d_digit=low; d_digit<=high; d_digit++ )
    {
      if ( FP[d_digit].mask == NULL )
	{
	  d_count = FP[d_digit].expectedCount;
	  sum += d_count;
	}
      else
	{
	  d_count = TrySequences( d_digit,
				  FP[d_digit].startDigits,
				  (const unsigned char (*)[2]) FP[d_digit].mask,
				  FP[d_digit].maskSize,
				  FP[d_digit].start,
				  FP[d_digit].end );
	  sum += d_count;
	}
      cout << "F(" << d_digit << ") = " << d_count << "\t\t" << sum << endl;
      assert( d_count == FP[d_digit].expectedCount, "FAIL" );
    }
}

