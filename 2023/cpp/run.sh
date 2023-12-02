clear
if [[ -e aoc ]]; then
	rm aoc
fi

if [[ $# -eq 0 ]]; then
	n=1
else
	n=$1
fi

clang++ -std=c++23 day$n/*.cpp -o aoc &&
	./aoc
