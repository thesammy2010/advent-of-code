package tests.day3.part2;

import advent_of_code.day3.part2.*;
import advent_of_code.helpers.Helpers;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {
    
    @Test                                               
    @DisplayName("slope")   
    public void testLoop() throws Exception {
        List<String> rows = Helpers.readFileLineByLine("tests/day3/input_test.txt");
        assertEquals(part2.slope(rows, 1, 1), 2);
        assertEquals(part2.slope(rows, 3, 1), 7);
        assertEquals(part2.slope(rows, 5, 1), 3);
        assertEquals(part2.slope(rows, 7, 1), 4);
        assertEquals(part2.slope(rows, 1, 2), 2);
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        long count = part2.main("tests/day3/input_test.txt");
        assertEquals(336, count);
    }

}