

#!/bin/env python
""" Advent of code """

from day_6 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,size", [
    (["1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"], 17),
])
def test_part_1(test_input, size):
    assert part_1(test_input) == size


@pytest.mark.parametrize("test_input,size", [
    (["1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"], 16),
])
def test_part_2(test_input, size):
    assert part_2(test_input) == size
