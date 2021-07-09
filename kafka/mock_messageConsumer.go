// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package kafka

import (
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	mock "github.com/stretchr/testify/mock"
)

// MockMessageConsumer is an autogenerated mock type for the MockMessageConsumer type
type MockMessageConsumer struct {
	mock.Mock
}

// Assign provides a mock function with given fields: partitions
func (_m *MockMessageConsumer) Assign(partitions []kafka.TopicPartition) error {
	ret := _m.Called(partitions)

	var r0 error
	if rf, ok := ret.Get(0).(func([]kafka.TopicPartition) error); ok {
		r0 = rf(partitions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockMessageConsumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Committed provides a mock function with given fields: partitions, timeoutMs
func (_m *MockMessageConsumer) Committed(partitions []kafka.TopicPartition, timeoutMs int) ([]kafka.TopicPartition, error) {
	ret := _m.Called(partitions, timeoutMs)

	var r0 []kafka.TopicPartition
	if rf, ok := ret.Get(0).(func([]kafka.TopicPartition, int) []kafka.TopicPartition); ok {
		r0 = rf(partitions, timeoutMs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.TopicPartition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]kafka.TopicPartition, int) error); ok {
		r1 = rf(partitions, timeoutMs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Events provides a mock function with given fields:
func (_m *MockMessageConsumer) Events() chan kafka.Event {
	ret := _m.Called()

	var r0 chan kafka.Event
	if rf, ok := ret.Get(0).(func() chan kafka.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan kafka.Event)
		}
	}

	return r0
}

// GetMetadata provides a mock function with given fields: topic, allTopics, timeoutMs
func (_m *MockMessageConsumer) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	ret := _m.Called(topic, allTopics, timeoutMs)

	var r0 *kafka.Metadata
	if rf, ok := ret.Get(0).(func(*string, bool, int) *kafka.Metadata); ok {
		r0 = rf(topic, allTopics, timeoutMs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafka.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, bool, int) error); ok {
		r1 = rf(topic, allTopics, timeoutMs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWatermarkOffsets provides a mock function with given fields: topic, partition, timeoutMs
func (_m *MockMessageConsumer) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (int64, int64, error) {
	ret := _m.Called(topic, partition, timeoutMs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, int32, int) int64); ok {
		r0 = rf(topic, partition, timeoutMs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string, int32, int) int64); ok {
		r1 = rf(topic, partition, timeoutMs)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int32, int) error); ok {
		r2 = rf(topic, partition, timeoutMs)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Subscribe provides a mock function with given fields: _a0, _a1
func (_m *MockMessageConsumer) Subscribe(_a0 string, _a1 kafka.RebalanceCb) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, kafka.RebalanceCb) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeTopics provides a mock function with given fields: _a0, _a1
func (_m *MockMessageConsumer) SubscribeTopics(_a0 []string, _a1 kafka.RebalanceCb) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, kafka.RebalanceCb) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unassign provides a mock function with given fields:
func (_m *MockMessageConsumer) Unassign() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
