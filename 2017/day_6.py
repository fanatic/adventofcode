#!/bin/env python
""" Advent of code """

import math


def part_1(banks):
    "Day 6 Algorithm"
    # print(banks)
    seen_states = []
    cycles = 0
    while bank_hash(banks) not in seen_states:
        seen_states.append(bank_hash(banks))

        # Find the memory bank with the most blocks
        max_block_idx = max(range(len(banks)), key=banks.__getitem__)

        cur_bank = (max_block_idx + 1) % len(banks)
        b = math.ceil(banks[max_block_idx] / len(banks))
        # print("Bank #", max_block_idx, "has the most blocks (",
        #       banks[max_block_idx], "), so chosen for redistribution")
        # print("The ", banks[max_block_idx],
        #       "blocks are spread out over the", len(banks), "memory banks:", b, "each")
        while cur_bank != max_block_idx:
            if banks[max_block_idx] >= b:
                banks[cur_bank] = banks[cur_bank] + b
                banks[max_block_idx] = banks[max_block_idx] - b
            else:
                banks[cur_bank] = banks[cur_bank] + banks[max_block_idx]
                banks[max_block_idx] = 0

            cur_bank = (cur_bank + 1) % len(banks)

        # print(banks)
        cycles += 1

    return cycles, banks


def bank_hash(banks):
    return ",".join(str(x) for x in banks)


def part_2(banks):
    "Day 6 Algorithm"
    seen_states = []
    cycles = 0
    while bank_hash(banks) not in seen_states:
        seen_states.append(bank_hash(banks))

        # Find the memory bank with the most blocks
        max_block_idx = max(range(len(banks)), key=banks.__getitem__)

        cur_bank = (max_block_idx + 1) % len(banks)
        b = math.ceil(banks[max_block_idx] / len(banks))
        while cur_bank != max_block_idx:
            if banks[max_block_idx] >= b:
                banks[cur_bank] = banks[cur_bank] + b
                banks[max_block_idx] = banks[max_block_idx] - b
            else:
                banks[cur_bank] = banks[cur_bank] + banks[max_block_idx]
                banks[max_block_idx] = 0

            cur_bank = (cur_bank + 1) % len(banks)

        cycles += 1

    return cycles


def main():
    "Main Entrypoint"
    input_str = open('input_6', 'r').read().strip()
    banks = [int(bank) for bank in input_str.split("\t")]

    answer, banks = part_1(banks)
    print("Part 1 Answer:", answer)
    answer = part_2(banks)
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
