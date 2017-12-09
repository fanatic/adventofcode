#!/bin/env python
""" Advent of code """


def part_1(stream):
    "Day 9 Algorithm"

    score = 0
    depth = 0
    in_garbage = False
    ignore_next = False

    for c in stream:
        if ignore_next:
            ignore_next = False
            continue
        elif c == '{' and not in_garbage:
            depth += 1
        elif c == '}' and not in_garbage:
            score += depth
            depth -= 1
        elif c == '<' and not in_garbage:
            in_garbage = True
        elif c == '>' and in_garbage:
            in_garbage = False
        elif c == '!':
            ignore_next = True

    return score


def part_2(stream):
    "Day 9 Algorithm"

    garbage_removed = 0
    depth = 0
    in_garbage = False
    ignore_next = False

    for c in stream:
        if ignore_next:
            ignore_next = False
            continue
        elif c == '{' and not in_garbage:
            depth += 1
        elif c == '}' and not in_garbage:
            depth -= 1
        elif c == '<' and not in_garbage:
            in_garbage = True
        elif c == '>' and in_garbage:
            in_garbage = False
        elif c == '!':
            ignore_next = True
        elif in_garbage:
            garbage_removed += 1

    return garbage_removed


def main():
    "Main Entrypoint"
    input_str = open('input_9', 'r').read().strip()
    answer = part_1(input_str)
    print("Part 1 Answer:", answer)
    answer = part_2(input_str)
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
