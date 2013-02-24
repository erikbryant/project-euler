#include <iostream>
#include <stack>
#include "graphlib.h++"

using namespace std;

Vertex::Vertex( int l ) :
  edges(NULL),
  outDegree(0),
  label(l),
  next(NULL)
{
}

Edge::Edge( Vertex *v ) :
  otherVertex(v),
  next(NULL)
{
}

Graph::Graph( bool directed ) :
vertices(NULL),
  numVertices(0),
  numEdges(0),
  isDirected(directed),
  isSimple(true)
{
  VALIDATE( this );
}

Graph::Graph( unsigned int width, unsigned int height, bool directed ) :
vertices(NULL),
  numVertices(0),
  numEdges(0),
  isDirected(directed),
  isSimple(true)
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
	      addEdge( v, v - 1 );
	    }
	  // Add an edge to the up neighbor
	  if ( h > 0 )
	    {
	      addEdge( v, v - width );
	    }
	}
    }

  VALIDATE( this );
}

Graph::~Graph( void )
{
  Vertex *vptr = vertices;

  while ( vptr != NULL )
    {
      vertices = vertices->next;
      Edge *eptr = vptr->edges;
      while ( eptr != NULL )
	{
	  vptr->edges = vptr->edges->next;
	  delete eptr;
	  eptr = vptr->edges;
	}
      delete vptr;
      vptr = vertices;
    }
}

Vertex *Graph::findVertex( int v ) const
{
  VALIDATE( this );

  Vertex *ptr = vertices;

  while ( ptr != NULL )
    {
      if ( ptr->label == v )
	{
	  return ptr;
	}
      ptr = ptr->next;
    }

  return NULL;
}

Edge *Graph::findEdge( int v1, int v2 ) const
{
  VALIDATE( this );

  Vertex *x = findVertex( v1 );
  Vertex *y = findVertex( v2 );

  if ( x == NULL || y == NULL )
    {
      return NULL;
    }

  Edge *ptr = x->edges;
  while ( ptr != NULL )
    {
      if ( ptr->otherVertex == y )
	{
	  return ptr;
	}
      ptr = ptr->next;
    }

  return NULL;
}

Vertex *Graph::addVertex( int v )
{
  VALIDATE( this );

  Vertex *ptr = findVertex( v );
  if ( ptr == NULL )
    {
      ptr = new Vertex();
      ptr->label = v;
      ptr->next = vertices;
      vertices = ptr;
      numVertices++;
    }

  VALIDATE( this );

  return ptr;
}

void Graph::addEdge( int v1, int v2 )
{
  VALIDATE( this );

  // Make sure both vertices exist
  Vertex *x = addVertex( v1 );
  Vertex *y = addVertex( v2 );

  if ( v1 == v2 )
    {
      isSimple = false;
    }

  if ( isSimple )
    {
      // See if there is already an equivalent edge
      if ( findEdge( v1, v2 ) != NULL )
	{
	  isSimple = false;
	}
    }

  Edge *ptr = x->edges;
  x->edges = new Edge( y );
  x->edges->next = ptr;
  x->outDegree++;

  // If this is not directed, we need to add the
  // reverse of this edge. Unless this is an
  // edge that starts and ends on the same node.
  if ( !isDirected && v1 != v2 )
    {
      Edge *ptr = y->edges;
      y->edges = new Edge( x );
      y->edges->next = ptr;
      y->outDegree++;
    }

  numEdges++;

  VALIDATE( this );
}

unsigned int Graph::countRoutes( int v1, int v2 ) const
{
  VALIDATE( this );
  set<int> visited;
  return countRoutes( v1, v2, visited );
}

unsigned int Graph::countRoutes( int v1, int v2, set<int> visited ) const
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
int Graph::countRoutesRightAndDown( int width, int height ) const
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

set<int> Graph::findConnectedVertices( int v ) const
{
  set<int> connected;
  stack<Vertex *> toVisit;
  Vertex *vptr = NULL;
  Edge   *eptr = NULL;

  // Initialize the list of vertices to visit with 'v'
  vptr = findVertex( v );
  if ( vptr == NULL )
    {
      return connected;
    }
  toVisit.push( vptr );

  while ( !toVisit.empty() )
    {
      // Take the head from the list, find all its connected
      // vertices and add them to the list (unless they are
      // already in there)
      vptr = toVisit.top();
      toVisit.pop();
      if ( vptr == NULL || connected.count( vptr->label ) != 0 )
	{
	  continue;
	}
      connected.insert( vptr->label );
      eptr = vptr->edges;
      while ( eptr != NULL )
	{
	  // Add vertex to list to vist
	  toVisit.push( eptr->otherVertex );
	  eptr = eptr->next;
	}
    }

  return connected;
}

bool Graph::isConnected( void ) const
{
  if ( vertices == NULL || numVertices <= 1 )
    {
      return true;
    }

  set<int> all;
  set<int> connected;
  Vertex *vptr = NULL;

  // Load all of the vertices into a set
  vptr = vertices;
  while ( vptr != NULL )
    {
      all.insert( vptr->label );
      vptr = vptr->next;
    }

  // Grab an arbitrary vertex and find all
  // that are connected to it
  vptr = vertices;
  connected = findConnectedVertices( vptr->label );

  // If 'all' and 'connected' are equal then
  // this is a connected graph.
  return all == connected;
}

bool Graph::isConnected( int v1, int v2 ) const
{
  if ( v1 == v2 )
    {
      return true;
    }

  // Find all vertices that are connected to v1
  set<int> connected = findConnectedVertices( v1 );

  return connected.count( v2 ) != 0;
}

void Graph::print( void ) const
{
  VALIDATE( this );

  Vertex *vptr = vertices;
  Edge   *eptr = NULL;

  cout << "Vertices    : " << numVertices << endl;
  cout << "Edges       : " << numEdges << endl;
  cout << "isDirected? : " << isDirected << endl;
  cout << "isSimple?   : " << isSimple << endl;

  while ( vptr != NULL )
    {
      cout << "Vertex " << vptr->label << ":";
      eptr = vptr->edges;
      while ( eptr != NULL )
	{
	  cout << " -> " << eptr->otherVertex->label;
	  eptr = eptr->next;
	}
      cout << endl;
      vptr = vptr->next;
    }

  cout << endl;
}

bool Graph::validate( const char *file, int line ) const
{
  Vertex *vptr = NULL;
  Edge   *eptr = NULL;
  int count = 0;

  // verify numVertices
  vptr = vertices;
  while ( vptr != NULL )
    {
      count++;
      vptr = vptr->next;
    }
  if ( count != numVertices )
    {
      cout << "ERROR (" << file << ":" << line << "): Expected " << numVertices << " vertices. Found: " << count << endl;
      return false;
    }

  // verify isDirected

  // verify outDegree
  vptr =  vertices;
  while ( vptr != NULL )
    {
      count = 0;
      eptr = vptr->edges;
      while ( eptr != NULL )
	{
	  count++;
	  eptr = eptr->next;
	}
      if ( count != vptr->outDegree )
	{
	  cout << "ERROR (" << file << ":" << line << "): OutDegree mismatch. Expected " << vptr->outDegree << " edges. Found: " << count << endl;
	  return false;
	}
      vptr = vptr->next;
    }

  // verify numEdges
  vptr =  vertices;
  count = 0;
  while ( vptr != NULL )
    {
      count += vptr->outDegree;
      vptr = vptr->next;
    }
  if ( isSimple )
    {
      if ( isDirected )
	{
	  if ( count != numEdges )
	    {
	      cout << "ERROR (" << file << ":" << line << "): numEdges (directed) mismatch. Expected " << numEdges << " edges. Found: " << count << endl;
	      return false;
	    }
	}
      else
	{
	  if ( count != numEdges * 2 )
	    {
	      cout << "ERROR (" << file << ":" << line << "): numEdges (!directed) mismatch. Expected " << numEdges * 2 << " edges. Found: " << count << endl;
	      return false;
	    }
	}
    }
  else
    {
      // TODO: Do something smart
      if ( isDirected )
	{
	  // numEdges = sum of all edges + 2 * #self-referential
	}
      else
	{
	  // numEdges = sum of all edges
	}
    }

  // verify isSimple
  //   no self-referential
  //   at most one arc from any V1 to any other V2

  return true;
}
