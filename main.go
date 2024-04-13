package main


import (
"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
"fmt"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		sgArgs := &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				ec2.SecurityGroupIngressArgs{
					Protocol: pulumi.String("tcp"),
					FromPort: pulumi.Int(8080),
					ToPort: pulumi.Int(8080),
					CidrBlocks: pulumi. StringArray{pulumi.String("0.0.0.0/0")},
				},
				ec2.SecurityGroupIngressArgs{
					Protocol: pulumi.String("tcp"),
					FromPort: pulumi.Int(22),
					ToPort: pulumi.Int(22),
					CidrBlocks: pulumi. StringArray{pulumi.String("0.0.0.0/0")},
				},	
			},	
			
			Egress: ec2.SecurityGroupEgressArray{
				ec2.SecurityGroupEgressArgs{
					Protocol: pulumi.String("-1"),
					FromPort: pulumi.Int(0),
					ToPort: pulumi.Int(0),
					CidrBlocks: pulumi. StringArray{pulumi.String("0.0.0.0/0")},
				},
			},	
		}
		
		sg, err := ec2.NewSecurityGroup(ctx, "jenkins-sg", sgArgs)
		if err != nil {
			return err
		}
		
		kp, err := ec2.NewKeyPair(ctx, "local-ssh", &ec2.KeyPairArgs{
			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCsI5N/aZHjyNQGfj1WQ5kGYHj1DJh2DxU9KOjWgShIVA4OFebnCjpd69tNEUE7fUwS9LKrVzApbZiPFwdPkJ5eqg+EXr8B5uW2wfxYDyljLmD7pShFg9v61QNQZYqrQRFaWMstdkMWy8iwTxUKYGylC98ZMyqEhNzyuTTFOT6WOCk8LupRjR7ZLM/J0UeMrsUGU7JnxnTMsGl+s5/sqe7T0YIW5edd3oq+dd40w+AQ6k4wWwUBmdw7kNRfPkY0i1bTf15pET3AlxS0IGJKRPlJapwFODM7SzFps6PKcFNFkBOp/eQSRj3/wWFecttPv0yJCyzQLJtp8fY+10sxxUyiYcw/L0rGXHbZePk2eABWRLKb/299Nn9KMmTaSuqh0Voi2GIbmzwv73WvvwDdNN/Vi08kasmpyTE4OyF9dhPausvlgpvBaJFmdxpO+nMng6EeHcW1zx6QnhVshCO/Mw6eSJF7c5OEtXp6TJRcOkJkY9hewzOipWIcfC9BuVjyGmGHaruy0jOEjdi33+c7lt0ovATG+hHh+SQQcqhruTz8lZR0HsVa3Eoz2EOmxw3ex+S+VDQU2fSu/qFWPj56CrmGkOyQvd1bixHQEnZ0edwj7HqxuxWfhA19QkeUMpPF5AN4Avq+Sq3kQXdZBEzifSbX1Rn5108X4+Mi9qghvXapOQ== ankitbhardwaj11411920@gmail.com"),
		})
		if err!=nil{
			return err
		}

		jenkinsServer, err := ec2.NewInstance(ctx, "jenkins-server", &ec2. InstanceArgs{
			InstanceType: pulumi. String("t3.micro"),
			VpcSecurityGroupIds: pulumi. StringArray{sg.ID()},
			Ami: pulumi.String("ami-0f0ec0d37d04440e3"),
			KeyName: kp.KeyName,
		})
		if err!=nil{
			return err
		}

		fmt.Println(jenkinsServer.PublicIp)
		fmt.Println(jenkinsServer.PublicDns)
		ctx. Export("publicIp", jenkinsServer.PublicIp)
		ctx. Export("publicHostName", jenkinsServer.PublicDns)

		
		return nil

	})
		
}