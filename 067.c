#include <iostream>
#include <iomanip>
#include <fstream>
#include <cstdlib>

using namespace std;

/*
 * This is the same as problem 018.
 */

#define MAX_DIM 100

void printArray( int array[MAX_DIM][MAX_DIM] )
{
  int row = 0;
  int col = 0;

  for ( row=0; row<MAX_DIM; row++ )
    {
      for ( col = 0; col<MAX_DIM; col++ )
	{
	  cout << setw(4) << array[row][col] << " ";
	}
      cout << endl;
    }
}

void zeroArray( int array[MAX_DIM][MAX_DIM] )
{
  int row = 0;
  int col = 0;

  for ( row=0; row<MAX_DIM; row++ )
    {
      for ( col = 0; col<MAX_DIM; col++ )
	{
	  array[row][col] = 0;
	}
    }
}

void readArray( int array[MAX_DIM][MAX_DIM] )
{
  ifstream myFile;
  string line;
  int row = 0;
  int col = 0;
  int i = 0;

  myFile.open( "067.data", ios::in );
  while ( myFile >> i )
    {
      array[row][col] = i;
      col++;
      if ( col > row )
	{
	  row++;
	  col = 0;
	}

      /*
      if ( row == MAX_DIM )
	{
	  cout << "ERROR: too many lines of input data!" << endl;
	  exit(1);
	}
      getline( myFile, line );


      row++;
      */
    }
  myFile.close();
}

void rollupArray( int array[MAX_DIM][MAX_DIM] )
{
  int row = 0;
  int col = 0;

  for ( row=MAX_DIM-2; row>=0; row-- )
    {
      for ( col=0; col<=row; col++ )
	{
	  array[row][col] += max( array[row+1][col], array[row+1][col+1] );
	}
    }
}

int main( int argc, char *argv[] )
{
  int array[MAX_DIM][MAX_DIM];

  zeroArray( array );
  readArray( array );
  //  cout << "Starting array..." << endl;
  //  printArray( array );
  rollupArray( array );
  //  cout << "Rolled up array..." << endl;
  //  printArray( array );
  cout << "Sum = " << array[0][0] << endl;
}
