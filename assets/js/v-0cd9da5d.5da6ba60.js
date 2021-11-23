"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[910],{1:(s,n,e)=>{e.r(n),e.d(n,{data:()=>a});const a={key:"v-0cd9da5d",path:"/intro/getting-started.html",title:"Quickstart",lang:"en-US",frontmatter:{},excerpt:"",headers:[],filePathRelative:"intro/getting-started.md",git:{updatedTime:1637704984e3,contributors:[{name:"Adam Tucker",email:"adamleetucker@outlook.com",commits:1}]}}},9322:(s,n,e)=>{e.r(n),e.d(n,{default:()=>o});const a=(0,e(6252).uE)('<h1 id="quickstart" tabindex="-1"><a class="header-anchor" href="#quickstart" aria-hidden="true">#</a> Quickstart</h1><p><em>(Note: This repository is under active development. Architecture and implementation may change without documentation)</em></p><p>This is what you&#39;d use to get a node up and running, fast. It assumes that it is starting on a blank ubuntu machine. It eschews a systemd unit, allowing automation to be up to the user. It assumes that installing Go is in-scope since Ubuntu&#39;s repositories aren&#39;t up to date and you&#39;ll be needing go to use osmosis. It handles the Go environment variables because those are a common pain point.</p><p><strong>Install go</strong></p><div class="language-bash ext-sh line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#ADBAC7;">wget -q -O - https://git.io/vQhTU </span><span style="color:#F47067;">|</span><span style="color:#ADBAC7;"> bash -s -- --version 1.17.2</span></span>\n<span class="line"></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Then exit and re-enter your shell.</p><p><strong>Install Osmosis and check that it is on $PATH</strong></p><div class="language-bash ext-sh line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#ADBAC7;">git clone https://github.com/osmosis-labs/osmosis</span></span>\n<span class="line"><span style="color:#6CB6FF;">cd</span><span style="color:#ADBAC7;"> osmosis</span></span>\n<span class="line"><span style="color:#ADBAC7;">git checkout v3.1.0</span></span>\n<span class="line"><span style="color:#ADBAC7;">make install</span></span>\n<span class="line"><span style="color:#ADBAC7;">which osmosisd</span></span>\n<span class="line"></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br></div></div><p><strong>Launch Osmosis</strong></p><div class="language-bash ext-sh line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#ADBAC7;">osmosisd init yourmonikerhere</span></span>\n<span class="line"><span style="color:#ADBAC7;">wget -O </span><span style="color:#F47067;">~</span><span style="color:#ADBAC7;">/.osmosisd/config/genesis.json https://github.com/osmosis-labs/networks/raw/main/osmosis-1/genesis.json</span></span>\n<span class="line"><span style="color:#ADBAC7;">osmosisd start</span></span>\n<span class="line"></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br></div></div><p>More Nodes ==&gt; More Network</p><p>More Network ==&gt; Faster Sync</p><p>Faster Sync ==&gt; Less Developer Friction</p><p>Less Developer Friction ==&gt; More Osmosis</p><p>Thank you for supporting a healthy blockchain network and community by running an Osmosis node!</p>',15),t={},o=(0,e(3744).Z)(t,[["render",function(s,n){return a}]])}}]);