// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

const (
	// RabbitMQService const
	RabbitMQService = "rabbitmq"

	// RabbitMQAMQPPort const
	RabbitMQAMQPPort = "5672"

	// RabbitMQDashboardPort const
	RabbitMQDashboardPort = "15672"

	// RabbitMQDockerImage const
	RabbitMQDockerImage = "rabbitmq:3.8-management"

	// RabbitMQRestartPolicy const
	RabbitMQRestartPolicy = "unless-stopped"

	// RabbitMQDefaultUsername const
	RabbitMQDefaultUsername = "guest"

	// RabbitMQDefaultPassword const
	RabbitMQDefaultPassword = "guest"
)

// GetRabbitMQConfig gets yaml definition object
func GetRabbitMQConfig(name string) DockerComposeConfig {
	services := make(map[string]Service)

	services[name] = Service{
		Image:   RabbitMQDockerImage,
		Restart: RabbitMQRestartPolicy,
		Ports: []string{
			RabbitMQAMQPPort,
			RabbitMQDashboardPort,
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
