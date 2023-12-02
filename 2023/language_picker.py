import argparse
import random as rand
from collections import Counter

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--testing", action="store_true")
    args = parser.parse_args()
    testing = args.testing

    languages = {
        "python": .25,
        "go": 1,
        "rust": 1,
        "java": 1,
        "js": 0.5,
        "ts": 0.5,
        "cpp": 0.1,
    }

    languages, weights = list(languages.keys()), list(languages.values())

    if testing:
        import matplotlib.pyplot as plt

        c = []
        for _ in range(10_000):
            l = rand.choices(languages, weights)
            c.extend(l)

        count = Counter(c)
        print(count['cpp']/count['python'])
        plt.bar(count.keys(), count.values())
        plt.show()
    else:
        l = rand.choices(languages, weights,k=1)[0]
        print(l)
