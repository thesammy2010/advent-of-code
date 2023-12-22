import os
from typing import List

def read_file(filepath: str) -> str:
    with open(file=filepath, mode="r") as file:
        return file.read()

def read_file_line_by_line(filepath: str) -> List[str]:
    with open(file=filepath, mode="r") as file:
        return [i.replace("\n", "") for i in file.readlines()]
