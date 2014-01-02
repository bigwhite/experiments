

all:
	javac -d build/class/lib1 src/lib1/com/helloworld/lib1/Bar1.java
	jar cf lib/Lib1.jar -C build/class/lib1 com/helloworld/lib1/Bar1.class
	javac -d build/class/lib2 src/lib2/com/helloworld/lib2/Bar2.java
	jar cf lib/Lib2.jar -C build/class/lib2 com/helloworld/lib2/Bar2.class
	javac -d build/class/main -cp 'src/lib1:src/lib2' src/main/HelloWorld.java 
	jar cmf src/main/Manifest.txt bin/HelloWorld.jar -C build/class/main com/helloworld/main/HelloWorld.class

clean:
	rm -r build/class/lib1/*
	rm -r build/class/lib2/*
	rm -r build/class/main/*
	rm -r lib/*.jar
	rm -r bin/*.jar
