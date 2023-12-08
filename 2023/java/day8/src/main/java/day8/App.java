package day8;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

class App {
    enum Part {
        Part1,
        Part2
    }

    public static void main(String[] args) throws Exception {
        System.out.println("*****");
        List<String> data = readData();
        data = testData();
        System.out.println(data);

    }

    private static List<String> readData() throws IOException {
        Path path = Paths.get("../../data/8.txt");
        // replace newline with space
        String s = Files.readString(path);
        List<String> data = Arrays.asList(s.split("\n"));
        return data;
    }

    private static List<String> testData() {
        String[] testData = """
                RL

                AAA = (BBB, CCC)
                BBB = (DDD, EEE)
                CCC = (ZZZ, GGG)
                DDD = (DDD, DDD)
                EEE = (EEE, EEE)
                GGG = (GGG, GGG)
                ZZZ = (ZZZ, ZZZ)
                """.split("\n");
        var rtn = Arrays.asList(testData);
        rtn = rtn.stream().filter(l -> l.length() > 0).collect(Collectors.toList());
        for (String s : rtn) {
            System.out.println(s);
        }
        return rtn;
    }

}
