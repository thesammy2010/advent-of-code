package advent_of_code.day5.part1;

import advent_of_code.helpers.Helpers;
import java.util.ArrayList;
import java.util.List;
import java.util.Collections;


public class part1 {

    public static int convert(String line) {
        String binaryString = line.replaceAll("F", "0").replaceAll("B", "1").replace("L", "0").replace("R", "1");
        int binaryNumber = Integer.parseInt(binaryString, 2);
        return binaryNumber;
    }

    public static Seat createSeat(String line) {
        int row = convert(line.substring(0, 7));
        int column = convert(line.substring(7, 10));
        Seat s = new Seat(row, column);        
        return s;
    }

    public static int main(String filepath) throws Exception {
            
        ArrayList<String> rawData = Helpers.readFileLineByLine(filepath);
        ArrayList<Integer> seatIDs = new ArrayList<Integer>();
        for (int i = 0; i < rawData.size(); i++) {
            Seat s = createSeat(rawData.get(i));
            seatIDs.add(s.seatID);
        }
        Collections.sort(seatIDs, Collections.reverseOrder());

        return seatIDs.get(0);
    
    }
    
}
