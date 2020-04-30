#!/bin/bash

mockgen -package=broker \
        -destination=./web/app/broker/mocks_lugo.go \
        github.com/lugobots/lugo4go/v2/lugo BroadcastClient,BroadcastServer,Broadcast_OnEventServer,Broadcast_OnEventClient