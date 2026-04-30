1. What is the sieve of Eratosthenes/Eratosthenes' algorithm?
    - Basically goes from 2 to some number and removes the multiples of 2, goes to next non removed number so 3 and removes it multiples, then 5 because 4 was removed, then 7 because 6 was removed, then 11 because 8, 9, 10 were removed, etc. Keep count of the number of prime numebers found until you reach N prime numbers.
    - Sieve of Eratosthenes tends to be good up to 10^9
        - Implentation: (1) define an upper bound using prime number theorem variation from Rosser and Schoenfeld 1962, which is n*(ln(n) + ln(ln(n))) for 6 <= n and is a corollary to the lower bound theorem. then (2) go from 2 to sqrt(upper_bound) since composites will all be caught less than sqrt of upper_bound, and (3) mark all multiples as composites and keep interating upwards.
2. Design an API that can be called to retrieve Nth prime number
3. Implement
4. Test using test suite

Passing Conditions
1. Existing Test Suite
    - python -m unittest test_sieve.py
2. Add some of my own tests