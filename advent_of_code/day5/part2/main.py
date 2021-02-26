from typing import List

from advent_of_code.day5.part1.main import Seat, create_seat
from advent_of_code.helpers.helpers import read_file_line_by_line


def get_missing_seat(seat_ids: List[int]) -> int:

    for i in range(len(seat_ids)):

        left_seat: int = seat_ids[i]
        potential_seat: int = left_seat + 1
        right_seat: int = seat_ids[i + 1]

        if potential_seat + 1 == right_seat:
            return potential_seat


    return 0


def main(filepath: str) -> int:
    lines: List[str] = read_file_line_by_line(filepath=filepath)
    seats: List[Seat] = [create_seat(line=line) for line in lines]
    seat_ids: List[int] = sorted([i.seat_id for i in seats])

    return get_missing_seat(seat_ids=seat_ids)
