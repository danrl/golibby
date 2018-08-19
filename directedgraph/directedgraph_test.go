package directedgraph

import "testing"

var (
	nodes = []struct {
		key   string
		value interface{}
	}{
		{
			key:   "foo",
			value: "bar",
		},
		{
			key:   "eleven",
			value: 11,
		},
		{
			key:   "friends",
			value: '🤩',
		},
		{
			key:   "scary",
			value: 1337,
		},
		{
			key:   "ocean's",
			value: "11!!!",
		},
	}
	edges = []struct {
		from, to string
	}{
		{from: "foo", to: "eleven"},
		{from: "friends", to: "eleven"},
		{from: "eleven", to: "scary"},
	}
)

func TestGraphNew(t *testing.T) {
	t.Run("initial values", func(t *testing.T) {
		g := New()
		if len(g.nodes) != 0 {
			t.Errorf("initial node list not empty")
		}
		if len(g.edges) != 0 {
			t.Errorf("initial edge list not empty")
		}
	})
}

func TestGraphNewNode(t *testing.T) {
	t.Run("one node", func(t *testing.T) {
		g := New()
		g.NewNode("foo", "bar")
		if len(g.nodes) != 1 {
			t.Errorf("unexpected node list length: %v", len(g.nodes))
		}
		if len(g.edges) != 1 {
			t.Errorf("unexpected edge list length: %v", len(g.edges))
		}
	})
	t.Run("duplicate nodes", func(t *testing.T) {
		g := New()
		g.NewNode("foo", "bar")
		err := g.NewNode("foo", "bar")
		if err != ErrorNodeAlreadyExists {
			t.Errorf("expected `%v` got `%v`", ErrorNodeAlreadyExists, err)
		}
		if len(g.nodes) != 1 {
			t.Errorf("unexpected node list length: %v", len(g.nodes))
		}
	})
	t.Run("multiple nodes", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, nd := range nodes {
			if value, ok := g.nodes[nd.key]; !ok {
				t.Errorf("expected node `%v` not found in node list", nd.key)
			} else if value != nd.value {
				t.Errorf("expected node value `%v`, got `%v`", nd.value, value)
			}
		}
	})
}

func TestGraphValue(t *testing.T) {
	t.Run("retrieve values", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, nd := range nodes {
			value, err := g.Value(nd.key)
			if err != nil {
				t.Errorf("node `%v`: %v", nd.key, err)
			}
			if value != nd.value {
				t.Errorf("expected node value `%v`, got `%v`", nd.value, value)
			}
		}
	})
	t.Run("accessing unknown node", func(t *testing.T) {
		g := New()
		_, err := g.Value("foo")
		if err != ErrorNodeNotFound {
			t.Errorf("expected `%v` got `%v`", ErrorNodeNotFound, err)
		}
	})
}

func TestGraphUpdateValue(t *testing.T) {
	t.Run("update value of existing node", func(t *testing.T) {
		g := New()
		g.NewNode("foo", nil)
		for _, nd := range nodes {
			err := g.UpdateValue("foo", nd.value)
			if err != nil {
				t.Errorf("node `%v`: %v", nd.key, err)
			}
			if value := g.nodes["foo"]; value != nd.value {
				t.Errorf("expected node value `%v`, got `%v`", nd.value, value)
			}
		}
	})
	t.Run("accessing unknown node", func(t *testing.T) {
		g := New()
		err := g.UpdateValue("foo", nil)
		if err != ErrorNodeNotFound {
			t.Errorf("expected `%v` got `%v`", ErrorNodeNotFound, err)
		}
	})
}

func TestGraphNewEdge(t *testing.T) {
	t.Run("existing nodes", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			err := g.NewEdge(e.from, e.to)
			if err != nil {
				t.Errorf("edge from `%v` to `%v`: %v", e.from, e.to, err)
			}
		}
	})
	t.Run("unknown source node", func(t *testing.T) {
		g := New()
		g.NewNode("foo", nil)
		err := g.NewEdge("unknown", "foo")
		if err != ErrorNodeNotFound {
			t.Errorf("expected `%v` got `%v`", ErrorNodeNotFound, err)
		}
	})
	t.Run("unknown destination node", func(t *testing.T) {
		g := New()
		g.NewNode("foo", nil)
		err := g.NewEdge("foo", "unknown")
		if err != ErrorNodeNotFound {
			t.Errorf("expected `%v` got `%v`", ErrorNodeNotFound, err)
		}
	})
}

func TestGraphEdges(t *testing.T) {
	t.Run("existing nodes", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			g.NewEdge(e.from, e.to)
		}
		for _, e := range edges {
			to, err := g.Edges(e.from)
			if err != nil {
				t.Errorf("node `%v`: %v", e.from, err)
			}
			found := false
			for i := range to {
				if to[i] == e.to {
					found = true
				}
			}
			if !found {
				t.Errorf("expected edge `%v`->`%v` not found.", e.from, e.to)
			}
		}
	})
	t.Run("unknown nodes", func(t *testing.T) {
		g := New()
		for _, e := range edges {
			_, err := g.Edges(e.from)
			if err != ErrorNodeNotFound {
				t.Errorf("expected `%v` got `%v`", ErrorNodeNotFound, err)
			}
		}
	})
}

func TestGraphNodes(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g := New()
		if len(g.Nodes()) != 0 {
			t.Errorf("expected empty graph, got non-empty graph")
		}
	})
	t.Run("regular graph", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		n := g.Nodes()
		if len(n) != len(nodes) {
			t.Errorf("expected `%v` nodes, got `%v`", len(nodes), len(n))
		}
		for _, nd := range nodes {
			found := false
			for i := range n {
				if n[i] == nd.key {
					found = true
				}
			}
			if !found {
				t.Errorf("expected node `%v` not found.", nd.key)
			}
		}
	})
}

func TestGraphIsCyclic(t *testing.T) {
	t.Run("acyclic graph", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			g.NewEdge(e.from, e.to)
		}
		if got := g.IsCyclic(); got {
			t.Errorf("expected `false` got `%v`", got)
		}
	})
	t.Run("cyclic graph (back edge)", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			g.NewEdge(e.from, e.to)
		}
		g.NewEdge("scary", "foo")
		if got := g.IsCyclic(); !got {
			t.Errorf("expected `true` got `%v`", got)
		}
	})
	t.Run("cyclic graph (self-referencing node)", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			g.NewEdge(e.from, e.to)
		}
		g.NewEdge("eleven", "eleven")
		if got := g.IsCyclic(); !got {
			t.Errorf("expected `true` got `%v`", got)
		}
	})
}

func TestGraphString(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g := New()
		got := g.String()
		if got != "" {
			t.Errorf("expected empty string, got `%s`", got)
		}
	})
	// not the greatest test, but good enough to detect non-cosmetic changes
	t.Run("regular graph", func(t *testing.T) {
		g := New()
		for _, nd := range nodes {
			g.NewNode(nd.key, nd.value)
		}
		for _, e := range edges {
			g.NewEdge(e.from, e.to)
		}
		got := g.String()
		if len(got) != 136 {
			t.Errorf("expected string length `%v`, got `%v`", 136, len(got))
		}
	})
}
