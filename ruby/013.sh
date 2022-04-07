#! /bin/zsh

echo "10 digits of input"
cut -c1-10 013.data | ./013.rb | cut -c1-10

echo "11 digits of input"
cut -c1-11 013.data | ./013.rb | cut -c1-10

echo "12 digits of input"
cut -c1-12 013.data | ./013.rb | cut -c1-10
