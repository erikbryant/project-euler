#include <stdlib.h>
#include <iostream>
#define DO_VALIDATION
#include "bigint.h"

using namespace std;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << "ERROR " << __FILE__ << ":" << __LINE__ << ": " << error << endl; errorCount++; }

void testDivisibility( unsigned int testNumber )
{
  BigInt value = 0;
  unsigned int i = 0;
  bool expected;

  for ( i=0; i<1000000; i++ )
  {
    value = i;
    expected = (i % testNumber) == 0;
    assert( value.isDivisibleBy( testNumber ) == expected, "isDivisibleBy failure value: " << value << " divisor: " << testNumber << " expected = " << expected );
  }
}

int main( int argc, char **argv )
{
    //
    // ctor
    //

    // void ctor
    BigInt a;
    assert( a == 0, "Void ctor is not zero" );
    assert( a.length() == 1, "length fail" );

    // int promotion ctor
    BigInt b = 1234567;
    assert( b == 1234567, "Int ctor failed" );
    assert( b.length() == 7, "length fail" );

    // string promotion ctor
    BigInt c = "4001";
    assert( c == 4001, "String ctor failed" );
    assert( c.length() == 4, "length fail" );

    // copy ctor
    BigInt z = c;
    assert( z == c, "copy ctor failed" );


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

    // Identity cases
    b = 0;
    c = 0;
    a = b + c;
    assert( a == (0 + 0), "operator+ fail" );

    // BigInt + int
    b = 1;
    a = b + 2;
    assert( a == (1 + 2), "operator+ fail" );

    // int + BigInt
    b = 1;
    a = 2 + b;
    assert( a == (2 + 1), "operator+ fail" );

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

    // BigInt - int
    b = 1;
    a = b - 2;
    assert( a == (1 - 2), "operator- fail" );

    // int - BigInt
    b = 1;
    a = 2 - b;
    assert( a == (2 - 1), "operator- fail" );

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

    // BigInt * int
    b = 10;
    a = b * 2;
    assert( a == (10 * 2), "operator* fail" );

    // int * BigInt
    b = 10;
    a = 2 * b;
    assert( a == (2 * 10), "operator* fail" );

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

//    Not supported yet
//    a = 4;
//    a = a.power( -1 );
//    assert( a == 2, "power fail" );

    // Verify that large exponents don't crash or hang
    a = "1234567890";
    b = "444";
    c = a.power( b );



    //
    // operator[]
    //

    a = "1543267890";
    assert( a[0] == 1, "operator[] fail" );
    assert( a[1] == 5, "operator[] fail" );
    assert( a[2] == 4, "operator[] fail" );
    assert( a[3] == 3, "operator[] fail" );
    assert( a[4] == 2, "operator[] fail" );
    assert( a[5] == 6, "operator[] fail" );
    assert( a[6] == 7, "operator[] fail" );
    assert( a[7] == 8, "operator[] fail" );
    assert( a[8] == 9, "operator[] fail" );
    assert( a[9] == 0, "operator[] fail" );


    //
    // isNegative
    //

    a = -3;
    assert( a.isNegative(), "isPositive fail" );
    a = 0;
    a--;
    assert( a.isNegative(), "isPositive fail" );

    //
    // isPositive
    //

    a = 3;
    assert( a.isPositive(), "isPositive fail" );
    a = -1;
    a++;
    assert( a.isPositive(), "isPositive fail" );



    //
    // isZero
    //

    a = 0;
    assert( a.isZero(), "isZero fail" );
    a = 3;
    a -=3;
    assert( a.isZero(), "isZero fail" );
    a = -4;
    a += 4;
    assert( a.isZero(), "isZero fail" );



    //
    // containsMultiple
    //

    a = "1003";
    assert( a.containsMultiple( 0 ), "containsMultiple fail" );
    a = "103";
    assert( !a.containsMultiple( 0 ), "containsMultiple fail" );

    a = "1003";
    assert( a.containsMultiple( 1, 3 ), "containsMultiple fail" );
    a = "103";
    assert( !a.containsMultiple( 1, 4 ), "containsMultiple fail" );



    //
    // countSequence
    //

    a = "1234123";
    assert( a.countSequence( 1 ) == 2, "countSequence fail" );
    a = "123423";
    assert( a.countSequence( 1 ) == 1, "countSequence fail" );
    a = "234203";
    assert( a.countSequence( 1 ) == 0, "countSequence fail" );

    a = "1234123";
    b = "123";
    assert( a.countSequence( b ) == 2, "countSequence fail" );
    a = "12341203";
    assert( a.countSequence( b ) == 1, "countSequence fail" );
    a = "1341203";
    assert( a.countSequence( b ) == 0, "countSequence fail" );




    //
    // isPowerOfTen
    //

    a = 1;
    assert( !a.isPowerOfTen(), "powerOfTen fail" );

    a = 100;
    assert( a.isPowerOfTen(), "powerOfTen fail" );



    //
    // isDivisibleBy
    //

    a = 4;
    assert( a.isDivisibleBy( 4 ) == true,  "isDivisbleBy failure" );

    a = 8;
    assert( a.isDivisibleBy( 4 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 8 ) == true,  "isDivisbleBy failure" );

    a = 9;
    assert( a.isDivisibleBy( 9 ) == true,  "isDivisbleBy failure" );

    a = 48;
    assert( a.isDivisibleBy( 0 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 1 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 2 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 3 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 4 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 8 ) == true,  "isDivisbleBy failure" );

    a = 30;
    assert( a.isDivisibleBy( 5 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 6 ) == true,  "isDivisbleBy failure" );

    a = 888;
    assert( a.isDivisibleBy( 8 ) == true,  "isDivisbleBy failure" );

    a = 999;
    assert( a.isDivisibleBy( 9 ) == true,  "isDivisbleBy failure" );

    a = 623;
    assert( a.isDivisibleBy( 7 ) == true,  "isDivisbleBy failure" );

    a = 127;
    assert( a.isDivisibleBy( 0 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 1 ) == true,  "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 2 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 3 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 4 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 5 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 6 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 7 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 8 ) == false, "isDivisbleBy failure" );
    assert( a.isDivisibleBy( 9 ) == false, "isDivisbleBy failure" );

    a = "77777777777777";
    assert( a.isDivisibleBy( 7 ) == true, "isDivisbleBy failure" );

    unsigned int num = 0;
    for ( num=1; num<=20; num++ )
    {
      testDivisibility( num );
    }

    //
    // sumDigits
    //

    a = 12345;
    assert( a.sumDigits() == 15, "sumDigits fail" );



    //
    // containsSequence
    //

    a = "1234567890";
    assert( a.containsSequence( "1" ) == true, "containsSequence failure" );
    assert( a.containsSequence( "456" ) == true, "containsSequence failure" );
    assert( a.containsSequence( "123456789" ) == true, "containsSequence failure" );
    assert( a.containsSequence( "234567890" ) == true, "containsSequence failure" );
    assert( a.containsSequence( "1234567890" ) == true, "containsSequence failure" );
    assert( a.containsSequence( "12345678901" ) == false, "containsSequence failure" );
    assert( a.containsSequence( "11234567890" ) == false, "containsSequence failure" );



    //
    // isPandigital
    //

    a = "0";
    assert( a.isPandigital( 0, 0 ), "isPandigital fail" );
    a = "456";
    assert( a.isPandigital( 4, 6 ), "isPandigital fail" );
    a = "0123456789";
    assert( a.isPandigital( 0, 9 ), "isPandigital fail" );
    a = "012345678";
    assert( !a.isPandigital( 0, 9 ), "isPandigital fail" );
    a = "01234567xi98";
    assert( !a.isPandigital( 0, 8 ), "isPandigital fail" );

  exit( errorCount );
}
