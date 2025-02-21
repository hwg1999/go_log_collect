package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	assert.NoError(t, err, "init etcd failed")
	fmt.Println("init etcd success.")
	defer cli.Close()

	// put操作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/name", "jerry")
	cancel()
	assert.NoError(t, err, "put to etcd failed")
	fmt.Println("put to etcd success.")

	// get操作
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/name")
	cancel()
	assert.NoError(t, err, "get from etcd failed")

	for _, v := range resp.Kvs {
		fmt.Printf("Key:%s,Value:%s\n", v.Key, v.Value)
	}
}
