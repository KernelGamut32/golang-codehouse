# Exercise 1: Functions

- write a function called `circleinfo` which accepts a float64 radius and returns two values, the area of the circle, and the circumference
- area = Pi * radius * radius
- circumference = Pi * radius * 2 
- Pi is defined in the math package

# Exercise 2: Pointers and Functions

Write a function `doubler` which takes a string argument and doubles it, i.e., "Golang" would become "GolangGolang"

- note that '+' is the way to concatenate two strings
- your function should not return anything, it should modify its argument directly
- write a main program to test your function
- once you've done that, make it so that `doubler` calls a second function, `doublerHelper` which actually does the doubling (to understand how to pass pointers around)

# Exercise 3: Command Line Args

- write a program which takes two command line args representing integers and prints out their sum
- test it on your machine since it won't run in the playground
- hint: remember Atoi in the strconv package (https://golang.org/pkg/strconv) and ignore errors for now

# Exercise 4: Fizzbuzz

Write a function which accepts an integer and returns

- "fizz" if the number is divisible by 3
- "buzz" if the number is divisible by 5
- "fizzbuzz" if the number is divisible by BOTH 3 and 5
- otherwise it just returns the string version of the number, e.g., "4"

Test your function with these inputs: 3, 5, 15, 4

# Exercise 5: for Loops

Implement a square root function using Newton's method:

- starting with some guess for the square root of x, we can adjust it based on how close guess² is to x, producing a better guess:
- guess = guess − (guess² − x) / (2 * guess)
- repeating the above makes the guess better and better
- use a starting guess of 1.0, regardless of the input (it works quite well)
- repeat the calculation 10 times and print each guess along the way

OR: implement the factorial function, n! = n * (n - 1) * ... * 1

# Exercise 6: for Loop as while Loop

Write a function which implements the Collatz Conjecture:

- it should accept an integer >= 1 (return false if < 1)
- if it's even, divide it by 2
- if it's odd, multiply it by 3 and add 1
- stop when the the result is 1
- return true for success

What happens if you use a REALLY, REALLY, REALLY, REALLY, REALLY big number? For example, 12341234123412341234123412341234123412341234123412341234123412341234123412342345234

# Exercise 7: Slices

1. Write a function minmaxavg which accepts a float64 slice and returns four values: the minimum value in the slice, the maximum value in the slice, the average of all the values, and an indication of success (as we did previously)

2. Write a function fib that returns a slice containing the first n Fibonacci numbers

- the Fibonacci sequence is 1, 1, 2, 3, 5, 8...
- each number is the sum of the previous Fibonacci numbers
- so fib(6) should return an int slice of size 6 containing 1, 1, 2, 3, 5, 8
- in the event of an error, fib should return the nil slice

# Exercise 8: Maps #1

- use a map to translate Roman numerals into their Arabic equivalents
- load the map with Roman numerals M (1000), D (500), C (100), L (50), X (10), V (5), I (1)
- read in a Roman numeral
- print Arabic equivalent
- try it with MCLX = 1000 + 100 + 50 + 10 = 1160
- complain if invalid Roman digits (e.g., 'Z') are found

**EXTRA**: Deal with the case where a smaller number precedes a larger number, e.g., XC = 100 - 10 = 90, or MCM = 1000 + (1000-100) = 1900

# Exercise 9: Maps #2

- create four calculator functions: add, sub, mul, div
- each function should take 2 integers and return an integer result
- create a map which maps the strings representing the operators to the functions
- that is, "+" would be mapped to add(), "-" would be mapped to sub()...
- have your program read in lines like "2 + 4" and have it determine the result by parsing the line and then using the operator ("+" in this case) to find the appropriate function to invoke

# Exercise 10: Reading from Files

- write a Go program to read a file and count the number of occurrences of each word in the file
- use a map indexed by word, to count the occurrences
- treat The and the as the same word when counting
- NOTE: the strings package has several functions that will be helpful
- NOTE 2: ideally we would sort by counts, from most common to least common, but this is much more difficult to do in Go than it is in Python
- EXTRA: remove punctuation, so Hamlet, == Hamlet

Test on Shakespeare's Hamlet (https://gist.githubusercontent.com/davewadestein/270167a513c42fc9798b6759911ee716/raw/ef89b09a0ee28e6c3754f95e6f57f1cd9bc0dad3/hamlet.txt)

# Exercise 11: HTTP API

# Exercise 12: Vue.js Front-End
