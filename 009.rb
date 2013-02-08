#! /usr/bin/ruby -W1

def find_sum( sum, values )
  i = 0
  while i < values.length do
    a = values[i]
    b = sum - a
    if values.index b then
      x = Math.sqrt( a ).to_i
      y = Math.sqrt( b ).to_i
      z = Math.sqrt( sum ).to_i
      if x + y + z == 1000 then
        puts "  #{x} + #{y} + #{z} = #{x + y + z}"
      end
#      return a, b
    end
    i += 1
  end
end

squares = []
(1..1000).each { |i| squares << i**2 }
squares.reverse!

while squares.length >= 3 do
  c = squares.shift
  a, b = find_sum c, squares
  if a && b then
    a = Math.sqrt( a ).to_i
    b = Math.sqrt( b ).to_i
    c = Math.sqrt( c ).to_i
    if a + b + c == 1000 then
      puts "  #{a} + #{b} + #{c} = #{a + b + c}"
    end
  end  
end
