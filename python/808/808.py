#
# Both 169 and 961 are the square of a prime. 169 is the reverse of 961.
# We call a number a reversible prime square if:
#   1. It is not a palindrome (single digits are palindromes), and
#   2. It is the square of a prime, and
#   3. Its reverse is also the square of a prime.
# 169 and 961 are not palindromes, so both are reversible prime squares.
#
# Find the sum of the first 50 reversible prime squares.

import math


# Is n prime?
def is_prime(n):
    if (n & 0x01) == 0:
        # Even numbers are not prime
        return False
    root = math.sqrt(n)
    if root == int(root):
        # This is a perfect square, so it isn't prime
        return False
    min_factor = 3
    max_factor = int(root)
    for factor in range(min_factor, max_factor + 1, 2):
        if n % factor == 0:
            return False
    return True


# Is n a perfect square? If so, what is its root?
def is_square(n):
    root = math.sqrt(n)
    return int(root), root == int(root)


# Is n a Reverse Square Prime?
def is_rsp(n):
    n_str = str(n)
    n_str_rev = n_str[::-1]

    if n_str == n_str_rev:
        # This is a palindrome
        return False

    root_n, square = is_square(n)
    if not square:
        # n is not a square
        return False

    n_rev = int(n_str_rev)
    root_n_rev, square = is_square(n_rev)
    if not square:
        # n reversed is not a square
        return False

    return is_prime(root_n) and is_prime(root_n_rev)


def main():
    print("Welcome to 808\n")

    n = 1
    count = 0
    total = 0

    while count < 50:
        square = n * n
        if is_rsp(square):
            print(count, "-", square)
            count += 1
            total += square
        n += 2  # Skip even numbers, as they are not prime

    print("Sum =", total)


main()
