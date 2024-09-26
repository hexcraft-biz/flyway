package flyway

import (
	"fmt"
	"os"
	"os/exec"
)

type Flyway struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DirSqls    string
	BinPath    string
}

func New(host, port, name, user, password, dirSqls, binPath string) *Flyway {
	return &Flyway{
		DbHost:     host,
		DbPort:     port,
		DbName:     name,
		DbUser:     user,
		DbPassword: password,
		DirSqls:    dirSqls,
		BinPath:    binPath,
	}
}

func (f Flyway) FlywayMigrate() error {
	cmd := exec.Command(
		f.BinPath,
		"-url=jdbc:"+fmt.Sprintf("mysql://%s:%s/%s", f.DbHost, f.DbPort, f.DbName),
		"-user="+f.DbUser,
		"-password="+f.DbPassword,
		"-locations=filesystem:"+f.DirSqls,
		"migrate",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (f Flyway) FlywayClean() error {
	cmd := exec.Command(
		f.BinPath,
		"-url=jdbc:"+fmt.Sprintf("mysql://%s:%s/%s", f.DbHost, f.DbPort, f.DbName),
		"-user="+f.DbUser,
		"-password="+f.DbPassword,
		"-locations=filesystem:"+f.DirSqls,
		"clean",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
