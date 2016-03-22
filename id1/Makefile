all: id1.go id1.cpp
	$(MAKE) id1_go
	$(MAKE) id1_cpp

id1_cpp: id1.cpp
	g++ $^ -o $@

id1_go: id1.go
	go build -o $@ $^
	
go_run: id1.go
	go run $^

python_run: id1.py
	python3 $^
