#include <iostream>
#include <fstream>
#include <string>
#include <vector>

#include "graphlib.h++"

using std::cout;
using std::endl;
using std::ifstream;
using std::string;
using std::ios;
using std::stoi;
using std::vector;

int makeLabel( unsigned int row, unsigned int col, unsigned int cols )
{
  return row * cols + col + 1;
}

void buildGraphFromFile( Graph<int> &graph, const char *filename, int &startVertex, int &endVertex )
{
  unsigned int rows = 0;
  unsigned int cols = 0;
  unsigned int r = 0;
  unsigned int c = 0;
  ifstream myFile;
  string line;
  unsigned int pos = 0;

  startVertex = -1;
  endVertex   = -2;

  myFile.open( filename, ios::in );

  // Read some setup information from the first line
  getline( myFile, line );
  pos = 0;
  cols = 1;
  for ( pos = 0; line[pos] != '\0'; ++pos )
    {
      if ( line[pos] == ',' )
        {
          cols++;
        }
    }
  rows = cols;

  // reset to start of file
  myFile.seekg( 0 );

  vector< vector<unsigned int> > matrix( rows );
  for ( r = 0; r < rows; ++r )
    {
      getline( myFile, line );
      for ( c = 0; c < cols; ++c )
        {
          pos = line.find( "," );
          matrix[r].push_back( stoi( line.substr( 0, pos ) ) );
          line = line.substr( pos + 1 );
        }
    }

  //
  // Start in any cell in the left column
  // and finish in any cell in the right column.
  // Valid moves are: up, down, and right
  //
  for ( r = 0; r < rows; ++r )
    {
      graph.addEdge( startVertex, makeLabel( r, 0, cols ), matrix[r][0] );
      for ( c = 0; c < cols; ++c )
        {
          // If there is a row above this, add an up link
          if ( r > 0 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r-1, c, cols ), matrix[r-1][c] );
            }
          // If there is a row below this, add a down link
          if ( r < rows - 1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r+1, c, cols ), matrix[r+1][c] );
            }
          // If there is a col to the right of this, add a right link
          if ( c < cols - 1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r, c+1, cols ), matrix[r][c+1] );
            }
        }
      graph.addEdge( makeLabel( r, cols - 1, cols ), endVertex, 0 );
    }
}

int main( int argc, char *argv[] )
{
  int startVertex = 0;
  int endVertex   = 0;
  Graph<int> g_mst;

  Graph<int> example( true );    // Create a directed graph
  buildGraphFromFile( example, "082.data.1", startVertex, endVertex );
  example.print();
  cout << "Total # of routes       : " << example.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << example.findLowestWeightRoute( startVertex, endVertex, true ) << endl;
  example.reduceWeightedDCGToMinimalPath( startVertex, endVertex );
  example.print();
  cout << "Total # of routes       : " << example.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << example.findLowestWeightRoute( startVertex, endVertex, false ) << endl;

  Graph<int> g2( true );    // Create a directed graph
  buildGraphFromFile( g2, "082.data.2", startVertex, endVertex );
  g2.print( false );
  cout << "Total # of routes       : " << "####" << endl;
  cout << "Weight of shortest route: " << "####" << endl;
  g2.reduceWeightedDCGToMinimalPath( startVertex, endVertex );
  g2.print();
  cout << "Total # of routes       : " << g2.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << g2.findLowestWeightRoute( startVertex, endVertex, false ) << endl;
}
