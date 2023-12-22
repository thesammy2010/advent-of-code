package advent_of_code.day5.part2;

import advent_of_code.day5.part1.*;
// import advent_of_code.day5.part1.Seat;
import advent_of_code.helpers.Helpers;

import java.util.ArrayList;
import java.util.List;
import java.util.Collections;


public class part2 {

    public static int getMissingSeat(ArrayList<Integer> seatIDs) {

        for (int i = 0; i < seatIDs.size(); i++) {
            int leftSeat = seatIDs.get(i);
            int potentialSeat = leftSeat + 1;
            int rightSeat = seatIDs.get(i + 1);

            if (potentialSeat + 1 == rightSeat) {
                return potentialSeat;
            }
        }

        return 0;

    }

    public static int main(String filepath) throws Exception {
            
        ArrayList<String> rawData = Helpers.readFileLineByLine(filepath);
        ArrayList<Integer> seatIDs = new ArrayList<Integer>();
        for (int i = 0; i < rawData.size(); i++) {
            Seat s = part1.createSeat(rawData.get(i));
            seatIDs.add(s.seatID);
        }
        Collections.sort(seatIDs);

        return getMissingSeat(seatIDs);
    
    } 
    
}
