#include <iostream>
#include <string.h>
#include "bigint.h"

using namespace std;

//
// Remove any character that is not a letter.
//
void strip( char *string )
{
    char *ptr = string;
    char *head = string;
    while ( *ptr != '\0' )
    {
      if ( toupper(*ptr) < 'A' || toupper(*ptr) > 'Z' )
      {
        ptr++;
        continue;
      }
      *head = *ptr;
      ptr++;
      head++;
    }
    *head = '\0';
}

void swap( char **a, char **b )
{
  char *temp = *a;
  *a = *b;
  *b = temp;
}

void bubbleSort( char **names )
{
  // A list with no elements is already sorted.
  // A list with one element is already sorted.
  if ( *names == NULL || *(names+1) == NULL )
  {
    return;
  }

  bool sorted = false;

  while ( !sorted )
  {
    sorted = true;
    char **ptr = names;
    while ( *(ptr+1) != NULL )
    {
      if ( strcmp( *ptr, *(ptr+1) ) > 0 )
      {
        swap( ptr, ptr+1 );
        sorted = false;
      }
      ptr++;
    }
  }
}

void insertionSort( char **names )
{
  // A list with no elements is already sorted.
  // A list with one element is already sorted.
  if ( *names == NULL || *(names+1) == NULL )
  {
    return;
  }

  unsigned int i = 1;

  while ( names[i] != NULL )
  {
    unsigned int j = i;
    while ( j >= 1 && strcmp( names[j], names[j-1] ) < 0 )
    {
      swap( names[j], names[j-1] );
      j--;
    }
    i++;
  }
}

unsigned int score( char *s )
{
  unsigned int i = 0;
  unsigned int sum = 0;

  for ( i=0; s[i] != '\0'; i++ )
  {
    sum += toupper( s[i] ) - 'A' + 1;
  }

  return sum;
}

int main( int argc, char **argv )
{
  char temp[40];
  char **names = new char*[6000];
  unsigned int i = 0;
  BigInt sum = 0;

  while ( cin.getline( temp, sizeof( temp ), ',' ) )
  {
    strip( temp );
    char *name = new char[strlen(temp)+1];
    strcpy( name, temp );
    names[i++] = name;
  }
  names[i] = NULL;

//  bubbleSort( names );
  insertionSort( names );

  for ( i=0; names[i] != NULL; i++ )
  {
//    cout << names[i] << endl;
    sum += score( names[i] ) * (i+1);
  }

  cout << "i: " << i << " sum: " << sum << endl;
}
