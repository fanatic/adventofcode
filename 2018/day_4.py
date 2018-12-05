#!/usr/bin/env python3

""" Advent of code """
import re
from collections import namedtuple

Guard = namedtuple('Guard', ['id', 'minute_ranges'])


def part_1(fh):
    fh.sort()
    guards = {}
    for line in fh:
        m = re.match(r"\[[\d-]+ \d{2}:(\d{2})\] (.*)", line)
        minute = int(m.group(1))
        activity = m.group(2)
        if activity == "falls asleep":
            try:
                guards[guard_id] = Guard(
                    guard_id, guards[guard_id].minute_ranges + [minute])
            except KeyError:
                guards[guard_id] = Guard(guard_id, [minute])
        elif activity == "wakes up":
            guards[guard_id] = Guard(
                guard_id, guards[guard_id].minute_ranges + [minute])
        else:
            m = re.match(r"Guard #([\d]+) begins shift", activity)
            guard_id = m.group(1)

    guard_asleep_the_most = Guard(-1, [])
    for guard_id, guard in guards.items():
        if minutes_asleep(guard.minute_ranges) > minutes_asleep(guard_asleep_the_most.minute_ranges):
            guard_asleep_the_most = guard

    minutes = {}
    ranges = guard_asleep_the_most.minute_ranges
    for i in range(0, len(ranges), 2):
        for j in range(ranges[i], ranges[i+1]):
            try:
                minutes[j] += 1
            except KeyError:
                minutes[j] = 1

    most_common_minute = max(minutes, key=minutes.get)

    return (int(guard_asleep_the_most.id), most_common_minute)


def minutes_asleep(minute_ranges):
    total = 0
    for i in range(0, len(minute_ranges), 2):
        total += minute_ranges[i+1] - minute_ranges[i]
    return total


def part_2(fh):
    fh.sort()
    guards = {}
    for line in fh:
        m = re.match(r"\[[\d-]+ \d{2}:(\d{2})\] (.*)", line)
        minute = int(m.group(1))
        activity = m.group(2)
        if activity == "falls asleep":
            try:
                guards[guard_id] = Guard(
                    guard_id, guards[guard_id].minute_ranges + [minute])
            except KeyError:
                guards[guard_id] = Guard(guard_id, [minute])
        elif activity == "wakes up":
            guards[guard_id] = Guard(
                guard_id, guards[guard_id].minute_ranges + [minute])
        else:
            m = re.match(r"Guard #([\d]+) begins shift", activity)
            guard_id = m.group(1)

    chosen_guard_id = -1
    chosen_guard_minute = 0
    chosen_guard_minute_count = 0
    for guard_id, guard in guards.items():
        minutes = {}
        ranges = guard.minute_ranges
        for i in range(0, len(ranges), 2):
            for j in range(ranges[i], ranges[i+1]):
                try:
                    minutes[j] += 1
                except KeyError:
                    minutes[j] = 1
        most_common_minute = max(minutes, key=minutes.get)
        most_common_minute_count = max(minutes.values())
        if most_common_minute_count > chosen_guard_minute_count:
            chosen_guard_id = int(guard_id)
            chosen_guard_minute = most_common_minute
            chosen_guard_minute_count = most_common_minute_count

    return (chosen_guard_id, chosen_guard_minute)


def main():
    "Main Entrypoint"
    ans = part_1(open('input_4', 'r').readlines())
    print('Part 1 is: ', ans[0] * ans[1])
    ans = part_2(open('input_4', 'r').readlines())
    print('Part 2 is: ', ans[0] * ans[1])


if __name__ == "__main__":
    main()
