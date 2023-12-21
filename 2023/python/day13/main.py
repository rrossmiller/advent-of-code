import p1
import p2


def get_data(test=False):
    with open("../data/13.txt") as f:
        rows = f.read().splitlines()

    if test:
        rows = """\
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#""".splitlines()

    data = []
    st = 0
    for i, r in enumerate(rows):
        r = r.strip()
        if len(r) == 0:
            data.append(rows[st:i])
            st = i + 1

    # add the last pattern
    data.append(rows[st:])

    return data


if __name__ == "__main__":
    print("****\nday13")
    test = False
    data = get_data(test)
    sum = p1.pt1(data)
    print()
    print(f"pt1: {sum}")
    print()
    sum = p2.pt2(data)
    if sum >= 39929:
        print("too high")

    print(f"pt2: {sum}")
