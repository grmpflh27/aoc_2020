if [ -z "$1" ]
  then
    echo "Provide day - e.g. $0 1"
    exit
fi

cur_day=day_$1

cp -R day_template $cur_day
mv day_template/aoc_goday.tmpl $cur_day/day$1.go
sed -i "" "s/<DAY>/$1/g" $cur_day/day$1.go  

echo "workdir for day $1 generated"