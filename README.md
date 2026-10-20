# Utilvo EdgeShield for Apache 2.4, LiteSpeed & cPanel (`v1.1.0`)

Production-ready **Apache HTTP Server 2.4 (`mod_authz_core`, `mod_remoteip`, `mod_setenvif`)** and **`.htaccess` (`mod_rewrite`)** bot firewall configurations powered by **[Utilvo EdgeShield](https://utilvo.com/edge-shield)**.

Blocks rotating-IP headless Chromium clusters from **Tencent Cloud (`AS132203`)**, **Alibaba Cloud (`AS45102`)**, and **Huawei Cloud (`AS136907`)** at the Apache edge while guaranteeing **100% SEO Safe-Harbor** for verified `Googlebot`, `Bingbot`, and `Applebot` crawlers.

- **Interactive Apache & Nginx Rule Customizer:** [https://utilvo.com/edge-shield](https://utilvo.com/edge-shield)
- **Direct Raw Apache 2.4 Config:** [https://utilvo.com/edge-shield/apache-bot-shield.conf](https://utilvo.com/edge-shield/apache-bot-shield.conf)
- **Live JSON Threat Intelligence API:** [https://utilvo.com/api/edge-shield/threats](https://utilvo.com/api/edge-shield/threats)
- **Live Plain-Text CIDR Feed:** [https://utilvo.com/api/edge-shield/threats?format=cidr](https://utilvo.com/api/edge-shield/threats?format=cidr)
- **Security Research Labs:** [https://utilvo.com/labs](https://utilvo.com/labs)

---

## 1. Automated 1-Line Installer (Ubuntu / Debian `apache2` & RHEL / AlmaLinux / cPanel `httpd`)

```bash
curl -sSL https://utilvo.com/edge-shield/install.sh | sudo bash -s -- --server=apache
```

- Automatically detects `apache2ctl` or `httpd`.
- Enables `mod_remoteip`, `mod_setenvif`, `mod_authz_core`, and `mod_rewrite`.
- Installs `/etc/apache2/conf-available/utilvo-edge-shield.conf` (or `/etc/httpd/conf.d/utilvo-edge-shield.conf`).
- Validates syntax via `apache2ctl -t` / `httpd -t` with automatic rollback and sets up daily Threat DB cron synchronization.

---

## 2. Native Apache 2.4 CIDR Radix-Tree Config (`apache-bot-shield.conf`)

Unlike regex-heavy `RewriteCond` chains, Apache 2.4's `mod_authz_core` compiles `Require not ip` CIDR prefixes (`43.128.0.0/10`, `47.74.0.0/15`, etc.) into an in-memory radix tree evaluated in microseconds:

```apache
<IfModule mod_authz_core.c>
    <Location "/">
        <RequireAny>
            Require env utilvo_seo_safe
            <RequireAll>
                Require all granted
                Require not env utilvo_bad_bot_ua
                Require not ip 43.128.0.0/10 101.32.0.0/15 119.28.0.0/16 124.156.0.0/16
                Require not ip 43.119.0.0/16 47.74.0.0/15 47.82.0.0/16 47.88.0.0/15 8.208.0.0/12
                Require not ip 149.232.0.0/16 159.138.0.0/16 114.119.0.0/16 46.250.0.0/16
            </RequireAll>
        </RequireAny>
    </Location>
</IfModule>
```

---

## 3. Shared Hosting / cPanel / LiteSpeed / WordPress (`.htaccess`)

For environments without root access to `/etc/apache2`, drop `.htaccess` into `public_html/.htaccess` to inspect both `CF-Connecting-IP` and `REMOTE_ADDR`.

## License

MIT License — [Utilvo Security Research](https://utilvo.com/edge-shield).
