# Second

## `sql.ErrNoRows` handling

Should we wrap the error when we encounter a `sql.ErrNoRows` in database operations, such as the dao layer, and throw it to the upper layer. Why and what should be done please write the code?

### Opaque error handling

```golang
// do not expose customize errors
var errNotFound = errors.New("not found")

// provide method for checking error type 
func IsNoRow(err error) bool {
	return errors.Is(err, errNotFound)
}
```