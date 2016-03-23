#all: id4.go id4.cpp
	#$(MAKE) id4_go
	#$(MAKE) id4_cpp

#id4_cpp: id4.cpp
	#g++ $^ -o $@

id4_go: id4.go
	go build -o $@ $^
	
go_run: id4.go
	go run $^

#python_run: id4.py
