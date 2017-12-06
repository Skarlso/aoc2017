"""
Part 1 Of Day 6.
"""
MEMORY = [2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14]

seen = [[]]

# if seen.count(current_config) > 1 -> combination has been already encountered
steps = 0
while True:
    blocks = max(MEMORY)
    block_index = MEMORY.index(blocks) # retrieves the index of the max value
    MEMORY[block_index] = 0
    i = (block_index + 1) % len(MEMORY)
    while blocks != 0:
        MEMORY[i] += 1
        blocks -= 1
        i = (i + 1) % len(MEMORY)
    if seen.count(MEMORY) > 0:
        break
    else:
        seen.append(MEMORY[:])
    steps += 1

# print(seen)
print("Took %d steps for repeated config to appear." % steps)
print("Current MEMORY: %s" % MEMORY)
