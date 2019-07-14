package main

import (
	_ "github.com/evisong/registrator/consul"
	_ "github.com/evisong/registrator/consulkv"
	_ "github.com/evisong/registrator/etcd"
	_ "github.com/evisong/registrator/skydns2"
	_ "github.com/evisong/registrator/zookeeper"
)
