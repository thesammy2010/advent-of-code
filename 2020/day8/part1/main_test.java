package tests.day8.part1;

import advent_of_code.day8.part1.*;

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
    public void testAction() {
        assertEquals(part1.Action("nop +0"), Arrays.asList(1, 0));
        assertEquals(part1.Action("jmp -3"), Arrays.asList(-3, 0));
        assertEquals(part1.Action("acc +6"), Arrays.asList(1, 6));
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(5, part1.main("tests/day8/input_test.txt"));
    }

}