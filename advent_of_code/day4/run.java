package advent_of_code.day4;

import advent_of_code.day4.part1.*;
import advent_of_code.day4.part2.*;

public class run {
    public static void main(String[] args) throws Exception {
        
        System.out.println("Part 1");
        int output1 = part1.main("advent_of_code/day4/input.txt");
        System.out.println(output1);

        System.out.println("Part 2");
        int output2 = part2.main("advent_of_code/day4/input.txt");
        System.out.println(output2);
    
    }

}
