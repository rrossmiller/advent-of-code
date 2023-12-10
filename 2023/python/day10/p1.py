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

    # find max dist
    # import json

    # print(json.dumps(visited, indent=1))
    # print()
    return max(visited.values())


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
