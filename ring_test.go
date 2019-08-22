package bestring

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_bestRing_Add(t *testing.T) {
	var secKey = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-./|=:#@!%^&$*&^~")

	type args struct {
		addr        []string
		vnodeNumber int
		requestKey  [][]byte
		randKey     int
	}
	tests := []struct {
		name           string
		args           args
		wantTotalNodes int
	}{
		{
			"test1",
			args{
				[]string{
					"192.168.0.1",
					"192.168.0.2",
					"192.168.0.3",
				},
				300,
				[][]byte{},
				10000,
			},
			300,
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
		})
	}
}
