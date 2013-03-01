#include <iostream>
#include <cstdlib>
#include "graphlib.h++"

using namespace std;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << __FILE__ << ":" << __LINE__ << ": error: " << error << endl; errorCount++; }

int main( int argc, char *argv[] )
{
  Graph<char> g;

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

  g.eraseVertex( 'f' );
  assert( g.isConnected(), "connected fail" );
  g.print();

  g.eraseEdge( 'c', 'd' );
  assert( g.isConnected(), "connected fail" );
  g.print();

  exit( errorCount );
}
