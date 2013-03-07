//
// Copyright Erik Bryant (erikbryantology@gmail.com)
// GPLv2 http://www.gnu.org/licenses/gpl-2.0.html
//

#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lib.h++"

using namespace std;

unsigned int power( unsigned int base, unsigned int exponent )
{
  unsigned int result = 1;

  while ( exponent > 0 )
  {
    result *= base;
    exponent--;
  }

  return result;
}

void printArray( unsigned int *array, unsigned int count )
{
  unsigned int i = 0;

  for ( i=0; i<count; i++ )
  {
    cout << array[i];
  }
  cout << endl;
}

unsigned int sumArray( unsigned int *array, unsigned int count )
{
  unsigned int i = 0;
  unsigned int sum = 0;

  for ( i=0; i<count; i++ )
  {
    sum += array[i];
  }

  return sum;
}

char arrayContains( unsigned int n, unsigned int *array, unsigned int count )
{
  unsigned int i = 0;

  for ( i=0; i<count; i++ )
  {
    if ( n == array[i] )
    {
      return 1;
    }
  }

  return 0;
}

unsigned int findDivisors( unsigned int n, unsigned int *divisors, char proper )
{
  unsigned int divisorsLen = 0;
  unsigned int i = 0;

  // Everything is divisble by 1
  divisors[divisorsLen++] = 1;

  for ( i=2; i<=n/2; i++ )
  {
    if ( n % i == 0 )
    {
      divisors[divisorsLen++] = i;
    }
  }

  if ( !proper && n != 1 )
  {
    divisors[divisorsLen++] = n;
  }

  return divisorsLen;
}

unsigned int sumOfDivisors( unsigned int n )
{
  unsigned int divisors[2000];
  unsigned int divisorsLen = 0;

  divisorsLen = findDivisors( n, divisors, 1 );

  return sumArray( divisors, divisorsLen );
}

unsigned int findAbundants( unsigned int max, unsigned int *array )
{
  unsigned int count = 0;
  unsigned int i=0;

  for ( i=1; i<=max; i++ )
  {
    if ( i < sumOfDivisors( i ) )
    {
      array[count++] = i;
    }
  }

  return count;
}

void englishify( unsigned int x )
{
  const char *english[] = {
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
    "seventeen", "eighteen", "nineteen"
  };

  if ( x > 1000 )
  {
    cout << "value is > 1000 (" << x << ")";
    return;
  }

  if ( x == 1000 )
  {
    cout << "one thousand";
    return;
  }

  if ( x > 100 )
  {
    unsigned int hundreds = x / 100;
    unsigned int tens = x - (hundreds * 100);

    englishify( hundreds );
    cout << " hundred";
    if ( tens != 0 )
    {
      cout << " and ";
      englishify( tens );
    }

    return;
  }

  if ( x == 100 )
  {
    cout << "one hundred";
    return;
  }

  if ( x >= 20 )
  {
    if ( x >= 90 )
    {
      cout << "ninety";
      x -= 90;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 80 )
    {
      cout << "eighty";
      x -= 80;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 70 )
    {
      cout << "seventy";
      x -= 70;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 60 )
    {
      cout << "sixty";
      x -= 60;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 50 )
    {
      cout << "fifty";
      x -= 50;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 40 )
    {
      cout << "forty";
      x -= 40;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 30 )
    {
      cout << "thirty";
      x -= 30;
      if ( x > 0 ) { cout << "-"; }
    }
    if ( x >= 20 )
    {
      cout << "twenty";
      x -= 20;
      if ( x > 0 ) { cout << "-"; }
    }

    if ( x > 0 )
    {
      englishify( x );
    }

    return;
  }

  cout << english[x];

  return;
}
