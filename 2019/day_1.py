#!/usr/bin/env python3
""" Advent of code """


def part_1(fh):
    sum = 0
    for line in fh:
        x = int(line)
        sum += int(x/3)-2
    return sum


def part_2(fh):
    sum = 0
    for line in fh:
        x = int(line)
        while True:
            x = int(x/3)-2

            if x <= 0:
                break
            else:
                sum += x

    return sum


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('day_1_input', 'r')))
    print('Part 2 is: ', part_2(open('day_1_input', 'r')))


if __name__ == "__main__":
    main()
