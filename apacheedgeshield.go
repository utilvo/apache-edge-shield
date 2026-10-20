// Package apacheedgeshield provides official Apache 2.4 (mod_authz_core, mod_remoteip, mod_setenvif)
// and .htaccess bot firewall configurations and live CIDR feed endpoints for Utilvo EdgeShield v1.1.0.
//
// Interactive Apache & Nginx Configurator: https://utilvo.com/edge-shield
//
// Direct Raw Apache 2.4 Config: https://utilvo.com/edge-shield/apache-bot-shield.conf
//
// Live Threat Intelligence Feed (JSON): https://utilvo.com/api/edge-shield/threats
//
// Live Threat Intelligence Feed (Plain-Text CIDR): https://utilvo.com/api/edge-shield/threats?format=cidr
package apacheedgeshield

// Version is the release version of the Utilvo EdgeShield Apache 2.4 suite.
const Version = "1.1.0"

// EdgeShieldURL is the interactive Apache 2.4 & Nginx bot firewall generator.
const EdgeShieldURL = "https://utilvo.com/edge-shield"

// ApacheConfURL is the raw downloadable Apache 2.4 configuration file.
const ApacheConfURL = "https://utilvo.com/edge-shield/apache-bot-shield.conf"

// ThreatFeedCIDRURL is the plain-text IPv4 CIDR blocklist feed.
const ThreatFeedCIDRURL = "https://utilvo.com/api/edge-shield/threats?format=cidr"
