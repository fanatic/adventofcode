

#!/bin/env python
""" Advent of code """

from day_11 import power_level, power_level_grid, part_1, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("x,y,serial_number,power", [
    (3, 5, 8, 4),
    (122, 79, 57, -5),
    (217, 196, 39, 0),
    (101, 153, 71, 4),
])
def test_power_level(x, y, serial_number, power):
    assert power_level(x, y, serial_number) == power


@pytest.mark.parametrize("serial_number,corner,power", [
    (18, (33, 45), 29),
    (42, (21, 61), 30),
])
def test_part_1(serial_number, corner, power):
    assert part_1(power_level_grid(serial_number), 3) == (corner, power)


@pytest.mark.parametrize("serial_number,corner,power", [
    (18, (90, 269, 16), 113),
    (42, (232, 251, 12), 119),
])
def test_part_2(serial_number, corner, power):
    assert part_2(serial_number) == (corner, power)
