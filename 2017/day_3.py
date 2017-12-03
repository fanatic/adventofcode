#!/bin/env python
""" Advent of code """


def part_1_steps(x):
    "Day 3 Algorithm"

    # Data
    #                     1,1,1,1,1,1,1,1,1,1,2,2,2,2,2,2, 2,2,2,2,3,3
    # 1, 2,3,4,5,6,7,8,9, 0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5, 6,7,8,9,0,1
    #
    # Ring
    # 1  2                3                                4
    #
    # Ring Position
    #                     0,1,2,3,4,5,6,7,8,9,...
    #
    # Edge Position
    #                     0,1,2,3,0,1,2,3,0,1,2,3,0,1,2,3, 0,1,2,3,4,5
    #
    # Distance to Center
    # 0, 0,1,0,1,0,1,0,1, 1,0,1,2,1,0,1,2,1,0,1,2,1,0,1,2, 2,1,0,1,2,3
    #
    # Solution
    # 0, 1,2,1,2,1,2,1,2, 3,2,3,4,3,2,3,4,3,2,3,4,3,2,3,4, 5,4,3,4,5,6

    # First find which ring
    ring = 1
    while (2 * ring - 1) ** 2 <= x - 1:
        ring += 1
    print("Ring of", x, "is", ring, "(", (2 * ring - 1) ** 2, ")")

    # Exit early so we don't have to check first ring
    if ring == 1:
        return 0

    # Find edge length of ring
    edge = (ring - 1) * 2
    print("Edge length of ring", ring, "is", edge)

    # Find position on ring
    pos = x - ((2 * (ring - 1) - 1) ** 2) - 1
    print("Position of", x, "is", pos)

    # Find position on edge
    edge_pos = pos % edge
    print("Edge position of", x, "is", edge_pos)

    # Number of steps is distance to edge center (ring-1) plus distance from center
    if edge_pos < edge / 2 - 1:
        dist_to_center = edge - edge_pos - int(edge / 2) - 1
    else:
        dist_to_center = edge_pos - int(edge / 2) + 1
    print("Distance to center of ", x, "is", dist_to_center)

    return dist_to_center + ring - 1


def part_2(stop):
    # Create temporary array and start in the middle
    SIZE = 100
    matrix = [[0 for x in range(SIZE)] for y in range(SIZE)]
    x, y = int(SIZE / 2), int(SIZE / 2)
    v = -1
    direction = "east"

    # Stop when the value reaches the passed stop point (problem input)
    while v < stop:
        # Determine value to write
        if v == -1:
            v = 1
        else:
            v = matrix[x - 1][y - 1]
            v += matrix[x - 1][y]
            v += matrix[x - 1][y + 1]
            v += matrix[x][y - 1]
            v += matrix[x][y + 1]
            v += matrix[x + 1][y - 1]
            v += matrix[x + 1][y]
            v += matrix[x + 1][y + 1]

        # Write value
        matrix[x][y] = v
        print("matrix[", x, "][", y, "] = ", v)

        # Change direction if necessary
        if not(x == 50 and y == 50):
            if direction == "east" and matrix[x][y + 1] == 0:
                direction = "north"
            elif direction == "north" and matrix[x - 1][y] == 0:
                direction = "west"
            elif direction == "west" and matrix[x][y - 1] == 0:
                direction = "south"
            elif direction == "south" and matrix[x + 1][y] == 0:
                direction = "east"

        # Move cursor
        if direction == "east":
            x = x + 1
        elif direction == "north":
            y = y + 1
        elif direction == "west":
            x = x - 1
        elif direction == "south":
            y = y - 1

    return v


def main():
    "Main Entrypoint"
    input_str = open('input_3', 'r').read().strip()
    #steps = part_1_steps(int(input_str))
    #print("Part 1 Steps:", steps)
    val = part_2(int(input_str))
    print("Part 2 Value:", val)


if __name__ == "__main__":
    main()
