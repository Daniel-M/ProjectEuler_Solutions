#all: id6.go id6.cpp
	#$(MAKE) id6_go
	#$(MAKE) id6_cpp

#id6_cpp: id6.cpp
	#g++ $^ -o $@

id6_go: id6.go
	go build -o $@ $^
	
go_run: id6.go
	go run $^

#python_run: id6.py
