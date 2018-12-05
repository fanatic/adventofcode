#!/usr/bin/env python3

""" Advent of code """

from string import ascii_lowercase


def part_1(input):
    stable = False
    while not stable:
        stable = True
        #print("Looping again", len(input))
        result = ""
        i = 0
        while i < len(input):
            if i < len(input) - 1 and input[i] == input[i+1].swapcase():
                #print("Reacting", input[i], input[i+1])
                i += 2
                stable = False
                continue
            result += input[i]
            i += 1
        input = result
    return result


def part_2(input):
    min_polymer = 999999999
    for c in ascii_lowercase:
        result = input.replace(c, '').replace(c.swapcase(), '')
        l = len(part_1(result))
        print(c, l)
        if l < min_polymer:
            min_polymer = l
    return min_polymer


def main():
    "Main Entrypoint"
    ans = part_1(open('input_5', 'r').read())
    print('Part 1 is: ', len(ans), ans)
    ans = part_2(open('input_5', 'r').read())
    print('Part 2 is: ', ans)


if __name__ == "__main__":
    main()
