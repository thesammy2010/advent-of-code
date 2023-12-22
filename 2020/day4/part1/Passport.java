package advent_of_code.day4.part1;

public class Passport {

    public String byr;
    public String iyr;
    public String eyr;
    public String hgt;
    public String hcl;
    public String ecl;
    public String pid;
    public String cid;

    public Passport() {

    }

    public Passport(String BYR, String IYR, String EYR, String HGT, String HCL, String ECL, String PID, String CID) {
        this.byr = BYR;
        this.iyr = IYR;
        this.eyr = EYR;
        this.hgt = HGT;
        this.hcl = HCL;
        this.ecl = ECL;
        this.pid = PID;
        this.cid = CID;
    }
    
}
