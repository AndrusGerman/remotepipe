package server

import (
	"fmt"
	"math/rand"
	"time"
)

func newID() []byte {
	var id = fmt.Sprintf("%d%d%d", time.Now().UnixNano(), rand.Int31(), rand.Int31n(20))
	var buffID = make([]byte, 1024)

	copy(buffID, []byte(id))

	return buffID
}
