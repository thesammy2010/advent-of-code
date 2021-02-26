package advent_of_code.day5.part1;

public class Seat {
    
    public int row;
    public int column;
    public int seatID;
    
    public Seat(int inputRow,int inputColumn) {
        this.row = inputRow;
        this.column = inputColumn;
        this.seatID = this.row*8 + this.column;
    }

}
