package tests.day1.part2;

import advent_of_code.day1.part2.*;

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
        
        assertEquals(0, part2.checkSum(1, 2, 3));          
        assertEquals(2018, part2.checkSum(2018, 1, 1));          

    }

    @Test
    @DisplayName("makeTriplets")
    public void testMakeTriplets(){

        List<List<Integer>> expectedResult = new ArrayList<List<Integer>>();
        
        expectedResult.add(Arrays.asList(1, 2, 3));
        expectedResult.add(Arrays.asList(1, 3, 2));
        expectedResult.add(Arrays.asList(2, 1, 3));
        expectedResult.add(Arrays.asList(2, 3, 1));
        expectedResult.add(Arrays.asList(3, 1, 2));
        expectedResult.add(Arrays.asList(3, 2, 1));


        List<List<Integer>> calculatedResult = part2.makeTriplets(Arrays.asList(1, 2, 3));
        assertEquals(expectedResult, calculatedResult);

    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(part2.main("tests/day1/input_test.txt"), Arrays.asList(979, 366, 675, 241861950));
    }

}
