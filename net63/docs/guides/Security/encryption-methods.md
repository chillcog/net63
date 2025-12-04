# Encryption Methods

Today we will learn about encryption methods, symmetric and asymmetric encryption, and their applications.

### Explanation
Encryption is the process of converting plaintext into ciphertext to protect data confidentiality. It ensures that only authorized parties can access the information.

#### Key Concepts:
##### Symmetric Encryption: Uses the same key for encryption and decryption (e.g., AES).
##### Asymmetric Encryption: Uses a public key for encryption and a private key for decryption (e.g., RSA).
##### Hashing: One-way encryption for integrity (e.g., SHA-256).

### Technical:
#### How Encryption Works:
Symmetric: Key is shared secretly. Fast for large data.
Asymmetric: Public key encrypts, private key decrypts. Slower, used for key exchange.
Hybrid systems combine both for efficiency.

#### Security Benefits:
Protects data at rest and in transit, prevents unauthorized access, and ensures integrity through hashing.

## How to setup one properly:

### General Best Practices
1. Use strong algorithms (AES-256, RSA-4096).
2. Manage keys securely (use HSMs or key managers).
3. Rotate keys regularly.
4. Implement end-to-end encryption where possible.

### Client (Encryption Tools)
Use encryption on your local machine.

#### Linux/macOS
1. Use GPG for asymmetric encryption:
   - Install: `sudo apt install gnupg` (Linux) or `brew install gnupg` (macOS).
   - Generate keys: `gpg --full-generate-key`
   - Encrypt file: `gpg --encrypt --recipient user@example.com file.txt`
2. For symmetric: Use OpenSSL: `openssl enc -aes-256-cbc -salt -in file.txt -out file.enc`
3. For full disk: Use LUKS on Linux or FileVault on macOS.

#### Windows
1. Use BitLocker for disk encryption:
   - Enable in Control Panel > BitLocker Drive Encryption.
2. For files: Use 7-Zip with AES encryption.
3. Install GPG4Win for asymmetric encryption.
   - Generate keys via Kleopatra.
   - Encrypt: Right-click file > Sign/Encrypt.

**Note:** Back up keys securely; loss means data loss.

### Server Setup
Implement encryption on servers.

#### Linux
1. Use LUKS for disk encryption during install or `cryptsetup`.
2. For network: Use TLS/SSL with certificates.
   - Generate certs with Let's Encrypt or OpenSSL.
3. Encrypt databases with tools like pgp_sym_encrypt for PostgreSQL.

#### Windows
1. Enable BitLocker on server drives.
2. Use EFS (Encrypting File System) for files.
3. For web servers: Configure SSL in IIS with certificates.

**Note:** Use certificate authorities for trust.

## Samples:

### Example GPG Encrypted File
```
-----BEGIN PGP MESSAGE-----
hQEMA5rX8QJzQ5QZAQgAk...
-----END PGP MESSAGE-----
```

### Example OpenSSL Symmetric Encryption
```
openssl enc -aes-256-cbc -salt -in plaintext.txt -out encrypted.enc
```
(Decrypt with -d flag.)

### Example RSA Key Pair
Public Key:
```
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...
-----END PUBLIC KEY-----
```
Private Key (DO NOT SHARE):
```
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC...
-----END PRIVATE KEY-----
```

## Recommended
- Use AES for symmetric, RSA/ECC for asymmetric.
- Tools: GPG, OpenSSL, VeraCrypt.
- Enable encryption by default.
- Stay updated on vulnerabilities (e.g., Heartbleed).
