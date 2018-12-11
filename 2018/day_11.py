#!/usr/bin/env python3

""" Advent of code """


def part_1(grid, square_size):
    grid_size = len(grid)
    largest_grid = ((1, 1), -99999)
    for x in range(grid_size-square_size):
        for y in range(grid_size-square_size):
            total_power = 0
            for i in range(square_size):
                total_power += sum(grid[x+i][y:y+square_size])

            if total_power > largest_grid[1]:
                largest_grid = ((x, y), total_power)
    return largest_grid


def power_level_grid(serial_number):
    grid_size = 300
    grid = []
    for x in range(grid_size):
        row = []
        for y in range(grid_size):
            row.append(power_level(x, y, serial_number))
        grid.append(row)
    return grid


def power_level(x, y, serial_number):
    rack_id = x+10
    power = rack_id * y
    power += serial_number
    power *= rack_id
    power = int(power % 1000/100)
    power -= 5
    return power


def part_2(serial_number):
    largest_grid = ((1, 1, 1), -99999)
    grid = power_level_grid(serial_number)

    for i in range(1, 300):
        sq = part_1(grid, i)
        print(i, sq)
        if sq[1] > largest_grid[1]:
            largest_grid = ((sq[0][0], sq[0][1], i), sq[1])
    return largest_grid[0]


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(2187, 3)[0])
    #print('Part 2 is: ', part_2(18))
    #print('Part 2 is: ', part_2(42))
    print('Part 2 is: ', part_2(2187))


if __name__ == "__main__":
    main()
