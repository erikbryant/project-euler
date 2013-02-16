#include <iostream>

using namespace std;

//
// TODO:
//   operators
//     -
//     /
//     -=
//     /=
//     []
//   const
//   real unit tests
//   Calculate length in more places instead of just marking dirty
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
    bool operator==( const BigInt &other ) const;
    bool operator!=( const BigInt &other ) const;
    bool operator<( const BigInt &other ) const;
    bool operator<=( const BigInt &other ) const;
    bool operator>( const BigInt &other ) const;
    bool operator>=( const BigInt &other ) const;
    const char operator[]( const int i ) const;

    bool isNegative( void ) const;
    bool isPositive( void ) const;
    bool isZero( void ) const;
    bool isOne( void ) const;
    bool isPowerOfTen( void ) const;
    bool isDivisibleBy( int divisor ) const;
    bool containsSequence( char value ) const;
    bool containsSequence( const BigInt &sequence ) const;
    unsigned int countSequence( char value ) const;
    unsigned int countSequence( const BigInt &sequence ) const;

    const BigInt power( const BigInt &exponent ) const;
    unsigned int length( void ) const;
    const BigInt sumDigits( void ) const;
    bool validate( const char *file, const int line ) const;
    void slice( unsigned int start, unsigned int length, BigInt &other ) const;

    void add( const BigInt &other );
    void subtract( const BigInt &other );
    void mul( const BigInt &other );

  private:
    char *bigint;
    unsigned int  buffLen;
    mutable unsigned int  dataLen;
    mutable unsigned char dirty;
    mutable char sign;

    void addStrings( char *s1, const char * const s2 );
    void subtractStrings( char *s1, const char * const s2 );
    void mulOneDigit( char *s1, const char digit );
    void import( const int x );
    void import( const char * const s );
    void extendBuffer( unsigned int length );
};

const BigInt operator+( const BigInt &lhs, const BigInt &rhs );
const BigInt operator-( const BigInt &lhs, const BigInt &rhs );
const BigInt operator*( const BigInt &lhs, const BigInt &rhs );

