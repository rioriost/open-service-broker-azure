package mysql

import (
	mysqlSDK "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql" // nolint: lll
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/open-service-broker-azure/pkg/azure/arm"
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

type module struct {
	allInOneServiceManager *allInOneManager
	dbmsManager            *dbmsManager
	databaseManager        *databaseManager
}

type allInOneManager struct {
	*dbmsManager
}

type dbmsManager struct {
	sqlDatabaseDNSSuffix        string
	armDeployer                 arm.Deployer
	checkNameAvailabilityClient mysqlSDK.CheckNameAvailabilityClient
	serversClient               mysqlSDK.ServersClient
}

type databaseManager struct {
	sqlDatabaseDNSSuffix string
	armDeployer          arm.Deployer
	databasesClient      mysqlSDK.DatabasesClient
}

// New returns a new instance of a type that fulfills the service.Module
// interface and is capable of provisioning MySQL DBMS and databases
// using "Azure Database for MySQL"
func New(
	azureEnvironment azure.Environment,
	armDeployer arm.Deployer,
	checkNameAvailabilityClient mysqlSDK.CheckNameAvailabilityClient,
	serversClient mysqlSDK.ServersClient,
	databaseClient mysqlSDK.DatabasesClient,
) service.Module {
	dm := &dbmsManager{
		sqlDatabaseDNSSuffix:        azureEnvironment.SQLDatabaseDNSSuffix,
		armDeployer:                 armDeployer,
		checkNameAvailabilityClient: checkNameAvailabilityClient,
		serversClient:               serversClient,
	}
	return &module{
		dbmsManager: dm,
		allInOneServiceManager: &allInOneManager{
			dbmsManager: dm,
		},
		databaseManager: &databaseManager{
			sqlDatabaseDNSSuffix: azureEnvironment.SQLDatabaseDNSSuffix,
			armDeployer:          armDeployer,
			databasesClient:      databaseClient,
		},
	}
}

func (m *module) GetName() string {
	return "mysql"
}
