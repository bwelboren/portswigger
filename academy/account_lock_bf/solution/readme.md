```

# What this code does.

Loops every username 5 times to trigger an account lock error which contains a logic flaw which is vulnerable to a bruteforce attack.

A normal bruteforce timeout looks like: "You have made too many incorrect login attempts. Please try again in 1 minute(s)."

One account returns a different error message: "You have made too many incorrect login attempts." which is the account lock error.

We use that account to continue our bruteforce and look in the body's response if it contains "Invalid password" if not it means we logged in and it returns USERNAME:PASSWORD




```
