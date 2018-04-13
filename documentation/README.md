# Structure

## pkg/api/...

IPMQ api implementation, you can query for a certain Kind

```go

// Handler maybe of ReadCloser & WriteCloser?
handler := api.Scheme.Lookup(kind, version)
handler.Handle(...)

```

## pkg/apis/...

