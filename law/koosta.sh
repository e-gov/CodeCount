#!bin/sh
cat  RKVS.txt KOVVS.txt EPVS.txt RaHS.txt | \
  tr -d [:punct:] | \
  tr -d '§-' | \
  tr "[:upper:]" "[:lower:]" | \
  tr "Ä" "ä" | \
  tr "Õ" "õ" | \
  tr "Ü" "ü" | \
  tr "Õ" "õ" | \
  tr -d '[:digit:]' |  \
  tr -s " " | \
  tr '[:space:]' '[\n*]' | \
  grep -v "^\s*$" | \
  sort | \
  uniq | \
  sort > \
  SÕNAD.txt

# Sidurda neli tekstifaili
# eemalda punktuatsioon
# eemalda paragrahivmärk ja -
# teisenda väiketähtedeks, ka täpitähed
# eemalda numbrid
# eemalda mitmekordsed tühikud (squeeze)
# jaga sõnad eraldi ridadele
# eemalda tühjad read (võib olla üleliigne?)
# sorteeri

# cat EVPS.txt | \
#   tr -d [:punct:] | \
#   tr -d '§-' | \
#   tr "[:upper:]" "[:lower:]" | \
#   tr -d '[:digit:]' | \
#   tr '[:space:]' '[\n*]' | \
#   grep -v "^\s*$" | \
#   sort | \
#   uniq -c | \
#   sort -bnr > \
#   EVPS-sõnad

# Märkmed
# https://unix.stackexchange.com/questions/39039/get-text-file-word-occurrence-count-of-all-words-print-output-sorted
