#! /usr/bin/ruby -W1

require 'pp.rb'

primes = [ 2 ]
candidate = 3

while primes.length < 10001 do
  prime = true
  primes.each do |p|
    if candidate % p == 0 then
      prime = false
      break
    end
  end
  primes << candidate if prime
  candidate += 2
end

puts "The 10,001th prime is #{primes[10000]}"
