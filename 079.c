#include <iostream>
#include <string.h>

using namespace std;

bool containsAttempt( const char *passcode, const char *attempt )
{
  unsigned int i = 0;
  unsigned int j = 0;

  for ( i=0; passcode[i] != '\0'; i++ )
    {
      if ( attempt[j] == passcode[i] )
	{
	  j++;
	  if ( attempt[j] == '\0' )
	    {
	      return true;
	    }
	}
    }

  return false;
}

bool evalAttempt( char *passcode, const char *attempt )
{
  if ( passcode[0] == '\0' )
    {
      strcpy( passcode, attempt );
      return true;
    }

  // Does this go at the start?
  if ( passcode[0] == attempt[1] && passcode[1] == attempt[2] )
    {
      strcpy( passcode + 1, passcode );
      passcode[0] = attempt[0];
      return true;
    }

  // Does this go at the end?
  unsigned int length = strlen( passcode );
  if ( passcode[length - 2] == attempt[0] && passcode[length - 1] == attempt[1] )
    {
      passcode[length] = attempt[2];
      passcode[length + 1] = '\0';
      return true;
    }

  // Does this go in the middle?
  unsigned int i = 0;
  for ( i=0; i<length - 1; i++ )
    {
      if ( passcode[i] == attempt[0] && passcode[i + 1] == attempt[2] )
	{
	  strcpy( passcode + i + 2, passcode + i + 1 );
	  passcode[i + 1] = attempt[1];
	  return true;
	}
    }

  return false;
}

int main( int argc, char *argv[] )
{
  char passcode[255];
  bool usedAll = true;
  unsigned int i = 0;
  const char *attempts[] = {
    "129",
    "160",
    "162",
    "168",
    "180",
    "289",
    "290",
    "316",
    "318",
    "319",
    "362",
    "368",
    "380",
    "389",
    "620",
    "629",
    "680",
    "689",
    "690",
    "710",
    "716",
    "718",
    "719",
    "720",
    "728",
    "729",
    "731",
    "736",
    "760",
    "762",
    "769",
    "790",
    "890",
  };
  passcode[0] = '\0';

  do
    {
      usedAll = true;
      for ( i=0; i<sizeof(attempts)/sizeof(char *); i++ )
	{
	  cout << "Passcode: " << passcode << endl;
	  if ( !containsAttempt( passcode, attempts[i] ) )
	    {
	      cout << " Attempt: " << attempts[i] << endl;
	      usedAll = false;
	      evalAttempt( passcode, attempts[i] );
	    }
	}
    } while ( !usedAll );

}
