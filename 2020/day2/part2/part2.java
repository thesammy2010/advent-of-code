package advent_of_code.day2.part2;

import java.util.ArrayList;
import java.util.List;
// import java.util.Exception;

import advent_of_code.helpers.Helpers;
import advent_of_code.day2.part1.Policy;
import advent_of_code.day2.part1.part1;


public class part2 {

    public static boolean checkPolicy(Policy p) {

        int[] indexes = new int[]{p.lowerBound - 1, p.upperBound - 1};
 
        if (
            p.password.charAt(indexes[0]) == p.character.charAt(0) && 
            p.password.charAt(indexes[1]) != p.character.charAt(0)
        ) {
            return true;
        } if (
            p.password.charAt(indexes[1]) == p.character.charAt(0) && 
            p.password.charAt(indexes[0]) != p.character.charAt(0)
        ) {
            return true;
        } else {
            return false;
        }

    }
    // public static void main(String[] args) {

    //     String test = "1-3 a: abcde";
    //     System.out.println(test);

    //     Policy p = part1.createPolicy(test);
    //     System.out.printf(
    //         "lowerBound: %s, upperBound: %s character: %s password %s", 
    //         p.lowerBound, p.upperBound, p.character, p.password
    //     );
    //     System.out.println("");
    //     System.out.print(checkPolicy(p));

    // }

    public static int main(String filepath) throws Exception {

        List<String> lines = Helpers.readFileLineByLine(filepath);

        int count = 0;

        for (int i = 0; i < lines.size(); i++) {
            Policy p = part1.createPolicy(lines.get(i));
            if (checkPolicy(p)) {
                count++;
            }
        }

        return count;
    }

}