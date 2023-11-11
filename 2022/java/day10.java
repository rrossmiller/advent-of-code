import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class day10 {
    private static List<String> crt = new ArrayList<>();

    public static void main(String[] args) throws IOException {
        List<String> data = getData();
        // List<String> data = getTestData();
        int x = 1;
        int ans = 0;
        List<Integer> checks = Arrays.asList(20, 60, 100, 140, 180, 220); // 20th, 60th, 100th, 140th,
                                                                          // 180th, and 220th cycles

        crt.add("#");
        for (int i = 1; i < data.size(); i++) {
            int s = i + 1;
            if (checks.contains(s)) {
                ans += x * s;
            }

            System.out.println(i % 40 + "|" + x);
            if (x >= i % 40 - 1 && x <= i % 40 + 1) {
                crt.add("#");
            } else {
                crt.add(".");
            }
            if (data.get(i - 1).equals("addx")) {
                int n = Integer.parseInt(data.get(i));
                x += n;
            }
        }

        System.out.println("Ans:" + ans);
        if (ans >= 14141) {
            System.out.println("too high");
        } else if (ans <= 11020) {
            System.out.println("too low");
        }

        drawCRT();
    }

    private static List<String> getData() throws IOException {
        Path path = Paths.get("../data/day_10.txt");
        // replace newline with space
        String s = Files.readString(path).replace('\n', ' ');
        List<String> data = Arrays.asList(s.split(" "));
        return data;
    }

    private static List<String> getTestData() {
        String x = "addx 15 addx -11 addx 6 addx -3 addx 5 addx -1 addx -8 addx 13 addx 4 noop addx -1 addx 5 addx -1 addx 5 addx -1 addx 5 addx -1 addx 5 addx -1 addx -35 addx 1 addx 24 addx -19 addx 1 addx 16 addx -11 noop noop addx 21 addx -15 noop noop addx -3 addx 9 addx 1 addx -3 addx 8 addx 1 addx 5 noop noop noop noop noop addx -36 noop addx 1 addx 7 noop noop noop addx 2 addx 6 noop noop noop noop noop addx 1 noop noop addx 7 addx 1 noop addx -13 addx 13 addx 7 noop addx 1 addx -33 noop noop noop addx 2 noop noop noop addx 8 noop addx -1 addx 2 addx 1 noop addx 17 addx -9 addx 1 addx 1 addx -3 addx 11 noop noop addx 1 noop addx 1 noop noop addx -13 addx -19 addx 1 addx 3 addx 26 addx -30 addx 12 addx -1 addx 3 addx 1 noop noop noop addx -9 addx 18 addx 1 addx 2 noop noop addx 9 noop noop noop addx -1 addx 2 addx -37 addx 1 addx 3 noop addx 15 addx -21 addx 22 addx -6 addx 1 noop addx 2 addx 1 noop addx -10 noop noop addx 20 addx 1 addx 2 addx 2 addx -6 addx -11 noop noop noop";
        return Arrays.asList(x.split(" "));
    }

    private static void drawCRT() {
        for (int i = 0; i < crt.size(); i++) {
            if (i % 40 == 0) {
                System.out.println();
            }
            System.out.print(crt.get(i));
        }

    }
}
