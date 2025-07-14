package main

import (
	"context"
	"fmt"
	infisical "github.com/infisical/go-sdk"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	//Get ENVs for step
	infisicalEnvs := StepInput{
		Client:       os.Getenv("infisical_client"),
		ClientSecret: os.Getenv("infisical_client_secret"),
		ProjectId:    os.Getenv("infisical_project_id"),
		Env:          os.Getenv("infisical_env"),
		ApiUrl:       os.Getenv("infisical_url"),
		Path:         os.Getenv("infisical_path"),
	}

	err := ValidateConfig(infisicalEnvs)
	if err != nil {
		log.Fatalln(err)
	}

	//Set config, Cloud API is default site URL.
	var config = infisical.Config{
		UserAgent: "Bitrise-infisical-step",
	}

	if infisicalEnvs.ApiUrl != "" {
		config.SiteUrl = infisicalEnvs.ApiUrl
	}

	client := infisical.NewInfisicalClient(context.Background(), config)

	_, err = client.Auth().UniversalAuthLogin(infisicalEnvs.Client, infisicalEnvs.ClientSecret)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	secretsList, err := client.Secrets().List(infisical.ListSecretsOptions{
		ProjectID:          infisicalEnvs.ProjectId,
		Environment:        infisicalEnvs.Env,
		SecretPath:         infisicalEnvs.Path,
		AttachToProcessEnv: false,
	})
	if err != nil {
		log.Fatalf("Secret Error: %v", err)
	}

	for _, secret := range secretsList {
		cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", secret.SecretKey, "--value",
			secret.SecretValue).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
			os.Exit(1)
		}

		if os.Getenv("BITRISE_DEBUG") == "true" {
			fmt.Printf("%s=%s\n", secret.SecretKey, secret.SecretValue)
		}
	}

	secretCountStr := strconv.Itoa(len(secretsList))

	fmt.Printf("\033[1;32m%s\033[0m", "Injecting "+secretCountStr+
		" Infisical secrets into your application process\n")

	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
