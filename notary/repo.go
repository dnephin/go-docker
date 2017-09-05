// +build !pkcs11

package notary // import "golang.docker.com/go-docker/notary"

import (
	"fmt"

	"github.com/docker/notary"
	"github.com/docker/notary/trustmanager"
)

func getKeyStores(baseDir string, retriever notary.PassRetriever) ([]trustmanager.KeyStore, error) {
	fileKeyStore, err := trustmanager.NewKeyFileStore(baseDir, retriever)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key store in directory: %s", baseDir)
	}
	return []trustmanager.KeyStore{fileKeyStore}, nil
}
