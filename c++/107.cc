//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include <cstdlib>
#include <fstream>
#include <string>
#include <iostream>
#include "graphlib.h"

using std::cout;
using std::endl;
using std::ifstream;
using std::string;
using std::ios;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << __FILE__ << ":" << __LINE__ << ": error: " << error << endl; exit( 1 ); }

void readArray( unsigned int array[40][40] )
{
  ifstream myFile;
  string line;
  int row = 0;
  int col = 0;
  int i = 0;

  myFile.open( "107.data", ios::in );
  while ( myFile >> i )
    {
      array[row][col] = i;
      col++;
      if ( col == 40 )
	{
	  row++;
	  col = 0;
	}
    }
  myFile.close();
}

int main( int argc, char *argv[] )
{
  Graph<char> g;
  Graph<char> minSpan;

  g.addEdge( 'a', 'b', 16 );
  g.addEdge( 'a', 'd', 21 );
  g.addEdge( 'a', 'c', 12 );
  g.addEdge( 'b', 'd', 17 );
  g.addEdge( 'b', 'e', 20 );
  g.addEdge( 'c', 'd', 28 );
  g.addEdge( 'c', 'f', 31 );
  g.addEdge( 'd', 'e', 18 );
  g.addEdge( 'd', 'f', 19 );
  g.addEdge( 'd', 'g', 23 );
  g.addEdge( 'e', 'g', 11 );
  g.addEdge( 'f', 'g', 27 );

  g.print();
  g.reduceToMST( minSpan );
  minSpan.print();

  cout << "-------------------------------------" << endl;

  unsigned int adjacencyMatrix[40][40];
  unsigned int numVertices = 40;
  unsigned int row = 0;
  unsigned int col = 0;

  for ( row = 0; row < numVertices; ++row )
    {
      for ( col = 0; col < numVertices; ++col )
        {
          adjacencyMatrix[row][col] = 0;
        }
    }

  readArray( adjacencyMatrix );
  assert( adjacencyMatrix[0][0]   ==   0, "read array fail" );
  assert( adjacencyMatrix[0][3]   == 427, "read array fail" );
  assert( adjacencyMatrix[0][39]  == 774, "read array fail" );
  assert( adjacencyMatrix[39][0]  == 774, "read array fail" );
  assert( adjacencyMatrix[39][38] == 540, "read array fail" );
  assert( adjacencyMatrix[39][39] ==   0, "read array fail" );

  Graph<char> connected( adjacencyMatrix, numVertices, 'A' );
  connected.print();
  connected.reduceToMST( minSpan );
  minSpan.print();

  return 0;
}
