#all: id7.go id7.cpp
	#$(MAKE) id7_go
	#$(MAKE) id7_cpp

#id7_cpp: id7.cpp
	#g++ $^ -o $@

id7_go: id7.go
	go build -o $@ $^
	
go_run: id7.go
	go run $^

#python_run: id7.py
