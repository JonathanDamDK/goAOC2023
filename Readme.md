# Advent of code
This is my attempt at solving Advent Of code 2023 and the go programming language.If you want to attempt AOC yourself go to the website [Advent of Code](https://adventofcode.com) for a new challenge every day.
# Exercise 1
during a rewrite, i somehow manage to botch day one pt 2 and it is currently not giving correct answers


# Code structure
Each day is its own struct that all implemenent the **Solve()** function. This is to make main as compressed as possible. All input files are gitignored in inputs/. All go files are in the main directory.

# running it 
if for any reason you want to run this the only dependency is having go installed. Please note three flags are available when using the compiler
**-today -includeFrom=[number] -includeTo[number]** these specify which days to run, includeFrom=5 will run all days after the 5th of december, and includeTo=5 will run all days up to the specified date. -today Only runs the most recently added day.
