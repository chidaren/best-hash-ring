package bestring

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func genAddr(cnt int) []string {
	res := make([]string, 0, cnt)

	for i := 0; i < cnt; i++ {
		res = append(res, fmt.Sprintf("192.168.0.%d", i))
	}
	return res
}

func Test_bestRing_Add(t *testing.T) {
	var secKey = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-./|=:#@!%^&$*&^~")

	type args struct {
		addr        []string
		vnodeNumber int
		requestKey  [][]byte
		randKey     int
	}

	var vNodeNumber int = 3

	tests := []struct {
		name           string
		args           args
		wantTotalNodes int
	}{
		{
			"test1",
			args{
				genAddr(20),
				vNodeNumber,
				[][]byte{},
				100000,
			},
			vNodeNumber,
		},
	}

	var randSeed = rand.NewSource(time.Now().Unix())
	var rander = rand.New(randSeed)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var tmpIndex int
			for i := 0; i < tt.args.randKey; i++ {
				a := rander.Intn(100)
				requestKey := make([]byte, 0, a)

				for i := a; i > 0; i-- {
					tmpIndex = rander.Intn(len(secKey))
					requestKey = append(requestKey, secKey[tmpIndex])
				}
				tt.args.requestKey = append(tt.args.requestKey, requestKey)
			}

			b := NewBestRing(tt.args.vnodeNumber)

			serverHit := make(map[string]int)
			for i := 0; i < len(tt.args.addr); i++ {
				serverHit[tt.args.addr[i]] = 0
				if gotTotalNodes := b.AddNode(tt.args.addr[i]); gotTotalNodes != tt.wantTotalNodes*(i+1) {
					t.Errorf("bestRing.AddNode() = %v, want %v", gotTotalNodes, tt.wantTotalNodes)
				}
			}

			for i := 0; i < len(tt.args.requestKey); i++ {
				server, _ := b.GetNode(tt.args.requestKey[i])
				serverHit[server] += 1
			}

			for s, c := range serverHit {
				fmt.Printf("Server:\t[%s]\thits  [%d]\ttimes\n", s, c)
			}

			for i := 0; i < len(tt.args.addr)/2; i++ {
				b.DeleteNode(tt.args.addr[i])
			}

			fmt.Println("After delete half nodes: ")

			serverHitAfterDelete := make(map[string]int)

			for i := 0; i < len(tt.args.requestKey); i++ {
				server, _ := b.GetNode(tt.args.requestKey[i])
				serverHitAfterDelete[server] += 1
			}

			for s, c := range serverHitAfterDelete {
				fmt.Printf("Server:\t[%s]\thits  [%d]\ttimes\n", s, c)
			}

			fmt.Println("After add all nodes: ")

			serverHitFinally := make(map[string]int)

			for i := 0; i < len(tt.args.addr); i++ {
				serverHitFinally[tt.args.addr[i]] = 0
				b.AddNode(tt.args.addr[i])
			}

			for i := 0; i < len(tt.args.requestKey); i++ {
				server, _ := b.GetNode(tt.args.requestKey[i])
				serverHitFinally[server] += 1
			}

			for s, c := range serverHitFinally {
				fmt.Printf("Server:\t[%s]\thits  [%d]\ttimes\n", s, c)
			}
		})
	}
}
