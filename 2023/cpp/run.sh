clear
if [[ -e aoc ]]; then
    rm aoc
fi

if [[ $# -eq 0 ]]; then
    day=1 
else
    day=$1
fi
clang++ -std=c++23  day$day.cpp  -o aoc &&
    ./aoc
