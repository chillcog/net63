# VPNs

Today we will learn about Virtual Private Networks (VPNs), how they work, and how to set them up for secure connections.

### Explanation
VPNs create a secure, encrypted connection over a less secure network, such as the internet. They allow users to send and receive data as if their devices were directly connected to a private network.

#### Key Concepts:
##### Tunneling: Encapsulates data packets within another protocol for secure transmission.
##### Protocols: Common ones include OpenVPN, WireGuard, IKEv2/IPsec.
##### Use Cases: Remote access, bypassing geo-restrictions, enhancing privacy.

### Technical:
#### How VPNs Work:
VPN clients encrypt traffic and send it through a tunnel to a VPN server. The server decrypts and forwards to the destination. This masks IP addresses and protects data from eavesdropping.

#### Security Benefits:
Encrypts internet traffic, hides IP addresses, prevents ISP tracking, and secures public Wi-Fi connections.

## How to setup one properly:

### General Best Practices
1. Choose reputable VPN providers with no-logs policies.
2. Use strong encryption protocols (e.g., WireGuard & OpenVPN).
3. Enable kill switch to prevent leaks.
4. Avoid free VPNs due to privacy risks!

### Client (VPN Setup)
Set up a VPN client on your local machine.

#### Linux/macOS
1. Install OpenVPN or WireGuard.
   - Ubuntu: `sudo apt install openvpn`
   - macOS: Use Tunnelblick for OpenVPN or WireGuard app.
2. Download config files from your VPN provider.
3. Connect: `sudo openvpn config.ovpn` (Linux) or use GUI.
4. For WireGuard: Install `wireguard-tools`, generate keys, and configure interface.

#### Windows
1. Install a VPN client like OpenVPN GUI or WireGuard.
   - Download from official sites.
2. Import config files.
3. Connect via the app's interface.
4. Enable split tunneling if needed.

**Note:** Test for DNS leaks using tools like dnsleaktest.com.

### Server Setup
Set up your own VPN server for custom control.

#### Linux
1. Install OpenVPN:
   - `sudo apt install openvpn easy-rsa`
2. Generate certificates: Use easy-rsa to create CA and keys.
3. Configure server.conf with encryption settings.
4. Start service: `sudo systemctl start openvpn@server`
5. For WireGuard: `sudo apt install wireguard`
   - Generate keys: `wg genkey | tee privatekey | wg pubkey > publickey`
   - Configure /etc/wireguard/wg0.conf

#### Windows
1. Use Windows built-in VPN or third-party software.
2. For server: Install OpenVPN server or use RRAS (Routing and Remote Access Service).
3. Configure certificates and tunnels.
4. Enable firewall rules for VPN traffic.

**Note:** Use strong keys and monitor logs.

## Samples:

### Example OpenVPN Config (client.ovpn)
```
client
dev tun
proto udp
remote vpn.example.com 1194
resolv-retry infinite
nobind
persist-key
persist-tun
ca ca.crt
cert client.crt
key client.key
remote-cert-tls server
cipher AES-256-CBC
auth SHA256
```

### Example WireGuard Config (/etc/wireguard/wg0.conf)
```
[Interface]
PrivateKey = <private_key>
Address = 10.0.0.2/24
DNS = 1.1.1.1

[Peer]
PublicKey = <server_public_key>
Endpoint = vpn.example.com:51820
AllowedIPs = 0.0.0.0/0
```

### Example Connection Command
```
sudo openvpn --config client.ovpn
```
(Connects to VPN server.)

## Recommended
- Use WireGuard for speed and security.
- Providers: Mullvad, ProtonVPN, ExpressVPN.
- Enable 2FA on VPN accounts.
- Regularly update software.
