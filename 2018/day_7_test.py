

#!/bin/env python
""" Advent of code """

from day_7 import part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,order", [
    ([
        "Step C must be finished before step A can begin.",
        "Step C must be finished before step F can begin.",
        "Step A must be finished before step B can begin.",
        "Step A must be finished before step D can begin.",
        "Step B must be finished before step E can begin.",
        "Step D must be finished before step E can begin.",
        "Step F must be finished before step E can begin.",
    ], "CABDFE"),
])
def test_part_1(test_input, order):
    assert part_1(test_input) == order


@pytest.mark.parametrize("test_input,time", [
    ([
        "Step C must be finished before step A can begin.",
        "Step C must be finished before step F can begin.",
        "Step A must be finished before step B can begin.",
        "Step A must be finished before step D can begin.",
        "Step B must be finished before step E can begin.",
        "Step D must be finished before step E can begin.",
        "Step F must be finished before step E can begin.",
    ], 15),
])
def test_part_2(test_input, time):
    assert part_2(test_input) == time
