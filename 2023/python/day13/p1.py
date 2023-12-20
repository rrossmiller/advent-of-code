def pt1(data: list[list[str]]):
    """
    find axis of symmetry and augment with score transform
        ><
    #.##..##. 1
    ..#.##.#. 2
    ##......# 3
    ##......# 4
    ..#.##.#. 5
    ..##..##. 6
    #.#.##.#. 7
        ><

    #...##..#
    #....#..#
    ..##..###
    #####.##. v
    #####.##. ^
    ..##..###
    #....#..#

    """

    sum = 0
    for d in data:
        # print("-*-*-*-*-*-*-*-*-")
        # print_d(d, 0)
        # check horizontal symmetry
        idx = check_symmetry(d)
        if idx is not None:
            # print("horiz", idx)
            sum += 100 * idx

        else:
            # print("vert")
            # transpose
            d = transpose(d)
            idx = check_symmetry(d)
            if idx is not None:
                sum += idx
            else:
                print_d(d, 0)
                raise Exception("should be symmetric")

    return sum


def check_symmetry(d):
    # for every row (from 1 to len(d)-1)
    # the pivot is in between rows. So, pivot=1 means the pivot is between lines 0 and 1 -- b needs to lag by 1
    for pivot in range(1, len(d)):
        is_sym = True
        # check if all of the rows above and below the pivot row are identical
        t = pivot - 1  # top
        b = pivot  # bottom
        while t >= 0 and b < len(d):
            # print(t, b, len(d))
            # print_d(d[t : b + 1], pivot)
            if d[t] != d[b]:
                is_sym = False
                break
            t -= 1
            b += 1

        if is_sym:
            # print(t + 1, b - 1, "is symm")
            return pivot
    return None


def transpose(d):
    t = [[] for _ in range(len(d[0]))]
    for j in range(len(d[0])):
        for i in range(len(d) - 1, -1, -1):
            t[j].append(d[i][j])
    return t


# check vertical symmetry
def print_d(d, p):
    print(p)
    print("-->")
    for i, r in enumerate(d):
        # if i == p:
        #     print(f"--\n  {r}")
        # else:
        # print("  " + r)
        print(r)

    print("-->")
    print()


