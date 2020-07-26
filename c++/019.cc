//
// Copyright Erik Bryant (erikbryantology@gmail.com)
//

#include <iostream>

using std::cout;
using std::endl;

bool isLeapYear( unsigned int year )
{
  bool isLeap = false;

  if ( year % 100 == 0 )
    {
      isLeap = ( year % 400 == 0 );
    }
  else
    {
      isLeap = ( year % 4 == 0 );
    }

  return isLeap;
}

unsigned int monthLen( unsigned int year, unsigned int month )
{
  // Months are 1-based
  const unsigned int monthLen[] = { 0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 };

  if ( isLeapYear( year ) && month == 2 )
    {
      return monthLen[month] + 1;
    }

  return monthLen[month];
}

unsigned int daysSinceEpoch( unsigned int year, unsigned int month, unsigned int day )
{
  unsigned int y = 0;
  unsigned int m = 0;
  unsigned int days = 0;

  if ( year < 1900 || month <= 0 || month > 12 || day > 31 )
    {
      return 0;
    }

  if ( day > monthLen( year, month ) )
    {
      return 0;
    }

  // Count forward until we get to the year specified
  for ( y = 1900; y < year; ++y )
    {
      for ( m = 1; m <= 12; ++m )
        {
          days += monthLen( y, m );
        }
    }

  // Count forward until we get to the month specified
  for ( m = 1; m < month; ++m )
    {
      days += monthLen( year, m );
    }

  // Add the days
  days += day - 1;

  return days;
}

int main( int argc, char *argv[] )
{
  unsigned int year = 0;
  unsigned int month = 0;
  unsigned int day = 0;
  unsigned int sundays = 0;

  day = 1;
  for ( year = 1901; year <= 2000; ++year )
    {
      for ( month = 1; month <= 12; ++month )
        {
          unsigned int daysSince = daysSinceEpoch( year, month, day );
          unsigned int dayOfWeek = ( daysSince + 8 ) % 7;

          if ( dayOfWeek == 0 )
            {
              cout << year << "/" << month << "/" << day << " is a Sunday" << endl;
              ++sundays;
            }
        }
    }

  cout << "The number of months between 1901/1/1 and 2000/12/31 that started on a Sunday is " << sundays << endl;
}
