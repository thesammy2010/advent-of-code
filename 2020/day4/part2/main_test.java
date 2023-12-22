package tests.day4.part2;

import advent_of_code.day4.part1.Passport;
import advent_of_code.day4.part2.*;

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
    @DisplayName("checkPassportIsValidV2")   
    public void testCheckPassportIsValidV2() {

        // byr iyr eyr hgt hcl ecl pid cid
        Passport p1 = new Passport("1980", "2012", "2030", "74in", "#623a2f", "grn", "087499704", "147");
        Passport p2 = new Passport("1989", "2014", "2029", "165cm", "#a97842", "blu", "896056539", "129");
        Passport p3 = new Passport("2001", "2015", "2022", "164cm", "#888785", "hzl", "545766238", "88");
        Passport p4 = new Passport("1944", "2010", "2021", "158cm", "#b6652a", "blu", "093154719", "88");
        Passport p5 = new Passport("1926", "2018", "1972", "186cm", "#18171d", "amb", "186cm", "100");
        Passport p6 = new Passport("1946", "2019", "1967", "170cm", "#602927", "grn", "012533040", null);
        Passport p7 = new Passport("1992", "2012", "2020", "182cm", "dab227", "brn", "021572410", "277");
        Passport p8 = new Passport("2007", "2023", "2038", "59cm", "74454a", "zzz", "3556412378", "277");
        
        Passport troubleMaker = new Passport("1932", null, "2024", "157cm", "#6b5442", "brn", "4275535874", null);

        assertTrue(part2.checkPassportIsValidV2(p1));
        assertTrue(part2.checkPassportIsValidV2(p2));
        assertTrue(part2.checkPassportIsValidV2(p3));
        assertTrue(part2.checkPassportIsValidV2(p4));
        assertFalse(part2.checkPassportIsValidV2(p5));
        assertFalse(part2.checkPassportIsValidV2(p6));
        assertFalse(part2.checkPassportIsValidV2(p7));
        assertFalse(part2.checkPassportIsValidV2(p8));
        assertFalse(part2.checkPassportIsValidV2(troubleMaker));

    }

    @Test
    @DisplayName("main")
    public void testMain() throws Exception {
        int count = part2.main("tests/day4/input_test.txt");
        assertEquals(2, count);
    }

}