---
title: 'Building AI-Powered Diagnostic Tools with MCP'
date: 2026-04-15T10:00:00+00:00
draft: false
params:
  slug: mcp-diagnostic-tools
layout: "post"
tags: ["mcp", "go", "typescript", "kubernetes", "postgresql", "devops", "ai"]
authors: ["Patrick Nelson"]
---

---

## Introduction

If you've spent any time troubleshooting a misbehaving Kubernetes cluster or chasing down slow queries in Postgres,
you know the drill — open six terminal tabs, run a dozen commands, mentally stitch together the results,
and hope you don't forget what you found in tab three while you're reading tab five.

I've been building tools to fix that problem. Not by replacing the investigation,
but by giving AI assistants the ability to do the legwork for you.

Enter the **Model Context Protocol (MCP)** — a standard for exposing structured tool capabilities to AI hosts
like Claude Desktop, Claude Code, and VS Code Copilot. I've been building MCP-powered diagnostic servers that
let you ask questions like *"why is this pod crashlooping?"* or *"which queries are killing my database?"*
and get real, structured answers pulled directly from live systems.

In this post, I'll walk through two tools I've been working on:
[**kube-doctor-mcp**](https://github.com/pat-nel87/kube-doctor-mcp) and [**pg-doctor**](https://github.com/pat-nel87/pg-doctor).

## What is MCP, and Why Should You Care?

MCP (Model Context Protocol) is essentially a standard interface between AI assistants and external tools.
Think of it like a USB port for AI — plug in a tool, and the AI can use it without needing custom integrations
for every host.

Before MCP, if you wanted an AI to interact with your infrastructure, you'd end up writing bespoke plugins
for each platform. Now you write one MCP server, and it works everywhere — Claude Code, VS Code Copilot Agent Mode,
Claude Desktop, you name it.

The key constraints I design around:
- **Read-only** — These tools inspect. They never mutate. Sleep well at night.
- **Structured output** — Human-readable text with severity tags, not raw JSON dumps.
- **Diagrams** — Mermaid diagrams for topology, dependencies, and security audits rendered inline.

## kube-doctor-mcp: Kubernetes Cluster Diagnostics

[**kube-doctor-mcp**](https://github.com/pat-nel87/kube-doctor-mcp) is a Kubernetes diagnostics MCP server written in Go.
It connects directly to the Kubernetes API via kubeconfig — no `kubectl` dependency required — and exposes
**48 read-only tools** across the full surface area of a cluster.

### What Can It Do?

Instead of rattling off all 48 tools, here's how they break down by category:

| Category | What You Get |
|----------|-------------|
| **Cluster & Nodes** | Context listing, namespace discovery, node details and metrics |
| **Workloads** | Pods, Deployments, StatefulSets, DaemonSets, Jobs — with detail views and logs |
| **Networking** | Services, Ingresses, Endpoints, Network Policies, connectivity analysis |
| **Storage** | PVCs, PVs — status and binding info |
| **Metrics & Resources** | Node/pod metrics, top resource consumers, resource allocation analysis |
| **Security** | Pod security analysis, RBAC bindings, namespace security audits |
| **FluxCD GitOps** | 8 dedicated tools for Kustomizations, HelmReleases, sources, image policies, and Flux system health |
| **Doctor Mode** | Composite diagnostics — `diagnose_pod`, `diagnose_namespace`, `diagnose_cluster` with remediation suggestions |

### The Doctor Tools

The individual inspection tools are useful, but the *doctor* tools are where the real value lives.
`diagnose_pod` doesn't just tell you a pod is unhealthy — it checks resource limits, restart counts,
image pull status, readiness probes, and recent events, then synthesizes findings tagged by severity:

```
[CRITICAL] Pod api-server-7b4d9 in CrashLoopBackOff — 47 restarts
[WARNING]  No memory limit set — risk of OOM on shared node
[INFO]     Last successful readiness probe: 3h ago
```

`diagnose_cluster` goes even bigger — it's a full cluster health check in one call.
Think of it as your morning standup, but the AI reads every namespace before you've finished your coffee.

### FluxCD Support

If you're running GitOps with FluxCD (and you should be), there are 8 dedicated Flux tools
for inspecting Kustomizations, HelmReleases, sources, and image policies.
`diagnose_flux_system` gives you a full system health check with a Mermaid topology diagram.
No more `flux get all -A` and squinting at terminal output.

### Mermaid Diagrams

Several tools generate Mermaid diagrams that render inline in supported AI hosts:
- Network connectivity maps from `analyze_pod_connectivity`
- Security audit visualizations from `audit_namespace_security`
- Resource allocation overviews from `analyze_resource_allocation`
- FluxCD topology from `get_flux_resource_tree` and `diagnose_flux_system`
- Workload dependency graphs from `get_workload_dependencies`

### Tech Stack

- Written in **Go** — because Kubernetes tooling belongs in Go
- Uses `client-go` directly — no shelling out to kubectl
- `controller-runtime` for FluxCD CRD support
- MCP stdio transport — single binary, no daemon, no dependencies
- 30-second timeout on all API calls — won't hang your AI session

## pg-doctor: Database Diagnostics in Your Editor

[**pg-doctor**](https://github.com/pat-nel87/pg-doctor) takes the same diagnostic philosophy and applies it to PostgreSQL.
It's a VS Code extension that exposes **14 Language Model Tools** to GitHub Copilot Chat,
so you can interrogate your databases without leaving your editor.

### What Can It Do?

**10 Diagnostic Tools:**

| Tool | What It Checks |
|------|---------------|
| **Health Summary** | Cache hit ratio, connection usage, checkpoint stats, XID age, deadlocks |
| **Slow Queries** | Top queries from `pg_stat_statements`, sortable by time or call count |
| **Index Health** | Unused, duplicate, and bloated indexes |
| **Lock Check** | Blocking chains with wait durations |
| **Vacuum Status** | Dead tuples, autovacuum timing, XID wraparound risk |
| **Connection Stats** | Pool breakdown, utilization vs `max_connections`, idle-in-transaction sessions |
| **Replication Status** | Replica lag, slot health, retained WAL |
| **Table Stats** | Sequential vs index scan ratios, dead tuples, cache hits per table |
| **Bloat Check** | Estimated table and index bloat with wasted bytes |
| **Triage** | Runs ALL checks at once, returns prioritized findings by severity |

**4 Mermaid Diagram Tools:**

| Tool | What It Renders |
|------|----------------|
| **Schema Diagram** | ER diagram with tables, columns, PKs, FKs, and relationships |
| **Lock Diagram** | Flowchart of blocking chains with root blockers highlighted |
| **Connection Diagram** | Pie chart of connection pool state |
| **Table Size Diagram** | Bar chart of top tables by size with data/index/bloat breakdown |

### The Triage Tool

The `#pgTriage` tool is the one I reach for first. It runs all 10 diagnostic checks in one shot
and returns a single prioritized report sorted by severity. Instead of running each check individually
and piecing together the story yourself, you get the full picture in one pass.

It's like having a DBA on call who already checked everything before you asked.

### Design Decisions

- **Read-only connections** — `default_transaction_read_only=on` is set at the connection level. No accidents.
- **10-second statement timeout** — Won't tank your production database with a runaway diagnostic query.
- **Multi-database support** — Configure multiple databases via `~/.pg-doctor.json` with aliases.
- **Minimal dependencies** — The only runtime dependency is the `pg` driver. That's it.

### Tech Stack

- Written in **TypeScript**
- VS Code extension using the Language Model Tools API
- Works in GitHub Copilot Chat Agent Mode
- Single runtime dependency (`pg`)
- esbuild for fast compilation

## The Pattern: Read-Only Diagnostic Servers

Both tools follow the same architectural pattern that I think works really well for infrastructure diagnostics:

1. **Read-only, always.** These tools never modify state. Every connection, every API call is strictly read-only with timeouts.
2. **Structured human-readable output.** Not JSON, not raw query results — formatted text with severity tags that an AI can interpret and a human can read.
3. **Composite diagnosis tools.** Individual checks are useful, but the real value is in tools that run multiple checks and synthesize findings. `diagnose_cluster` and `#pgTriage` are the ones you'll actually use day-to-day.
4. **Inline diagrams.** Mermaid diagrams rendered in the AI chat give you visual context without switching windows.
5. **Zero external dependencies at runtime.** One binary (Go) or one extension (VS Code). No CLI tools, no agents, no sidecars.

## What's Next

These two tools are the first pieces of something bigger — what I've been thinking of as an **AI Observability Mesh**.

The idea is an MCP gateway that acts as a universal diagnostic control plane across Azure resources.
Instead of individual MCP servers that each know about one system, you'd have a single entry point
that can route diagnostic queries across the full stack — from Postgres query performance up through
AKS pod health, FluxCD reconciliation state, and Azure resource metrics.
An AI assistant could triage an incident end-to-end without you needing to context-switch between tools
or even know which layer the problem lives in.

kube-doctor and pg-doctor are proving out the pattern. The next step is the connective tissue.

If you want to try them out or contribute, check them out on GitHub:
- [kube-doctor-mcp](https://github.com/pat-nel87/kube-doctor-mcp)
- [pg-doctor](https://github.com/pat-nel87/pg-doctor)

## Let's Connect

If you're building MCP tools or working on AI-assisted operations, I'd love to hear about it.
Reach out via contacts on my [resume page](/resume), and let's connect.
