#
# Project Euler problem #684
#

fCache = {
    0: 0,
    1: 1
}

Smax = 0
Ssum = 0

# Define s(n) to be the smallest number that has a
# digit sum of n. For example, s(10) = 19.
#
# To make the math more efficient, this function
# actually returns a number 1 higher than the
# answer. The caller has the responsibility of
# subtracting 1 from this answer.
def s(n):
    nines = n // 9
    remainder = n % 9
    return remainder * pow(10, nines)

# Let S(k) = the summation from n=1 to k of s(n).
# You are given S(20) = 1074.
def S(k):
    # We assume this is called in ascending order so that
    # we can continue summing where we left off.
    global Smax, Ssum

    count = 0
    for n in range(Smax+1, k+1):
        Ssum += s(n)
        count += 1

    # s(n) returns a number that is 1 higher than
    # the actual answer. Subtract all those extras.
    Ssum -= count
    Smax = k

    return Ssum

# Let f(i) be the Fibonacci sequence defined by
# f(0) = 0
# f(1) = 1
# f(i) = f(i-2) + f(i-1)
def f(i):
    cached = fCache.get(i, -1)
    if cached != -1:
        return cached

    fCache[i] = f(i-2) + f(i-1)

    return fCache[i]

# Find the summation from i=2 to 90 of S(f(i)).
# Give the answer modulo 1,000,000,007.
def main():
    sum = 0

    end = 90
    for i in range(2, end+1):
        print(i)
        sum += S(f(i))
        sum = sum % 1000000007

    print("answer = ", sum)

    print(max9)

main()
