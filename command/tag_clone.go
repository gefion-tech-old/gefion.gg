package command

import (
	"flag"
	"fmt"
	"os"

	m "github.com/gefion-tech/gefion.gg/model"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type TagClone struct {
	fs *flag.FlagSet
	tag,
	remote,
	destination,
	hash,
	size,
	username,
	password string
}

func TagCloneCommand() *TagClone {
	tc := &TagClone{
		fs: flag.NewFlagSet("tag:clone", flag.ContinueOnError),
	}

	tc.fs.StringVar(&tc.tag, "t", "", "Tag.")
	tc.fs.StringVar(&tc.remote, "r", "", "Remote host.")
	tc.fs.StringVar(&tc.destination, "d", "", "Destination.")
	tc.fs.StringVar(&tc.hash, "h", "", "Hash.")
	tc.fs.StringVar(&tc.size, "s", "", "Remote repo size.")
	tc.fs.StringVar(&tc.username, "username", "", "Git uses a username to associate commits with an identity.")
	tc.fs.StringVar(&tc.password, "password", "", "Your account password or token.")

	return tc
}

// Получить имя субкоманды
func (c *TagClone) Name() string {
	return c.fs.Name()
}

// Получить инициализировать субкоманду
func (g *TagClone) Init(args []string) error {
	return g.fs.Parse(args)
}

// Выполнить субкоманду
func (c *TagClone) Run() interface{} {
	fmt.Println(c.remote)
	if c.remote != "" {
		_, err := git.PlainClone(c.destination, false, &git.CloneOptions{
			URL:      c.remote,
			Progress: os.Stdout,
			Auth: &http.BasicAuth{
				Username: c.username,
				Password: c.password,
			},
			Tags: git.TagFollowing,
		})

		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.GIT__ERROR,
					Message: err.Error(),
				},
			}
		}
	}

	return nil
}
