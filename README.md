# go-errors

A simple error handling package for Go that provides error wrapping with additional context.

## Installation

```bash
go get github.com/iam-kevin/go-errors
```

## Usage

### Creating Errors

```go
err := errors.New("something went wrong")
```

### Wrapping Errors

Add context to existing errors:

```go
func readConfig(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return errors.Wrap(err, "failed to read config")
    }
    // ...
    return nil
}
```

### Formatted Wrapping

```go
func getUser(id int) error {
    user, err := db.Find(id)
    if err != nil {
        return errors.Wrapf(err, "user %d not found", id)
    }
    // ...
    return nil
}
```

### Joining Multiple Errors

```go
func validate(data Data) error {
    var errs []error

    if data.Name == "" {
        errs = append(errs, errors.New("name required"))
    }
    if data.Email == "" {
        errs = append(errs, errors.New("email required"))
    }

    return errors.Join(errs...)
}
```

### Pre-defined Errors

```go
func notReady() error {
    return errors.ErrNotImplemented
}

func unsupportedFormat() error {
    return errors.ErrUnsupported
}
```

## API

- `New(message string) error` - Create new error
- `Wrap(err error, message string) error` - Wrap error with context
- `Wrapf(err error, format string, args ...any) error` - Wrap with formatted message
- `Join(errs ...error) error` - Combine multiple errors
- `ErrNotImplemented` - Pre-defined "not implemented" error
- `ErrUnsupported` - Pre-defined "unsupported" error
