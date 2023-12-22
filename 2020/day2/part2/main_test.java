package tests.day2.part2;

import advent_of_code.day2.part2.*;
import advent_of_code.day2.part1.Policy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {

    @Test                                               
    @DisplayName("checkPolicy")   
    public void testCheckPolicy() {

        Policy p1 = new Policy(1, 3, "a", "abcde");
        Policy p2 = new Policy(1, 3, "a", "abade");

        assertEquals(part2.checkPolicy(p1), true);
        assertEquals(part2.checkPolicy(p2), false);

    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        int count = part2.main("tests/day2/input_test.txt");
        assertEquals(count, 1);
    }

}