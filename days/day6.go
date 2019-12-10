package days

import (
	"fmt"
	"strings"
)

// Day6Part1 solves Day 6, Part 1
func Day6Part1(input []string) (string, error) {

	rootNodes := buildOrbitMap(input)
	allNodes := getAllNodes(rootNodes)

	count := 0
	for _, node := range allNodes {
		count += len(node.Chain())
	}

	return fmt.Sprintf("%d", count), nil
}

type orbitNode struct {
	identifier string
	orbits     *orbitNode
}

func (n orbitNode) Chain() []orbitNode {
	if n.orbits == nil {
		return []orbitNode{}
	}
	return append(n.orbits.Chain(), n)
}

func (n orbitNode) InOrbitOf(nodes []orbitNode) []orbitNode {
	orbiters := []orbitNode{}
	for _, node := range nodes {
		if node.orbits == nil {
			continue
		}

		if node.orbits.identifier == n.identifier {
			orbiters = append(orbiters, node)
			break
		}
	}

	return orbiters
}

func (n orbitNode) HasPathTo(to *orbitNode) (path bool, numSteps int) {
	chain := n.Chain()
	for i, node := range chain {
		if node.identifier == to.identifier {
			return true, len(chain) - i - 1
		}
	}
	return false, 0
}

func getNodeWithIdentifier(nodes []*orbitNode, identifier string) *orbitNode {
	for _, node := range nodes {
		if node.identifier == identifier {
			return node
		}
	}
	return nil
}

func getNodeWithIdentifierFromObjects(nodes []orbitNode, identifier string) *orbitNode {
	for _, node := range nodes {
		if node.identifier == identifier {
			return &node
		}
	}
	return nil
}

func getAllNodes(rootNodes []orbitNode) (nodes []orbitNode) {
	nodes = []orbitNode{}
	for _, rootNode := range rootNodes {

		currNode := rootNode
		nodes = append(nodes, currNode)

		isAtBase := false
		for !isAtBase {
			isAtBase = (currNode.orbits == nil)
			if isAtBase {
				break
			}

			currNode = *currNode.orbits

			alreadyAdded := false
			for _, addedNode := range nodes {
				if addedNode == currNode {
					alreadyAdded = true
					break
				}
			}

			if !alreadyAdded {
				nodes = append(nodes, currNode)
			}
		}
	}
	return nodes
}

func buildOrbitMap(input []string) (nodes []orbitNode) {
	// Start by building a slice of all the nodes. We'll remove duplicates later
	allNodes := []*orbitNode{}
	for _, orbitInstruction := range input {
		parts := strings.Split(orbitInstruction, ")")
		orbited, orbitedBy := parts[0], parts[1]
		allNodes = append(allNodes, &orbitNode{
			identifier: orbited,
		})
		allNodes = append(allNodes, &orbitNode{
			identifier: orbitedBy,
		})
	}

	// Now remove any duplicates
	uniqueNodes := []*orbitNode{}
	for i := 0; i < len(allNodes); i++ {
		nodeIsUnique := true
		for j := i + 1; j < len(allNodes); j++ {
			if allNodes[i].identifier == allNodes[j].identifier {
				nodeIsUnique = false
				break
			}
		}

		if nodeIsUnique {
			uniqueNodes = append(uniqueNodes, allNodes[i])
		}
	}

	// Now build the relationships
	var orbitedByNode, orbitedNode *orbitNode
	for _, orbitInstruction := range input {
		parts := strings.Split(orbitInstruction, ")")
		orbited, orbitedBy := parts[0], parts[1]

		orbitedByNode = getNodeWithIdentifier(uniqueNodes, orbitedBy)
		orbitedNode = getNodeWithIdentifier(uniqueNodes, orbited)
		orbitedByNode.orbits = orbitedNode
	}

	// Finally find all the nodes that aren't orbited by anything else
	nodes = []orbitNode{}
	for _, orbitedNode := range uniqueNodes {
		isOrbited := false
		for _, orbiter := range uniqueNodes {
			if orbiter.orbits != nil && *orbiter.orbits == *orbitedNode {
				isOrbited = true
				break
			}
		}
		if !isOrbited {
			nodes = append(nodes, *orbitedNode)
		}
	}

	return nodes
}

// Day6Part2 solves Day 6, Part 2
func Day6Part2(input []string) (string, error) {
	rootNodes := buildOrbitMap(input)
	allNodes := getAllNodes(rootNodes)

	you := getNodeWithIdentifierFromObjects(allNodes, "YOU")
	san := getNodeWithIdentifierFromObjects(allNodes, "SAN")

	_, d := firstCommonNodeBetween(allNodes, you, san)

	// minus 2 because the distance returned includes YOU and SAN
	return fmt.Sprintf("%d", d-2), nil
}

func firstCommonNodeBetween(allNodes []orbitNode, a, b *orbitNode) (node orbitNode, distance int) {
	for _, node := range allNodes {
		pathToA, stepsToA := a.HasPathTo(&node)
		pathToB, stepsToB := b.HasPathTo(&node)
		if pathToA && pathToB {
			return node, stepsToA + stepsToB
		}
	}
	fmt.Printf("No common node between %+v and %+v\n", a, b)
	return orbitNode{}, 0
}
