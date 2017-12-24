"""
day 23 Part 2

For future me: Why is this the solution?
At first, I re-wrote that sample into a Python program and run that. It run for a little while,
but spat out 909. I tried and it did work. Than I thought, that it has to be faster than that,
so I kept on checking what it was actually doing, now that I re-wrote it in Python.
After I saw the code better, I saw that it actually is counting non prime numbers.
And thus, this is the correct solution.
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
