#!/bin/env python
""" Advent of code """


def part_1_sum(input_str):
    "Day 1 Algorithm"
    sum = 0
    for i, c in enumerate(input_str):
        x = int(c)
        i2 = (i + 1) % len(input_str)
        if x == int(input_str[i2]):
            sum += x
    print("Part 1 Sum is", sum)


def part_2_sum(input_str):
    "Day 1 Algorithm"
    sum = 0
    s = int(len(input_str) / 2)
    for i, c in enumerate(input_str):
        x = int(c)
        i2 = (i + s) % len(input_str)
        if x == int(input_str[i2]):
            sum += x
    print("Part 2 Sum is", sum)


def main():
    "Main Entrypoint"
    input_str = open('input_1', 'r').read().strip()
    part_1_sum(input_str)
    part_2_sum(input_str)


if __name__ == "__main__":
    main()
