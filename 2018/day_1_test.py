#!/bin/env python
""" Advent of code """

from day_1 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,expected", [
    ("+1\n+1\n+1", 3),
    ("+1\n+1\n-2", 0),
    ("-1\n-2\n-3", -6),
])
def test_part_1(test_input, expected):
    assert part_1(test_input.splitlines()) == expected


@pytest.mark.parametrize("test_input,expected", [
    ("+1\n-1", 0),
    ("+3\n+3\n+4\n-2\n-4", 10),
    ("-6\n+3\n+8\n+5\n-6", 5),
    ("+7\n+7\n-2\n-7\n-4", 14),
])
def test_part_2(test_input, expected):
    assert part_2(test_input.splitlines()) == expected
