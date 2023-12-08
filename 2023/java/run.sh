clear

if [[ $# -eq 0 ]]; then
	n=8
else
	n=$1
fi

# if [[ $n -eq 3 ]]; then
# javac day$n/*.java
# java day$n/*.java
# else
cd day$n
mvn clean package &&
	clear &&
	java -jar target/day$n-1.0.jar
# fi
