#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lib.h"

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
    printf( "%ld ", array[i] );
  }
  printf( "\n" );
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
