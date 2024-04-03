package config

import (
	"log"
	"main/controllers"
	repository "main/repositories"
	"main/services"

	"go.uber.org/dig"
)

type Dependency struct {
	Instance interface{}
	Class    interface{}
	Name     string
}

var deps = []Dependency{
	{
		Instance: repository.DatabaseInstance,
		Class:    new(repository.IDatabase),
		Name:     "Database",
	},
	{
		Instance: repository.CompanyRepositoryInstance,
		Class:    new(repository.ICompanyRepository),
		Name:     "CompanyRepository",
	},
	{
		Instance: repository.NotificationRepositoryInstance,
		Class:    new(repository.INotificationRepository),
		Name:     "NotificationRepository",
	},
	{
		Instance: services.AuthenticationServiceInstance,
		Class:    new(services.IAuthenticationService),
		Name:     "AuthenticationService",
	},
	{
		Instance: services.NotificationServiceInstance,
		Class:    new(services.INotificationService),
		Name:     "NotificationService",
	},
	{
		Instance: controllers.NotificationControllerInstance,
		Class:    new(controllers.INotificationController),
		Name:     "NotificationController",
	},
	{
		Instance: controllers.HealthControllerInstance,
		Class:    new(controllers.IHealthController),
		Name:     "HealthController",
	},
	{
		Instance: controllers.AuthenticationControllerInstance,
		Class:    new(controllers.IAuthenticationController),
		Name:     "AuthenticationController",
	},
	{
		Instance: AuthenticatorInstance,
		Class:    new(IAuthenticator),
		Name:     "Authenticator",
	},
}

func Inject() *dig.Container {
	container := dig.New()

	for _, dep := range deps {
		err := container.Provide(
			dep.Instance,
			dig.As(dep.Class),
			dig.Name(dep.Name),
		)
		if err != nil {
			log.Fatalf("Failed to create the dependencies %s", err.Error())
		}
	}

	return container
}
