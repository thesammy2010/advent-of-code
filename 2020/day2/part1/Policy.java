package advent_of_code.day2.part1;

public class Policy {
    
    public int lowerBound;
    public int upperBound;
    public String character;
    public String password;

    public Policy() {};
    
    public Policy(int x, int y, String z, String t) {
        lowerBound = x;
        upperBound = y;
        character = z;
        password = t;
    }

}
