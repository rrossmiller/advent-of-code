package day8;

class TreeNode {
    public String id;
    public TreeNode left;
    public TreeNode right;

    public String toString() {
        String l = this.left == null ? "null" : this.left.id;
        String r = this.right == null ? "null" : this.right.id;
        return String.format("%s: %s %s", this.id, l, r);
    }
}
