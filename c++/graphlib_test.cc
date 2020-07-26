//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include <iostream>
#include <cstdlib>
#include "graphlib.h"

using namespace std;

unsigned int errorCount = 0;

#define assert( cond, error ) if ( !(cond) ) { cout << __FILE__ << ":" << __LINE__ << ": error: " << error << endl; errorCount++; }

int main( int argc, char *argv[] )
{

  //
  //
  //
  Graph<char> d( true );
  d.addEdge( 'a', 'b', 0 );
  assert( d.numEdges() == 1, "edge count fail" );
  assert( d.numVertices() == 2, "vertex count fail" );
  assert( d.isSimple(), "isSimple fail" );
  assert( d.isDirected() == true, "isDirected fail" );
  assert( d.hasVertex( 'b' ) == true, "has vertex fail" );
  assert( d.hasEdge( 'a', 'c' ) == false, "has edge fail" );
  assert( d.isConnected(), "connected fail" );



  //
  // Int promotion ctor
  // (There is no void ctor and the int promotion has a default value)
  //

  Graph<int> g;
  assert( g.numEdges() == 0, "edge count fail" );
  assert( g.numVertices() == 0, "vertex count fail" );
  assert( g.isSimple(), "isSimple fail" );
  assert( g.isDirected() == false, "isDirected fail" );
  assert( g.hasVertex( 99 ) == false, "has vertex fail" );
  assert( g.hasEdge( 12, 67 ) == false, "has edge fail" );
  assert( g.isConnected(), "connected fail" );
  assert( g.isConnected( 3, 5 ) == false, "connected fail" );



  //
  // Copy ctor
  //

  Graph<int> g_copy( g );
  assert( g_copy.numEdges() == 0, "edge count fail" );
  assert( g_copy.numVertices() == 0, "vertex count fail" );
  assert( g_copy.isSimple(), "isSimple fail" );
  assert( g_copy.isDirected() == false, "isDirected fail" );
  assert( g_copy.hasVertex( 99 ) == false, "has vertex fail" );
  assert( g_copy.hasEdge( 12, 67 ) == false, "has edge fail" );
  assert( g_copy.isConnected(), "connected fail" );
  assert( g_copy.isConnected( 3, 5 ) == false, "connected fail" );



  //
  // addVertex
  //

  g.addVertex( 3 );
  assert( g.numEdges() == 0, "edge count fail" );
  assert( g.numVertices() == 1, "vertex count fail" );
  assert( g.isSimple(), "isSimple fail" );
  assert( g.isDirected() == false, "isDirected fail" );
  assert( g.hasVertex( 3 ), "has vertex fail" );
  assert( g.hasEdge( 3, 67 ) == false, "has edge fail" );
  assert( g.isConnected(), "connected fail" );
  assert( g.isConnected( 3, 3 ), "connected fail" );
  assert( g.isConnected( 3, 5 ) == false, "connected fail" );

  g.addVertex( 3 );
  assert( g.numEdges() == 0, "edge count fail" );
  assert( g.numVertices() == 1, "vertex count fail" );
  assert( g.isSimple(), "isSimple fail" );
  assert( g.isDirected() == false, "isDirected fail" );
  assert( g.hasVertex( 3 ), "has vertex fail" );
  assert( g.hasEdge( 3, 67 ) == false, "has edge fail" );
  assert( g.isConnected(), "connected fail" );
  assert( g.isConnected( 3, 3 ), "connected fail" );
  assert( g.isConnected( 3, 5 ) == false, "connected fail" );



  //
  // addEdge
  //

  g.addEdge( 3, 5 );
  assert( g.numEdges() == 1, "edge count fail" );
  assert( g.numVertices() == 2, "vertex count fail" );
  assert( g.isSimple(), "isSimple fail" );
  assert( g.isDirected() == false, "isDirected fail" );
  assert( g.hasVertex( 3 ), "has vertex fail" );
  assert( g.hasVertex( 5 ), "has vertex fail" );
  assert( g.hasEdge( 3, 5 ), "has edge fail" );
  assert( g.isConnected(), "connected fail" );
  assert( g.isConnected( 3, 3 ), "connected fail" );
  assert( g.isConnected( 5, 5 ), "connected fail" );
  assert( g.isConnected( 3, 5 ), "connected fail" );
  assert( g.isConnected( 5, 3 ), "connected fail" );
  assert( g.isConnected( 9, 5 ) == false, "connected fail" );

  g.addEdge( 3, 5 );
  assert( g.numEdges() == 2, "edge count fail" );
  assert( g.numVertices() == 2, "vertex count fail" );
  assert( !g.isSimple(), "isSimple fail" );
  assert( g.isDirected() == false, "isDirected fail" );
  assert( g.hasVertex( 3 ), "has vertex fail" );
  assert( g.hasVertex( 5 ), "has vertex fail" );
  assert( g.hasEdge( 3, 5 ), "has edge fail" );
  assert( g.isConnected(), "connected fail" );
  assert( g.isConnected( 3, 3 ), "connected fail" );
  assert( g.isConnected( 5, 5 ), "connected fail" );
  assert( g.isConnected( 3, 5 ), "connected fail" );
  assert( g.isConnected( 5, 3 ), "connected fail" );
  assert( g.isConnected( 9, 5 ) == false, "connected fail" );



  //
  // Copy ctor of complex graph
  //
  Graph<int> g_bigcopy( g );
  assert( g_bigcopy.numEdges() == 2, "edge count fail" );
  assert( g_bigcopy.numVertices() == 2, "vertex count fail" );
  assert( !g_bigcopy.isSimple(), "isSimple fail" );
  assert( g_bigcopy.isDirected() == false, "isDirected fail" );
  assert( g_bigcopy.hasVertex( 3 ), "has vertex fail" );
  assert( g_bigcopy.hasVertex( 5 ), "has vertex fail" );
  assert( g_bigcopy.hasEdge( 3, 5 ), "has edge fail" );
  assert( g_bigcopy.isConnected(), "connected fail" );
  assert( g_bigcopy.isConnected( 3, 3 ), "connected fail" );
  assert( g_bigcopy.isConnected( 5, 5 ), "connected fail" );
  assert( g_bigcopy.isConnected( 3, 5 ), "connected fail" );
  assert( g_bigcopy.isConnected( 5, 3 ), "connected fail" );
  assert( g_bigcopy.isConnected( 9, 5 ) == false, "connected fail" );



  //
  // dtor
  //

  Graph<int> *gptr = new Graph<int>();
  gptr->addEdge( 10, 20 );
  delete gptr;



  Graph<int> g2;

  assert( g2.numEdges() == 0, "edge count fail" );
  assert( g2.numVertices() == 0, "vertex count fail" );
  assert( g2.isSimple(), "isSimple fail" );
  assert( g2.isDirected() == false, "isDirected fail" );
  assert( g2.hasVertex( 14 ) == false, "has vertex fail" );
  assert( g2.hasEdge( 14, 14 ) == false, "has edge fail" );
  assert( g2.hasEdge( 9, 34 ) == false, "has edge fail" );
  assert( g2.isConnected(), "connected fail" );
  assert( g2.isConnected( 0, 3 ) == false, "connected fail" );

  g2.addEdge( 9, 9 );
  assert( g2.numEdges() == 1, "edge count fail" );
  assert( g2.numVertices() == 1, "vertex count fail" );
  assert( !g2.isSimple(), "isSimple fail" );
  assert( g2.isDirected() == false, "isDirected fail" );
  assert( g2.hasVertex( 9 ), "has vertex fail" );
  assert( g2.hasEdge( 9, 9 ), "has edge fail" );
  assert( g2.hasEdge( 9, 12 ) == false, "has edge fail" );
  assert( g2.isConnected(), "connected fail" );
  assert( g2.isConnected( 9, 9 ), "connected fail" );
  assert( g2.isConnected( 9, 3 ) == false, "connected fail" );

  unsigned int w = 10;
  unsigned int h = 10;
  Graph<int> grid( w, h );
  assert( grid.numEdges() == w * (h + 1) + h * (w + 1), "edge count fail" );
  assert( grid.numVertices() == (w + 1) * (h + 1), "vertex count fail" );
  assert( grid.isSimple(), "isSimple fail" );
  assert( grid.isDirected() == false, "isDirected fail" );
  assert( grid.hasVertex( 0 ), "has vertex fail" );
  assert( grid.hasEdge( 0, w + 1 ), "has edge fail" );
  assert( grid.hasEdge( 0, 2 ) == false, "has edge fail" );
  assert( grid.isConnected(), "connected fail" );
  unsigned int i = 0;
  for ( i=0; i<=(w + 1) * (h + 1) - 1; i++ )
    {
      assert( grid.isConnected( 0, i ), "connected fail" );
    }

  //
  // BFS
  //
  grid.BFS( 0 );



  w = 1;
  h = 1;
  Graph<int> grid1( w, h );
  assert( grid1.numEdges() == w * (h + 1) + h * (w + 1), "edge count fail" );
  assert( grid1.numVertices() == (w + 1) * (h + 1), "vertex count fail" );
  assert( grid1.isSimple(), "isSimple fail" );
  assert( grid1.isDirected() == false, "isDirected fail" );
  //  assert( grid1.countRoutes( 0, ((w + 1) * (h + 1) - 1) ) == 2, "count routes fail" );
  assert( grid1.hasVertex( 1 ), "has vertex fail" );
  assert( grid1.hasEdge( 1, w - 1 ), "has edge fail" );
  assert( grid1.hasEdge( 0, ((w + 1) * (h + 1) - 1) ) == false, "has edge fail" );
  assert( grid1.isConnected(), "connected fail" );
  for ( i=0; i<=(w + 1) * (h + 1) - 1; i++ )
    {
      assert( grid1.isConnected( 0, i ), "connected fail" );
    }

  w = 2;
  h = 2;
  Graph<int> grid2( w, h );
  assert( grid2.numEdges() == w * (h + 1) + h * (w + 1), "edge count fail" );
  assert( grid2.numVertices() == (w + 1) * (h + 1), "vertex count fail" );
  assert( grid2.isSimple(), "isSimple fail" );
  assert( grid2.isDirected() == false, "isDirected fail" );
  //  assert( grid2.countRoutes( 0, ((w + 1) * (h + 1) - 1) ) == 12, "count routes fail" );
  assert( grid2.hasVertex( w ), "has vertex fail" );
  assert( grid2.hasEdge( w + 1, 0 ), "has edge fail" );
  assert( grid2.hasEdge( 0, ((w + 1) * (h + 1) - 1) ) == false, "has edge fail" );
  assert( grid2.isConnected(), "connected fail" );
  for ( i=0; i<=(w + 1) * (h + 1) - 1; i++ )
    {
      assert( grid2.isConnected( 0, i ), "connected fail" );
    }



  //
  // findTriangle
  //

  Graph<int> network;
  int v3 = 0;

  network.addEdge( 1, 2 );
  network.addEdge( 2, 3 );
  assert( network.findTriangle( 1, 3, v3 ) == false, "findTriangle fail" );
  assert( v3 == 0, "findTriangle fail" );   // v3 must be unchanged

  network.addEdge( 3, 1 );
  assert( network.findTriangle( 1, 2, v3 ), "findTriangle fail" );
  assert( v3 == 3, "findTriangle fail" );
  assert( network.findTriangle( 2, 3, v3 ), "findTriangle fail" );
  assert( v3 == 1, "findTriangle fail" );
  assert( network.findTriangle( 1, 3, v3 ), "findTriangle fail" );
  assert( v3 == 2, "findTriangle fail" );

  network.addEdge( 3, 4 );
  network.addEdge( 4, 1 );
  assert( network.findTriangle( 1, 3, v3 ), "findTriangle fail" );
  assert( v3 == 4, "findTriangle fail" );

  // Three possible triangles
  network.addEdge( 3, 5 );
  network.addEdge( 1, 5 );
  assert( network.findTriangle( 1, 3, v3 ), "findTriangle fail" );

  // Vertices that don't exist
  assert( network.findTriangle( 99, 88, v3 ) == false, "findTriangle fail" );



  //
  // Weights
  //
  Graph<int> weighted;

  weighted.addEdge( 12, 24, 100 );
  assert( weighted.sumWeights() == 100, "sum weights fail" );
  weighted.addEdge( 1, 2, 300 );
  assert( weighted.sumWeights() == 100 + 300, "sum weights fail" );
  weighted.addEdge( 5, 5, 250 );
  assert( weighted.sumWeights() == 100 + 300 + 250, "sum weights fail" );
  weighted.addEdge( 2, 2, 125 );
  assert( weighted.sumWeights() == 100 + 300 + 250 + 125, "sum weights fail" );



  //
  // eraseEdge
  //
  Graph<int> eraseMe;
  eraseMe.addEdge( 1, 2, 0 );
  eraseMe.addEdge( 2, 3, 0 );
  eraseMe.addEdge( 3, 4, 0 );
  assert( eraseMe.numEdges() == 3, "numEdges fail" );
  assert( eraseMe.numVertices() == 4, "numVertices fail" );

  eraseMe.eraseEdge( 2, 3 );
  assert( eraseMe.numEdges() == 2, "numEdges fail" );
  assert( eraseMe.numVertices() == 4, "numVertices fail" );



  //
  // eraseVertex
  //
  eraseMe.eraseVertex( 2 );
  assert( eraseMe.numEdges() == 1, "numEdges fail" );
  assert( eraseMe.numVertices() == 3, "numVertices fail" );



  //
  // erase
  //
  eraseMe.erase();
  assert( eraseMe.numEdges() == 0, "numEdges fail" );
  assert( eraseMe.numVertices() == 0, "numVertices fail" );

  exit( errorCount );
}
