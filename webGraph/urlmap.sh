HTML_SRC="$HOME/Desktop/ir_proj/html"
DESC="$HOME/Desktop"
RESULT="$DESC/urlmap.txt"

pat=$(echo $HTML_SRC | sed -e 's/[]\/$*.^|[]/\\&/g')

find $HTML_SRC -type f | sed -e "s/^$pat\///g" | sed -e "s/^/http:\/\//" | sort > $RESULT