#all: id8.go id8.cpp
	#$(MAKE) id8_go
	#$(MAKE) id8_cpp

#id8_cpp: id8.cpp
	#g++ $^ -o $@

id8_go: id8.go
	go build -o $@ $^
	
go_run: id8.go
	go run $^

#python_run: id8.py
