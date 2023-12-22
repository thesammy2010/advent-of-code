package advent_of_code.day3;

import advent_of_code.day3.part1.*;
import advent_of_code.day3.part2.*;

import java.util.List;

public class run {
    public static void main(String[] args) throws Exception {
        
        System.out.println("Part 1");
        int output1 = part1.main("advent_of_code/day3/input.txt");
        System.out.println(output1);

        System.out.println("Part 2");
        long output2 = part2.main("advent_of_code/day3/input.txt");
        System.out.println(output2);
    }

}
