/*
	The Service frontend is UI on a page. Each Service has a ServiceRepository
	and its usecases data. The architecture groundwork is the usecases that
	orchestrate everything relevant to the frontend. The usecases do not
	need to know underlying database nor transpotation layer. 
	At directory we run ~/usecases$ go test
*/
package usecases

import (
//	"github.com/khaiphong/mu/backend/entities"
)

func GetItems() {
}

/*
	The MuHub usecases may contain (1) coded integration of WiFi and IoT for Home
	automation and Car transport, (2) backup for offline from laptop, (3) alert
	messaging, (4) studio, (5) collaboration and production services.
*/