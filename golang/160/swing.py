import bisect
import math


# Sieve of Eratosthenes
# Code by David Eppstein, UC Irvine, 28 Feb 2002
# http://code.activestate.com/recipes/117119/

def gen_primes():
    """ Generate an infinite sequence of prime numbers.
    """
    # Maps composites to primes witnessing their compositeness.
    # This is memory efficient, as the sieve is not "run forward"
    # indefinitely, but only as long as required by the current
    # number being tested.
    #
    D = {}

    # The running integer that's checked for primeness
    q = 2

    while True:
        if q not in D:
            # q is a new prime.
            # Yield it and mark its first multiple that isn't
            # already marked in previous iterations
            #
            yield q
            D[q * q] = [q]
        else:
            # q is composite. D[q] is the list of primes that
            # divide it. Since we've reached q, we no longer
            # need it in the map, but we'll mark the next
            # multiples of its witnesses to prepare for larger
            # numbers
            #
            for p in D[q]:
                D.setdefault(p + q, []).append(p)
            del D[q]

        q += 1

def factorial(n):
    def prime_range(start, stop):
        """Return a list of all primes between start and stop - 1, inclusive."""
        if stop < start:
            print("start must be greater than stop", start, stop)
            return []
        primes = []
        for p in gen_primes():
            if p > stop-1:
                break
            if p >= start:
                primes.append(p)
                continue
        return primes

    def isqrt(a):
        """Returns the greatest integer <= to math.sqrt(n)."""
        return int(math.sqrt(a))

    def product(s, n, m):
        return math.prod(s[n:m+1])

    def swing(m, primes):
        if m < 4:
            return [1,1,1,3][m]

        s = bisect.bisect_left(primes, 1 + isqrt(m))
        d = bisect.bisect_left(primes, 1 + m // 3)
        e = bisect.bisect_left(primes, 1 + m // 2)
        g = bisect.bisect_left(primes, 1 + m)
        print(m, s, d, e, g)

        factors = primes[e:g]
        factors += filter(lambda x: (m//x)&1 == 1, primes[s:d])
        for prime in primes[1:s]:
            p, q = 1, m
            while True:
                q //= prime
                if q == 0:
                    break
                if q & 1 == 1:
                    p *= prime
            if p > 1:
                factors.append(p)

        return product(factors, 0, len(factors) - 1)

    def odd_factorial(n, primes):
        if n < 2: return 1
        return (odd_factorial(n//2,primes)**2)*swing(n,primes)

    def compute(n):
        # if n < 10:
        #     return product(range(2, n + 1), 0, n-2)
        bits = n - n.bit_count()
        primes = prime_range(2, n + 1)
        return odd_factorial(n, primes) * 2**bits

    return compute(n)


n = 11
f = factorial(n)
print("%12d! = %12d" % (n, f))

# for k in range(1, 8):
#     n = 10**k
#     f = factorial(n)
#     while f%10 == 0:
#         f = f // 10
#     f = f%10000000
#     print("%12d! = %12d" % (n, f))

