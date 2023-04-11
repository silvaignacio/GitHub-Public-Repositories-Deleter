package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func main() {
	// Define los parámetros de autenticación
	token := "Access Token"

	// Crea un cliente de autenticación con el token personal
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	// Crea un cliente de Github con el cliente de autenticación
	client := github.NewClient(tc)

	// Lista todos los repositorios del usuario "username"
	repos, _, err := client.Repositories.List(context.Background(), "username", &github.RepositoryListOptions{
		Type: "public",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Elimina cada repositorio de la lista
	for _, repo := range repos {
		err := deleteRepo(repo.GetOwner().GetLogin(), repo.GetName(), token)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func deleteRepo(repoOwner, repoName, token string) error {
	// Crea un cliente de autenticación con el token personal
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	// Crea un cliente de Github con el cliente de autenticación
	client := github.NewClient(tc)

	// Elimina el repositorio especificado
	_, err := client.Repositories.Delete(context.Background(), repoOwner, repoName)
	if err != nil {
		return err
	}

	// Imprime un mensaje de confirmación
	fmt.Printf("El repositorio %s/%s ha sido eliminado correctamente.\n", repoOwner, repoName)
	return nil
}
