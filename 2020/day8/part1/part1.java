package advent_of_code.day8.part1;

import java.util.ArrayList;
import java.util.List;
import java.util.Arrays;
import java.util.HashSet;
import advent_of_code.helpers.Helpers;

public class part1 {

    public static List<Integer> Action(String data) {
        List<Integer> action = new ArrayList<>();
        String[] newData = data.split("\\s+", -1);
        String act = newData[0];
        Integer val = Integer.parseInt(newData[1]);

        // System.out.println(val);

        if (act.charAt(0) == "nop".charAt(0)) {
            action = Arrays.asList(1, 0);
        } else if (act.charAt(0) == "acc".charAt(0)) {
            action = Arrays.asList(1, val);
        } else {
            action = Arrays.asList(val, 0);
        }
        
        return action;
    }

    public static int main(String filepath) throws Exception {
        ArrayList<String> rows = Helpers.readFileLineByLine(filepath);
        HashSet<Integer> indexes = new HashSet<Integer>();
        int ind = 0;
        int acc = 0;

        while (true) {

            List<Integer> vals = Action(rows.get(ind));
            int newIndex = vals.get(0);
            int newAcc = vals.get(1);
            ind += newIndex;

            if (indexes.contains(ind)) {
                break;
            }

            indexes.add(ind);
            acc += newAcc;

        }
        return acc;
    }
    
}
