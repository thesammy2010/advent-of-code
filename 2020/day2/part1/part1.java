package advent_of_code.day2.part1;

import java.util.ArrayList;
import java.util.List;

import advent_of_code.helpers.Helpers;
public class part1 {
    
    public static Policy createPolicy(String input) {
       
        // input = "1-4 l: hbljb"
        String[] parts = input.split("\\s");
        // parts = ["1-4", "l:", "hbljb"]
        String[] bounds = parts[0].split("-");
        // bounds = ["1", "4"]
        int lowerBound = Integer.parseInt(bounds[0]);
        // lowerBound = 1
        int upperBound = Integer.parseInt(bounds[1]);
        // upperBound = 4
        String character = parts[1].replace(":", "");
        // character = "l"
        String password = parts[2];
        // password = "hbljb"
        
        // this.lowerBound = lowerBound;

        Policy p = new Policy(lowerBound, upperBound, character, password);

        return p;
    
    }

    public static boolean checkPolicy(Policy p) {
        
        int count = 0;

        for(int i=0; i < p.password.length(); i++) {
            if(p.password.charAt(i) == p.character.charAt(0)) {
                count++;
            }
        }
        if (count >= p.lowerBound && count <= p.upperBound) {
            return true;
        } else {
            return false;
        }
    }
    // public static void main(String[] args) {

    //     String test = "1-4 l: hbljb";
    //     System.out.println(test);

    //     Policy p = createPolicy(test);
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
            Policy p = createPolicy(lines.get(i));
            if (checkPolicy(p)) {
                count++;
            }
        }

        return count;
    }

}