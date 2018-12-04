#!/usr/bin/env python3

""" Advent of code """
import re


def part_1(fh):
    fabric = []
    for i in range(2000):
        row = []
        for j in range(2000):
            row.append([])
        fabric.append(row)

    overlaps = 0
    for line in fh:
        m = re.match(r"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", line)
        for x in range(int(m.group(2)), int(m.group(2)) + int(m.group(4))):
            for y in range(int(m.group(3)), int(m.group(3)) + int(m.group(5))):
                # print("#%s - %d, %d" % (m.group(1), x, y))
                fabric[x][y].append(m.group(1))
                if len(fabric[x][y]) == 2:
                    overlaps += 1
                    #print("Found overlap at", x, y, fabric[x][y])
    return overlaps


def part_2(fh):
    fabric = []
    for i in range(2000):
        row = []
        for j in range(2000):
            row.append([])
        fabric.append(row)

    overlaps = 0
    groups = []
    for line in fh:
        m = re.match(r"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", line)
        groups.append(m.group(1))
        for x in range(int(m.group(2)), int(m.group(2)) + int(m.group(4))):
            for y in range(int(m.group(3)), int(m.group(3)) + int(m.group(5))):
                fabric[x][y].append(m.group(1))
                if len(fabric[x][y]) == 2:
                    overlaps += 1

    for i in range(2000):
        for j in range(2000):
            if len(fabric[i][j]) > 1:
                for group in fabric[i][j]:
                    try:
                        groups.remove(group)
                    except ValueError:
                        pass

    return groups


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_3', 'r')))
    print('Part 2 is: ', part_2(open('input_3', 'r')))


if __name__ == "__main__":
    main()
