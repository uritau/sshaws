package login

import (
	"fmt"
	"os"
	"time"
)

func pushTempKeyPair(destInstance, az string, IP string, username string) {
	destination := destInstance
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	time.Sleep(1 * time.Second)
	awscliPath := getCommandPath("aws")
	sshKeygenPath := getCommandPath("ssh-keygen")

	tempKeyPath := "/tmp/id_ed25519_temp"
	tempPubKeyPath := fmt.Sprintf("%s.pub", tempKeyPath)
	err := os.Remove(tempKeyPath)
	err = os.Remove(tempPubKeyPath)
	fmt.Println("Instance: ",destination)

	fmt.Println(">> Generating temporal key pair")
	proc, err := os.StartProcess(sshKeygenPath, []string{
		"ssh-keygen",
		"-t", "ed25519",
		"-f", tempKeyPath,
		"-N", "",
		"-q"},
		&pa)
	if err != nil {
		panic(err)
	}
	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Println(">> Pushing temporal pub key to instance ", destination)
	proc, err = os.StartProcess(awscliPath, []string{
		"aws",
		"ec2-instance-connect",
		"send-ssh-public-key",
		"--region", getRegionFromAZ(az),
		"--instance-id", destination,
		"--availability-zone", az,
		"--instance-os-user", username,
		"--ssh-public-key", "file://" + tempPubKeyPath},
		&pa)
	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println(">> Configure your SSH tunel with this parameters:")
	fmt.Println(`
		HOST/IP: `, IP,`
		User: `, username,`
		Autentication method: "Public key"
		Private Key:`, tempKeyPath, `
		`)
}
