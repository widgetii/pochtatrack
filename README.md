## Quick examples

### History service

```go
	ws := pochtatrack.NewRTM34("LOGIN", "PASSWORD")

	res, err := ws.GetOperationHistory("19003132224427")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
```

### Batch service

```go
	ws := pochtatrack.NewFC("LOGIN", "PASSWORD")
	ticket := ws.GetTicket([]string{"19003132224427", "12727630039983", "80088332714549"})
	time.Sleep(900 * time.Second)
	fmt.Println(ws.GetResponse(ticket))
```

