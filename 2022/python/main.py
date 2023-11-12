import sys

from day11 import run

if __name__ == "__main__":
    day = sys.argv[1] if len(sys.argv) > 1 else "11"
    match day:
        case "11":
            run()
