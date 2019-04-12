# tcas-3rd-round-resolver

## Prerequisites

- [Go](https://golang.org/)
- [Dep](https://github.com/golang/dep)

## Installation

```
$ git clone https://github.com/kohpai/tcas-3rd-round-resolver
$ cd tcas-3rd-round-resolver/
$ dep ensure -v
$ go build -o resolver
```

## Usage

### Getting help

```
$ ./resolver --help
```

### Running

```
$ ./resolver --students data/TC01/con1_student_enroll.json --courses data/TC01/all_course.json --rankings data/TC01/con1_course_accept.csv --output output.csv
```
