package day8;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

class App {
    enum Test {
        Part1, Part2
    }

    enum Direction {
        L, R
    }

    public static void main(String[] args) throws Exception {
        System.out.println("*****");
        List<String> data = readData();
        // data = testData(Test.Part1);
        // data = testData(Test.Part2);
        String dirString = data.remove(0);
        List<Direction> dirs = prepData(dirString);
        TreeNode root = prepData(data);
        Part1.main(root, dirs);

    }

    private static List<Direction> prepData(String data) {
        List<Direction> dirs = new ArrayList<>();
        for (String d : data.split("")) {
            switch (d) {
                case "L":
                    dirs.add(Direction.L);
                    break;
                case "R":
                    dirs.add(Direction.R);
                    break;
            }

        }
        return dirs;
    }

    private static TreeNode prepData(List<String> data) {
        Map<String, TreeNode> treeMap = new HashMap<>();
        TreeNode node;
        for (String s : data) {
            var spl1 = s.split(" = "); // id, LR
            var spl2 = spl1[1].split(","); // LR
            String lID = spl2[0].trim().replace("(", "");
            String rID = spl2[1].trim().replace(")", "");
            // if the id node already exists, update its neighbors

            if (treeMap.containsKey(spl1[0])) {
                node = treeMap.get(spl1[0]);
            } else {
                // else create the node
                node = new TreeNode();
                node.id = spl1[0];
                treeMap.put(spl1[0], node);
            }
            // left neighbor
            addNeighbor(node, lID, treeMap, Direction.L);
            // right neighbor
            addNeighbor(node, rID, treeMap, Direction.R);
            // System.out.println(node);
        }
        // System.out.println();
        // for (TreeNode v : treeMap.values()) {
        // System.out.println(v);
        // }
        // System.out.println();

        return treeMap.get("AAA");
    }

    private static void addNeighbor(TreeNode node, String neighborID, Map<String, TreeNode> treeMap, Direction d) {
        TreeNode neighb;
        if (treeMap.containsKey(neighborID)) {
            neighb = treeMap.get(neighborID);
        } else {
            neighb = new TreeNode();
            neighb.id = neighborID;
        }

        switch (d) {
            case L:
                node.left = neighb;
                break;
            case R:
                node.right = neighb;
                break;
        }
        treeMap.put(neighb.id, neighb);

    }

    private static List<String> readData() throws IOException {
        Path path = Paths.get("../../data/8.txt");
        // replace newline with space
        String s = Files.readString(path);
        List<String> data = Arrays.asList(s.split("\n"));
        data = data.stream().filter(l -> l.length() > 0).collect(Collectors.toList());
        return data;
    }

    private static List<String> testData(Test t) {
        String[] testData = null;
        switch (t) {
            case Part1:
                testData = """
                        RL

                        AAA = (BBB, CCC)
                        BBB = (DDD, EEE)
                        CCC = (ZZZ, GGG)
                        DDD = (DDD, DDD)
                        EEE = (EEE, EEE)
                        GGG = (GGG, GGG)
                        ZZZ = (ZZZ, ZZZ)
                        """.split("\n");
                break;
            case Part2:
                testData = """
                        LLR

                        AAA = (BBB, BBB)
                        BBB = (AAA, ZZZ)
                        ZZZ = (ZZZ, ZZZ)""".split("\n");
                break;

        }
        var rtn = Arrays.asList(testData);
        rtn = rtn.stream().filter(l -> l.length() > 0).collect(Collectors.toList());
        return rtn;
    }

}
