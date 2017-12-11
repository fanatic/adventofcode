#!/bin/env python
""" Advent of code """


def part_1(input_str):
    "Day 10 Algorithm"

    path = input_str.split(',')

    x, y, z = 0, 0, 0

    for step in path:
        if step == 'n':
            z -= 1
            y += 1
        elif step == 'ne':
            z -= 1
            x += 1
        elif step == 'se':
            y -= 1
            x += 1
        elif step == 's':
            y -= 1
            z += 1
        elif step == 'sw':
            x -= 1
            z += 1
        elif step == 'nw':
            y += 1
            x -= 1

    return cube_distance(x, y, z)


def cube_distance(x, y, z):
    "https://www.redblobgames.com/grids/hexagons/#distances"

    return int((abs(x) + abs(y) + abs(z)) / 2)


def part_2(input_str):
    "Day 10 Algorithm"

    path = input_str.split(',')

    x, y, z = 0, 0, 0
    max_distance = 0

    for step in path:
        if step == 'n':
            z -= 1
            y += 1
        elif step == 'ne':
            z -= 1
            x += 1
        elif step == 'se':
            y -= 1
            x += 1
        elif step == 's':
            y -= 1
            z += 1
        elif step == 'sw':
            x -= 1
            z += 1
        elif step == 'nw':
            y += 1
            x -= 1
        max_distance = max(max_distance, cube_distance(x, y, z))

    return max_distance


def main():
    "Main Entrypoint"
    input_str = open('input_11', 'r').read().strip()
    answer = part_1(input_str)
    print("Part 1 Answer:", answer)
    answer = part_2(input_str)
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
