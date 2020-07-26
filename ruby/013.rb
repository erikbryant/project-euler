#! /usr/bin/ruby -W1

def sum( fh )
  a = 0
  fh.map do |line|
    a += line.to_i
  end
  a
end

a = sum ARGV.first ? File.open(ARGV.shift) : $stdin

puts a
