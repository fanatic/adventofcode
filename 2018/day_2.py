#!/usr/bin/env python3
""" Advent of code """
from collections import defaultdict


def part_1(fh):
    two_count, three_count = 0, 0
    for line in fh:
        c2, c3 = part_1_sub(line)
        if c2 == 1:
            two_count += 1
        if c3 == 1:
            three_count += 1

    return (two_count, three_count)


def part_1_sub(line):
    has_two, has_three = False, False
    counts = {}
    counts = defaultdict(lambda: 0, counts)
    for c in line:
        counts[c] += 1
    for c, count in counts.items():
        if count == 2:
            has_two = True
        elif count == 3:
            has_three = True

    return (has_two, has_three)


def part_2(lines):
    max_letters_in_common = []
    for a in lines:
        for b in lines:
            if a != b:
                letters_in_common = part_2_sub(a, b)
                if len(letters_in_common) > len(max_letters_in_common):
                    max_letters_in_common = letters_in_common
    return ''.join(max_letters_in_common)


def part_2_sub(a, b):
    letters_in_common = []
    for i in range(0, len(a)-1):
        if a[i] == b[i]:
            letters_in_common.append(a[i])
    return letters_in_common


def main():
    "Main Entrypoint"
    two_count, three_count = part_1(open('input_2', 'r'))
    print('Part 1 is: ', two_count * three_count)
    print('Part 2 is: ', part_2(open('input_2', 'r').readlines()))


if __name__ == "__main__":
    main()
