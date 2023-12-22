package advent_of_code.day1.part2;

import java.util.ArrayList;
import java.util.List;

import advent_of_code.helpers.Helpers;
import advent_of_code.day1.part1.*;


public class part2 {
    
    public static int checkSum(int x, int y, int z) {

        if (x + y + z == 2020) {
            return x*y*z;
        } else {
            return 0;
        }

    }

    public static List<List<Integer>> makeTriplets(List<Integer> data) {
        
        List<List<Integer>> listToReturn = new ArrayList<>();
        int length = data.size();

		for(int i = 0; i < length; i++) {
            for(int j = 0; j < length; j++) {
                for(int k = 0; k < length; k++) {
                    if (i != j && i != k && j != k) {
                        ArrayList<Integer> tempList = new ArrayList<Integer>();
                        tempList.add(data.get(i)); 
                        tempList.add(data.get(j)); 
                        tempList.add(data.get(k));
                        listToReturn.add(tempList);
                    }
                }
            }
        }
    
        return listToReturn;
    
    }

    public static List<Integer> main(String filepath) throws Exception {
        
        // read file
        List<String> lines = Helpers.readFileLineByLine(filepath);

        // list of numbers
        List<Integer> numbers = new ArrayList<Integer>();
        for (int i = 0; i < lines.size(); i++) {
            numbers.add(part1.castToInt(lines.get(i)));
        }

        // generate pairs
        List<List<Integer>> triples = makeTriplets(numbers);
        List<Integer> returnValues = new ArrayList<Integer>();

        for (int i = 0; i < triples.size(); i++) {
            List<Integer> row = triples.get(i);
            int temp = checkSum(row.get(0), row.get(1), row.get(2));

            if (temp != 0) {
                returnValues.add(row.get(0));
                returnValues.add(row.get(1));
                returnValues.add(row.get(2));
                returnValues.add(temp);
                return returnValues;
            }
            
        }

        return returnValues;

    }

}
