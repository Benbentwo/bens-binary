package github

import (
	"strings"
)

func (c *AuthConfig) SetUserAuth(url string, auth *UserAuth) {
	username := auth.Username
	for i, server := range c.Servers {
		if urlsEqual(server.URL, url) {
			for j, a := range server.Users {
				if a.Username == auth.Username {
					c.Servers[i].Users[j] = auth
					c.Servers[i].CurrentUser = username
					c.DefaultUsername = username
					c.CurrentServer = url
					return
				}
			}
			c.Servers[i].Users = append(c.Servers[i].Users, auth)
			c.Servers[i].CurrentUser = username
			c.DefaultUsername = username
			c.CurrentServer = url
			return
		}
	}
	c.Servers = append(c.Servers, &AuthServer{
		URL:         url,
		Users:       []*UserAuth{auth},
		CurrentUser: username,
	})
	c.DefaultUsername = username
	c.CurrentServer = url

}

func urlsEqual(url1, url2 string) bool {
	return url1 == url2 || strings.TrimSuffix(url1, "/") == strings.TrimSuffix(url2, "/")
}

// DeleteServer deletes the server for the given URL and updates the current server
// if is the same with the deleted server
func (c *AuthConfig) DeleteServer(url string) {
	for i, s := range c.Servers {
		if urlsEqual(s.URL, url) {
			c.Servers = append(c.Servers[:i], c.Servers[i+1:]...)
		}
	}
	if urlsEqual(c.CurrentServer, url) && len(c.Servers) > 0 {
		c.CurrentServer = c.Servers[0].URL
	} else {
		c.CurrentServer = ""
	}
}
