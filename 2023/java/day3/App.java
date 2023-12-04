import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class App {
    enum Part {
        Part1,
        Part2
    }

    public static void main(String[] args) throws Exception {
        System.out.println("*****");
        List<String> data = readData();
        // data = testData();
        run(data, Part.Part1);
        run(data, Part.Part2);

    }

    private static void run(List<String> data, Part p) {
        List<Integer> nums = new ArrayList<>();
        List<Coordinate> visited = new ArrayList<>();

        // for every line
        for (int i = 0; i < data.size(); i++) {
            char[] line = data.get(i).toCharArray();
            // for every char in the line
            for (int j = 0; j < line.length; j++) {
                char c = line[j];

                // if the char is a symbol, look around it for numbers
                if (!Character.isDigit(c) && c != '.') {
                    switch (p) {
                        case Part1:
                            p1(data, i, j, nums, visited);
                            break;
                        case Part2:
                            if (c == '*')
                                p2(data, i, j, nums, visited);
                            break;
                    }
                }
            }
        }

        int part = 0;
        switch (p) {
            case Part1:
                part = 1;
                break;
            case Part2:
                part = 2;
                break;
        }
        int i = 0;
        for (int x : nums) {
            i += x;
        }
        System.out.println(String.format("pt%d: %d", part, i));
    }

    private static void p1(List<String> data, int i, int j, List<Integer> nums, List<Coordinate> visited) {
        int[][] directions = new int[][] {
                { -1, -1 }, { -1, 0 }, { -1, 1 },
                { 0, -1 }, { 0, 1 },
                { 1, -1 }, { 1, 0 }, { 1, 1 }
        };
        int cols = data.get(0).length();

        // bfs to look around the symbol
        for (int[] dir : directions) {
            int row = i + dir[0];
            int col = j + dir[1];
            var coord = new Coordinate(row, col);
            String line = data.get(row);

            // if the char at the coord is a digit
            if (row >= 0 // checks to not cause IOB exception
                    && row < data.size()
                    && col >= 0
                    && col < cols
                    && Character.isDigit(line.charAt(col)) // char is digit
                    && !coordinateChecked(visited, coord) // idx hasn't been checked already
            ) {
                // visit this coordinate
                visited.add(coord);

                // find the rest of the number
                int l = col - 1;
                int r = col + 1;
                // find leftmost limit
                while (l > 0 && Character.isDigit(line.charAt(l))) {
                    visited.add(new Coordinate(row, l));
                    l--;
                }
                if (!Character.isDigit(line.charAt(l)))
                    l++;
                // find rightmost limit
                while (r < line.length() && Character.isDigit(line.charAt(r))) {
                    visited.add(new Coordinate(row, r));
                    r++;
                }
                nums.add(Integer.parseInt(line.substring(l, r)));
            }

        }
    }

    private static void p2(List<String> data, int i, int j, List<Integer> nums, List<Coordinate> visited) {
        int[][] directions = new int[][] {
                { -1, -1 }, { -1, 0 }, { -1, 1 },
                { 0, -1 }, { 0, 1 },
                { 1, -1 }, { 1, 0 }, { 1, 1 }
        };
        int cols = data.get(0).length();

        int[] gears = new int[2];
        int n = 0;

        // bfs to look around the symbol
        for (int[] dir : directions) {
            int row = i + dir[0];
            int col = j + dir[1];
            var coord = new Coordinate(row, col);
            String line = data.get(row);

            // if the char at the coord is a digit
            if (row >= 0 // checks to not cause IOB exception
                    && row < data.size()
                    && col >= 0
                    && col < cols
                    && Character.isDigit(line.charAt(col)) // char is digit
                    && !coordinateChecked(visited, coord) // idx hasn't been checked already
            ) {
                // visit this coordinate
                visited.add(coord);

                // find the rest of the number
                int l = col - 1;
                int r = col + 1;
                // find leftmost limit
                while (l > 0 && Character.isDigit(line.charAt(l))) {
                    visited.add(new Coordinate(row, l));
                    l--;
                }
                if (!Character.isDigit(line.charAt(l)))
                    l++;
                // find rightmost limit
                while (r < line.length() && Character.isDigit(line.charAt(r))) {
                    visited.add(new Coordinate(row, r));
                    r++;
                }
                gears[n] = Integer.parseInt(line.substring(l, r));
                n++;
            }
        }
        nums.add(gears[0] * gears[1]);
    }

    private static boolean coordinateChecked(List<Coordinate> visited, Coordinate c) {
        for (Coordinate v : visited) {
            if (v.row() == c.row() && v.col() == c.col()) {
                return true;
            }
        }
        return false;
    }

    private static List<String> readData() throws IOException {
        Path path = Paths.get("../data/3.txt");
        // replace newline with space
        String s = Files.readString(path);
        List<String> data = Arrays.asList(s.split("\n"));
        return data;
    }

    private static List<String> testData() {
        String[] testData = """
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..""".split("\n");
        return Arrays.asList(testData);
    }

}

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
