package redis

func TestRedisString(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	val, err := rds.Get(ctx, "test").Result()
}

func TestRedisInt(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()
	key := "count-Jeff"

	count, err := rds.Incr(ctx, key).Result()

	if err != nil {
		log.Printf("增加计数器失败: %v", err)
	}

	rds.Expire(ctx, key, time.Second*30)
	ttl, _ := rds.TTL(ctx, key).Result()

	fmt.Printf("count: %d", count)
	fmt.Printf("ttl: %d", ttl)
}

func TestRedisHash(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	userId := "user:1001"
	key := "cart:" + userId

	// Hash塞值
	rds.HIncrBy(ctx, key, "product:1001", 1)
	rds.HIncrBy(ctx, key, "product:1002", 2)

	// Hash取值
	cartHash, _ := rds.HGetAll(ctx, key).Result()
	fmt.Printf("cartHash: %v\n", cartHash)

	// Hash更新
	rds.HIncrBy(ctx, key, "product:1001", 1)
	cartHash, _ = rds.HGetAll(ctx, key).Result()
	fmt.Printf("updated cartHash: %v", cartHash)

	// 刪除
	rds.Del(ctx, key)
}

func TestRedisBitmap(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	keyFormat := "login-count:%s{group}" //login-count:20230701

	// 做幾個登入
	rds.SetBit(ctx, fmt.Sprintf(keyFormat, "20230801"), 10001, 1)
	rds.SetBit(ctx, fmt.Sprintf(keyFormat, "20230802"), 10001, 1)

	// 做幾個重複登入
	rds.SetBit(ctx, fmt.Sprintf(keyFormat, "20230801"), 10002, 1)
	rds.SetBit(ctx, fmt.Sprintf(keyFormat, "20230801"), 10002, 1)

	// 20230801 有多少人登入過
	monthCount, _ := rds.BitCount(ctx, fmt.Sprintf(keyFormat, "20230801"), nil).Result()
	fmt.Printf("%s count:%d\n", "20230801", monthCount)

	// 202308 有多少人登入過
	// 先聚合
	rds.BitOpOr(ctx, "login-count:202308{group}", "login-count:20230801{group}", "login-count:20230802{group}").Result()

	countTest, _ := rds.BitCount(ctx, "login-count:202308{group}", nil).Result()
	fmt.Printf("%s count:%d\n", "202308", countTest)
}

func TestRedisZset(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	rds.ZAdd(ctx, "game:leaderboard", redis.Z{
		Score:  1000,
		Member: "user:1001",
	})

	rds.ZAdd(ctx, "game:leaderboard", redis.Z{
		Score:  2000,
		Member: "user:1002",
	})

	var zset []redis.Z
	for i := 1; i < 50; i++ {
		zset = append(zset, redis.Z{
			Score:  float64(i * 10),           // 示例分数
			Member: fmt.Sprintf("user:%d", i), // 示例用户
		})
	}

	rds.ZAdd(ctx, "game:leaderboard", zset...)

	// get rankings
	leaderboard, _ := rds.ZRevRangeWithScores(ctx, "game:leaderboard", 0, 10).Result()
	fmt.Printf("leaderboard: %v\n", leaderboard)

	// 加分數
	rds.ZAddArgsIncr(ctx, "game:leaderboard", redis.ZAddArgs{
		Members: []redis.Z{
			// 3000
			{
				Score:  2000,
				Member: "user:1001",
			},
			// 2500
			{
				Score:  -500,
				Member: "user:1001",
			},
		},
	}).Result()

	// get rankings
	leaderboard, _ = rds.ZRevRangeWithScores(ctx, "game:leaderboard", 0, 10).Result()
	fmt.Printf("leaderboard: %v\n", leaderboard)
}

func TestRedisZsetIncr(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	members :=
		[]redis.Z{
			{
				Score:  500,
				Member: "user:1001",
			},
		}

	// ZIncr僅支援單一member
	_, err := rds.ZAddArgsIncr(ctx, "game:leaderboard", redis.ZAddArgs{
		Members: members,
	}).Result()

	if err != nil {
		log.Printf("ZAddArgsIncr error: %v", err)
	}

	// get rankings

	leaderboard, _ := rds.ZRevRangeWithScores(ctx, "game:leaderboard", 0, 10).Result()
	fmt.Printf("leaderboard: %v\n", leaderboard)
}

func TestRedisList(t *testing.T) {
	ctx := context.Background()
	rds := conn.GetRds()

	// 最新公告 , 保留5筆

	rds.LPush(ctx, "post", "msg1", "msg2", "msg3", "msg4", "msg5", "msg6")

	values, err := rds.LTrim(ctx, "post", 0, 4).Result()
	if err != nil {
		log.Printf("LRange error: %v", err)
	}

	fmt.Printf("values: %v\n", values)\
}