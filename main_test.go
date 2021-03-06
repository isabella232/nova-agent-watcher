package main

import (
	"testing"
)

func TestSSH(t *testing.T) {
	test_keys := `ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.loc22
ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.locasdfasd
ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.loc`
	config, err := handleSSH(test_keys, "./scripts")
	if err != nil {
		t.Fatal(err)
	}
	expect := `ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.loc22`
	if config.SSHAuthorizedKeys[0] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[0])
	}
	if config.Users[0].SSHAuthorizedKeys[0] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[0])
	}
	expect = `ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.locasdfasd`
	if config.SSHAuthorizedKeys[1] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[0])
	}
	if config.Users[0].SSHAuthorizedKeys[1] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[0])
	}
	expect = `ssh-dss AAAAB3NzaC1kc3MAAACBANxpzIbTzKTeBRaOIdUxwwGwvDasTfU/PonhbNIuhYjc+xFGvBRTumox2F+luVAKKs4WdvA4nJXaY1OFi6DZftk5Bp4E2JaSzp8ulAzHsMexDdv6LGHGEJj/qdHAL1vHk2K89PpwRFSRZI8XRBLjvkr4ZgBKLG5ZILXPJEPP2j3lAAAAFQCtxoTnV8wy0c4grcGrQ+1sCsD7WQAAAIAqZsW2GviMe1RQrbZT0xAZmI64XRPrnLsoLxycHWlS7r6uUln2c6Ae2MB/YF0d4Kd1XZii9GHj7rrypqEo7MW8uSabhu70nmu1J8m2O3Dsr+4oJLeat9vwPsJV92IKO0jQwjKnAOHOiB9JKGeCw+NfXfogbti9/q38Q6XcS+SI5wAAAIEA1803Y2h+tOOpZXAsNIwl9mRfExWzLQ3L7knwJdznQu/6SW1H/1oyoYLebuk187Qj2UFI5qQ6AZNc49DvohWx0Cg6ABcyubNyoaCjZKWIdxVnItHWNbLe//+tyTu0I2eQwJOORsEPK5gMpf599C7wXQ//DzZOWbTWiHEX52gCTmk= polvi@Alexs-MacBook-Air.loc`
	if config.SSHAuthorizedKeys[2] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[2])
	}
	if config.Users[0].SSHAuthorizedKeys[2] != expect {
		t.Fatalf("parsing failed for %v, got %v", expect, config.SSHAuthorizedKeys[0])
	}
	return
}
func TestNet(t *testing.T) {
	gentoo_net := `# Creator: NOVA AGENT:
        # This file was autogenerated at 2014-03-11 20:44:51.442429 by /usr/share/nova-agent/nova-agent.py.
        # While it can still be managed manually, definitely not recommended.
        

modules=( "ifconfig" )


# Label private
config_eth47=(
  "10.209.171.177 netmask 255.255.224.0"
)

routes_eth1=(
  "10.208.0.0 netmask 255.240.0.0 gw 10.209.160.1"
  "10.176.0.0 netmask 255.240.0.0 gw 10.209.160.1"
)

dns_servers_eth1=(
 "173.203.4.9"
 "173.203.4.8"
)

# Label public
config_eth0=(
  "23.253.212.214 netmask 255.255.255.0"
  "2001:4801:7824:0104:fce4:0b49:ff11:6246/64"
)

routes_eth0=(
  "default via 23.253.212.1"
  "default via fe80::def"
)

dns_servers_eth0=(
 "173.203.4.9"
 "173.203.4.8"
)
`
	config, err := handleNet(gentoo_net, "./scripts")
	if err != nil {
		t.Fatal(err)
	}
	if len(config.Coreos.Units) != 3 {
		t.Fatalf("got %d units, expected %d", len(config.Coreos.Units), 3)
	}
	if config.Coreos.Units[0].Name != "50-eth47.network" {
		t.Fatalf("got %v, expected %v", config.Coreos.Units[0].Name, "50-eth47.network")
	}
	if config.Coreos.Units[1].Name != "50-eth1.network" {
		t.Fatalf("got %v, expected %v", config.Coreos.Units[1].Name, "50-eth1.network")
	}
	if config.Coreos.Units[2].Name != "50-eth0.network" {
		t.Fatalf("got %v, expected %v", config.Coreos.Units[2].Name, "50-eth0.network")
	}
}
func TestHostname(t *testing.T) {
	hostname := `# Set to the hostname of this machine
# Creator: NOVA AGENT:
        # This file was autogenerated at 2014-03-14 18:36:09.168884 by /usr/share/nova-agent/nova-agent.py.
        # While it can still be managed manually, definitely not recommended.
        
HOSTNAME="polvi-test"
`
	config, err := handleHostname(hostname, "./scripts")
	if err != nil {
		t.Fatal(err)
	}
	if config.Hostname != "polvi-test" {
		t.Fatalf("got %v, expected %v", config.Hostname, "polvi-test")
	}

}
func TestShadow(t *testing.T) {
	shadow := `root:$1$NyBnu0Gl$GBoj9u6lx3R8nyqHuxPwz/:15839:0:::::
`
	config, err := handleShadow(shadow, "./scripts")
	if err != nil {
		t.Fatal(err)
	}
	if config.Users[0].Name != "root" {
		t.Fatalf("failed to set root user, got", config.Users[0].Name)
	}
	if config.Users[0].PasswordHash != "$1$NyBnu0Gl$GBoj9u6lx3R8nyqHuxPwz/" {
		t.Fatalf("failed to parse root user password, got", config.Users[0].PasswordHash)
	}
	if config.Users[1].Name != "core" {
		t.Fatalf("failed to set core user, got", config.Users[1].Name)
	}
	if config.Users[1].PasswordHash != "$1$NyBnu0Gl$GBoj9u6lx3R8nyqHuxPwz/" {
		t.Fatalf("failed to parse core user password, got", config.Users[1].PasswordHash)
	}

}
