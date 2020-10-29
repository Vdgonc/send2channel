# Send2channel
---
Send notifications to Slack channel.

## How to config:
 * Create file `/etc/send2channel/config.yml`


```yaml
title: "Deploy Application: "
channels:
  - "YOUR INCOMING WEBHOOK"
footer: "Send2channel"
fallback:
  init: "Deploy init"
  success: "Success success"
  fail: "Deploy fail"
message:
  init: |
    Deploy init
  success: |
    Deploy success :)
  fail: |
    Deploy failed ;(
color:
  init: "fffb01"
  success: "#2dfc34"
  fail: "#fc0505"

```

 * Move binary to `/usr/local/bin`
 
```bash
 mv send2channel /usr/local/bin
```
## How to use

```bash
send2channel -appname="Main API" -status="init"
send2channel -appname="Main API" -status="success"
send2channel -appname="Main API" -status="fail"
```

### One tip :)
Create an alias to send2channel

```bash
alias s2c="send2channel"
```
