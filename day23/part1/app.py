"""
day 23 Part 1
"""
# pylint: disable-msg=C0103
import sys

registers = {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0}
filename = sys.argv[1]
instructions = []

mul_count = 0

with open(filename, "r") as f:
    for line in f:
        instructions.append(line.rstrip('\n'))


transform = lambda val: registers[val] if val.isalpha() else int(val)


position = 0
while position < len(instructions):
    inst = instructions[position].split(" ")
    if inst[0] == "set":
        registers[inst[1]] = transform(inst[2])
    elif inst[0] == "sub":
        registers[inst[1]] -= transform(inst[2])
    elif inst[0] == "mul":
        registers[inst[1]] *= transform(inst[2])
        mul_count += 1
    elif inst[0] == "jnz":
        if transform(inst[1]) != 0:
            position += transform(inst[2])
            continue
    position += 1

print("Mul called: %d times" % mul_count)
