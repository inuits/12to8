#!/bin/bash -xe

docker run --name ninetofiver -p 8000:8000 -v $PWD/tests:/tests --rm roidelapluie/925r /tests/run-925r.sh &>925r.log &

stop(){
    cat 925r.log
    cat results.log
    docker kill ninetofiver
}

max=300
timeout $max /bin/bash -c "i=0
while ! curl 127.0.0.1:8000 -s;
do
    let i++
    echo wait for 925r... \$i/$max
    sleep 1
done;
" || stop

export TWELVE_TO_EIGHT_USER=user
export TWELVE_TO_EIGHT_PASSWORD=pass
export TWELVE_TO_EIGHT_ENDPOINT=http://127.0.0.1:8000/api

> results.log

run(){
    echo $1
    if bash -xec "$2"; then
        echo $1 .. ok | tee -a results.log
    else
        echo $1 .. fail  | tee -a results.log
        return 1
    fi
}
ok(){
    if ! run "$1" "$2"; then
        stop
        return 1
    fi
}
fail(){
    if run "$1" "$2"; then
        echo "$1 should have failed"  | tee -a results.log
        stop
        return 1
    else
        echo $1 .. expected failure  | tee -a results.log
    fi
}

# TIMESHEETS TESTS
# we run multiple times list timesheets to help debugging
ok list_timesheets "./12to8 list timesheets"
ok zero_timesheet "[[ \$(./12to8 list timesheets|wc -l) -eq 0 ]]"
ok new_timesheet "./12to8 new timesheet"
ok one_timesheet "[[ \$(./12to8 list timesheets|wc -l) -eq 1 ]]"
ok active_timesheet "./12to8 list timesheets|grep ACTIVE"
fail release_timesheet_no_appr "echo n | ./12to8 release timesheet"
ok list_timesheets "./12to8 list timesheets"
ok active_timesheet_no_appr "./12to8 list timesheets|grep ACTIVE"
ok release_timesheet "./12to8 release timesheet -f"
ok list_timesheets "./12to8 list timesheets"
ok active_timesheet "./12to8 list timesheets|grep PENDING"


stop || true
