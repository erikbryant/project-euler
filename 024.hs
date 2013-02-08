module Euler where

import Data.List

perm [] = [ [] ]
perm xs = [ x:ps | x <- xs, ps <- perm (delete x xs) ]
