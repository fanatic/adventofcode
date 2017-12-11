#!/bin/env python
""" Advent of code """

import os
from day_11 import part_1
import pytest
pytest.main([os.path.basename(__file__)])


@pytest.mark.parametrize("test_input,expected", [
    ('ne,ne,ne', 3),
    ('ne,ne,sw,sw', 0),
    ('ne,ne,s,s', 2),
    ('se,sw,se,sw,sw', 3),
])
def test_part_1(test_input, expected):
    assert part_1(test_input) == expected
