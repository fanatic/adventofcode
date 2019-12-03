

#!/bin/env python
""" Advent of code """

from day_13 import part_1
import pytest
pytest.main()


@pytest.mark.parametrize("input,location", [
    ([
        '/->-\\        ',
        '|   |  /----\\',
        '| /-+--+-\\  |',
        '| | |  | v  |',
        '\\-+-/  \\-+--/',
        '  \\------/   '
    ], (7, 3)),
])
def test_part_1(input, location):
    assert part_1(input) == location
