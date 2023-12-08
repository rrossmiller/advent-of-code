package day3;

class Coordinate {
    private int row;
    private int col;

    public Coordinate(int row, int col) {
        this.row = row;
        this.col = col;
    }

    public int row() {
        return this.row;
    }

    public int col() {
        return this.col;
    }
}
