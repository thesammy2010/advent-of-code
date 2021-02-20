package tests.day4.part1;

import advent_of_code.day4.part1.*;

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
    @DisplayName("cleanupInput")   
    public void testcleanupInput() {

        String rawResult = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929";
        ArrayList<String> expectedResult = new ArrayList<String>(2);
        expectedResult.add("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm");
		expectedResult.add("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929");

        assertEquals(part1.cleanupInput(rawResult), expectedResult);
    }

    @Test                                               
    @DisplayName("createPassport")   
    public void testCreatePassport() {

        String rawResult = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm";
        Passport expectedResult = new Passport("1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147");
        Passport calculatedResult = part1.createPassport(rawResult);
        assertEquals(calculatedResult.byr, expectedResult.byr);
        assertEquals(calculatedResult.iyr, expectedResult.iyr);
        assertEquals(calculatedResult.eyr, expectedResult.eyr);
        assertEquals(calculatedResult.hgt, expectedResult.hgt);
        assertEquals(calculatedResult.hcl, expectedResult.hcl);
        assertEquals(calculatedResult.ecl, expectedResult.ecl);
        assertEquals(calculatedResult.pid, expectedResult.pid);
        assertEquals(calculatedResult.cid, expectedResult.cid);

    }

    @Test                                               
    @DisplayName("checkPassportIsValid")   
    public void testCheckPassportIsValid() {

        Passport validPassport = new Passport("1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147");
        Passport invalidPassport = new Passport("1937", "2017", "2020", "183cm", "#fffffd", null, "860033327", "147");
        assertTrue(part1.checkPassportIsValid(validPassport));
        assertFalse(part1.checkPassportIsValid(invalidPassport));

    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        int count = part1.main("tests/day4/input_test.txt");
        assertEquals(2, count);
    }

}