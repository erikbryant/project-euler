#include <iostream>
#include "graphlib.h++"

using namespace std;

#if 1
#define VALIDATE( obj ) (obj)->validate( __FILE__, __LINE__ );
#else
#define VALIDATE(obj)
#endif

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

Graph::Graph( int width, int height, bool directed ) :
vertices(NULL),
  numVertices(0),
  numEdges(0),
  isDirected(directed),
  isSimple(true)
{
  // A wxh grid has (w+1)x(h+1) vertices
  width++;
  height++;

  int w = 0;
  int h = 0;
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

void Graph::addVertex( int v )
{
  VALIDATE( this );

  Vertex *ptr = new Vertex();
  ptr->label = v;
  ptr->next = vertices;
  vertices = ptr;
  numVertices++;

  VALIDATE( this );
}

void Graph::addEdge( int v1, int v2 )
{
  VALIDATE( this );

  Vertex *x = findVertex( v1 );

  if ( x == NULL )
    {
      addVertex( v1 );
      x = findVertex( v1 );
      if ( x == NULL )
	{
	  cout << "ERROR: Unable to find Vertex: " << v1 << endl;
	  return;
	}
    }

  Vertex *y = findVertex( v2 );

  if ( y == NULL )
    {
      addVertex( v2 );
      y = findVertex( v2 );
      if ( y == NULL )
	{
	  cout << "ERROR: Unable to find Vertex: " << v2 << endl;
	  return;
	}
    }

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
      isDirected = true;
      addEdge( v2, v1 );
      isDirected = false;
    }
  else
    {
      numEdges++;
    }

  VALIDATE( this );
}

int Graph::countRoutes( int v1, int v2 ) const
{
  VALIDATE( this );
  set<int> visited;
  return countRoutes( v1, v2, visited );
}

int Graph::countRoutes( int v1, int v2, set<int> visited ) const
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
  // verify numVertices
  Vertex *ptr = vertices;
  int count = 0;
  while ( ptr != NULL )
    {
      count++;
      ptr = ptr->next;
    }
  if ( count != numVertices )
    {
      cout << "ERROR (" << file << ":" << line << "): Expected " << numVertices << " vertices. Found: " << count << endl;
      return false;
    }

  // verify isDirected
  // verify outDegree
  // verify numEdges
  // verify isSimple

  return true;
}
