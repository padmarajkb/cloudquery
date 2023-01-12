package support

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createServices(router *mux.Router) error {
	var item armsupport.ServicesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/providers/Microsoft.Support/services", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestServices(t *testing.T) {
	client.MockTestHelper(t, Services(), createServices)
}
