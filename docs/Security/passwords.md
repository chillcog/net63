# Passwords

Today we will learn about passwords, their importance, and best practices for creating and managing them.

### Explanation
Passwords are the most common form of authentication used to protect access to accounts, devices, and systems. A strong password acts as a barrier against unauthorized access, but weak or poorly managed passwords can lead to security breaches.

#### Key Concepts:
##### Password Strength: Determined by length, complexity (mix of uppercase, lowercase, numbers, symbols), and unpredictability.
##### Password Managers: Tools that securely store and generate complex passwords, reducing the need to remember them.
##### Multi-Factor Authentication (MFA): Adds an extra layer of security beyond passwords, such as codes sent to a phone or biometric verification.

### Technical:
#### How Passwords Work:
Passwords are hashed (transformed into a fixed-size string) using algorithms like bcrypt or Argon2 before storage. During login, the entered password is hashed and compared to the stored hash. This ensures passwords aren't stored in plain text.

#### Security Benefits:
Strong passwords prevent brute-force attacks, dictionary attacks, and credential stuffing. Combined with MFA, they significantly reduce the risk of account compromise.

## How to setup one properly:

### General Best Practices
1. Use a minimum of 12-16 characters.
2. Include a mix of character types.
3. Avoid common words, personal info, or patterns.
4. Use unique passwords for each account.
5. Change passwords periodically or after breaches.

### Client (Password Management)
Manage passwords securely on your local machine or device.

#### Linux/macOS
1. Install a password manager like KeePassXC or Bitwarden.
   - Ubuntu/Debian: `sudo apt install keepassxc`
   - macOS: `brew install keepassxc` (if Homebrew installed)
2. Generate a strong master password for the manager.
3. Create entries for each account with auto-generated passwords.
4. Enable auto-fill in browsers or use browser extensions.

#### Windows
1. Install a password manager like Bitwarden or KeePass.
   - Download from official sites.
2. Set up a master password.
3. Use the manager to generate and store passwords.
4. Integrate with Windows Credential Manager or browser extensions.

**Note:** Never reuse passwords. Use MFA wherever possible.

### Server/Account Setup
For systems requiring password authentication:

#### Linux
1. Set strong passwords for user accounts:
   - Use `passwd` to change passwords.
2. Enforce password policies with PAM:
   - Edit `/etc/security/pwquality.conf` for complexity rules.
3. Disable root login over SSH; use sudo.

#### Windows
1. Use Group Policy to enforce password policies:
   - Go to Computer Configuration > Windows Settings > Security Settings > Account Policies > Password Policy.
   - Set minimum length, complexity, etc.
2. Enable password history and lockout policies.

**Note:** For web accounts, enable MFA and use password managers.

## Samples:

### Example Strong Password
```
Tr!ckyP@ssw0rd2023!
```
(This is an example; generate unique ones.)

### Example Password Hash (bcrypt)
```
$2b$10$abcdefghijklmnopqrstuvwx.yzABCDEFGHIJKLMNOPQRSTUVWXYZ123456
```
(Hashed passwords are irreversible.)

### Example Password Policy Configuration (Linux /etc/security/pwquality.conf)
```
minlen = 12
dcredit = -1
ucredit = -1
lcredit = -1
ocredit = -1
```
(Requires at least 12 chars, one digit, uppercase, lowercase, special char.)

## Recommended
- Use password managers like Bitwarden or LastPass.
- Enable MFA on all accounts.
- Regularly audit and update passwords.
- Avoid writing passwords down; use secure storage.
