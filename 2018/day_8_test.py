

#!/bin/env python
""" Advent of code """

from day_8 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,sum", [
    ("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 138),
])
def test_part_1(test_input, sum):
    assert part_1(test_input) == sum


@pytest.mark.parametrize("test_input,value", [
    ("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 66),
])
def test_part_2(test_input, value):
    assert part_2(test_input) == value
