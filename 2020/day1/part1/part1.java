package advent_of_code.day1.part1;

import java.util.ArrayList;
import java.util.List;

import advent_of_code.helpers.Helpers;

public class part1 {
    
    public static int checkSum(int x, int y)  {

        if (x + y == 2020) {
            return x*y;
        } else {
            return 0;
        }
    }

    public static List<List<Integer>> makePairs(List<Integer> data) {
        
        List<List<Integer>> listToReturn = new ArrayList<>();
        int length = data.size();

		for(int i = 0; i < length; i++) {
            for(int j = 0; j < length; j++) {
                if (i != j) {
                    ArrayList<Integer> tempList = new ArrayList<Integer>();
                    tempList.add(data.get(i)); tempList.add(data.get(j));
                    listToReturn.add(tempList);
                }
            }
        }
        return listToReturn;
    }

    public static int castToInt(String line) throws Exception {
        return Integer.parseInt(line);
    }

    public static List<Integer> main(String filepath) throws Exception {

        // read file
        List<String> lines = Helpers.readFileLineByLine(filepath);

        // list of numbers
        List<Integer> numbers = new ArrayList<Integer>();
        for (int i = 0; i < lines.size(); i++) {
            numbers.add(castToInt(lines.get(i)));
        }

        // generate pairs
        List<List<Integer>> pairs = makePairs(numbers);
        List<Integer> returnValues = new ArrayList<Integer>();

        for (int i = 0; i < pairs.size(); i++) {
            List<Integer> row = pairs.get(i);
            int temp = checkSum(row.get(0), row.get(1));

            if (temp != 0) {
                returnValues.add(row.get(0));
                returnValues.add(row.get(1));
                returnValues.add(temp);
                return returnValues;
            }

        }

        return returnValues;

    }

}
