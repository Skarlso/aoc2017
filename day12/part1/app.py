"""
day 12 Part 1
"""
# pylint: disable-msg=C0103
import re

pipes = {}

with open("../sample.txt", "r") as f:
    data = f.readlines()

for line in data:
    m = re.match(r"(?P<pipe>\d+)\s<->\s(?P<connections>[\d+,\s]+)", line)
    pipe = m.group("pipe")
    connections = m.group("connections")
    pipes[int(pipe)] = map(int, connections.strip().split(','))
    print(pipe)
    print(connections.strip().split(','))
