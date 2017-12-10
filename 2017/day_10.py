#!/bin/env python
""" Advent of code """


def part_1(input_lengths):
    "Day 10 Algorithm"

    circular_list = list(range(256))
    current_position = 0
    skip_size = 0

    # Knot Hash Algorithm Round
    for length in input_lengths:
        print_list(circular_list, current_position)

        print("\nLength is %d" % length)

        if length > 0:
            # Find the sublist to be reversed
            substr_end = (length + current_position - 1) % len(circular_list)
            print_list(circular_list, current_position, substr_end)

            # Reverse that section
            reverse_sublist(circular_list, current_position, substr_end)
            print_list(circular_list, current_position, substr_end)

        # Move current position forward
        current_position = (current_position + length +
                            skip_size) % len(circular_list)
        skip_size += 1

    print_list(circular_list, current_position)
    return circular_list[0] * circular_list[1]


def part_2(input_lengths):
    "Day 10 Algorithm"

    input_lengths = ascii_code(input_lengths)
    input_lengths += [17, 31, 73, 47, 23]
    print(input_lengths)

    # Perform 64 rounds reach sparse_hash
    circular_list = list(range(256))
    current_position = 0
    skip_size = 0

    for i in range(64):
        # Knot Hash Algorithm Round
        for length in input_lengths:
            if length > 0:
                # Find the sublist to be reversed
                substr_end = (length + current_position -
                              1) % len(circular_list)

                # Reverse that section
                reverse_sublist(circular_list, current_position, substr_end)

            # Move current position forward
            current_position = (current_position + length +
                                skip_size) % len(circular_list)
            skip_size += 1

    sparse_hash = circular_list
    print(sparse_hash)

    # Reduce into dense_hash
    dense_hash = sparse_to_dense(sparse_hash)
    print(dense_hash)

    hex_str = ""
    for i in dense_hash:
        hex_str += "%0.2x" % i
    return hex_str


def ascii_code(s):
    return [ord(c) for c in s]


def sparse_to_dense(hsh):
    dense_hash = []
    for i in range(16):
        tmp = 0
        for j in range(16):
            tmp ^= hsh[i * 16 + j]
        dense_hash.append(tmp)
    return dense_hash


def reverse_sublist(lst, start, end):
    if start <= end:
        lst[start:end + 1] = lst[start:end + 1][::-1]
    else:
        sublist = lst[start:] + lst[:end + 1]
        sublist = sublist[::-1]
        lst[start:] = sublist[:len(lst) - start]
        lst[:end + 1] = sublist[len(lst) - start:]
    return lst


def print_list(l, current_position, substr_end=-1):
    for idx, i in enumerate(l):
        if idx == current_position and substr_end == current_position:
            print("([%d])" % i, end=' ')
        elif idx == current_position and substr_end != -1:
            print("([%d]" % i, end=' ')
        elif idx == current_position:
            print("[%d]" % i, end=' ')
        elif idx == substr_end:
            print("%d)" % i, end=' ')
        else:
            print("%d" % i, end=' ')
    print()


def main():
    "Main Entrypoint"
    input_str = open('input_10', 'r').read().strip()
    input_lengths = [int(i) for i in input_str.split(',')]
    answer = part_1(input_lengths)
    print("Part 1 Answer:", answer)
    answer = part_2(input_str)
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
