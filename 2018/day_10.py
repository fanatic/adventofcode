#!/usr/bin/env python3

""" Advent of code """

import re
from collections import namedtuple

Point = namedtuple('Point', ['position', 'velocity'])


def part_1(fh):
    # Parse notes
    points = []
    for line in fh:
        m = re.match(
            r"position=< *([-\d]+), *([-\d]+)> velocity=< *([-\d]+), *([-\d]+)>", line)
        position = (int(m.group(1)), int(m.group(2)))
        velocity = (int(m.group(3)), int(m.group(4)))

        points.append(Point(position, velocity))

    i = 0
    while True:
        print("On second", i)
        print_positions(points)
        for j in range(len(points)):
            p = points[j]
            new_position = (p.position[0] + p.velocity[0],
                            p.position[1] + p.velocity[1])
            points[j] = Point(new_position, p.velocity)
        i += 1


def print_positions(points):
    graph = {}
    x_range = (0, 0)
    y_range = (0, 0)
    for p in points:
        if p.position[0] < x_range[0]:
            x_range = (p.position[0], x_range[1])
        if p.position[0] > x_range[1]:
            x_range = (x_range[0], p.position[0])
        if p.position[1] < y_range[0]:
            y_range = (p.position[1], y_range[1])
        if p.position[1] > y_range[1]:
            y_range = (y_range[0], p.position[1])
        try:
            graph[p.position[0]][p.position[1]] = True
        except KeyError:
            graph[p.position[0]] = {p.position[1]: True}

    if x_range[1] - x_range[0] > 200:
        return

    for y in range(y_range[0], y_range[1]+1):
        for x in range(x_range[0], x_range[1]+1):
            try:
                if graph[x][y]:
                    print('#', end='')
                    continue
            except KeyError:
                pass
            print('.', end='')
        print()


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_10', 'r')))


if __name__ == "__main__":
    main()
