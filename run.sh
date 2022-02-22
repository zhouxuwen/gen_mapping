if [[ $# < 1 ]]; then
  echo "args count not match"
  echo "usage: ${0} [build|run]"
  exit -1
fi

case $1 in
"build")
  go build .
  ./gen_mapping -h
  ;;
"run")
  go build .
  ./gen_mapping -p $2 -i $3 -o $4
  ;;
*)
esac
