#!/bin/env python
""" Advent of code """


def part_1_checksum(input_str):
    "Day 2 Algorithm"
    sum = 0
    for line in input_str.splitlines():
        chars = line.split("\t")
        digits = [int(c) for c in chars]
        diff = max(digits) - min(digits)
        sum += diff

    print("Part 1 Sum is", sum)


def part_2_checksum(input_str):
    "Day 2 Algorithm"
    sum = 0
    for line in input_str.splitlines():
        chars = line.split("\t")
        digits = [int(c) for c in chars]

        def outer():
            for i in digits:
                for j in digits:
                    if i != j and i % j == 0:
                        return int(i / j)

        sum += outer()

    print("Part 2 Sum is", sum)


def main():
    "Main Entrypoint"
    input_str = open('input_2', 'r').read().strip()
    part_1_checksum(input_str)
    part_2_checksum(input_str)


if __name__ == "__main__":
    main()
