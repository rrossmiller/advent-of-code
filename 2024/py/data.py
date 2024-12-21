def get_data(day: str, test: bool) -> list[str]:
    with open(f"../data/{day}") as f:
        lines = f.read().splitlines()

    test_line = 0
    for i in range(len(lines)):
        if lines[i] == "-----TEST-----":
            test_line = i
            break

    if test:
        return lines[:test_line]

    return lines[test_line + 1 :]
