
.PHONY: grammar

grammar: Algol60V2.g4
	wget -N -P /tmp/ https://www.antlr.org/download/antlr-4.10.1-complete.jar
	java -jar /tmp/antlr-4.10.1-complete.jar -Dlanguage=Go -package algol60_parser -no-listener -visitor $<
