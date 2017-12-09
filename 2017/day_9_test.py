#!/bin/env python
""" Advent of code """

from day_9 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,expected", [
    ("{}", 1),
    ("{{{}}}", 6),
    ("{{},{}}", 5),
    ("{{{},{},{{}}}}", 16),
    ("{<a>,<a>,<a>,<a>}", 1),
    ("{{<ab>},{<ab>},{<ab>},{<ab>}}", 9),
    ("{{<!!>},{<!!>},{<!!>},{<!!>}}", 9),
    ("{{<a!>},{<a!>},{<a!>},{<ab>}}", 3),
])
def test_part_1(test_input, expected):
    assert part_1(test_input) == expected


@pytest.mark.parametrize("test_input,expected", [
    ("<>", 0),
    ("<random characters>", 17),
    ("<<<<>", 3),
    ("<{!>}>", 2),
    ("<!!>", 0),
    ("<!!!>>", 0),
    ("<{o\"i!a,<{i<a>", 10),
])
def test_part_2(test_input, expected):
    assert part_2(test_input) == expected
