# 12to8

A client for [9to5r](https://github.com/kalmanolah/925r)

## Install

```
$ go get github.com/Inuits/12to8
```

## Usage

### Lists
List users:

```
12to8 list users
```

List timesheets:

```
12to8 list timesheets
```

List companies:

```
12to8 list companies
```


### Actions

Create new timesheet:

```
12to8 new timesheet
12to8 new timesheet 9
12to8 new timesheet 9/2017
```

Release timesheet:
```
12to8 release timesheet
12to8 release timesheet -f
12to8 release timesheet 9
12to8 release timesheet 9/2017
```

Delete performance:
```
t list performances -P
t delete performance 3
```

Create a performance:
```
12to8 new performance 10/09/2017 8.0 -c "Consult [Zero Ci]"
```


### Configuration

You can create a config file `~/.12to8.yml` (or another path with `--config`).
We also support JSON, TOML, HCL, and Java properties config files. The choice is
yours.

```
---
endpoint: https://ninetofiver.example.com/api
user: your_username
password: your_password
```

If you do not want to set your password in a file, you can also set the password
as an env variable:

```
export TWELVE_TO_EIGHT_PASSWORD=your_password
```


### Completion

In your shell (& ~/.bashrc)

```
. <(12to8 completion bash)
```

or:

```
12to8 completion bash > ~/.12to8.complete
echo . ~/.12to8.complete >> ~/.bashrc
```
