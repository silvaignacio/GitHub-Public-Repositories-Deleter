package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

// main es la función principal del programa.
//
// Utiliza el token de acceso personal para autenticarse en GitHub y eliminar todos los repositorios públicos del usuario "username".
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

// deleteRepo elimina un repositorio de GitHub.
//
// Parámetros:
// - repoOwner: El nombre de usuario del propietario del repositorio.
// - repoName: El nombre del repositorio que se eliminará.
// - token: El token de acceso personal para autenticar la solicitud de eliminación.
//
// Valor de retorno:
// - error: Si se produce un error durante la eliminación del repositorio, se devuelve un objeto de error que describe la causa del problema. Si la eliminación se realiza correctamente, se devuelve nil.
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
