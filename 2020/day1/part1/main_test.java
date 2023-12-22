package tests.day1.part1;

import advent_of_code.day1.part1.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class main_test {
    
    @Test                                               
    @DisplayName("checkSum")   
    public void testCheckSum() {
        
        assertEquals(0, part1.checkSum(1, 2));          
        assertEquals(2019, part1.checkSum(2019, 1));          

    }

    @Test
    @DisplayName("makePairs")
    public void testMakePairs(){

        List<List<Integer>> expectedResult = new ArrayList<List<Integer>>();
        expectedResult.add(Arrays.asList(1, 2));
        expectedResult.add(Arrays.asList(1, 3));
        expectedResult.add(Arrays.asList(2, 1));
        expectedResult.add(Arrays.asList(2, 3));
        expectedResult.add(Arrays.asList(3, 1));
        expectedResult.add(Arrays.asList(3, 2));
        
        List<List<Integer>> calculatedResult = part1.makePairs(Arrays.asList(1, 2, 3));
        assertEquals(expectedResult, calculatedResult);

    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(part1.main("tests/day1/input_test.txt"), Arrays.asList(1721, 299, 514579));
    }

}