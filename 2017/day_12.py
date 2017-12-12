#!/bin/env python
""" Advent of code """

import re

ps = {}


def parse(lines):
    for line in lines:
        parsed = re.findall(
            r'(\d+) <-> ([,\d ]+)', line)
        pid = int(parsed[0][0])
        pid_peers = parsed[0][1].split(', ')
        ps[pid] = [int(i) for i in pid_peers]


def part_1():
    "Day 12 Algorithm"

    group = set()
    find_pid_peers(group, 0)
    return len(group)


def find_pid_peers(group_members, pid):
    group_members.add(pid)
    for peer in ps[pid]:
        if peer not in group_members:
            find_pid_peers(group_members, peer)
    return group_members


def part_2():
    "Day 12 Algorithm"

    groups = []
    pids_remaining = set(ps.keys())
    while pids_remaining:
        group = set()
        find_pid_peers(group, pids_remaining.pop())
        pids_remaining -= group
        groups.append(group)
    return len(groups)


def main():
    "Main Entrypoint"
    input_str = open('input_12', 'r').read().strip()
    parse(input_str.splitlines())
    answer = part_1()
    print("Part 1 Answer:", answer)
    answer = part_2()
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
