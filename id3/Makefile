#all: id3.go id3.cpp
	#$(MAKE) id3_go
	#$(MAKE) id3_cpp

#id3_cpp: id3.cpp
	#g++ $^ -o $@

id3_go: id3.go
	go build -o $@ $^
	
go_run: id3.go
	go run $^

#python_run: id3.py
