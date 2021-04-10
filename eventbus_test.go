package eventbus

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBus(t *testing.T) {
	// Initialize the event bus.
	bus := New()

	// Create some test data.
	eventName := "testEvent"
	eventTags := []string{"tag1", "tag2"}
	eventContent := "test content"

	// Create some listeners with various priorities.
	var listenerCalls []string
	listenerFunc1 := func(data EventData) {
		listenerCalls = append(listenerCalls, "listener1")
		assert.Equal(t, eventName, data.Name)
		assert.Equal(t, eventTags, data.Tags)
		assert.Equal(t, eventContent, data.Content)
	}
	listenerFunc2 := func(data EventData) {
		listenerCalls = append(listenerCalls, "listener2")
		assert.Equal(t, eventName, data.Name)
		assert.Equal(t, eventTags, data.Tags)
		assert.Equal(t, eventContent, data.Content)
	}
	listenerFunc3 := func(data EventData) {
		listenerCalls = append(listenerCalls, "listener3")
		assert.Equal(t, eventName, data.Name)
		assert.Equal(t, eventTags, data.Tags)
		assert.Equal(t, eventContent, data.Content)
	}

	// Add the listeners to the event bus with various priorities.
	bus.AddListener(eventName, listenerFunc1, 3)
	bus.AddListener(eventName, listenerFunc2, 1)
	bus.AddListener(eventName, listenerFunc3, 2)

	// Emit the event.
	bus.Emit(EventData{Name: eventName, Tags: eventTags, Content: eventContent})

	// Check that the listeners were called in the correct order.
	expectedListenerCalls := []string{"listener1", "listener3", "listener2"}
	assert.Equal(t, expectedListenerCalls, listenerCalls)

	// Remove old listeners.
	bus.Reset()

	// Test registering events from multiple goroutines.
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i uint) {
			bus.AddListener(eventName, func(data EventData) {
				listenerCalls = append(listenerCalls, fmt.Sprintf("l%d", i))
			}, i)
			wg.Done()
		}(uint(i))
	}
	wg.Wait()

	// Emit the event again.
	listenerCalls = []string{}
	bus.Emit(EventData{Name: eventName, Tags: eventTags, Content: eventContent})

	// Check that the listeners were called in the correct order.
	expectedListenerCalls = []string{"l9", "l8", "l7", "l6", "l5", "l4", "l3", "l2", "l1", "l0"}
	assert.Equal(t, expectedListenerCalls, listenerCalls)
}
