#!/bin/bash

if [[ ! $(go run . "line" ../testFiles/* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'line'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'line' "
fi

if [[ ! $(go run . "world" ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'world'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'world' "
fi

if [[ ! $(go run . -quality 0 "xyz" ../testFiles/* | wc -l) -eq 17 ]];then
    echo "FAILED: Pattern matching failed where it should match every line"
    exit 1
else
    echo "PASS: Pattern matching with quality 0 "
fi

if [[ ! $(go run . -quality 100 "xyz" ../testFiles/* | wc -l) -eq 3 ]];then
    echo "FAILED: Pattern matching failed where it should match nothing and print only the file paths"
    exit 1
else
    echo "PASS: Pattern matching for forced no result "
fi

if [[ ! $(go run . -quality 100 "world0" ../testFiles/* | grep world | wc -l) -eq 1 ]];then
    echo "FAILED: Pattern matching failed where it should match only one line"
    exit 1
else
    echo "PASS: Pattern matching for forced one line result "
fi

if [[ ! $(cat ../testFiles/testFile1.txt | go run . line | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Reading from StdIn through pipe failed"
    exit 1
else
    echo "PASS: Pattern matching for input from StdIn "
fi

if [[ ! $(go run . -quality 75 wold ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Deletion 'world' -> 'wold'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wold'"
fi 

if [[ ! $(go run . -quality 75 wod ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Deletion 'world' -> 'wod'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wod'"
fi

if [[ ! $(go run . -quality 75 lixxne ../testFiles/testFile* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxne'"
fi

if [[ ! $(go run . -quality 75 lixxxne ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxxne'"
fi

if [[ ! $(go run . -recursive=../testFiles FileContent | grep FileContent | wc -l) -eq 3 ]];then
    echo "FAILED: Only recursevly searching in not hidden files"
    exit 1
else
    echo "PASS: Only recursevly searching in not hidden files"
fi

if [[ ! $(go run . -recursive=../testFiles -recH FileContent | grep FileContent | wc -l) -eq 6 ]];then
    echo "FAILED: Recursevly searching also in hidden files"
    exit 1
else
    echo "PASS: Recursevly searching also in hidden files"
fi

if [[ ! $(go run . -recursive=../testFiles -type=n test | wc -l) -eq 3 ]];then
    echo "FAILED: Recursive file name search"
    exit 1
else
    echo "PASS: Recursive file name search"
fi

if [[ ! $(go run . -recursive="../testFiles" -type="n" -recH test | wc -l) -eq 6 ]];then
    echo "FAILED: Recursive file name search including hidden files"
    exit 1
else
    echo "PASS: Recursive file name search including hidden files"
fi

if [[ ! $(go run . -recursive="../testFiles" -type="n" file | grep -v bin | wc -l) -eq 3 ]];then
    echo "FAILED: Recursive file name search excluding binary files"
    exit 1
else
    echo "PASS: Recursive file name search excluding binary files"
fi

if [[ ! $(go run . -recursive="../testFiles" -type="n" -binary file | grep bin | wc -l) -eq 1 ]];then
    echo "FAILED: Recursive file name search including binary files"
    exit 1
else
    echo "PASS: Recursive file name search including binary files"
fi
