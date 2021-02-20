package advent_of_code.day4.part2;

import advent_of_code.day4.part1.Passport;
import advent_of_code.day4.part1.part1;


import advent_of_code.helpers.Helpers;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class part2 {

    public static boolean checkPassportIsValidV2(Passport p) {

        int byr;
        int iyr;
        int eyr;
        int hgt;

        // byr between (1920 & 2002)
        try {
            byr = Integer.parseInt(p.byr);
            if (!(1920 <= byr) || !(byr <= 2002)) {
                return false;
            }
        } catch (NumberFormatException e) {
            return false;
        }
        
        // iyr between (2010 & 2020)
        try {
            iyr = Integer.parseInt(p.iyr);
            if (!(2010 <= iyr) || !(iyr <= 2020)) {
                return false;
            }
        } catch (NumberFormatException e) {
            return false;
        }
        
        // eyr between (2020 & 2030)
        try {
            eyr = Integer.parseInt(p.eyr);
            if (!(2020 <= eyr) || !(eyr <= 2030)) {
                return false;
            }
        } catch (NumberFormatException e) {
            return false;
        }
        
        // hgt
        Pattern heightPattern = Pattern.compile("[0-9]{2,3}(cm|in)");
        Matcher heightMatcher = heightPattern.matcher(p.hgt);
        if (!heightMatcher.find()) {
            return false;
        } else {
            hgt = Integer.parseInt(p.hgt.substring(0, p.hgt.length() - 2));
            if (p.hgt.contains("cm")) {
                if (!(150 <= hgt) || !(hgt <= 193)) {
                    return false;
                }
            } 
            else if (p.hgt.contains("in")) {
                if (!(59 <= hgt) || !(hgt <= 76)) {
                    return false;
                } 
            } else {
                return false;
            }
        }

        // hcl
        Pattern colourPattern = Pattern.compile("#[0-9,a-f]{6}");
        Matcher colourMatcher = colourPattern.matcher(p.hcl);
        if (!colourMatcher.find()) {
            return false;
        }

        // ecl
        ArrayList<String> allowedColours = new ArrayList<>(7);
        allowedColours.add("amb"); allowedColours.add("blu"); allowedColours.add("brn"); allowedColours.add("gry");
        allowedColours.add("grn"); allowedColours.add("hzl"); allowedColours.add("oth");
        if (!allowedColours.contains(p.ecl)) {
            return false;
        }

        // pid
        Pattern passportPattern = Pattern.compile("[0-9]{9}");
        Matcher passportMatcher = passportPattern.matcher(p.pid);
        if (!passportMatcher.find()) {
            return false;
        } else {
            if (passportMatcher.group(0) != p.pid) {
                return false;
            }
        }

        return true;
    }

    public static int main(String filepath) throws Exception {
        String rawData = Helpers.readFile(filepath);
        ArrayList<String> lines = part1.cleanupInput(rawData);
        int count = 0;
        for (int i = 0; i < lines.size(); i++) {
            Passport p = part1.createPassport(lines.get(i));
            if (part1.checkPassportIsValid(p)) {
                if (checkPassportIsValidV2(p)) {
                    count++;
                }
            }
        }
        return count;
    }
    
}
