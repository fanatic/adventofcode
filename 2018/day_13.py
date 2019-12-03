#!/usr/bin/env python3

""" Advent of code """


def part_1(fh):
    m = []
    for line in fh:
        m.append(list(line))

    for y in range(len(m)):
        print(''.join(m[y]))

    while True:
        skip_next_y = -1
        for y in range(len(m)):
            skip_next_x = False
            for x in range(len(m[y])):
                if skip_next_x:
                    skip_next_x = False
                elif skip_next_y > -1 and skip_next_y == x:
                    skip_next_y = -1
                elif m[y][x] == '^':
                    m[y][x] = "|"
                    m[y-1][x] = "^"
                elif m[y][x] == 'v':
                    m[y][x] = "|"
                    m[y+1][x] = "v"
                    skip_next_y = x
                elif m[y][x] == '<':
                    m[y][x] = "-"
                    m[y][x-1] = "<"
                elif m[y][x] == '>':
                    m[y][x] = "-"
                    m[y][x+1] = ">"
                    skip_next_x = True

        for y in range(len(m)):
            print(''.join(m[y]))

        break


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_13', 'r')))


if __name__ == "__main__":
    main()
