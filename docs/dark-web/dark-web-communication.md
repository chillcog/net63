# Dark Web Communication

Today we will learn about communicating securely on the dark web.

### Explanation
Communication on the dark web must be encrypted to prevent interception. Use tools designed for anonymity.

#### Key Concepts:
##### PGP: Pretty Good Privacy for email encryption.
##### Secure Messengers: Ricochet, Session.
##### Forums: Dread for discussions.

### Technical:
#### End-to-End Encryption:
Messages encrypted on sender's device, decrypted on receiver's.

#### Risks:
Metadata, key compromise.

## How to setup one properly:

### General Best Practices
1. Use PGP for all emails.
2. Verify keys in person or via trusted channels.
3. Avoid unencrypted chats.
4. Use anonymous accounts.

### Client (Communication Tools)
Set up secure channels.

#### Linux/macOS
1. Install GPG: `sudo apt install gnupg`
2. Generate keys: `gpg --full-generate-key`
3. Use Thunderbird with Enigmail for email.
4. For IM: Ricochet or Signal over Tor.

#### Windows
1. Install GPG4Win.
2. Use Kleopatra for key management.
3. ProtonMail for email.
4. Session app for messaging.

**Note:** Exchange keys securely.

### Samples:
Example PGP Key
```
-----BEGIN PGP PUBLIC KEY BLOCK-----
...
-----END PGP PUBLIC KEY BLOCK-----
```

Example Ricochet ID
```
ricochet:abcd...@ricochet.example.onion
```

## Recommended
- Use Ricochet for peer-to-peer.
- Avoid centralized services.
- Practice key hygiene.
- Use OTR for chat encryption.
