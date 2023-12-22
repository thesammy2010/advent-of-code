from typing import List, Iterable

from advent_of_code.day8.part1.main import action
from advent_of_code.helpers.helpers import read_file_line_by_line


def change(rs: List[str], index_to_change: int) -> List[str]:
    rsn: List[str] = rs
    if rs[index_to_change][:3] == "nop":
        rsn[index_to_change] = rs[index_to_change].replace("nop", "jmp")
    else:
        rsn[index_to_change] = rs[index_to_change].replace("jmp", "nop")

    return rsn


def main(filepath: str) -> int:

    rows: List[str] = read_file_line_by_line(filepath=filepath)
    orig_rows: List[str] = read_file_line_by_line(filepath=filepath)
    changable_indexes: Iterable[int] = iter(
        [i for i in range(len(rows)) if rows[i].split()[0] in ["nop", "jmp"]]
    )
    indexes: List[int] = []; ind: int = 0; acc: int = 0

    while True:
        new_index, new_acc = action(row=rows[ind])
        ind += new_index

        # swap condition
        if ind in indexes:
            switched_index: int = next(changable_indexes)
            indexes = []; ind = 0; acc = 0
            rows = orig_rows
            rows = change(rs=rows, index_to_change=switched_index)
            # there's some random memory leak here
            orig_rows = read_file_line_by_line(filepath=filepath)  
            assert orig_rows == read_file_line_by_line(filepath=filepath)
            continue

        # terminated successfully
        elif ind >= len(rows):
            acc += new_acc
            break
        
        else:
            indexes.append(ind)
            acc += new_acc

    return acc