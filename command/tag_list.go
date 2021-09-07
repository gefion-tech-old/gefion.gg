package command

import (
	"flag"
	"fmt"
	"net/url"

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
	// Проверка на валидность переданнных парметров
	for _, param := range []string{c.remote, c.password, c.username, c.ssh_key} {
		_, err := url.QueryUnescape(param)
		if err != nil {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.UTIL__ERROR,
					Message: fmt.Sprintf("Unable to decode passed parameter: `%s`", param),
				},
			}
		}
	}

	// Если указан репозиторий
	if c.remote != "" {
		if c.username != "" && c.password != "" && c.ssh_key == "" {
			// Подключаюсь к удаленному репозиторию
			rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
				Name: "origin",
				URLs: []string{fmt.Sprintf("https://%s:%s@%s", c.username, c.password, c.remote)},
			})

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
		} else if c.username == "" && c.password == "" && c.ssh_key != "" {
			// file, err := os.Create(os.Getenv("HOME") + "/.ssh/id_rsa")
			// if err != nil {
			// 	return m.Error{
			// 		Error: &m.ErrorBody{
			// 			Type:    m.UTIL__ERROR,
			// 			Message: fmt.Sprintf("Unable to create SSH key file: %s", err),
			// 		},
			// 	}
			// }
			// defer file.Close()

			// k, err := url.QueryUnescape(c.ssh_key)
			// if err != nil {
			// 	return m.Error{
			// 		Error: &m.ErrorBody{
			// 			Type:    m.UTIL__ERROR,
			// 			Message: fmt.Sprintf("Unable to decode SSH key: %s", err),
			// 		},
			// 	}
			// }

			// file.Write([]byte(k))
			// file.Chmod(0400)

			// ЗДЕСЬ ПОДКЛЮЧЕНИЕ К РЕП

			// err = os.Remove(os.Getenv("HOME") + "/.ssh/id_rsa")
			// if err != nil {
			// 	return m.Error{
			// 		Error: &m.ErrorBody{
			// 			Type:    m.GIT__ERROR,
			// 			Message: err.Error(),
			// 		},
			// 	}
			// }
		} else {
			return m.Error{
				Error: &m.ErrorBody{
					Type:    m.UTIL__ERROR,
					Message: "Invalid combination of arguments passed",
				},
			}
		}

	}

	return m.Error{
		Error: &m.ErrorBody{
			Type:    m.UTIL__ERROR,
			Message: "You must pass a remote host",
		},
	}
}
