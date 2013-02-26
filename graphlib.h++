#include <set>

//
// TODO:
// deleteVertex()
// deleteEdge()
// isConnected()
// vertex fcns
//   max label
//   min label
//   iterator over all labels
// have dtor call deleteVertex() and deleteEdge() ???
//

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

private:
  Vertex( const Vertex &other );
};

class Edge
{
 public:
  Edge( Vertex *v );

  Vertex *otherVertex;
  Edge *next;

 private:
  Edge( void );
  Edge( const Edge &other );
};

#if 1
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ );
#else
#define VALIDATE(obj)
#endif

class Graph
{
 public:
  // Create an empty graph
  Graph( bool directed = false );

  Graph( const Graph &other );

  // Create a simple widthxheight grid graph
  // Note that a 2x2 grid has 3x3 vertices
  Graph( unsigned int width, unsigned int height, bool directed = false );

  ~Graph();

  Graph operator=( const Graph &rhs );
  Vertex *findVertex( int v ) const;
  Edge *findEdge( int v1, int v2 ) const;
  Vertex *addVertex( int v );
  void addEdge( int v1, int v2 );
  unsigned int countRoutes( int v1, int v2 ) const;
  unsigned int countRoutes( int v1, int v2, set<int> visited ) const;
  set<int> findConnectedVertices( int v ) const;
  bool isConnected( void ) const;
  bool isConnected( int v1, int v2 ) const;
  bool findTriangle( int v1, int v2, int &v3 ) const;
  Vertex *findTriangle( Vertex *v1, Vertex *v2 ) const;

  void print( void ) const;
  bool validate( const char *file, int line ) const;

  int vertexCount( void ) const
  {
    VALIDATE( this );
    return numVertices;
  }

  int edgeCount( void ) const
  {
    VALIDATE( this );
    return numEdges;
  }

  bool directed( void ) const
  {
    VALIDATE( this );
    return isDirected;
  }

  bool simple( void ) const
  {
    VALIDATE( this );
    return isSimple;
  }

 private:
  Vertex *vertices;
  int numVertices;
  int numEdges;
  bool isDirected;
  bool isSimple;
};
