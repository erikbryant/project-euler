//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include <iostream>
#include <fstream>
#include <string>

using std::cout;
using std::endl;
using std::ifstream;
using std::ios;
using std::string;

unsigned int rtod( char digit )
{
  unsigned int value = 0;

  switch ( toupper( digit ) )
    {
    case 'I':
      value = 1;
      break;
    case 'V':
      value = 5;
      break;
    case 'X':
      value = 10;
      break;
    case 'L':
      value = 50;
      break;
    case 'C':
      value = 100;
      break;
    case 'D':
      value = 500;
      break;
    case 'M':
      value = 1000;
      break;
    default:
      cout << "Invalid character '" << digit << "' in input." << endl;
      return 0;
      break;
    }

  return value;
}

int compareRomanDigits( char digit1, char digit2 )
{
  unsigned int d1 = rtod( digit1 );
  unsigned int d2 = rtod( digit2 );
  int result = 0;

  if ( d1 < d2 )
    {
      result = -1;
    }
  else if ( d1 > d2 )
    {
      result = 1;
    }

  return result;
}

unsigned int rtod( const char *roman )
{
  unsigned int value = 0;
  unsigned int temp = 0;
  const char *ptr = roman;
  char lastDigit = *ptr;

  for ( ptr = roman; *ptr != '\0'; ++ptr )
    {
      if ( rtod( *ptr ) > rtod( lastDigit ) )
        {
          value -= temp;
          temp = 0;
        }
      else if ( rtod( *ptr ) < rtod( lastDigit ) )
        {
          value += temp;
          temp = 0;
        }
      temp += rtod( *ptr );
      lastDigit = *ptr;
    }

  value += temp;

  return value;
}

void dtor( unsigned int value, char *outBuff, unsigned int buffLen )
{
  unsigned int i = 0;

  while ( value >= 1000 )
    {
      outBuff[i++] = 'M';
      value -= 1000;
    }

  if ( value >= 900 )
    {
      outBuff[i++] = 'C';
      outBuff[i++] = 'M';
      value -= 900;
    }

  while ( value >= 500 )
    {
      outBuff[i++] = 'D';
      value -= 500;
    }

  if ( value >= 400 )
    {
      outBuff[i++] = 'C';
      outBuff[i++] = 'D';
      value -= 400;
    }

  while ( value >= 100 )
    {
      outBuff[i++] = 'C';
      value -= 100;
    }

  if ( value >= 90 )
    {
      outBuff[i++] = 'X';
      outBuff[i++] = 'C';
      value -= 90;
    }

  while ( value >= 50 )
    {
      outBuff[i++] = 'L';
      value -= 50;
    }

  if ( value >= 40 )
    {
      outBuff[i++] = 'X';
      outBuff[i++] = 'L';
      value -= 40;
    }

  while ( value >= 10 )
    {
      outBuff[i++] = 'X';
      value -= 10;
    }

  if ( value >= 9 )
    {
      outBuff[i++] = 'I';
      outBuff[i++] = 'X';
      value -= 9;
    }

  while ( value >= 5 )
    {
      outBuff[i++] = 'V';
      value -= 5;
    }

  if ( value >= 4 )
    {
      outBuff[i++] = 'I';
      outBuff[i++] = 'V';
      value -= 4;
    }

  while ( value >= 1 )
    {
      outBuff[i++] = 'I';
      value -= 1;
    }

  outBuff[i] = '\0';
}

void shrinkRoman( const string &input, string &output )
{
  unsigned int value = rtod( input.c_str() );
  char outBuff[100];

  dtor( value, outBuff, 100 );

  output = outBuff;
}

int main( int argc, char *argv[] )
{
#if 0
  cout << "I = " << rtod( "I" ) << endl;
  cout << "V = " << rtod( "V" ) << endl;
  cout << "X = " << rtod( "X" ) << endl;
  cout << "L = " << rtod( "L" ) << endl;
  cout << "C = " << rtod( "C" ) << endl;
  cout << "D = " << rtod( "D" ) << endl;
  cout << "M = " << rtod( "M" ) << endl;
  cout << endl;
  cout << "DCXXV    (625) = " << rtod( "DCXXV" ) << endl;
  cout << "MM      (2000) = " << rtod( "MM" ) << endl;
  cout << "MCMXLIV (1944) = " << rtod( "MCMXLIV" ) << endl;
  cout << "           (0) = " << rtod( "" ) << endl;
  cout << endl;
  cout << "IIIIIIIIIIIIIIII  (16) = " << rtod( "IIIIIIIIIIIIIIII" ) << endl;
  cout << "VIIIIIIIIIII      (16) = " << rtod( "VIIIIIIIIIII" ) << endl;
  cout << "VVIIIIII          (16) = " << rtod( "VVIIIIII" ) << endl;
  cout << "XIIIIII           (16) = " << rtod( "XIIIIII" ) << endl;
  cout << "VVVI              (16) = " << rtod( "VVVI" ) << endl;
  cout << "XVI               (16) = " << rtod( "XVI" ) << endl;
#endif

  ifstream myFile;
  myFile.open( "089.data", ios::in );

  string input;
  string output;
  unsigned int oldSize = 0;
  unsigned int newSize = 0;
  while ( getline( myFile, input ) )
    {
      oldSize += input.length();
      shrinkRoman( input, output );
      newSize += output.length();
      if ( input.length() < output.length() )
        {
          cout << "WARNING: Space lost converting '" << input << "' to '" << output << "'" << endl;
        }
      if ( rtod( input.c_str() ) != rtod( output.c_str() ) )
        {
          cout << "ERROR: Input value '" << input << "' (" << rtod( input.c_str() ) << ") and output value '"
               << output << "' (" << rtod( output.c_str() ) << ") differ!" << endl;
        }
    }

  cout << "Space saved: " << oldSize << " - " << newSize << " = " << oldSize - newSize << endl;
}
