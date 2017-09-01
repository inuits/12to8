# 12to8

A client for [9to5r](https://github.com/kalmanolah/925r)

# Install

```
$ go get github.com/Inuits/12to8
```

# Usage

List users:

```
12to8 users -e https://9TO5URL/api --user YOURUSER -p YOURPASS
```

List timesheets:

```
12to8 timesheet list
```

Create new timesheet:

```
12to8 timesheet new
12to8 timesheet new 9
12to8 timesheet new 9/2017
```
