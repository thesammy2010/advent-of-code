from typing import List

from advent_of_code.helpers.helpers import read_file_line_by_line

class Seat(object):
    def __init__(self, row: int = None, column: int = None) -> None:
        super().__init__()
        self.row: int = row
        self.column: int = column
        self.seat_id = self.get_seat_id()

    def get_seat_id(self) -> int:
        return (self.row * 8) + self.column


def convert(line: str) -> int:
    s: str = line.replace("F", "0").replace("B", "1").replace("L", "0").replace("R", "1")
    n: int = int(s, base=2)
    return n


def create_seat(line: str) -> Seat:
    row_string: str = line[:7]
    column_string: str = line[-3:]
    s: Seat = Seat(row=convert(row_string), column=convert(column_string))
    return s


def main(filepath: str) -> int:
    lines: List[str] = read_file_line_by_line(filepath=filepath)
    seats: List[Seat] = [create_seat(line=line) for line in lines]
    sorted_seat: Seat = sorted(seats, key=lambda s: s.seat_id, reverse=True)[0]
    return sorted_seat.seat_id
