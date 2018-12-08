#!/usr/bin/env python3

""" Advent of code """

import re
from collections import namedtuple

Node = namedtuple('Node', ['children', 'metadata'])


def part_1(str):
    input = list(map(int, str.split()))
    root_node, l = parse_node(input)

    print(root_node)

    return metadata_sum(root_node)


def parse_node(input):
    num_children = input[0]
    num_metadata = input[1]

    children = []
    i = 2
    for j in range(num_children):
        child, l = parse_node(input[i:])
        i += l
        children.append(child)

    metadata = []
    for j in range(i, i+num_metadata):
        metadata.append(input[j])

    return Node(children, metadata), i + num_metadata


def metadata_sum(node):
    s = sum(node.metadata)
    for c in node.children:
        s += metadata_sum(c)
    return s


def part_2(str):
    input = list(map(int, str.split()))
    root_node, l = parse_node(input)

    return value(root_node)


def value(node):
    if len(node.children) == 0:
        return sum(node.metadata)

    s = 0
    for i in node.metadata:
        try:
            s += value(node.children[i-1])
        except IndexError:
            pass

    return s


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_8', 'r').read()))
    print('Part 2 is: ', part_2(open('input_8', 'r').read()))


if __name__ == "__main__":
    main()
