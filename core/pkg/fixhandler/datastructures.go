package fixhandler

import (
	"github.com/armosec/armoapi-go/armotypes"
	metav1 "github.com/kubescape/kubescape/v2/core/meta/datastructures/v1"
	"github.com/kubescape/opa-utils/reporthandling"
	reporthandlingv2 "github.com/kubescape/opa-utils/reporthandling/v2"
	"gopkg.in/yaml.v3"
)

// FixHandler is a struct that holds the information of the report to be fixed
type FixHandler struct {
	fixInfo       *metav1.FixInfo
	reportObj     *reporthandlingv2.PostureReport
	localBasePath string
}

// ResourceFixInfo is a struct that holds the information about the resource that needs to be fixed
type ResourceFixInfo struct {
	YamlExpressions map[string]*armotypes.FixPath
	Resource        *reporthandling.Resource
	FilePath        string
}

// NodeInfo holds extra information about the node
type NodeInfo struct {
	node   *yaml.Node
	parent *yaml.Node

	// position of the node among siblings
	index int
}

// FixInfoMetadata holds the arguments "getFixInfo" function needs to pass to the
// functions it uses
type FixInfoMetadata struct {
	originalList        *[]NodeInfo
	fixedList           *[]NodeInfo
	originalListTracker int
	fixedListTracker    int
	contentToAdd        *[]ContentToAdd
	contentToRemove     *[]ContentToRemove
}

// ContentToAdd holds the information about where to insert the new changes in the existing yaml file
type ContentToAdd struct {
	// Line where the fix should be applied to
	Line int
	// Content is a string representation of the YAML node that describes a suggested fix
	Content string
}

// ContentToRemove holds the line numbers to remove from the existing yaml file
type ContentToRemove struct {
	startLine int
	endLine   int
}
