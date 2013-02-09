#include <stdlib.h>
#include <iostream>
#include "bigint.h"

using namespace std;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << "ERROR " << __FILE__ << ":" << __LINE__ << ": " << error << endl; errorCount++; }

int main( int argc, char **argv )
{
  BigInt a;
  assert( a == 0, "Void ctor is not zero" );

  BigInt b = 1234567;
  assert( b == 1234567, "Int ctor failed" );

  BigInt c = "4001";
  assert( c == 4001, "String ctor failed" );

  b = 1234567;
  c = b;
  assert( c == 1234567, "operator= fail" );

  b = 1234567;
  b = b;
  assert( b == 1234567, "operator= fail" );

  b = 1234567;
  a = c = b;
  assert( c == 1234567, "operator= fail" );
  assert( a == 1234567, "operator= fail" );

  b = 1234;
  c = 1111;
  a = b + c;
  assert( a == 2345, "operator+ fail" );

  a = 9999;
  ++a;
  assert( a == 10000, "prefix operator++ fail" );

  a = 9999;
  a++;
  assert( a == 10000, "postfix operator++ fail" );

  a = 4;
  b = ++a;
  assert( b == 5, "prefix operator++ fail" );

  a = 4;
  b = a++;
  assert( b == 5, "postfix operator++ fail" );

  a = 9999;
  a += 1;
  assert( a == 10000, "operator+= fail" );

  a = 100;
  a += a;
  assert( a == 200, "operator+= fail" );

  c = 900;
  a = c * 2;
  assert( a == 1800, "operator* fail" );

  a = 430;
  a *= 2;
  assert( a == 860, "operator*= fail" );

  a = 200;
  a *= a;
  assert( a == 40000, "operator*= fail" );

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

  a = 2;
  b = 5;
  c = a.power( b );
  assert( c == 32, "power fail" );

  a = 3;
  c = a.power( a );
  assert( c == 27, "power fail" );

  a = 1;
  assert( !a.powerOfTen(), "powerOfTen fail" );

  a = 100;
  assert( a.powerOfTen(), "powerOfTen fail" );

  a = 0;
  assert( a.length() == 1, "length fail" );

  a = 98765;
  assert( a.length() == 5, "length fail" );

  a = 12345;
  assert( a.sumDigits() == 15, "sumDigits fail" );

  exit( errorCount );
}
