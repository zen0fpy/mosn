/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
*Host/主机：能够进行网络通信的实体（如移动设备、服务器上的应用程序）。
*
*Downstream/下游：下游主机连接到 Mosn，发送请求并接收响应。
*
*Upstream/上游：上游主机接收来自 Mosn 的连接和请求，并返回响应。
*
*Listener/监听器：监听器是命名网地址（例如，端口、unix domain socket等)，可以被下游客户端连接。Mosn 暴露一个或者多个监听器给下游主机连接。
*
*Cluster/集群：集群是指 Mosn 连接到的逻辑上相同的一组上游主机。Mosn 通过服务发现来发现集群的成员。Mosn 通过负载均衡策略决定将请求路由到哪个集群成员。
 */

package proxy

import (
	"context"

	"mosn.io/mosn/pkg/buffer"
	"mosn.io/mosn/pkg/network"
)

func init() {
	buffer.RegisterBuffer(&ins)
}

var ins = proxyBufferCtx{}

type proxyBufferCtx struct {
	buffer.TempBufferCtx
}

func (ctx proxyBufferCtx) New() interface{} {
	return new(proxyBuffers)
}

func (ctx proxyBufferCtx) Reset(i interface{}) {
	buf, _ := i.(*proxyBuffers)
	*buf = proxyBuffers{}
}

type proxyBuffers struct {
	stream  downStream
	request upstreamRequest
	info    network.RequestInfo
}

func proxyBuffersByContext(ctx context.Context) *proxyBuffers {
	poolCtx := buffer.PoolContext(ctx)
	return poolCtx.Find(&ins, nil).(*proxyBuffers)
}
