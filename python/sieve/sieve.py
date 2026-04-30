import math

class Sieve:

    def __init__(self) -> None:
        pass

    def nth_prime(self, n: int) -> int:
        if n < 0:
            raise ValueError("n must be non-negative")

        # Estimate upper bound for nth prime using the prime number theorem
        if n < 6:
            upper_bound = 15
        else:
            # TODO: apparently this is a loose upper bound, other tighter bounds exist
            upper_bound = int(n * (math.log(n) + math.log(math.log(n))))

        # Sieve of Eratosthenes
        # Create a list of index 0 - n of True's
        sieve = [True] * (upper_bound + 1)
        # 0 and 1 are not prime
        sieve[0] = sieve[1] = False
        
        # Start at 2, go to sqrt(n)
        for num in range(2, int(math.sqrt(upper_bound))+1):
            if sieve[num] == True:
                # Mark all multiples False, indexing by num
                for composite_index in range(num * num, upper_bound, num):
                    sieve[composite_index] = False

        # Get list of primes
        prime_numbers = []
        for prime_index in range(0, len(sieve)):
            if sieve[prime_index] == True:
                prime_numbers.append(prime_index)
            
        # Return nth prime
        return prime_numbers[n]



if __name__ == "__main__":
    sieve = Sieve()
    sieve.nth_prime(10)