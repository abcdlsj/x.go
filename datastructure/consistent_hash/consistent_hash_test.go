package consistenthash

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	c := New(3, nil)

	// Add three nodes
	c.Add("node1", "node2", "node3")

	// Expect the ring to have 3 * replicas = 9 entries
	expected := 9
	actual := len(c.Ring)
	if actual != expected {
		t.Errorf("Expected %d ring entries, but got %d", expected, actual)
	}

	// Expect nodes to be distributed evenly across the ring
	nodeCounts := map[string]int{}
	for _, node := range c.Map {
		nodeCounts[node]++
	}
	for node, count := range nodeCounts {
		if count != c.Replicas {
			t.Errorf("Expected %d replicas for node %s, but got %d", c.Replicas, node, count)
		}
	}
}

func TestGet(t *testing.T) {
	c := New(100, nil)
	size := 100

	for i := 0; i < size; i++ {
		// Add three nodes
		c.Add(fmt.Sprintf("node%d", i))
	}

	// Generate random keys and record hit probabilities for each node
	numKeys := 10000
	hits := map[string]int{}
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		node := c.Get(key)
		hits[node]++
	}

	// Expect hit probabilities to be close to uniform
	expected := numKeys / size
	errors := 0
	for _, count := range hits {
		if count < expected/2 || count > expected*2 {
			errors++
		}
	}

	if errors > 5 {
		t.Errorf("Expected hit probabilities to be close to uniform, but got %d errors", errors)
	}
}

func TestRemove(t *testing.T) {
	c := New(3, nil)

	// Add three nodes
	c.Add("node1", "node2", "node3")

	// Remove a node
	c.Remove("node2")

	// Expect the ring to have 2 * replicas = 6 entries
	expected := 6
	actual := len(c.Ring)
	if actual != expected {
		t.Errorf("Expected %d ring entries, but got %d", expected, actual)
	}

	// Expect the removed node to have been removed from the ring
	for _, node := range c.Map {
		if node == "node2" {
			t.Errorf("Expected node2 to have been removed from the ring, but it was still present")
		}
	}
}
