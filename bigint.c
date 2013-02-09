#include <algorithm>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "bigint.h"

/************************************************************
 *
 * Big integers
 *
 * Representation: 0xFF terminated array of digits.
 * Digits are stored in reverse of the printed order
 * to optimize for efficiency when the number grows
 * or shrinks.
 *
 ************************************************************/

#if 1
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ )
#else
#define VALIDATE( obj )
#endif

BigInt::BigInt() : bigint(NULL), buffLen(0)
{
  import( 0 );
}

BigInt::BigInt( const int x ) : bigint(NULL), buffLen(0)
{
  import( x );
}

BigInt::BigInt( const char * const s ) : bigint(NULL), buffLen(0)
{
  import( s );
}

BigInt::BigInt( const BigInt &other ) : bigint(NULL), buffLen(0)
{
  VALIDATE( &other );

  extendBuffer( other.length() );
  memcpy( bigint, other.bigint, sizeof(unsigned char) * other.length() + 1 );
}

BigInt::~BigInt()
{
  free( bigint );
}

ostream &operator<<( ostream &os, const BigInt &bi )
{
  VALIDATE( &bi );

  int i = bi.length() - 1;

  while ( i >= 0 )
  {
    os << (int) bi.bigint[i];
    i--;
  }

  return os;
}

const BigInt &BigInt::operator=( const BigInt &rhs )
{
  VALIDATE( &rhs );

  // Skip the assignment if lhs and rhs are the same object.
  if ( this != &rhs )
  {
    extendBuffer( rhs.length() );
    memcpy( bigint, rhs.bigint, sizeof( unsigned char ) * (rhs.length() + 1) );
  }

  return *this;
}

const BigInt BigInt::operator+( const BigInt &other ) const
{
  VALIDATE( this );
  VALIDATE( &other );
  BigInt result = *this;
  result.add( other );
  return result;
}

const BigInt &BigInt::operator++( void )
{
  VALIDATE( this );
  add( 1 );
  return *this;
}

const BigInt &BigInt::operator++( int )
{
  VALIDATE( this );
  add( 1 );
  return *this;
}

const BigInt &BigInt::operator+=( const BigInt &rhs )
{
  VALIDATE( this );
  VALIDATE( &rhs );
  this->add( rhs );
  return *this;
}

const BigInt BigInt::operator*( const BigInt &other ) const
{
  VALIDATE( this );
  VALIDATE( &other );
  BigInt result = *this;
  result.mul( other );
  return result;
}

const BigInt &BigInt::operator*=( const BigInt &rhs )
{
  VALIDATE( this );
  VALIDATE( &rhs );
  this->mul( rhs );
  return *this;
}

int BigInt::compare( const BigInt &other ) const
{
  int i        = length();
  int otherLen = other.length();

  if ( i < otherLen )
  {
    return -1;
  }

  if ( i > otherLen )
  {
    return 1;
  }

  i--;

  while ( i >= 0 )
  {
    if ( bigint[i] < other.bigint[i] )
    {
      return -1;
    }
    if ( bigint[i] > other.bigint[i] )
    {
      return 1;
    }
    i--;
  }

  return 0;
}

bool BigInt::operator==( const BigInt &other ) const
{
  return ( compare( other ) == 0 );
}

bool BigInt::operator!=( const BigInt &other ) const
{
  return ( compare( other ) != 0 );
}

bool BigInt::operator<( const BigInt &other ) const
{
  return ( compare( other ) == -1 );
}

bool BigInt::operator<=( const BigInt &other ) const
{
  return ( compare( other ) != 1 );
}

bool BigInt::operator>( const BigInt &other ) const
{
  return ( compare( other ) == 1 );
}

bool BigInt::operator>=( const BigInt &other ) const
{
  return ( compare( other ) != -1 );
}

void BigInt::addStrings( unsigned char *s1, const unsigned char * const s2 )
{
  unsigned int i = 0;
  unsigned int tempSum = 0;

  while ( s2[i] != 0xFF )
  {
    if ( s1[i] == 0xFF )
    {
      s1[i] = 0;
      s1[i+1] = 0xFF;
    }
    tempSum = s1[i] + s2[i];
    while ( tempSum >= 10 )
    {
      // Do the 'carry' math
      unsigned char carry[] = { 1, 0xFF };
      addStrings( s1+i+1, carry );
      tempSum -= 10;
    }
    s1[i] = tempSum;
    i++;
  }
}

void BigInt::add( const BigInt &other )
{
  VALIDATE( this );
  VALIDATE( &other );

  // For addition, the sum is at most one digit
  // longer than the longer of the two addends.
  extendBuffer( std::max( length(), other.length() ) + 1 );

  addStrings( bigint, other.bigint );
}

void BigInt::mulOneDigit( unsigned char *s1, const unsigned char digit )
{
  if ( digit == 0 )
  {
    s1[0] = 0;
    s1[1] = 0xFF;
    return;
  }

  if ( digit == 1 ) {
    return;
  }

  unsigned int i = 0;
  unsigned int carry = 0;

  while ( s1[i] != 0xFF )
  {
    unsigned int temp = s1[i] * digit + carry;
    carry = 0;
    while ( temp >= 10 )
    {
      carry++;
      temp -= 10;
    }
    s1[i] = temp;
    i++;
  }
  if ( carry > 0 )
  {
    s1[i] = carry;
    s1[i+1] = 0xFF;
  }
}

void BigInt::mul( const BigInt &other )
{
  VALIDATE( this );
  VALIDATE( &other );

  // Short-circuit on identity multiplication.
  if ( length() == 1 )
  {
    if ( bigint[0] == 0 )
    {
      return;
    }
    if ( bigint[1] == 1 )
    {
      *this = other;
      return;
    }
  }

  // Short-circuit on identity multiplication.
  if ( other.length() == 1 )
  {
    if ( other.bigint[0] == 0 )
    {
      bigint[0] = 0;
      bigint[1] = 0xFF;
      return;
    }
    if ( other.bigint[1] == 1 )
    {
      return;
    }
  }

  // Short-circuit on a factor of 10, 100, 1000, etc.
  if ( other.powerOfTen() )
  {
    unsigned int zeroCount = other.length() - 1;

    extendBuffer( zeroCount + length() );

    memcpy( bigint + zeroCount, bigint, sizeof( unsigned char) * (length() + 1) );
    memset( bigint, 0, zeroCount );
    VALIDATE( this );
    return;
  }

  // For multiplication, the length of the product is
  // at most the sum of the lengths of the mutiplicands.
  extendBuffer( length() + other.length() );

  //
  // Long multiplication.
  // Product of bigint and other.bigint is sum of each of the terms:
  //   bigint * 10**i * other.bigint[i]
  //
  // Example:
  //        123
  //       *123
  //       ----
  //        369   123 * 10**0 * 3
  //       2460   123 * 10**1 * 2
  //      12300   123 * 10**2 * 1
  //      -----
  //      15129
  //
  unsigned int i = 0;
  BigInt accumulator = 0;
  while ( other.bigint[i] != 0xFF )
  {
    BigInt temp = 0;
    temp.extendBuffer( length() + other.length() );

    // temp = bigint * 10**i
    unsigned int j = 0;
    while ( j < i ) {
      temp.bigint[j] = 0;
      j++;
    }
    temp.bigint[j] = 0xFF;
    memcpy( temp.bigint+j, bigint, sizeof( unsigned char ) * (length() + 1) );

    // temp *= other.bigint[i]
    mulOneDigit( temp.bigint, other.bigint[i] );

    // Sum the terms
    accumulator += temp;
    i++;
  }

  *this = accumulator;
}

//
// It is a power of ten if it is a series
// of 1 or more zeroes follwed by a 1 and
// then a terminator.
//
bool BigInt::powerOfTen( void ) const
{
  unsigned int i = 0;

  while ( bigint[i] == 0 )
  {
    i++;
  }

  return ( i >= 1 && bigint[i] == 1 && bigint[i+1] == 0xFF );
}

const BigInt BigInt::power( BigInt const &exponent ) const
{
  BigInt result = *this;

  // Short-circuit identities
  if ( exponent.length() == 1 )
  {
    if ( exponent.bigint[0] == 0 )
    {
      result = 1;
      return result;
    }
    if ( exponent.bigint[0] == 1 )
    {
      return result;
    }
  }

  BigInt i      = 1;
  BigInt factor = *this;

  for ( ; i < exponent; i += 1 )
  {
    result *= factor;
  }

  return result;
}

void BigInt::import( const int x )
{
  char temp[20];
  sprintf( temp, "%ld", x );
  import( temp );
  VALIDATE( this );
}

void BigInt::import( const char * const s )
{
  unsigned int length = strlen( s );

  extendBuffer( length );

  const char *sptr = s + (length - 1);
  unsigned int i = 0;

  while ( sptr >= s )
  {
    bigint[i++] = *sptr - '0';
    sptr--;
  }
  bigint[i] = 0xFF;

  VALIDATE( this );
}

void BigInt::extendBuffer( unsigned int length )
{
  // Account for the terminator character
  length += 1;

  // Extending can be expensive. Do it in large blocks
  //  so we don't need to do this often.
  if ( length % 1024 != 0 )
  {
    length = 1024 * ( 1 + length / 1024 );
  }

  if ( length > buffLen )
  {
    buffLen = length;
    bigint = (unsigned char *) realloc( bigint, buffLen * sizeof( unsigned char ) );
  }
}

unsigned int BigInt::length( void ) const
{
  unsigned int i = 0;

  while ( bigint[i] != 0xFF )
  {
    i++;
  }

  return i;
}

const BigInt BigInt::sumDigits( void ) const
{
  unsigned int i = 0;
  BigInt sum = 0;

  while ( bigint[i] != 0xFF )
  {
    sum += bigint[i];
    i++;
  }

  return sum;
}

//
// !!!!!WARNING!!!!!
// Don't call any other functions from validate().
// All other functions call validate() and that
// would trigger infinite recursion.
//
bool BigInt::validate( const char *file, const int line ) const
{
  unsigned int i = 0;
  unsigned char found = 0;

  // Verify there is a termination character.
  // Verify each digit 0 <= digit <= 9.
  i = 0;
  found = 0;
  while ( i < buffLen )
  {
    if ( bigint[i] == 0xFF )
    {
      found = 1;
      break;
    }
    if ( bigint[i] < 0 || bigint[i] > 9 )
    {
      cout << endl << "ERROR " << file << ":" << line << ": bigint has a digit out of range: bigint[" << i << "] == " << (int) bigint[i] << endl;
    }
    i++;
  }
  if ( !found )
  {
    cout << endl << "ERROR " << file << ":" << line << ": bigint is lacking a termination character." << endl;
    return false;
  }

  unsigned int length = i;

  // Verify there are no [extraneous] leading zeroes.
  if ( length > 1 )
  {
    if ( bigint[length - 1] == 0 )
    {
      cout << endl << "ERROR " << file << ":" << line << ": bigint has extra leading zeroes." << endl;
      return false;
    }
  }

  // Verify there is at least one digit (even if it is just zero).
  if ( length < 1 )
  {
    cout << endl << "ERROR " << file << ":" << line << ": bigint has no digits." << endl;
    return false;
  }

  return true;
}

