# tclient

Simple telnet client lib, written in golang.

Example usage:

main.go:
```
	client := tclient.New(5, "")
	err := client.Open("10.10.10.10", 23)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// you can omit this, or do auth stuff manually by calling `ReadUntil` with login/password prompts
	out, err := client.Login("script2", "pw3")
	if err != nil {
		panic(err)
	}
	fmt.Printf(out)

	out, err = client.Cmd("show time")
	if err != nil {
		panic(err)
	}
	fmt.Printf(out)
```

Output: 

![Output](https://i.imgur.com/2M91MEN.png)




# TODO

~~Implement pagination parsing/manipulating (various network devices paging their output)~~

Implement parsing regexps with callbacks. For output pagination purposes, etc.
