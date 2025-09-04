package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/jbenet/go-base58"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	client := medialive.NewFromConfig(cfg)
	// flowArn := ""
	input, err := client.DescribeSchedule(ctx, &medialive.DescribeScheduleInput{
		MaxResults: aws.Int32(100),
		ChannelId:  aws.String("2865342"),
	})
	if err != nil {
		log.Fatalf("failed to describe flow: %v", err)
	}
	for _, action := range input.ScheduleActions {
		fmt.Printf("%+v\n", action.ScheduleActionStartSettings.ImmediateModeScheduleActionStartSettings)
	}

	// cw := cloudwatch.NewFromConfig(cfg)
	// endTime := time.Now()
	// startTime := endTime.Add(-10 * time.Second)
	// resp, err := cw.GetMetricData(ctx, &cloudwatch.GetMetricDataInput{
	// 	StartTime: aws.Time(startTime),
	// 	EndTime:   aws.Time(endTime),
	// 	MetricDataQueries: []types.MetricDataQuery{
	// 		{
	// 			Id: aws.String("m1"),
	// 			MetricStat: &types.MetricStat{
	// 				Metric: &types.Metric{
	// 					Namespace:  aws.String("AWS/MediaConnect"),
	// 					MetricName: aws.String("SourceBitRate"),
	// 					Dimensions: []types.Dimension{
	// 						{Name: aws.String("SourceARN"), Value: aws.String("")},
	// 					},
	// 				},
	// 				Period: aws.Int32(1),
	// 				Stat:   aws.String("Minimum"),
	// 			},
	// 		},
	// 		{
	// 			Id: aws.String("m2"),
	// 			MetricStat: &types.MetricStat{
	// 				Metric: &types.Metric{
	// 					Namespace:  aws.String("AWS/MediaConnect"),
	// 					MetricName: aws.String("SourceBitRate"),
	// 					Dimensions: []types.Dimension{
	// 						{Name: aws.String("SourceARN"), Value: aws.String("")},
	// 					},
	// 				},
	// 				Period: aws.Int32(1),
	// 				Stat:   aws.String("Minimum"),
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to get metric data: %v", err))
	// }
	// for _, m := range resp.MetricDataResults {
	// 	fmt.Println("Metric ID:", *m.Id)
	// 	for _, v := range m.Values {
	// 		fmt.Printf("MBps: %.2fMBps\n", v/1e6)
	// 	}
	// }
}

func NewFlowID(flowArn string) string {
	ss := strings.Split(flowArn, "flow:")
	rawID := []byte(ss[len(ss)-1])
	return base58.Encode(rawID)
}
