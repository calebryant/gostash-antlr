#!/bin/bash

export PATH=$PATH:/Users/caleb.bryant/Library/Python/3.11/bin
outfile=/Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/chronicle-parser-language-server/test/results.txt

echo "" > $outfile

for folder in $(ls /Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/cyderes-parsers/standard)
do
    path=/Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/cyderes-parsers/standard/$folder/conf/*.conf
    echo $folder | tee -a $outfile
    antlr4-parse /Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/chronicle-parser-language-server/antlr/ChronicleLogstashLexer.g4 /Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/chronicle-parser-language-server/antlr/ChronicleLogstashParser.g4 filterblock $path | tee -a $outfile
done