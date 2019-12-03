#!/usr/bin/env python3
""" Advent of code """


def run(input, noun, verb):
    data = [int(x) for x in input.split(',')]
    if noun > 0:
        data[1] = noun
    if verb > 0:
        data[2] = verb

    for i in range(0, len(data), 4):
        op = data[i]
        if op == 99:
            break
        in1 = data[i+1]
        in2 = data[i+2]
        out = data[i+3]

        if op == 1:
            data[out] = data[in1] + data[in2]
        elif op == 2:
            data[out] = data[in1] * data[in2]

    return data[0]


def part_1(input):
    return run(input, 12, 2)


def part_2(input):
    for noun in range(100):
        for verb in range(100):
            output = run(input, noun, verb)
            if output == 19690720:
                return 100 * noun + verb
    return 0


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('day_2_input', 'r').read()))
    print('Part 2 is: ', part_2(open('day_2_input', 'r').read()))


if __name__ == "__main__":
    main()
