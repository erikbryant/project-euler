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
    const BigInt operator+( const BigInt &other ) const;
    const BigInt &operator++( void );
    const BigInt &operator++( int );
    const BigInt &operator+=( const BigInt &rhs );
    const BigInt operator*( const BigInt &other ) const;
    const BigInt &operator*=( const BigInt &rhs );
    int compare( const BigInt &other ) const;
    bool operator==( const BigInt &other ) const;
    bool operator!=( const BigInt &other ) const;
    bool operator<( const BigInt &other ) const;
    bool operator<=( const BigInt &other ) const;
    bool operator>( const BigInt &other ) const;
    bool operator>=( const BigInt &other ) const;

    const BigInt power( const BigInt &exponent ) const;
    bool powerOfTen( void ) const;
    unsigned int length( void ) const;
    const BigInt sumDigits( void ) const;
    bool validate( const char *file, const int line ) const;

  private:
    unsigned char *bigint;
    unsigned int  buffLen;

    void addStrings( unsigned char *s1, const unsigned char * const s2 );
    void add( const BigInt &other );
    void mulOneDigit( unsigned char *s1, const unsigned char digit );
    void mul( const BigInt &other );
    void import( const int x );
    void import( const char * const s );
    void extendBuffer( unsigned int length );
};

