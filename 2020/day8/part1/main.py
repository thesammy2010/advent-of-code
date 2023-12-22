from typing import List, Tuple

from advent_of_code.helpers.helpers import read_file_line_by_line


def action(row: str) -> Tuple[int, int]:
    act, val = row.split()
    if act == "nop":
        return (1, 0)
    elif act == "acc":
        return (1, int(val))
    else:
        return (int(val), 0)


def main(filepath: str) -> int:

    rows: List[str] = read_file_line_by_line(filepath=filepath)
    indexes: List[int] = []
    ind: int = 0
    acc: int = 0

    while True:
        new_index, new_acc = action(row=rows[ind])
        ind += new_index

        if ind in indexes:
            break

        indexes.append(ind)
        acc += new_acc

    return acc