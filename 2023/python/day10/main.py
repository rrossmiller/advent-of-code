import p1
import p2


def get_data(test=False) -> tuple[list[str], list[int]]:
    start = []
    with open("../data/10.txt") as f:
        rows = f.readlines()

    if test:
        rows = str(  # this is dumb, pyright. Why?
            """-L|F7
7S-7|
L|7||
-L-J|
L|-JF"""
        ).split("\n")
        #         rows = str(
        #             """.....
        # .S-7.
        # .|.|.
        # .L-J.
        # ....."""
        #         ).split("\n")

        rows = str(
            """..F7.
.FJ|.
SJ.L7
|F--J
LJ..."""
        ).split("\n")
        rows = str(
            """7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ"""
        ).split("\n")

        # pt2
        rows = str(
            """...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
..........."""
        ).splitlines()
        rows = str(
            """.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ..."""
        ).splitlines()

    for i, l in enumerate(rows):
        for j, c in enumerate(l):
            if c == "S":
                return rows, [i, j]

    return rows, start


if __name__ == "__main__":
    """
    How many steps along the loop does it take to get from the starting position to the point farthest from the starting position?

    | is a vertical pipe connecting north and south.
    - is a horizontal pipe connecting east and west.
    L is a 90-degree bend connecting north and east.
    J is a 90-degree bend connecting north and west.
    7 is a 90-degree bend connecting south and west.
    F is a 90-degree bend connecting south and east.
    . is ground; there is no pipe in this tile.
    S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
    """
    print("****")
    testing = True
    data, start = get_data(testing)
    # print(f"{start= }")
    # print(data)
    # print()
    # dist = p1.run(start, data)
    # print(f"p1: {dist}")

    enclosed = p2.run(start, data)
    print(f"p2: {enclosed}")
