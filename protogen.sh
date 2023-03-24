#!/usr/bin/env bash

gobin=$(go env GOBIN)

if test -z $gobin ; then
    gopath=$(go env GOPATH)
    if test -z $gopath; then
      gopath="$HOME/go"
    fi
    gobin="$gopath/bin"
fi

export PATH="$PATH:$gobin"

protoc=$(which protoc)

if test -z "$protoc"; then
  printf "protoc not found"
  exit 1
fi

gengo=$(which protoc-gen-go)
if test -z "$gengo"; then
  printf "gengo not found "
  exit  1
fi

root_dir="./protos"

function getdir(){
    for element in `ls $1`
    do
        dir_or_file=$1"/"$element
        if [ -d $dir_or_file ]
        then
            getdir $dir_or_file
        else
            echo $1
            echo  $dir_or_file
            ##注意：
                          #
                          #1> 提取文件后缀名： ${file##*.}
                          #
                          #    ##是贪婪操作符，从左至右匹配，匹配到最右边的.号，移除包含.号的左边内容。
                          #2> 是=，而且其两边有空格，如果没有空格，会报错
                          #
                          #3> 多加了x，是为了防止字符串为空时报错。

            if [ "${dir_or_file##*.}"x = "proto"x ]; then
              protoc -I $root_dir  --proto_path=$1 --validate_out="lang=go:$1" --go_out=$root_dir  --doc_out=$root_dir/doc --doc_opt=html,$dir_or_file.html $dir_or_file  $root_dir/*.proto
            fi
        fi
    done
}

getdir $root_dir

exec cp -r protos/github.com/phuhao00/greatestworks-proto/*  ./

