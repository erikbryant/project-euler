#include <algorithm>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "bigint.h"

/************************************************************
 *
 * Big integers
 *
 * Representation: EOS terminated array of digits.
 * Digits are stored in reverse of the printed order
 * to optimize for efficiency when the number grows
 * or shrinks.
 *
 ************************************************************/

#if 0
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ )
#else
#define VALIDATE( obj )
#endif

#define EOS 0x7F

BigInt::BigInt() : bigint(NULL), buffLen(0), dataLen(0), dirty(true), sign(1)
{
  import( 0 );
}

BigInt::BigInt( const int x ) : bigint(NULL), buffLen(0), dataLen(0), dirty(true), sign(1)
{
  import( x );
}

BigInt::BigInt( const char * const s ) : bigint(NULL), buffLen(0), dataLen(0), dirty(true), sign(1)
{
  import( s );
}

BigInt::BigInt( const BigInt &other ) : bigint(NULL), buffLen(0), dataLen(0), dirty(true), sign(1)
{
  VALIDATE( &other );

  extendBuffer( other.length() );
  memcpy( bigint, other.bigint, sizeof(unsigned char) * other.length() + 1 );

  dataLen = other.length();
  dirty = false;
  sign = other.sign;
}

BigInt::~BigInt()
{
  free( bigint );
}

ostream &operator<<( ostream &os, const BigInt &bi )
{
  VALIDATE( &bi );

  int i = bi.length() - 1;

  if ( bi.isNegative() )
  {
    os << "-";
  }

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
    dataLen = rhs.dataLen;
    dirty   = rhs.dirty;
    sign    = rhs.sign;
  }

  return *this;
}

const BigInt &BigInt::operator++( void )
{
  VALIDATE( this );
  this->add( 1 );
  return *this;
}

const BigInt &BigInt::operator++( int )
{
  VALIDATE( this );
  this->add( 1 );
  return *this;
}

const BigInt &BigInt::operator+=( const BigInt &rhs )
{
  VALIDATE( this );
  VALIDATE( &rhs );
  this->add( rhs );
  return *this;
}

const BigInt &BigInt::operator--( void )
{
  VALIDATE( this );
  this->subtract( 1 );
  return *this;
}

const BigInt &BigInt::operator--( int )
{
  VALIDATE( this );
  this->subtract( 1 );
  return *this;
}

const BigInt &BigInt::operator-=( const BigInt &rhs )
{
  VALIDATE( this );
  VALIDATE( &rhs );
  this->subtract( rhs );
  return *this;
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
  if ( isNegative() && !other.isNegative() )
  {
    return -1;
  }

  if ( !isNegative() && other.isNegative() )
  {
    return 1;
  }

  int i        = length();
  int otherLen = other.length();
  int result   = 0;

  if ( i < otherLen )
  {
    result = -1;
  }
  else if ( i > otherLen )
  {
    result = 1;
  } else {
    i--;
    while ( i >= 0 )
    {
      if ( bigint[i] < other.bigint[i] )
      {
        result = -1;
        break;
      }
      if ( bigint[i] > other.bigint[i] )
      {
        result = 1;
        break;
      }
      i--;
    }
  }

  if ( isNegative() && other.isNegative() )
  {
    result *= -1;
  }

  return result;
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

// This is a read-only operator. We don't want people
// trying to change the internal structure.
const char BigInt::operator[]( const int i ) const
{
  unsigned int maxIndex = this->length() - 1;

  if ( i >= 0 && i <= maxIndex )
  {
    return bigint[maxIndex - i];
  } else {
    return 0xFF;
  }
}

void BigInt::addStrings( char *s1, const char * const s2 )
{
  unsigned int i = 0;
  unsigned int tempSum = 0;

  while ( s2[i] != EOS )
  {
    if ( s1[i] == EOS )
    {
      s1[i] = 0;
      s1[i+1] = EOS;
    }
    tempSum = s1[i] + s2[i];
    while ( tempSum >= 10 )
    {
      // Do the 'carry' math
      char carry[] = { 1, EOS };
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

  if ( other.isZero() )
  {
    return;
  }

  if ( isPositive() && other.isPositive() )
  {
    // this->add(other)

    // For addition, the sum is at most one digit
    // longer than the longer of the two values.
    extendBuffer( std::max( length(), other.length() ) + 1 );
    addStrings( bigint, other.bigint );
    dirty = true;
  }
  else if ( isPositive() && other.isNegative() )
  {
    // this->subtract(other)
    other.sign = 1;
    this->subtract(other);
    other.sign = -1;
  }
  else if ( isNegative() && other.isPositive() )
  {
    // other->subtract(this)
    this->sign = 1;
    *this = other - *this;
  }
  else if ( isNegative() && other.isNegative() )
  {
    // this->add(other), sign = -1
    this->sign = 1;
    other.sign = 1;
    this->add( other );
    other.sign = -1;
    this->sign = isZero() ? 1 : -1;
  }

  VALIDATE( this );
  VALIDATE( &other );
}

//
// PREREQUISITE:
// s1 >= s2
//
void BigInt::subtractStrings( char *s1, const char * const s2 )
{
  unsigned int i = 0;

  while ( s2[i] != EOS )
  {
    s1[i] -= s2[i];
    i++;
  }

  i = 0;
  while ( s1[i] != EOS )
  {
    // Normalize the negative values
    if ( s1[i] < 0 )
    {
      if ( s1[i+1] == EOS )
      {
        // ERROR: we should never get here
        cout << "Internal error in subtractStrings. Found a negative cell at end of string: " << (int) s1[i] << endl;
      } else {
        s1[i+1]--;
        s1[i] += 10;
      }
    }
    i++;
  }

  // That process may have left leading zeroes. Remove those.
  i--;
  while ( i > 0 )
  {
    if ( s1[i] == 0 )
    {
      s1[i] = EOS;
    } else {
      break;
    }
    i--;
  }
}

void BigInt::subtract( const BigInt &other )
{
  VALIDATE( this );
  VALIDATE( &other );

  if ( isPositive() && other.isPositive() )
  {
    // this->subtract(other)

    // For subtraction, the result can never be longer
    // than either of the two values.
    extendBuffer( std::max( length(), other.length() ) );

    if ( *this >= other )
    {
      subtractStrings( bigint, other.bigint );
    } else {
      *this = other - *this;
      this->sign = isZero() ? 1 : -1;
    }
    dirty = true;
  }
  else if ( isPositive() && other.isNegative() )
  {
    // this->add(other)
    other.sign = 1;
    this->add( other );
    other.sign = -1;
  }
  else if ( isNegative() && other.isPositive() )
  {
    // this->add(other), this->sign = -1
    this->sign = 1;
    this->add( other );
    this->sign = isZero() ? 1 : -1;
  }
  else if ( isNegative() && other.isNegative() )
  {
    // this->add(other)
    other.sign = 1;
    this->add( other );
    other.sign = -1;
  }

  VALIDATE( this );
  VALIDATE( &other );
}

void BigInt::mulOneDigit( char *s1, const char digit )
{
  if ( digit == 0 )
  {
    s1[0] = 0;
    s1[1] = EOS;
    return;
  }

  if ( digit == 1 ) {
    return;
  }

  unsigned int i = 0;
  unsigned int carry = 0;

  while ( s1[i] != EOS )
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
    s1[i+1] = EOS;
  }
}

void BigInt::mul( const BigInt &other )
{
  VALIDATE( this );
  VALIDATE( &other );

  // Short-circuit on identity multiplication.
  if ( isZero() || other.isOne() )
  {
    return;
  }

  if ( isOne() || other.isZero() )
  {
    *this = other;
    return;
  }

  char resultSign = ( sign == other.sign ) ? 1 : -1;

  // Short-circuit on a factor of 10, 100, 1000, etc.
  if ( other.isPowerOfTen() )
  {
    unsigned int zeroCount = other.length() - 1;

    extendBuffer( zeroCount + length() );

    memcpy( bigint + zeroCount, bigint, sizeof( unsigned char) * (length() + 1) );
    memset( bigint, 0, zeroCount );
    if ( !dirty )
    {
      dataLen += zeroCount;
    }

    sign = resultSign;

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
  while ( other.bigint[i] != EOS )
  {
    BigInt temp = 0;
    temp.extendBuffer( length() + other.length() );

    // temp = bigint * 10**i
    unsigned int j = 0;
    while ( j < i ) {
      temp.bigint[j] = 0;
      j++;
    }
    temp.bigint[j] = EOS;
    memcpy( temp.bigint+j, bigint, sizeof( unsigned char ) * (length() + 1) );
    temp.dirty = true;

    // temp *= other.bigint[i]
    mulOneDigit( temp.bigint, other.bigint[i] );

    // Sum the terms
    accumulator += temp;
    i++;
  }

  *this = accumulator;
  sign = resultSign;
  dirty = true;

  VALIDATE( this );
}

bool BigInt::isNegative( void ) const
{
  if ( isZero() && sign != 1 )
  {
    sign = 1;
  }

  return sign == -1;
}

bool BigInt::isPositive( void ) const
{
  return !isNegative();
}

bool BigInt::isZero( void ) const
{
  VALIDATE( this );
  return ( bigint[0] == 0 && bigint[1] == EOS );
}

bool BigInt::isOne( void ) const
{
  VALIDATE( this );
  return ( bigint[0] == 1 && bigint[1] == EOS && sign == 1 );
}

//
// It is a power of ten if it is a series
// of 1 or more zeroes follwed by a 1 and
// then a terminator.
//
bool BigInt::isPowerOfTen( void ) const
{
  unsigned int i = 0;

  while ( bigint[i] == 0 )
  {
    i++;
  }

  return ( i >= 1 && bigint[i] == 1 && bigint[i+1] == EOS );
}

bool BigInt::isDivisibleBy( int divisor ) const
{
  switch (divisor)
  {
    case 0:
      // Nothing is divisible by 0
      return false;
      break;
    case 1:
      // Everything is divisible by 1
      return true;
      break;
    case 2:
      // If the last digit is divisible by 2 then the entire number is
      return ( bigint[0] % 2 == 0 );
      break;
    case 3:
      // If the sum of the digits of a number are divisible by 3 then the entire number is
      if ( this->length() == 1 )
      {
        return ( *this == 3 || *this == 6 || *this == 9 );
      } else {
        return ( this->sumDigits().isDivisibleBy( 3 ) );
      }
      break;
    case 4:
      // If the last two digits are divisible by 4 then the entire number is
      if ( this->length() >= 2 )
      {
        unsigned int lastTwo = bigint[0] + bigint[1] * 10;
        return ( lastTwo % 4 == 0 );
      } else {
        return ( bigint[0] == 4 || bigint[0] == 8 );
      }
      break;
    case 5:
      // Numbers ending in 0 or 5 are divisible by 5
      return ( bigint[0] == 0 || bigint[0] == 5 );
      break;
    case 6:
      // Numbers divisible by 2 AND by 3 are divisble by 6
      return (this->isDivisibleBy( 2 ) && this->isDivisibleBy( 3 ) );
      break;
    case 7:
      // 7 is hard!
      if ( *this <= 189 )
      {
        return ( *this ==   7 || *this ==  14 || *this ==  21 || *this ==  28 ||
                 *this ==  35 || *this ==  42 || *this ==  49 || *this ==  56 ||
                 *this ==  63 || *this ==  70 || *this ==  77 || *this ==  84 ||
                 *this ==  91 || *this ==  98 || *this == 105 || *this == 112 ||
                 *this == 119 || *this == 126 || *this == 133 || *this == 140 ||
                 *this == 147 || *this == 154 || *this == 161 || *this == 168 ||
                 *this == 175 || *this == 182 || *this == 189 );
      } else {
        // Remove the trailing digit from the candidate.
        // Double it and subtract it from the remaining
        // digits. Example:
        //    773 --> 77 - (3 * 2) --> 71
        BigInt temp = *this;
        temp.sign = 1;
        memcpy( temp.bigint, (temp.bigint)+1, temp.length() );
        temp.dataLen--;
        VALIDATE( &temp );
        temp -= bigint[0] * 2;
        return ( temp.isDivisibleBy( 7 ) );
      }
      break;
    case 8:
      // If the last 3 digits are divisible by 8 then the entire number is
      if ( this->length() == 1 )
      {
        return ( bigint[0] == 8 );
      }
      else if ( this->length() == 2 )
      {
        unsigned int lastTwo = bigint[0] + bigint[1] * 10;
        return ( lastTwo % 8 == 0 );
      } else {
        unsigned int lastThree = bigint[0] + bigint[1] * 10 + bigint[2] * 100;
        return ( lastThree % 8 == 0 );
      }
      break;
    case 9:
      // If the sum of the digits of a number are divisible by 9 then the entire number is
      if ( this->length() == 1 )
      {
        return ( *this == 9 );
      } else {
        return ( this->sumDigits().isDivisibleBy( 9 ) );
      }
      break;
    default:
      cout << "ERROR: number out of range for isDivisibleBy(). Expected 0..9, got: " << divisor << endl;
      return false;
  }

  // We shouldn't get here. If we did, it is an error.
  return false;
}

bool BigInt::containsSequence( const BigInt &sequence ) const
{
  if ( this->length() < sequence.length() )
  {
    return false;
  }
  else if ( this->length() == sequence.length() )
  {
    return ( *this == sequence );
  }

  unsigned int i = 0;

  while ( i <= this->length() - sequence.length() )
  {
    if ( memcmp( &(this->bigint[i]), sequence.bigint, sequence.length() ) == 0 )
    {
      return true;
    }
    i++;
  }

  return false;
}

unsigned int BigInt::countSequence( const BigInt &sequence ) const
{
  if ( this->length() < sequence.length() )
  {
    return 0;
  }
  else if ( this->length() == sequence.length() )
  {
    return ( *this == sequence ) ? 1 : 0;
  }

  unsigned int i = 0;
  unsigned int count = 0;

  while ( i <= this->length() - sequence.length() )
  {
    if ( memcmp( &(this->bigint[i]), sequence.bigint, sequence.length() ) == 0 )
    {
      count++;
    }
    i++;
  }

  return count;
}

const BigInt BigInt::power( BigInt const &exponent ) const
{
  BigInt result = *this;

  // Short-circuit identities
  if ( exponent.isZero() )
  {
    result = 1;
    return result;
  }
  if ( exponent.isOne() )
  {
    return result;
  }
  if ( exponent.isNegative() )
  {
    cout << endl << "Negative exponentiation is not supported." << endl;
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
  extendBuffer( 20 );

  if ( x == 0 )
  {
    bigint[0] = 0;
    bigint[1] = EOS;
    dataLen = 1;
    dirty = false;
    sign = 1;
    return;
  }

  int value = x;
  unsigned int i = 0;

  if ( value < 0 )
  {
    sign = -1;
    value *= -1;
  } else {
    sign = 1;
  }

  while ( value > 0 )
  {
    unsigned int mod = value % 10;
    bigint[i] = mod;
    value = value / 10;
    i++;
  }
  bigint[i] = EOS;

  dataLen = i;
  dirty = false;

  VALIDATE( this );
}

void BigInt::import( const char * const s )
{
  unsigned int length = strlen( s );

  extendBuffer( length );

  const char *sptr = s + (length - 1);
  const char *head = s;
  unsigned int i = 0;

  if ( *head == '-' )
  {
    sign = -1;
    head++;
    length--;
  } else {
    sign = 1;
  }

  while ( sptr >= head )
  {
    bigint[i++] = *sptr - '0';
    sptr--;
  }
  bigint[i] = EOS;

  dataLen = length;
  dirty = false;

  VALIDATE( this );
}

void BigInt::extendBuffer( unsigned int length )
{
  // Account for the terminator character
  length += 1;

  // Extending can be expensive. Do it in large blocks
  //  so we don't need to do this often.
  unsigned int blocksize = 512;
  length = blocksize * ( (length + (blocksize - 1)) / blocksize );

  if ( length > buffLen )
  {
    char *old_bigint = bigint;

    buffLen = length;
    bigint = (char *) realloc( bigint, buffLen * sizeof(char) );

    if ( bigint == NULL )
    {
      cout << "ERROR: Failed to allocate " << buffLen << " bytes of RAM." << endl;
      bigint = old_bigint;
    }
  }
}

unsigned int BigInt::length( void ) const
{
  if ( dirty )
  {
    dataLen = 0;
    while ( bigint[dataLen] != EOS )
    {
      dataLen++;
    }
    dirty = false;
  }

  return dataLen;
}

const BigInt BigInt::sumDigits( void ) const
{
  unsigned int i = 0;
  BigInt sum = 0;

  while ( bigint[i] != EOS )
  {
    sum += bigint[i];
    i++;
  }

  return sum;
}

//
// !!!!!WARNING!!!!!
//
// Don't call any other functions from validate().
// All other functions call validate() and that
// would trigger infinite recursion.
//
bool BigInt::validate( const char *file, const int line ) const
{
  bool valid = true;
  unsigned int i = 0;
  unsigned char found = 0;

  // Verify there is a termination character.
  // Verify each digit 0 <= digit <= 9.
  i = 0;
  found = 0;
  while ( i < buffLen )
  {
    if ( bigint[i] == EOS )
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

  if ( !dirty && length != dataLen )
  {
    cout << endl << "ERROR " << file << ":" << line << ": internal length cache is incorrect. Should be: " << length << " is: " << dataLen << endl;
    valid = false;
  }

  // Verify there are no [extraneous] leading zeroes.
  if ( length > 1 )
  {
    if ( bigint[length - 1] == 0 )
    {
      cout << endl << "ERROR " << file << ":" << line << ": bigint has extra leading zeroes." << endl;
      valid = false;
    }
  }

  // Verify there is at least one digit (even if it is just zero).
  if ( length < 1 )
  {
    cout << endl << "ERROR " << file << ":" << line << ": bigint has no digits." << endl;
    valid = false;
  }

  // Verify the sign bit...remember not to call any other functions!
  if ( sign != 1 && sign != -1 )
  {
    cout << endl << "ERROR " << file << ":" << line << ": bigint has invalid sign: " << (int) sign << endl;
    valid = false;
  }

  // Zero can't be negative.
  if ( bigint[0] == 0 && bigint[1] == EOS && sign != 1 )
  {
    cout << endl << "ERROR " << file << ":" << line << ": bigint is zero, but has negative sign: " << (int) sign << endl;
    valid = false;
  }

  return valid;
}

const BigInt operator+( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  BigInt result = lhs;
  result.add( rhs );
  return result;
}

const BigInt operator-( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  BigInt result = lhs;
  result.subtract( rhs );
  return result;
}

const BigInt operator*( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  BigInt result = lhs;
  result.mul( rhs );
  return result;
}

