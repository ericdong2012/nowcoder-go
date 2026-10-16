"""
质数(素数)相关。
质数又称素数。一个大于 1 的自然数，如果除了 1 和它自身外，不能被其他自然数整除的数。
"""
import math


def is_prime(n):
    if n <= 1:
        return False
    i = 2
    while i <= int(math.sqrt(n)):
        if n % i == 0:
            return False
        i += 1
    return True


def find_primes(n):
    primes = []
    for i in range(2, n + 1):
        if is_prime(i):
            primes.append(i)
    return primes


if __name__ == "__main__":
    n = 100
    primes = find_primes(n)
    print(f"The primes less than or equal to {n} are: {primes}")
