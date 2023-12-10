clear

if [[ $# -eq 0 ]]; then
	n=10
else
	n=$1
fi

python3 day$n/main.py
