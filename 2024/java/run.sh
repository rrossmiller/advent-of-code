#!/bin/bash
# if [[ $1 == 'r' ]]; then
#   clear
#   java -jar ./target/aoc-0.0.1.jar
#   exit $?
# elif [[ $1 == 'c' ]]; then
if [[ $1 == 'c' ]]; then
  mvn clean package || exit 1
fi
clear
java -jar ./target/aoc-0.0.1.jar
