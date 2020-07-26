#! /usr/bin/ruby -W1

def sum_of_squares( value )
  sum = 0
  while value >= 1 do
    sum += value**2
    value -= 1
  end

  sum
end

def square_of_sums( value )
  sum = value * ( 1 + value ) / 2
  sum**2
end

a = sum_of_squares 100
b = square_of_sums 100
c = b - a
puts "Sum of squares = #{a}"
puts "Square of sums = #{b}"
puts "Difference     = #{c}"

