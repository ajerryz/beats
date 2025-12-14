package main

import "math/rand"

var ApiLogMockSlice = []ApiLogMock{
	{
		Endpoint: "https://api.internal.com/v1",
		Url:      "/v1/logical/test",
		Method:   "GET",
		Status:   200,
		Qps:      3,
	},
	{
		Endpoint: "https://api.internal.com/v1",
		Url:      "/v1/logical/blb/{blbId}",
		Method:   "POST",
		Status:   200,
		Qps:      10,
	},
	{
		Endpoint: "https://api.internal.com/v1",
		Url:      "/v1/blb/create",
		Method:   "POST",
		Status:   500,
		Qps:      2,
	},
	{
		Endpoint: "https://logical.internal.com/v1",
		Url:      "/logical/t1",
		Method:   "GET",
		Status:   200,
		Qps:      10,
	},
	{
		Endpoint: "https://logical.internal.com/v1",
		Url:      "/logical/t2/{productId}}",
		Method:   "GET",
		Status:   200,
		Qps:      33,
	},
	{
		Endpoint: "https://logical.internal.com/v1",
		Url:      "/logical/t3/{productId}}",
		Method:   "PUT",
		Status:   404,
		Qps:      5,
	},
}

type ApiLogMock struct {
	Endpoint string
	Url      string
	Method   string
	Status   int
	Qps      int
}

func CostTimeRandByMaxMin(min, max int) int {
	r := rand.Intn(max)
	if r < min {
		return min
	}
	return r - min
}

func CostTimeRand() int {
	return CostTimeRandByMaxMin(2, 300)
}
