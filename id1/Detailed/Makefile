all: id1_detailed.go id1_detailed.cpp
	$(MAKE) id1_detailed_go
	$(MAKE) id1_detailed_cpp

id1_detailed_cpp: id1_detailed.cpp
	g++ $^ -o $@

id1_detailed_go: id1_detailed.go
	go build -o $@ $^
	
go_run: id1_detailed.go
	go run $^

python_run: id1_detailed.py
	python3 $^
