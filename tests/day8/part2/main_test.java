package tests.day8.part2;

import advent_of_code.day8.part2.*;
import advent_of_code.helpers.Helpers;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {
    
    @Test                                               
    @DisplayName("change")   
    public void testChange() throws Exception {
        ArrayList<String> rawResult = Helpers.readFileLineByLine("tests/day8/input_test.txt");
        ArrayList<String> expectedResult = new ArrayList<String>();
        expectedResult.add("jmp +0"); expectedResult.add("acc +1"); expectedResult.add("jmp +4");
        expectedResult.add("acc +3"); expectedResult.add("jmp -3"); expectedResult.add("acc -99");
        expectedResult.add("acc +1"); expectedResult.add("jmp -4"); expectedResult.add("acc +6");
        assertEquals(part2.Change(rawResult, 0), expectedResult);
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(8, part2.main("tests/day8/input_test.txt"));
    }

}