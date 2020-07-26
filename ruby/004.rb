#! /usr/bin/ruby -W1

def find_two_3digit_factors( value )
  value = value.to_i
  i = 100
  while i <= 999 do
    if value % i == 0 then
      factor = value / i
      if factor.to_s.length == 3 then
        puts "#{value} has two 3-digit factors: #{i} and #{factor}"
exit
      end
    end
    i += 1
  end
end

def palindromify_6( value )
  value.to_s + value.to_s.reverse
end

def palindromify_5( value )
  s = value.to_s
  t = value.to_s.reverse
  t[0] = ""
  s + t
end

def is_palindrome?( value )
  value.to_s == value.to_s.reverse
end

#
# 6-digit palindromes
#
i = 999
while i >= 100 do
  p = palindromify_6 i
  find_two_3digit_factors p
  i -= 1
end

#
# 5-digit palindromes
#
i = 999
while i >= 100 do
  p = palindromify_5 i
  find_two_3digit_factors p
  i -= 1
end

