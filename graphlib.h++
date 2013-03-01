#include <list>
#include <map>
#include <set>

using namespace std;

//
// Graph
//
// All edges have weights. If not used, they are initialized to zero.
// All vertices are named. Duplicate vertices are not allowed.
//

//
// TODO:
// deleteVertex()
// deleteEdge()
// isConnected()
// max edge weight
// min edge weight
// iterator over all edges
// iterator over all edges for a given vertex
// iterator over all vertices
// have dtor call deleteVertex() and deleteEdge() ???
//

#if 1
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ );
#else
#define VALIDATE(obj)
#endif

class Graph
{
public:
  typedef int Label;

  class Edge
  {
  public:
    Label myV1;
    Label myV2;
    int   myWeight;
    Edge( Label v1, Label v2, int weight ) :
      myV1( v1 ),
      myV2( v2 ),
      myWeight( weight )
    {
    }
  };

  typedef list<Edge> Vertex;

  // Create an empty graph
  Graph( bool directed = false );

  Graph( const Graph &other );

  // Create a simple widthxheight grid graph
  // Note that a 2x2 grid has 3x3 vertices
  Graph( unsigned int width, unsigned int height, bool directed = false, int weight = 0 );

  ~Graph()
  {
  }

  Graph operator=( const Graph &rhs );

  void addVertex( Label v1 );

  void addEdge( Label v1, Label v2, int weight = 0 );

  bool hasVertex( const Label v1 ) const;

  // findEdge() is very literal. It only looks for edges from
  // v1 --> v2. Edges from v2 --> v1 don't count.
  bool hasEdge( Label v1, Label v2 ) const;

  unsigned int countRoutes( Label v1, Label v2 ) const;

  unsigned int countRoutes( Label v1, Label v2, set<int> visited ) const;

  set<int> findConnectedVertices( Label v1 ) const;

  bool isConnected( void ) const;

  bool isConnected( Label v1, Label v2 ) const;

  bool findTriangle( Label v1, Label v2, Label &v3 ) const;

  void print( void ) const;

  bool validate( const char *file, int line ) const;

  int numVertices( void ) const
  {
    VALIDATE( this );
    return myVertices.size();
  }

  int numEdges( void ) const
  {
    VALIDATE( this );
    return myNumEdges;
  }

  bool isDirected( void ) const
  {
    VALIDATE( this );
    return myIsDirected;
  }

  bool isSimple( void ) const
  {
    VALIDATE( this );
    return myIsSimple;
  }

  int outDegree( Label v1 ) const
  {
    const Vertex *v = findVertex( v1 );
    return v == NULL ? 0 : v->size();
  }

private:
  Vertex *addVertexGetPtr( Label v1 );
  Vertex *findVertex( const Label v1 );
  const Vertex *findVertex( const Label v1 ) const;

  typedef map< Label, list<Edge> > Vertices;
  Vertices myVertices;
  int  myNumEdges;
  bool myIsDirected;
  bool myIsSimple;
};
