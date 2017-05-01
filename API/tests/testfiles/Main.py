from functools import reduce
S = input()

print(reduce(lambda result, x: result + x, range(1, S+1)))
