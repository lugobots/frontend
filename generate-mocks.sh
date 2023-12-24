#!/bin/bash

mockgen -package=broker \
        -destination=./web/app/broker/mocks_lugo.go \
        github.com/lugobots/lugo4go/v3/lugo BroadcastClient,BroadcastServer,Broadcast_OnEventServer,Broadcast_OnEventClient

mockgen -package=broker \
        -destination=./web/app/broker/mocks_internal.go \
        github.com/lugobots/frontend/web/app/broker HitsCounter
