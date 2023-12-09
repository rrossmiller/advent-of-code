package day8;

import java.util.List;

import day8.App.Direction;

class Part2 {

    // try 3... Lowest Common Multiple
    public static void main(TreeNode[] nodes, List<Direction> dirs) {
        int[] stepsArr = new int[nodes.length];

        for (int i = 0; i < nodes.length; i++) {
            int steps = 0;
            int idx = 0;
            while (!nodes[i].id.endsWith("Z")) {
                steps++;
                // wrap the index
                if (idx == dirs.size()) {
                    idx = 0;
                }

                var dir = dirs.get(idx);
                // one step for each node
                switch (dir) {
                    case L:
                        nodes[i] = nodes[i].left;
                        break;
                    case R:
                        nodes[i] = nodes[i].right;
                        break;
                }
                idx++;
            }
            stepsArr[i] = steps;
        }
        System.out.println(String.format("Part 2: %d", lcm(stepsArr)));
    }

    private static long lcm(int[] steps) {
        long result = steps[0];
        for (int i = 1; i < steps.length; i++) {
            result = lcm((long) result, (long) steps[i]);
        }

        return result;
    }

    private static long lcm(long a, long b) {
        return (a * b) / gcd(a, b);
    }

    private static long gcd(long a, long b) {
        if (a == 0 || b == 0) {
            return a + b;
        }
        long biggerValue = Math.max(Math.abs(a), Math.abs(b));
        long smallerValue = Math.min(Math.abs(a), Math.abs(b));
        return gcd(biggerValue % smallerValue, smallerValue);
    }

    // // try 1 -- don't brute force this...
    // public static void main(TreeNode[] nodes, List<Direction> dirs) {
    // int steps = 0;
    // int idx = 0;
    // var start = 0.0+System.currentTimeMillis();
    // // follow the directions until node ZZZ is reached
    // while (!allZs(nodes)) {
    // steps++;
    // // wrap the index
    // if (idx == dirs.size()) {
    // idx = 0;
    // }
    //
    // var dir = dirs.get(idx);
    // // one step for each node
    // for (int i = 0; i < nodes.length; i++) {
    // // System.out.println(i + " id: " + nodes[i]);
    // switch (dir) {
    // case L:
    // nodes[i] = nodes[i].left;
    // break;
    // case R:
    // nodes[i] = nodes[i].right;
    // break;
    // }
    // }
    // idx++;
    // // long lim = 100_000_000;
    // // if (steps > lim) {
    // // System.out.println(String.format("nope.... lim: %s",
    // // NumberFormat.getInstance(Locale.US).format(lim)));
    // // return;
    // // }
    // if ((System.currentTimeMillis() - start) / 1000.0 / 60 > 15) {// mins
    // System.out.println("15 mins exceeded");
    // return;
    // }
    // }
    // System.out.println(String.format("Part 2: %d", steps));
    //
    // }
    //
    // private static boolean allZs(TreeNode[] nodes) {
    // // for (TreeNode n : nodes) {
    // // if (!n.id.endsWith("Z"))
    // // return false;
    // // }
    // // return true;
    // return Arrays.stream(nodes).allMatch(n -> n.id.endsWith("Z"));
    // }
    //
    // // try 2
    // public static void main(Map<String, String[]> nodeMap, List<Direction> dirs)
    // {
    // int steps = 0;
    // int idx = 0;
    // int nodesSize = nodeMap.keySet().stream().filter(k ->
    // k.endsWith("A")).toArray().length;
    // String[] nodes = new String[nodesSize];
    // nodes = nodeMap.keySet().stream().filter(s ->
    // s.endsWith("A")).collect(Collectors.toList()).toArray(nodes);
    // // System.out.println();
    // // for (String n : nodes)
    // // System.out.println("n:" + n);
    // // follow the dirs until all nodes end in Z
    // while (!allZs(nodes)) {
    // steps++;
    // // wrap the index
    // if (idx == dirs.size()) {
    // idx = 0;
    // }
    // // one step for each node
    // for (int i = 0; i < nodes.length; i++) {
    // switch (dirs.get(idx)) {
    // case L:
    // nodes[i] = nodeMap.get(nodes[i])[0];
    // break;
    // case R:
    // nodes[i] = nodeMap.get(nodes[i])[1];
    // break;
    // }
    // }
    // idx++;
    // long lim = 100_000_000;
    // if (steps > lim) {
    // System.out.println(String.format("nope.... lim: %s",
    // NumberFormat.getInstance(Locale.US).format(lim)));
    // return;
    // }
    // }
    //
    // System.out.println(String.format("Part 2: %d", steps));
    //
    // }
    //
    // private static boolean allZs(String[] nodes) {
    // return Arrays.stream(nodes).allMatch(n -> n.endsWith("Z"));
    // // return Arrays.stream(nodes).anyMatch(n -> n.endsWith("Z"));
    // }
    //
}
