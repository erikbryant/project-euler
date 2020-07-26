#! /usr/bin/ruby -W1

target=600851475143

def is_prime?( value )
  f = find_factors value
  f.length == 0
end

def find_factors( value )
  a = []

  factor = 2

  while factor <= Math.sqrt( value ) do
    a << factor if value % factor == 0
    factor += 1
  end

  a
end

factors = find_factors target

puts "The factors of #{target} are #{factors}"

factors.each { |f| puts "Prime: #{f}" if is_prime? f }
