#!/bin/bash



BANNER="############################################################
# testbed selection                                        #
############################################################"

getNetByNum() {
  local ind=$1 i=0
  [ -z "$ind" ] && echo "ERR: no argument passed [index of network]" return 0
  for it in $__NETWORKS__; do
    ((i++))
    [ "$i" -eq "$ind" ] && echo "$it" && return 0
  done
  echo "ERR: docker network not found by index"
  return 1
}

#
choiceNet() {
  local name=$1 ind max n list i answer msg num=0 reg_num='^[0-9]+$'

  list=$(filterNet "$name")
  max="$?"

  if [ "$max" -eq 0 ]; then
    list=$(filterNet "")
    max="$?"
  fi

  [ "$max" -eq 0 ] && echo "ERR: no networks" && return 0
  if [ "$max" -eq 1 ]; then
    getIndNet "$list"
    return $?
  fi
  msg=$(strUnion "$BANNER" \
    "$(msgChoiceNet "$list")" \
    "List of docker networks: Choose a number [1 - $max] ? ")

  while true; do
    read -r -p "$msg" answer
    if [ "$answer" ] && [[ $answer =~ $reg_num ]]; then
      i=0
      for it in $list; do
        ((i++))
        [ "$i" -eq "$answer" ] && num=$i && break
      done
    fi
    [ "$num" ] && break
  done

  return "$num"
}

msgChoiceNet() {
  for it in $1; do
    ((ind++))
    echo "$ind) $it"
  done
}

strUnion() {
  [ "$1" ] && echo "$1"
  [ "$2" ] && echo "$2"
  [ "$3" ] && echo "$3"
  [ "$4" ] && echo "$4"
}

filterNet() {
  local name n count ind
  name=$1
  count=0
  for it in $__NETWORKS__; do
    if [ "$name" ]; then
      n=$(echo "$it" | grep -i "$name")
      [ -z "$n" ] && continue
    fi
    echo "$it"
    ((count++))
  done
  return "$count"
}

# returns the docker network number. Index starts at 1
getIndNet() {
  local name=$1 ind=0
  [ -z "$name" ] && echo "ERR: no argument passed [name of network]" return 0
  for it in $__NETWORKS__; do
    ((ind++))
    [ "$it" = "$name" ] && return "$ind"
  done
  echo "ERR: could not determine the network [$name]" return 0
}


