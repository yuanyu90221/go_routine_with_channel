# go_with_channel

## un-buffered-channel with goroutine

    un-buffered channel

    every read will wait until something write into channel

    so we could use it as block

```golang===
c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		// time.Sleep(1 * time.Second)
		c <- true
	}()
	<-c
```

## buffered channel will work as async