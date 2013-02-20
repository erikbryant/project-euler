#include <iostream>
#include <string.h>

using namespace std;

#ifdef DO_VALIDATION
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ );
#else
#define VALIDATE(obj)
#endif

#define EOS 0x7F
#define STARTER_LEN 24

//
// TODO:
//   operators
//     -
//     /
//     -=
//     /=
//   real unit tests
//   Add a subOne function for operator-- to call(?)
//
class BigInt
{
  public:
    BigInt();
    BigInt( const int x );
    BigInt( const char * const s );
    BigInt( const BigInt &other );
    ~BigInt();

    friend ostream &operator<<( ostream &os, const BigInt &bi );
    const BigInt &operator=( const BigInt &rhs );
    const BigInt &operator++( void );
    const BigInt &operator++( int );
    const BigInt &operator+=( const BigInt &rhs );
    const BigInt &operator--( void );
    const BigInt &operator--( int );
    const BigInt &operator-=( const BigInt &rhs );
    const BigInt &operator*=( const BigInt &rhs );
    int compare( const BigInt &other ) const;

    bool operator==( const BigInt &other ) const
    {
      return ( compare( other ) == 0 );
    }

    bool operator!=( const BigInt &other ) const
    {
      return ( compare( other ) != 0 );
    }

    bool operator<( const BigInt &other ) const
    {
      return ( compare( other ) == -1 );
    }

    bool operator<=( const BigInt &other ) const
    {
      return ( compare( other ) != 1 );
    }

    bool operator>( const BigInt &other ) const
    {
      return ( compare( other ) == 1 );
    }

    bool operator>=( const BigInt &other ) const
    {
      return ( compare( other ) != -1 );
    }

    char &operator[]( const int i ) const
    {
      return bigint[this->length() - 1 - i];
    }

    bool isNegative( void ) const
    {
      VALIDATE( this );
      if ( isZero() && sign != 1 )
      {
        sign = 1;
      }
      return sign == -1;
    }
    bool isPositive( void ) const
    {
      VALIDATE( this );
      return sign == 1;
    }
    bool isZero( void ) const
    {
      VALIDATE( this );
      return ( bigint[0] == 0 && bigint[1] == EOS );
    }
    bool isOne( void ) const;
    bool isPowerOfTen( void ) const;
    bool isDivisibleBy( int divisor ) const;
    bool containsSequence( char value ) const;
    bool containsSequence( const BigInt &sequence ) const;

    bool containsMultiple( char v1 ) const
    {
      unsigned int i = 0;
      bool found = false;

      while ( bigint[i] != EOS )
      {
        if ( bigint[i] == v1 )
        {
          if ( found ) { return true; }
          found = true;
        }
        i++;
      }

      return false;
    }

    bool containsMultiple( char v1, char v2 ) const
    {
      unsigned int i = 0;
      bool found = false;

      while ( bigint[i] != EOS )
      {
        if ( bigint[i] == v1 || bigint[i] == v2 )
        {
          if ( found ) { return true; }
          found = true;
        }
        i++;
      }

      return false;
    }

    unsigned int countSequence( char value ) const;
    unsigned int countSequence( const BigInt &sequence ) const;

    const BigInt power( const BigInt &exponent ) const;
    unsigned int length( void ) const
    {
      return dataLen;
    }

    unsigned int sumDigits( void ) const;
    bool validate( const char *file, const int line ) const;
    void slice( unsigned int start, unsigned int length, BigInt &other ) const;
    bool testSliceDivisible( unsigned int start, unsigned int length, unsigned int divisor);

    void add( const BigInt &other );
    void addOne( void );
    void subtract( const BigInt &other );
    void mul( const BigInt &other );
    void mulByTen( void );
    unsigned int divByTen( void );
    void chop( void )
    {
      memcpy( bigint, bigint+1, length() );
      dataLen--;
      VALIDATE( this );
    }

  private:
    char *bigint;
    char starter[STARTER_LEN];
    unsigned int  buffLen;
    mutable unsigned int  dataLen;
    mutable char sign;

    unsigned int addStrings( char *s1, const char * const s2 );
    unsigned int subtractStrings( char *s1, const char * const s2 );
    unsigned int mulOneDigit( char *s1, const char digit );
    void import( const int x );
    void import( const char * const s );
    void extendBuffer( unsigned int length );
};

const BigInt operator+( const BigInt &lhs, const BigInt &rhs );
const BigInt operator-( const BigInt &lhs, const BigInt &rhs );
const BigInt operator*( const BigInt &lhs, const BigInt &rhs );

