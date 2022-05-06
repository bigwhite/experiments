antlr4_exe = java -jar /usr/local/lib/antlr-4.10.1-complete.jar
grun_exe = java org.antlr.v4.gui.TestRig
target = demo1.csv

all: build

build:
	go build

gen:
	$(antlr4_exe) -Dlanguage=Go -o parser CSV.g4

gen_java:
	$(antlr4_exe) CSV.g4

gui: gen_java
	javac *.java
	$(grun_exe) CSV csvFile $(target) -gui

tokens: gen_java
	javac *.java
	$(grun_exe) CSV csvFile $(target) -tokens

trace: gen_java
	javac *.java
	$(grun_exe) CSV csvFile $(target) -trace

tree: gen_java
	javac *.java
	$(grun_exe) CSV csvFile $(target) -tree

diag: gen_java
	javac *.java
	$(grun_exe) CSV csvFile $(target) -diagnostics

test: gen
	go run main.go $(target)

clean:
	rm -fr ./csvparser
