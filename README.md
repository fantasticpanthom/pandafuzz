```
                                                  🅿🅰🅽🅳🅰 🅵🆄🆉🆉
```

:panda_face:fuzz is a minimalistic web fuzzer written in go.

# :scroll: 🅼🅴🅽🆄 

- [Usage](#Usage)
- [Installation](#Installation)

## Usage

Examples:

Simple jobs that can be performed using `pandafuzz`.

### Directory Discovery

```bash
>> pandafuzz -url https://www.example.com/fuzz -wl common.txt
```

### Get based parameter fuzzing
```bash
>> pandafuzz -url https://www.example.com?param=fuzz -wl common.txt
```

>Custom fuzz word can be provided by using the flag -fw 

Example :

```bash
>>pandafuzz -url https://www.example.com?param=panda -wl common.txt -fw panda
```

## Installation 
### From source:

```
>>go get -u -v github.com/fantasticpanthom/pandafuzz

or

>>cd $GOPATH
>>git clone https://github.com/fantasticpanthom/pandafuzz.git
>>go build statuscode.go -o statuscode
>>mv statuscode /usr/local/bin
```
### From binary:
You can download the pre-built binaries from the [releases](https://github.com/fantasticpanthom/pandafuzz/releases/) page and then move them into your $PATH.
Binaries have not been released yet.
Coming soon.
