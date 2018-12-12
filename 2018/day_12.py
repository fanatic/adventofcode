#!/usr/bin/env python3

""" Advent of code """


def part_1(fh, generations):
    offset = 10
    initial_state = [False] * offset
    spreads = []
    for line in fh:
        if len(initial_state) == offset:
            for c in line.split(' ')[2]:
                if c == '#':
                    initial_state.append(True)
                elif c == '.':
                    initial_state.append(False)
            continue

        if len(line.rstrip()) == 0:
            continue

        spread = []
        for c in line.split(' ')[0]:
            if c == '#':
                spread.append(True)
            elif c == '.':
                spread.append(False)
        result = line.split(' ')[2][0] == '#'
        spreads.append((spread, result))

    state = initial_state
    print_gen(0, state, offset)
    print()
    for i in range(generations):
        state = next(state, spreads)
        print_gen(i+1, state, offset)

        pot_sum = 0
        for i in range(len(state)):
            if state[i]:
                pot_sum += i-offset
        print(' ', pot_sum)

    return pot_sum


def next(state, spreads):
    next_state = []
    next_state.append(next_pot(([False]*4) + state[0:1], spreads))
    next_state.append(next_pot(([False]*3) + state[0:2], spreads))
    for i in range(len(state)):
        s = state[i:i+5]
        if len(s) < 5:
            s += [False] * (5 - len(s))
        # print(''.join(map(lambda x: '#' if x else '.', s)))
        next_state.append(next_pot(s, spreads))

    return next_state


def next_pot(s, spreads):
    if s == [False, False, False, False]:
        return False
    for spread in spreads:
        if spread[0] == s:
            return spread[1]
    return False


def print_gen(i, state, offset):
    print('%2d:' % i, ''.join(
        map(lambda x: '#' if x else '.', state[7:34])), end='')


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_12', 'r'), 20))
    print('Part 2 is: ', part_1(open('input_12', 'r'), 200))
    # Look for pattern in the output


if __name__ == "__main__":
    main()
