/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	goflag "flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/version/verflag"
)

// ip = target for /whoami query
// rip = returned ip
// pip = this pod's ip
// nip = this node's ip

type MasqTester struct {
	Port string
}

func NewMasqTester() *MasqTester {
	return &MasqTester{
		Port: "8080",
	}
}

func (m *MasqTester) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&m.Port, "port", m.Port, "The port to serve /checknosnat and /whoami endpoints on.")
}

// NewMasqCommand creates a *cobra.Command object with default parameters
func NewMasqCommand() *cobra.Command {
	m := NewMasqTester()
	cmd := &cobra.Command{
		Use:  "MasqTest",
		Long: `MasqTest`,
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			cmdutil.CheckErr(m.Run())
		},
	}
	m.AddFlags(pflag.CommandLine)

	cmd.MarkFlagFilename("config", "yaml", "yml", "json")

	return cmd

}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := NewMasqCommand()
	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

func (m *MasqTester) Run() error {
	// pip is the current pod's IP and nip is the current node's IP
	// pull the pip and nip out of the env
	pip, ok := os.LookupEnv("POD_IP")
	if !ok {
		return fmt.Errorf("POD_IP env var was not present in the environment")
	}
	nip, ok := os.LookupEnv("NODE_IP")
	if !ok {
		return fmt.Errorf("NODE_IP env var was not present in the environment")
	}

	// validate that pip and nip are ip addresses.
	if net.ParseIP(pip) == nil {
		return fmt.Errorf("POD_IP env var contained %q, which is not an IP address", pip)
	}
	if net.ParseIP(nip) == nil {
		return fmt.Errorf("NODE_IP env var contained %q, which is not an IP address", nip)
	}

	// register handlers
	http.HandleFunc("/whoami", whoami)
	http.HandleFunc("/checknosnat", mkChecknosnat(pip, nip))

	// spin up the server
	return http.ListenAndServe(":"+m.Port, nil)
}

type handler func(http.ResponseWriter, *http.Request)

func joinErrors(errs []error, sep string) string {
	strs := make([]string, len(errs))
	for i, err := range errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, sep)
}

// Builds checknosnat handler, using pod and node ip of current location
func mkChecknosnat(pip string, nip string) handler {
	// Queries /whoami for each provided ip, resp 200 if all resp bodies match this Pod's ip, 500 otherwise
	return func(w http.ResponseWriter, req *http.Request) {
		errs := []error{}
		ips := strings.Split(req.URL.Query().Get("ips"), ",")
		for _, ip := range ips {
			if err := check(ip, pip, nip); err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", joinErrors(errs, ", "))
			return
		}
		w.WriteHeader(200)
	}
}

// Writes the req.RemoteAddr into the response, req.RemoteAddr is the address of the incoming connection
func whoami(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s", req.RemoteAddr)
}

// Queries ip/whoami and compares response to pip, uses nip to differentiate SNAT from other potential failure modes
func check(ip string, pip string, nip string) error {
	url := fmt.Sprintf("http://%s/whoami", ip)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rips := strings.Split(string(body), ":")
	if rips == nil || len(rips) == 0 {
		return fmt.Errorf("Invalid returned ip %q from %q", string(body), url)
	}
	rip := rips[0]
	if rip != pip {
		if rip == nip {
			return fmt.Errorf("Returned ip %q != my Pod ip %q, == my Node ip %q - SNAT", rip, pip, nip)
		} else {
			return fmt.Errorf("Returned ip %q != my Pod ip %q or my Node ip %q - SNAT to unexpected ip (possible SNAT through unexpected interface on the way into another node)", rip, pip, nip)
		}
	}
	return nil
}
