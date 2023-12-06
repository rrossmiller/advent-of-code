def run(data: list[tuple[int, int]]):
    """
    Determine how long to hold down the button to surpass distance d if you have t seconds and
    for each second you hold the button, the spead increases by 1m/s

    Ex:
        7 second race, 9 meters to beat
        x: Don't hold the button at all at the start of the race, the boat won't move; it will have traveled 0 millimeters by the end of the race.
        x: Hold the button for 1 millisecond at the start of the race. Then, the boat will travel at a speed of 1 meter per second for 6 seconds
        √: Hold for 2 seconds: 2m/s for 5 seconds = 10m > 9m
        √: Hold for 3 seconds: 3m/s for 4 seconds = 12m > 9m
        √: Hold for 4 seconds: 4m/s for 3 seconds = 12m > 9m
        √: Hold for 5 seconds: 5m/s for 2 seconds = 10m > 9m
        x: Hold for 6 seconds: 6m/s for 1 seconds = 6m < 9m
        x: Hold for 7 seconds: 7m/s for 0 seconds = 0m < 9m

    Solve by starting in the middle and working back until holding for that many seconds won't beat the distance
    """

    rounds = []
    for t, d in data:
        wins = 0
        winning = True
        hold_down = int(t / 2)
        minus_1 = (t - hold_down) == hold_down
        while winning:
            time_left = t - hold_down
            # print(f"{wins+1}:", hold_down, time_left, time_left * hold_down, ">?", d)
            if hold_down * time_left > d:
                wins += 1
            else:
                winning = False
            hold_down = hold_down - 1

        # -1 if odd wins to not double count middle winning hold down time
        wins = wins * 2 - 1 if minus_1 else wins * 2
        rounds.append(wins)
        # print(f"->{t} {d}: {wins}")
        # print()

    # print()
    # print(rounds)
    ans = 1
    for r in rounds:
        ans *= r

    print(ans)
