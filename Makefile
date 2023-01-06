.DEFAULT_GOAL:= all

all: 1 2 3 4 5 6 7 8 9 10 11 12 13 14

.PHONY: 1
1:
	cd one && go test -v

.PHONY: 2
2:
	cd two && go test -v

.PHONY: 3
3:
	cd three && go test -v

.PHONY: 4
4:
	cd four && go test -v

.PHONY: 5
5:
	cd five && go test -v

.PHONY: 6
6:
	cd six && go test -v

.PHONY: 7
7:
	cd seven && go test -v

.PHONY: 8
8:
	cd eight && go test -v

.PHONY: 9
9:
	cd nine && go test -v

.PHONY: 10
10:
	cd ten && go test -v

.PHONY: 11
11:
	cd eleven && go test -v

.PHONY: 12
12:
	cd twelve && go test -v

.PHONY: 13
13:
	cd thirteen && go test -v

.PHONY: 14
14:
	cd fourteen && go test -v

.PHONY: fourteen
fourteen:
	cd fourteen && go run .