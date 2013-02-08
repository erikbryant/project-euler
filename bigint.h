#include <iostream>

using namespace std;

//
// TODO:
//   operators
//     -
//     /
//     -=
//     /=
//   const
//   unit tests
//
class BigInt
{
  public:
    BigInt();
    BigInt( const int x );
    BigInt( const char *s );
    BigInt( const BigInt &other );
    ~BigInt();

    friend ostream &operator<<( ostream &os, const BigInt &bi );
    BigInt &operator=( const BigInt &other );
    BigInt operator+( BigInt const &other ) const;
    BigInt &operator++( void );
    BigInt &operator++( int );
    BigInt &operator+=( BigInt const &other );
    BigInt operator*( BigInt const &other ) const;
    BigInt &operator*=( BigInt const &other );
    int compare( const BigInt &other ) const;
    bool operator==( const BigInt &other ) const;
    bool operator!=( const BigInt &other ) const;
    bool operator<( const BigInt &other ) const;
    bool operator<=( const BigInt &other ) const;
    bool operator>( const BigInt &other ) const;
    bool operator>=( const BigInt &other ) const;

    BigInt power( BigInt const &exponent ) const;
    bool powerOfTen( void ) const;
    unsigned int length( void ) const;
    BigInt sumDigits( void ) const;
    bool validate( const char *file, const int line ) const;

  private:
    unsigned char *bigint;
    unsigned int  buffLen;

    void addStrings( unsigned char *s1, const unsigned char *s2 );
    void add( const BigInt &other );
    void mulOneDigit( unsigned char *s1, const unsigned char digit );
    void mul( const BigInt &other );
    void import( const int x );
    void import( const char *s );
    void extendBuffer( unsigned int length );
};

