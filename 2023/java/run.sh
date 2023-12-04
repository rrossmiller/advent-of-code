clear

if [[ $# -eq 0 ]]; then
	n=3
else
	n=$1
fi

javac day$n/*.java
java day$n/*.java
