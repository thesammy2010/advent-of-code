package advent_of_code.day1;

import advent_of_code.day1.part1.*;
import advent_of_code.day1.part2.*;

import java.util.List;

public class run {
    public static void main(String[] args) throws Exception {
        
        System.out.println("Part 1");
        List<Integer> output1 = part1.main("advent_of_code/day1/input.txt");
        System.out.println(output1);

        System.out.println("Part 2");
        List<Integer> output2 = part2.main("advent_of_code/day1/input.txt");
        System.out.println(output2);
    }

}
