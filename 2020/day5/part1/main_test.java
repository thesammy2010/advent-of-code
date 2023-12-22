package tests.day5.part1;

import advent_of_code.day5.part1.*;

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
    @DisplayName("convert")   
    public void testconvert() {
        assertEquals(part1.convert("L"), 0);
        assertEquals(part1.convert("R"), 1);
        assertEquals(part1.convert("LRLRLR"), 21);
    }

    @Test                                               
    @DisplayName("createSeat")   
    public void testcreateSeat() {
        Seat expectedResult = new Seat(44, 5);
        Seat calculatedResult = part1.createSeat("FBFBBFFRLR");
        assertEquals(expectedResult.seatID, 357);
        assertEquals(calculatedResult.row, expectedResult.row);
        assertEquals(calculatedResult.column, expectedResult.column);
        assertEquals(calculatedResult.seatID, expectedResult.seatID);
    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        assertEquals(820, part1.main("tests/day5/input_test.txt"));
    }

}