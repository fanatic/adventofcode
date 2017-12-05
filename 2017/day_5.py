#!/bin/env python
""" Advent of code """


def part_1(lines):
    "Day 4 Algorithm"
    pc = 0
    offsets = [int(line) for line in lines]
    steps = 0
    while True:
        if pc >= len(offsets):
            break

        # print_offsets(pc, offsets)

        old_pc = pc
        pc = old_pc + offsets[old_pc]
        offsets[old_pc] += 1
        steps += 1

    # print_offsets(pc, offsets)

    return steps


def print_offsets(pc, offsets):
    for i, offset in enumerate(offsets):
        if pc == i:
            print("(%d) " % offset, end='')
        else:
            print("%d " % offset, end='')
    print("")


def part_2(lines):
    "Day 4 Algorithm"
    pc = 0
    offsets = [int(line) for line in lines]
    steps = 0
    while True:
        if pc >= len(offsets):
            break

        # print_offsets(pc, offsets)

        old_pc = pc
        pc = old_pc + offsets[old_pc]
        if offsets[old_pc] >= 3:
            offsets[old_pc] -= 1
        else:
            offsets[old_pc] += 1
        steps += 1

    # print_offsets(pc, offsets)

    return steps


def main():
    "Main Entrypoint"
    input_str = open('input_5', 'r').read().strip()
    answer = part_1(input_str.splitlines())
    print("Part 1 Answer:", answer)
    answer = part_2(input_str.splitlines())
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
