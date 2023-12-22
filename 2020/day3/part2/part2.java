package advent_of_code.day3.part2;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import advent_of_code.helpers.Helpers;
import advent_of_code.day3.part1.*;

public class part2 {

    public static int slope(List<String> rows, int slide, int lineSkip) {
        
        int count = 0;
        int index = 0;
        int lineNumber = 0;

        while (lineNumber < rows.size()) {

            String row = rows.get(lineNumber);
            boolean res = part1.loop(row.length(), index, row);
            if (res) {
                count++;
            }
            index += slide;
            lineNumber += lineSkip;

        }

        return count;

    }

    public static long main(String filepath) throws Exception {

        List<String> inputFile = Helpers.readFileLineByLine(filepath);

        List<List<Integer>> conditions = new ArrayList<>();
        conditions.add(Arrays.asList(1, 1));
        conditions.add(Arrays.asList(1, 3));
        conditions.add(Arrays.asList(1, 5));
        conditions.add(Arrays.asList(1, 7));
        conditions.add(Arrays.asList(2, 1));

        List<Integer> counts = new ArrayList<Integer>();

        for (int i = 0; i < conditions.size(); i++) {
            counts.add(
                slope(
                    inputFile, 
                    conditions.get(i).get(1), 
                    conditions.get(i).get(0)
                )
            );
        }

        long returnValue = 1; // to handle 32 bit integer overflow
        for (int i = 0; i < counts.size(); i++) {
            returnValue *= counts.get(i);
        }

        return returnValue;
    
    }

}