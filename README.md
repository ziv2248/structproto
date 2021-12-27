structproto - StructPrototype
=============================


## Synopsis
```go
import (
  "github.com/structproto"
  "github.com/structproto/valuebinder"
)


type mockCharacter struct {
  Name       string    `demo:"*NAME"`
  Age        *int      `demo:"*AGE"`
  Alias      []string  `demo:"ALIAS"`
  DatOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark     string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
  Numbers    []int     `demo:"NUMBERS"`
}

func main() {
  c := mockCharacter{}
  prototype, err := structproto.Prototypify(&c, &structproto.StructProtoOption{
    TagName: "demo",
  })
  if err != nil {
    panic(err)
  }
  
  err = prototype.BindValues(NamedValues{
    "NAME":          "luffy",
    "ALIAS":         "lucy",
    "DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
  }, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
}
```
