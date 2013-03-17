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

void buildGraphFromFile( Graph<int> &graph, const char *filename, int &startVertex, int &endVertex, unsigned long long int &hint )
{
  unsigned int rows = 0;
  unsigned int cols = 0;
  unsigned int r = 0;
  unsigned int c = 0;
  unsigned int firstWeight = 0;
  ifstream myFile;
  string line;
  unsigned int pos = 0;

  startVertex = -1;
  endVertex   = -2;

  myFile.open( filename, ios::in );

  // Read some setup information from the first line
  getline( myFile, line );
  pos = line.find( "," );
  firstWeight = stoi( line.substr( 0, pos ) );
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
  // Start at the top left and finish in
  // the bottom right. Valid moves are:
  // up, down, left, and right
  //
  for ( r = 0; r < rows; ++r )
    {
      for ( c = 0; c < cols; ++c )
        {
          // If there is a row above this, add an up link
          // On the left col there can be no up-links
          // On the right col there can be no up-links
          if ( r > 0 && c > 0 && c != cols-1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r-1, c, cols ), matrix[r-1][c] );
            }
          // If there is a row below this, add a down link
          if ( r < rows - 1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r+1, c, cols ), matrix[r+1][c] );
            }
          // If there is a col to the left of this, add a left link
          // On the top row there can be no left-links
          // On the bottom row there can be no left-links
          if ( c > 0 && r != 0 && r != rows - 1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r, c-1, cols ), matrix[r][c-1] );
            }
          // If there is a col to the right of this, add a right link
          if ( c < cols - 1 )
            {
              graph.addEdge( makeLabel( r, c, cols ), makeLabel( r, c+1, cols ), matrix[r][c+1] );
            }
        }
    }

  graph.addEdge( startVertex, makeLabel( 0, 0, cols ), firstWeight );
  graph.addEdge( makeLabel( rows-1, cols-1, cols ), endVertex, 0 );

  //
  // The shortest path searches work much better
  // if they have an upper bound for the path
  // length to start with. A 'hint', if you will.
  // The 'hint' path we will sum is down the
  // left edge and across the bottom.
  //
  hint = 0;
  for ( r = 0; r < rows; ++r )
    {
      hint += matrix[r][0];
    }
  for ( c = 0; c < cols; ++c )
    {
      hint += matrix[rows-1][c];
    }
}

int main( int argc, char *argv[] )
{
  int startVertex = 0;
  int endVertex   = 0;
  Graph<int> g_mst;
  unsigned long long int hint = 0;

  Graph<int> example( true );    // Create a directed graph
  buildGraphFromFile( example, "083.data.1", startVertex, endVertex, hint );
  example.print();
  cout << "Total # of routes       : " << example.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << example.findLowestWeightRoute( startVertex, endVertex, true, hint ) << endl;
  example.reduceWeightedDCGToMinimalPath( startVertex, endVertex );
  example.print();
  cout << "Total # of routes       : " << example.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << example.findLowestWeightRoute( startVertex, endVertex, false, hint ) << endl;

  Graph<int> g2( true );    // Create a directed graph
  buildGraphFromFile( g2, "083.data.2", startVertex, endVertex, hint );
  g2.print( false );
  cout << "Total # of routes       : " << "####" << endl;
  cout << "Weight of shortest route: " << "####" << endl;
  g2.reduceWeightedDCGToMinimalPath( startVertex, endVertex );
  g2.print();
  cout << "Total # of routes       : " << g2.countRoutes( startVertex, endVertex ) << endl;
  cout << "Weight of shortest route: " << g2.findLowestWeightRoute( startVertex, endVertex, false, hint ) << endl;
}
