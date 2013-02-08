#! /usr/bin/ruby -W1

require 'pp.rb'

primes = [ 2, 3, 5 ]
candidate = 7

while candidate <= 2000000 do
  max = Math.sqrt( candidate ).to_i + 1
  i = 0
  prime = true
  while primes[i] <= max do
    if candidate % primes[i] == 0 then
      prime = false
      break
    end
    i += 1
  end
  primes << candidate if prime
  candidate += 2
end

total = primes.inject { |p,sum| p + sum }
puts total
