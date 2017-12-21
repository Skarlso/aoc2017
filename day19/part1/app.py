"""
day 19 Part 1
"""
# pylint: disable-msg=C0103
import sys

pipes = {}

filename = sys.argv[1]

PUZZLE = ""
grid = []

UP = [1, 0]
DOWN = [-1, 0]
LEFT = [0, -1]
RIGHT = [0, 1]


def next_move(direction, x, y):
    """
    Next move return None if there are no more moves left.
    """
    global PUZZLE
    char = grid[y][x]
    # Extract these two as they are doing the same thing
    if char != "+":
        if char.isalpha():
            PUZZLE += char
        y, x = y + direction[0], x + direction[1]
        if grid[y][x] == " ":
            return None, 0, 0
        return direction, x, y
    elif char == "+":
        check_directions = [UP, DOWN, LEFT, RIGHT]
        for chk_dir in check_directions:
            if direction[0] + chk_dir[0] == 0 and direction[1] + chk_dir[1] == 0:
                continue
            if y + chk_dir[0] < len(grid) and x + chk_dir[1] < len(grid[y]) and grid[y + chk_dir[0]][x + chk_dir[1]] != " ":
                y, x = y + chk_dir[0], x + chk_dir[1]
                return chk_dir, x, y
        return None, 0, 0


# TODO: Currently these are arrays of rows. Use strings instead. There is no real need to explode them.
with open(filename, "r") as f:
    for line in f:
        grid.append([c for c in line.rstrip('\n')])

x, y = 0, 0
for i, col in enumerate(grid[0]):
    if col == "|":
        x, y = i, 0
        break

direction, x, y = next_move(UP, x, y)
while direction != None:
    direction, x, y = next_move(direction, x, y)

print("Gathered letters are: %s" % PUZZLE)
