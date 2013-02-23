#include <set>

using namespace std;

class Edge;

class Vertex
{
 public:
  Vertex( int l = -1 );

  Edge *edges;
  int outDegree;
  int label;
  Vertex *next;
};

class Edge
{
 public:
  Edge( Vertex *v );

  Vertex *otherVertex;
  Edge *next;

 private:
  Edge( void );
};

class Graph
{
 public:
  // Create an empty graph
  Graph( bool directed = false );

  // Create a simple widthxheight grid graph
  // Note that a 2x2 grid has 3x3 vertices
  Graph( int width, int height, bool directed = false );

  ~Graph();

  Vertex *findVertex( int v ) const;
  Edge *findEdge( int v1, int v2 ) const;
  void addVertex( int v );
  void addEdge( int v1, int v2 );
  int countRoutes( int v1, int v2 ) const;
  int countRoutes( int v1, int v2, set<int> visited ) const;

  void print( void ) const;
  bool validate( const char *file, int line ) const;

 private:
  Vertex *vertices;
  int numVertices;
  int numEdges;
  bool isDirected;
  bool isSimple;
};
