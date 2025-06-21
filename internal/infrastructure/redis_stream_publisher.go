package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"todo-app/internal/entities"
	"github.com/go-redis/redis/v8"
)

const (
	streamName = "todo-stream"
	groupName  = "todo-group"
	consumerID = "todo-consumer"
)

type RedisStreamPublisher struct {
	client *redis.Client
}

func NewRedisStreamPublisher(client *redis.Client) *RedisStreamPublisher {
	return &RedisStreamPublisher{client: client}
}

func (r *RedisStreamPublisher) Publish(ctx context.Context, item *entities.TodoItem) error {
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal todo item: %w", err)
	}

	_, err = r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"todoItem":  itemJSON,
			"timestamp": time.Now().Unix(),
		},
	}).Result()

	if err != nil {
		return fmt.Errorf("failed to publish to redis stream: %w", err)
	}

	return nil
}

func (r *RedisStreamPublisher) Subscribe(ctx context.Context) (<-chan *entities.TodoItem, error) {

	err := r.client.XGroupCreate(ctx, streamName, groupName, "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	todoChan := make(chan *entities.TodoItem)

	go func() {
		defer close(todoChan)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				
				streams, err := r.client.XReadGroup(ctx, &redis.XReadGroupArgs{
					Group:    groupName,
					Consumer: consumerID,
					Streams:  []string{streamName, ">"},
					Count:    1,
					Block:    0,
				}).Result()

				if err != nil {
					if err == redis.Nil {
						continue
					}
					
					fmt.Printf("Error reading from stream: %v\n", err)
					continue
				}
				
				for _, stream := range streams {
					for _, message := range stream.Messages {
						
						todoItemJSON, ok := message.Values["todoItem"].(string)
						if !ok {
							continue
						}
						
						var todoItem entities.TodoItem
						if err := json.Unmarshal([]byte(todoItemJSON), &todoItem); err != nil {
							fmt.Printf("Error unmarshaling todo item: %v\n", err)
							continue
						}
						
						select {
						case todoChan <- &todoItem:
						case <-ctx.Done():
							return
						}
					}
				}
			}
		}
	}()

	return todoChan, nil
}

func (r *RedisStreamPublisher) Close() error {
	return r.client.Close()
}
