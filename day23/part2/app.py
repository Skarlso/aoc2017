"""
day 23 Part 2
"""
from math import sqrt, floor


def isPrime(n):
    for j in range(2, floor(sqrt(n))):
        if (n % j) == 0:
            return False
    return True


# pylint: disable-msg=C0103
h = 0
b = 81 * 100 + 100000
c = b + 17000
for i in range(b, c + 1, 17):
    if not isPrime(i):
        h += 1

print("Number of primes: %d" % h)
