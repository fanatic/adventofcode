

#!/bin/env python
""" Advent of code """

from day_9 import part_1
import pytest
pytest.main()


@pytest.mark.parametrize("players,points,high_score", [
    (9, 25, 32),
    (9, 48, 63),
    (1, 48, 95),
    (10, 1618, 8317),
    (13, 7999, 146373),
    (17, 1104, 2764),
    (21, 6111, 54718),
    (30, 5807, 37305),
    (419, 71052, 412117),
    (419, 71052*100, 3444129546),
])
def test_part_1(players, points, high_score):
    assert part_1(players, points) == high_score
