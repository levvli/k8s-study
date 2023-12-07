package abac

import (
	"bufio"
	"fmt"
	"k8s.io/kubernetes/pkg/apis/abac"
	"os"
)

type policyLoadError struct {
	path string
	line int
	data []byte
	err  error
}

func (p policyLoadError) Error() string {
	if p.line >= 0 {
		return fmt.Sprintf("error reading policy file %s, line %d: %s: %v", p.path, p.line, string(p.data), p.err)
	}
	return fmt.Sprintf("error reading policy file %s: %v", p.path, p.err)
}

// PolicyList is simply a slice of Policy structs
type PolicyList []*abac.Policy

func NewFromFile(path string) (PolicyList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pl := make(PolicyList, 0)
	decoder := abac.Codecs.Univer
}
