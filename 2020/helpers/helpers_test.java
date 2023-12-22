package tests.helpers;

import advent_of_code.helpers.Helpers;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;



public class helpers_test {

    @Test                                               
    @DisplayName("ReadFileLineByLine")   
    public void testReadFileLineByLine() throws Exception {

        ArrayList<String> expectedResult = new ArrayList<String>() { 
            { 
                add("this"); 
                add("is"); 
                add("a"); 
                add("test"); 
            } 
        }; 
        ArrayList <String> calculatedOutput = Helpers.readFileLineByLine("tests/helpers/helpers_test.txt");

        assertEquals(expectedResult, calculatedOutput);          
    }

    @Test                                               
    @DisplayName("readfile") 
    public void testReadFile() throws Exception {
        String expectedResult = "this\nis\na\ntest\n";
        String calculatedOutput = Helpers.readFile("tests/helpers/helpers_test.txt");
        assertEquals(expectedResult, calculatedOutput);          
    }
    
}