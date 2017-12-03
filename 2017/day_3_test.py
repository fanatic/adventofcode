#!/bin/env python
""" Advent of code """

from day_3 import part_1_steps
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,expected", [
    (1, 0),
    (2, 1),
    (3, 2),
    (4, 1),
    (5, 2),
    (6, 1),
    (7, 2),
    (8, 1),
    (9, 2),
    (10, 3),
    (11, 2),
    (12, 3),
    (13, 4),
    (14, 3),
    (15, 2),
    (16, 3),
    (17, 4),
    (18, 3),
    (19, 2),
    (20, 3),
    (21, 4),
    (22, 3),
    (23, 2),
    (24, 3),
    (25, 4),
    (26, 5),
    (1024, 31),
])
def test_part_1_steps(test_input, expected):
    assert part_1_steps(test_input) == expected

# def test_part_1_steps_manual():
#     for i in range(25):
#         steps = part_1_steps(i + 1)
#         print("\n")
