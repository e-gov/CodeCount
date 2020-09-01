#!bin/sh
cat  ÜhtestatudSõnad.txt | \
  sort | \
  uniq | \
  sort > \
  ÜhtestatudKordumatud.txt

# eemalda tühjad read (võib olla üleliigne?)
# sorteeri

# Märkmed
# https://unix.stackexchange.com/questions/39039/get-text-file-word-occurrence-count-of-all-words-print-output-sorted
