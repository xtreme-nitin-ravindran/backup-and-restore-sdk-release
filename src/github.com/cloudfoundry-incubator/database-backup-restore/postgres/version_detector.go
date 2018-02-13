package postgres

import (
	"log"

	"github.com/cloudfoundry-incubator/database-backup-restore/config"
	"github.com/cloudfoundry-incubator/database-backup-restore/version"
)

type ServerVersionDetector struct {
	psqlPath string
}

func NewServerVersionDetector(psqlPath string) ServerVersionDetector {
	return ServerVersionDetector{psqlPath: psqlPath}
}

func (d ServerVersionDetector) GetVersion(config config.ConnectionConfig) (version.DatabaseServerVersion, error) {
	cmdArgs := []string{
		"--tuples-only",
		config.Database,
		`--command=SELECT VERSION()`,
	}

	stdout, stderr, err := NewPostgresCommand(config, d.psqlPath).WithParams(cmdArgs...).Run()

	if err != nil {
		log.Fatalf("Unable to check version of Postgres: %v\n%s\n%s", err, string(stdout), string(stderr))
	}

	semVer, err := ParseVersion(string(stdout))
	if err != nil {
		log.Fatalf("Unable to check version of Postgres: %v", err)
	}

	return version.DatabaseServerVersion{
		Implementation:  "postgres",
		SemanticVersion: semVer,
	}, nil
}
