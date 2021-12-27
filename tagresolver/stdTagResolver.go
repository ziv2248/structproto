package tagresolver

import (
	"fmt"
	"strings"

	"github.com/structproto/internal"
)

var _ internal.TagResolver = StdTagResolver

func StdTagResolver(fieldname, token string) (*internal.Tag, error) {
	if len(token) > 0 {
		parts := strings.SplitN(token, ";", 2)
		var desc string
		if len(parts) == 2 {
			parts, desc = strings.Split(parts[0], ","), parts[1]
		} else {
			parts = strings.Split(token, ",")
		}
		name, flags := parts[0], parts[1:]

		for ii := 0; ii < len(name); ii++ {
			ch := name[ii]

			if ch == '_' || ch == '-' ||
				(ch >= 'a' && ch <= 'z') ||
				(ch >= 'A' && ch <= 'Z') ||
				(ch >= '0' && ch <= '9') {
				name = name[ii:]
				break
			}

			switch ch {
			case '*':
				flags = append(flags, internal.RequiredFlag)
			default:
				return nil, fmt.Errorf("unknow attribute symbole '%c'", ch)
			}
		}

		var tag *internal.Tag
		if len(name) > 0 && name != "-" {
			tag = &internal.Tag{
				Name:  name,
				Flags: flags,
				Desc:  desc,
			}
		}
		return tag, nil
	}
	return nil, nil
}
