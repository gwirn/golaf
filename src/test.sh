#!/bin/bash

if [[ ! $(go run main.go "line" ../testFiles/* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'line'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'line' "
fi

if [[ ! $(go run main.go "world" ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'world'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'world' "
fi

if [[ ! $(go run main.go -quality 0 "xyz" ../testFiles/* | wc -l) -eq 14 ]];then
    echo "FAILED: Pattern matching failed where it should match every line"
    exit 1
else
    echo "PASS: Pattern matching with quality 0 "
fi

if [[ ! $(go run main.go -quality 100 "xyz" ../testFiles/* | wc -l) -eq 2 ]];then
    echo "FAILED: Pattern matching failed where it should match nothing and print only the file paths"
    exit 1
else
    echo "PASS: Pattern matching for forced no result "
fi

if [[ ! $(go run main.go -quality 100 "world0" ../testFiles/* | grep world | wc -l) -eq 1 ]];then
    echo "FAILED: Pattern matching failed where it should match only one line"
    exit 1
else
    echo "PASS: Pattern matching for forced one line result "
fi

if [[ ! $(cat ../testFiles/testFile1.txt | go run main.go line | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Reading from StdIn through pipe failed"
    exit 1
else
    echo "PASS: Pattern matching for input from StdIn "
fi

if [[ ! $(go run main.go wold ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Reduced 'world' -> 'wold'"
    exit 1
else
    echo "PASS: Reduced 'world' -> 'wold'"
fi 

if [[ ! $(go run main.go wod ../testFiles/* | grep world | wc -l) -eq 0 && ! $(go run main.go wod ../testFiles/* | wc -l) -eq 0 ]];then
    echo "FAILED: Reduced 'world' -> 'wold'"
    exit 1
else
    echo "PASS: Reduced 'world' -> 'wod'"
fi
