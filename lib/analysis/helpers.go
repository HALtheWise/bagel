package analysis

import (
	"errors"
	"fmt"

	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

func GetDefaultFiles(providers []loading.Provider) ([]refs.CFileRef, error) {
	var defaultInfo *loading.Provider

	for i, provider := range providers {
		if provider.Kind == loading.DefaultInfo {
			if defaultInfo != nil {
				return nil, fmt.Errorf("Multiple copies of %s", provider.Kind.Name())
			}
			defaultInfo = &providers[i]
		}
	}

	if defaultInfo == nil {
		return nil, errors.New("No DefaultInfo provider")
	}

	files, ok := defaultInfo.Data["files"].(*loading.Depset)
	if !ok {
		return nil, errors.New("DefaultInfo.files is not a depset")
	}

	var result []refs.CFileRef
	for _, value := range files.Items {
		result = append(result, value.(*BzlFile).Ref)
	}
	return result, nil
}
