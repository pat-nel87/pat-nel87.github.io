---
title: 'Building My Portfolio with Hugo and Go'
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

I've recently launched my personal portfolio site, powered by Hugo—a fast and powerful static site generator written in Go. In this post, I'll walk you through the key steps I took, sharing insights and tips from my experience.

## Why Hugo and Go? 🚀

I chose Hugo primarily for its speed, simplicity, and powerful templating capabilities. Hugo uses Go under the hood, which ensures rapid site generation and deployment. Plus, as someone deeply involved in DevOps, automation, and infrastructure-as-code, Hugo aligns perfectly with my workflow.

## Setting Up Hugo 📦

Getting started with Hugo is straightforward:

```bash
# Install Hugo
brew install hugo

# Create a new site
hugo new site my-portfolio
```

## Choosing a Theme 🎨

I selected the [Terminal Hugo theme](https://themes.gohugo.io/themes/hugo-theme-terminal/) for its minimalistic yet impactful presentation, matching my personality and professional style:

```bash
git submodule add https://github.com/panr/hugo-theme-terminal.git themes/terminal
```

Then, I customized the `hugo.yaml` to fine-tune the appearance.

## Customizing My Site 🛠️

One of the benefits of Hugo is its flexibility. I leveraged Hugo’s powerful templating system:

- Added custom shortcodes for SVG icons:

{{< svg logo="go" >}}

- Highlighted code snippets clearly:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Hugo!")
}
```

## Automating Deployment 🤖

Deployments are automated via GitHub Actions, reflecting my DevOps expertise. Every push triggers a workflow that:

- Runs `hugo` to build the site
- Publishes content automatically to GitHub Pages

This simplifies publishing updates tremendously.

## Lessons Learned 📚

- **Rapid Prototyping**: Hugo’s speed allowed quick iterations.
- **Customization Power**: Shortcodes and templating make dynamic content easy.
- **Integration with DevOps Tools**: Hugo integrates seamlessly with CI/CD pipelines.

## What's Next? 🔮

Looking forward, I plan to enhance my site further:

- Adding interactive JavaScript components.
- Exploring deeper integrations with cloud-native services.

## Let's Connect! 🌐

I'm always excited to discuss Hugo, DevOps, or technology in general. Reach out via my [contact page](/contact), and let’s connect!

Happy coding! 🎉

