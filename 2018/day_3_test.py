

#!/bin/env python
""" Advent of code """

from day_3 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,overlaps", [
    (["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"], 4),
])
def test_part_1(test_input, overlaps):
    assert part_1(test_input) == overlaps


@pytest.mark.parametrize("test_input,id", [
    (["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"], ['3']),
])
def test_part_2(test_input, id):
    assert part_2(test_input) == id
