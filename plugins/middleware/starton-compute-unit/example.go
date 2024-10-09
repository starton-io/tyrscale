package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"

	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
)

type UseComputeUnit struct {
	projectId string
	amount    int
}

var startonRedisUri string
var redisClient redis.UniversalClient
var ctx context.Context = context.Background()
var channel chan UseComputeUnit = make(chan UseComputeUnit, 100)

// ComputeUnitMiddleware is an example middleware that logs requests
func ComputeUnitMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		projectId := string(ctx.Request.Header.Peek("x-consumer-custom-id"))
		if projectId == "" {
			log.Println("Could not get projectId")
			// TODO return json
			ctx.Error("You must be authenticated", fasthttp.StatusUnauthorized)
			return
		}

		projectSetting, err := redisClient.HGetAll(ctx, fmt.Sprintf("project:setting.%s", projectId)).Result()
		if err != nil {
			log.Println("Could not fetch projectSettings", err)
		}
		if projectSetting["maxComputeUnitReach"] == "true" {
			log.Printf("maxComputeUnitReach: %s", projectSetting["maxComputeUnitReach"])
			// TODO return json
			ctx.Error("You have reached your maximum comput unit.", fasthttp.StatusPaymentRequired)
			return
		}
		next(ctx)
		// TODO CUSTOMIZE PRICING BY METHOD
		channel <- UseComputeUnit{projectId, 5}
	}
}

type ComputeUnitMiddlewareRegister struct{}

func (p *ComputeUnitMiddlewareRegister) RegisterMiddleware(registerFunc func(name string, middleware types.MiddlewareFunc), payload []byte) error {
	startonRedisUri = os.Getenv("startonRedisUri")
	if startonRedisUri == "" {
		log.Fatalln("you must set startonRedisUri")
	}

	redisOptions := &redis.UniversalOptions{
		Addrs:      []string{startonRedisUri},
		MaxRetries: 1,
	}

	redisClient = redis.NewUniversalClient(redisOptions)
	go func() {
		for msg := range channel {
			countComputeUnit(msg)
		}
	}()
	registerFunc("ComputeUnitMiddleware", ComputeUnitMiddleware)
	return nil
}

func (p *ComputeUnitMiddlewareRegister) Validate(configPayload []byte) error {
	return nil
}

// Exported symbol
var Middleware ComputeUnitMiddlewareRegister

func main() {}

func countComputeUnit(msg UseComputeUnit) {
	res, err := redisClient.HIncrBy(ctx, "project:compute-unit", msg.projectId, int64(msg.amount)).Result()
	if err != nil {
		log.Println("Could not increase compute unit", err)
	}
	log.Println("Increased compute unit", res)

	usage, err := redisClient.ZIncrBy(ctx, "project:usage", float64(msg.amount), "RPC_CALL").Result()
	if err != nil {
		log.Println("Could not count project usage", err)
	}
	log.Println("Project usage updated", usage)
}
