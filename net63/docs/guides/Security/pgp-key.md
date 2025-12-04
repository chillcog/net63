# PGP Keys

Today we will learn about PGP keys both public and private ones

### Explanation
PGP (Pretty Good Privacy) keys provide a secure way to encrypt and decrypt messages, files, and emails using public key cryptography. It ensures that your data is encrypted during transit and at rest, and allows for digital signatures to verify authenticity.

#### Key Generation:
##### You generate a pair of cryptographic keys on your local machine:
Private Key: Kept secret on your client (e.g., your computer). This is used for decrypting messages and signing data.
Public Key: Shared with others (e.g., uploaded to keyservers or exchanged directly). This is used by others to encrypt messages for you or verify your signatures.

### Technical:
#### Encryption/Decryption Process:
To encrypt a message: The sender uses the recipient's public key to encrypt the data.
To decrypt: The recipient uses their private key to decrypt it.
For signing: You use your private key to sign a message, and others use your public key to verify the signature.
This process is asymmetric: The public key can't derive the private key, ensuring security.

#### Security Benefits:
PGP provides end-to-end encryption, protects against eavesdropping, allows for non-repudiation through signatures, and is widely used for secure communication (e.g., email with tools like Thunderbird + Enigmail or GPG).

## How to setup one properly:

### Client (Key Generation)
Generate your PGP key pair using GPG (GNU Privacy Guard). This creates a private key (keep secret) and a public key (share with others).

#### Linux/macOS
1. Install GPG if not already available:
   - Ubuntu/Debian: `sudo apt install gnupg`
   - macOS: `brew install gnupg` (if Homebrew is installed)
2. Open a terminal.
3. Run the following command to generate a key pair:
```bash
gpg --full-generate-key
```
   - Choose RSA and RSA (default).
   - Key size: 4096 bits.
   - Expiration: Set to 0 (no expiration) or choose a period.
   - Enter your name, email, and comment.
   - Set a strong passphrase.
4. Your keys are generated. List them with `gpg --list-keys`.

#### Windows
1. Install GPG4Win (https://www.gpg4win.org/) which includes Kleopatra (GUI) and command-line tools.
2. Open Command Prompt or PowerShell.
3. Run the same command as Linux:
```bash
gpg --full-generate-key
```
   - Follow the prompts as above.
4. Alternatively, use Kleopatra GUI to generate keys.

**Note:** Always use a strong passphrase and back up your private key securely (e.g., on an encrypted drive).

### Usage
#### Encrypting a File
```bash
gpg --encrypt --recipient "recipient@example.com" file.txt
```
This creates `file.txt.gpg`.

#### Decrypting a File
```bash
gpg --decrypt file.txt.gpg > decrypted.txt
```

#### Signing a File
```bash
gpg --sign file.txt
```

#### Verifying a Signature
```bash
gpg --verify file.txt.gpg
```

#### Exporting Public Key
```bash
gpg --export --armor your_email@example.com > public_key.asc
```

#### Importing a Public Key
```bash
gpg --import public_key.asc
```

#### Uploading to Keyserver
```bash
gpg --send-keys --keyserver hkps://keys.openpgp.org your_key_id
```

**Note:** For email encryption, integrate with email clients like Thunderbird using Enigmail or use command-line for manual encryption.

## Samples:

### Example Public Key (ASCII Armor)
```
-----BEGIN PGP PUBLIC KEY BLOCK-----
mQENBFYJL4oBCACrKJ... (shortened example)
-----END PGP PUBLIC KEY BLOCK-----
```
(This is a shortened example; actual keys are much longer.)

### Example Private Key - DO NOT SHARE
```
-----BEGIN PGP PRIVATE KEY BLOCK-----
lQPFBFYJL4oBCACrKJ... (redacted for security; never expose private keys)
-----END PGP PRIVATE KEY BLOCK-----
```
**Warning:** Private keys must remain secret. Use password managers or secure storage.

### Example Signed Message
```
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

This is a signed message.
-----BEGIN PGP SIGNATURE-----
iQEcBAEBCAAGBQJWCS+KHSB5b3VyX2VtYWlsQGV4YW1wbGUuY29t...
-----END PGP SIGNATURE-----
```

## Recommended
- Use GPG for all sensitive communications.
- Regularly update and revoke keys if compromised.
- Use keyservers for public key distribution.
- For advanced usage, consider tools like Keybase or ProtonMail for integrated PGP.
