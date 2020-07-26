#! /usr/bin/ruby -W1

def precompute
  a = [ [] ]
  for n in 1..100 do
    a << collatz( n )
  end
  a
end

def collatz( n, precompute=nil )
  a = []

  until n == 1 do
    if precompute && n < 100 then
      return a + precompute[n]
    end

    a << n
    if n == (n >> 1) << 1 then
      # Even
      n = n >> 1
    else
      # Odd
      n = 1 + n + n + n
    end
  end

  a << 1
end

pre = precompute

max_i   = 0
max_len = 0

for i in 1...1000000 do
  a = collatz i, pre
  if a.length > max_len then
    max_len = a.length
    max_i = i
  end
end

puts "Max chain: i = #{max_i} len = #{max_len}"
