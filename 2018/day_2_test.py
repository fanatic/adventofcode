

#!/bin/env python
""" Advent of code """

from day_2 import part_1, part_1_sub, part_2
import pytest
pytest.main()


@pytest.mark.parametrize("test_input,has_two,has_three", [
    ("abcdef", False, False),
    ("bababc", True, True),
    ("abbcde", True, False),
    ("abcccd", False, True),
    ("aabcdd", True, False),
    ("abcdee", True, False),
    ("ababab", False, True),
])
def test_part_1_sub(test_input, has_two, has_three):
    assert part_1_sub(test_input) == (has_two, has_three)


@pytest.mark.parametrize("test_input,two_count,three_count", [
    (["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"], 4, 3),
])
def test_part_1(test_input, two_count, three_count):
    assert part_1(test_input) == (two_count, three_count)


@pytest.mark.parametrize("test_input,common_letters", [
    (["abcde",
      "fghij",
      "klmno",
      "pqrst",
      "fguij",
      "axcye",
      "wvxyz"], "fgij"),
])
def test_part_2(test_input, common_letters):
    assert part_2(test_input) == common_letters
