#!/bin/env python
""" Advent of code """

import os
from day_10 import part_2, sparse_to_dense, ascii_code
import pytest
pytest.main([os.path.basename(__file__)])


def test_sparse_to_dense():
    assert sparse_to_dense(
        [65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22] * 16) == [64] * 16


def test_ascii_code():
    assert ascii_code("1,2,3") == [49, 44, 50, 44, 51]


@pytest.mark.parametrize("test_input,expected", [
    ("", "a2582a3a0e66e6e86e3812dcb672a272"),
    ("AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"),
    ("1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"),
    ("1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"),
])
def test_part_2(test_input, expected):
    assert part_2(test_input) == expected
