package tests.day3.part1;

import advent_of_code.day3.part1.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {
    
    @Test                                               
    @DisplayName("loop")   
    public void testLoop() {
        assertEquals(part1.loop(11, 0, "..##......."), false);
        assertEquals(part1.loop(11, 3, "..##......."), true);
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        int count = part1.main("tests/day3/input_test.txt");
        assertEquals(count, 7);
    }

}