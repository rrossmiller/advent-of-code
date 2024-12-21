import argparse

import day4

from data import get_data

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-t", "--test", action="store_true")
    parser.add_argument("-d", "--day", default="4")

    args = parser.parse_args()

    data = get_data(f"{args.day}.txt", args.test)
    day4.run(data)
