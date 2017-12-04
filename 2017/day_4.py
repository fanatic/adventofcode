#!/bin/env python
""" Advent of code """


def part_1(lines):
    "Day 4 Algorithm"
    valid_lines = 0
    for line in lines:
        if not has_dup_words(line):
            valid_lines += 1
    return valid_lines


def has_dup_words(line):
    seen_words = {}
    for word in line.split(" "):
        if word in seen_words:
            return True
        seen_words[word] = True
    return False


def part_2(lines):
    "Day 4 Algorithm"
    valid_lines = 0
    for line in lines:
        if not has_anagram_words(line):
            valid_lines += 1
    return valid_lines


def has_anagram_words(line):
    seen_words = {}
    for word in line.split(" "):
        w = ''.join(sorted(word))
        if w in seen_words:
            return True
        seen_words[w] = True
    return False


def main():
    "Main Entrypoint"
    input_str = open('input_4', 'r').read().strip()
    answer = part_1(input_str.splitlines())
    print("Part 1 Answer:", answer)
    answer = part_2(input_str.splitlines())
    print("Part 2 Answer:", answer)


if __name__ == "__main__":
    main()
