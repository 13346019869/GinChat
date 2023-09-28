/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 14:57:55
 */
package models

import (
	"context"
	"fmt"
	"ginchat/utils"
	"time"
)

/*
*
设置在线用户到redis缓存
*
*/
func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	set := utils.Red.Set(ctx, key, val, timeTTL)
	fmt.Println(set.Err())
}
