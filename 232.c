#include <iostream>
#include <cmath>
#include <cstdlib>

using namespace std;

//
// true  == HEADS
// false == TAILS
//
bool flip( void )
{
  return rand() % 2 == 1;
}

unsigned int mirror( void )
{
  return flip() ? 1 : 0;
}

unsigned int T( unsigned int t )
{
  unsigned int i = 0;

  for ( i=1; i<=t; i++ )
    {
      if ( !flip() )
	{
	  return 0;
	  break;
	}
    }

  return pow( 2, t - 1 );
}

unsigned int greedy( unsigned int to_win )
{
  unsigned int t = 0;

  if ( to_win <= 1 )
    {
      t = 1;
    }
  else if ( to_win <= 2 )
    {
      t = 2;
    }
  else if ( to_win <= 4 )
    {
      t = 3;
    }
  else if ( to_win <= 8 )
    {
      t = 4;
    }
  else if ( to_win <= 16 )
    {
      t = 5;
    }
  else if ( to_win <= 32 )
    {
      t = 6;
    }
  else if ( to_win <= 64 )
    {
      t = 7;
    }
  else
    {
      t = 8;
    }

  return T( t );
}

//
// true  == P1 win
// false == P2 win
//
bool playGame( void )
{
  unsigned int p1_score = 0;
  unsigned int p2_score = 0;

  do
    {
      // Player 1 turn
      p1_score += flip() ? 1 : 0;
      if ( p1_score >= 100 ) { break; }

      // Player 2 turn

// Player 2 plays just like player 1
//      p2_score += mirror();                          // 49.0%

// Player 2 goes for the win each turn
//      p2_score += greedy( 100 - p2_score );          // 54.1%

// Player 2 tries to fill half the gap each turn
//      p2_score += greedy( ( 100 - p2_score ) / 2 );  // 60.1%

// Player 2 tries to aggressively catch up if she is behind
//      p2_score += greedy( p1_score > p2_score ? p1_score - p2_score : 1 );    // 79.3%
//      p2_score += greedy( p1_score > p2_score ? p1_score - p2_score + 1 : 1 );    // 82.3%
//      p2_score += greedy( p1_score > p2_score ? p1_score - p2_score + 2 : 1 );    // 82.4%
//      p2_score += greedy( p1_score > p2_score ? p1_score - p2_score + 3 : 1 );    // 82.2%

// Player 2 tries to maintain an X-point lead
//      p2_score += greedy( (p1_score+1) > p2_score ? (p1_score+1) - p2_score + 1 : 1 );    // 82.3%
//      p2_score += greedy( (p1_score+2) > p2_score ? (p1_score+2) - p2_score + 1 : 1 );    // 82.3%
//      p2_score += greedy( (p1_score+4) > p2_score ? (p1_score+4) - p2_score + 1 : 1 );    // 81.5%
//      p2_score += greedy( (p1_score+6) > p2_score ? (p1_score+6) - p2_score + 1 : 1 );    // 80.3%
//      p2_score += greedy( (p1_score+8) > p2_score ? (p1_score+8) - p2_score + 1 : 1 );    // 78.7%      
//      p2_score += greedy( (p1_score+10) > p2_score ? (p1_score+10) - p2_score + 1 : 1 );    // 77.7%

// Player 2 tries a constant-T strategy
//      p2_score += T( 1 );                       // 49.1%
//      p2_score += T( 2 );                       // 50.4%
//      p2_score += T( 3 );                       // 51.6%
//      p2_score += T( 4 );                       // 47.0%
//      p2_score += T( 5 );                       // 42.8%
//      p2_score += T( 6 );                       // 37.6%
//      p2_score += T( 7 );                       // 46.4%
//      p2_score += T( 8 );                       // 54.2%
    } while ( p2_score < 100 );

  return p1_score > p2_score;
}

int main( int argc, char *argv[] )
{
  srand( time(NULL) );

  unsigned int i = 0;
  unsigned int j = 0;
  unsigned int p1_wins = 0;
  unsigned int p2_wins = 0;

  for ( i=1; i<=10000; i++ )
    {
      for ( j=1; j<=10; j++ )
	{
	  if ( playGame() )
	    {
	      p1_wins++;
	    }
	  else
	    {
	      p2_wins++;
	    }
	}
    }

  cout << "P1 wins: " << p1_wins << endl;
  cout << "P2 wins: " << p2_wins << endl;
}
