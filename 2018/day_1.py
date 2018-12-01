#!/usr/bin/env python3
""" Advent of code """


def part_1(fh):
    sum = 0
    for line in fh:
        x = int(line[1:])
        if line[0] == '-':
            x = -x
        sum += x
    return sum


def part_2(fh):
    seen_frequencies = [0]
    sum = 0
    while True:
        for line in fh:
            x = int(line[1:])
            if line[0] == '-':
                x = -x
            sum += x
            if sum in seen_frequencies:
                return sum
            seen_frequencies.append(sum)
        if not isinstance(fh, list):
            fh.seek(0)


def main():
    "Main Entrypoint"
    fh = open('input_1', 'r')
    print('Part 1 is: ', part_1(fh))
    print('Part 2 is: ', part_2(fh))


if __name__ == "__main__":
    main()
