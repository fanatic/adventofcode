#!/usr/bin/env python3

""" Advent of code """

import re
from string import ascii_lowercase

grid_size = 400
sum_max = 10000


def part_1(fh):
    grid = []
    for i in range(grid_size):
        row = []
        for j in range(grid_size):
            row.append('.')
        grid.append(row)

    # Parse Coords
    coordinates = []
    for line in fh:
        m = re.match(r"(\d+), (\d+)", line)
        coordinates.append(((int(m.group(1)), int(m.group(2)))))

    # Fill in
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            grid[i][j] = closest_coordinate((i, j), coordinates)

    # Draw grid
    if grid_size < 40:
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[j][i] == -1:
                    print('.', end='')
                else:
                    print(ascii_lowercase[grid[j][i]], end='')
            print()

    areas = {}
    # Find largest area
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            cell_coord = grid[i][j]
            if cell_coord == '.':
                continue
            elif i == 0 or i == len(grid)-1 or j == 0 or j == len(grid[i])-1:
                areas[cell_coord] = -1
            else:
                try:
                    if areas[cell_coord] != -1:
                        areas[cell_coord] += 1
                except KeyError:
                    areas[cell_coord] = 1

    print(areas)
    return max(areas.values())


def closest_coordinate(a, coords):
    closest_coord_index = -1
    closest_distance = 999999
    for i, b in enumerate(coords):
        if a == b:
            # print(a, ascii_lowercase[i].upper())
            return i
        distance = manhattan_distance(a, b)
        if distance < closest_distance:
            closest_coord_index = i
            closest_distance = distance

    # Check for another coordinate tied in distance
    for i, b in enumerate(coords):
        distance = manhattan_distance(a, b)
        if distance == closest_distance and closest_coord_index != i:
            return -1

    # print(a, ascii_lowercase[closest_coord_index])
    return closest_coord_index


def manhattan_distance(a, b):
    return abs(b[0] - a[0]) + abs(b[1] - a[1])


def part_2(fh):
    grid = []
    for i in range(grid_size):
        row = []
        for j in range(grid_size):
            row.append('.')
        grid.append(row)

    # Parse Coords
    coordinates = []
    for line in fh:
        m = re.match(r"(\d+), (\d+)", line)
        coordinates.append(((int(m.group(1)), int(m.group(2)))))

    # Fill in
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if in_region((i, j), coordinates):
                grid[i][j] = '#'

    # Draw grid
    if grid_size < 40:
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[j][i] == '#':
                    print(grid[j][i], end='')
                else:
                    print(pixel((j, i), coordinates), end='')
            print()

    region_count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[j][i] == '#':
                region_count += 1

    return region_count


def in_region(a, coords):
    sum_to_all_coordinates = 0
    for b in coords:
        distance = manhattan_distance(a, b)
        sum_to_all_coordinates += distance

    if sum_to_all_coordinates < sum_max:
        return True

    return False


def pixel(a, coords):
    for i, b in enumerate(coords):
        if a == b:
            return ascii_lowercase[i].upper()
    return '.'


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_6', 'r')))
    print('Part 2 is: ', part_2(open('input_6', 'r')))


if __name__ == "__main__":
    main()
