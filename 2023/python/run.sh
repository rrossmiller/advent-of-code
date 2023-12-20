clear

if [[ $# -eq 0 ]]; then
	n=13
else
	n=$1
fi

python3 day$n/main.py
