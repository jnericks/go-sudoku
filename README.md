# Sudoku Solver

## Build & Run

Run `make all`, which will produce a runnable `./solve` binary

Run `./solve /path/to/sudoku.csv` or `./solve < /path/to/sudoku.csv`

You should see the something similar to the following:

```
$ ./solve ~/path/to/sudoku.csv
Enter Sudoku:
-,-,3,-,2,-,6,-,-
9,-,-,3,-,5,-,-,1
-,-,1,8,-,6,4,-,-
-,-,8,1,-,2,9,-,-
7,-,-,-,-,-,-,-,8
-,-,6,7,-,8,2,-,-
-,-,2,6,-,9,5,-,-
8,-,-,2,-,3,-,-,9
-,-,5,-,1,-,3,-,-
Solution:
4,8,3,9,2,1,6,5,7
9,6,7,3,4,5,8,2,1
2,5,1,8,7,6,4,9,3
5,4,8,1,3,2,9,7,6
7,2,9,5,6,4,1,3,8
1,3,6,7,9,8,2,4,5
3,7,2,6,8,9,5,1,4
8,1,4,2,5,3,7,6,9
6,9,5,4,1,7,3,8,2
```