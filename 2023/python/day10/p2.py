# map heading direction to a new direction and how to update the point
dirs = {
    "U": [-1, 0],  # up
    "D": [1, 0],  # down
    "L": [0, -1],  # left
    "R": [0, 1],  # right
}

# pipe type: {heading coming in: (heading going out, transform on coordinate}
pipe_dirs = {
    "|": {"U": ("U", dirs["U"]), "D": ("D", dirs["D"])},  # continue in the same dir
    "-": {"R": ("R", dirs["R"]), "L": ("L", dirs["L"])},  # continue in the same dir
    "L": {
        "D": ("R", dirs["R"]),
        "L": ("U", dirs["U"]),
    },  # heading down, go right. Left, go up
    "J": {
        "R": ("U", dirs["U"]),
        "D": ("L", dirs["L"]),
    },  # heading right, go up. Down, go left
    "7": {
        "U": ("L", dirs["L"]),
        "R": ("D", dirs["D"]),
    },  # heading up, go left. Right down
    "F": {
        "U": ("R", dirs["R"]),
        "L": ("D", dirs["D"]),
    },  # heading up, go right. Left, down
}


def run(start: list[int], data: list[str]):
    heading = ""
    # dist, heading, coordinate
    q: list[tuple[int, str, list[int]]] = []
    # init q
    for h, d in dirs.items():
        s = [start[0] + d[0], start[1] + d[1]]
        # if the next 'piece' isn't '.'
        piece = data[s[0]][s[1]]
        if next_piece_valid(piece, h):
            q.append((1, h, s))

    # print(f"{q= }")

    visited: dict[str, int] = {}
    while len(q) > 0:
        # pop q
        t, q = q[0], q[1:]
        (
            dist,
            heading,
            coord,
        ) = t  # distance of this piece, the heading T, the current coordinate
        k = f"{coord[0]}:{coord[1]}"

        # if the piece hasn't already been visited
        if k not in visited:
            # print(f"{heading}-> {k}")
            visited[k] = dist
            cur_piece = data[coord[0]][coord[1]]

            # step through the piece and add result to the q
            # add its neighbors to the queue
            heading, move = pipe_dirs[cur_piece][heading]
            coord[0] += move[0]
            coord[1] += move[1]
            next_piece = data[coord[0]][coord[1]]
            if next_piece_valid(next_piece, heading):
                q.append((dist + 1, heading, coord))

    # write check file
    out = ""
    to_key = lambda x: f"{x[0]}:{x[1]}"
    for i, l in enumerate(data):
        for j, c in enumerate(l):
            k = to_key((i, j))
            if k in visited or c == "S":
                out += c
            else:
                out += "*"
        out += "\n"
    with open("out.txt", "w") as f:
        f.write(out)
    # return flood_fill(data, set(visited.keys()))
    return pip(data, set(visited.keys()))


def next_piece_valid(next_piece: str, heading: str) -> bool:
    match heading:
        case "U":
            return next_piece in ["|", "7", "F"]
        case "D":
            return next_piece in ["|", "L", "J"]
        case "L":
            return next_piece in ["-", "F", "L"]
        case "R":
            return next_piece in ["-", "J", "7"]

    return False


# point in polygon
def pip(data: list[str], loop: set[str]):
    rows, cols = len(data), len(data[0])
    cnt = 0
    to_key = lambda x: f"{x[0]}:{x[1]}"
    out = ""
    # count vertical loop edges along each row (anything that isn't '-')
    for i, l in enumerate(data):
        for j, c in enumerate(l):
            k = to_key((i, j))
            # if a point is not a part of the loop and the edge count is odd, it's inside
            if k not in loop:
                # get edge count
                edges = 0
                lx, rx = j, j
                uy, dy = i, i
                while lx > 0 and rx < cols and uy > 0 and dy < rows:
                    # horiz
                    if lx > 0:
                        lx -= 1
                    if rx < cols - 1:
                        rx += 1
                    # if to_key((i, lx)) in loop and data[i][lx] != "-":
                    if data[i][lx] == "|":
                        edges += 1
                    if data[i][rx] == "|":
                        edges += 1

                    # vert
                    # while uy >= 0 and dy <= rows:
                    if uy > 0:
                        uy -= 1
                    if dy < rows - 1:
                        dy += 1
                    if data[uy][j] != "-":
                        edges += 1
                    if data[dy][j] != "-":
                        edges += 1
                if edges % 2 != 0 and edges > 0:
                    cnt += 1
                    out += "*"
                else:
                    out += " "
            else:
                out += c
        #     break
        # break

        out += "\n"
    with open("fill.txt", "w") as f:
        f.write(out)
    return cnt


# doesn't cover the case where two loop pieces touch each other but are not connected
# def flood_fill(data: list[str], loop: set[str]):
#     rows, cols = len(data), len(data[0])
#     ttl = rows * cols
#     to_key = lambda x: f"{x[0]}:{x[1]}"
#     q: list[tuple[int, int]] = [
#         (0, 0),  # top left
#         (0, cols - 1),  # top right
#         (rows - 1, cols - 1),  # bottom right
#         (rows - 1, 0),  # bottom left
#     ]
#     visited = set()
#
#     # from wikipedia: https://en.wikipedia.org/wiki/Flood_fill
#     # slightly different because this is the reverse... fill what isn't in the shape
#     while len(q) > 0:
#         # pop queue
#         n, q = q[0], q[1:]
#         n_key = to_key(n)
#         if n_key not in visited:  # not in loop and char != "S":
#             # visit the location
#             visited.add(n_key)
#             # add the valid neighbors
#             for _, d in dirs.items():
#                 x, y = n[0] + d[0], n[1] + d[1]
#                 # if s is within the boundary of the grid and hasn't already been filled (visted)
#                 k = to_key((x, y))
#                 if (
#                     x >= 0
#                     and x < rows
#                     and y >= 0
#                     and y < cols
#                     and k not in loop
#                     and data[x][y] != "S"
#                     and k not in visited
#                 ):
#                     q.append((x, y))
#     # write check file
#     out = ""
#     abc = 0
#     for i, l in enumerate(data):
#         for j, c in enumerate(l):
#             k = to_key((i, j))
#             if k in loop or c == "S":
#                 out += c
#             elif k in visited:
#                 # out += "O"
#                 out += " "
#             else:
#                 # check that eventually, in all directions, this point will hit the loop
#                 ok = True
#                 for _, d in dirs.items():
#                     x = i
#                     y = j
#                     while ok:
#                         x += d[0]
#                         y += d[1]
#                         k = to_key((x, y))
#                         # if it's a border, it's not inside
#                         if x == 0 or x >= rows - 1 or y == 0 or y >= cols - 1:
#                             ok = False
#                         elif k in loop:
#                             break
#
#                 if ok:
#                     abc += 1
#                     # out += "I"
#                     # out += "ø"
#                     # out += "ˆ"
#                     # out += "◊"
#                     out += "*"
#                 else:
#                     out += " "
#         out += "\n"
#     with open("fill.txt", "w") as f:
#         f.write(out)
#
#     from collections import Counter
#
#     print(Counter(out)["*"])
#     print(abc)
#     return ttl - len(visited) - len(loop)
