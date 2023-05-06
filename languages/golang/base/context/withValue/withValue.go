package withValue

import (
	"context"
	"fmt"
	"time"
)

func withValueNewContext(key, val string) context.Context {
	return context.WithValue(context.Background(), key, val)
}

func getValueWithString(ctx context.Context, key interface{}) string {
	// 断言
	val, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return val
}

func withValuePrintLog(ctx context.Context, key, msg string) {
	val := getValueWithString(ctx, key)
	fmt.Printf("%v|%s:%s|%s", time.Now().Format("2006-01-02 15:04:05"), key, val, msg)
}
