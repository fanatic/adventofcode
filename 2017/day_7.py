#!/bin/env python
""" Advent of code """

import re


class Program:
    name = ''
    weight = 0
    children = []
    parent = None

    def __repr__(self):
        return "<Program name:%s weight:%d children:%s>" % (self.name, self.weight, self.children)

    def total_weight(self):
        children_weights = []
        for child in self.children:
            children_weights.append(child.total_weight())

        # A little bit of manual work here for speed..
        # Look at the first printed result - that's the child we care about
        # Then fix the parent weight to equalize and balance the disk based on the output
        if len(set(children_weights)) > 1:
            print(self.name, "disc is unbalanced.")
            print("Children:", [o.name for o in self.children])
            print("Children:", [o.weight for o in self.children])
            print("Weights:", children_weights)
        return self.weight + sum(children_weights)


def part_1(lines):
    "Day 7 Algorithm"

    # Parse file
    programs = {}
    for line in lines:
        parsed = re.findall('([^ ]+) \((\d+)\)(?: -> ([\w, ]+))?', line)
        p = Program()
        p.name = parsed[0][0]
        p.weight = int(parsed[0][1])
        if parsed[0][2] != '':
            p.children = parsed[0][2].split(', ')
        programs[p.name] = p

    # Populate parents
    for _, p in programs.items():
        for child in p.children:
            programs[child].parent = p

    # Find the program with no parent
    for _, p in programs.items():
        if not p.parent:
            return p

    return None


def part_2(lines):
    "Day 7 Algorithm"

    # Parse file
    programs = {}
    for line in lines:
        parsed = re.findall('([^ ]+) \((\d+)\)(?: -> ([\w, ]+))?', line)
        p = Program()
        p.name = parsed[0][0]
        p.weight = int(parsed[0][1])
        if parsed[0][2] != '':
            p.children = parsed[0][2].split(', ')
        programs[p.name] = p

    # Populate children
    for _, p in programs.items():
        c = []
        for child in p.children:
            c.append(programs[child])
        p.children = c

    # Populate parents
    for _, p in programs.items():
        for child in p.children:
            child.parent = p

    # Find the program with no parent
    root = None
    for _, p in programs.items():
        if not p.parent:
            root = p

    root.total_weight()

    return None


def main():
    "Main Entrypoint"
    input_str = open('input_7', 'r').read().strip()

    answer = part_1(input_str.splitlines())
    print("Part 1 Answer:", answer)
    answer = part_2(input_str.splitlines())
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
