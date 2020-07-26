//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include "bigint.h"

#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <algorithm>

using std::swap;
using std::cout;
using std::endl;

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

BigInt::BigInt() : bigint(starter), buffLen(STARTER_LEN), dataLen(0), sign(1)
{
  import( 0 );
}

BigInt::BigInt( const int x ) : bigint(starter), buffLen(STARTER_LEN), dataLen(0), sign(1)
{
  import( x );
}

BigInt::BigInt( const char * const s ) : bigint(starter), buffLen(STARTER_LEN), dataLen(0), sign(1)
{
  import( s );
}

BigInt::BigInt( const BigInt &other ) :
  bigint(starter),
  buffLen(STARTER_LEN),
  dataLen(other.length()),
  sign(other.sign)
{
  VALIDATE( &other );

  extendBuffer( other.length() );
  memcpy( bigint, other.bigint, sizeof(char) * other.length() + 1 );

  VALIDATE( this );
}

BigInt::~BigInt()
{
  if ( bigint != starter )
  {
    free( bigint );
  }
}

void BigInt::slice( unsigned int start, unsigned int length, BigInt &other ) const
{
  VALIDATE( this );

  // Remember...the internal storage is in reverse...
  start = ( this->length() - 1 ) - start - ( length - 1 );

  // Remove extraneous leading zeroes
  while ( bigint[start + length - 1] == 0 && length > 1 )
  {
    length--;
  }

  other.extendBuffer( length );
  memcpy( other.bigint, bigint + start, sizeof(char) * length );
  other.bigint[length] = EOS;
  other.dataLen = length;
  other.sign = 1;

  VALIDATE( &other );
}

bool BigInt::testSliceDivisible( unsigned int start, unsigned int length, unsigned int divisor ) const
{
  VALIDATE( this );

  // Remember...the internal storage is in reverse...
  start = this->length() - start - length;

  // Remove extraneous leading zeroes
  while ( bigint[start + length - 1] == 0 && length > 1 )
  {
    length--;
  }

  BigInt slice;
  slice.extendBuffer( length );
  memcpy( slice.bigint, this->bigint + start, length );
  slice.bigint[length] = EOS;
  slice.dataLen = length;
  VALIDATE( &slice );

  return slice.isDivisibleBy( divisor );
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
    memcpy( bigint, rhs.bigint, sizeof(char) * (rhs.length() + 1) );
    dataLen = rhs.dataLen;
    sign    = rhs.sign;
  }

  VALIDATE( this );

  return *this;
}

const BigInt &BigInt::operator++( void )
{
  VALIDATE( this );
  this->addOne();
  return *this;
}

const BigInt BigInt::operator++( int )
{
  VALIDATE( this );
  this->addOne();
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

const BigInt BigInt::operator--( int )
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

const BigInt &BigInt::operator/=( const BigInt &rhs )
{
  VALIDATE( this );
  VALIDATE( &rhs );
  *this = this->div( rhs );
  return *this;
}

int BigInt::compare( const BigInt &other ) const
{
  VALIDATE( &other );
  VALIDATE( this );

  if ( sign == other.sign )
  {
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

  if ( isNegative() && !other.isNegative() )
  {
    return -1;
  }

  // The only other case left is:
  //   ( !isNegative() && other.isNegative() )
  return 1;
}

bool BigInt::isPalindrome( void ) const
{
  VALIDATE( this );

  char *head = bigint;
  char *tail = bigint + length() - 1;

  while ( head < tail )
  {
    if ( *head != *tail )
    {
      return false;
    }
    head++;
    tail--;
  }

  return true;
}

BigInt BigInt::reverse( void ) const
{
  VALIDATE( this );

  BigInt result = *this;
  char *head = result.bigint;
  char *tail = result.bigint + result.length() - 1;

  while ( head < tail )
  {
    swap( *head, *tail );
    head++;
    tail--;
  }

  // Remove any leading zeroes
  tail = result.bigint + result.length() - 1;
  while ( *tail == 0 && length() > 1 )
  {
    *tail = EOS;
    result.dataLen--;
    tail--;
  }

  return result;
}

#if 1
//
// This is better than the alternate implementation below.
// The other uses recursion and if you have a 10,000 digit
// BigInt you could end up recursing that deep. Better to
// loop instead of recursing.
//
unsigned int BigInt::addStrings( char *s1, const char * const s2 )
{
  unsigned int i = 0;
  unsigned int carry = 0;

  while ( s2[i] != EOS )
  {
    if ( s1[i] == EOS )
    {
      s1[i] = 0;
      s1[i+1] = EOS;
    }
    s1[i] = s1[i] + s2[i] + carry;
    carry = 0;
    if ( s1[i] >= 10 )
    {
      carry = 1;
      s1[i] -= 10;
    }
    i++;
  }

  while ( carry > 0 )
  {
    if ( s1[i] == EOS )
    {
      s1[i] = 0;
      s1[i+1] = EOS;
    }
    s1[i] = s1[i] + carry;
    carry = 0;
    if ( s1[i] >= 10 )
    {
      carry = 1;
      s1[i] -= 10;
    }
    i++;
  }

  // Find the length, since we are already so close to the end
  while ( s1[i] != EOS ) { i++; }

  return i;
}
#else
unsigned int BigInt::addStrings( char *s1, const char * const s2 )
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
    if ( tempSum >= 10 )
    {
      // Do the 'carry' math
      char carry[] = { 1, EOS };
      addStrings( s1+i+1, carry );
      tempSum -= 10;
    }
    s1[i] = tempSum;
    i++;
  }

  // Find the length, since we are already so close to the end
  while ( s1[i] != EOS ) { i++; }

  return i;
}
#endif

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
    dataLen = addStrings( bigint, other.bigint );
  }
  else if ( isPositive() && other.isNegative() )
  {
    // this->subtract(abs(other))
    this->subtract( other * -1 );
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
    this->add( other * -1 );
    this->sign = isZero() ? 1 : -1;
  }

  VALIDATE( this );
  VALIDATE( &other );
}

void BigInt::addOne( void )
{
  VALIDATE( this );

  if ( isPositive() )
  {
    // this->add(1)

    // For addition, the sum is at most one digit
    // longer than the longer of the two values.
    extendBuffer( length() + 1 );
    (*bigint)++;
    if ( *bigint >= 10 )
    {
      *bigint -= 10;
      // Do the 'carry' math
      char carry[] = { 1, EOS };
      dataLen = addStrings( bigint + 1, carry ) + 1;
    }
  }
  else
  {
    // 1->subtract(this)
    this->sign = 1;
    *this = 1 - *this;
  }

  VALIDATE( this );
}

//
// PREREQUISITE:
// s1 >= s2
//
unsigned int BigInt::subtractStrings( char *s1, const char * const s2 )
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
        exit( 1 );
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

  // Find the length, since we are already so close to the end
  while ( s1[i] != EOS ) { i++; }

  return i;
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
      dataLen = subtractStrings( bigint, other.bigint );
    } else {
      *this = other - *this;
      this->sign = isZero() ? 1 : -1;
    }
  }
  else if ( other.isNegative() )
  {
    // this->add(abs(other))
    this->add( other * -1 );
  }
  else if ( isNegative() && other.isPositive() )
  {
    // this->add(other), this->sign = -1
    this->sign = 1;
    this->add( other );
    this->sign = isZero() ? 1 : -1;
  }

  VALIDATE( this );
  VALIDATE( &other );
}

unsigned int BigInt::mulOneDigit( char *s1, const char digit )
{
  if ( digit == 0 )
  {
    s1[0] = 0;
    s1[1] = EOS;
    return 1;
  }

  unsigned int i = 0;

  if ( digit == 1 ) {
    while ( s1[i] != EOS ) { i++; }
    return i;
  }

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

  // Find the length, since we are already so close to the end
  while ( s1[i] != EOS ) { i++; }

  return i;
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

    memcpy( bigint + zeroCount, bigint, sizeof(char) * (length() + 1) );
    memset( bigint, 0, zeroCount );
    dataLen += zeroCount;

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
    memcpy( temp.bigint+j, bigint, sizeof(char) * (length() + 1) );
    temp.dataLen = j + dataLen;
    VALIDATE( &temp );

    // temp *= other.bigint[i]
    if ( other.bigint[i] != 1 )
    {
      temp.dataLen = mulOneDigit( temp.bigint, other.bigint[i] );
    }
    VALIDATE( &temp );

    // Sum the terms
    accumulator += temp;
    i++;
  }

  *this = accumulator;
  sign = resultSign;

  VALIDATE( this );
}

void BigInt::mulByTen( void )
{
  VALIDATE( this );
  if ( !isZero() )
    {
      extendBuffer( length() + 1 );
      memcpy( bigint + 1, bigint, sizeof(char) * (length() + 1) );
      bigint[0] = 0;
      dataLen++;
      VALIDATE( this );
    }
}

// Integer division
BigInt BigInt::div( const BigInt &denominator ) const
{
  VALIDATE( this );
  VALIDATE( &denominator );

  // Short-circuit cases
  if ( denominator == 1 )
  {
    return *this;
  }

  if ( denominator == -1 )
  {
    return *this * -1;
  }

  if ( this->isPositive() && *this < denominator )
  {
    return 0;
  }

  BigInt result = *this;
  result.sign = denominator.sign;
  if ( result == denominator )
  {
    result = 1;
    result.sign = this->sign * denominator.sign;
    return result;
  }
  result.sign = this->sign;

  BigInt i = 1;

  result = denominator;

  // Go coarsely through the numbers...
  unsigned int factor = 1;
  while ( result < *this )
  {
    result.mulByTen();
    i += 9 * factor;
    factor *= 10;
  }
  if ( result > *this )
  {
    result.divByTen();
    factor /= 10;
    i -= 9 * factor;
  }

  while ( result < *this )
  {
    result += denominator;
    i++;
  }

  if ( result > *this )
  {
    result -= denominator;
    i--;
  }

  VALIDATE( &i );

  return i;
}

unsigned int BigInt::divByTen( void )
{
  VALIDATE( this );
  unsigned int lowDigit = bigint[0] * sign;
  if ( dataLen == 1 )
    {
      bigint[0] = 0;
      sign = 1;
    }
  else
    {
      memcpy( bigint, bigint + 1, sizeof(char) * (length() + 1) );
      dataLen--;
    }
  VALIDATE( this );
  return lowDigit;
}

unsigned int BigInt::uniqueDigits( void ) const
{
  unsigned int result = 0;
  unsigned int i = 0;

  while ( i < length() )
  {
    result |= 1 << bigint[i];
    i++;
  }

  return result;
}

// Is a number pandigital in the range of low..high
bool BigInt::isPandigital( unsigned int low, unsigned int high ) const
{
  if ( length() != high - low + 1 ) { return false; }

  bool foundDigits[10] = { false, false, false, false, false, false, false, false, false, false };
  unsigned int i = 0;

  for ( i=0; i<length(); i++ )
  {
    unsigned int digit = (int) bigint[i];
    if ( digit < low || digit > high ) { return false; }
    if ( foundDigits[digit] )
    {
      return false;
    }
    foundDigits[digit] = true;
  }

  return true;
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
  VALIDATE( this );

  unsigned int i = 0;

  while ( bigint[i] == 0 )
  {
    i++;
  }

  return ( i >= 1 && bigint[i] == 1 && bigint[i+1] == EOS );
}

bool BigInt::isDivisibleBy( int divisor ) const
{
  VALIDATE( this );

  unsigned int lastDigits = 0;
  unsigned int i = 0;
  int j = 0;
  int sum = 0;

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
      return ( !(bigint[0] & 0x01)  );
      break;
    case 3:
      // If the sum of the digits of a number are divisible by 3 then the entire number is
      return ( this->sumDigits() % 3 == 0 );
      break;
    case 4:
      // If the last two digits are divisible by 4 then the entire number is
      lastDigits = bigint[0];
      if ( bigint[1] != EOS )
      {
        lastDigits += bigint[1] * 10;
      }
      return ( !(lastDigits & 0x03) );
      break;
    case 5:
      // Numbers ending in 0 or 5 are divisible by 5
      return ( bigint[0] == 0 || bigint[0] == 5 );
      break;
    case 6:
      // Numbers divisible by 2 AND by 3 are divisble by 6
      return ( this->isDivisibleBy( 2 ) && this->isDivisibleBy( 3 ) );
      break;
    case 7:
      // http://en.wikipedia.org/wiki/Divisibility_rule
      i = 0;
      while ( i < length() )
      {
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 1; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 3; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 2; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -1; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -3; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -2; }
      }
      return ( sum % 7 == 0 );
      break;
    case 8:
      // If the last 3 digits are divisible by 8 then the entire number is
      lastDigits = bigint[0];
      if ( bigint[1] != EOS )
      {
        lastDigits += bigint[1] * 10;
        if ( bigint[2] != EOS )
        {
          lastDigits += bigint[2] * 100;
        }
      }
      return ( !(lastDigits & 0x07) );
      break;
    case 9:
      // If the sum of the digits of a number are divisible by 9 then the entire number is
      return ( this->sumDigits() % 9 == 0 );
      break;
    case 10:
      // If the last digit is zero then the entire number is
      return ( bigint[0] == 0 );
      break;
    case 11:
      // http://en.wikipedia.org/wiki/Divisibility_rule
      j = length() - 1;
      while ( j >= 0 )
      {
        sum += bigint[j];
        j -= 2;
      }
      j = length() - 2;
      while ( j >= 0 )
      {
        sum -= bigint[j];
	j -= 2;
      }
      return ( sum % 11 == 0 );
      break;
    case 12:
      // You have to do 3 & 4. You can't do 2 & 6. 2 & 6 expands to 2 & (3 & 2)
      // which then simplifies to 3 & 2. That just ends up checking whether the
      // number is divisible by 6, which isn't the same
      return ( this->isDivisibleBy( 4 ) && this->isDivisibleBy( 3 ) );
      break;
    case 13:
      // http://en.wikipedia.org/wiki/Divisibility_rule
      i = 0;
      while ( i < length() )
      {
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -3; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -4; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * -1; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 3; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 4; }
        if ( bigint[i] != EOS ) { sum += bigint[i++] * 1; }
      }
      return ( sum % 13 == 0 );
      break;
    case 14:
      // Numbers divisible by 2 AND by 7 are divisble by 14
      return ( this->isDivisibleBy( 2 ) && this->isDivisibleBy( 7 ) );
      break;
    case 15:
      // Numbers divisible by 3 AND by 5 are divisble by 15
      return ( this->isDivisibleBy( 5 ) && this->isDivisibleBy( 3 ) );
      break;
    case 16:
      // If the last 4 digits are divisible by 16 then the entire number is
      lastDigits = bigint[0];
      if ( bigint[1] != EOS )
      {
        lastDigits += bigint[1] * 10;
        if ( bigint[2] != EOS )
        {
          lastDigits += bigint[2] * 100;
          if ( bigint[3] != EOS )
          {
            lastDigits += bigint[3] * 1000;
          }
        }
      }
      return ( !(lastDigits & 0x0F) );
      break;
    case 17:
      // http://en.wikipedia.org/wiki/Divisibility_rule
      if ( *this <= 17 )
      {
        return ( *this == 0 || *this == 17 );
      } else {
        BigInt temp = *this;
        do
        {
          unsigned int digit = temp.bigint[0];
          temp.chop();
          temp -= digit * 5;
          temp.sign = 1;
        } while ( temp > 51 );
        return ( temp == 0 || temp == 17 || temp == 34 || temp == 51 );
      }
      break;
    case 18:
      // You have to do 2 & 9. You can't do 3 & 6. 3 & 6 expands to 3 & (3 & 2)
      // which then simplifies to 3 & 2. That just ends up checking whether the
      // number is divisible by 6, which isn't the same
      return ( this->isDivisibleBy( 2 ) && this->isDivisibleBy( 9 ) );
      break;
    case 19:
      // http://en.wikipedia.org/wiki/Divisibility_rule
      if ( *this <= 19 )
      {
        return ( *this == 0 || *this == 19 );
      } else {
        BigInt temp = *this;
        do
        {
          unsigned int digit = temp.bigint[0];
          temp.chop();
          temp += digit << 1;
        } while ( temp > 19 );
        return ( temp == 0 || temp == 19 );
      }
      break;
    case 20:
      // You have to do 4 & 5. You can't do 2 & 10. 2 & 10 expands to 2 & (5 & 2)
      // which then simplifies to 5 & 2. That just ends up checking whether the
      // number is divisible by 10, which isn't the same
      return ( this->isDivisibleBy( 5 ) && this->isDivisibleBy( 4 ) );
      break;
    default:
      cout << "ERROR: isDivisibleBy(). Expected 0..20, got: " << divisor << endl;
      return false;
  }

  // We shouldn't get here. If we did, it is an error.
  return false;
}

bool BigInt::containsSequence( char value ) const
{
  VALIDATE( this );

  unsigned int i = 0;

  while ( bigint[i] != EOS )
  {
    if ( bigint[i] == value )
    {
      return true;
    }
    i++;
  }

  return false;
}

bool BigInt::containsSequence( const BigInt &sequence ) const
{
  VALIDATE( this );

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

unsigned int BigInt::countSequence( char value ) const
{
  VALIDATE( this );

  unsigned int i = 0;
  unsigned int count = 0;

  while ( bigint[i] != EOS )
  {
    if ( bigint[i] == value )
    {
      count++;
    }
    i++;
  }

  return count;
}

unsigned int BigInt::countSequence( const BigInt &sequence ) const
{
  VALIDATE( this );

  unsigned int i = 0;
  unsigned int count = 0;

  for ( i=0; i <= this->length() - sequence.length(); i++ )
  {
    if ( memcmp( &(this->bigint[i]), sequence.bigint, sequence.length() ) == 0 )
    {
      count++;
    }
  }

  return count;
}

const BigInt BigInt::power( BigInt const &exponent ) const
{
  VALIDATE( this );

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
    cout << endl << "ERROR: Negative exponentiation is not supported." << endl;
  }

  BigInt i      = 1;
  BigInt factor = *this;

  for ( ; i < exponent; i += 1 )
  {
    result *= factor;
  }

  VALIDATE( &result );

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
    sign = 1;
    VALIDATE( this );
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
    if ( *sptr < '0' || *sptr > '9' )
      {
	cout << "ERROR: Non-numeric character '" << *sptr << "' in input: " << s << endl;
	bigint[i++] = 0;
      }
    else
      {
	bigint[i++] = *sptr - '0';
      }
    sptr--;
  }
  bigint[i] = EOS;

  dataLen = length;

  // Remove any leading zeroes
  i = dataLen - 1;
  while ( dataLen > 1 && bigint[i] == 0 )
    {
      bigint[i] = EOS;
      dataLen--;
      i--;
    }

  VALIDATE( this );
}

void BigInt::extendBuffer( unsigned int length )
{
  // Account for the terminator character
  length++;

  // As long as we keep blocksize >= STARTER_LEN we
  // don't need to check whether we are using the
  // starter buffer or the allocated buffer,
  // we can just short-circuit on sizes that we
  // know will fit in the starter buffer
  if ( length <= STARTER_LEN )
  {
    return;
  }

  // Extending can be expensive. Do it in large blocks
  // so we don't need to do this often.
  unsigned int blocksize = STARTER_LEN << 1;
  length = blocksize * ( (length + (blocksize - 1)) / blocksize );

  if ( length > buffLen )
  {
    char *old_bigint = bigint;

    if ( bigint == starter )
    {
      bigint = NULL;
    }

    buffLen = length;
    bigint = (char *) realloc( bigint, buffLen * sizeof(char) );

    if ( bigint == NULL )
    {
      cout << "ERROR: Failed to allocate " << buffLen << " bytes of RAM." << endl;
      bigint = old_bigint;
    }

    if ( old_bigint == starter )
    {
      memcpy( bigint, starter, sizeof( starter ) );
    }
  }
}

/*
unsigned int BigInt::length( void ) const
{
  return dataLen;
}
*/

unsigned int BigInt::sumDigits( void ) const
{
  VALIDATE( this );

  char *ptr = bigint;
  unsigned int sum = *ptr++;

  while ( *ptr != EOS )
  {
    sum += *ptr++;
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
  bool found = false;

  // Verify there is a termination character.
  // Verify each digit 0 <= digit <= 9.
  i = 0;
  while ( i < buffLen )
  {
    if ( bigint[i] == EOS )
    {
      found = true;
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

  if ( length != dataLen )
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
  VALIDATE( &result );
  return result;
}

const BigInt operator-( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  BigInt result = lhs;
  result.subtract( rhs );
  VALIDATE( &result );
  return result;
}

const BigInt operator*( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  BigInt result = lhs;
  result.mul( rhs );
  VALIDATE( &result );
  return result;
}

const BigInt operator/( const BigInt &lhs, const BigInt &rhs )
{
  VALIDATE( &lhs );
  VALIDATE( &rhs );
  return lhs.div( rhs );
}
