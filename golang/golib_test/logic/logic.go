package logic

import (
	"context"
	"fmt"
	"sync"
)

var Seats = map[int]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
}

var m sync.Mutex

func Reserve(ctx context.Context, uid int, sid int) {
	//m.Lock()
	isUse := Seats[sid]
	if isUse {
		takeSeat(sid)
		fmt.Printf("uid:%d，sid:%d\n", uid, sid)
		//g.Log().Infof(ctx, "uid:%d，sid:%d", uid, sid)
	}
	//m.Unlock()
}

func takeSeat(sid int) {
	Seats[sid] = false
}
