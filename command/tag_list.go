package command

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	m "github.com/gefion-tech/gefion.gg/model"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

type TagList struct {
	fs *flag.FlagSet
	remote,
	username,
	ssh_key,
	password string
}

func TagListCommand() *TagList {
	tl := &TagList{
		fs: flag.NewFlagSet("tag:list", flag.ContinueOnError),
	}

	tl.fs.StringVar(&tl.remote, "r", "", "Remote host.")
	tl.fs.StringVar(&tl.username, "username", "", "Git uses a username to associate commits with an identity.")
	tl.fs.StringVar(&tl.password, "password", "", "Your account password or token.")
	tl.fs.StringVar(&tl.ssh_key, "ssh_key", "", "Private SSH key.")

	return tl
}

func (c *TagList) Name() string {
	return c.fs.Name()
}

func (c *TagList) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *TagList) Run() interface{} {
	if c.remote != "" {
		file, err := os.Create("/home/I0HuKc/.ssh/id_rsa")
		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.UTIL__ERROR,
					Message: fmt.Sprintf("Unable to create SSH key file: %s", err),
				},
			}
		}
		defer file.Close()

		k, err := url.QueryUnescape(c.ssh_key)
		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.UTIL__ERROR,
					Message: fmt.Sprintf("Unable to decode SSH key: %s", err),
				},
			}
		}

		file.Write([]byte(k))
		file.Chmod(0400)

		// Подключаюсь к удаленному репозиторию
		rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
			// Name: "origin",
			URLs: []string{"git@github.com:I0HuKc/vanga.git"},
		})

		err = os.Remove("/home/I0HuKc/.ssh/id_rsa")
		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.GIT__ERROR,
					Message: err.Error(),
				},
			}
		}

		refs, err := rem.List(&git.ListOptions{})
		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.GIT__ERROR,
					Message: err.Error(),
				},
			}
		}

		// Фильтрую список и сохраняю нужные данные о тегах
		var tags = make([]m.TagRemote, 0)
		for _, ref := range refs {
			if ref.Name().IsTag() {
				tags = append(tags, m.TagRemote{
					Name: ref.Name().Short(),
					Hash: ref.Hash().String(),
				})
			}
		}
		return m.TagResponse{
			Tags: tags,
		}
	}

	return m.Error{
		Error: &m.ErrorBody{
			Type:    m.UTIL__ERROR,
			Message: "You must pass a remote host",
		},
	}
}

func keyEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
