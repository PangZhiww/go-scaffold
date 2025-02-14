// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lestrrat-go/file-rotatelogs"
	"go-scaffold/internal/app"
	"go-scaffold/internal/app/command"
	greet4 "go-scaffold/internal/app/command/handler/greet"
	"go-scaffold/internal/app/command/script"
	"go-scaffold/internal/app/component/casbin"
	"go-scaffold/internal/app/component/client/grpc"
	"go-scaffold/internal/app/component/discovery"
	"go-scaffold/internal/app/component/orm"
	"go-scaffold/internal/app/component/redis"
	"go-scaffold/internal/app/component/trace"
	config2 "go-scaffold/internal/app/config"
	"go-scaffold/internal/app/cron"
	"go-scaffold/internal/app/cron/job"
	"go-scaffold/internal/app/repository/user"
	"go-scaffold/internal/app/service/greet"
	user2 "go-scaffold/internal/app/service/user"
	"go-scaffold/internal/app/transport"
	grpc2 "go-scaffold/internal/app/transport/grpc"
	greet3 "go-scaffold/internal/app/transport/grpc/handler/v1/greet"
	user4 "go-scaffold/internal/app/transport/grpc/handler/v1/user"
	"go-scaffold/internal/app/transport/http"
	greet2 "go-scaffold/internal/app/transport/http/handler/v1/greet"
	trace2 "go-scaffold/internal/app/transport/http/handler/v1/trace"
	user3 "go-scaffold/internal/app/transport/http/handler/v1/user"
	"go-scaffold/internal/app/transport/http/router"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func initApp(rotateLogs *rotatelogs.RotateLogs, logLogger log.Logger, zapLogger *zap.Logger, configConfig *config2.Config) (*app.App, func(), error) {
	ormConfig := configConfig.DB
	db, cleanup2, err := orm.New(ormConfig, logLogger, zapLogger)
	if err != nil {
		return nil, nil, err
	}
	traceConfig := configConfig.Trace
	tracer, cleanup3, err := trace.New(traceConfig, logLogger)
	if err != nil {
		cleanup2()
		return nil, nil, err
	}
	redisConfig := configConfig.Redis
	client, cleanup4, err := redis.New(redisConfig, logLogger)
	if err != nil {
		cleanup3()
		cleanup2()
		return nil, nil, err
	}
	example := job.NewExample(logLogger)
	cronCron, err := cron.New(logLogger, db, client, example)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		return nil, nil, err
	}
	configApp := configConfig.App
	configHTTP := configConfig.HTTP
	jwt := configConfig.JWT
	casbinConfig := configConfig.Casbin
	enforcer, err := casbin.New(casbinConfig, db)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		return nil, nil, err
	}
	service := greet.NewService(logLogger)
	handler := greet2.NewHandler(logLogger, service)
	discoveryConfig := configConfig.Discovery
	discoveryDiscovery, err := discovery.New(discoveryConfig, zapLogger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		return nil, nil, err
	}
	grpcClient := grpc.New(logLogger, discoveryDiscovery)
	traceHandler := trace2.NewHandler(logLogger, configConfig, tracer, grpcClient)
	repository := user.NewRepository(db, client)
	userService := user2.NewService(logLogger, repository)
	userHandler := user3.NewHandler(logLogger, userService)
	apiV1Group := router.NewAPIV1Group(handler, traceHandler, userHandler)
	apiGroup := router.NewAPIGroup(logLogger, configApp, configHTTP, jwt, enforcer, apiV1Group)
	httpHandler := router.New(configApp, configHTTP, rotateLogs, zapLogger, logLogger, apiGroup)
	server := http.NewServer(logLogger, configHTTP, httpHandler)
	configGRPC := configConfig.GRPC
	greetHandler := greet3.NewHandler(logLogger, service)
	handler2 := user4.NewHandler(logLogger, userService, repository)
	grpcServer := grpc2.NewServer(logLogger, configGRPC, greetHandler, handler2)
	transportTransport := transport.New(logLogger, configApp, server, grpcServer, discoveryDiscovery)
	appApp := app.New(logLogger, db, tracer, cronCron, transportTransport, enforcer)
	return appApp, func() {
		cleanup4()
		cleanup3()
		cleanup2()
	}, nil
}

func initCommand(rotateLogs *rotatelogs.RotateLogs, logLogger log.Logger, zapLogger *zap.Logger, configConfig *config2.Config) (*command.Command, func(), error) {
	handler := greet4.NewHandler(logLogger)
	s0000000000 := script.NewS0000000000(logLogger)
	commandCommand := command.New(handler, s0000000000)
	return commandCommand, func() {
	}, nil
}
