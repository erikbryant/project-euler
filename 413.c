#include <stdlib.h>
#include <stdio.h>

#define assert( cond, error ) if ( !(cond) ) { printf( "ERROR %s:%d: %s\n", __FILE__, __LINE__, error ); exit(1); }

//
// TODO:
//
// Zero optimization:
//   There can be no 1-child numbers that have two or more zeroes in them.
//     * Enumerate all numbers with one zero in them. Test if they are 1-child.
//     * Enumerate all numbers with no zeroes in them. Test.
//

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

typedef struct
{
  const char *start;
  const char *end;
  unsigned long int expectedCount;
} f_params;

f_params FP[] =
  {
    // 0
    { NULL, NULL,  0 },

    // 1
    { NULL, NULL,  9 },

    // 2
    { NULL, NULL, 20 },

    // 3
    // Number contains one of: { 0, 3, 6, 9 }
    //   and number does not sum to 3
    // Number does NOT contain one of: { 0, 3, 6, 9 }
    //   and number DOES sum to 3
    { NULL, NULL, (3*6*6-18)+(6*4*6-18)+(6*6*4-18) + 6*3*1 },

    // 4
    { "1011", "9998", 2701 },

    // 5
    // if ( !x.containsSequence( 0 ) && x[0] == 5 && x.countSequence( 5 ) == 1 ) { sum++; }
    // I.E., All numbers that begin with a 5 and have no other 5's or 0's in them
    { NULL, NULL, 8 * 8 * 8 * 8 },

    // 6
    { "101111", "999986", 109466 },

    // 7
    { "10", "99", 161022 },

    // 8
    { "1011", "9999", 13068583 },

    // 9
    { "101", "988", 2136960 },

    // 10
    { NULL, NULL, 0 },

    // 11
    { "1013456", "9989865", 71101800 },   // 3m21

    // 12
    { "1011111", "9999998", 3582069103 },

    // 13
    { "10111112", "99999989", 55121700430 },

    // 14
    { "10111112", "99999989", 55121700430 },

    // 15
    { "10111112", "99999989", 55121700430 },

    // 16
    { "10111112", "99999989", 55121700430 },

    // 17
    { "10111112", "99999989", 55121700430 },

    // 18
    { "10111112", "99999989", 55121700430 },

    // 19
    { "10111112", "99999989", 55121700430 },
  };

char x[30];
unsigned int xLength = 0;
unsigned int d_digit = 0;
unsigned int count = 0;

unsigned int AddOneDigit( void )
{
  // Should we really be in here? Check the
  // termination conditions to make sure.
  if ( xLength == d_digit )
    {
      return count == 1 ? 1 : 0;
    }

  // Make room for an extra digit on the end
  xLength++;

  unsigned int initialCount = count;
  unsigned int i = 0;
  int start = 0;
  unsigned int sum = 0;

  // Try each [0-9] ending digit unless we have
  // already found a child, in which case we can
  // skip zero because it would add another child
  for ( i = (count>0); i <= 9; ++i )
  {
    count = initialCount;
    x[xLength - 1] = i;

    unsigned long int value = 0;
    unsigned long int power = 1;
    for ( start = xLength-1; start >= 0; --start, power*=10 )
    {
      value += x[start] * power;
      if ( value % d_digit == 0 )
      {
        count++;
        if ( count > 1 ) { break; }
      }
    }
    if ( count <= 1 )
      {
        sum += AddOneDigit();
      }
  }

  // Remove that extra digit we put on the end
  xLength--;

  return sum;
}

int main( int argc, char **argv )
{
  unsigned int d_count = 0;
  unsigned int sum = 0;
  unsigned int d_min = 1;
  unsigned int d_max = 19;

  if ( argc > 1 )
    {
      d_min = atoi( argv[1] );
    }
  if ( argc > 2 )
    {
      d_max = atoi( argv[2] );
    }

  for ( d_digit = d_min; d_digit <= d_max; d_digit++ )
    {
      if ( FP[d_digit].start == NULL )
	{
	  d_count = FP[d_digit].expectedCount;
	}
      else
	{
	  d_count = 0;
	  unsigned int i = 0;
	  for ( i = 1; i <= 9; ++i )
	    {
	      x[0] = i;
	      xLength = 1;
	      count = i % d_digit == 0;
	      d_count += AddOneDigit();
	    }
	}
      sum += d_count;
      printf( "F(%d) = %d\t\t%d\n", d_digit, d_count, sum );
      assert( d_count == FP[d_digit].expectedCount, "FAIL" );
    }

  return 0;
}
