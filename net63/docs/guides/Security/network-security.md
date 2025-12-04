# Network Security

Today we will learn about network security, threats, and defenses.

### Explanation
Network security protects data during transmission and prevents unauthorized access to networks.

#### Key Concepts:
##### Threats: Man-in-the-middle, DDoS, eavesdropping.
##### Defenses: Encryption, firewalls, IDS/IPS.
##### Protocols: TLS, IPsec.

### Technical:
#### How It Works:
Uses encryption, access controls, monitoring to secure traffic.

#### Security Benefits:
Protects against interception, ensures integrity.

## How to setup one properly:

### General Best Practices
1. Use WPA3 for Wi-Fi.
2. Segment networks.
3. Monitor traffic.

### Client
Configure secure connections.

#### Linux/macOS
1. Use VPN for public Wi-Fi.
2. Enable firewall.

#### Windows
1. Use Windows Firewall.
2. Connect to secure networks.

### Server
Implement network security.

#### Linux
1. Use iptables for rules.
2. Enable TLS.

#### Windows
1. Use Windows Firewall.
2. Configure IPSec.

## Samples:

### Example iptables Rule
```
-A INPUT -p tcp --dport 443 -j ACCEPT
```

### Example TLS Config
```
ssl_protocols TLSv1.2 TLSv1.3;
```

## Recommended
- Use VPNs.
- Implement IDS.
- Regular audits.
