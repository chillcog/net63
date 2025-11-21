# Firewalls

Today we will learn about firewalls, their types, and how to configure them for network security.

### Explanation
Firewalls are security systems that monitor and control incoming and outgoing network traffic based on predetermined security rules. They act as a barrier between trusted internal networks and untrusted external networks.

#### Key Concepts:
##### Hardware Firewalls: Dedicated devices (e.g., routers with firewall capabilities).
##### Software Firewalls: Applications running on hosts (e.g., Windows Firewall).
##### Next-Generation Firewalls (NGFW): Advanced firewalls with deep packet inspection, intrusion prevention, etc.

### Technical:
#### How Firewalls Work:
Firewalls inspect packets at the network layer (stateless) or application layer (stateful). They allow or block traffic based on rules like IP addresses, ports, and protocols.

#### Security Benefits:
Prevent unauthorized access, block malicious traffic, and protect against attacks like DDoS and port scanning.

## How to setup one properly:

### General Best Practices
1. Use the principle of least privilege: Only allow necessary traffic.
2. Regularly update rules and firmware.
3. Enable logging for monitoring.
4. Combine with other security measures like IDS/IPS.

### Client (Host-Based Firewalls)
Configure firewalls on your local machine.

#### Linux/macOS
1. Use UFW (Uncomplicated Firewall) on Linux:
   - Install: `sudo apt install ufw` (Ubuntu).
   - Enable: `sudo ufw enable`
   - Allow SSH: `sudo ufw allow ssh`
   - Deny all incoming: `sudo ufw default deny incoming`
2. On macOS: Use built-in firewall in System Preferences > Security & Privacy > Firewall.

#### Windows
1. Enable Windows Defender Firewall:
   - Go to Control Panel > System and Security > Windows Defender Firewall.
2. Configure rules:
   - Allow specific apps or ports.
   - Block inbound connections by default.
3. Use PowerShell for advanced config:
   - Enable: `Set-NetFirewallProfile -Profile Domain,Public,Private -Enabled True`

**Note:** Test rules to avoid locking yourself out.

### Server Setup
For network-level firewalls:

#### Linux
1. Use iptables or nftables:
   - Example iptables rule: `sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT`
   - Save rules: `sudo apt install iptables-persistent`
2. For servers, use fail2ban to block brute-force attacks.

#### Windows
1. Use Windows Firewall with Advanced Security:
   - Open Windows Defender Firewall > Advanced settings.
2. Create inbound/outbound rules based on ports and IPs.
3. For servers, integrate with Azure Firewall or third-party appliances.

**Note:** For cloud environments, use cloud provider firewalls (e.g., AWS Security Groups).

## Samples:

### Example UFW Status Output
```
Status: active

To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
22/tcp (v6)                ALLOW       Anywhere (v6)
80/tcp (v6)                ALLOW       Anywhere (v6)
443/tcp (v6)               ALLOW       Anywhere (v6)
```

### Example iptables Rule
```
-A INPUT -i eth0 -p tcp --dport 80 -j ACCEPT
```
(Allows HTTP traffic on eth0.)

### Example Windows Firewall Rule (PowerShell)
```
New-NetFirewallRule -DisplayName "Allow SSH" -Direction Inbound -Protocol TCP -LocalPort 22 -Action Allow
```

## Recommended
- Use UFW/iptables on Linux, Windows Firewall on Windows.
- Enable stealth mode to hide ports.
- Monitor logs regularly.
- Consider NGFW for advanced protection.
