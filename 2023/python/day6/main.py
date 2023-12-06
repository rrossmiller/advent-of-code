import p1


def get_data() -> list[tuple[int, int]]:
    with open("../data/6.txt") as f:
        lines = [s.replace("\n", "") for s in f.readlines()]
    # test data
    # lines = ["Time:      7  15   30", "Distance:  9  40  200"]

    time = [int(s) for s in lines[0].split(":")[1].strip().split(" ") if s != ""]
    dist = [int(s) for s in lines[1].split(":")[1].strip().split(" ") if s != ""]
    data = [(a, b) for a, b in zip(time, dist)]

    return data


if __name__ == "__main__":
    print("****")
    data = get_data()
    p1.run(data)
    t = ""
    d = ""
    for n, x in data:
        t += str(n)
        d += str(x)

    # pt2
    p1.run([(int(t), int(d))])
