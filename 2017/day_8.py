#!/bin/env python
""" Advent of code """

import re


def compute(lines):
    "Day 8 Algorithm"

    # Parse file
    registers = {}
    largest_value_evar = -999999
    for line in lines:
        parsed = re.findall(
            r'(\w+) (\w+) ([-0-9]+) if (\w+) ([><=!]+) ([-0-9]+)', line)
        reg = parsed[0][0]
        action = parsed[0][1]
        amount = int(parsed[0][2])
        condition_reg = parsed[0][3]
        condition_cond = parsed[0][4]
        condition_amount = int(parsed[0][5])

        print("%s %s %d %s %s %d" % (
            reg, action, amount, condition_reg, condition_cond, condition_amount))

        # Check condition
        if ((condition_cond == ">" and registers.get(condition_reg, 0) <= condition_amount) or
                (condition_cond == "<" and registers.get(condition_reg, 0) >= condition_amount) or
                (condition_cond == ">=" and registers.get(condition_reg, 0) < condition_amount) or
                (condition_cond == "<=" and registers.get(condition_reg, 0) > condition_amount) or
                (condition_cond == "==" and registers.get(condition_reg, 0) != condition_amount) or
                (condition_cond == "!=" and registers.get(condition_reg, 0) == condition_amount)):
            print("Skip")
            continue

        # Take action
        if action == "inc":
            registers[reg] = registers.get(reg, 0) + amount
        elif action == "dec":
            registers[reg] = registers.get(reg, 0) - amount

        print("%s = %d" % (reg, registers[reg]))

        if registers[reg] > largest_value_evar:
            largest_value_evar = registers[reg]

    # Find largest
    print("Part 1: Largest value in any register is",
          registers[max(registers, key=registers.get)])
    print("Part 2: Largest value held in any register during the process",
          largest_value_evar)


def main():
    "Main Entrypoint"
    input_str = open('input_8', 'r').read().strip()
    compute(input_str.splitlines())


if __name__ == "__main__":
    main()
