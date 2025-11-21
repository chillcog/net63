# Dark Web Introduction

Today we will learn about the dark web, what it is, and how to access it safely.

### Explanation
The dark web is a part of the internet that is not indexed by search engines and requires special software to access. It is often associated with anonymity and illegal activities, but it also serves legitimate purposes like privacy advocacy and secure communication.

#### Key Concepts:
##### Deep Web vs Dark Web: Deep web is any content not indexed; dark web uses overlay networks like Tor.
##### Tor: The Onion Router, a network that routes traffic through multiple nodes to hide IP addresses.
##### .onion Sites: Websites with .onion domain, only accessible via Tor.

### Technical:
#### How Tor Works:
Traffic is encrypted and routed through volunteer nodes (entry, middle, exit). Each layer peels off encryption, ensuring no single node knows the full path.

#### Safety Benefits:
Provides anonymity, protects against surveillance, but also harbors risks like malware and scams.

## How to setup one properly:

### General Best Practices
1. Use dedicated tools; never use regular browser.
2. Enable security settings in Tor Browser.
3. Avoid logging in with personal accounts.
4. Be aware of exit node vulnerabilities.

### Client (Accessing the Dark Web)
Set up Tor for safe access.

#### Linux/macOS
1. Download Tor Browser from torproject.org.
2. Extract and run: `./start-tor-browser.desktop`
3. Connect to Tor network.
4. For advanced: Use Tails OS (live USB) for amnesia.

#### Windows
1. Download Tor Browser installer.
2. Run and follow setup.
3. Use bridges if Tor is blocked: In Tor Browser, Settings > Tor > Use a bridge.

**Note:** Update Tor regularly; use the latest version.

### Server (Hosting .onion Sites)
For advanced users hosting services.

#### Linux
1. Install Tor: `sudo apt install tor`
2. Configure /etc/tor/torrc with HiddenServiceDir and Port.
3. Restart Tor: `sudo systemctl restart tor`
4. Get .onion address from hostname file.

#### Windows
1. Use Tor expert bundle.
2. Similar config in torrc.

**Note:** Hosting requires technical knowledge; use for legal purposes only.

## Samples:

### Example .onion URL
```
http://dreadytofatroptsdj6io7l3xptbet6onoyno2yv7jicoxknyazubrad.onion/
```
(This is a placeholder; real URLs are long and random.)

### Example Torrc Config for Hidden Service
```
HiddenServiceDir /var/lib/tor/hidden_service/
HiddenServicePort 80 127.0.0.1:80
```

### Example Tor Browser Security Level
```
Safest: Disables JavaScript, images, etc.
```

## Recommended
- Use Tor Browser exclusively.
- Learn about OPSEC (operational security).
- Avoid illegal activities.
- Use Tails for maximum anonymity.
- Stay informed on Tor updates and vulnerabilities.
