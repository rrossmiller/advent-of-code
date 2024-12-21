XMAS = "XMAS"
MAS = "MAS"
visited = set()

search_dirs = [
    (-1, -1),  # UL
    (-1, 0),  # U
    (-1, 1),  # UR
    (0, -1),  # L
    (0, 1),  # R
    (1, -1),  # DL
    (1, 0),  # D
    (1, 1),  # DR
]
x_dirs = [
    (-1, -1),  # UL
    (-1, 1),  # UR
    (1, -1),  # DL
    (1, 1),  # DR
]


def run(data: list[str]):
    p1(data)

    visited.clear()
    if len(data) < 100:
        for row in data:
            print(" ".join(c for c in row))
        print()
    p2(data)
    if len(data) < 100:
        print_visited(data)


def p1(data: list[str]):
    words = []
    for i, row in enumerate(data):
        for j, c in enumerate(row):
            if c == XMAS[0]:
                # print(i, j, c)
                visited.add((i, j))
                words.extend([search(data, i, j, d) for d in search_dirs])

    words = list(filter(lambda x: x == XMAS, words))
    # print(words)
    print("P1:", len(words))


def p2(data: list[str]):
    words = 0
    for i, row in enumerate(data):
        for j, c in enumerate(row):
            if c == "A":
                diag = False
                anti_diag = False
                if (
                    i - 1 >= 0
                    and j - 1 >= 0
                    and i + 1 < len(data)
                    and j + 1 < len(data[0])
                ):
                    # search diag
                    w = f"{data[i-1][j-1]}{data[i][j]}{data[i+1][j+1]}"
                    if w == MAS or w == MAS[::-1]:
                        diag = True
                    # search anti-diag
                    w = f"{data[i-1][j+1]}{data[i][j]}{data[i+1][j-1]}"
                    if w == MAS or w == MAS[::-1]:
                        anti_diag = True

                if anti_diag and diag:
                    words += 1
                    visited.add((i, j))
                    visited.add((i - 1, j - 1))
                    visited.add((i + 1, j + 1))
                    visited.add((i - 1, j + 1))
                    visited.add((i + 1, j - 1))
    print("P2:", words)


def search(
    data: list[str],
    i: int,
    j: int,
    dir: tuple[int, int],
    step=1,
    word="X",
    search_word=XMAS,
):
    # return the word if we're searching too far
    if step == len(search_word):
        return word

    letter_to_find = search_word[step]

    # if the next step is within the bounds of the grid
    x = i + dir[0]
    y = j + dir[1]
    if x >= 0 and x < len(data) and y >= 0 and y < len(data[0]):
        visited.add((x, y))
        letter = data[x][y]  # current letter
        if letter == letter_to_find:
            word += letter
            return search(data, x, y, dir, step + 1, word=word, search_word=search_word)

    return word


def print_visited(data):
    print(visited)
    print()
    for i in range(len(data)):
        r = ""
        for j in range(len(data[0])):
            if (i, j) in visited:
                r += data[i][j]
            else:
                r += "."
        print(r)
    print()
