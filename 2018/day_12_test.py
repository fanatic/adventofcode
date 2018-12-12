

#!/bin/env python
""" Advent of code """

from day_12 import part_1
import pytest
pytest.main()

sample_input = [
    "initial state: #..#.#..##......###...###",
    "",
    "...## => #",
    "..#.. => #",
    ".#... => #",
    ".#.#. => #",
    ".#.## => #",
    ".##.. => #",
    ".#### => #",
    "#.#.# => #",
    "#.### => #",
    "##.#. => #",
    "##.## => #",
    "###.. => #",
    "###.# => #",
    "####. => #"
]


@pytest.mark.parametrize("input,sum", [
    (sample_input, 325),
])
def test_part_1(input, sum):
    assert part_1(input, 20) == sum
