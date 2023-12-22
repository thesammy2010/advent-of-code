package advent_of_code.day3.part1;

import java.util.ArrayList;
import java.util.List;

import advent_of_code.helpers.Helpers;

public class part1 {

    public static boolean loop(int lengthOfRow, int index, String row) {
        
        char pos = row.charAt(index % lengthOfRow);

        if (pos == "#".charAt(0)) {
            return true;
        } else {
            return false;
        }
    }

    public static int main(String filepath) throws Exception {

        List<String> lines = Helpers.readFileLineByLine(filepath);

        int count = 0;
        int index = 0;

        for (int i = 0; i < lines.size(); i++) {
            boolean res = loop(lines.get(i).length(), index, lines.get(i));
            if (res) {
                count++;
            }
            index += 3;
        }

        return count;
    
    }

}