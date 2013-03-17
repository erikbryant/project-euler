//
// Copyright Erik Bryant (erikbryantology@gmail.com)
// GPLv2 http://www.gnu.org/licenses/gpl-2.0.html
//

#pragma once

#include <list>
#include <map>
#include <set>
#include <stack>
#include <vector>
#include <iostream>

using std::cout;
using std::endl;
using std::list;
using std::set;
using std::map;
using std::pair;
using std::stack;
using std::vector;
using std::min;

//
// Graph
//
// All edges have weights. If not used, they are initialized to zero.
// All vertices are named. Duplicate vertices are not allowed.
// Vertex names can be of any type (that's the 'Label' parameter
// in the class template).
//

#if 0
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
    unsigned int   myWeight;
    Edge( Label v1, Label v2, unsigned int weight ) :
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
  Graph( unsigned int width,
         unsigned int height,
         bool directed = false,
         unsigned int weight = 0 );

  Graph( unsigned int adjacencyMatrix[40][40], unsigned int numVertices, const Label &startLabel, bool directed = false );

  ~Graph()
  {
  }

  Graph operator=( const Graph &rhs );  

  void addVertex( Label v1 );

  void addEdge( Label v1, Label v2, unsigned int weight = 0 );

  bool hasVertex( const Label v1 ) const;

  // hasEdge() is very literal. It only looks for edges from
  // v1 --> v2. Edges from v2 --> v1 don't count.
  bool hasEdge( Label v1, Label v2 ) const;

  unsigned int getEdgeWeight( Label v1, Label v2 ) const;

  void setEdgeWeight( Label v1, Label v2, unsigned int weight );
  void incEdgeWeight( Label v1, Label v2, unsigned int delta );
  void decEdgeWeight( Label v1, Label v2, unsigned int delta );

  void eraseVertex( Label v1 );

  void eraseEdge( Label v1, Label v2 );

  void eraseAllEdgesOutOf( Label v1 );

  void eraseAllEdgesInto( Label v1 );

  void eraseDuplicateEdges( Label v1 );

  void erase( void );

  unsigned long long int sumWeights( void ) const;

  unsigned long long int countRoutes( Label v1, Label v2 ) const;

  unsigned long long int findLowestWeightRoute( Label v1, Label v2, bool cyclic = true, unsigned long long int hint = 0xFFFFFFFFFFFFFFFF ) const;

  // Find the full set of vertices that are connected (no matter how remotely) to v1
  set<Label> findConnectedVertices( Label v1 ) const;

  // Find the set of vertices that have edges
  // directly into the given vertex.
  set<Label> findEdgesInto( Label v1 ) const;

  // Find the set of vertices that have edges
  // directly out of the given vertex.
  set<Label> findEdgesOutOf( Label v1 ) const;

  // If all of the edges into a vertex having matching edges
  // out of that vertex then that vertex can be removed and
  // the route cost through the graph will be unaffected
  void removeSymmetricallyConnectedVertices( void );

  // Given a weighted DAG, a start vertex, and a terminus vertex,
  // remove all vertices between start and terminus, rolling the
  // minimal path weight up into a single edge between start and
  // terminus.
  void reduceWeightedDAGToMinimalPath( const Label start, const Label terminus );

  // Given a weighted DCG, a start vertex, and a terminus vertex,
  // remove all vertices between start and terminus, rolling the
  // minimal path weight up into a single edge between start and
  // terminus.
  void reduceWeightedDCGToMinimalPath( const Label start, const Label terminus );

  // Is the entire graph a connected graph?
  // Allow the caller to provide a starting node
  // critical for directed, acyclic graphs).
  bool isConnected( const Label *startV = NULL ) const;

  // Is there some path that connects v1 to v2?
  bool isConnected( Label v1, Label v2 ) const;

  // Given two vertices that are directly connected, find whether
  // there is a third that makes up a trianlge with them.
  bool findTriangle( Label v1, Label v2, Label &v3 ) const;

  // Make a cut through a connected graph. Some vertices will be
  // on one side and some will be on the other. This function finds
  // the lowest-weighted edge that connects the two, otherwise,
  // disjointed sets.
  bool findLightestEdge( const set<Label> &s1, const set<Label> &s2, Label &v1, Label &v2, unsigned int &minWeight ) const;

  void reduceToMST( Graph<Label> &mst, const Label *startV = NULL ) const;

  void print( bool verbose = true ) const;

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
    for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
      {
        typename list<Edge>::const_iterator e_it;
        for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
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

  unsigned int outDegree( Label v1 ) const
  {
    const Vertex *v = findVertex( v1 );
    return v == NULL ? 0 : v->size();
  }

private:
  Vertex *addVertexGetPtr( Label v1 );
  Vertex *findVertex( const Label v1 );
  const Vertex *findVertex( const Label v1 ) const;
  unsigned long long int countRoutes( Label v1, Label v2, set<Label> visited ) const;
  void findLowestWeightRouteDFS( Label v1,
                                 Label v2,
                                 set<Label> visited,
                                 unsigned long long int &minFound,
                                 unsigned long long int foundSoFar ) const;
  void findLowestWeightRouteAcyclicDFS( Label v1,
                                        Label v2,
                                        unsigned long long int &minFound,
                                        unsigned long long int foundSoFar ) const;
  bool hasEdge( const Vertex *v1, Label v2 ) const;
  void pruneConnectionlessVertices( set<Label> &s1, const set<Label> &s2 ) const;

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
  for ( v_it = other.myVertices.begin(); v_it != other.myVertices.end(); ++v_it )
    {
      addVertex( v_it->first );
    }
  // Once we have the vertices, copy each edge list
  for ( v_it = other.myVertices.begin(); v_it != other.myVertices.end(); ++v_it )
    {
      typename Vertex::const_iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          // Don't use addEdge(). It gets confused.
          addVertexGetPtr( v_it->first )->push_front( Edge( v_it->first, e_it->myV2, e_it->myWeight ) );
        }
    }

  VALIDATE( this );
}

template <typename Label>
Graph<Label>::Graph( unsigned int width,
                     unsigned int height,
                     bool directed,
                     unsigned int weight ) :
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

  for ( h = 0; h < height; h++ )
    {
      for ( w = 0; w < width; w++ )
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
Graph<Label>::Graph( unsigned int adjacencyMatrix[40][40], unsigned int numVertices, const Label &startLabel, bool directed ) :
  myVertices(),
  myIsDirected( directed ),
  myIsSimple( true )
{
  unsigned int row = 0;
  unsigned int col = 0;

  if ( directed )
    {
      for ( row = 0; row < numVertices; ++row )
        {
          for ( col = 0; col < numVertices; ++col )
            {
              if ( adjacencyMatrix[row][col] != 0 )
                {
                  addEdge( startLabel + row, startLabel + col, adjacencyMatrix[row][col] );
                }
            }
        }
    }
  else
    {
      for ( row = 0; row < numVertices; ++row )
        {
          for ( col = row; col < numVertices; ++col )
            {
              if ( adjacencyMatrix[row][col] != 0 )
                {
                  addEdge( startLabel + row, startLabel + col, adjacencyMatrix[row][col] );
                }
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
  const Vertex *v1ptr = findVertex( v1 );
  return hasEdge( v1ptr, v2 );
}

template <typename Label>
bool Graph<Label>::hasEdge( const Vertex *v1, Label v2 ) const
{
  VALIDATE( this );

  if ( v1 == NULL )
    {
      return false;
    }

  typename Vertex::const_iterator it;
  for ( it = v1->begin(); it != v1->end(); ++it )
    {
      if ( it->myV2 == v2 )
        {
          return true;
        }
    }

  return false;
}

template <typename Label>
unsigned int Graph<Label>::getEdgeWeight( Label v1, Label v2 ) const
{
  VALIDATE( this );

  const Vertex *v = findVertex( v1 );

  if ( v == NULL )
    {
      return 0;
    }

  typename Vertex::const_iterator it;
  for ( it = v->begin(); it != v->end(); ++it )
    {
      if ( it->myV2 == v2 )
        {
          return it->myWeight;
        }
    }

  return 0;
}

template <typename Label>
void Graph<Label>::setEdgeWeight( Label v1, Label v2, unsigned int weight )
{
  VALIDATE( this );

  Vertex *v = findVertex( v1 );

  if ( v == NULL || !hasVertex( v2 ) )
    {
      return;
    }

  typename Vertex::iterator it;
  for ( it = v->begin(); it != v->end(); ++it )
    {
      if ( it->myV2 == v2 )
        {
          it->myWeight = weight;
        }
    }
}

template <typename Label>
void Graph<Label>::incEdgeWeight( Label v1, Label v2, unsigned int delta )
{
  VALIDATE( this );

  Vertex *v = findVertex( v1 );

  if ( v == NULL || !hasVertex( v2 ) )
    {
      return;
    }

  typename Vertex::iterator it;
  for ( it = v->begin(); it != v->end(); ++it )
    {
      if ( it->myV2 == v2 )
        {
          it->myWeight += delta;
        }
    }
}

template <typename Label>
void Graph<Label>::decEdgeWeight( Label v1, Label v2, unsigned int delta )
{
  VALIDATE( this );

  Vertex *v = findVertex( v1 );

  if ( v == NULL || !hasVertex( v2 ) )
    {
      return;
    }

  typename Vertex::iterator it;
  for ( it = v->begin(); it != v->end(); ++it )
    {
      if ( it->myV2 == v2 )
        {
          if ( it->myWeight >= delta )
            {
              it->myWeight -= delta;
            }
          else
            {
              it->myWeight = 0;
            }
        }
    }
}

template <typename Label>
void Graph<Label>::addVertex( Label v1 )
{
  VALIDATE( this );

  Vertex *v1ptr = findVertex( v1 );
  if ( v1ptr == NULL )
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

  Vertex *v1ptr = findVertex( v1 );
  if ( v1ptr == NULL )
    {
      Vertex v;
      myVertices.insert( pair<Label, Vertex>( v1, v ) );
      v1ptr = findVertex( v1 );
    }

  VALIDATE( this );

  return v1ptr;
}

template <typename Label>
void Graph<Label>::addEdge( Label v1, Label v2, unsigned int weight )
{
  VALIDATE( this );

  if ( v1 == v2 )
    {
      myIsSimple = false;
    }

  Vertex *v1ptr = addVertexGetPtr( v1 );
  Vertex *v2ptr = addVertexGetPtr( v2 );

  if ( isSimple() )
    {
      // See if there is already an equivalent edge
      if ( hasEdge( v1ptr, v2 ) )
        {
          myIsSimple = false;
        }
    }

  v1ptr->push_front( Edge( v1, v2, weight ) );

  // If this is not directed then edges go both ways, so we
  // also need to add the reverse of this edge.
  if ( !isDirected() )
    {
      v2ptr->push_front( Edge( v2, v1, weight ) );
    }

  VALIDATE( this );
}

template <typename Label>
void Graph<Label>::eraseVertex( Label v1 )
{
  VALIDATE( this );

  bool found = false;

  // Remove all edges that point to this vertex
  typename Vertices::iterator v_it;
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      do
        {
          found = false;
          typename list<Edge>::iterator e_it;
          for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
            {
              if ( e_it->myV2 == v1 )
                {
                  e_it = v_it->second.erase( e_it );
                  // Our iterator is now invalid. Start it over.
                  found = true;
                  break;
                }
            }
        } while ( found );
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
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      typename list<Edge>::iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          if ( isDirected() )
            {
              if ( e_it->myV1 == v1 && e_it->myV2 == v2 )
                {
                  e_it = v_it->second.erase( e_it );
                  // Our iterators are now invalid. Escape!
                  VALIDATE( this );
                  return;
                }
            }
          else
            {
              if ( ( e_it->myV1 == v1 && e_it->myV2 == v2 ) ||
                   ( e_it->myV1 == v2 && e_it->myV2 == v1 ) )
                {
                  e_it = v_it->second.erase( e_it );
                  // Our iterators are now invalid. Escape!
                  VALIDATE( this );
                  return;
                }
            }
        }
    }

  VALIDATE( this );
}

template <typename Label>
void Graph<Label>::eraseAllEdgesOutOf( Label v1 )
{
  VALIDATE( this );

  Vertex *vptr = findVertex( v1 );
  if ( vptr != NULL )
    {
      vptr->clear();
    }

  VALIDATE( this );
}

template <typename Label>
void Graph<Label>::eraseDuplicateEdges( Label v1 )
{
  Vertex *v = findVertex( v1 );

  if ( v == NULL )
    {
      return;
    }

  typename Vertex::iterator it;
  for ( it = v->begin(); it != v->end(); ++it )
    {
      typename Vertex::iterator it2 = it;
      it2++;
      for ( ; it2 != v->end(); ++it2 )
        {
          if ( it->myV2 == it2->myV2 )
            {
              it->myWeight = min( it->myWeight, it2->myWeight );
              v->erase( it2 );
              // We have invalidated our iterators!
              // Start again with new ones then get out.
              eraseDuplicateEdges( v1 );
              return;
            }
        }
    }
}

template <typename Label>
void Graph<Label>::erase( void )
{
  while ( myVertices.size() > 0 )
    {
      eraseVertex( myVertices.begin()->first );
    }
}

template <typename Label>
unsigned long long int Graph<Label>::sumWeights( void ) const
{
  unsigned long long int sum = 0;

  typename Vertices::const_iterator v_it;
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      typename list<Edge>::const_iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          sum += e_it->myWeight;
        }
    }

  return isDirected() ? sum : sum / 2;
}

template <typename Label>
unsigned long long int Graph<Label>::countRoutes( Label v1, Label v2 ) const
{
  VALIDATE( this );
  set<Label> visited;
  return countRoutes( v1, v2, visited );
}

template <typename Label>
unsigned long long int Graph<Label>::countRoutes( Label v1, Label v2, set<Label> visited ) const
{
  if ( v1 == v2 )
    {
      return 1;
    }

  visited.insert( v1 );

  unsigned long long int total = 0;

  const Graph<Label>::Vertex *edgeList = findVertex( v1 );

  typename Vertex::const_iterator edge_it;
  for ( edge_it = edgeList->begin(); edge_it != edgeList->end(); ++edge_it )
    {
      if ( visited.count( edge_it->myV2 ) == 0 )
        {
          total += countRoutes( edge_it->myV2, v2, visited );
        }
    }

  return total;
}

template <typename Label>
unsigned long long int Graph<Label>::findLowestWeightRoute( Label v1, Label v2, bool cyclic, unsigned long long int hint ) const
{
  VALIDATE( this );

  if ( numVertices() <= 1 )
    {
      return 0;
    }

  unsigned long long int minFound = hint;
  if ( cyclic )
    {
      set<Label> visited;
      findLowestWeightRouteDFS( v1, v2, visited, minFound, 0 );
    }
  else
    {
      findLowestWeightRouteAcyclicDFS( v1, v2, minFound, 0 );
    }

  return minFound;
}

template <typename Label>
void Graph<Label>::findLowestWeightRouteDFS( Label v1,
                                             Label v2,
                                             set<Label> visited,
                                             unsigned long long int &minFound,
                                             unsigned long long int foundSoFar ) const
{
  if ( v1 == v2 )
    {
      if ( foundSoFar < minFound )
        {
          minFound = foundSoFar;
        }
      return;
    }

  visited.insert( v1 );

  const Graph<Label>::Vertex *edgeList = findVertex( v1 );

  typename Vertex::const_iterator edge_it;
  for ( edge_it = edgeList->begin(); edge_it != edgeList->end(); ++edge_it )
    {
      if ( visited.count( edge_it->myV2 ) == 0 )
        {
          if ( foundSoFar + edge_it->myWeight < minFound )
            {
              // This edge's weight does not exhaust our budget. Try it.
              findLowestWeightRouteDFS( edge_it->myV2, v2, visited, minFound, foundSoFar + edge_it->myWeight );
            }
        }
    }
}

template <typename Label>
void Graph<Label>::findLowestWeightRouteAcyclicDFS( Label v1,
                                                    Label v2,
                                                    unsigned long long int &minFound,
                                                    unsigned long long int foundSoFar ) const
{
  if ( v1 == v2 )
    {
      if ( foundSoFar < minFound )
        {
          minFound = foundSoFar;
        }
      return;
    }

  const Graph<Label>::Vertex *edgeList = findVertex( v1 );

  typename Vertex::const_iterator edge_it;

  for ( edge_it = edgeList->begin(); edge_it != edgeList->end(); ++edge_it )
    {
      if ( foundSoFar + edge_it->myWeight < minFound )
        {
          // This edge's weight does not exhaust our budget. Try it.
          findLowestWeightRouteAcyclicDFS( edge_it->myV2, v2, minFound, foundSoFar + edge_it->myWeight );
        }
    }
}

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
      for ( it = vptr->begin(); it != vptr->end(); ++it )
        {
          // Add vertex to list to vist
          toVisit.push( it->myV2 );
        }
    }

  return connected;
}

template <typename Label>
set<Label> Graph<Label>::findEdgesInto( Label v1 ) const
{
  set<Label> edges;

  typename Vertices::const_iterator v_it;
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      typename list<Edge>::const_iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          if ( e_it->myV2 == v1 )
            {
              edges.insert( e_it->myV1 );
            }
        }
    }

  return edges;
}

template <typename Label>
set<Label> Graph<Label>::findEdgesOutOf( Label v1 ) const
{
  set<Label> edges;

  const Vertex *v1ptr = findVertex( v1 );
  if ( v1ptr == NULL )
    {
      return edges;
    }

  typename list<Edge>::const_iterator e_it;
  for ( e_it = v1ptr->begin(); e_it != v1ptr->end(); ++e_it )
    {
      edges.insert( e_it->myV2 );
    }

  return edges;
}

template <typename Label>
void Graph<Label>::reduceWeightedDAGToMinimalPath( const Label start, const Label terminus )
{
  if ( !findVertex( start ) || !findVertex( terminus ) )
    {
      cout << "start or terminus did not exist!" << endl;
      return;
    }

  if ( !isConnected( start, terminus ) )
    {
      cout << "start and terminus are not connected!" << endl;
      return;
    }

  while ( !hasEdge( start, terminus ) )
    {
      set<Label> terminusEdges = findEdgesInto( terminus );

      // for each vertex in terminusEdges...
      typename set<Label>::iterator t_it;
      for ( t_it = terminusEdges.begin(); t_it != terminusEdges.end(); ++t_it )
        {
          unsigned int terminusEdgeWeight = getEdgeWeight( *t_it, terminus );
          set<Label> incoming = findEdgesInto( *t_it );
          typename set<Label>::iterator i_it;
          for ( i_it = incoming.begin(); i_it != incoming.end(); ++i_it )
            {
              unsigned int edgeWeight = terminusEdgeWeight  + getEdgeWeight( *i_it, *t_it );
              if ( hasEdge( *i_it, terminus ) )
                {
                  setEdgeWeight( *i_it, terminus, min( edgeWeight, getEdgeWeight( *i_it, terminus ) ) );
                }
              else
                {
                  addEdge( *i_it, terminus, edgeWeight );
                }
            }
          eraseVertex( *t_it );
        }
    }
}

template <typename Label>
void Graph<Label>::removeSymmetricallyConnectedVertices( void )
{
#if 0
  // Remove about 1/4 of the vertices. Find a known pattern
  // (a vertex with exactly 4 neighbors) and remove the
  // central vertex, rewiring the routes accordingly.
  typename Vertices::iterator it;
  for ( it = myVertices.begin(); it != myVertices.end(); ++it )
    {
      if ( it->second.size() == 4 )
        {
          if ( findEdgesInto( it->first ) != findEdgesOutOf( it->first ) )
            {
              continue;
            }
          cout << "Removing a symmetric vertex: " << it->first << endl;
          Label v[4];
          unsigned int i = 0;
          unsigned int j = 0;
          typename list<Edge>::iterator lit = it->second.begin();
          for ( i = 0; i < 4; ++i, ++lit )
            {
              v[i] = lit->myV2;
            }
          for ( i = 0; i < 4; ++i )
            {
              for ( j = i+1; j < 4; ++j )
                {
                  addEdge( v[i], v[j], getEdgeWeight( v[i], it->first ) + getEdgeWeight( it->first, v[j] ) );
                  addEdge( v[j], v[i], getEdgeWeight( v[j], it->first ) + getEdgeWeight( it->first, v[i] ) );
                }
            }
          eraseVertex( it->first );
        }
    }
#else
  // Because we are going to be deleting vertices as we go along
  // we can't use an iterator-over-vertices because it blows up
  // if you delete things from what it is iterating on.
  unsigned int maxEdges = 4;
  do
    {
      set<Label> all;
      typename Vertices::iterator v_it;
      for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
        {
          all.insert( v_it->first );
        }

      typename set<Label>::iterator it;
      for ( it = all.begin(); it != all.end(); ++it )
        {
          Vertex *vptr = findVertex( *it );
          if ( findEdgesInto( *it ) == findEdgesOutOf( *it ) )
            {
              unsigned int outD = vptr->size();
              if ( outD > maxEdges ) { continue; }
              cout << "Removing symmetric vertex (size: " << outD << "/" << maxEdges << ")" << "  [" << numVertices() << "]" << endl;
              vector<Label> v;
              vector<unsigned int> w_i_it;
              vector<unsigned int> w_it_i;
              unsigned int i = 0;
              unsigned int j = 0;
              typename list<Edge>::iterator lit = vptr->begin();
              for ( lit = vptr->begin(); lit != vptr->end(); ++lit )
                {
                  v.push_back( lit->myV2 );
                  w_i_it.push_back( getEdgeWeight( lit->myV2, *it ) );
                  w_it_i.push_back( lit->myWeight );
                  ++i;
                }
              for ( i = 0; i < outD; ++i )
                {
                  for ( j = i+1; j < outD; ++j )
                    {
                      addEdge( v[i], v[j], w_i_it[i] + w_it_i[j] );
                      eraseDuplicateEdges( v[i] );
                      addEdge( v[j], v[i], w_i_it[j] + w_it_i[i] );
                      eraseDuplicateEdges( v[j] );
                    }
                }
              eraseVertex( *it );
            }
        }
      maxEdges += 10;
    } while ( numVertices() > 10 );
#endif
}

template <typename Label>
void Graph<Label>::reduceWeightedDCGToMinimalPath( const Label start, const Label terminus )
{
  if ( !findVertex( start ) || !findVertex( terminus ) )
    {
      cout << "start or terminus did not exist!" << endl;
      return;
    }

  if ( !isConnected( start, terminus ) )
    {
      cout << "start and terminus are not connected!" << endl;
      return;
    }

#if 1
  removeSymmetricallyConnectedVertices();
#endif

  while ( numVertices() > 2 )
    {
      typename set<Label>::iterator t_it;
      set<Label> edges;

      edges = findEdgesInto( terminus );
      edges.erase( start );   // Don't nuke the start vertex!
      // for each vertex T in terminus' edges...
      for ( t_it = edges.begin(); t_it != edges.end(); ++t_it )
        {
          cout << "Removing T vertex: " << *t_it << "  [" << numVertices() << "]" << endl;
          // ...find the vertices I that lead into it...
          set<Label> incoming = findEdgesInto( *t_it );
          typename set<Label>::iterator i_it;
          // ...then remove T and connect I directly to terminus...
          unsigned int tweight = getEdgeWeight( *t_it, terminus );
          for ( i_it = incoming.begin(); i_it != incoming.end(); ++i_it )
            {
              unsigned long long int hint = getEdgeWeight( *i_it, *t_it ) + tweight;
              unsigned int edgeWeight = findLowestWeightRoute( *i_it, terminus, true, hint );
              eraseAllEdgesOutOf( *i_it );
              addEdge( *i_it, terminus, edgeWeight );
            }
          eraseVertex( *t_it );
        }

      edges.clear();

      edges = findEdgesOutOf( start );
      edges.erase( terminus );   // Don't nuke the terminus vertex!
      // for each vertex S in start's edges...
      for ( t_it = edges.begin(); t_it != edges.end(); ++t_it )
        {
          cout << "Removing S vertex: " << *t_it << "  [" << numVertices() << "]" << endl;
          // ...find the vertices I that T leads out to...
          set<Label> outgoing = findEdgesOutOf( *t_it );
          typename set<Label>::iterator i_it;
          // ...then remove S and connect start directly to I...
          unsigned int t_weight = getEdgeWeight( start, *t_it );
          for ( i_it = outgoing.begin(); i_it != outgoing.end(); ++i_it )
            {
              unsigned long long int hint = t_weight + getEdgeWeight( *t_it, *i_it );
              unsigned int edgeWeight = findLowestWeightRoute( start, *i_it, true, hint );
              addEdge( start, *i_it, edgeWeight );
            }
          eraseVertex( *t_it );
        }

      typename Vertices::iterator v_it;
      for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
        {
          eraseDuplicateEdges( v_it->first );
        }
    }
}

template <typename Label>
void Graph<Label>::eraseAllEdgesInto( const Label v1 )
{
  set<Label> edges = findEdgesInto( v1 );
  typename set<Label>::iterator i_it;
  for ( i_it = edges.begin(); i_it != edges.end(); ++i_it )
    {
      eraseEdge( *i_it, v1 );
    }
}

template <typename Label>
bool Graph<Label>::isConnected( const Label *startV ) const
{
  if ( numVertices() <= 1 )
    {
      return true;
    }

  set<Label> all;
  set<Label> connected;

  // Load all of the vertices into a set
  typename Vertices::const_iterator it;
  for ( it = myVertices.begin(); it != myVertices.end(); ++it )
    {
      all.insert( it->first );
    }

  // Grab an arbitrary vertex and find all
  // that are connected to it. If the caller
  // gave a suggested start, use that.
  Label v;
  if ( startV != NULL )
    {
      v = *startV;
    }
  else
    {
      v = myVertices.begin()->first;
    }
  connected = findConnectedVertices( v  );

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
  for ( v_it = vptr->begin(); v_it != vptr->end(); ++v_it )
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
  for ( v_it = vptr->begin(); v_it != vptr->end(); ++v_it )
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
bool Graph<Label>::findLightestEdge( const set<Label> &s1, const set<Label> &s2, Label &v1, Label &v2, unsigned int &minWeight ) const
{
  typename set<Label>::iterator s1_it;
  bool found = false;

  minWeight = 0xFFFFFFFF;

  for ( s1_it = s1.begin(); s1_it != s1.end(); ++s1_it )
    {
      const Vertex *v1ptr = findVertex( *s1_it );
      typename Vertex::const_iterator edge_it;
      for ( edge_it = v1ptr->begin(); edge_it != v1ptr->end(); ++edge_it )
        {
          if ( s2.count( edge_it->myV2 ) != 0 )
            {
              if ( edge_it->myWeight < minWeight )
                {
                  minWeight = edge_it->myWeight;
                  v1        = edge_it->myV1;
                  v2        = edge_it->myV2;
                  found     = true;
                }
            }
        }
    }

  return found;
}

template <typename Label>
void Graph<Label>::pruneConnectionlessVertices( set<Label> &s1, const set<Label> &s2 ) const
{
  set<Label> s1Copy = s1;
  typename set<Label>::iterator s1_it;

  for ( s1_it = s1Copy.begin(); s1_it != s1Copy.end(); ++s1_it )
    {
      bool hasV2edge = false;
      const Vertex *v1ptr = findVertex( *s1_it );
      typename Vertex::const_iterator edge_it;
      for ( edge_it = v1ptr->begin(); edge_it != v1ptr->end(); ++edge_it )
        {
          if ( s2.count( edge_it->myV2 ) != 0 )
            {
              hasV2edge = true;
              break;
            }
        }
      if ( !hasV2edge )
        {
          s1.erase( *s1_it );
        }
    }
}

template <typename Label>
void Graph<Label>::reduceToMST( Graph<Label> &mst, const Label *startV ) const
{
  // Delete everything from mst. We will build it from scratch.
  mst.erase();
  // We need to tell mst whether to be directed or not.
  mst.myIsDirected = this->myIsDirected;
  // It will decide whether to stay simple based on the edges added
  // (in this case, a minimal spanning tree is always simple).
  mst.myIsSimple = true;

  if ( numVertices() == 0 )
    {
      return;
    }

  // WARNING: Watch out for non-connected graphs!
  if ( !isConnected( startV ) )
    {
      return;
    }

  set<Label> s1;    // The vertices in mst
  set<Label> s2;    // The vertices not yet in mst

  typename Vertices::const_iterator it;
  for ( it = myVertices.begin(); it != myVertices.end(); ++it )
    {
      s2.insert( it->first );
    }

  //
  // Start: Pick a starting vertex from 'this'.
  //        Copy it to 'mst'. If the caller gave
  //        us a suggested start, use that.
  Label v;
  if ( startV != NULL )
    {
      v = *startV;
    }
  else
    {
      v = myVertices.begin()->first;
    }
  mst.addVertex( v );
  s1.insert( v );
  s2.erase( v );
  // Loop:  Find the least-cost route from the vertices in 'mst'
  //        to the vertices not in 'mst'.
  //        Add that edge to mst.
  // End:   All vertices have been copied to 'mst'.
  //
  while ( !s2.empty() )
    {
      Label v1;
      Label v2;
      unsigned int weight = 0;
      findLightestEdge( s1, s2, v1, v2, weight );
      mst.addEdge( v1, v2, weight );      // Adding the edge also adds the missing vertex
      s1.insert( v1 );
      s1.insert( v2 );
      s2.erase( v1 );
      s2.erase( v2 );
      // After a while, many of the vertices in s1 will no longer have
      // any edges into s2. When that happens we no longer need to
      // consider those vertices and they can be pruned. Pruning is
      // time consuming. Only prune s1 has grown by a sizable amount.
      if ( s2.size() % 100 == 0 )
        {
          pruneConnectionlessVertices( s1, s2 );
        }
    }
}

template <typename Label>
void Graph<Label>::print( bool verbose ) const
{
  unsigned long long int weight = 0;

  VALIDATE( this );

  cout << "Vertices    : " << numVertices() << endl;
  cout << "Edges       : " << numEdges() << endl;
  cout << "isDirected? : " << isDirected() << endl;
  cout << "isSimple?   : " << isSimple() << endl;

  typename Vertices::const_iterator v_it;
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      if ( verbose )
        {
          cout << "Vertex " << v_it->first << "(" << v_it->second.size() << ") :";
        }
      typename list<Edge>::const_iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          if ( verbose )
            {
              cout << " --" << e_it->myWeight << "--> " << e_it->myV2;
            }
          weight += e_it->myWeight;
        }
      if ( verbose )
        {
          cout << endl;
        }
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
  for ( v_it = myVertices.begin(); v_it != myVertices.end(); ++v_it )
    {
      typename Vertex::const_iterator e_it;
      for ( e_it = v_it->second.begin(); e_it != v_it->second.end(); ++e_it )
        {
          if ( v_it->first != e_it->myV1 )
            {
              cout << file << ":" << line << ": error: A vertex ( "
                   << v_it->first
                   << " ) has an edge that does not belong to it ( "
                   << e_it->myV1 << ", " << e_it->myV2 << ", "
                   << e_it->myWeight << " )" << endl;
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
