# tagt: Friend of table driven tests
tagt is test helper that prints tags with original error message.

## Usage

```go
testdata := []struct{
	id  int
	exists bool
	tag string
}{
	{1, true, "valid user"},
	{9, false, "unknown user"},
}
for _, data := range testdata {
	tt := tagt.New(t, data.tag)
	user := repo.FindByID(data.id)
	if d.exists && user == nil {
		t.Errorf("user should be found by id is %d", data.id)
		// => user should be found by id is 1
		tt.Errorf("user should be found by id is %d", data.id)
		// => [valid user] user should be found by id is 1
	} else if !d.exists && user != nil {
		t.Errorf("user should not be found by id is %d", data.id)
		// => user should not be found by id is 9
		tt.Errorf("user should not be found by id is %d", data.id)
		// => [unknown user] user should not be found by id is 9
	}
}
```
