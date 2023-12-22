package advent_of_code.day8.part2;

import java.util.ArrayList;
import java.util.List;
import java.util.Arrays;
import java.util.HashSet;

import advent_of_code.day8.part1.*;
import advent_of_code.helpers.Helpers;

public class part2 {

    public static ArrayList<String> Change(ArrayList<String> rs, int indexToChange) {        
        
        if (rs.get(indexToChange).charAt(0) == "nop".charAt(0)) {
            rs.set(indexToChange, rs.get(indexToChange).replaceAll("nop", "jmp"));
        } else if (rs.get(indexToChange).charAt(0) == "jmp".charAt(0)) {
            rs.set(indexToChange, rs.get(indexToChange).replaceAll("jmp", "rs"));
        }
        else {
            return new ArrayList<>();
        }
        return rs;
    }

    public static int main(String filepath) throws Exception {
        ArrayList<String> rows = Helpers.readFileLineByLine(filepath);
        HashSet<Integer> indexes = new HashSet<Integer>();
        int ind = 0;
        int acc = 0;

        ArrayList<Integer> changableIndexes = new ArrayList<Integer>();
        for (int i = 0; i < rows.size(); i++) {
            if (rows.get(i).charAt(0) == "nop".charAt(0) || rows.get(i).charAt(0) == "jmp".charAt(0)) {
                changableIndexes.add(i);
            }
        }
        int nextIndex = -1;

        while (true) {

            List<Integer> vals = part1.Action(rows.get(ind));
            int newIndex = vals.get(0);
            int newAcc = vals.get(1);
            ind += newIndex;

            // swap
            if (indexes.contains(ind)) {
                nextIndex += 1;
                if (nextIndex >= rows.size()) {
                    break;
                }
                rows = Change(Helpers.readFileLineByLine(filepath), changableIndexes.get(nextIndex));
                ind = 0;
                acc = 0;
                indexes = new HashSet<Integer>();               
                continue;
            }

            // terminate
            if (ind >= rows.size()) {
                acc += newAcc;
                return acc;
                // break;
            }

            indexes.add(ind);
            acc += newAcc;

        }
        return acc;
    }
    
}
