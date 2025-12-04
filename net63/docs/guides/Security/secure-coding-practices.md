# Secure Coding Practices

Today we will learn about secure coding practices to prevent vulnerabilities in software development.

### Explanation
Secure coding involves writing code that resists attacks like injection, XSS, and buffer overflows by following best practices and principles.

#### Key Concepts:
##### OWASP Top 10: Common vulnerabilities like SQL injection, broken authentication.
##### Principles: Least privilege, input validation, error handling.
##### Tools: Static analysis (e.g., SonarQube), dependency scanning.

### Technical:
#### Common Vulnerabilities:
- Injection: Unsanitized inputs execute code.
- XSS: Malicious scripts in web pages.
- CSRF: Forged requests.

#### Security Benefits:
Reduces bugs, prevents data breaches, ensures compliance.

## How to setup one properly:

### General Best Practices
1. Validate and sanitize all inputs.
2. Use parameterized queries.
3. Implement proper authentication/authorization.
4. Keep dependencies updated.

### Development Environment
Set up tools for secure coding.

#### Linux/macOS
1. Install SonarQube or ESLint for code analysis.
   - Ubuntu: `sudo apt install sonar-scanner`
2. Use OWASP ZAP for testing.
3. For Python: Use Bandit (`pip install bandit`).
4. Integrate into CI/CD with GitHub Actions.

#### Windows
1. Install Visual Studio with security analyzers.
2. Use OWASP tools or Fortify.
3. For .NET: Enable Code Analysis in project settings.

**Note:** Review code in pull requests.

### Application Setup
Implement security in code.

#### Web Apps (General)
1. Use HTTPS everywhere.
2. Implement CSP (Content Security Policy).
3. Avoid storing secrets in code; use environment variables.

#### Examples
- SQL: Use prepared statements.
- JS: Escape outputs for XSS.

## Samples:

### Example Parameterized Query (Python)
```python
cursor.execute("SELECT * FROM users WHERE id = %s", (user_id,))
```

### Example CSP Header
```
Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline'
```

### Example Input Validation (JS)
```javascript
if (!/^[a-zA-Z0-9]+$/.test(input)) {
  throw new Error('Invalid input');
}
```

## Recommended
- Follow OWASP guidelines.
- Use SAST/DAST tools.
- Train developers on security.
- Conduct regular audits.
