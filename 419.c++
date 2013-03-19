#include <cstring>
#include <cstdlib>
#include <cassert>
#include <cstdio>
#include <iostream>

using std::cout;
using std::endl;

unsigned int lookAndSay( char *buff1, char *buff2 )
{
  unsigned int i = 0;
  char digit = buff1[0];
  unsigned int count = 1;
  unsigned int j = 0;

  buff2[0] = '\0';

  do
    {
      ++i;
      if ( buff1[i] == digit )
        {
          count++;
        }
      else
        {
          j += sprintf( buff2 + j, "%d%d", count, digit - '0' );
          digit = buff1[i];
          count = 1;
        }
    } while ( buff1[i] != '\0' );

  return j;
}

void countDigits( char *buff, unsigned long long int &one, unsigned long long int &two, unsigned long long int &three )
{
  while ( *buff != '\0' )
    {
      switch ( *buff )
        {
        case '1':
          ++one;
          break;
        case '2':
          ++two;
          break;
        case '3':
          ++three;
          break;
        }
      buff++;
    }
}

void printCount( char *buff, unsigned int len, unsigned int i )
{
  unsigned long long int one = 0;
  unsigned long long int two = 0;
  unsigned long long int three = 0;
  countDigits( buff, one, two, three );

  cout << " (" << i << ")\t\t" << len << "\t" << one << "\t" << two << "\t" << three << endl;
}

void printJustified( char *buff, unsigned int len, unsigned int i )
{
#if 0
  if ( len < 180 )
    {
      unsigned int pad = 180 - len;

      for ( ; pad >= 1; --pad )
        {
          cout << " ";
        }
    }
#endif

  printCount( buff, len, i );

#if 0
  cout << buff << " (" << i << ")" << endl;
#endif
}

void verifyPatterns( char *buff, unsigned int len, unsigned int i )
{
  char *ptr = buff + len - 1;

  // The final digit is always '1'
  assert ( *ptr == '1' );

  if ( ptr == buff ) { return; } else { --ptr; }

  // The second-to-final digit is a 1 or a 2
  if ( i % 2 == 0 )
    {
      assert( *ptr == '1' );
    }
  else
    {
      assert( *ptr == '2' );
    }

  if ( ptr == buff ) { return; } else { --ptr; }

  // The third-to-final digit is always 2
  assert ( *ptr == '2' );

  if ( ptr == buff ) { return; } else { --ptr; }

  // The fourth-to-final digit goes: 1, 2, 2, 3
  // (starting at i == 5)
  if ( i >= 5 )
    {
      switch ( i % 4 )
        {
        case 1:
          assert( *ptr == '1' );
          break;
        case 2:
          assert( *ptr == '2' );
          break;
        case 3:
          assert( *ptr == '2' );
          break;
        case 0:
          assert( *ptr == '3' );
          break;
        }
    }

  if ( ptr == buff ) { return; } else { --ptr; }

  // The fifth-to-final digit goes: 1, 1, 1, 3
  // (starting at i == 6)
  if ( i >= 6 )
    {
      switch ( i % 4 )
        {
        case 2:
          assert( *ptr == '1' );
          break;
        case 3:
          assert( *ptr == '1' );
          break;
        case 0:
          assert( *ptr == '1' );
          break;
        case 1:
          assert( *ptr == '3' );
          break;
        }
    }

  if ( ptr == buff ) { return; } else { --ptr; }

  // The sixth-to-final digit goes: 1, 1, 2, 1
  // (starting at i == 9)
  if ( i >= 9 )
    {
      switch ( i % 4 )
        {
        case 1:
          assert( *ptr == '1' );
          break;
        case 2:
          assert( *ptr == '1' );
          break;
        case 3:
          assert( *ptr == '2' );
          break;
        case 0:
          assert( *ptr == '1' );
          break;
        }
    }

}

void discoverPatterns( void )
{
  unsigned int BUFF_COUNT = 12;
  unsigned int MAX_BUFF   = 20000000;
  char *buffers[BUFF_COUNT];
  unsigned int len[BUFF_COUNT];
  unsigned int i = 0;

  for ( i = 0; i < BUFF_COUNT; ++i )
    {
      buffers[i] = new char[MAX_BUFF];
      buffers[i][0] = '\0';
      len[i] = 0;
    }

  buffers[0][0] = '1';
  buffers[0][1] = '\0';

  unsigned int f = 1;
  while ( 1 ) 
    {
      for ( i = 1; i < BUFF_COUNT; ++i )
        {
          len[i] = lookAndSay( buffers[i-1], buffers[i] );
        }

      unsigned int x = 0;
      bool foundFail = false;
      for ( x = 0; x < len[0]; ++x )
        {
          for ( i = 0; i < 4; ++i )
            {
              if ( buffers[i][len[i]-x] != buffers[i+4][len[i+4]-x] )
                {
                  cout << "(" << f+i << ")\t" << len[0] << "\tfail digit:\t" << x << "\tmatch %: " << 100 * x / len[0] << " -> ";
                  foundFail = true;
                  break;
                }
            }
          if ( foundFail )
            {
              for ( i = 0; i < BUFF_COUNT; ++i )
                {
                  cout << buffers[i][len[i]-x] << ", ";
                }
              cout << endl;
              break;
            }
        }

      memcpy( buffers[0], buffers[1], len[1] );
      len[0] = len[1];
      ++f;
    }
}

int main( int argc, char *argv[] )
{
  unsigned int i = 1;
  unsigned int MAX_BUFF = 20000000;
  char *buff1 = new char[MAX_BUFF];
  char *buff2 = new char[MAX_BUFF];
  unsigned int len = 0;

  buff1[0] = '1';
  buff1[1] = '\0';
  len = 1;

  discoverPatterns();

  while ( 1 )
    {
      printJustified( buff1, len, i );
      verifyPatterns( buff1, len, i );
      ++i;
      len = lookAndSay( buff1, buff2 );
      printJustified( buff2, len, i );
      verifyPatterns( buff2, len, i );
      ++i;
      len = lookAndSay( buff2, buff1 );
    }

  delete[] buff1;
  delete[] buff2;
}
