---
title: 'A Blog about a Blog'
date: 2025-03-22T10:00:00+00:00
draft: false
params:
  slug: building-hugo-portfolio
layout: "post"
tags: ["hugo", "go", "devops", "portfolio", "tutorial"]
authors: ["Patrick Nelson"]
---

---

## Introduction 👋

As a Software Engineer, You can’t write a blog without first building a blog.
So here I am, writing a blog about building a blog...
I recently launched my personal site, powered by Hugo, Go, and GitHub Actions.

In this post, I’ll walk you through how I set everything up,
from installing Hugo to picking a theme that didn’t make my site look like it was designed in 1998. 

## Why Hugo and Go? 🚀

I chose Hugo for the same reason I choose black coffee and CLI tools.. it’s fast, simple, and gets the job done.
This thing is so fast, I could rebuild my entire site in the time it takes WordPress to load a settings page.
Hugo fits into my workflow perfectly. No databases, no bloated UI, just markdown, Go, and lightning-fast builds. 

## Setup 📦

### Prerequisites
- I installed [Go](https://go.dev/dl/).
- I installed [Hugo](https://gohugo.io/installation/).
- I reviewed Hugo's [Quick Start Guide](https://gohugo.io/getting-started/quick-start/).
- I reviewed [GitHub Pages Docs](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site).

## Getting Started:

### Building the Blog with Hugo
  
To get started:
- I created a new Git repository to host my GitHub Page following [GitHub Pages Docs](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site).
- I created a local repository and set upstream to my new pages repo's main branch.
- I followed the steps described in [Quick Start Guide](https://gohugo.io/getting-started/quick-start/) to start my hugo project in my new repository.
- I then browsed [Hugo Themes](https://themes.gohugo.io/) to find a template for my page.

#### Picking the Perfect Hugo Theme: So Many Choices, So Little Time 🎨
One of the best things about using Hugo is the amazing selection of open-source themes.
Seriously, there’s a theme for everything.
Want a minimalist, typography-focused blog? There’s a theme for that. 
Need something that looks like it belongs in a cyberpunk movie? There’s a theme for that too.
Want your site to look like a 1999 Geocities page? Well… you could do that, but let’s not.

#### The Hugo Theme Hunt: A Developer's Paradox 🔄
Picking a theme should be simple, right? Wrong. This is the ultimate developer dilemma.. Spend actual time writing blog posts,
or endlessly tweak the theme until it’s perfect? (Spoiler: It’s never perfect.)

I started off browsing the official [Hugo theme gallery](https://themes.gohugo.io/).

I found sleek, modern themes with dark mode (a necessity, obviously ☕).

Themes that had just the right amount of typography finesse to make me feel like a professional writer.

And themes so minimalist that they were almost just a blank page. ("Simplicity!" the theme description proudly declared. "Where’s my content?" I asked.)

**Customizing: Because We Can’t Leave Well Enough Alone 🛠️**
Of course, I could have just picked a theme and been done with it. But where’s the fun in that? 

- ✅ I created some [custom tooling](https://github.com/pat-nel87/pat-nel87.github.io/blob/main/scripts/svg-to-html.go) to import custom svg from [CoreUI](https://github.com/coreui/coreui-icons) as html. (Also in Go 🚀😆) 
- ✅ I configured this tool be driven via GitHub Actions Pipeline. See my pipeline [here](https://github.com/pat-nel87/pat-nel87.github.io/blob/main/.github/workflows/svg-to-html.yaml) 
- ✅ I imported custom svg for use with my template to add the Go Logo as a shortcode. 

{{< svg logo="go" text="Powered by Go" styling="inline-block">}}

This is how all Hugo users end up with a half-stock, half-custom Frankenstein theme that mostly works until the next theme update breaks something.
But hey, that’s the price of perfection.

**The Final Choice: Function Over Perfection 🚀**
After much overthinking, I finally landed on [Terminal Hugo theme](https://themes.gohugo.io/themes/hugo-theme-terminal/).
I decided on this theme because it was:
- ✅ Fast (because slow blogs make me irrationally angry)
- ✅ Clean & modern (because I refuse to use a Comic Sans header)
- ✅ Easily customizable (because I will tweak it again next week)

It was as easy as,

```bash
git submodule add https://github.com/panr/hugo-theme-terminal.git themes/terminal
```
Then, I customized the `hugo.yaml` to fine-tune the appearance.
And just like that, my blog finally had a look that fit my style.
Well… for now.

### Hosting on GitHub Pages & DNS with Cloudflare 🌍
Once I had built my blog with Hugo, the next big question was: Where do I host this thing? I mean, I could spin up a Kubernetes cluster,
set up a load balancer, configure auto-scaling, and make my personal blog more resilient than most production applications…
or I could just use GitHub Pages.

#### Why GitHub Pages? 🤓
GitHub Pages is like the free lunch of static site hosting—simple, reliable, and doesn’t ask for much in return.
No servers to manage, no databases to worry about, just push to a repo and let GitHub Actions do its thing.

Also, let’s be real, nothing feels more DevOps-y than having my blog live in a git repo, with every update triggered by a commit.
CI/CD for blog posts? Absolutely.

#### Cloudflare for DNS: Because Why Not? 🚀
Now, even though GitHub Pages gives me free hosting,
I still wanted my own custom domain because nothing screams "I take myself seriously" like a .net.

Enter Cloudflare—the Swiss Army knife of DNS and web performance.

With Cloudflare, I get:
- ✅ Free, fast, and reliable DNS
- ✅ CDN caching (so my blog loads faster even when I barely have any visitors)
- ✅ Automatic HTTPS (because even personal blogs deserve encryption)

Setting it up was as simple as:

- Pointing my custom domain to Cloudflare.

- Adding DNS records to direct traffic to GitHub Pages.

- Enabling HTTPS and feeling like a cybersecurity expert.

And just like that—boom! My blog is online, and secured.

#### Setting Up an Apex Domain with GitHub Pages & Cloudflare 🛠️

Since I wanted my site to be accessible at `patricknelson-devops.net` and `www.patricknelson-devops.net`,
I had to configure an apex domain (a.k.a. root domain), and `A` and `CNAME` records.
Here’s how I set it up:

**Add an A Record for the apex domain**

In Cloudflare’s DNS settings, I added an A record pointing to GitHub Pages' IP addresses:

```bash
185.199.108.153
185.199.109.153
185.199.110.153
185.199.111.153
(These are GitHub’s official IPs for Pages hosting.)
```

**Set up CNAME Flattening for the apex domain**

Cloudflare supports CNAME Flattening, which lets me point `patricknelson-devops.net` to `pat-nel87.github.io` without violating DNS rules.

I added a `CNAME` record for `www` that points to `pat-nel87.github.io`, so `www` works too!

**Enable HTTPS**

In Cloudflare, I switched SSL/TLS mode to "Full (Strict)" to ensure secure HTTPS connections.

**Tell GitHub Pages to use my custom domain**

In my GitHub Pages settings, I entered `patricknelson-devops.net` as the custom domain and checked “Enforce HTTPS.”

And just like that—boom! 

My blog was live at `patricknelson-devops.net`, loading fast, and secured by HTTPS,

### Automating Deployment: GitHub Actions to the Rescue 🤖🚀
Once I had my Hugo-powered blog looking the way I wanted (after way too much theme tweaking), the next step was getting it deployed automatically.

I turned to GitHub Actions, because why deploy manually when you can write YAML and let the robots 🤖 handle it?

#### Starting with the Stock Workflow 🏗️

Hugo makes it ridiculously easy to deploy to GitHub Pages with GitHub Actions.

I started with their [recommended workflow](https://gohugo.io/host-and-deploy/host-on-github-pages/#step-9) and customized it to fit my repo.
Whenever I push changes to main, the GitHub Actions workflow:

- Checks out the repo

- Installs Go & Hugo

- Builds the site

- Deploys it to GitHub Pages

**The Result: Push, and It’s Live! 🎉**
Now, every time I update my blog and push to main:
- ✅ GitHub Actions builds my site
- ✅ The public/ directory gets pushed to the gh-pages branch
- ✅ GitHub Pages deploys the new version automagically 🤖

No manual builds. No copy-pasting files. No stress. Just instant deployment, powered by YAML, GitHub, and Hugo.

Because if I wanted to deploy my site manually, I’d still be using FTP like it’s 1989.

## Wrapping It Up: A Blog About a Blog About a Blog 🔄
So, let’s take a step back and appreciate the beautiful recursion we’ve just lived through:

I built a blog ✨

To write a blog 📝

About building a blog 🔄

If this isn’t the most developer thing ever, I don’t know what is.

- ✅ I chose Hugo for its speed, simplicity, and my growing obsession with Go.
- ✅ I hosted it on GitHub Pages because free, automated, and Git-powered is my love language.
- ✅ I set up Cloudflare for DNS, security, and occasional gaslighting via cached pages.
- ✅ I automated deployments with GitHub Actions, because if I have to scp files in 2025, I’ve failed as a DevOps engineer.
- ✅ I fell into the theme customization rabbit hole and barely made it out alive.

And at the end of it all? 
I now have a working, fully automated, beautifully over-engineered personal blog that I can… finally use to write about things other than building the blog itself.

If there’s one key takeaway here, it’s this:
🚀 Just start. Hugo makes it easy. GitHub Pages makes it free. 
And writing about your journey—even if it’s just about setting up the blog is half the fun.

That is, until I inevitably rewrite everything in some hot new framework in six months.
But for now? I’m calling this a success.

Welcome to my blog. Next post: something other than a blog about a blog… 😆

## Let's Connect! 🌐

I'm always excited to discuss Go, DevOps, or technology in general. Reach out via contacts on my resume page [resume page](/resume), and let’s connect!

Happy coding! 🎉

