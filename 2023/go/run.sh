clear
if [[ -e aoc ]]; then
    rm aoc
fi

if [[ $# -eq 0 ]]; then
    n=2 
else
    n=$1
fi

cd day$n &&
    go build -o aoc &&
    mv aoc .. && 
    cd ..&&
    ./aoc

