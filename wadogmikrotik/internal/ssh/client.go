package ssh

import (
    "fmt"
    "io"
    "os"

    "golang.org/x/crypto/ssh"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/config"
)

type Client interface {
    RunCommand(host string, cfg config.SSHConfig, command string) error
}

type client struct{}

func New() Client {
    return &client{}
}

func (c *client) RunCommand(host string, cfg config.SSHConfig, command string) error {
    authMethods := []ssh.AuthMethod{}
    if cfg.Password != "" {
        authMethods = append(authMethods, ssh.Password(cfg.Password))
    }
    if cfg.PrivateKey != "" {
        privateKey, err := os.ReadFile(cfg.PrivateKey)
        if err != nil {
            return fmt.Errorf("no se puede leer private_key: %w", err)
        }
        signer, err := ssh.ParsePrivateKey(privateKey)
        if err != nil {
            return fmt.Errorf("error parseando clave privada: %w", err)
        }
        authMethods = append(authMethods, ssh.PublicKeys(signer))
    }
    if len(authMethods) == 0 {
        return fmt.Errorf("no auth method configured")
    }

    sshConfig := &ssh.ClientConfig{
        User:            cfg.User,
        Auth:            authMethods,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Timeout:         cfg.Timeout,
    }

    conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", host), sshConfig)
    if err != nil {
        return fmt.Errorf("error conectando por SSH a %s: %w", host, err)
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        return fmt.Errorf("error creando sesión SSH a %s: %w", host, err)
    }
    defer session.Close()

    stdout, err := session.StdoutPipe()
    if err != nil {
        return fmt.Errorf("error obteniendo stdout SSH: %w", err)
    }
    stderr, err := session.StderrPipe()
    if err != nil {
        return fmt.Errorf("error obteniendo stderr SSH: %w", err)
    }

    if err := session.Start(command); err != nil {
        return fmt.Errorf("error iniciando comando SSH: %w", err)
    }

    io.Copy(io.Discard, stdout)
    io.Copy(io.Discard, stderr)

    if err := session.Wait(); err != nil {
        return fmt.Errorf("comando SSH falló: %w", err)
    }

    return nil
}
