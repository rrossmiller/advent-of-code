package day8;

import java.util.List;

import day8.App.Direction;

class Part1 {
    public static void main(TreeNode node, List<Direction> dirs) {
        int steps = 0;
        int idx = 0;
        // follow the directions until node ZZZ is reached
        while (!node.id.equals("ZZZ")) {
            // System.out.println("id: "+node);
            steps++;
            // wrap the index
            if (idx == dirs.size()) {
                idx = 0;
            }

            switch (dirs.get(idx)) {
                case L:
                    node = node.left;
                    break;
                case R:
                    node = node.right;
                    break;
            }
            idx++;
        }
        System.out.println(String.format("Part 1: %d", steps));

    }
}
