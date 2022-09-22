package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	sumpb2 "server/sumpb"
	"strconv"
	"time"
)

type process struct {
	rdb  *redis.Client
	mong *mongo.Client
}

func (p *process) Init() {
	p.rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	p.InitMongoDB()
}

func (p *process) InitMongoDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	var err error
	p.mong, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo connect err:", err.Error())
		return
	}

	// 检查连接
	err = p.mong.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("mongo ping err:", err.Error())
		return
	}
	fmt.Println("Connected to MongoDB succ!")

	// init database & collection
	coll := p.mong.Database("trinity").Collection("roledef")
	var rreq sumpb2.SumRequest
	rreq.Numbers = &sumpb2.Numbers{
		A: 100,
		B: 200,
	}
	_, err = coll.InsertOne(context.Background(), rreq)
	if err != nil {
		fmt.Println("insertone failed", err)
		return
	}
}

func makeRedisKey(req *sumpb2.SumRequest) string {
	return fmt.Sprintf("%d", req.Numbers.A)
}

// Add returns sum of two integers
func (p *process) Add(ctx context.Context, req *sumpb2.SumRequest) (*sumpb2.SumResponse, error) {
	//todo get jwt token from ctx
	var resp sumpb2.SumResponse

	//get result from redis
	ret, err := p.rdb.Get(ctx, makeRedisKey(req)).Result()
	if err == nil && ret != "" {
		// get from redis succ
		ival, _ := strconv.Atoi(ret)
		resp.Result = int64(ival)
		p.rdb.Expire(ctx, makeRedisKey(req), time.Second*10)
		fmt.Println("get succ from redis")
		return &resp, nil
	}

	//get result from mongodb, update redis
	var reqRecord sumpb2.Numbers
	coll := p.mong.Database("trinity").Collection("roledef")
	err = coll.FindOne(ctx, bson.D{{"A", req.Numbers.A}}).Decode(&reqRecord)
	if err == nil {
		resp.Result = reqRecord.B
		fmt.Println("get succ from mongo")
		return &resp, nil
	}

	//calc the request, update redis, mongodb
	reqRecord.B = req.GetNumbers().GetA() + req.GetNumbers().GetB()

	//todo,
	p.mong.Database("trinity").Collection("roledef").InsertOne(ctx, reqRecord)
	p.rdb.Set(ctx, makeRedisKey(req), reqRecord.B, time.Second*10)
	fmt.Println("update mongo redis")
	return &sumpb2.SumResponse{Result: reqRecord.B}, nil
}
