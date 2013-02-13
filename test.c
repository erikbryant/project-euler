#include <stdlib.h>
#include <iostream>
#include "bigint.h"

using namespace std;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << "ERROR " << __FILE__ << ":" << __LINE__ << ": " << error << endl; errorCount++; }

int main( int argc, char **argv )
{
  unsigned int i = 0;

//  for ( i=0; i<100000; i++ )
  {
    //
    // ctor
    //

    BigInt a;
    assert( a == 0, "Void ctor is not zero" );
    assert( a.length() == 1, "length fail" );

    BigInt b = 1234567;
    assert( b == 1234567, "Int ctor failed" );
    assert( b.length() == 7, "length fail" );

    BigInt c = "4001";
    assert( c == 4001, "String ctor failed" );
    assert( c.length() == 4, "length fail" );



    //
    // negative numbers
    //

    BigInt d = -1;
    BigInt e = -2;
    assert( d.isNegative(), "negative fail" );
    assert( d.length() == 1, "length fail" );
    assert( d < 0, "comparison fail" );
    assert( d > e, "comparison fail" );

    d = "-123";
    assert( d.isNegative(), "negative fail" );
    assert( d.length() == 3, "length fail" );
    assert( d == -123, "comparison fail" );
    assert( d < e, "comparison fail" );
    assert( e > d, "comparison fail" );

    e = "123";
    assert( d != e, "comparison fail" );

    e = "-122";
    assert( d < e, "comparison fail" );

    e = "-124";
    assert( d > e, "comparison fail" );



    //
    // operator=
    //

    b = 1234567;
    c = b;
    assert( c == 1234567, "operator= fail" );
    assert( c.length() == 7, "length fail" );

    b = 1234567;
    b = b;
    assert( b == 1234567, "operator= fail" );
    assert( b.length() == 7, "length fail" );

    b = 1234567;
    a = c = b;
    assert( c == 1234567, "operator= fail" );
    assert( a == 1234567, "operator= fail" );
    assert( a.length() == 7, "length fail" );
    assert( b.length() == 7, "length fail" );
    assert( c.length() == 7, "length fail" );



    //
    // operator+
    //
    // comment notation:
    //   -b > +c
    //   b is a negative value
    //   c is a positive value
    //   abs(b) > abs(c)
    //

    // +b < +c
    b = 1234;
    c = 119911;
    a = b + c;
    assert( a == (1234 + 119911), "operator+ fail" );

    // +b > +c
    b = 129934;
    c = 1111;
    a = b + c;
    assert( a == (129934 + 1111), "operator+ fail" );

    // +b < -c
    b = 129934;
    c = -987654321;
    a = b + c;
    assert( a == (129934 + -987654321), "operator+ fail" );

    // +b > -c
    b = 129934;
    c = -987;
    a = b + c;
    assert( a == (129934 + -987), "operator+ fail" );

    // -b < +c
    b = -123;
    c = 98007;
    a = b + c;
    assert( a == (-123 + 98007), "operator+ fail" );

    // -b > +c
    b = -657123;
    c = 98007;
    a = b + c;
    assert( a == (-657123 + 98007), "operator+ fail" );

    // -b < -c
    b = -65;
    c = -98007;
    a = b + c;
    assert( a == (-65 + -98007), "operator+ fail" );

    // -b > -c
    b = -8675309;
    c = -42;
    a = b + c;
    assert( a == (-8675309 + -42), "operator+ fail" );



    //
    // operator-
    //
    // comment notation:
    //   -b > +c
    //   b is a negative value
    //   c is a positive value
    //   abs(b) > abs(c)
    //

    // +b < +c
    b = 1;
    c = 3645;
    a = b - c;
    assert( a == (1 - 3645), "operator- fail" );

    b = 9;
    c = 21189;
    a = b - c;
    assert( a == (9 - 21189), "operator- fail" );

    // +b > +c
    b = 20654;
    c = 1;
    a = b - c;
    assert( a == (20654 - 1), "operator- fail" );

    b = 2118302;
    c = 9;
    a = b - c;
    assert( a == (2118302 - 9), "operator- fail" );

    // +b < -c
    b = 211;
    c = -300657;
    a = b - c;
    assert( a == (211 - -300657), "operator- fail" );

    // +b > -c
    b = 211911;
    c = -9;
    a = b - c;
    assert( a == (211911 - -9), "operator- fail" );

    // -b < +c
    b = -12;
    c = 806219;
    a = b - c;
    assert( a == (-12 - 806219), "operator- fail" );

    // -b > +c
    b = -2115545;
    c = 9;
    a = b - c;
    assert( a == (-2115545 - 9), "operator- fail" );

    // -b < -c
    b = -18;
    c = -2448808;
    a = b - c;
    assert( a == (-18 - -2448808), "operator- fail" );

    // -b > -c
    b = -568830;
    c = -39;
    a = b - c;
    assert( a == (-568830 - -39), "operator- fail" );

    // Manipulations where lhs and rhs are the same objects

    // a = a - c
    a = 10;
    c = 20;
    a = a - c;
    assert( a == (10 - 20), "operator- fail" );

    // a = a - a
    a = 10;
    a = a - a;
    assert( a == (10 - 10), "operator- fail" );

    // a = a - a - a - a
    a = 10;
    a = a - a - a - a;
    assert( a == (10 - 10 - 10 - 10), "operator- fail" );



    //
    // operator++ (prefix)
    //

    a = 9999;
    ++a;
    assert( a == (9999 + 1), "prefix operator++ fail" );

    a = 4;
    b = ++a;
    assert( a == (4 + 1), "prefix operator++ fail" );
    assert( b == (4 + 1), "prefix operator++ fail" );

    a = -2;
    ++a;
    assert( a == (-2 + 1), "prefix operator++ fail" );

    a = -1;
    ++a;
    assert( a == (-1 + 1), "prefix operator++ fail" );



    //
    // operator++ (postfix)
    //

    a = 9999;
    a++;
    assert( a == (9999 + 1), "postfix operator++ fail" );

    a = 4;
    b = a++;
    assert( a == (4 + 1), "postfix operator++ fail" );
    assert( b == (4 + 1), "postfix operator++ fail" );

    a = -2;
    a++;
    assert( a == (-2 + 1), "postfix operator++ fail" );

    a = -1;
    a++;
    assert( a == (-1 + 1), "postfix operator++ fail" );



    //
    // operator-- (prefix)
    //

    a = 9999;
    --a;
    assert( a == (9999 - 1), "prefix operator-- fail" );

    a = 4;
    b = --a;
    assert( a == (4 - 1), "prefix operator-- fail" );
    assert( b == (4 - 1), "prefix operator-- fail" );

    a = 0;
    --a;
    assert( a == (0 - 1), "prefix operator-- fail" );

    a = -1;
    --a;
    assert( a == (-1 - 1), "prefix operator-- fail" );



    //
    // operator-- (postfix)
    //

    a = 9999;
    a--;
    assert( a == (9999 - 1), "postfix operator-- fail" );

    a = 4;
    b = a--;
    assert( a == (4 - 1), "postfix operator-- fail" );
    assert( b == (4 - 1), "postfix operator-- fail" );

    a = 0;
    a--;
    assert( a == (0 - 1), "postfix operator-- fail" );

    a = -1;
    a--;
    assert( a == (-1 - 1), "postfix operator-- fail" );



    //
    // operator+=
    //

    a = 9999;
    a += 1;
    assert( a == (9999 + 1), "operator+= fail" );

    a = 100;
    a += a;
    assert( a == (100 + 100), "operator+= fail" );



    //
    // operator*
    //

    c = 900;
    a = c * 2;
    assert( a == 1800, "operator* fail" );

    a = c * 1;
    assert( a == 900, "operator* fail" );

    a = c * 0;
    assert( a == 0, "operator* fail" );

    c = 1;
    a = c * 900;
    assert( a == (1 * 900), "operator* fail" );

    c = 0;
    a = c * 900;
    assert( a == (0 * 900), "operator* fail" );

    c = 101;
    d = -1;
    a = c * d;
    assert( a == (101 * -1), "operator* fail" );

    c = 101;
    d = -12;
    a = c * d;
    assert( a == (101 * -12), "operator* fail" );

    c = -45;
    d = 892;
    a = c * d;
    assert( a == (-45 * 892), "operator* fail" );

    c = -689;
    d = -45897;
    a = c * d;
    assert( a == (-689 * -45897), "operator* fail" );



    //
    // operator*=
    //

    a = 430;
    a *= 2;
    assert( a == 860, "operator*= fail" );

    a = 200;
    a *= a;
    assert( a == 40000, "operator*= fail" );



    //
    // Comparison operators
    //

    a = 5;
    b = 5;
    assert( a == b, "operator== fail" );
    assert( !(a != b), "operator!= fail" );

    a = 9999;
    b = 7;
    assert( !(a == b), "operator== fail" );
    assert( a != b, "operator!= fail" );
    assert( !(a < b), "operator< fail" );
    assert( !(a <= b), "operator<= fail" );
    assert( a > b, "operator> fail" );
    assert( a >= b, "operator>= fail" );



    //
    // Exponentiation
    //

    a = 2;
    b = 5;
    c = a.power( b );
    assert( c == 32, "power fail" );

    a = 3;
    c = a.power( a );
    assert( c == 27, "power fail" );

    a = 0;
    a = a.power( 0 );
    assert( a == 1, "power fail" );

    a = 0;
    a = a.power( 1 );
    assert( a == 0, "power fail" );

    a = 1;
    a = a.power( 0 );
    assert( a == 1, "power fail" );

    a = 1;
    a = a.power( 1 );
    assert( a == 1, "power fail" );

    a = 999;
    a = a.power( 1 );
    assert( a == 999, "power fail" );

    a = -1;
    a = a.power( 0 );
    assert( a == 1, "power fail" );

    a = -1;
    a = a.power( 1 );
    assert( a == -1, "power fail" );

    a = -1;
    a = a.power( 2 );
    assert( a == 1, "power fail" );

    a = -1;
    a = a.power( 3 );
    assert( a == -1, "power fail" );

    a = 4;
    a = a.power( -1 );
    assert( a == 2, "power fail" );




    a = "1234567890";
    b = "4";
    c = a.power( b );



    //
    // isPowerOfTen
    //

    a = 1;
    assert( !a.isPowerOfTen(), "powerOfTen fail" );

    a = 100;
    assert( a.isPowerOfTen(), "powerOfTen fail" );



    //
    // sumDigits
    //

    a = 12345;
    assert( a.sumDigits() == 15, "sumDigits fail" );
  }

  exit( errorCount );
}
