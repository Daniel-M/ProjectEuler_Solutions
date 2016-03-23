#all: id5.go id5.cpp
	#$(MAKE) id5_go
	#$(MAKE) id5_cpp

#id5_cpp: id5.cpp
	#g++ $^ -o $@

id5_go: id5.go
	go build -o $@ $^
	
go_run: id5.go
	go run $^

#python_run: id5.py
