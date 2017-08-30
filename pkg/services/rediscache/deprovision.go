package rediscache

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-broker/pkg/service"
)

func (m *module) GetDeprovisioner(
	string,
	string,
) (service.Deprovisioner, error) {
	return service.NewDeprovisioner(
		service.NewDeprovisioningStep("deleteARMDeployment", m.deleteARMDeployment),
		// krancour: This next step is a workaround because, currently, deleting
		// the ARM deployment is NOT deleting the PostgreSQL server. This seems to
		// be a problem not with ARM, but with the Postgres RP.
		service.NewDeprovisioningStep(
			"deleteRedisServer",
			m.deleteRedisServer,
		),
	)
}

func (m *module) deleteARMDeployment(
	ctx context.Context, // nolint: unparam
	instanceID string, // nolint: unparam
	serviceID string, // nolint: unparam
	planID string, // nolint: unparam
	provisioningContext service.ProvisioningContext,
) (service.ProvisioningContext, error) {
	pc, ok := provisioningContext.(*redisProvisioningContext)
	if !ok {
		return nil, fmt.Errorf(
			"error casting provisioningContext as redisProvisioningContext",
		)
	}
	if err := m.armDeployer.Delete(
		pc.ARMDeploymentName,
		pc.ResourceGroupName,
	); err != nil {
		return nil, fmt.Errorf("error deleting ARM deployment: %s", err)
	}
	return pc, nil
}

func (m *module) deleteRedisServer(
	ctx context.Context, // nolint: unparam
	instanceID string, // nolint: unparam
	serviceID string, // nolint: unparam
	planID string, // nolint: unparam
	provisioningContext service.ProvisioningContext,
) (service.ProvisioningContext, error) {
	pc, ok := provisioningContext.(*redisProvisioningContext)
	if !ok {
		return nil, fmt.Errorf(
			"error casting provisioningContext as redisProvisioningContext",
		)
	}
	if err := m.redisManager.DeleteServer(
		pc.ServerName,
		pc.ResourceGroupName,
	); err != nil {
		return nil, fmt.Errorf("error deleting redis server: %s", err)
	}
	return pc, nil
}