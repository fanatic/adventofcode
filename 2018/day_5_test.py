

#!/bin/env python
""" Advent of code """

from day_5 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,polymer", [
    ("", ""),
    ("a", "a"),
    ("aA", ""),
    ("aAa", "a"),
    ("abBA", ""),
    ("abAB", "abAB"),
    ("aabAAB", "aabAAB"),
    ("dabAcCaCBAcCcaDA", "dabCBAcaDA"),
])
def test_part_1(test_input, polymer):
    assert part_1(test_input) == polymer


@pytest.mark.parametrize("test_input,l", [
    ("dabAcCaCBAcCcaDA", 4),
])
def test_part_2(test_input, l):
    assert part_2(test_input) == l
