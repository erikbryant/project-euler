#include <list>
#include <map>
#include <set>
#include <stack>

using namespace std;

//
// Graph
//
// All edges have weights. If not used, they are initialized to zero.
// All vertices are named. Duplicate vertices are not allowed.
// Vertex names can be of any type (that's the 'Label' parameter
// in the class template).
//

#if 1
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ );
#else
#define VALIDATE(obj)
#endif

template <typename Label>
class Graph
{
public:
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

  void eraseVertex( Label v1 );

  void eraseEdge( Label v1, Label v2 );

  int sumWeights( void ) const;

  unsigned int countRoutes( Label v1, Label v2 ) const;

  unsigned int countRoutes( Label v1, Label v2, set<Label> visited ) const;

  set<Label> findConnectedVertices( Label v1 ) const;

  bool isConnected( void ) const;

  bool isConnected( Label v1, Label v2 ) const;

  bool findTriangle( Label v1, Label v2, Label &v3 ) const;

  void print( void ) const;

  bool validate( const char *file, int line ) const;

  unsigned int numVertices( void ) const
  {
    VALIDATE( this );
    return myVertices.size();
  }

  unsigned int numEdges( void ) const
  {
    unsigned int edges = 0;

    VALIDATE( this );

    typename Vertices::const_iterator v_it;
    for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
      {
	typename list<Edge>::const_iterator e_it;
	for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	  {
	    edges++;
	  }
      }

    return isDirected() ? edges : edges / 2;
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
  bool myIsDirected;
  bool myIsSimple;
};

template <typename Label>
Graph<Label>::Graph( bool directed ) :
  myVertices(),
  myIsDirected( directed ),
  myIsSimple( true )
{
  VALIDATE( this );
}

template <typename Label>
Graph<Label>::Graph( const Graph &other ) :
  myVertices(),
  myIsDirected( other.myIsDirected ),
  myIsSimple( other.myIsSimple )
{
  // Loop through other, copying each vertex
  typename Vertices::const_iterator v_it;
  for ( v_it=other.myVertices.begin(); v_it!=other.myVertices.end(); ++v_it )
    {
      addVertex( v_it->first );
    }
  // Once we have the vertices, copy each edge list
  for ( v_it=other.myVertices.begin(); v_it!=other.myVertices.end(); ++v_it )
    {
      typename Vertex::const_iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  // Don't use addEdge(). It gets confused.
	  addVertexGetPtr( v_it->first )->push_front( Edge( v_it->first, e_it->myV2, e_it->myWeight ) );
	}
    }

  VALIDATE( this );
}

template <typename Label>
Graph<Label>::Graph( unsigned int width, unsigned int height, bool directed, int weight ) :
  myVertices(),
  myIsDirected( directed ),
  myIsSimple( true )
{
  // A wxh grid has (w+1)x(h+1) vertices
  width++;
  height++;

  unsigned int w = 0;
  unsigned int h = 0;
  int v = 0;

  for ( h=0; h<height; h++ )
    {
      for ( w=0; w<width; w++ )
	{
	  v = (h * width) + w;
	  addVertex( v );
	  // Add an edge to the left neighbor
	  if ( w > 0 )
	    {
	      addEdge( v, v - 1, weight );
	    }
	  // Add an edge to the up neighbor
	  if ( h > 0 )
	    {
	      addEdge( v, v - width, weight );
	    }
	}
    }

  VALIDATE( this );
}

template <typename Label>
typename Graph<Label>::Vertex *Graph<Label>::findVertex( const Label v1 )
{
  VALIDATE( this );
  typename Vertices::iterator it = myVertices.find( v1 );
  return it != myVertices.end() ? &(it->second) : NULL;
}

template <typename Label>
const typename Graph<Label>::Vertex *Graph<Label>::findVertex( const Label v1 ) const
{
  VALIDATE( this );
  typename Vertices::const_iterator it = myVertices.find( v1 );
  return it != myVertices.end() ? &(it->second) : NULL;
}

template <typename Label>
bool Graph<Label>::hasVertex( const Label v1 ) const
{
  VALIDATE( this );
  typename Vertices::const_iterator it = myVertices.find( v1 );
  return it != myVertices.end();
}

template <typename Label>
bool Graph<Label>::hasEdge( Label v1, Label v2 ) const
{
  VALIDATE( this );

  const Vertex *v = findVertex( v1 );

  if ( v == NULL || !hasVertex( v2 ) )
    {
      return NULL;
    }

  typename Vertex::const_iterator it;
  for ( it=v->begin(); it!=v->end(); ++it )
    {
      if ( it->myV2 == v2 )
	{
	  return &(*it);
	}
    }

  return NULL;
}

template <typename Label>
void Graph<Label>::addVertex( Label v1 )
{
  VALIDATE( this );

  if ( !hasVertex( v1 ) )
    {
      Vertex v;
      myVertices.insert( pair<Label, Vertex>( v1, v ) );
    }

  VALIDATE( this );
}

template <typename Label>
typename Graph<Label>::Vertex *Graph<Label>::addVertexGetPtr( Label v1 )
{
  VALIDATE( this );

  Vertex *ptr = findVertex( v1 );
  if ( ptr == NULL )
    {
      Vertex v;
      myVertices.insert( pair<Label, Vertex>( v1, v ) );
      ptr = findVertex( v1 );
    }

  VALIDATE( this );

  return ptr;
}

template <typename Label>
void Graph<Label>::addEdge( Label v1, Label v2, int weight )
{
  VALIDATE( this );

  if ( v1 == v2 )
    {
      myIsSimple = false;
    }

  if ( isSimple() )
    {
      // See if there is already an equivalent edge
      if ( hasEdge( v1, v2 ) )
	{
	  myIsSimple = false;
	}
    }

  addVertexGetPtr( v1 )->push_front( Edge( v1, v2, weight ) );

  // If this is not directed then edges go both ways, so we
  // also need to add the reverse of this edge.
  if ( !isDirected() )
    {
      addVertexGetPtr( v2 )->push_front( Edge( v2, v1, weight ) );
    }

  VALIDATE( this );
}

template <typename Label>
void Graph<Label>::eraseVertex( Label v1 )
{
  VALIDATE( this );

  // Remove all edges that point to this vertex
  typename Vertices::iterator v_it;
  for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
    {
      typename list<Edge>::iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  if ( e_it->myV2 == v1 )
	    {
	      e_it = v_it->second.erase( e_it );
	    }
	}
    }

  // Remove the vertex itself
  myVertices.erase( v1 );

  VALIDATE( this );
}

template <typename Label>
void Graph<Label>::eraseEdge( Label v1, Label v2 )
{
  VALIDATE( this );

  typename Vertices::iterator v_it;
  for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
    {
      typename list<Edge>::iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  if ( isDirected() )
	    {
	      if ( e_it->myV1 == v1 && e_it->myV2 == v2 )
		{
		  e_it = v_it->second.erase( e_it );
		}
	    }
	  else
	    {
	      if ( ( e_it->myV1 == v1 && e_it->myV2 == v2 ) ||
		   ( e_it->myV1 == v2 && e_it->myV2 == v1 ) )
		{
		  e_it = v_it->second.erase( e_it );
		}
	    }
	}
    }

  VALIDATE( this );
}

template <typename Label>
int Graph<Label>::sumWeights( void ) const
{
  int sum = 0;

  typename Vertices::const_iterator v_it;
  for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
    {
      typename list<Edge>::const_iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  sum += e_it->myWeight;
	}
    }

  return isDirected() ? sum : sum / 2;
}

#if 0
template <typename Label>
unsigned int Graph<Label>::countRoutes( Label v1, Label v2 )
{
  VALIDATE( this );
  set<Label> visited;
  return countRoutes( v1, v2, visited );
}

template <typename Label>
unsigned int Graph<Label>::countRoutes( Label v1, Label v2, set<Label> visited )
{
  if ( v1 == v2 )
    {
      return 1;
    }

  visited.insert( v1 );

  Vertex *start = findVertex( v1 );
  Edge   *ptr   = start->edges;
  int total = 0;

  while ( ptr != NULL )
    {
      if ( visited.count( ptr->otherVertex->label ) == 0 )
	{
	  total += countRoutes( ptr->otherVertex->label, v2, visited );
	}
      ptr = ptr->next;
    }

  return total;
}

/*
template <typename Label>
int Graph<Label>::countRoutesRightAndDown( int width, int height ) const
{
  if ( width == 1 )
    {
      return height + 1;
    }

  if ( height == 1 )
    {
      return width + 1;
    }

  return countRoutesRightAndDown( width - 1, height ) +
    countRoutesRightAndDown( width, height - 1 );
}
*/
#endif

template <typename Label>
set<Label> Graph<Label>::findConnectedVertices( Label v1 ) const
{
  set<Label> connected;
  stack<Label> toVisit;

  if ( !hasVertex( v1 ) )
    {
      return connected;
    }

  // Initialize the list of vertices to visit with 'v1'
  toVisit.push( v1 );

  while ( !toVisit.empty() )
    {
      // Take the head from the list, find all its connected
      // vertices and add them to the list (unless they are
      // already in there)
      Label l = toVisit.top();
      toVisit.pop();
      if ( connected.count( l ) != 0 )
	{
	  continue;
	}
      connected.insert( l );
      const Vertex *vptr = findVertex( l );
      typename Vertex::const_iterator it;
      for ( it=vptr->begin(); it!=vptr->end(); ++it )
	{
	  // Add vertex to list to vist
	  toVisit.push( it->myV2 );
	}
    }

  return connected;
}

template <typename Label>
bool Graph<Label>::isConnected( void ) const
{
  if ( numVertices() <= 1 )
    {
      return true;
    }

  set<Label> all;
  set<Label> connected;

  // Load all of the vertices into a set
  typename Vertices::const_iterator it;

  for ( it=myVertices.begin(); it!=myVertices.end(); ++it )
    {
      all.insert( it->first );
    }

  // Grab an arbitrary vertex and find all
  // that are connected to it
  connected = findConnectedVertices( myVertices.begin()->first );

  // If 'all' and 'connected' are equal then
  // this is a connected graph.
  return all == connected;
}

template <typename Label>
bool Graph<Label>::isConnected( Label v1, Label v2 ) const
{
  if ( v1 == v2 )
    {
      return true;
    }

  // Find all vertices that are connected to v1
  set<Label> connected = findConnectedVertices( v1 );

  return connected.count( v2 ) != 0;
}

template <typename Label>
bool Graph<Label>::findTriangle( Label v1, Label v2, Label &v3 ) const
{
  if ( !hasVertex( v1 ) ||
       !hasVertex( v2 ) ||
       !hasEdge( v1, v2 ) )
    {
      return false;
    }

  //
  // Look for an edge that connects v1 <-> ?? <-> v2
  //
  const Vertex *vptr = findVertex( v1 );
  typename Vertex::const_iterator v_it;
  for ( v_it=vptr->begin(); v_it!=vptr->end(); ++v_it )
    {
      if ( hasEdge( v_it->myV2, v2 ) )
	{
	  v3 = v_it->myV2;
	  return true;
	}
    }

  //
  // Look for an edge that connects v2 <-> ?? <-> v1
  //
  vptr = findVertex( v2 );
  for ( v_it=vptr->begin(); v_it!=vptr->end(); ++v_it )
    {
      if ( hasEdge( v_it->myV2, v1 ) )
	{
	  v3 = v_it->myV2;
	  return true;
	}
    }

  return false;
}

template <typename Label>
void Graph<Label>::print( void ) const
{
  unsigned int weight = 0;

  VALIDATE( this );

  cout << "Vertices    : " << numVertices() << endl;
  cout << "Edges       : " << numEdges() << endl;
  cout << "isDirected? : " << isDirected() << endl;
  cout << "isSimple?   : " << isSimple() << endl;

  typename Vertices::const_iterator v_it;
  for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
    {
      cout << "Vertex " << v_it->first << "(" << v_it->second.size() << ") :";
      typename list<Edge>::const_iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  cout << " --" << e_it->myWeight << "--> " << e_it->myV2;
	  weight += e_it->myWeight;
	}
      cout << endl;
    }
  cout << "total weight : " << (isDirected() ? weight : weight / 2) << endl;
  cout << endl;
}

template <typename Label>
bool Graph<Label>::validate( const char *file, int line ) const
{
  // verify isDirected

  // verify the edges are attached to the right vertices
  // i.e., edge.myV1 == vertex label
  typename Vertices::const_iterator v_it;
  for ( v_it=myVertices.begin(); v_it!=myVertices.end(); ++v_it )
    {
      typename Vertex::const_iterator e_it;
      for ( e_it=v_it->second.begin(); e_it!=v_it->second.end(); ++e_it )
	{
	  if ( v_it->first != e_it->myV1 )
	    {
	      cout << file << ":" << line << ": error: A vertex ( " << v_it->first<< " ) has an edge that does not belong to it ( " << e_it->myV1 << ", " << e_it->myV2 << ", " << e_it->myWeight << " )" << endl;
	      return false;
	    }
	}
    }

  // verify the edges refer to vertices that exist
  // i.e., Graph contains a vertex with label edge.myV2
  // TODO

  // verify isSimple
  //   no self-referential
  //   at most one arc from any V1 to any other V2
  // TODO

  // if myIsDirected == true then verify that for
  // every edge V1 --> V2 there is a corresponding
  // edge V2 --> V1
  // TODO

  return true;
}
