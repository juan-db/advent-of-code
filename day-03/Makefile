# This is a terrible Makefile but it's good enough
# I really need to use implicit rules for these

part-one.jar:
	kotlinc -include-runtime -d part-one.jar part-one.kt

part-two.jar:
	kotlinc -include-runtime -d part-two.jar part-two.kt

fclean:
	rm -f part-one.jar part-two.jar
