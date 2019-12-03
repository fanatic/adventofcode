#!/bin/env python
""" Advent of code """

from day_1 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,expected", [
    ("12", 2),
    ("14", 2),
    ("1969", 654),
    ("100756", 33583),
])
def test_part_1(test_input, expected):
    assert part_1(test_input.splitlines()) == expected


@pytest.mark.parametrize("test_input,expected", [
    ("14", 2),
    ("1969", 966),
    ("100756", 50346),
])
def test_part_2(test_input, expected):
    assert part_2(test_input.splitlines()) == expected
