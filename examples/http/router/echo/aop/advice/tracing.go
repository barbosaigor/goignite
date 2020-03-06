package advice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/wesovilabs/beyond/api"
	c "github.com/wesovilabs/beyond/api/context"
)

type TracingAdvice struct {
	prefix string
}

func (a *TracingAdvice) Before(ctx *c.BeyondContext) {

	log.Println("aqui")

	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *c.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	printTrace(ctx, a.prefix, params)
}

func NewTracingAdvice() api.Before {
	return &TracingAdvice{}
}

func NewTracingAdviceWithPrefix(prefix string) func() api.Before {
	return func() api.Before {
		return &TracingAdvice{
			prefix: prefix,
		}
	}
}

func printTrace(ctx *c.BeyondContext, prefix string, params []string) {

	l := logrus.FromContext(context.Background())

	if prefix == "" {
		l.Infof("%s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
		return
	}
	l.Infof("%s %s.%s(%s)\n", prefix, ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}
