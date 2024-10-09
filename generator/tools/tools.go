//go:build generate

package tools

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)

// list of tags:
// $ cat swagger_flat.yaml| yq e '.. | select(has("tags")) | .tags[]' | sort -u | paste -sd ','                                                                                                                                                                                                                                   INT ✘  system 
// accounts,activeupdate,adminaccount,algorithms,api-policy,auth,auth-policy,autobackup,backup,certificates,certification authorities,cfgcheck,cfgdiff,cli,config,custom properties,deployment,deployment-warnings,disclaimer,encryption profiles,export,feature-toggling,firewalls,folders,ipsec-dr,ldap,license,lock,logs,message-boxes,misc,network,networkInterfaces,nsrpc,objects,package,proxy,QoS,radius,routing,rules,rule-sets,snsdiff,topologies,tunnels,update,usage,variables

// TODO: This doesn't work, because I don't know how to escape the spaces for some tags?
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --config=../client.yaml --include-tags="default,$(cat ../swagger_flat.yaml | yq e '.. | select(has("tags")) | .tags[]' | sort -u | grep -Ev 'firewalls|objects|proxy|rules|rule-sets' | paste -sd ',')" ../swagger_flat.yaml
