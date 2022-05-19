antlr4_exe = java -jar /usr/local/lib/antlr-4.10.1-complete.jar
grun_exe = java org.antlr.v4.gui.TestRig
target = 

all:
	go build

gen:
	$(antlr4_exe) -Dlanguage=Go -o parser Calc.g4
	
gen_visitor:
	$(antlr4_exe) -Dlanguage=Go -visitor -o parser Calc.g4

gen_java:
	$(antlr4_exe) Calc.g4

gui: gen_java
	javac *.java
	$(grun_exe) Tdat prog $(target) -gui

trace: gen_java
	javac *.java
	$(grun_exe) Tdat prog $(target) -trace

tokens: gen_java
	javac *.java
	$(grun_exe) Tdat prog $(target) -tokens

tree: gen_java
	javac *.java
	$(grun_exe) Tdat prog $(target) -tree

clean:
	rm -fr *.java *.class calc


