"""
day 12 Part 2
"""
# pylint: disable-msg=C0103
import re

pipes = {}

with open("../input.txt", "r") as f:
    data = f.readlines()

for line in data:
    m = re.match(r"(?P<pipe>\d+)\s<->\s(?P<connections>[\d+,\s]+)", line)
    pipe = m.group("pipe")
    connections = m.group("connections")
    pipes[int(pipe)] = list(map(int, connections.strip().split(',')))


seen = []


def travers(program):
    if seen.count(program) != 0:
        return
    seen.append(program)
    for program in pipes[program]:
        travers(program)


groups = 0
for pipe in pipes:
    if seen.count(pipe) != 0:
        continue
    travers(pipe)
    groups += 1

print(groups)
