#BJCLI

### Attempt to create a program to spit out best action for blackjack hand based on the following table:

```
P = Split
H = Hit
D = Double
S = Stand
```

```
┌─────────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐
│ Dealer> │  2  │  3  │  4  │  5  │  6  │  7  │  8  │  9  │ 10  |  A  |
└─────────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘
┌─────────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐
| Player  │     │     │     │     │     │     │     │     │     |     |
├─────────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤
│   1 1   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │
│   1 2   │ 'H' │ 'H' │ 'H' │ 'H' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   1 3   │ 'H' │ 'H' │ 'H' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   1 4   │ 'H' │ 'H' │ 'H' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   1 5   │ 'H' │ 'H' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   1 6   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   1 7   │ 'S' │ 'D' │ 'D' │ 'D' │ 'D' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │
│   1 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   1 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  1 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   2 1   │ 'H' │ 'H' │ 'H' │ 'H' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 2   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 3   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 4   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 5   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 6   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 7   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   2 8   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   2 9   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│  2 10   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 1   │ 'H' │ 'H' │ 'H' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 2   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 3   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 4   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 5   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 6   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   3 7   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   3 8   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   3 9   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  3 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 1   │ 'H' │ 'H' │ 'H' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 2   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 3   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 4   │ 'H' │ 'H' │ 'H' │ 'P' │ 'P' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 5   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 6   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   4 7   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   4 8   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   4 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  4 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 1   │ 'H' │ 'H' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 2   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 3   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 4   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 5   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   5 6   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   5 7   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   5 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  5 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 1   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 2   │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 3   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 4   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   6 5   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   6 6   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 7   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   6 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  6 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 1   │ 'S' │ 'D' │ 'D' │ 'D' │ 'D' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │
│   7 2   │ 'H' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 3   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   7 4   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   7 5   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 6   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 7   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   7 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  7 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   8 1   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   8 2   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │ 'H' │
│   8 3   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   8 4   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   8 5   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   8 6   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   8 7   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   8 8   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │
│   8 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  8 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   9 1   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   9 2   │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'D' │ 'H' │
│   9 3   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   9 4   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   9 5   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   9 6   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   9 7   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│   9 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│   9 9   │ 'P' │ 'P' │ 'P' │ 'P' │ 'P' │ 'S' │ 'P' │ 'P' │ 'S' │ 'S' │
│  9 10   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  10 1   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  10 2   │ 'H' │ 'H' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  10 3   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  10 4   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  10 5   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  10 6   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'H' │ 'H' │ 'H' │ 'H' │ 'H' │
│  10 7   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  10 8   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  10 9   │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
│  10 10  │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │ 'S' │
└─────────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘
```