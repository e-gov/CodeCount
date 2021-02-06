#!bin/sh
cat main.go | \
  awk ' /^\s*(if|else|for|switch|case)/ { printf "\n%s\n", $0 } 
    !/^\s*(if|else|for|switch|case)/ { printf ". " }
  ' 

# cat main.go | \
#  awk ' \
#    /^\s*(if|else|for|switch|case)/ { print $0 } \
#    !/^\s*(if|else|for|switch|case)/ { print "..." } ' 

# Asenda 'if' mittesisaldavad laused punktiga.
