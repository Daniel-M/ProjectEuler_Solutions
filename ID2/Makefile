all: id2.go id2.cpp
	$(MAKE) id2_go
	$(MAKE) id2_cpp

id2_cpp: id2.cpp
	g++ $^ -o $@

id2_go: id2.go
	go build -o $@ $^
	
go_run: id2.go
	go run $^

python_run: id2.py
	python3 $^
