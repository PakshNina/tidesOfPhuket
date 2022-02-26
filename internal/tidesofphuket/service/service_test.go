package service

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	apiClient "tidesOfPhuket/internal/tidesofphuket/client"
	"tidesOfPhuket/internal/tidesofphuket/coordinates"
	"tidesOfPhuket/internal/tools/testing/client"
	"tidesOfPhuket/internal/tools/testing/redis"
)

func getTidesExtremeResult() *apiClient.TidesExtremeResult {
	return &apiClient.TidesExtremeResult{
		Extremes: []apiClient.TidesExtreme{
			{
				Date:     1136203445,
				Height:   1.0005,
				TideType: High,
			},
			{
				Date:     1136203400,
				Height:   1.0005,
				TideType: Low,
			},
		},
	}
}
func TestGetTidesInfo(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	currentTime := time.Date(2006, time.January, 2, 16, 4, 5, 0, loc)
	monkey.Patch(time.Now, func() time.Time { return currentTime })

	// Mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_client.NewMockWorldTidesClient(ctrl)
	redis := mock_redis.NewMockRedis(ctrl)
	redis.EXPECT().GetExtremesForToday(PatongCommand).Return(getTidesExtremeResult(), nil)

	result := make(chan string)
	coords := coordinates.GetPatongCoordinates()
	go getTidesInfo(PatongCommand, coords, result, client, redis)
	var replyMessage string
	replyMessage = <-result
	reply := "Now it is 20:04 (Bangkok time). rLast tide was 02.01 19:03, 1.0005 meters Low.\n\nUpcoming tides on /patong\n\nHigh:\n\nLow:\n"
	assert.NotNil(t, replyMessage)
	assert.Equal(t, reply, replyMessage)
}
