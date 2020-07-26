#! /usr/bin/ruby -W1

def triangle( n )
  ( (1 + n) * n ) >> 1
end

def divisors( n )
  a = [ 1 ]
  b = []
  for i in 2..Math.sqrt(n) do
    a << i if n % i == 0
  end
  for d in a do
    b << d
    b << n/d
  end
  b
end

i = 1
t = triangle i-1

for n in i..(i+10000) do
  t += n
  d = divisors t
  puts "#{n}: #{t} #{d.length}" if d.length >= 500
end
