# SSH Keys
Today we will learn about ssh keys both public and private ones

### Explanation
SSH (Secure Shell) keys provide a secure way to authenticate and login to remove servers without using insecure passwords. It relies on public key cryptography
to ensure that your connection is encrypted during transit.
#### Key Generation:
##### You generate a pair of cryptographic keys on your local machine:
Private Key: Kept secret on your client (e.g., your computer). This is like a digital signature tool.
Public Key: Shared with the remote server (e.g., added to ~/.ssh/authorized_keys on the server). This is like a lock that only the private key can open.
### Technical:
#### Authentication Process:
When you attempt to connect (e.g., via ssh user@server), the server sends a challenge (random data) encrypted with your public key.
Your client uses the private key to decrypt/sign the challenge and sends it back.
The server verifies the signature using the public key. If it matches, access is granted—no password needed.
This process is asymmetric: The public key can't be used to get the private key, ensuring security even if the public key is compromised.

#### Security Benefits:
Keys are ***WAY*** stronger than passwords (longer, extremely hard to brute-force), support passphrases for added protection, and prevent man-in-the-middle attacks when properly set up (e.g., via SSH host keys).

## How to setup one properly:

### Client (Key Generation)
Generate your SSH key pair on your local machine. This creates a private key (keep secret) and a public key (share with servers).

#### Linux/macOS
1. Open a terminal.
2. Run the following command to generate an RSA key pair (replace `your_email@example.com` with your email for identification):
```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```
   - Press Enter to accept the default file location (`~/.ssh/id_rsa`).
   - Optionally, enter a passphrase for added security (recommended).
3. Your keys are generated: `~/.ssh/id_rsa` (private) and `~/.ssh/id_rsa.pub` (public).

#### Windows
1. Install OpenSSH if not already available (via Windows Settings > Apps > Optional features > Add a feature > OpenSSH Client).
2. Open PowerShell or Command Prompt.
3. Run the same command as Linux:
```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```
   - Press Enter for defaults.
   - Add a passphrase if wanted, this isnt needed.
4. Keys are saved in `C:\Users\YourUsername\.ssh\id_rsa` and `id_rsa.pub`.

**Note:** Always use a strong passphrase and back up your private key securely.

### Server
#### Linux
To set up SSH key authentication on a Linux server:

1. Install OpenSSH server if not installed:
<br>
For Ubuntu/Debian:

```bash
sudo apt update && sudo apt install openssh-server
```
   Or for CentOS/RHEL:

```bash
sudo yum install openssh-server
``` 
Or for Alpine:
```
sudo apk add openssh-server
```
2. Ensure SSH service is running:
` sudo systemctl start ssh && sudo systemctl enable ssh `
3. On the server, create the '.ssh' directory for the user (replace 'username' with the target user):
```
sudo mkdir -p /home/username/.ssh
sudo chown username:username /home/username/.ssh
sudo chmod 700 /home/username/.ssh
```
4. Add your public key to the `authorized_keys` file:
   - Copy your `id_rsa.pub` content (from client).
   - Append it to `/home/username/.ssh/authorized_keys`:
```bash
echo "your_public_key_here" >> /home/username/.ssh/authorized_keys
```
5. Set correct permissions:
```bash
sudo chown username:username /home/username/.ssh/authorized_keys
sudo chmod 600 /home/username/.ssh/authorized_keys
```
6. Restart SSH service:
```
sudo systemctl restart ssh
```

Now, you can SSH from your client without a password: `ssh username@server_ip`.

#### Windows
To set up SSH on a Windows server using OpenSSH:

1. Install OpenSSH Server via Windows Settings > Apps > Optional features > Add a feature > OpenSSH Server.
2. Start the service:
   - Open PowerShell as Administrator.
```
Start-Service sshd
Set-Service -Name sshd -StartupType 'Automatic'
```
3. Create `.ssh` directory for the user (replace `Username`):
```
New-Item -ItemType Directory -Path "C:\Users\Username\.ssh" -Force
```
4. Add your public key to `authorized_keys`:
   - Copy `id_rsa.pub` content.
   - Create/edit `C:\Users\Username\.ssh\authorized_keys` and paste the key.
5. Set permissions (use icacls for security):
```
icacls "C:\Users\Username\.ssh" /inheritance:r /grant "Username:(OI)(CI)F" /grant "SYSTEM:(OI)(CI)F" /grant "Administrators:(OI)(CI)F"
icacls "C:\Users\Username\.ssh\authorized_keys" /inheritance:r /grant "Username:F" /grant "SYSTEM:F" /grant "Administrators:F"
```
6. Restart SSH:
```
Restart-Service sshd
```

Connect from client: `ssh username@server_ip`.

**Note:** For both, verify host keys on first connect to prevent MITM attacks.

## Samples:

### Example Public Key (id_rsa.pub)
```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7VbLjBCqz0xhBqQX8Ak... your_email@example.com
```
(This is a shortened example; actual keys are much longer.)

### Example Private Key (id_rsa) - DO NOT SHARE
```
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABlwAAAAdzc2...
... (redacted for security; never expose private keys)
-----END OPENSSH PRIVATE KEY-----
```
**Warning:** Private keys must remain secret. Use tools like `termius` to manage them securely.

### Example authorized_keys Entry
```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7VbLjBCqz0xhBqQX8Ak... your_email@example.com
```
(This file contains one or more public keys, one per line.)