package advent_of_code.day4.part1;

import advent_of_code.day4.part1.Passport;
import advent_of_code.helpers.Helpers;
import java.util.ArrayList;
import java.util.List;

public class part1 {

    public static ArrayList<String> cleanupInput(String rawInput) {
        ArrayList<String> dataToReturn = new ArrayList<String>();
        String[] splitData = rawInput.split("\n\n");
        for (int i = 0; i < splitData.length; i++) {
            dataToReturn.add(
                splitData[i].replace("\n", " ")
            );
        }
        return dataToReturn;
    }

    public static Passport createPassport(String rawString) {
        Passport p = new Passport();
        String[] keyValuePairsString = rawString.split(" ");
        for (int i = 0; i < keyValuePairsString.length; i++) {
            String[] temp = keyValuePairsString[i].split(":");
            String key = temp[0];
            String value = null;
            try {
                value = temp[1];
            } catch (ArrayIndexOutOfBoundsException e) {
                System.out.println(temp);
            }
            switch (key) {
                case "byr":
                    p.byr = value; break;
                case "iyr":
                    p.iyr = value; break;
                case "eyr":
                    p.eyr = value; break;
                case "hgt":
                    p.hgt = value; break;
                case "hcl":
                    p.hcl = value; break;
                case "ecl":
                    p.ecl = value; break;
                case "pid":
                    p.pid = value; break;
                case "cid":
                    p.cid = value; break;
            }
        }
        return p;
    }

    public static boolean checkPassportIsValid(Passport p) {
        if (p.byr == null) {
            return false;
        }
        if (p.iyr == null) {
            return false;
        }
        if (p.eyr == null) {
            return false;
        }
        if (p.hgt == null) {
            return false;
        }
        if (p.hcl == null) {
            return false;
        }
        if (p.ecl == null) {
            return false;
        }
        if (p.pid == null) {
            return false;
        }
        return true;
    }

    public static int main(String filepath) throws Exception {
        String rawData = Helpers.readFile(filepath);
        ArrayList<String> lines = cleanupInput(rawData);
        int count = 0;
        for (int i = 0; i < lines.size(); i++) {
            Passport p = createPassport(lines.get(i));
            if (checkPassportIsValid(p)) {
                count++;
            }
        }
        return count;
    }
    
}
