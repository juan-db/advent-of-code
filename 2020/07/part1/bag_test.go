package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleSimpleRule(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("dark orange bags contain no other bags.", &bags)

	assert.Len(t, bags, 1)

	assert.Contains(t, bags, "dark orange")

	b := bags["dark orange"]
	assert.Empty(t, b.children)
	assert.Empty(t, b.parents)
	assert.Equal(t, "dark orange", b.color)
}

func TestTwoSimpleRules(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("dark orange bags contain no other bags.", &bags)
	graphRule("dark purple bags contain no other bags.", &bags)

	assert.Len(t, bags, 2)

	assert.Contains(t, bags, "dark orange")
	assert.Contains(t, bags, "dark purple")

	orange := bags["dark orange"]
	assert.Empty(t, orange.children)
	assert.Empty(t, orange.parents)
	assert.Equal(t, "dark orange", orange.color)

	purple := bags["dark purple"]
	assert.Empty(t, purple.children)
	assert.Empty(t, purple.parents)
	assert.Equal(t, "dark purple", purple.color)
}

func TestSingleChild(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("light green bags contain 2 pale cyan bags.", &bags)

	assert.Len(t, bags, 2)

	assert.Contains(t, bags, "light green")
	assert.Contains(t, bags, "pale cyan")

	green := bags["light green"]
	assert.Len(t, green.children, 1)
	assert.Empty(t, green.parents)

	cyan := bags["pale cyan"]
	assert.Len(t, cyan.parents, 1)
	assert.Empty(t, cyan.children)

	l := link{green, cyan, 2}
	assert.Equal(t, green.children[0], &l)
	assert.Equal(t, cyan.parents[0], &l)
}

func TestTwoChildren(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("light green bags contain 2 pale cyan bags, 3 mirrored lavender bags.", &bags)

	assert.Len(t, bags, 3)

	assert.Contains(t, bags, "light green")
	assert.Contains(t, bags, "pale cyan")
	assert.Contains(t, bags, "mirrored lavender")

	green := bags["light green"]
	cyan := bags["pale cyan"]
	lavender := bags["mirrored lavender"]
	gcLink := &link{green, cyan, 2}
	glLink := &link{green, lavender, 3}

	greenExpected := &bag{"light green", []*link{}, []*link{gcLink, glLink}}
	assert.Equal(t, greenExpected, green)

	cyanExpected := &bag{"pale cyan", []*link{gcLink}, []*link{}}
	assert.Equal(t, cyanExpected, cyan)

	lavenderExpected := &bag{"mirrored lavender", []*link{glLink}, []*link{}}
	assert.Equal(t, lavenderExpected, lavender)
}

func TestTwoParents(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("light green bags contain 2 pale cyan bags.", &bags)
	graphRule("mirrored lavender bags contain 1 pale cyan bags.", &bags)

	assert.Len(t, bags, 3)

	assert.Contains(t, bags, "light green")
	assert.Contains(t, bags, "pale cyan")
	assert.Contains(t, bags, "mirrored lavender")

	green := bags["light green"]
	cyan := bags["pale cyan"]
	lavender := bags["mirrored lavender"]
	gcLink := &link{green, cyan, 2}
	lcLink := &link{lavender, cyan, 1}

	greenExpected := &bag{"light green", []*link{}, []*link{gcLink}}
	assert.Equal(t, greenExpected, green)

	cyanExpected := &bag{"pale cyan", []*link{gcLink, lcLink}, []*link{}}
	assert.Equal(t, cyanExpected, cyan)

	lavenderExpected := &bag{"mirrored lavender", []*link{}, []*link{lcLink}}
	assert.Equal(t, lavenderExpected, lavender)
}

func TestOneLevelNesting(t *testing.T) {
	bags := map[string]*bag{}

	graphRule("light green bags contain 2 pale cyan bags.", &bags)
	graphRule("pale cyan bags contain 3 mirrored lavender bags.", &bags)

	assert.Len(t, bags, 3)

	assert.Contains(t, bags, "light green")
	assert.Contains(t, bags, "pale cyan")
	assert.Contains(t, bags, "mirrored lavender")

	green := bags["light green"]
	cyan := bags["pale cyan"]
	lavender := bags["mirrored lavender"]
	gcLink := &link{green, cyan, 2}
	clLink := &link{cyan, lavender, 3}

	greenExpected := &bag{"light green", []*link{}, []*link{gcLink}}
	assert.Equal(t, greenExpected, green)

	cyanExpected := &bag{"pale cyan", []*link{gcLink}, []*link{clLink}}
	assert.Equal(t, cyanExpected, cyan)

	lavenderExpected := &bag{"mirrored lavender", []*link{clLink}, []*link{}}
	assert.Equal(t, lavenderExpected, lavender)
}