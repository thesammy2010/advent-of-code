package tests.day6.part2;

import advent_of_code.day6.part2.*;
import advent_of_code.helpers.Helpers;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.HashMap;
import static java.util.Map.entry;    


import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {
    
    @Test                                               
    @DisplayName("compileGroups")   
    public void testCompileGroups() throws Exception {
        String file = Helpers.readFile("tests/day6/input_test.txt");
        List<List<String>> expectedResult = Arrays.asList(
            Arrays.asList("abc"),
            Arrays.asList("a", "b", "c"),
            Arrays.asList("ab", "ac"),
            Arrays.asList("a", "a", "a", "a"),
            Arrays.asList("b"),
            Arrays.asList("abcx", "abcy", "abcz")
        );
        assertEquals(part2.compileGroups(file), expectedResult);
    }

    @Test                                               
    @DisplayName("groupData")   
    public void testGroupData() {
        HashMap<String, Integer> expectedResult1 = new HashMap<String, Integer>();
        expectedResult1.put("a", 1);
        expectedResult1.put("b", 1);
        expectedResult1.put("c", 1);
        expectedResult1.put("count", 1);
        assertEquals(part2.groupData(Arrays.asList("abc")), expectedResult1);    
        HashMap<String, Integer> expectedResult2 = new HashMap<String, Integer>();
        expectedResult2.put("a", 1);
        expectedResult2.put("b", 1);
        expectedResult2.put("c", 1);
        expectedResult2.put("count", 3);
        assertEquals(part2.groupData(Arrays.asList("a", "b", "c")), expectedResult2);  
        HashMap<String, Integer> expectedResult3 = new HashMap<String, Integer>();
        expectedResult3.put("a", 2);
        expectedResult3.put("b", 1);
        expectedResult3.put("c", 1);
        expectedResult3.put("count", 2);
        assertEquals(part2.groupData(Arrays.asList("ab", "ac")), expectedResult3);
        HashMap<String, Integer> expectedResult4 = new HashMap<String, Integer>();
        expectedResult4.put("a", 4);
        expectedResult4.put("count", 4);
        assertEquals(part2.groupData(Arrays.asList("a", "a", "a", "a")), expectedResult4);
        HashMap<String, Integer> expectedResult5 = new HashMap<String, Integer>();
        expectedResult5.put("b", 1);
        expectedResult5.put("count", 1);
        assertEquals(part2.groupData(Arrays.asList("b")), expectedResult5);
        HashMap<String, Integer> expectedResult6 = new HashMap<String, Integer>();
        expectedResult6.put("a", 3);
        expectedResult6.put("b", 3);
        expectedResult6.put("c", 3);
        expectedResult6.put("count", 3);
        expectedResult6.put("x", 1);
        expectedResult6.put("y", 1);
        expectedResult6.put("z", 1);
        assertEquals(part2.groupData(Arrays.asList("abcx", "abcy", "abcz")), expectedResult6);
    }
    
    @Test
    @DisplayName("getCountFromGroup")
    public void testGetCountFromGroup() throws Exception {
        HashMap<String, Integer> inputData1 = new HashMap<String, Integer>();
        inputData1.put("a", 1);
        inputData1.put("b", 1);
        inputData1.put("c", 1);
        inputData1.put("count", 1);
        assertEquals(3, part2.getCountFromGroup(inputData1));
        HashMap<String, Integer> inputData2 = new HashMap<String, Integer>();
        inputData2.put("a", 1);
        inputData2.put("b", 1);
        inputData2.put("c", 1);
        inputData2.put("count", 3);
        assertEquals(0, part2.getCountFromGroup(inputData2));
        HashMap<String, Integer> inputData3 = new HashMap<String, Integer>();
        inputData3.put("a", 4);
        inputData3.put("count", 4);
        assertEquals(1, part2.getCountFromGroup(inputData3));
        HashMap<String, Integer> inputData4 = new HashMap<String, Integer>();
        inputData4.put("b", 1);
        inputData4.put("count", 1);
        assertEquals(1, part2.getCountFromGroup(inputData4));
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(9, part2.main("tests/day6/input_test.txt"));
    }

}