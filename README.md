# Patchway

A JSON Patch library that makes it easy to construct patch operations by hand.

## Get It!

```bash
go get -u github.com/jamiecuthill/patchway@main
```

## Example Usage

```go
ops := patchway.Operations{
    patchway.Replace(first_name, "name", "first"),
    patchway.Replace(last_name, "name", "last"),
}
```

Use in conjunction with a [JSON Patch SDK](https://jsonpatch.com/#go) as follows

```go
patch, err := jsonpatch.DecodePatch(ops.Bytes())
if err != nil {
    panic(err)
}

modified, err := patch.Apply(original)
if err != nil {
    panic(err)
}
```
