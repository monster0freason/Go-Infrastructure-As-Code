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
			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCwqLUy2W+7SiKPpUW3GnIf+E4EkMLZ0I33lDooMUq4YQ3Muh1vbBlp3TjdkZO46ZsZXoKqp2wmCE9uzmKX7YzadBYOJVJZthuK95cnYK+EI7nBQezrMx5VdyzmwSl3+gDdmr0camuB4PWNYSjAO3C4xdcUob9H6Y2MzSMoS8pPe2MaGlrd5I8nqwBdvLAM1GroBKjX420O303lYwSsY5TGbYzjuRTNSl74PsHKqAOricAvMLa5V1gYbrQmoptvUQ6/jl3mg6fDvs6S/5c9dM+07bm245WBOzxBNNpL9ewsLIWlHoLKMspzaGJ5JMsPjyDvlhogoX2P1IEfmaSjh5+XluPqMsbzn7s+bcmRkhcdXfov0LntkXKKeRdI/TuEwwfcChHev3YLQ/blQ3qPeuShoCD9hb5jiD2EkN+9MHYG0HvKMY8GrPIfZ9JvPl5Aqztwe6LPVrT9monYqVTRSCU6Dj/gQc/YZf51dwo5MEP4iqMNP7GVLmSBNnqSAPohbM0= ankit@DESKTOP-GR9E2KS"),
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