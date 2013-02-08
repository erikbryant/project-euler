#! /usr/bin/ruby -W1

i = 1
f_2 = 0
f_1 = 1
f_0 = 1
sum = 0

puts "#{i} : #{f_0}"

while f_0 < 4000000 do
  i += 1
  f_0 = f_1 + f_2
  f_2, f_1 = f_1, f_0
  puts "#{i} : #{f_0}"

  i += 1
  f_0 = f_1 + f_2
  f_2, f_1 = f_1, f_0
  puts "#{i} : #{f_0}"

  if f_0 <= 4000000
    puts "Summing where i = #{i}, f_0 = #{f_0}"
    sum += f_0
  end

  i += 1
  f_0 = f_1 + f_2
  f_2, f_1 = f_1, f_0
  puts "#{i} : #{f_0}"
end

puts "Sum : #{sum}"
