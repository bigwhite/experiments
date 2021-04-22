# ./http_client_load.sh 3 10000 10 GET http://10.10.195.181:8080
echo "$0 task_num count_per_hey conn_per_hey method url"
task_num=$1
count_per_hey=$2
conn_per_hey=$3
method=$4
url=$5

start=$(date +%s%N)
for((i=1; i<=$task_num; i++)); do {
	tm=$(date +%T.%N)
        echo "$tm: task $i start"
	hey -n $count_per_hey -c $conn_per_hey -m $method $url > hey_$i.log
	tm=$(date +%T.%N)
        echo "$tm: task $i done"
} & done
wait
end=$(date +%s%N)

count=$(( $task_num * $count_per_hey ))
runtime_ns=$(( $end - $start ))
runtime=`echo "scale=2; $runtime_ns / 1000000000" | bc`
echo "runtime: "$runtime
speed=`echo "scale=2; $count / $runtime" | bc`
echo "speed: "$speed 
